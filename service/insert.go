package service

import (
	"context"
	"fmt"
	"github.com/xh-polaris/gopkg/util/log"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/data"
	"platform-data/infra"
	"strings"
)

type IInsertServer interface {
	Insert(ctx context.Context, req *data.InsertReq) (bool, error)
}

type InsertServer struct {
	IInsertServer
	mapper infra.IEsMapper
}

func NewInsertServer() *InsertServer {
	return &InsertServer{
		mapper: infra.NewEsMapper(),
	}
}

func (server *InsertServer) Insert(_ context.Context, req *data.InsertReq) (bool, error) {

	var buf strings.Builder

	for _, doc := range req.Documents {

		index := "business_event-" + doc.EventName

		meta := []byte(fmt.Sprintf(`{ "index" : { "_index" : "%s" } }%s`, index, "\n"))

		content := []byte(doc.Tags)
		content = append(content, '\n')

		log.Info("%s : %s\n", string(meta), content)

		buf.Write(meta)
		buf.Write(content)
	}

	err := server.mapper.Insert([]byte(buf.String()))

	if err != nil {
		return false, err
	}
	return true, nil
}
