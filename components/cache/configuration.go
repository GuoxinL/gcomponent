/*
   Created by guoxin in 2020/4/13 1:34 下午
*/
package cache

import (
	"github.com/GuoxinL/gcomponent/components/environment"
	"github.com/GuoxinL/gcomponent/components/logging"
	"github.com/GuoxinL/gcomponent/components/tools"
	"github.com/gomodule/redigo/redis"
	"github.com/mna/redisc"
	"time"
)

func init() {
	new(Configuration).Initialize()
}

var redisInstances *dbInstanceMap

type Configuration struct {
	RedisSources []RedisSource `yaml:"rds"`
}

type Pool struct {
	MaxActive       int `yaml:"maxActive"`
	MaxIdle         int `yaml:"maxIdle"`
	IdleTimeout     int `yaml:"idleTimeout"`
	MaxConnLifetime int `yaml:"maxConnLifetime"`
	TestOnBorrow    int `yaml:"testOnBorrow"`
}

type Option struct {
	ConnectTimeout int `yaml:"connectTimeout"`
	WriteTimeout   int `yaml:"writeTimeout"`
	ReadTimeout    int `yaml:"readTimeout"`
}

type RedisSource struct {
	Name     string   `yaml:"name"`
	Nodes    []string `yaml:"nodes"`
	Password string   `yaml:"password"`
	Pool     Pool     `yaml:"pool"`
	Option   Option   `yaml:"option"`
}

type clusterWrapper struct {
	cluster redisc.Cluster
}

func (this *Configuration) Initialize(params ...interface{}) interface{} {
	err := environment.GetConfig("components.redis", &this)
	if err != nil {
		logging.Exitf("GComponent [Mysql]读取配置异常, 退出程序！！！\n异常信息: %v", err.Error())
	}
	redisInstances = newRedisInstanceMap()
	for _, source := range this.RedisSources {
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

		redisInstances.Put(source.Name, &clusterWrapper{cluster})
	}
	GetInstance("xxx")
	return nil
}

func (this *clusterWrapper) Get(key string) (string, error) {
	return redis.String(this.cluster.Get().Do("GET", key))
}

func (this *clusterWrapper) Set(key string, value interface{}) (string, error) {
	return redis.String(this.cluster.Get().Do("SET", key, value))
}

func getDurationSecond(i int) time.Duration {
	return time.Duration(i) * time.Second
}

/**
通过该方法获得*redisc.Cluster对象
*/
func GetInstance(name string) *redisc.Cluster {
	instance := redisInstances.Get(name)
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
