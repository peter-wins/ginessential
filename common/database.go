package common

import (
	"fmt"
	"ginessential/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB

func InitDB() *gorm.DB {
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password :=viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&model.User{})


	DB = db
	return db
}
// GetDB 定义一个方法获取DB实例
func GetDB() *gorm.DB {
	return DB
}