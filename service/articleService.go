package service

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"stvsljl.com/SSIMP/cos"
	"stvsljl.com/SSIMP/db"
	"stvsljl.com/SSIMP/redis"
)

func ArticleList(c *gin.Context) {

	type Article struct {
		Aid          int       `gorm:"column:aid;primary_key" db:"aid" json:"aid" form:"aid"`                         //  文章编号
		Coverimg     string    `gorm:"column:coverimg" db:"coverimg" json:"coverimg" form:"coverimg"`                 //  封面图片
		Title        string    `gorm:"column:title" db:"title" json:"title" form:"title"`                             //  标题
		Introduction string    `gorm:"column:introduction" db:"introduction" json:"introduction" form:"introduction"` //  简介
		Updatetime   time.Time `gorm:"column:updatetime" db:"updatetime" json:"updatetime" form:"updatetime"`         //  更新日期
		Pageviews    int64     `gorm:"column:pageviews" db:"pageviews" json:"pageviews" form:"pageviews"`             //  浏览量
		Status       int64     `gorm:"column:status" db:"status" json:"status" form:"status"`                         //  文章状态
	}
	ArticleList := []Article{}
	db := db.GetConn()
	db.Table("Article").Select("aid", "title", "coverimg", "updatetime", "pageviews", "status", "introduction").Find(&ArticleList)
	// fmt.Println(ArticleList)
	json, _ := json.Marshal(ArticleList)
	c.JSON(200, gin.H{
		"code": "SE200",
		"data": string(json),
	})
}

func ArticleCarousel(c *gin.Context) {
	type Article struct {
		Aid          int       `gorm:"column:aid;primary_key" db:"aid" json:"aid" form:"aid"`                         //  文章编号
		Coverimg     string    `gorm:"column:coverimg" db:"coverimg" json:"coverimg" form:"coverimg"`                 //  封面图片
		Title        string    `gorm:"column:title" db:"title" json:"title" form:"title"`                             //  标题
		Introduction string    `gorm:"column:introduction" db:"introduction" json:"introduction" form:"introduction"` //  简介
		Updatetime   time.Time `gorm:"column:updatetime" db:"updatetime" json:"updatetime" form:"updatetime"`         //  更新日期
		Pageviews    int64     `gorm:"column:pageviews" db:"pageviews" json:"pageviews" form:"pageviews"`             //  浏览量
		Status       int64     `gorm:"column:status" db:"status" json:"status" form:"status"`                         //  文章状态
	}
	ArticleList := []Article{}
	db := db.GetConn()
	db.Table("Article").Select("aid", "title", "coverimg", "updatetime", "pageviews", "status", "introduction").Where("status = ?", 2).Find(&ArticleList)
	// fmt.Println(ArticleList)
	json, _ := json.Marshal(ArticleList)
	c.JSON(200, gin.H{
		"code": "SE200",
		"data": string(json),
	})
}

func ArticleRecommendList(c *gin.Context) {
	// 从post请求中获取参数
	var aidstruct struct {
		Name string `json:"name" form:"name"`
	}
	if err := c.Bind(&aidstruct); err != nil {
		Code.SE400(c)
		return
	}
	type Article struct {
		Aid          int       `gorm:"column:aid;primary_key" db:"aid" json:"aid" form:"aid"`                         //  文章编号
		Coverimg     string    `gorm:"column:coverimg" db:"coverimg" json:"coverimg" form:"coverimg"`                 //  封面图片
		Title        string    `gorm:"column:title" db:"title" json:"title" form:"title"`                             //  标题
		Introduction string    `gorm:"column:introduction" db:"introduction" json:"introduction" form:"introduction"` //  简介
		Updatetime   time.Time `gorm:"column:updatetime" db:"updatetime" json:"updatetime" form:"updatetime"`         //  更新日期
		Pageviews    int64     `gorm:"column:pageviews" db:"pageviews" json:"pageviews" form:"pageviews"`             //  浏览量
		Status       int64     `gorm:"column:status" db:"status" json:"status" form:"status"`                         //  文章状态
	}
	ArticleList := []Article{}
	db := db.GetConn()
	db.Table("Article").Select("aid", "title", "coverimg", "updatetime", "pageviews", "status", "introduction").Where("status = ?", 2).Where("title like ?", "%"+aidstruct.Name+"%").Find(&ArticleList)
	// fmt.Println(ArticleList)
	json, _ := json.Marshal(ArticleList)
	c.JSON(200, gin.H{
		"code": "SE200",
		"data": string(json),
	})
}

