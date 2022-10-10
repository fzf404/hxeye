package config

import (
	"log"

	"github.com/spf13/viper"
)

// InitConfig 初始化配置文件
func InitConfig() {
	// 设置配置文件信息
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	// 搜索路径
	viper.AddConfigPath("./config")
	// 自动根据类型来读取配置
	err := viper.ReadInConfig()
	if err != nil {
		log.Print("read config failed: ", err)
	}
}
