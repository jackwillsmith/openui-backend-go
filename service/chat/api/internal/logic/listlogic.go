package logic

import (
	"context"
	"github.com/openui-backend-go/service/chat-api/internal/svc"
	"github.com/openui-backend-go/service/chat-api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

/***
func (l *ListLogic) List(req *types.ListRequest) (resp []types.DetailResponse, err error) {
	res, err := l.svcCtx.ChatRpc.ListChat(l.ctx, &chat.Empty{})
	if err != nil {
		logc.Error(l.ctx, "ListChat error: %v", err)
		return nil, err
	}
	resp = make([]types.DetailResponse, 0)
	chats := []types.DetailResponse{}
	for _, v := range res.List {
		chat := types.DetailResponse{
			Id:       v.Id,
			UserId:   v.UserId,
			Title:    v.Title,
			Chat:     v.Chat,
			ShareId:  v.ShareId,
			Archived: v.Archived,
		}
		chats = append(chats, chat)
	}
	resp = chats
	return
}
***/

func (l *ListLogic) List(req *types.ListRequest) (resp []types.Chat, err error) {
	/***
	  [
	    {
	        "id": "4c1aa140-2387-48ef-9151-ed198b2a9f6c",
	        "title": "New Chat",
	        "updated_at": 1716795687,
	        "created_at": 1716795625
	    }
	  ]
	  ***/
	resp = make([]types.Chat, 0)
	chat := types.Chat{
		ID:        "4c1aa140-2387-48ef-9151-ed198b2a9f6c",
		Title:     "New Chat",
		UpdatedAt: 1716795687,
		CreatedAt: 1716795625,
	}
	resp = append(resp, chat)
	return
}