func ArticlePublicList(c *gin.Context) {
	type Article struct {
		Aid          int       `gorm:"column:aid;primary_key" db:"aid" json:"aid" form:"aid"`                         //  文章编号
		Coverimg     string    `gorm:"column:coverimg" db:"coverimg" json:"coverimg" form:"coverimg"`                 //  封面图片
		Title        string    `gorm:"column:title" db:"title" json:"title" form:"title"`                             //  标题
		Introduction string    `gorm:"column:introduction" db:"introduction" json:"introduction" form:"introduction"` //  简介
		Updatetime   time.Time `gorm:"column:updatetime" db:"updatetime" json:"updatetime" form:"updatetime"`         //  更新日期
		Pageviews    int64     `gorm:"column:pageviews" db:"pageviews" json:"pageviews" form:"pageviews"`             //  浏览量
		Status       int64     `gorm:"column:status" db:"status" json:"status" form:"status"`                         //  文章状态
	}
	ArticleList := []Article{}
	db := db.GetConn()
	db.Table("Article").Select("aid", "title", "coverimg", "updatetime", "pageviews", "status", "introduction").Where("status != ?", 3).Find(&ArticleList)
	// fmt.Println(ArticleList)
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
	if err := c.Bind(&aidstruct); err != nil {
		Code.SE400(c)
		return
	}
	articlemgr := db.ArticleMgr(db.GetConn())
	var aid int
	fmt.Sscan(aidstruct.Aid, &aid)
	article, err := articlemgr.GetByOption(articlemgr.WithAid(aid))
	if err != nil {
		Code.SE602(c)
		return
	}
	json, _ := json.Marshal(article)
	// fmt.Println(string(json))
	c.JSON(200, gin.H{
		"code": "SE200",
		"data": string(json),
	})
	article.Pageviews++
	articlemgr.Model(article).Update("pageviews", article.Pageviews)
	// 更新Redis浏览量
	red := redis.RunningData{}
	red.AddNewViewCount()
}

func ArticleSearch(c *gin.Context) {
	var namestruct struct {
		Name string `json:"name" form:"name"`
	}
	if err := c.Bind(&namestruct); err != nil {
		Code.SE400(c)
		return
	}
	namestruct.Name = "%" + namestruct.Name + "%" // 模糊查询
	dbcoon := db.GetConn()
	type Article struct {
		Aid          int       `gorm:"column:aid;primary_key" db:"aid" json:"aid" form:"aid"`                         //  文章编号
		Coverimg     string    `gorm:"column:coverimg" db:"coverimg" json:"coverimg" form:"coverimg"`                 //  封面图片
		Title        string    `gorm:"column:title" db:"title" json:"title" form:"title"`                             //  标题
		Introduction string    `gorm:"column:introduction" db:"introduction" json:"introduction" form:"introduction"` //  简介
		Updatetime   time.Time `gorm:"column:updatetime" db:"updatetime" json:"updatetime" form:"updatetime"`         //  更新日期
		Pageviews    int64     `gorm:"column:pageviews" db:"pageviews" json:"pageviews" form:"pageviews"`             //  浏览量
		Status       int64     `gorm:"column:status" db:"status" json:"status" form:"status"`                         //  文章状态
	}
	articles := []Article{}
	// SELECT title FROM `Article` WHERE `title` LIKE '%%';改写成
	dbcoon.Model(&db.Article{}).Where("title LIKE ?", namestruct.Name).Find(&articles)
	json, _ := json.Marshal(articles)
	c.JSON(200, gin.H{
		"code": "SE200",
		"data": string(json),
	})
}

