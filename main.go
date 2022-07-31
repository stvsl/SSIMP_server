package main

import (
	"fmt"

	"stvsljl.com/SSIMP/service"
	"stvsljl.com/SSIMP/utils"
)

func main() {
	// 初始化数据库连接池
	// 启动服务器
	fmt.Println(utils.GetSqlConnConfig())
	service.Start()
}
