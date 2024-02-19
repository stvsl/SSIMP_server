package service

import (
	"encoding/base64"

	"github.com/gin-gonic/gin"
	"stvsljl.com/SSIMP/db"
	"stvsljl.com/SSIMP/redis"
	"stvsljl.com/SSIMP/security"
)

func AdminLogin(c *gin.Context) {
	// 从请球体中获取用户名和密码
	var json struct {
		ID      string `json:"id"`
		Passwd  string `json:"passwd"`
		Feature string `json:"feature"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		Code.SE400(c)
		return
	}
	jid, err1 := base64.StdEncoding.DecodeString(json.ID)
	jpasswd, err2 := base64.StdEncoding.DecodeString(json.Passwd)
	if err1 != nil || err2 != nil {
		Code.SE401(c)
		return
	}
	json.ID = string(jid)
	json.Passwd = string(jpasswd)
	// 从redis中获取用户信息
	aeswait := redis.AESWaitClient{Feature: json.Feature}
	aes, err := aeswait.ReadAndRemoveFromRedis()
	if err != nil {
		Code.SE001(c)
		return
	}
	// 解密id和passwd(先转换成[]byte)
	idByte, err := security.AesDecrypt([]byte(json.ID), []byte(aes))
	if err != nil {
		Code.SE002(c)
		return
	}
	passwdByte, err := security.AesDecrypt([]byte(json.Passwd), []byte(aes))
	if err != nil {
		Code.SE002(c)
		return
	}
	// 从数据库中获取用户信息
	mgr := db.AdminMgr(db.GetConn())
	admin, err := mgr.GetFromAdminID(string(idByte))
	if err != nil {
		Code.SE406(c)
		return
	}
	// 验证密码
	if admin.Passwd != string(passwdByte) {
		Code.SE406(c)
		return
	}
	// 生成token
	token, _ := GenToken(aes)
	// 返回token
	c.JSON(200, gin.H{
		"code":  "SE200",
		"msg":   "登录成功",
		"token": token,
	})
}
