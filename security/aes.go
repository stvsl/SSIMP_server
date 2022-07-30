package security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AES加密
func AesEncrypt(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	var iv = []byte{'S', 'S', 'I', 'M', 'P', 'S', 'T', 'V', 'S', 'L', 'J', 'C', 'K', 'E', 'Y', 'S'}
	plaintext = PKCS7Padding(plaintext, 16)
	blockMode := cipher.NewCBCEncrypter(block, iv[:block.BlockSize()])
	crypted := make([]byte, len(plaintext))
	blockMode.CryptBlocks(crypted, plaintext)
	return crypted, nil
}

// AesDecrypt 解密函数
func AesDecrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	var iv = []byte{'S', 'S', 'I', 'M', 'P', 'S', 'T', 'V', 'S', 'L', 'J', 'C', 'K', 'E', 'Y', 'S'}
	blockMode := cipher.NewCBCDecrypter(block, iv[:block.BlockSize()])
	origData := make([]byte, len(ciphertext))
	blockMode.CryptBlocks(origData, ciphertext)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}
