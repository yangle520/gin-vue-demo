package global

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"github.com/yangle520/gin-vue-demo/server/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_DB        *gorm.DB
	GVA_DBList    map[string]*gorm.DB
	GVA_REDIS     redis.UniversalClient
	GVA_REDISList map[string]redis.UniversalClient
	// GVA_MONGO               *qmgo.QmgoClient
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	GVA_LOG    *zap.Logger
	// GVA_Timer  timer.Timer = timer.NewTimerTask()
	// GVA_Concurrency_Control             = &singleflight.Group{}
	GVA_ROUTERS       gin.RoutesInfo
	GVA_ACTIVE_DBNAME *string
	lock              sync.RWMutex
)

func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return GVA_DBList[dbname]
}

func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GVA_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}

func GetRedis(name string) redis.UniversalClient {
	redis, ok := GVA_REDISList[name]
	if !ok || redis == nil {
		panic(fmt.Sprintf("redis `%s` no init", name))
	}
	return redis
}
