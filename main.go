package main

import (
	"fmt"
	"time"

	"stvsljl.com/SSIMP/db"
	"stvsljl.com/SSIMP/security"
	"stvsljl.com/SSIMP/service"
	"stvsljl.com/SSIMP/utils"
)

func main() {
	fmt.Println("请输入运行模式 1. 测试模式 2. 生产模式")
	var mode = "1"
	fmt.Scanln(&mode)
	if mode == "1" {
		db.Connect()
		mgr := db.AdminMgr(db.GetConn())
		test, _ := mgr.GetFromAdminID("1234567890")
		utils.Log.Init()
		utils.Log.Info(test)
		security.Init()
		utils.Log.Info("服务器RSA密钥更新")
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
	} else if mode == "3" {
		// 安全组件相关事务初始化
		security.Init()
		// 加密
		data, _ := security.RsaEncrypt([]byte("1234567890"))
		fmt.Println(data)
		// 解密
		data, _ = security.RsaDecrypt(data)
		fmt.Println(string(data))
	}

}
