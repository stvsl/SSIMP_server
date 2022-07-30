package redis

import (
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func Start() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "120.48.17.34,5210",
		Password: "stvsl2060", // 密码
		DB:       0,           // 数据库
		PoolSize: 20,          // 连接池大小
	})
}
