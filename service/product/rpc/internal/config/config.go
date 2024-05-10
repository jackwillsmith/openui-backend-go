package config

import (
	"github.com/openui-backend-go/common/database"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Mysql database.DbConfig

	CacheRedis database.RedisConfig
	Salt       string
}
