package main

import (
	"ginessential/common"
	"ginessential/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)


func main() {
	// 读取配置文件
	config.InitConfig()
	// 初始化数据库
	common.InitDB()

	r := gin.Default()
	r = CollectRouter(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

