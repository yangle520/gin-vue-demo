package initialize

import (
	"os"

	"github.com/yangle520/gin-vue-demo/server/config"
	"github.com/yangle520/gin-vue-demo/server/global"
	"github.com/yangle520/gin-vue-demo/server/model/example"
	"github.com/yangle520/gin-vue-demo/server/model/system"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化数据库连接
func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		global.GVA_ACTIVE_DBNAME = &global.GVA_CONFIG.Mysql.Dbname
		return initMysqlDB(global.GVA_CONFIG.Mysql)
	case "pgsql":
		global.GVA_ACTIVE_DBNAME = &global.GVA_CONFIG.Pgsql.Dbname
		return nil
	default:
		global.GVA_ACTIVE_DBNAME = &global.GVA_CONFIG.Mysql.Dbname
		return initMysqlDB(global.GVA_CONFIG.Mysql)
	}
}

// 初始化表结构
func RegisterTables() {
	db := global.GVA_DB

	err := db.AutoMigrate(
		system.Equipment{},
		system.User{},
		example.ExaCustomer{},
	)

	if err != nil {
		global.GVA_LOG.Error("创建表失败", zap.Error(err))
		os.Exit(0)
	}

	global.GVA_LOG.Info("创建表成功")
}

// 初始化 mysql 数据库连接
func initMysqlDB(m config.Mysql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}

	if db, err := gorm.Open(mysql.Open(m.Dsn())); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "Engine="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
