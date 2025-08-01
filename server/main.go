package main

import (
	"github.com/yangle520/gin-vue-demo/server/core"
	"github.com/yangle520/gin-vue-demo/server/global"
	"github.com/yangle520/gin-vue-demo/server/initialize"
)

func main() {

	// 初始化系统
	initializeSystem()

	// 运行服务器
	core.RunServer()
}

func initializeSystem() {

	// 初始化Viper
	global.GVA_VP = core.Viper()
	// 初始化zap日志库
	global.GVA_LOG = core.Zap()
	// 初始化数据库
	global.GVA_DB = initialize.Gorm()
	// 初始化表
	initialize.RegisterTables()
}
