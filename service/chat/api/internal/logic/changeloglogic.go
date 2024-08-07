package logic

import (
	"context"

	"github.com/openui-backend-go/service/chat-api/internal/svc"
	"github.com/openui-backend-go/service/chat-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangelogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangelogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangelogLogic {
	return &ChangelogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangelogLogic) Changelog() (resp *types.ChangelogResponse, err error) {
	resp = &types.ChangelogResponse{
		Changelog: "1.0.0",
	}
	return
}
