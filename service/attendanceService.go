package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"stvsljl.com/SSIMP/db"
)

func EmployerTaskStatus(c *gin.Context) {
	var info struct {
		Employid string `json:"eid"`
		Tid      int64  `json:"tid"`
	}
	err := c.BindJSON(&info)
	if err != nil {
		Code.SE400(c)
		return
	}
	dbconn := db.GetConn()
	attendance := db.Attendance{}
	// 查找数据库中是否有当天的该员工的签到记录
	// 获得当天凌晨的时间以及第二天凌晨的时间
	now := time.Now()
	year, month, day := now.Date()
	zero := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	zero2 := zero.AddDate(0, 0, 1)
	// 查询数据库
	dbconn.Model(&attendance).Where("employid = ? and tid = ? and startTime >= ? and startTime < ?", info.Employid, info.Tid, zero, zero2).First(&attendance)
	// 如果没有找到
	if attendance.InspectionTrack == "" {
		c.JSON(200, gin.H{
			"code": "SE200",
			"msg":  "未签到",
		})
	} else {
		c.JSON(200, gin.H{
			"code": "SE200",
			"msg":  attendance.TaskCompletion,
		})
	}
}
func AttendaceInfo(c *gin.Context) {
	var info struct {
		Employid string `json:"eid"`
		Tid      int64  `json:"tid"`
		Datestr  string `json:"date"`
	}
	err := c.BindJSON(&info)
	if err != nil {
		Code.SE400(c)
		return
	}
	//datestr转换为date
	date, err := time.Parse("2006-01-02", info.Datestr)
	if err != nil {
		Code.SE400(c)
		return
	}
	dbconn := db.GetConn()
	attendance := db.Attendance{}
	// 查找数据库中是否有当天的该员工的签到记录
	// 获取date当天凌晨的时间以及第二天凌晨的时间
	year, month, day := date.Date()
	zero := time.Date(year, month, day, 0, 0, 0, 0, date.Location())
	zero2 := zero.AddDate(0, 0, 1)

	// 查询数据库
	dbconn.Model(&attendance).Where("employid = ? and tid = ? and startTime >= ? and startTime < ?", info.Employid, info.Tid, zero, zero2).First(&attendance)
	// 如果没有找到
	if attendance.InspectionTrack == "" {
		Code.SE401(c)
	} else {
		attendancesjson, _ := json.Marshal(attendance)
		c.JSON(200, gin.H{
			"code": "SE200",
			"msg":  "查询成功",
			"data": string(attendancesjson),
		})
	}
}
func EmployerTaskSign(c *gin.Context) {
	var info struct {
		Employid string `json:"eid"`
		Tid      int64  `json:"tid"`
	}
	err := c.BindJSON(&info)
	if err != nil {
		Code.SE400(c)
		return
	}
	dbconn := db.GetConn()
	attendance := db.Attendance{}
	// 查找数据库中是否有当天的该员工的签到记录
	// 获得当天凌晨的时间以及第二天凌晨的时间
	now := time.Now()
	year, month, day := now.Date()
	zero := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	zero2 := zero.AddDate(0, 0, 1)
	// 查询数据库
	if dbconn.Model(&attendance).Where("employid = ? and tid = ? and startTime >= ? and startTime <= ?",
		info.Employid, info.Tid, zero, zero2).First(&attendance).Error != nil {
		if attendance.InspectionTrack == "" {

			fmt.Println("没有该员工的签到记录")
			attendance.Employid = info.Employid
			attendance.Tid = info.Tid
			attendance.StartTime = time.Now()
			attendance.TaskCompletion = "未完成"
			attendance.InspectionTrack = "{}"
			// 将签到记录插入数据库
			if dbconn.Create(&attendance).Error != nil {
				Code.SE600(c)
				return
			}
			c.JSON(200, gin.H{
				"code": "SE200",
				"msg":  "签到成功",
			})
		}
	} else {
		attendancesjson, _ := json.Marshal(attendance)
		c.JSON(200, gin.H{
			"code": "SE200",
			"msg":  "已签到",
			"data": string(attendancesjson),
		})
	}

}

