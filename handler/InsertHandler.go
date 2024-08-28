package handler

import (
	"context"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/data"
	"platform-data/service"
)

type IInsertHandler interface {
	Insert(ctx context.Context, req *data.InsertReq) (*data.InsertResp, error)
}

type InsertHandler struct {
	IInsertHandler
	insertService service.IInsertServer
}

func NewInsertHandler() *InsertHandler {
	return &InsertHandler{
		insertService: service.NewInsertServer(),
	}
}

func (h *InsertHandler) Insert(ctx context.Context, req *data.InsertReq) (*data.InsertResp, error) {

	success, err := h.insertService.Insert(ctx, req)

	if err != nil {
		return nil, err
	}

	return &data.InsertResp{
		Done: success,
	}, nil
}
