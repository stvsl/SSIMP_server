package utils

import (
	"crypto/rand"
	"errors"
)

// GetRandom 生成随机数(正整数)
func GetRandom(leng int) (int, error) {
	if leng <= 0 {
		return 0, errors.New("长度非法")
	}
	b := make([]byte, leng)
	_, err := rand.Read(b)
	if err != nil {
		return 0, err
	}
	return int(b[0]), nil
}

// 生成指定范围的随机数
func GetRandomRange(min, max int) (int, error) {
	if min > max {
		return 0, errors.New("范围非法")
	}
	if min == max {
		return min, nil
	}
	// 使用 crypto/rand 生成随机小数，随后乘以范围，再加上最小值
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		return 0, err
	}
	r := int(b[0])%max + min
	return r, nil
}

// GetRandomString 生成随机字符串
// leng 随机字符串的长度
func GetRandomString(leng int) (string, error) {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if leng < 0 {
		return "", errors.New("长度非法")
	}
	// 使用getRandomRange生成随机数，然后取出对应的字符
	var result string
	for i := 0; i < leng; i++ {
		r, err := GetRandomRange(0, len(str)-1)
		if err != nil {
			return "", err
		}
		result += string(str[r])
	}
	return result, nil
}
