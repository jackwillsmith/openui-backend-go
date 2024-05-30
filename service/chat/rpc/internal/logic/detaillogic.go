package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/openui-backend-go/service/chat-rpc/chat"
	"github.com/openui-backend-go/service/chat-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *chat.DetailRequest) (*chat.DetailResponse, error) {
	id := in.Id
	// 调用 mysql 方法
	res, err := l.svcCtx.ChatModel.FindOne(l.ctx, id)
	if err != nil {
		logc.Error(l.ctx, "DetailLogic.Detail", "err: %v", err)
		return nil, err
	}

	resp := &chat.DetailResponse{
		Id:       res.Id,
		Title:    res.Title,
		ShareId:  res.ShareId,
		UserId:   res.UserId,
		Chat:     res.Chat,
		Archived: res.Archived,
	}

	return resp, nil
}
