package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"stvsljl.com/SSIMP/db"
)

func ArticleList(c *gin.Context) {

	type Article struct {
		Aid        int       `gorm:"column:aid;primary_key" db:"aid" json:"aid" form:"aid"`                 //  文章编号
		Coverimg   string    `gorm:"column:coverimg" db:"coverimg" json:"coverimg" form:"coverimg"`         //  封面图片
		Title      string    `gorm:"column:title" db:"title" json:"title" form:"title"`                     //  标题
		Updatetime time.Time `gorm:"column:updatetime" db:"updatetime" json:"updatetime" form:"updatetime"` //  更新日期
		Pageviews  int64     `gorm:"column:pageviews" db:"pageviews" json:"pageviews" form:"pageviews"`     //  浏览量
	}
	ArticleList := []Article{}
	db := db.GetConn()
	db.Table("Article").Select("aid", "title", "writetime", "updatetime", "pageviews").Find(&ArticleList)
	fmt.Println(ArticleList)
	json, _ := json.Marshal(ArticleList)
	c.JSON(200, gin.H{
		"code": "SE200",
		"data": string(json),
	})
}

func ArticleDetail(c *gin.Context) {
	var aidstruct struct {
		Aid string `json:"aid" form:"aid"`
	}
	c.Bind(&aidstruct)
	articlemgr := db.ArticleMgr(db.GetConn())
	var aid int
	fmt.Sscan(aidstruct.Aid, &aid)
	article, err := articlemgr.GetByOption(articlemgr.WithAid(aid))
	if err != nil {
		Code.SE602(c)
		return
	}
	json, _ := json.Marshal(article)
	fmt.Println(string(json))
	c.JSON(200, gin.H{
		"code": "SE200",
		"data": string(json),
	})
}

func ArticleAdd(c *gin.Context) {
	article := db.Article{}
	c.Bind(&article)
	fmt.Println(string(func() []byte {
		json, _ := json.Marshal(article)
		return json
	}()))
	article.Writetime = time.Now()
	article.Updatetime = time.Now()
	article.Pageviews = 0
	db := db.GetConn()
	db.Create(&article).Select("aid").Scan(&article)
	// 判断aid是否非空
	if article.Aid == 0 {
		Code.SE602(c)
		return
	}
	c.JSON(200, gin.H{
		"code": "SE200",
		"aid":  article.Aid,
	})
}

func ArticleUpdate(c *gin.Context) {}

func ArticleDelete(c *gin.Context) {}
