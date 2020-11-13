/*
   Created by guoxin in 2020/4/13 1:34 下午
*/
package cache_redis

import (
	"github.com/GuoxinL/gcomponent/environment"
	"github.com/GuoxinL/gcomponent/logging"
	"github.com/GuoxinL/gcomponent/tools"
	"github.com/gomodule/redigo/redis"
	"github.com/mna/redisc"
	"time"
)

func init() {
	new(Configuration).Initialize()
}

var instances *dbInstanceMap

type Configuration struct {
	RedisSources []RedisSource `mapstructure:"rds"`
}

type Pool struct {
	MaxActive       int `mapstructure:"maxActive"`
	MaxIdle         int `mapstructure:"maxIdle"`
	IdleTimeout     int `mapstructure:"idleTimeout"`
	MaxConnLifetime int `mapstructure:"maxConnLifetime"`
	TestOnBorrow    int `mapstructure:"testOnBorrow"`
}

type Option struct {
	ConnectTimeout int `mapstructure:"connectTimeout"`
	WriteTimeout   int `mapstructure:"writeTimeout"`
	ReadTimeout    int `mapstructure:"readTimeout"`
}

type RedisSource struct {
	Name     string   `mapstructure:"name"`
	Nodes    []string `mapstructure:"nodes"`
	Password string   `mapstructure:"password"`
	Pool     Pool     `mapstructure:"pool"`
	Option   Option   `mapstructure:"option"`
}

func (c *Configuration) Initialize(params ...interface{}) interface{} {
	err := environment.GetProperty("components.redis", &c)
	if err != nil {
		logging.Exitf("GComponent [Redis]读取配置异常, 退出程序！！！\n异常信息: %v", err.Error())
	}
	instances = newRedisInstanceMap()
	for _, source := range c.RedisSources {
		cluster := redisc.Cluster{
			StartupNodes: source.Nodes,
			DialOptions: []redis.DialOption{
				redis.DialConnectTimeout(getDurationSecond(source.Option.ConnectTimeout)),
				redis.DialWriteTimeout(getDurationSecond(source.Option.WriteTimeout)),
				redis.DialReadTimeout(getDurationSecond(source.Option.ReadTimeout)),
				redis.DialPassword(source.Password),
			},
			CreatePool: func(address string, options ...redis.DialOption) (pool *redis.Pool, err error) {
				redispool := &redis.Pool{
					MaxActive:       source.Pool.MaxActive,
					MaxIdle:         source.Pool.MaxIdle,
					IdleTimeout:     getDurationSecond(source.Pool.IdleTimeout),
					MaxConnLifetime: getDurationSecond(source.Pool.MaxConnLifetime),
					Dial: func() (conn redis.Conn, e error) {
						c, err := redis.Dial("tcp", address, options...)
						if err != nil {
							return nil, err
						}
						if _, err := c.Do("SELECT", 0); err != nil {
							_ = c.Close()
							return nil, err
						}
						return c, nil
					},
					TestOnBorrow: func(c redis.Conn, t time.Time) error {
						if time.Since(t) < getDurationSecond(source.Pool.TestOnBorrow) {
							return nil
						}
						_, err := c.Do("PING")
						return err
					},
				}
				return redispool, nil
			},
		}

		instances.Put(source.Name, &clusterWrapper{cluster})
	}
	return nil
}

func getDurationSecond(i int) time.Duration {
	return time.Duration(i) * time.Second
}

/**
通过该方法获得*redisc.Cluster对象
*/
func GetInstance(name string) *redisc.Cluster {
	instance := instances.Get(name)
	if instance == nil {
		logging.Error0("未找到`" + name + "`对应的数据库连接，请核对配置文件")
		return nil
	}
	return instance.(*redisc.Cluster)
}

type dbInstanceMap struct {
	*tools.ConcurrentMap
}

func newRedisInstanceMap() *dbInstanceMap {
	return &dbInstanceMap{tools.NewConcurrentMap()}
}
