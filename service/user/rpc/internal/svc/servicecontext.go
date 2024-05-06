package svc

import (
	"openui-backend-go/common/database"
	"openui-backend-go/service/user/model"
	"openui-backend-go/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := database.New(c.Mysql)
	redisCli := database.NewDcRedisClient(c.CacheRedis)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, redisCli),
	}
}
