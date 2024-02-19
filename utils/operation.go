package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// 异或运算
func Xor(a, b string) string {
	// 保证a,b长度相同
	if len(a) > len(b) {
		b = b + a[len(b):]
	}
	if len(a) < len(b) {
		a = a + b[len(a):]
	}
	// 异或运算
	var result string
	for i := 0; i < len(a); i++ {
		result = result + string(a[i]^b[i])
	}
	return result
}

// MD5
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
