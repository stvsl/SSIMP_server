package service

import (
	"encoding/json"
	"fmt"
	"strings"

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
	// fmt.Println(string(employeejson))
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

func EmployeeLogin(c *gin.Context) {
	// 获取账号密码
	var employee db.Employer
	err := c.BindJSON(&employee)
	if err != nil {
		Code.SE400(c)
		return
	}
	// 验证账号密码
	mgr := db.EmployerMgr(db.GetConn())
	employee, err = mgr.GetFromEmployid(employee.Employid)
	if err != nil {
		Code.SE406(c)
		return
	}
	if strings.Compare(employee.Passwd, employee.Passwd) != 0 {
		Code.SE406(c)
		return
	}
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "success",
		"data": employee,
	})
}

func EmployeeTaskList(c *gin.Context) {
	// 获取id
	var employee db.Employer
	err := c.BindJSON(&employee)
	if err != nil {
		Code.SE400(c)
		return
	}
	tasksets := []db.TaskSet{}
	// 获取任务
	db := db.GetConn()

}
