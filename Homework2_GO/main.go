package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func mysqlConn() {
	DB, _ := sql.Open("mysql", "root:@a20040207@tcp(127.0.0.1:3306)/articles")
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connect success")
}
func main() {
	mysqlConn()
	
	ginServer := gin.Default()
	ginServer.POST("/login", func(context *gin.Context) {
		context.JSON(200, nil)
	})

	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello,gin"})
	})

	ginServer.Run(":80")
}
