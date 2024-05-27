package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/openui-backend-go/service/chat-rpc/chat"
	"github.com/openui-backend-go/service/chat-rpc/internal/config"
	"github.com/openui-backend-go/service/chat-rpc/internal/server"
	"github.com/openui-backend-go/service/chat-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/chat.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logc.MustSetup(c.LogConf)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		chat.RegisterChatServer(grpcServer, server.NewChatServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
