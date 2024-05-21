package svc

import (
	"github.com/openui-backend-go/common/database"
	"github.com/openui-backend-go/common/model"
	"github.com/openui-backend-go/service/chat-rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	ChatModel model.ChatModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := database.New(c.Mysql)
	redisCli := database.NewDcRedisClient(c.CacheRedis)
	return &ServiceContext{
		Config:    c,
		ChatModel: model.NewChatModel(conn, redisCli),
	}
}
