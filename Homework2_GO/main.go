package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func main() {
	r := gin.New()
	r.Use(Cors())
	Dba = mysqlConn("articles")
	Dbl = mysqlConn("login")
	go login(r)
	go signup(r)
	go getArticles(r)
	go getIdentity(r)
	go logout(r)
	go uploadArticle(r)
	go removeArticle(r)
	r.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello,gin"})
	})

	r.Run(":8880")
}
