/**
  Create by guoxin 2020.11.13
*/
package cache_redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/mna/redisc"
)

type clusterWrapper struct {
	cluster redisc.Cluster
}

func (w *clusterWrapper) Get(key string) (string, error) {
	return redis.String(w.cluster.Get().Do("GET", key))
}

func (w *clusterWrapper) Set(key string, value interface{}) (string, error) {
	return redis.String(w.cluster.Get().Do("SET", key, value))
}
