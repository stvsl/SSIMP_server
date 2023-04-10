package main

import (
	"stvsljl.com/SSIMP/db"
	"stvsljl.com/SSIMP/service"
)

func main() {
	// fmt.Println("请输入运行模式 1. 测试模式 2. 生产模式")
	// var mode = "1"
	// fmt.Scanln(&mode)
	// if mode == "1" {
	// 	redis.Init()
	// 	rundata := redis.GetRunningDataStruct()
	// 	rundata.ReadAllFromRedis()
	// } else if mode == "2" {
	db.Connect()
	// 启动服务器
	service.Start()
	// }

}
