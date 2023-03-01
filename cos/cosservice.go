package cos

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
	"stvsljl.com/SSIMP/utils"
)

var CosClient *cos.Client
var BucketName = ""

func Init() {
	CosClient = NewCos()
	BucketName = utils.GetCosConfig().Bucket
}

func NewCos() *cos.Client {
	u, _ := url.Parse(utils.GetCosConfig().Domain)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  utils.GetCosConfig().SecretId,
			SecretKey: utils.GetCosConfig().SecretKey,
		},
	})
	return client
}

func GetCosClient() *cos.Client {
	return CosClient
}

// 上传文件
func UploadFile(base64img string) (string, error) {
	ior, filename, err := utils.Base64toFileStream(base64img)
	if err != nil {
		return "文件转换失败", err
	}
	// 获取年月日，自动创建文件夹，如果文件夹已存在则不会创建
	path := time.Now().Format("2006/01/02/")
	// 上传文件
	_, err = CosClient.Object.Put(context.Background(), path+filename, ior, nil)
	if err != nil {
		return "上传失败", err
	}
	return utils.GetCosConfig().Domain + "/" + path + filename, nil
}

// 删除文件
func DeleteFile(key string) error {
	_, err := CosClient.Object.Delete(context.Background(), key, nil)
	return err
}
