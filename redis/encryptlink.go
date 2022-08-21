package redis

import (
	"context"
	"time"

	"stvsljl.com/SSIMP/utils"
)

type AESWaitClient struct {
	Feature      string
	ClientPubKey string
}

func (a *AESWaitClient) WriteToRedis() error {
	err := rdb.HSet(context.Background(), "AESWaitClient", a.Feature, a.ClientPubKey, time.Duration(120)*time.Second).Err()
	if err != nil {
		utils.Log.Error("redis存储失败：", err)
	}
	return err
}

func (a *AESWaitClient) ReadAndRemoveFromRedis() (string, error) {
	clientPubKey, err := rdb.HGet(context.Background(), "AESWaitClient", a.Feature).Result()
	if err != nil {
		utils.Log.Error("键值不存在", err)
		return "", err
	}
	a.ClientPubKey = clientPubKey
	status, err := rdb.HDel(context.Background(), "AESWaitClient", a.Feature).Result()
	if err != nil || status == 0 {
		utils.Log.Error("redis删除失败：", err)
		return "", err
	}
	return clientPubKey, err
}
