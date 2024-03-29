package service

import (
	"encoding/json"
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
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "success",
		"data": string(employeejson),
	})
}

func EmployeeInfo(c *gin.Context) {

	eidstr := c.Query("eid")
	if eidstr == "" {
		Code.SE400(c)
		return
	}
	dbconn := db.GetConn()
	employee := db.Employer{}
	dbconn.Model(&employee).Where("employid = ?", eidstr).First(&employee)
	employeejson, err := json.Marshal(employee)
	if err != nil {
		Code.SE602(c)
		return
	}
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "success",
		"data": string(employeejson),
	})
}
func EmployeeUpdatePasswd(c *gin.Context) {
	var info struct {
		Eid       string `json:"eid"`
		Oldpasswd string `json:"oldpasswd"`
		Newpasswd string `json:"newpasswd"`
	}
	err := c.BindJSON(&info)
	if err != nil {
		Code.SE400(c)
		return
	}
	dbconn := db.GetConn()
	employee := db.Employer{}
	dbconn.Model(&employee).Where("employid = ?", info.Eid).First(&employee)
	if employee.Passwd != info.Oldpasswd {
		Code.SE406(c)
		return
	}
	employee.Passwd = info.Newpasswd
	dbconn.Model(&employee).Updates(&employee)
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "success",
	})
}

func EmployeeAdd(c *gin.Context) {
	// 获取数据
	var employee db.Employer
	err := c.BindJSON(&employee)
	if err != nil {
		Code.SE400(c)
		return
	}
	// 添加数据
	db := db.GetConn()
	err = db.Create(&employee).Error
	if err != nil {
		Code.SE601(c)
		return
	}
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "success",
		"data": employee,
	})
}

func EmployeeUpdate(c *gin.Context) {
	// 获取数据
	var employee db.Employer
	err := c.BindJSON(&employee)
	if err != nil {
		Code.SE400(c)
		return
	}
	// 更新数据
	db := db.GetConn()
	err = db.Model(&employee).Updates(&employee).Error
	if err != nil {
		Code.SE601(c)
		return
	}
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "success",
	})
}

func EmployeeDelete(c *gin.Context) {
	// 获取数据
	var employee db.Employer
	err := c.BindJSON(&employee)
	if err != nil {
		Code.SE400(c)
		return
	}
	// 删除数据
	db := db.GetConn()
	err = db.Exec("DELETE FROM Employer WHERE employid = '" + employee.Employid + "';").Error
	if err != nil {
		Code.SE601(c)
		return
	}
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "success",
	})
}

func EmployeeLogin(c *gin.Context) {
	// 获取账号密码
	var employee db.Employer
	err := c.BindJSON(&employee)
	if err != nil {
		Code.SE400(c)
		return
	}
	// 验证账号密码
	passwd := employee.Passwd
	mgr := db.EmployerMgr(db.GetConn())
	employee, err = mgr.GetFromEmployid(employee.Employid)
	if err != nil {
		Code.SE406(c)
		return
	}
	if strings.Compare(employee.Passwd, passwd) != 0 {
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
	err = db.Exec("SELECT * FROM Task INNER JOIN TaskSet ON Task.tid = TaskSet.tid WHERE Task.employid = '" + employee.Employid + "';").Scan(&tasksets).Error
	if err != nil {
		Code.SE601(c)
		return
	}
	tasksetsjson, err := json.Marshal(tasksets)
	if err != nil {
		Code.SE500(c)
		return
	}
	c.JSON(200, gin.H{
		"code": "SE200",
		"msg":  "success",
		"data": string(tasksetsjson),
	})
}
