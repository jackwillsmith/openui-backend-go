package database

import (
	"context"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/zeromicro/go-zero/core/logc"
)

var InitFlag bool = false

type DcRedisClient struct {
	name string      "" //redis库名
	db   int         //redis库索引
	pool *redis.Pool //Redis连接池
}

// RedisDoExeCmd 执行redis命令（底层方法）
func (c *DcRedisClient) RedisDoExeCmd(ctx context.Context, cmd string, vals ...interface{}) (interface{}, error) {
	rc := c.pool.Get()
	replay, err := rc.Do(cmd, vals...)
	if err != nil {
		logc.Errorf(ctx, "%v", err)
		return replay, err
	}
	// 用完后将连接放回连接池
	defer rc.Close()
	return replay, err
}

// RedisDoDelKey 删除Redis数据
func (c *DcRedisClient) RedisDoDelKey(ctx context.Context, key string) bool {
	_, err := c.RedisDoExeCmd(ctx, "DEL", key)
	if err != nil {
		return false
	} else {
		return true
	}
}

// RedisDoExpireKey 给Redis数据设置过期时间
func (c *DcRedisClient) RedisDoExpireKey(ctx context.Context, name string, sec int) bool {
	_, err := c.RedisDoExeCmd(ctx, "EXPIRE", name, sec)
	if err != nil {
		return false
	} else {
		return true
	}
}

// RedisDoSetExKey 为指定的 key 设置值及其过期时间。如果 key 已经存在， SETEX 命令将会替换旧的值
func (c *DcRedisClient) RedisDoSetExKey(ctx context.Context, key string, sec int, val string) bool {
	_, err := c.RedisDoExeCmd(ctx, "SETEX", key, sec, val)
	if err != nil {
		return false
	} else {
		return true
	}
}

// RedisDoSetKey Redis添加数据
func (c *DcRedisClient) RedisDoSetKey(ctx context.Context, key string, val string) bool {
	_, err := c.RedisDoExeCmd(ctx, "SET", key, val)
	if err != nil {
		return false
	} else {
		return true
	}
}

// RedisDoSetKeyWithExpire Redis给指定key设置过期时间
func (c *DcRedisClient) RedisDoSetKeyWithExpire(ctx context.Context, key string, val string, sec int64) bool {
	_, err := c.RedisDoExeCmd(ctx, "SET", key, val, "EX", sec)
	if err != nil {
		return false
	} else {
		return true
	}
}

// RedisDoGetKey 获取Redis数据（string）
func (c *DcRedisClient) RedisDoGetKey(ctx context.Context, key string) (string, bool) {
	resValue, resErr := c.RedisDoExeCmd(ctx, "GET", key)
	if resErr != nil {
		return "", false
	}
	if resValue == nil {
		logc.Errorf(ctx, "key: %s token is null", key)
		return "", false
	}
	strval, err := redis.String(resValue, resErr)
	if err != nil {
		logc.Errorf(ctx, "%v", err)
		return "", false
	}
	logc.Infof(ctx, "key: %s token: %s", key, strval)
	return strval, true
}

// RedisDoGetInt64 获取Redis数据（Int64）
func (c *DcRedisClient) RedisDoGetInt64(ctx context.Context, key string) (val int64, ok bool) {
	resValue, resErr := c.RedisDoExeCmd(ctx, "GET", key)
	if resErr != nil {
		return 0, false
	}
	if resValue == nil {
		return 0, false
	}
	val, err := redis.Int64(resValue, resErr)
	if err != nil {
		logc.Errorf(ctx, "%v", err)
		return 0, false
	}
	logc.Infof(ctx, "%v", val)
	return val, true
}

// RedisExistKey 判断Redis中是否有指定key的数据
func (c *DcRedisClient) RedisExistKey(ctx context.Context, key string) bool {
	resValue, resErr := c.RedisDoExeCmd(ctx, "EXISTS", key)
	if resErr != nil {
		return false
	}
	val, err := redis.Bool(resValue, resErr)
	if err != nil {
		logc.Errorf(ctx, "%v", err)
		return false
	}
	logc.Infof(ctx, "%v", val)
	return val
}

// RedisDoHMSetAll Redis批量设置数据
func (c *DcRedisClient) RedisDoHMSetAll(ctx context.Context, key string, valmap map[string]string) bool {
	rc := c.pool.Get()
	resValue, resErr := rc.Do("hmset", redis.Args{}.Add(key).AddFlat(valmap)...)
	if resErr != nil {
		logc.Infof(ctx, "%v", resErr)
		return false
	}
	logc.Infof(ctx, "%v", resValue)
	return true
}

