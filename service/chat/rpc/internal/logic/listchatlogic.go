package logic

import (
	"context"

	"github.com/openui-backend-go/service/chat-rpc/chat"
	"github.com/openui-backend-go/service/chat-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListChatLogic {
	return &ListChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListChatLogic) ListChat(in *chat.Empty) (*chat.ListChats, error) {
	// todo: add your logic here and delete this line

	return &chat.ListChats{}, nil
}
