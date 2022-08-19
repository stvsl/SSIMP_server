package service

import (
	"github.com/gin-gonic/gin"
	"stvsljl.com/SSIMP/db"
	"stvsljl.com/SSIMP/security"
)

func Start() {
	// 安全组件相关事务初始化
	security.Init()
	// 数据库事务初始化
	db.Connect()
	// 服务器服务启动
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/", func(c *gin.Context) {
		Code.SE001(c)
	})
	router.Run(":8080")
}
