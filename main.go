/*
   Created by guoxin in 2020/4/10 3:26 下午
*/
package main

import (
	cache_redis "github.com/GuoxinL/gcomponent/cache-redis"
	"github.com/GuoxinL/gcomponent/environment"
	"github.com/GuoxinL/gcomponent/logging"
)

func main() {
	new(environment.Configuration).Initialize()
	new(logging.Configuration).Initialize()
	new(cache_redis.Configuration).Initialize()

}
