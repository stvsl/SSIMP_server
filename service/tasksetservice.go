package service

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"stvsljl.com/SSIMP/db"
)

func TaskSetList(c *gin.Context) {
	mgr := db.TaskSetMgr(db.GetConn())
	taskset, err := mgr.Gets()
	if err != nil {
		Code.SE601(c)
		return
	}
	tasksetjson, err := json.Marshal(taskset)
	fmt.Println(string(tasksetjson))
	if err != nil {
		Code.SE602(c)
		return
	}
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "success",
		"data": string(tasksetjson),
	})
}

func TaskSetAdd(c *gin.Context) {
	taskset := db.TaskSet{}
	err := c.BindJSON(&taskset)
	if err != nil {
		Code.SE400(c)
		return
	}
	db := db.GetConn()
	db.Create(&taskset).Select("tid").Scan(&taskset)
	if err != nil {
		Code.SE602(c)
		return
	}
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "success",
	})
}

func TaskSetUpdate(c *gin.Context) {

	taskset := db.TaskSet{}
	err := c.BindJSON(&taskset)
	if err != nil {
		Code.SE400(c)
		return
	}
	db := db.GetConn()
	db.Model(&taskset).Updates(taskset)
	if err != nil {
		Code.SE602(c)
		return
	}
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "success",
	})

}

func TaskSetDelete(c *gin.Context) {

	tid := c.Query("tid")
	fmt.Println(tid)
	tidint, err := strconv.Atoi(tid)
	if err != nil {
		Code.SE400(c)
		return
	}
	tsmgr := db.TaskSetMgr(db.GetConn())
	tsmgr.Delete(db.TaskSet{Tid: tidint})
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "success",
	})
}
