package logic

import (
	"context"
	"github.com/openui-backend-go/service/chat-rpc/chat"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/openui-backend-go/service/chat-api/internal/svc"
	"github.com/openui-backend-go/service/chat-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPromptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPromptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPromptLogic {
	return &GetPromptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPromptLogic) GetPrompt() (resp *types.PromptResponse, err error) {
	res, err := l.svcCtx.ChatRpc.ListPrompt(l.ctx, &chat.Empty{})
	if err != nil {
		logc.Error(l.ctx, "ListPrompt error: %v", err)
		return
	}

	// res.List 转换为 []Prompt
	var prompts []types.Prompt
	for _, v := range res.List {
		prompt := types.Prompt{
			Id:      v.Id,
			Command: v.Command,
			Content: v.Content,
			Title:   v.Title,
			UserId:  v.UserId,
		}
		prompts = append(prompts, prompt)
	}
	resp = &types.PromptResponse{
		Prompts: prompts,
	}

	return
}
