package logic

import (
	"context"

	"github.com/openui-backend-go/service/chat-api/internal/svc"
	"github.com/openui-backend-go/service/chat-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatLogic {
	return &ChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatLogic) Chat(req *types.NewChatEntity) (resp *types.ChatRespone, err error) {
	// todo: add your logic here and delete this line

	return
}
