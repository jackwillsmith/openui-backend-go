package logic

import (
	"context"

	"github.com/openui-backend-go/service/chat-api/internal/svc"
	"github.com/openui-backend-go/service/chat-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompleteLogic {
	return &CompleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompleteLogic) Complete(req *types.CompleteRequest) (resp *types.ChatRespone, err error) {
	// todo: add your logic here and delete this line

	return
}
