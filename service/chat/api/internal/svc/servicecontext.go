package svc

import (
	"github.com/openui-backend-go/service/chat-api/internal/config"
	"github.com/openui-backend-go/service/chat-rpc/chatclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	ChatRpc chatclient.Chat
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		ChatRpc: chatclient.NewChat(zrpc.MustNewClient(c.ChatRpc)),
	}
}
