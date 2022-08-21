package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"stvsljl.com/SSIMP/utils"
)

var rdb *redis.Client

func Init() error {
	config := utils.GetRedisConfig()
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Pwd,      // 密码
		DB:       0,               // 数据库
		PoolSize: config.PoolSize, // 连接池大小
	})
	defer rdb.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		utils.Log.Error("redis连接失败", err)
	}
	return err
}

func GetRedis() *redis.Client {
	return rdb
}
