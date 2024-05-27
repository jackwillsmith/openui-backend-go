package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"github.com/openui-backend-go/common/jwtx"
	"github.com/openui-backend-go/service/user-rpc/userclient"

	"github.com/openui-backend-go/service/user-api/internal/svc"
	"github.com/openui-backend-go/service/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	res, err := l.svcCtx.UserRpc.Login(l.ctx, &userclient.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		logc.Error(l.ctx, "user login failed", err)
		return nil, err
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire

	accessToken, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, res.Id)
	if err != nil {
		logc.Error(l.ctx, "get access token failed", err)
		return nil, err
	}

	return &types.LoginResponse{
		Id:    res.Id,
		Name:  res.Name,
		Role:  res.Role,
		Token: accessToken,
	}, nil
}
