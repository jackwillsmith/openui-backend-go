package logic

import (
	"context"

	"github.com/openui-backend-go/product-rpc/internal/svc"
	"github.com/openui-backend-go/product-rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *product.DetailRequest) (*product.DetailResponse, error) {
	// 获取product详情
	// 1. 获取in中的参数
	id := in.Id
	// 2. 调用model层获取数据
	res, err := l.svcCtx.ProductModel.FindOne(l.ctx, id)
	if err != nil {
		return nil, err
	}
	return &product.DetailResponse{
		Id:     id,
		Name:   res.Name,
		Desc:   res.Desc,
		Status: res.Status,
	}, nil
}
