package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"openui-backend-go/common/database"
)

type Config struct {
	zrpc.RpcServerConf

	Mysql database.DbConfig

	CacheRedis database.RedisConfig
	Salt       string
}
