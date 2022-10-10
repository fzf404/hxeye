package main

import (
	"App/config"
	"App/database"
	"App/route"

	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"
)

func main() {
	// 初始化
	config.InitConfig()
	database.InitSQL()

	r := gin.Default()
	// 设置最大文件大小
	r.MaxMultipartMemory = 2 << 20
	r = route.CollectRoute(r)
	port := viper.GetString("common.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
	
}

