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
	// 允许跨域
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	/**********************
	 * 加载路由
	 **********************/
	// 通信加密相关
	router.GET("/api/encryption/rsapubkey", Rsapubkey)                            // 获取服务端公钥
	router.POST("/api/encryption/rsatoaes", ClientRsakey)                         // 获取客户端公钥
	router.POST("/api/account/admin/login", AdminLogin)                           // 管理员登录
	router.POST("/api/account/employee/list", AuthMiddleware(), EmployeeList)     // 获取员工信息列表
	router.POST("/api/account/employee/add", AuthMiddleware(), EmployeeAdd)       // 添加员工信息
	router.POST("/api/account/employee/update", AuthMiddleware(), EmployeeUpdate) // 更新员工信息
	router.POST("/api/account/employee/delete", AuthMiddleware(), EmployeeDelete) // 删除员工信息
	router.POST("/api/article/list", ArticleList)                                 // 获取文章列表
	router.POST("/api/article/detail", ArticleDetail)                             // 获取文章详情
	router.POST("/api/article/add", AuthMiddleware(), ArticleAdd)                 // 添加文章
	router.POST("/api/article/update", AuthMiddleware(), ArticleUpdate)           // 更新文章
	router.POST("/api/article/delete", AuthMiddleware(), ArticleDelete)           // 删除文章
	router.Run(":6521")
}
