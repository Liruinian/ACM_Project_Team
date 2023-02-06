package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	ginServer := gin.Default()
	ginServer.POST("/login", func(context *gin.Context) {
		context.JSON()
	})

	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello,gin"})
	})

	ginServer.Run(":80")
}
