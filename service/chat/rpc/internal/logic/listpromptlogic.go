package logic

import (
	"context"

	"github.com/openui-backend-go/service/chat-rpc/chat"
	"github.com/openui-backend-go/service/chat-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPromptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPromptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPromptLogic {
	return &ListPromptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListPromptLogic) ListPrompt(in *chat.Empty) (*chat.ListPrompts, error) {

	return &chat.ListPrompts{}, nil
}
