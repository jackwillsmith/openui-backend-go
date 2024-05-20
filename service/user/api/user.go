package main

import (
	"flag"
	"fmt"
	"github.com/openui-backend-go/service/user-api/internal/config"
	"github.com/openui-backend-go/service/user-api/internal/handler"
	"github.com/openui-backend-go/service/user-api/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logc.MustSetup(c.LogConf)

	server := rest.MustNewServer(c.RestConf, rest.WithCors("*"))
	defer server.Stop()

	//server.Use(middleware.NewCorsMiddleware().Handle)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
