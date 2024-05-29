package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/openui-backend-go/service/chat-rpc/chatclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/openui-backend-go/service/chat-api/internal/svc"
	"github.com/openui-backend-go/service/chat-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.NewChatRequest) (resp *types.CreateResponse, err error) {
	resp = &types.CreateResponse{}

	chatReq := req.Chat

	uid, _ := l.ctx.Value("uid").(json.Number).Float64()

	chatByte, _ := json.Marshal(chatReq)

	// 构建 newChat
	newChat := chatclient.CreateRequest{
		UserId: fmt.Sprintf("%d", int64(uid)),
		Title:  chatReq.Title,
		Chat:   string(chatByte),
	}
	res, err := l.svcCtx.ChatRpc.Create(l.ctx, &newChat)
	if err != nil {
		logc.Error(l.ctx, "create chat error: %v", err)
		return nil, err
	}
	resp.Id = res.Id
	return
}
