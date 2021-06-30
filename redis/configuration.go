// Package gredis
// Created by guoxin in 2020/4/13 1:34 下午
package gredis

import (
    "github.com/GuoxinL/gcomponent/core"
    "github.com/GuoxinL/gcomponent/core/tools"
    "github.com/GuoxinL/gcomponent/environment"
    "github.com/GuoxinL/gcomponent/zap"
    "go.uber.org/zap"
    "gopkg.in/redis.v5"
    "time"
)

var (
    initializeLock = core.NewInitLock()
    instances *instanceMap
)

func New(params ...interface{}) {
    c := &Configuration{
        InitializeLock: initializeLock,
    }
    c.Initialize(params...)
}

type Configuration struct {
    core.InitializeLock
    RedisSources []RedisSource `mapstructure:"rds"`
}

type RedisSource struct {
    core.InitializeLock
    Name          string                `mapstructure:"name"`
    Cluster       *redis.ClusterOptions `mapstructure:"Cluster"`
    redis.Options `mapstructure:",squash"`
}

func (c *Configuration) Initialize(params ...interface{}) interface{} {
    if c.IsInit() {
        return &instances
    }
    gzap.New()
    err := environment.GetProperty("components.redis", &c)
    if err != nil {
        zap.L().Fatal("GComponent [redis]Read configuration exceptions and exit.", zap.Error(err))
    }

    instances = &instanceMap{tools.NewConcurrentMap()}
    for _, source := range c.RedisSources {
        var client redis.Cmdable
        if len(source.Cluster.Addrs) != 0 {
            // 集群
            source.Cluster.DialTimeout = source.Cluster.DialTimeout * time.Second
            source.Cluster.IdleTimeout = source.Cluster.IdleTimeout * time.Second
            source.Cluster.PoolTimeout = source.Cluster.PoolTimeout * time.Second
            source.Cluster.ReadTimeout = source.Cluster.ReadTimeout * time.Second
            source.Cluster.WriteTimeout = source.Cluster.WriteTimeout * time.Second

            client = redis.NewClusterClient(source.Cluster)
        } else if len(source.Addr) != 0 {
            // 单机
            source.Options.DialTimeout = source.Options.DialTimeout * time.Second
            source.Options.IdleTimeout = source.Options.IdleTimeout * time.Second
            source.Options.PoolTimeout = source.Options.PoolTimeout * time.Second
            source.Options.ReadTimeout = source.Options.ReadTimeout * time.Second
            source.Options.WriteTimeout = source.Options.WriteTimeout * time.Second
            client = redis.NewClient(&source.Options)
        } else {
            // 都没有设置走默认
            client = redis.NewClient(DefaultOptions)
        }
        if err := client.Ping().Err(); err != nil {
            zap.L().Fatal("GComponent [redis]connection to Redis failed.", zap.Error(err))
        } else {
            zap.L().Info("GComponent [redis] ping success.")
        }
        instances.Put(source.Name, &client)
    }
    return nil
}

// GetInstance
// 通过该方法获得 *redis.Cmdable 对象
func GetInstance(name string) redis.Cmdable {
    instance := instances.Get(name)
    if instance == nil {
        zap.L().Error("The connection for `"+name+"` was not found, please check the configuration file", zap.String("name", name))
        return nil
    }
    return instance.(redis.Cmdable)
}

type instanceMap struct {
    *tools.ConcurrentMap
}

var DefaultOptions = &redis.Options{
    //连接信息
    Network:  "tcp",            //网络类型，tcp or unix，默认tcp
    Addr:     "127.0.0.1:6379", //主机名+冒号+端口，默认localhost:6379
    Password: "",               //密码
    DB:       0,                // redis数据库index

    //连接池容量
    PoolSize: 15, // 连接池最大socket连接数，默认为4倍CPU数， 4 * runtime.NumCPU

    //超时
    DialTimeout:  5 * time.Second, //连接建立超时时间，默认5秒。
    ReadTimeout:  3 * time.Second, //读超时，默认3秒， -1表示取消读超时
    WriteTimeout: 3 * time.Second, //写超时，默认等于读超时
    PoolTimeout:  4 * time.Second, //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。

    //闲置连接检查包括IdleTimeout，MaxConnAge
    IdleCheckFrequency: 60 * time.Second, //闲置连接检查的周期，默认为1分钟，-1表示不做周期性检查，只在客户端获取连接时对闲置连接进行处理。
    IdleTimeout:        5 * time.Minute,  //闲置超时，默认5分钟，-1表示取消闲置超时检查

    MaxRetries: 0,     // 命令执行失败时，最多重试多少次，默认为0即不重试
    ReadOnly:   false, // 只读
    TLSConfig:  nil,   // Redis从v6开始支持TLS https://redis.io/topics/encryption
    Dialer:     nil,
}
