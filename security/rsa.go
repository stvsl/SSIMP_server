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
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
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
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
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
	ciphertextByte, err := rsa.EncryptPKCS1v15(rand.Reader, pubInterface, origData)
	if err != nil {
		ciphertextByte = make([]byte, 0)
		for i := 0; i < len(origData); i += 245 {
			var end int
			if i+245 > len(origData) {
				end = len(origData)
			} else {
				end = i + 245
			}
			// 加密
			ciphertextByteTemp, err := rsa.EncryptPKCS1v15(rand.Reader, pubInterface, origData[i:end])
			if err != nil {
				return nil, err
			}
			ciphertextByte = append(ciphertextByte, ciphertextByteTemp...)
		}
		return ciphertextByte, nil
	}
	return ciphertextByte, nil
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
		log.Println("私钥解析错误")
		return nil, errors.New("私钥解析错误:" + err.Error())
	}
	// 判断数据长度是否超过单个切片的长度
	if len(ciphertext) > 256 {
		// 切片解密
		var result []byte
		for i := 0; i < len(ciphertext); i += 256 {
			// 判断切片是否超出长度
			if i+256 > len(ciphertext) {
				// 解密
				decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext[i:])
				if err != nil {
					return nil, err
				}
				result = append(result, decrypted...)
			} else {
				// 解密
				decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext[i:i+256])
				if err != nil {
					return nil, err
				}
				result = append(result, decrypted...)
			}
		}
		return result, nil
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
