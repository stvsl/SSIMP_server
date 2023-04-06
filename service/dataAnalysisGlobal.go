package service

import (
	"github.com/gin-gonic/gin"
	"stvsljl.com/SSIMP/db"
)

func DataAnalysisGlobal(c *gin.Context) {
	// 获取员工总数，文章总数，新增浏览量，任务完成率，网站访问量（7天），文章访问量（7天）
	// 当日缺勤人数，已工作完成人数，未工作完成人数，工作中人数，正在工作的人的位置数据

	dbconn := db.GetConn()
	// 获取员工总数
	var employeecount int64
	dbconn.Model(&db.Employer{}).Count(&employeecount)
	// 获取文章总数
	var articlecount int64
	dbconn.Model(&db.Article{}).Count(&articlecount)
	// 获取新增浏览量

}
