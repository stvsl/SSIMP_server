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
	var posinfo []struct {
		Posli float64 `json:"posli"`
		Poslo float64 `json:"poslo"`
	}
	err = json.Unmarshal([]byte(info.Track), &posinfo)
	if err != nil {
		Code.SE401(c)
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
	// 追加轨迹数据
	var posinfoold []struct {
		Posli float64 `json:"posli"`
		Poslo float64 `json:"poslo"`
	}
	err = json.Unmarshal([]byte(attendance.InspectionTrack), &posinfoold)
	if err != nil {
		Code.SE401(c)
		return
	}
	posinfo = append(posinfoold, posinfo...)
	posinfojson, _ := json.Marshal(posinfo)
	// 更新数据库
	dbconn.Model(&attendance).Update("inspectionTrack", string(posinfojson)).Where("employid = ? and tid = ? and startTime >= ? and startTime <= ?")
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
