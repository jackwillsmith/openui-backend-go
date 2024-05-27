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
	// 1. 调用 chat list方法
	res, err := l.svcCtx.ChatModel.List(l.ctx)
	if err != nil {
		logx.Error("List error: %v", err)
		return nil, err
	}
	// 2. 将返回的 chats封装成响应
	var chats []*chat.DetailResponse
	for _, v := range *res {
		chat := &chat.DetailResponse{
			Id:       v.Id,
			UserId:   v.UserId,
			Title:    v.Title,
			Chat:     v.Chat,
			ShareId:  v.ShareId,
			Archived: v.Archived,
		}
		chats = append(chats, chat)
	}
	resp := &chat.ListChats{
		List: chats,
	}

	return resp, nil
}
