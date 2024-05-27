package logic

import (
	"context"

	"github.com/openui-backend-go/common/consts"
	"github.com/openui-backend-go/common/cryptx"
	"github.com/openui-backend-go/service/user-rpc/internal/svc"
	"github.com/openui-backend-go/service/user-rpc/user"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// 查询用户是否存在
	res, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		if err == consts.ERROR_NOT_FOUND {
			return nil, status.Error(consts.ACCESS_NOT_FOUND, consts.WrongMessageEn[consts.ACCESS_NOT_FOUND])
		}
		return nil, status.Error(consts.ERROR_INTERNAL_SERVER_OPERATION, err.Error())
	}

	// 判断密码是否正确
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if password != res.Password {
		return nil, status.Error(consts.ACCESS_NOT_FOUND, consts.WrongMessageEn[consts.ACCESS_PWD_WRONG])
	}

	return &user.LoginResponse{
		Id:   res.Id,
		Name: res.Name,
		Role: res.Role,
	}, nil
}
