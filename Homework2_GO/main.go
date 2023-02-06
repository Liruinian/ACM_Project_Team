package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {
	database, err := sqlx.Open("mysql", "root:localhost@tcp(127.0.0.1:3306)/mytest")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	fmt.Println(database)
	ginServer := gin.Default()
	ginServer.POST("/login", func(context *gin.Context) {
		context.JSON(200, nil)
	})

	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello,gin"})
	})

	ginServer.Run(":80")
}
