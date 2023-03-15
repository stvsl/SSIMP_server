package service

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"stvsljl.com/SSIMP/db"
)

func EmployerTaskList(c *gin.Context) {
	// 从Query中获取参数
	eid := c.Query("eid")
	// 查询数据库
	taskmgr := db.TaskMgr(db.GetConn())
	tasklist, err := taskmgr.GetByOptions(taskmgr.WithEmployid(eid))
	if err != nil {
		Code.SE602(c)
		return
	}
	taskslist, err := json.Marshal(tasklist)
	if err != nil {
		Code.SE600(c)
		return
	}
	fmt.Println(string(taskslist))
	// 返回数据
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "查询成功",
		"data": string(taskslist),
	})
}
