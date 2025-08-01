package initialize

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/yangle520/gin-vue-demo/server/config"
	"github.com/yangle520/gin-vue-demo/server/global"
)

func initRedisClient(redisCfg config.Redis) (redis.UniversalClient, error) {
	var client redis.UniversalClient

	if redisCfg.UseCluster {
		// 使用集群模式
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    redisCfg.ClusterAddrs,
			Password: redisCfg.Password,
		})
	} else {
		// 使用单例模式
		client = redis.NewClient(&redis.Options{
			Addr:     redisCfg.Addr,
			Password: redisCfg.Password,
			DB:       redisCfg.DB,
		})
	}

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("redis 连接异常,err: %v", err)
	}

	fmt.Println("redis 连接成功")

	return client, nil
}

func Redis() {
	redisClient, err := initRedisClient(global.GVA_CONFIG.Redis)
	if err != nil {
		panic(err)
	}
	global.GVA_REDIS = redisClient
}
