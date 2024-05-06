package database

import (
	"context"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/zeromicro/go-zero/core/logc"
)

type RedisConfig struct {
	Name    string `json:"name"`
	Nettype string `json:"nettype"`
	Address string `json:"address"`
	Auth    string `json:"auth"`
	DB      int    `json:"db"`
}

func initRedis(conf RedisConfig, c *DcRedisClient) {
	// 建立连接池
	logc.Info(context.Background(), "########### initRedis #########")

	c.pool = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle: 2, //最大的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,                 //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300 * time.Second, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			logc.Infof(context.Background(), "%v %v %v", conf.Nettype, conf.Address, c.db)
			c, err := redis.Dial(conf.Nettype, conf.Address, redis.DialDatabase(c.db))
			if err != nil {
				logc.Errorf(context.Background(), "%v", err)
				return nil, err
			}
			if conf.Auth != "" {
				if _, err := c.Do("AUTH", conf.Auth); err != nil {
					logc.Errorf(context.Background(), "%v", err)
					c.Close()
					return nil, err
				}
			}
			return c, nil
		},
	}
}

func NewDcRedisClient(conf RedisConfig) *DcRedisClient {
	var c DcRedisClient
	c.name = conf.Name
	c.db = conf.DB
	c.pool = nil

	initRedis(conf, &c)
	// c.RedisExistKey("001")
	return &c
}
