package logic

import (
	"context"

	"github.com/openui-backend-go/common/callmodel"
	"github.com/openui-backend-go/common/consts"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/openui-backend-go/product-rpc/internal/svc"
	"github.com/openui-backend-go/product-rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallLogic {
	return &CallLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CallLogic) Call(in *product.CallRequest) (*product.CallResponse, error) {
	// 1. 获取in中的参数
	modname := in.Name
	prom := in.Prompt

	// 2. 获取model client链接
	llm := callmodel.NewOllamaModel()
	err := llm.GetModelClient(l.ctx, consts.Inter, modname)
	if err != nil {
		logc.Error(l.ctx, "ollama get model client failed", err)
		return nil, err
	}
	res, err := llm.CallModel(l.ctx, prom)
	if err != nil {
		logc.Error(l.ctx, "ollama call model failed", err)
		return nil, err
	}

	return &product.CallResponse{
		Text: res,
	}, nil
}
