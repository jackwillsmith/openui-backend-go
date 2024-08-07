package main

import (
	"flag"
	"fmt"
	"github.com/openui-backend-go/common/middleware"
	"github.com/openui-backend-go/common/utils"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/openui-backend-go/service/chat-api/internal/config"
	"github.com/openui-backend-go/service/chat-api/internal/handler"
	"github.com/openui-backend-go/service/chat-api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/chat.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logc.MustSetup(c.LogConf)

	utils.SetOllUrl(c.OllUrl)

	server := rest.MustNewServer(c.RestConf, rest.WithCors("*"))
	// 设置静态文件中间件
	middleware.MiddelwareStatic(server)

	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
