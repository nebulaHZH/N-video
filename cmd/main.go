package main

import (
	router "N-video/router"
	_ "N-video/utils" //自动执行mysql连接
	"fmt"

	"github.com/gin-gonic/gin"
	//需要安装 go get -u github.com/swaggo/files
	//需要安装 go get -u github.com/swaggo/gin-swagger
)

var err error

// @title N-video API
// @version 1.0
// @description This is a sample server for Your Project.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := gin.Default()
	router.Router(r)
	if err != nil {
		fmt.Println("mysql connect failed")
	} else {
		fmt.Println("mysql connect success")
	}
	// 设置api文档
	// http://127.0.0.1:8080/swagger/index.html
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
