package logic

import (
	"context"

	"github.com/openui-backend-go/service/chat-api/internal/svc"
	"github.com/openui-backend-go/service/chat-api/internal/types"
	"github.com/openui-backend-go/service/chat-rpc/chatclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	resp = &types.DetailResponse{}
	res, err := l.svcCtx.ChatRpc.Detail(l.ctx, &chatclient.DetailRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.DetailResponse{
		Id:       res.Id,
		Title:    res.Title,
		UserId:   res.UserId,
		Chat:     res.Chat,
		ShareId:  res.ShareId,
		Archived: res.Archived,
	}
	return resp, nil
}
