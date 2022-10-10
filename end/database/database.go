package database

import (
	"App/model"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 获取DB
var DB *gorm.DB

// InitSQL 初始化mysql数据库
func InitSQL() *gorm.DB {
	// 获取配置文件信息
	username := viper.Get("mysql.username")
	password := viper.Get("mysql.password")
	host := viper.Get("mysql.host")
	port := viper.Get("mysql.port")
	database := viper.Get("mysql.database")
	charset := viper.Get("mysql.charset")
	// 信息格式化
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)

	// fmt.Printf(args)

	// 连接及错误处理
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		log.Print("fail connect mysql: ", err)
	}

	// 自动建表
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Article{})

	DB = db
	return db
}

// GetDB 获取初始化完的DB
func GetDB() *gorm.DB {
	return DB
}
