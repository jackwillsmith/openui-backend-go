package logic

import (
	"context"
	"time"

	"github.com/openui-backend-go/common/consts"
	"github.com/openui-backend-go/common/cryptx"
	"github.com/openui-backend-go/common/model"
	"google.golang.org/grpc/status"

	"github.com/openui-backend-go/service/user-rpc/internal/svc"
	"github.com/openui-backend-go/service/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// 判断手机号是否已经注册
	_, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err == nil {
		return nil, status.Error(consts.ACCOUNT_ALREADY_EXISTS, consts.WrongMessageEn[consts.ACCOUNT_ALREADY_EXISTS])
	}

	if err == consts.ERROR_NOT_FOUND {
		newUser := model.User{
			Name:       in.Name,
			Email:      in.Email,
			Password:   cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}

		err = l.svcCtx.UserModel.Insert(l.ctx, &newUser)
		if err != nil {
			return nil, status.Error(consts.ERROR_INTERNAL_SERVER_OPERATION, err.Error())
		}

		return &user.RegisterResponse{
			Id:   newUser.Id,
			Name: newUser.Name,
		}, nil

	}

	return nil, status.Error(500, err.Error())
}
