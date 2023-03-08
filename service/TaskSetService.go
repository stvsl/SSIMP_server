package service

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"stvsljl.com/SSIMP/db"
)

func TaskSetList(c *gin.Context) {
	mgr := db.TaskSetMgr(db.GetConn())
	taskset, err := mgr.Gets()
	tasksetjson, err := json.Marshal(taskset)
	fmt.Println(string(tasksetjson))
	if err != nil {
		Code.SE602(c)
		return
	}
	fmt.Println("tasksetjson" + string(tasksetjson))
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "success",
		"data": string(tasksetjson),
	})
}