// RedisDoHMSetField Redis设置单个数据
func (c *DcRedisClient) RedisDoHMSetField(ctx context.Context, keyname, field, value string) bool {
	_, resErr := c.RedisDoExeCmd(ctx, "HMSET", keyname, field, value)
	if resErr != nil {
		return false
	}
	return true
}

// RedisDoHMSetStruct Redis设置结构体数据
func (c *DcRedisClient) RedisDoHMSetStruct(ctx context.Context, key string, v interface{}) bool {
	rc := c.pool.Get()
	resValue, resErr := rc.Do("hmset", redis.Args{}.Add(key).AddFlat(v)...)
	if resErr != nil {
		logc.Errorf(ctx, "%v", resErr)
		return false
	}
	logc.Errorf(ctx, "%v", resValue)
	return true
}

// RedisDoHGetAll Redis获取多条数据
func (c *DcRedisClient) RedisDoHGetAll(ctx context.Context, key string) (*map[string]string, bool) {
	resValue, resErr := c.RedisDoExeCmd(ctx, "HGETALL", key)
	if resErr != nil {
		return nil, false
	}
	strvalmap, err := redis.StringMap(resValue, resErr)
	if err != nil {
		logc.Errorf(ctx, "%v", err)
		return nil, false
	}
	// logc.Info(strvalmap)
	return &strvalmap, true
}

// RedisDoHGetStruct Redis获取结构体数据
func (c *DcRedisClient) RedisDoHGetStruct(ctx context.Context, key string, p2 interface{}) bool {
	resValue, resErr := c.RedisDoExeCmd(ctx, "HGETALL", key)
	if resErr != nil {
		return false
	}
	v, err := redis.Values(resValue, resErr)
	if len((v)) == 0 {
		logc.Info(ctx, "resValue NULL", resValue)
		return false
	}
	if err != nil {
		logc.Errorf(ctx, "%v", err)
		return false
	}
	if err := redis.ScanStruct(v, p2); err != nil {
		logc.Errorf(ctx, "%v", err)
		return false
	}
	return true
}

// RedisDoRpush 将一个值插入到列表的尾部（最右边）
func (c *DcRedisClient) RedisDoRpush(ctx context.Context, keyname, value string) bool {
	_, resErr := c.RedisDoExeCmd(ctx, "RPUSH", keyname, value)
	if resErr != nil {
		return false
	}
	return true
}

// RedisDoGetList 获取列表指定范围内的元素
func (c *DcRedisClient) RedisDoGetList(ctx context.Context, key string, begin string, end string) ([]string, bool) {
	resValue, resErr := c.RedisDoExeCmd(ctx, "LRANGE", key, begin, end)
	if resErr != nil {
		return nil, false
	}
	strvalmap, err := redis.Strings(resValue, resErr)
	if err != nil {
		logc.Errorf(ctx, "%v", err)
		return nil, false
	}
	// logc.Info(strvalmap)
	return strvalmap, true
}

// RedisDoSADD 向集合添加一个成员
func (c *DcRedisClient) RedisDoSADD(ctx context.Context, keyname, value string) bool {
	_, resErr := c.RedisDoExeCmd(ctx, "SADD", keyname, value)
	if resErr != nil {
		return false
	}
	return true
}

// RedisDoSISMEMBER 判断member元素是否是集合key的成员
func (c *DcRedisClient) RedisDoSISMEMBER(ctx context.Context, keyname, value string) bool {
	resValue, resErr := c.RedisDoExeCmd(ctx, "SISMEMBER", keyname, value)
	if resErr != nil {
		return false
	}
	val, err := redis.Bool(resValue, resErr)
	if err != nil {
		logc.Errorf(ctx, "%v", err)
		return false
	}
	logc.Infof(ctx, "%v", val)
	return val
}

// RedisDoOneZADD 向有序集合添加一个或多个成员，或者更新已存在成员的分数
func (c *DcRedisClient) RedisDoOneZADD(ctx context.Context, keyname string, score int64, value string) bool {
	_, resErr := c.RedisDoExeCmd(ctx, "ZADD", keyname, score, value)
	if resErr != nil {
		return false
	}
	return true
}

// RedisDoOneZCARD 获取有序集合的成员数量
func (c *DcRedisClient) RedisDoOneZCARD(ctx context.Context, keyname string) (int, bool) {
	resValue, resErr := c.RedisDoExeCmd(ctx, "ZCARD", keyname)
	if resErr != nil {
		return 0, false
	}
	val, err := redis.Int(resValue, resErr)
	if err != nil {
		logc.Errorf(ctx, "%v", err)
		return 0, false
	}
	logc.Infof(ctx, "%v", val)
	return val, true
}

