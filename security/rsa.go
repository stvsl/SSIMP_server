package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
)

// 服务器RSA私钥
var RSA_PRIVATE_LOCAL string
var RSA_PRIVATE_LOCAL_BYTES *rsa.PrivateKey

// 服务器RSA公钥
var RSA_PUBLIC_LOCAL string
var RSA_PUBLIC_LOCAL_BYTES *rsa.PublicKey

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
	RSA_PRIVATE_LOCAL = string(privateKeyPem)
	RSA_PRIVATE_LOCAL_BYTES = privateKey
	RSA_PUBLIC_LOCAL = string(publicKeyPem)
	RSA_PUBLIC_LOCAL_BYTES = publicKey
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

// 返回服务器公钥字符串
func GetPublicKey() string {
	return RSA_PUBLIC_LOCAL
}
