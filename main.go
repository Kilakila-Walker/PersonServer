package main

import (
	"perServer/core"
	"perServer/global"
	"perServer/initialize"
	//"runtime"
)

// 一切的开始与结束
// 架构提供者地址：
// https://github.com/flipped-aurora/gin-vue-admin
func main() {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		initialize.Mysql()
	default:
		initialize.Mysql()
	}
	initialize.DBTables()
	// 程序结束前关闭数据库链接
	defer global.GVA_DB.Close()

	core.RunWindowsServer()
}
