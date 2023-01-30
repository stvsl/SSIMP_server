package service

import (
	"github.com/gin-gonic/gin"
	"stvsljl.com/SSIMP/db"
	"stvsljl.com/SSIMP/redis"
	"stvsljl.com/SSIMP/security"
)

func AdminLogin(c *gin.Context) {
	// 从请球体中获取用户名和密码
	id := c.PostForm("id")
	passwd := c.PostForm("passwd")
	feature := c.PostForm("feature")
	// 从redis中获取用户信息
	aeswait := redis.AESWaitClient{Feature: feature}
	aes, err := aeswait.ReadAndRemoveFromRedis()
	if err != nil {
		Code.SE001(c)
		return
	}
	// 解密id和passwd(先转换成[]byte)
	idByte, err := security.AesDecrypt([]byte(id), []byte(aes))
	if err != nil {
		Code.SE002(c)
		return
	}
	passwdByte, err := security.AesDecrypt([]byte(passwd), []byte(aes))
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
		"msg":   "登录成功",
		"token": token,
	})
}
