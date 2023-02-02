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
		chapertext := "1234567890"
		passed := "e806256002f490a56333d2b66d48d20b"
		fmt.Println("加密前原文" + chapertext)
		fmt.Println("加密前密码" + passed)
		// 二进制
		fmt.Println("加密前原文二进制" + fmt.Sprintf("%b", []byte(chapertext)))
		fmt.Println("加密前密码二进制" + fmt.Sprintf("%b", []byte(passed)))
		security.Init()
		crypted, err := security.AesEncrypt([]byte(chapertext), []byte(passed))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("加密后" + string(crypted))
		decrypted, err := security.AesDecrypt(crypted, []byte(passed))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("解密后" + string(decrypted))
	}

}
