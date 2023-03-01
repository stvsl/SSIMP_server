package utils

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"io/ioutil"
	"os"
)

// base64转换成文件流
func Base64toFileStream(base64img string) (io.Reader, string, error) {
	// 解码base64字符串为二进制数据
	decoded, err := base64.StdEncoding.DecodeString(base64img)
	if err != nil {
		return nil, "", err
	}

	// 生成随机文件名
	filename, err := generateRandomFilename()
	if err != nil {
		return nil, filename, err
	}

	// 将二进制数据写入临时文件
	err = ioutil.WriteFile(filename, decoded, 0644)
	if err != nil {
		return nil, filename, err
	}

	// 打开临时文件并返回一个io.Reader
	file, err := os.Open(filename)
	if err != nil {
		return nil, filename, err
	}

	// 使用defer语句在函数返回时删除临时文件
	defer os.Remove(filename)

	return file, filename, nil
}

func generateRandomFilename() (string, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
