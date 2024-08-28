package service

import (
	"context"
	"fmt"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/data"
	"platform-data/domain"
	"strings"
)

type IInsertServer interface {
	Insert(ctx context.Context, req *data.InsertReq) (bool, error)
}

type InsertServer struct {
	IInsertServer
	mapper domain.IEsMapper
}

func NewInsertServer() *InsertServer {
	return &InsertServer{
		mapper: domain.NewEsMapper(),
	}
}

func (server *InsertServer) Insert(_ context.Context, req *data.InsertReq) (bool, error) {

	var buf strings.Builder

	for _, doc := range req.Documents {

		index := "business_event-" + doc.EventName

		meta := []byte(fmt.Sprintf(`{ "index" : { "_index" : "%s" } }%s`, index, "\n"))

		// 将单引号换成双引号避免违背es的语法
		content := []byte(strings.ReplaceAll(doc.Tags, "'", "\""))

		content = append(content, '\n')

		buf.Write(meta)
		buf.Write(content)
	}

	err := server.mapper.Insert([]byte(buf.String()))

	if err != nil {
		return false, err
	}
	return true, nil
}
