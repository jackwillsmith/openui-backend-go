package logic

import (
	"context"
	"github.com/openui-backend-go/common/callmodel"
	"github.com/openui-backend-go/common/utils"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/openui-backend-go/service/chat-rpc/chat"
	"github.com/openui-backend-go/service/chat-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenPromptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenPromptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenPromptLogic {
	return &GenPromptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenPromptLogic) GenPrompt(in *chat.NewChatEntity) (*chat.CallResponse, error) {
	// 1. 初始化 oll client
	ollCli := callmodel.NewOllamaModel()
	uri := utils.GetOllUrl()

	model := in.Model

	err := ollCli.GetModelClient(l.ctx, uri, model)
	if err != nil {
		logc.Error(l.ctx, "ollama get model client failed", err)
		return nil, err
	}

	// 2. 调用 oll call 接口
	var prompt string = ""
	for _, message := range in.Messages {
		if message.Role == "user" {
			prompt = message.Content
			break
		}
	}
	res, err := ollCli.CallModel(l.ctx, prompt)
	if err != nil {
		logc.Error(l.ctx, "ollama generate content failed", err)
		return nil, err
	}

	resp := &chat.CallResponse{
		Text: res,
	}

	return resp, nil
}
