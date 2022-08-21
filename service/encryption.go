package service

import (
	"github.com/gin-gonic/gin"
	"stvsljl.com/SSIMP/redis"
	"stvsljl.com/SSIMP/security"
)

func Rsapubkey(c *gin.Context) {
	c.JSON(200, gin.H{
		"pubkey": security.SERVER_RSA.PUBLIC_KEY,
	})
}

func ClientRsakey(c *gin.Context) {
	feature := c.PostForm("feature")
	clientPubKey := c.PostForm("pubkey")
	dFeature, err1 := security.RsaDecrypt([]byte(feature))
	dClientPubKey, err2 := security.RsaDecrypt([]byte(clientPubKey))
	if err1 != nil || err2 != nil {
		Code.SE400(c)
	}
	awc := redis.AESWaitClient{
		Feature:      string(dFeature),
		ClientPubKey: string(dClientPubKey),
	}
	awc.WriteToRedis()
}
