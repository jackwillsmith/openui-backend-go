package config

import (
	"github.com/openui-backend-go/common/database"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Mysql   database.DbConfig
	LogConf logc.LogConf //logx与logc的区别，logc是对logx的封装，可以加上context进行日志打印

	CacheRedis database.RedisConfig
	Salt       string
}
