package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"openui-backend-go/common/consts"
	"openui-backend-go/service/user/rpc/internal/svc"
	"openui-backend-go/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	// 查询用户是否存在
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == consts.ERROR_NOT_FOUND {
			return nil, status.Error(consts.ACCESS_NOT_FOUND, consts.WrongMessageEn[consts.ACCESS_NOT_FOUND])
		}
		return nil, status.Error(consts.ERROR_INTERNAL_SERVER_OPERATION, err.Error())
	}

	return &user.UserInfoResponse{
		Id:     res.Id,
		Name:   res.Name,
		Gender: res.Gender,
		Mobile: res.Mobile,
	}, nil
}
