package common

import (
	"fmt"
	"ginessential/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB

func InitDB() *gorm.DB {
	host := "192.192.191.244"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "73da0b2bf5e36958"
	charset := "utf8mb4"
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
// 定义一个方法获取DB实例
func GetDB() *gorm.DB {
	return DB
}