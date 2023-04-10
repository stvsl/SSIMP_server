package service

import (
	"github.com/gin-gonic/gin"
	"stvsljl.com/SSIMP/cos"
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
	// 对象存储初始化
	cos.Init()
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
	router.GET("/api/encryption/rsapubkey", Rsapubkey)                                // 获取服务端公钥
	router.POST("/api/encryption/rsatoaes", ClientRsakey)                             // 获取客户端公钥
	router.POST("/api/account/admin/login", AdminLogin)                               // 管理员登录
	router.POST("/api/account/employee/login", EmployeeLogin)                         // 员工登录
	router.POST("/api/account/employee/list", AuthMiddleware(), EmployeeList)         // 获取员工信息列表
	router.POST("/api/account/employee/add", AuthMiddleware(), EmployeeAdd)           // 添加员工信息
	router.POST("/api/account/employee/update", AuthMiddleware(), EmployeeUpdate)     // 更新员工信息
	router.POST("/api/account/employee/delete", AuthMiddleware(), EmployeeDelete)     // 删除员工信息
	router.POST("/api/employee/tasklist", AuthMiddleware(), EmployeeTaskList)         // 获取员工任务列表
	router.POST("/api/article/list", ArticleList)                                     // 获取文章列表
	router.POST("/api/article/recommendlist", ArticleRecommendList)                   // 获取推荐文章列表
	router.POST("/api/article/carousel", ArticleCarousel)                             // 获取轮播图文章列表
	router.POST("/api/article/publiclist", ArticlePublicList)                         // 获取公开文章列表
	router.POST("/api/article/detail", ArticleDetail)                                 // 获取文章详情
	router.POST("/api/article/add", AuthMiddleware(), ArticleAdd)                     // 添加文章
	router.POST("/api/article/update", AuthMiddleware(), ArticleUpdate)               // 更新文章
	router.POST("/api/article/search", ArticleSearch)                                 // 搜索文章
	router.GET("/api/article/tonocarousel", AuthMiddleware(), ArticleToNoCarousel)    // 将文章转为非轮播图
	router.GET("/api/article/delete", AuthMiddleware(), ArticleDelete)                // 删除文章
	router.POST("/api/taskset/list", TaskSetList)                                     // 获取任务集列表
	router.GET("/api/taskset/info", TaskSetInfo)                                      // 获取任务集信息
	router.POST("/api/taskset/add", AuthMiddleware(), TaskSetAdd)                     // 添加任务集
	router.POST("/api/taskset/update", AuthMiddleware(), TaskSetUpdate)               // 更新任务集
	router.GET("/api/taskset/delete", AuthMiddleware(), TaskSetDelete)                // 删除任务集
	router.POST("/api/task/employertasklist", AuthMiddleware(), EmployerTaskList)     // 获取雇员任务列表
	router.POST("/api/task/employertaskdelete", AuthMiddleware(), EmployerTaskDelete) // 删除雇员任务
	router.POST("/api/task/employertaskadd", AuthMiddleware(), EmployerTaskAdd)       // 添加雇员任务
	router.POST("/api/employee/task/list", EmployerTaskListFull)                      // 雇员任务列表(员工端)
	router.POST("/api/employee/task/status", EmployerTaskStatus)                      // 雇员任务状态(员工端)
	router.POST("/api/employee/task/sign", EmployerTaskSign)                          // 雇员签到(员工端)
	router.POST("/api/employee/task/posupload", EmployerTaskSposUpload)               // 雇员上传位置(员工端)
	router.POST("/api/employee/task/finish", EmployerTaskFinish)                      // 雇员完成任务(员工端)
	router.POST("/api/employee/info", EmployeeInfo)                                   // 获取员工信息(员工端)
	router.POST("/api/attendance/list/day", AttendanceListDay)                        // 获取考勤日期集合
	router.POST("/api/attendance/list", AuthMiddleware(), AttendanceList)             // 获取考勤列表
	router.POST("/api/attendance/record", AuthMiddleware(), AttendanceRecord)         // 获取考勤记录
	router.POST("/api/employee/updatepasswd", EmployeeUpdatePasswd)                   // 修改员工密码(员工端)
	router.POST("/api/employee/attendace/info", AttendaceInfo)                        // 修改员工信息(员工端)
	router.POST("/api/feedback/list", FeedbackList)                                   // 查询反馈信息列表(指定员工)
	router.POST("/api/feedback/list/all", AuthMiddleware(), FeedbackListAll)          // 查询反馈信息列表(所有员工)
	router.POST("/api/feedback/add", FeedbackAdd)                                     // 添加反馈信息
	router.POST("/api/feedback/set/orange", AuthMiddleware(), FeedbackOrange)         // 设置反馈信息为已接受
	router.POST("/api/feedback/set/accept", AuthMiddleware(), FeedbackAccept)         // 设置反馈信息为已接受
	router.POST("/api/feedback/set/solved", AuthMiddleware(), FeedbackSolved)         // 设置反馈信息为已完成
	router.POST("/api/feedback/set/doing", AuthMiddleware(), FeedbackDoing)           // 设置反馈信息为未完成
	router.POST("/api/feedback/set/reject", AuthMiddleware(), FeedbackReject)         // 拒绝反馈信为废弃
	router.POST("/api/feedback/set/delegate", AuthMiddleware(), FeedbackDelegate)     // 设置反馈信息的委派人
	router.POST("/api/feedback/delete", AuthMiddleware(), FeedbackDelete)             // 删除反馈信息
	router.POST("/api/data/analysis/global", DataAnalysisGlobal)                      // 获取全局数据分析
	router.POST("/api/data/analysis/employee", DataAnalysisEmployee)                  // 获取员工数据分析
	router.GET("/api/data/web/visit", WebVisit)                                       // 增加文章访问量
	router.Run(":6521")
}
