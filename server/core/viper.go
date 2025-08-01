package core

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/yangle520/gin-vue-demo/server/core/internal"
	"github.com/yangle520/gin-vue-demo/server/global"
)

// Viper 是 Go 语言生态中最流行的配置管理库,支持多源配置融合与动态更新
func Viper() *viper.Viper {
	config := getConfigPath()

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("配置文件读取失败： %w", err))
	}

	v.WatchConfig()

	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件变动：", in.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		panic(fmt.Errorf("配置文件加载失败： %w", err))
	}

	// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}

// 获取配置文件路径，优先级： 命令行 > 环境变量 > 默认值
func getConfigPath() (config string) {

	// 命令行
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()
	if config != "" {
		fmt.Printf("正在使用命令行 -c 参数传递的值，config 路径为 %s\n", config)
		return
	}

	// 环境变量
	if env := os.Getenv(internal.ConfigEnv); env != "" {
		config = env
		fmt.Printf("正在使用 %s 环境，config 路径为 %s \n", internal.ConfigEnv, config)
		return
	}

	// 根据 gin 模式文件名
	switch gin.Mode() {
	case gin.DebugMode:
		config = internal.ConfigDebugFile
	case gin.TestMode:
		config = internal.ConfigTestFile
	case gin.ReleaseMode:
		config = internal.ConfigReleaseFile
	}
	fmt.Printf("正在使用 gin 的 %s 模式运行，config 路径为 %s\n", gin.Mode(), config)

	_, err := os.Stat(config)
	if err != nil || os.IsNotExist(err) {
		config = internal.ConfigDefaultFile
		fmt.Printf("配置文件路径不存在，使用默认配置文件 %s \n", config)
	}

	return
}
