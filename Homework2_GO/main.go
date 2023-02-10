package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/unrolled/secure"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "http://acmdev.liruinian.top")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func main() {
	useTLS := false
	r := gin.New()
	r.Use(Cors())
	if useTLS {
		r.Use(TlsHandler())
	}

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
	if useTLS {
		r.RunTLS(":8880", "api.liruinian.top.pem", "api.liruinian.top.key")
	} else {
		r.Run(":8880")
	}
}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "api.liruinian.top:8880",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
