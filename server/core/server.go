package core

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yangle520/gin-vue-demo/server/global"
	"github.com/yangle520/gin-vue-demo/server/initialize"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

func RunServer() {

	// 初始化 redis
	if global.GVA_CONFIG.System.UseRedis {
		initialize.Redis()
	}

	// 初始化 mongo
	if global.GVA_CONFIG.System.UseMongo {
		initialize.Mongo.Initialization()
	}

	// 初始化路由
	Router := initialize.Routers()

	// 端口号
	port := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)

	// 初始化服务
	initServer(port, Router, 10*time.Minute, 10*time.Minute)
}

func initServer(port string, router *gin.Engine, readTimeout, writeTimeout time.Duration) {

	// 创建服务
	srv := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err)
			zap.L().Error("server启动失败", zap.Error(err))
			os.Exit(1)
		}
	}()

	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGALRM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("web服务关闭异常", zap.Error(err))
	}

	zap.L().Info("web服务已关闭")
}
