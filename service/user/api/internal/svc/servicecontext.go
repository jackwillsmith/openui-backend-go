package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"openui-backend-go/service/user/api/internal/config"
	"openui-backend-go/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config

	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
