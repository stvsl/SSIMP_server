package security

import (
	"crypto/rsa"
	"fmt"
	"time"

	"stvsljl.com/SSIMP/utils"
)

// 主RSA密钥
type Security struct {
	// 私钥
	PRIVATE_KEY       []byte
	PRIVATE_KEY_BYTES *rsa.PrivateKey
	// 公钥
	PUBLIC_KEY       []byte
	PUBLIC_KEY_BYTES *rsa.PublicKey
}

var SERVER_RSA Security
var SERVER_RSA_LAST Security

// 服务器通讯RSA初始化
func Init() {
	GenerateLocalRsaKey()
	SERVER_RSA_LAST = SERVER_RSA
	go AutoUpdate()
}

func AutoUpdate() {
	tracker := time.NewTicker(time.Minute * time.Duration(utils.GetSecurityConfig().RSAUpdateLifecycle))
	defer tracker.Stop()
	for t := range tracker.C {
		fmt.Println("RSA Update Time: ", t)
		RsaUpdate()
		tracker.Reset(time.Minute * time.Duration(utils.GetSecurityConfig().RSAUpdateLifecycle))
	}
}

// 服务器Rsa密钥更新
func RsaUpdate() {
	SERVER_RSA_LAST = SERVER_RSA
	GenerateLocalRsaKey()
}

// 单向通信主动加密
func RsaEncrypt(data []byte) ([]byte, error) {
	return Encrypt(data, SERVER_RSA.PUBLIC_KEY)
}

// 单向通信主动解密
func RsaDecrypt(data []byte) ([]byte, error) {
	data, err := Decrypt(data, SERVER_RSA.PRIVATE_KEY)
	if err != nil {
		data, err = Decrypt(data, SERVER_RSA_LAST.PRIVATE_KEY)
	}
	return data, err
}
