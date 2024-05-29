package logic

import (
	"context"

	"github.com/openui-backend-go/service/chat-api/internal/svc"
	"github.com/openui-backend-go/service/chat-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigLogic {
	return &ConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigLogic) Config() (resp *types.ConfigResponse, err error) {
	var defaultPromptSuggestions []types.DefaultPromptSuggestions
	defaultPromptSuggestions = append(defaultPromptSuggestions, types.DefaultPromptSuggestions{
		Title:   []string{"Help me study", "vocabulary for a college entrance exam"},
		Content: "Help me study vocabulary: write a sentence for me to fill in the blank, and I'll try to pick the correct option.",
	}, types.DefaultPromptSuggestions{
		Title:   []string{"Give me ideas", "for what to do with my kids' art"},
		Content: "What are 5 creative things I could do with my kids' art? I don't want to throw them away, but it's also so much clutter.",
	}, types.DefaultPromptSuggestions{
		Title:   []string{"Tell me a fun fact", "about the Roman Empire"},
		Content: "Tell me a random fun fact about the Roman Empire",
	})

	resp = &types.ConfigResponse{
		Status:                   true,
		Name:                     "Open WebUI",
		Version:                  "0.1.120",
		DefaultLocale:            "en-US",
		Images:                   false,
		DefaultModels:            nil,
		DefaultPromptSuggestions: defaultPromptSuggestions,
		TrustedHeaderAuth:        false,
	}

	return
}
