package logic

import (
	"context"
	"github.com/openui-backend-go/common/model"
	"github.com/openui-backend-go/service/chat-rpc/chat"
	"github.com/openui-backend-go/service/chat-rpc/internal/svc"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *chat.CreateRequest) (*chat.CreateResponse, error) {
	// 1. 入库
	newChat := model.Chat{
		UserId:     in.UserId,
		Title:      in.Title,
		Chat:       in.Chat,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err := l.svcCtx.ChatModel.Insert(l.ctx, &newChat)
	if err != nil {
		return nil, err
	}

	return &chat.CreateResponse{
		Id: newChat.Id,
	}, nil
}
