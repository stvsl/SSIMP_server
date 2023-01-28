package service

import (
	"encoding/base64"
	"fmt"

	"github.com/gin-gonic/gin"
	"stvsljl.com/SSIMP/redis"
	"stvsljl.com/SSIMP/security"
	"stvsljl.com/SSIMP/utils"
)

func Rsapubkey(c *gin.Context) {
	fmt.Println("获取服务端公钥")
	c.JSON(200, gin.H{
		"pubkey": security.SERVER_RSA.GetPublicKey(),
	})
}

func ClientRsakey(c *gin.Context) {
	var clientRsaKey struct {
		Feature   string `json:"feature"`
		PublicKey string `json:"pubkey"`
		AESP      string `json:"aesp"`
	}
	err := c.ShouldBindJSON(&clientRsaKey)
	if err != nil {
		utils.Log.Error("客户端数据解析失败", err.Error())
		Code.SE400(c)
		return
	}
	// 数据从base64转回byte
	bFeature, _ := base64.StdEncoding.DecodeString(clientRsaKey.Feature)
	bPublicKey, _ := base64.StdEncoding.DecodeString(clientRsaKey.PublicKey)
	bAESP, _ := base64.StdEncoding.DecodeString(clientRsaKey.AESP)
	dFeature, err1 := security.RsaDecrypt(bFeature)
	dClientPubKey, err2 := security.RsaDecrypt(bPublicKey)
	dAESP, err3 := security.RsaDecrypt(bAESP)
	// fmt.Println("特征值" + string(dFeature))
	// fmt.Println("客户端公钥" + string(dClientPubKey))
	fmt.Println("AES密钥（部分）" + string(dAESP))
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("客户端公钥解析失败" + "1" + err1.Error() + "2" + err2.Error() + "3" + err3.Error())
		utils.Log.Error("客户端公钥解析失败" + err1.Error() + err2.Error() + err3.Error())
		Code.SE401(c)
	}
	aesp2, _ := utils.GetRandomString(64)
	// 计算dAESP和aesp2的异或值
	aesp := utils.Xor(string(dAESP), aesp2)
	aes := utils.Md5(aesp)
	// 暂时保存客户端公钥
	awc := redis.AESWaitClient{
		Feature: string(dFeature),
		AES:     string(aes),
	}
	// fmt.Println(awc)
	err = awc.WriteToRedis()
	if err != nil {
		utils.Log.Error("redis写入失败", err.Error())
		Code.SE610(c)
		return
	}
	fmt.Println("AES密钥2", aesp2)
	fmt.Println("AES密钥 " + string(aes))
	aesp2Byte := []byte(aesp2)
	// 将aesp2加密后返回
	aesp2byte, err := security.Encrypt(aesp2Byte, dClientPubKey)
	if err != nil {
		utils.Log.Error("客户端公钥加密失败", err.Error())
		Code.SE500(c)
	}
	// 转为base64
	base64Aesp2 := base64.StdEncoding.EncodeToString(aesp2byte)
	c.JSON(200, gin.H{
		"aesp2": base64Aesp2,
	})
}
