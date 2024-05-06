package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"openui-backend-go/common/consts"
	"openui-backend-go/common/cryptx"
	"openui-backend-go/service/user/model"
	"time"

	"openui-backend-go/service/user/rpc/internal/svc"
	"openui-backend-go/service/user/rpc/user"

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
	_, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err == nil {
		return nil, status.Error(consts.ACCOUNT_ALREADY_EXISTS, consts.WrongMessageEn[consts.ACCOUNT_ALREADY_EXISTS])
	}

	if err == consts.ERROR_NOT_FOUND {
		newUser := model.User{
			Name:       in.Name,
			Gender:     in.Gender,
			Mobile:     in.Mobile,
			Password:   cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}

		err = l.svcCtx.UserModel.Insert(l.ctx, &newUser)
		if err != nil {
			return nil, status.Error(consts.ERROR_INTERNAL_SERVER_OPERATION, err.Error())
		}

		return &user.RegisterResponse{
			Id:     newUser.Id,
			Name:   newUser.Name,
			Gender: newUser.Gender,
			Mobile: newUser.Mobile,
		}, nil

	}

	return nil, status.Error(500, err.Error())
}
