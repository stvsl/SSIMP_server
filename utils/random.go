package utils

import (
	"crypto/rand"
	"math/big"
	"time"
)

// GetRandomString 生成随机数
func GetRandom(leng int, key string) int {
	var keyInt int
	for _, v := range key {
		if v >= 48 && v <= 57 {
			keyInt = keyInt*10 + int(v) - 48
		}
	}
	if keyInt == 0 {
		keyInt = time.Now().Nanosecond()
	}
	r2, _ := rand.Int(rand.Reader, big.NewInt(int64(keyInt)))
	rand := r2.Int64() % int64(leng)
	return int(rand)
}

// GetRandomString 生成随机字符串
func GetRandomString(leng int, key string) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < leng; i++ {
		rand := GetRandom(62, key)
		key = key + string(str[rand])
	}
	return key
}