func ArticleAdd(c *gin.Context) {
	article := db.Article{}
	c.Bind(&article)
	// 读取base64图片数据，调用cos上传图片，返回图片地址
	coverpath, err := cos.UploadFile(article.Coverimg)
	if err != nil {
		fmt.Println("对象存储上传失败" + err.Error())
		Code.SE620(c)
		return
	}
	article.Coverimg = coverpath
	contentpath, err := cos.UploadFile(article.Contentimg)
	if err != nil {
		fmt.Println("对象存储上传失败" + err.Error())
		Code.SE620(c)
		return
	}
	article.Contentimg = contentpath
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

func ArticleUpdate(c *gin.Context) {
	var article struct {
		Aid          string    `gorm:"column:aid;primary_key" db:"aid" json:"aid" form:"aid"`                         //  文章编号
		Coverimg     string    `gorm:"column:coverimg" db:"coverimg" json:"coverimg" form:"coverimg"`                 //  封面图片
		Contentimg   string    `gorm:"column:contentimg" db:"contentimg" json:"contentimg" form:"contentimg"`         //  内容大图
		Title        string    `gorm:"column:title" db:"title" json:"title" form:"title"`                             //  标题
		Introduction string    `gorm:"column:introduction" db:"introduction" json:"introduction" form:"introduction"` //  简介
		Text         string    `gorm:"column:text" db:"text" json:"text" form:"text"`                                 //  正文
		Updatetime   time.Time `gorm:"column:updatetime" db:"updatetime" json:"updatetime" form:"updatetime"`         //  更新日期
		Status       int64     `gorm:"column:status" db:"status" json:"status" form:"status"`                         //  文章状态
	}
	c.Bind(&article)
	// 读取base64图片数据，判断内容是否为base64格式的图片数据，调用cos上传图片，返回图片地址
	// 判断两个图片数据是不是前缀是http
	if !strings.HasPrefix(article.Coverimg, "http") {
		coverpath, err := cos.UploadFile(article.Coverimg)
		if err != nil {
			fmt.Println("对象存储上传失败" + err.Error())
			Code.SE620(c)
			return
		}
		article.Coverimg = coverpath
	}
	if !strings.HasPrefix(article.Contentimg, "http") {
		contentpath, err := cos.UploadFile(article.Contentimg)
		if err != nil {
			fmt.Println("对象存储上传失败" + err.Error())
			Code.SE620(c)
			return
		}
		article.Contentimg = contentpath
	}
	article.Updatetime = time.Now()
	db := db.GetConn()
	aid, err := strconv.Atoi(article.Aid)
	if err != nil {
		Code.SE401(c)
		return
	}
	db.Exec("update Article set title=?,coverimg=?,contentimg=?,text=?,updatetime=?,status=? where aid=?", article.Title, article.Coverimg, article.Contentimg, article.Text, article.Updatetime, article.Status, aid)
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "更新成功",
	})

}

// ArticleToCarousel 取消轮播图设置
func ArticleToNoCarousel(c *gin.Context) {
	// 从GET请求中获取aid
	aid := c.Query("aid")
	// 将aid转换为int类型
	aidint, err := strconv.Atoi(aid)
	if err != nil {
		Code.SE401(c)
		return
	}
	fmt.Println(aidint)
	db := db.GetConn()
	db.Exec("update Article set status=0 where aid=?", aidint)
	fmt.Println("取消轮播图设置成功")
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "取消轮播图设置成功",
	})
}

func ArticleDelete(c *gin.Context) {
	// 从GET请求中获取aid
	aid := c.Query("aid")
	// 将aid转换为int类型
	aidint, err := strconv.Atoi(aid)
	if err != nil {
		Code.SE401(c)
		return
	}
	// 从数据库中删除aid对应的文章
	// 查询aid对应的文章，获取文章的封面图片和内容图片地址，调用cos删除图片
	articlemgr := db.ArticleMgr(db.GetConn())
	article, err := articlemgr.GetByOption(articlemgr.WithAid(aidint))
	if err != nil {
		fmt.Println("查询失败")
		Code.SE602(c)
		return
	}
	// 删除封面图片
	fmt.Println(article.Contentimg)
	err = cos.DeleteFile(article.Coverimg)
	if err != nil {
		fmt.Println("对象存储删除失败" + err.Error())
		Code.SE620(c)
		return
	}
	// 删除内容图片
	err = cos.DeleteFile(article.Contentimg)
	if err != nil {
		Code.SE620(c)
		return
	}
	// 删除文章
	db.GetConn().Delete(&article)
	c.JSON(200, gin.H{
		"code": "SE200",
	})
}
