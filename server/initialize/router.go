package initialize

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yangle520/gin-vue-demo/server/global"
	"github.com/yangle520/gin-vue-demo/server/router"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrPermission
	}

	return f, nil
}

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	// systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example

	// 提供静态文件访问
	Router.StaticFile("/favicon.ico", "./dist/favicon.ico")
	Router.Static("/assets", "./dist/assets")   // dist里面的静态资源
	Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面
	Router.StaticFS(global.GVA_CONFIG.Local.StorePath, justFilesFilesystem{http.Dir(global.GVA_CONFIG.Local.StorePath)})

	// 跨域
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)

	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	{
		exampleRouter.InitCustomerRouter(PrivateGroup) // 客户路由
	}

	global.GVA_ROUTERS = Router.Routes()

	return Router
}
