package logic

import (
	"context"
	"encoding/json"

	"github.com/openui-backend-go/service/user-api/internal/svc"
	"github.com/openui-backend-go/service/user-api/internal/types"
	"github.com/openui-backend-go/service/user-rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	// 初始化 UserInfoLogic
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// 1. 根据上下文获取uid
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	// 2. 调用 user rpc 服务, 获取用户信息
	res, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userclient.UserInfoRequest{
		Id: uid,
	})
	if err != nil {
		return nil, err
	}
	// 3. 返回
	return &types.UserInfoResponse{
		Id:              res.Id,
		Name:            res.Name,
		Email:           res.Email,
		Role:            res.Role,
		ProfileImageUrl: res.ProfileImageUrl,
	}, nil
}