func EmployerTaskSposUpload(c *gin.Context) {
	var info struct {
		Employid string `json:"eid"`
		Tid      int64  `json:"tid"`
		Track    string `json:"track"`
	}
	err := c.BindJSON(&info)
	if err != nil {
		Code.SE400(c)
		return
	}
	dbconn := db.GetConn()
	// 查找数据库中是否有当天的该员工的签到记录
	// 获得当天凌晨的时间以及第二天凌晨的时间
	now := time.Now()
	year, month, day := now.Date()
	zero := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	zero2 := zero.AddDate(0, 0, 1)
	// 查询数据库
	attendance := db.Attendance{}
	dbconn.Model(&attendance).Where("employid = ? and tid = ? and startTime >= ? and startTime <= ?",
		info.Employid, info.Tid, zero, zero2).First(&attendance)
	// 判断原有的轨迹是否为空
	if attendance.InspectionTrack == "{}" {
		// 如果为空则直接将新的轨迹插入数据库
		dbconn.Model(&attendance).Update("inspection_track", info.Track).Where("employid = ? and tid = ? and startTime >= ? and startTime <= ?",
			info.Employid, info.Tid, zero, zero2)
		c.JSON(200, gin.H{
			"code": "SE200",
			"msg":  "上传成功",
		})
		return
	}
	var posinfoold []struct {
		Posli float64 `json:"lat"`
		Poslo float64 `json:"lng"`
	}
	err = json.Unmarshal([]byte(attendance.InspectionTrack), &posinfoold)
	if err != nil {
		fmt.Println("server", err)
		Code.SE401(c)
		return
	}
	var posinfo []struct {
		Posli float64 `json:"lat"`
		Poslo float64 `json:"lng"`
	}
	err = json.Unmarshal([]byte(info.Track), &posinfo)
	if err != nil {
		fmt.Println("app", err)
		Code.SE401(c)
		return
	}
	posinfo = append(posinfoold, posinfo...)
	posinfojson, _ := json.Marshal(posinfo)
	// 更新数据库
	dbconn.Model(&attendance).Update("inspection_track", string(posinfojson)).Where("employid = ? and tid = ? and startTime >= ? and startTime <= ?")
	// 判断时长是否超过了任务时长
	taskset := db.TaskSet{}
	dbconn.Model(&taskset).Where("tid = ?", info.Tid).First(&taskset)
	// 计算轨迹的时长
	// attendance的开始时间加上 taskset.Duration 小时与当前时间比较

	TaskDuration := attendance.StartTime.Add(time.Hour * time.Duration(taskset.Duration))
	if TaskDuration.Before(time.Now()) {
		dbconn.Model(&attendance).Update("task_completion", "已完成").Where("employid = ? and tid = ? and startTime >= ? and startTime <= ?",
			info.Employid, info.Tid, zero, zero2)
		c.JSON(200, gin.H{
			"code": "SE200",
			"msg":  "任务已完成",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "轨迹上传成功",
	})

}

func EmployerTaskFinish(c *gin.Context) {
	var info struct {
		Employid string `json:"eid"`
		Tid      int64  `json:"tid"`
	}
	err := c.BindJSON(&info)
	if err != nil {
		Code.SE400(c)
		return
	}
	dbconn := db.GetConn()
	// 查找数据库中是否有当天的该员工的签到记录
	// 获得当天凌晨的时间以及第二天凌晨的时间
	now := time.Now()
	year, month, day := now.Date()
	zero := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	zero2 := zero.AddDate(0, 0, 1)
	// 查询数据库
	attendance := db.Attendance{}
	dbconn.Model(&attendance).Where("employid = ? and tid = ? and startTime >= ? and startTime <= ?",
		info.Employid, info.Tid, zero, zero2).First(&attendance)
	// 添加结束时间
	attendance.EndTime = time.Now()
	// 更新数据库
	dbconn.Model(&attendance).Update("taskCompletion", "已完成").Where("employid = ? and tid = ? and startTime >= ? and startTime <= ?")
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "任务完成",
	})

}

func AttendanceListDay(c *gin.Context) {
	// 获取eid,返回该员工的签到日期列表
	eid := c.Query("eid")
	dbconn := db.GetConn()

	// 查询数据库
	attendances := []db.Attendance{}
	dbconn.Model(&attendances).Where("employid = ?", eid).Find(&attendances)
	// 只保留开始时间的日期
	date := []string{}
	for _, v := range attendances {
		date = append(date, v.StartTime.Format("2006-01-02"))
	}
	// 去重
	date = removeDuplicateElement(date)
	datejson, _ := json.Marshal(date)
	// 返回
	fmt.Println(date)
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "获取成功",
		"data": string(datejson),
	})
}

func AttendanceList(c *gin.Context) {
	// 获取eid,返回该员工的签到日期列表
	eid := c.Query("eid")
	date := c.Query("date")
	fmt.Println(eid, date)
	dbconn := db.GetConn()
	// date格式为 2006-01-02,转换为时间
	zero, _ := time.Parse("2006-01-02", date)
	zero2 := zero.AddDate(0, 0, 1)

	// 查询数据库
	attendances := []db.Attendance{}
	dbconn.Model(&attendances).Select(`employid`, `tid`, `startTime`, `endTime`, `task_completion`).Where("employid = ? and startTime >= ? and startTime <= ?",
		eid, zero, zero2).Find(&attendances)
	fmt.Println(attendances)
	json, _ := json.Marshal(attendances)
	// 返回
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "获取成功",
		"data": string(json),
	})

}

func AttendanceRecord(c *gin.Context) {
	// 获取eid,返回该员工的签到日期列表
	eid := c.Query("eid")
	tid := c.Query("tid")
	date := c.Query("date")
	dbconn := db.GetConn()
	// date格式为 2006-01-02,转换为时间
	zero, _ := time.Parse("2006-01-02", date)
	zero2 := zero.AddDate(0, 0, 1)

	// 查询数据库
	attendance := db.Attendance{}
	dbconn.Model(&attendance).Where("employid = ? and tid = ? and startTime >= ? and startTime <= ?",
		eid, tid, zero, zero2).First(&attendance)
	json, _ := json.Marshal(attendance)
	// 返回
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "获取成功",
		"data": string(json),
	})

}

func removeDuplicateElement(arr []string) []string {
	result := make([]string, 0, len(arr))
	temp := map[string]struct{}{}
	for _, item := range arr {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
