package service

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"stvsljl.com/SSIMP/db"
)

func EmployeeList(c *gin.Context) {
	mgr := db.EmployerMgr(db.GetConn())
	employee, err := mgr.Gets()
	if err != nil {
		Code.SE601(c)
		return
	}
	employeejson, err := json.Marshal(employee)
	fmt.Println(string(employeejson))
	if err != nil {
		Code.SE602(c)
		return
	}
	fmt.Println(employee)
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "success",
		"data": string(employeejson),
	})
}

func EmployeeAdd(c *gin.Context) {}

func EmployeeUpdate(c *gin.Context) {}

func EmployeeDelete(c *gin.Context) {}
