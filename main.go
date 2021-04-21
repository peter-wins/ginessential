package main

import (
	"ginessential/common"
	"github.com/gin-gonic/gin"
)


func main() {
	// 初始化数据库
	common.InitDB()

	r := gin.Default()
	r = CollectRouter(r)
	panic(r.Run(":8080"))
}



