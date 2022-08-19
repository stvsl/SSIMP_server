package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
)

// 生成RSA密钥对
func GenerateRsaKey() (string, string, error) {
	// 生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return "", "", err
	}
	// 生成公钥
	publicKey := &privateKey.PublicKey
	// 生成pem格式的密钥
	privateKeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)
	publicKeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(publicKey),
		},
	)
	return string(privateKeyPem), string(publicKeyPem), nil
}

// 生成服务器本地密钥对
func GenerateLocalRsaKey() {
	// 生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return
	}
	// 生成公钥
	publicKey := &privateKey.PublicKey
	// 生成pem格式的密钥
	privateKeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)
	publicKeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(publicKey),
		},
	)
	SERVER_RSA.PRIVATE_KEY = privateKeyPem
	SERVER_RSA.PUBLIC_KEY = publicKeyPem
	SERVER_RSA.PRIVATE_KEY_BYTES = privateKey
	SERVER_RSA.PUBLIC_KEY_BYTES = publicKey
}

// 加密算法（使用PKCS1密钥）
func Encrypt(origData []byte, publickey []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(publickey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 加密
	return rsa.EncryptPKCS1v15(rand.Reader, pubInterface, origData)
}

// 解密
func Decrypt(ciphertext []byte, privateKey []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("私钥错误")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Println("!!!")
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