// DoOneZRANK 返回有序集合中指定成员的索引
func (c *DcRedisClient) DoOneZRANK(ctx context.Context, keyname, value string, desc bool) (int, bool) {
	cmd := "ZREVRANK"
	if desc {
		cmd = "ZREVRANK"
	} else {
		cmd = "ZRANK"
	}
	resValue, resErr := c.RedisDoExeCmd(ctx, cmd, keyname, value)
	if resErr != nil {
		return 0, false
	}
	val, err := redis.Int(resValue, resErr)
	if err != nil {
		logc.Errorf(ctx, "%v", err)
		return 0, false
	}
	logc.Infof(ctx, "%v", val)
	return val, true
}

// DoZRANGE 通过索引区间返回有序集合指定区间内的string类型成员（根据参数排序）
func (c *DcRedisClient) DoZRANGE(ctx context.Context, keyname string, beg, end int, desc bool) ([]string, bool) {
	var cmd = "ZREVRANGE"
	if desc {
		cmd = "ZREVRANGE"
	} else {
		cmd = "ZRANGE"
	}
	resValue, resErr := c.RedisDoExeCmd(ctx, cmd, keyname, beg, end)
	if resErr != nil {
		return nil, false
	}
	val, err := redis.Strings(resValue, resErr)
	if err != nil {
		logc.Errorf(ctx, "%v", err)
		return nil, false
	}
	logc.Infof(ctx, "%v", val)
	return val, true
}

// DoZrangeInt64 通过索引区间返回有序集合指定区间内的int64类型成员（根据参数排序）
func (c *DcRedisClient) DoZrangeInt64(ctx context.Context, keyname string, beg, end int, desc bool) ([]int64, bool) {
	var cmd = "ZREVRANGE"
	if desc {
		cmd = "ZREVRANGE"
	} else {
		cmd = "ZRANGE"
	}
	resValue, resErr := c.RedisDoExeCmd(ctx, cmd, keyname, beg, end)
	if resErr != nil {
		return nil, false
	}
	val, err := redis.Int64s(resValue, resErr)
	if err != nil {
		logc.Errorf(ctx, "%v", err)
		return nil, false
	}
	logc.Infof(ctx, "%v", val)
	return val, true
}

// DoINCR 将key中储存的数字值增一
func (c *DcRedisClient) DoINCR(ctx context.Context, keyname string) (int64, error) {
	resValue, resErr := c.RedisDoExeCmd(ctx, "INCR")
	if resErr != nil {
		return 0, resErr
	}
	val, err := redis.Int64(resValue, resErr)
	if err != nil {
		logc.Errorf(ctx, "%v", err)
		return val, err
	}
	logc.Infof(ctx, "%v", val)
	return val, err
}

// DoINCRBY 将key中储存的数字加上指定的增量值
func (c *DcRedisClient) DoINCRBY(ctx context.Context, keyname string, v int64) (int64, error) {
	resValue, resErr := c.RedisDoExeCmd(ctx, "INCRBY", keyname, v)
	if resErr != nil {
		return 0, resErr
	}
	val, err := redis.Int64(resValue, resErr)
	if err != nil {
		logc.Errorf(ctx, "%v", err)
		return val, err
	}
	logc.Infof(ctx, "%v", val)
	return val, err
}

// StoreStructIntoRedis 将结构体存储到Redis
func StoreStructIntoRedis(ctx context.Context, redisclient *DcRedisClient,
	key string, info interface{}, timelen int) bool {
	if info == nil {
		return false
	}
	// redisclient := GetTokenRedisHandle()
	if redisclient != nil {
		// reginfoPoint = make(interface{})
		if redisclient.RedisDoHMSetStruct(ctx, key, info) {
			if timelen > 0 {
				if redisclient.RedisDoExpireKey(ctx, key, timelen) {
					return true
				} else {
					logc.Error(ctx, "RedisDoExpireKey failed")
				}
			} else {
				return true
			}
		} else {
			logc.Error(ctx, "RedisDoHMSetAll failed")
		}
	} else {
		logc.Error(ctx, "redisclient nil")
	}
	return false
}

// NewDcRedisConnClient
func NewDcRedisConnClient(ctx context.Context, network, addr, auth string, db int) *DcRedisClient {
	var c DcRedisClient
	c.name = addr
	c.db = db
	c.pool = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle: 2, //最大的空闲连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,                 //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300 * time.Second, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			logc.Infof(ctx, "%v %v %v", network, addr, db)
			c, err := redis.Dial(network, addr, redis.DialDatabase(db))
			if err != nil {
				logc.Errorf(ctx, "%v", err)
				return nil, err
			}
			if auth != "" {
				if _, err := c.Do("AUTH", auth); err != nil {
					logc.Errorf(ctx, "%v", err)
					c.Close()
					return nil, err
				}
			}
			return c, nil
		},
	}
	return &c
}
