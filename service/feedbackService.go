package service

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"stvsljl.com/SSIMP/db"
)

func FeedbackList(c *gin.Context) {
	eid := c.Query("eid")
	if eid == "" {
		Code.SE401(c)
		return
	}
	var feedbacks []db.Feedback
	dbconn := db.GetConn()
	if dbconn.Table("Feedback").Where("sponsor = ?", eid).Find(&feedbacks).Error != nil {
		Code.SE500(c)
		return
	}
	feedbacksJson, _ := json.Marshal(feedbacks)
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "success",
		"data": string(feedbacksJson),
	})

}

func FeedbackListAll(c *gin.Context) {
	var feedbacks []db.Feedback
	dbconn := db.GetConn()
	dbconn.Table("Feedback").Find(&feedbacks)

	feedbacksJson, _ := json.Marshal(feedbacks)
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "success",
		"data": string(feedbacksJson),
	})
}

func FeedbackAdd(c *gin.Context) {
	var info struct {
		Question    string `gorm:"column:question;type:varchar(50);not null" json:"question"` // 问题描述
		Description string `gorm:"column:description;type:varchar(100)" json:"description"`   // 问题详细描述
		Sponsor     string `gorm:"column:sponsor;type:varchar(100);not null" json:"sponsor"`  // 发起人
		Status      int    `gorm:"default:0"`                                                 // 处理进度

	}
	if c.BindJSON(&info) != nil {
		Code.SE400(c)
		return
	}
	dbconn := db.GetConn()
	var result db.Employer
	if err := db.GetConn().Model(&db.Employer{}).Where("employid = ?", info.Sponsor).First(&result).Error; err != nil {
		Code.SE500(c)
		return
	}
	now := time.Now()
	year, month, day := now.Date()
	zero := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	err := dbconn.Exec("insert into Feedback (question, description, sponsor, status, create_date, teleinfo) values (?, ?, ?, ?, ?, ?)",
		info.Question, info.Description, info.Sponsor, info.Status, zero, result.Telephone).Error
	if err != nil {
		Code.SE500(c)
		return
	}
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "success",
	})
}
