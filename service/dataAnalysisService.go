package service

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"stvsljl.com/SSIMP/db"
	"stvsljl.com/SSIMP/redis"
)

type TaskInfo struct {
	// 任务完成率, 当日缺勤人数，已工作完成人数，未工作完成人数，工作中人数, 正在工作的人的位置数据

}

func DataAnalysisGlobal(c *gin.Context) {
	// 获取员工总数，文章总数，新增浏览量，任务完成率，网站访问量（7天），文章访问量（7天）
	// 当日缺勤人数，已工作完成人数，工作中人数，正在工作的人的位置数据

	dbconn := db.GetConn()
	// 获取员工总数
	var employeecount int64
	dbconn.Model(&db.Employer{}).Count(&employeecount)

	// 获取当日缺勤人数
	var absentCount int64
	if err := dbconn.
		Model(&db.Task{}).
		Select("COUNT(DISTINCT Task.employid) AS absent_count").
		Joins("LEFT JOIN Attendance ON Task.employid = Attendance.employid AND DATE(Attendance.starttime) = CURDATE()").
		Where("Attendance.employid IS NULL").
		Scan(&absentCount).
		Error; err != nil {
		Code.SE602(c)
	}

	var finishedCount int64
	if err := dbconn.
		Model(&db.Task{}).
		Select("COUNT(Task.employid) AS finished_count").
		Joins("LEFT JOIN Attendance ON Task.employid = Attendance.employid AND Task.tid = Attendance.tid AND DATE(Attendance.starttime) = CURDATE() AND Attendance.task_completion = '已完成'").
		Group("Task.employid").
		Having("COUNT(Task.tid) = COUNT(Attendance.tid)").
		Scan(&finishedCount).
		Error; err != nil {
		Code.SE602(c)
	}

	workingCount := employeecount - absentCount - finishedCount

	// 获取文章总数
	var articlecount int64
	dbconn.Model(&db.Article{}).Count(&articlecount)

	// 计算任务相关数据
	var weekTaskNum int64
	if err := dbconn.Model(&db.Task{}).
		Select("SUM(ts.cycle)").
		Joins("INNER JOIN TaskSet ts ON Task.tid = ts.tid").
		Where("ts.cycle BETWEEN 1 AND 7").
		Scan(&weekTaskNum).
		Error; err != nil {
		Code.SE602(c)
	}

	var completedTasks int64
	if err := dbconn.Model(&db.Attendance{}).
		Select("COUNT(*)").
		Where("task_completion = ?", "已完成").
		Where("starttime BETWEEN DATE_SUB(NOW(), INTERVAL 1 WEEK) AND NOW()").
		Scan(&completedTasks).
		Error; err != nil {
		Code.SE602(c)
		return
	}

	// 获取员工工作位置数据
	var inspectionTracks []string
	if err := dbconn.
		Model(&db.Attendance{}).
		Select("DISTINCT inspection_track").
		Joins("LEFT JOIN Task ON Attendance.employid = Task.employid AND Attendance.tid = Task.tid").
		Where("Attendance.task_completion != '已完成'").
		Where("DATE(Attendance.starttime) = CURDATE()").
		Where("NOT EXISTS (SELECT 1 FROM Attendance a2 WHERE Attendance.employid = a2.employid AND a2.tid = Attendance.tid AND a2.starttime > Attendance.starttime AND DATE(a2.starttime) = CURDATE())").
		Pluck("inspection_track", &inspectionTracks).
		Error; err != nil {
		Code.SE602(c)
	}
	var locations []map[string]float64
	for i := 0; i < len(inspectionTracks); i++ {
		err := json.Unmarshal([]byte(inspectionTracks[i]), &locations)
		if err != nil {
			fmt.Printf("Failed to unmarshal JSON: %v\n", err)
			continue // 继续迭代下一个元素
		}

		if len(locations) > 0 {
			inspectionTracks[i] = fmt.Sprintf(`{"lat": %f, "lng": %f}`, locations[len(locations)-1]["lat"], locations[len(locations)-1]["lng"])
		}
	}

	completionRate := float64(completedTasks) / float64(weekTaskNum) * 100.0
	data := redis.GetRunningDataStruct()
	data.ReadAllFromRedis()

	var Info struct {
		EmployeeCount             int64    `json:"employee_count"`
		ArticleCount              int64    `json:"article_count"`
		CompletionRate            float64  `json:"completion_rate"`
		NewViewCount              int      `json:"new_view_count"`
		SevenDaysViewCount        []int    `json:"seven_days_view_count"`
		SevenDaysArticleViewCount []int    `json:"seven_days_article_view_count"`
		AbsentCount               int64    `json:"absent_count"`
		FinishedCount             int64    `json:"finished_count"`
		WorkingCount              int64    `json:"working_count"`
		WorkingLocations          []string `json:"working_locations"`
	}
	Info.EmployeeCount = employeecount
	Info.ArticleCount = articlecount
	Info.CompletionRate = completionRate
	Info.NewViewCount = data.NewViewCount
	Info.SevenDaysViewCount = data.SevenDaysViewCount
	Info.SevenDaysArticleViewCount = data.SevenDaysArticleViewCount
	Info.AbsentCount = absentCount
	Info.FinishedCount = finishedCount
	Info.WorkingCount = workingCount
	Info.WorkingLocations = inspectionTracks
	infojson, _ := json.Marshal(Info)
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "success",
		"data": string(infojson),
	})

}

func WebVisit(c *gin.Context) {
	redis.SetUpdateWebViewCount()
}

func DataAnalysisEmployee(c *gin.Context) {
	eid := c.Query("eid")
	// // 获取一周出勤情况（一周内签到时间，一周内签退时间），一个月内缺勤天数，执勤天数
	dbconn := db.GetConn()
	var result []map[string]interface{}
	err := dbconn.Table("Attendance").
		Select("DATE(startTime) AS `日期`, COALESCE(MIN(startTime), CAST(CONCAT(DATE(startTime), ' 23:59:59') AS DATETIME)) AS `签到时间`, COALESCE(MAX(endTime), CAST(CONCAT(DATE(startTime), ' 23:59:59') AS DATETIME)) AS `签退时间`").
		Where("employid = ? AND startTime >= DATE_SUB(NOW(), INTERVAL 1 WEEK)", eid).
		Group("DATE(startTime)").
		Order("DATE(startTime) DESC").
		Scan(&result).Error
	if err != nil {
		Code.SE602(c)
		return
	}
	var info struct {
		Zq int `json:"zq"`
		Qq int `json:"qq"`
	}
	err = dbconn.Table("Attendance").
		Select("COUNT(DISTINCT DATE(startTime)) AS `zq`, (SELECT DAY(LAST_DAY(CURDATE()))) - COUNT(DISTINCT DATE(startTime)) AS `qq`").
		Where("employid = ? AND startTime >= DATE_SUB(CURDATE(), INTERVAL 1 MONTH)", eid).
		Scan(&info).Error
	if err != nil {
		Code.SE602(c)
	}

	finalresult := make(map[string]interface{})
	finalresult["data"] = result
	finalresult["info"] = info

	resultjson, err := json.Marshal(finalresult)
	if err != nil {
		Code.SE602(c)
		return
	}
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "success",
		"data": string(resultjson),
	})
}
