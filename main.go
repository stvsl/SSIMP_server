package main

import (
	"fmt"
	"time"

	"stvsljl.com/SSIMP/db"
	"stvsljl.com/SSIMP/security"
	"stvsljl.com/SSIMP/service"
)

func main() {
	fmt.Println("请输入运行模式 1. 测试模式 2. 生产模式")
	var mode = "1"
	fmt.Scanln(mode)
	if mode == "1" {
		db.Connect()
		mgr := db.AdminMgr(db.GetConn())
		test, _ := mgr.GetFromAdminID("1234567890")
		fmt.Println(test)
		security.Init()
		fmt.Println(string(security.SERVER_RSA.PRIVATE_KEY))
		fmt.Println(string(security.SERVER_RSA.PUBLIC_KEY))
		time.Sleep(120 * time.Second)
	} else if mode == "2" {
		db.Connect()
		// 启动服务器
		service.Start()
		mgr := db.AdminMgr(db.GetConn())
		for {
			test, _ := mgr.GetFromAdminID("1234567890")
			fmt.Println(test)
		}
	}

}
