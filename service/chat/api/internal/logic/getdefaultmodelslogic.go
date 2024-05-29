package logic

import (
	"context"

	"github.com/openui-backend-go/service/chat-api/internal/svc"
	"github.com/openui-backend-go/service/chat-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDefaultModelsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDefaultModelsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDefaultModelsLogic {
	return &GetDefaultModelsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDefaultModelsLogic) GetDefaultModels(req *types.DefaultModels) (resp *types.DefaultModels, err error) {
	// todo: add your logic here and delete this line

	return
}
