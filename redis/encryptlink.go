package redis

import (
	"context"
	"time"

	"stvsljl.com/SSIMP/utils"
)

type AESWaitClient struct {
	Feature string
	AES     string
}

func (a *AESWaitClient) WriteToRedis() error {
	status, err := rdb.HSet(context.Background(), "AESWaitClient", a.Feature, a.AES).Result()
	if err != nil || status == 0 {
		utils.Log.Error("redis写入失败", err)
		return err
	}
	// 设置过期时间
	rdb.Expire(context.Background(), "AESWaitClient", time.Minute*5)
	return nil
}

func (a *AESWaitClient) ReadAndRemoveFromRedis() (string, error) {
	aes, err := rdb.HGet(context.Background(), "AESWaitClient", a.Feature).Result()
	if err != nil {
		utils.Log.Error("键值不存在", err)
		return "", err
	}
	a.AES = aes
	status, err := rdb.HDel(context.Background(), "AESWaitClient", a.Feature).Result()
	if err != nil || status == 0 {
		utils.Log.Error("redis删除失败：", err)
		return "", err
	}
	return aes, err
}
