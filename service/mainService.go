package service

import (
	"github.com/gin-gonic/gin"
	"stvsljl.com/SSIMP/db"
	"stvsljl.com/SSIMP/redis"
	"stvsljl.com/SSIMP/security"
	"stvsljl.com/SSIMP/utils"
)

func Start() {
	// 安全组件相关事务初始化
	security.Init()
	// 数据库事务初始化
	db.Connect()
	// 日志组件初始化
	utils.Log.Init()
	// redis初始化
	redis.Init()
	// 服务器服务启动
	router := gin.Default()
	router.SetTrustedProxies(nil)
	/**********************
	 * 加载路由
	 **********************/
	// 通信加密相关
	router.GET("/api/encryption/rsapubkey", Rsapubkey)    // 获取服务端公钥
	router.POST("/api/encryption/rsatoaes", ClientRsakey) // 获取客户端公钥
	router.POST("/api/account/admin/login", AdminLogin)   // 管理员登录
	router.Run(":6521")
}
