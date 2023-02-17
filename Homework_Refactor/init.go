package main

import (
	"Homework_Refactor/logic"
	"Homework_Refactor/tools"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"github.com/unrolled/secure"
	"log"
	"net/http"
)

type Config struct {
	Port    string
	UseCors bool
	Domain  string
	UseTLS  bool
	TLSPem  string
	TLSKey  string
	SSLHost string
}

var (
	Conf = Config{
		Port:    ":8880",                 // 网站访问端口
		UseCors: true,                    // 是否允许跨域访问
		Domain:  ".liruinian.top",        // setcookie 时使用的域名
		UseTLS:  false,                   // 是否使用TLS加密（https）*使用加密需要填写以下字段
		TLSPem:  "api.liruinian.top.pem", // pem路径
		TLSKey:  "api.liruinian.top.key", // key路径
		SSLHost: "api.liruinian.top",     // ssl证书生效的域名
	}
)

func Cors() gin.HandlerFunc {
	// 跨域设置： 前后端分离在不同端口上，需要进行跨域处理
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")

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
func TlsHandler() gin.HandlerFunc {
	// https 支持
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     Conf.SSLHost,
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			return
		}

		c.Next()
	}
}
func UserVerifyMiddleware() gin.HandlerFunc {
	// 用户权限检测中间件： use bcrypt and jwt
	return func(c *gin.Context) {
		ip := c.ClientIP()

		username := c.GetHeader("username")
		token := c.GetHeader("authorization")
		log.Println(username, token)
		if tools.VerifyUser(username, token) {
			log.Println("UserVerifyMiddleware: " + color.FgGreen.Render("允许用户名："+username+" 访问文章接口"))
			c.Next()
		} else {
			log.Println("UserVerifyMiddleware: " + color.FgYellow.Render("不允许ip为"+ip+"的游客访问文章接口"))
			c.JSON(200, gin.H{"code": 201, "msg": "访问受限，请重新登录后再进行操作"})
			c.Abort()
			return
		}
	}
}
func main() {

	tools.DB = tools.DbConn("root", "e89r245z", "127.0.0.1:3306", "homework")
	r := gin.New()
	r.Use(Cors())
	if Conf.UseTLS {
		r.Use(TlsHandler())
	}
	user := r.Group("/user")
	{
		user.POST("/login", logic.Login)
		user.POST("/logout", logic.Logout)
		user.POST("/register", logic.Signup)

		user.POST("/info", logic.GetUserInfo)
	}
	article := r.Group("/article")
	article.Use(UserVerifyMiddleware())
	{
		article.POST("/list", logic.GetArticles)
		article.POST("/create", logic.CreateArticle)
		article.DELETE("/delete/:id", logic.RemoveArticle)

		article.POST("/edit", logic.EditArticle)
		article.POST("/:id", logic.GetArticle) // 单行
		article.POST("/comments/:id", logic.GetComments)
	}
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 200, "msg": "Welcome to api.liruinian.top!", "version": 0.1, "description": ""})
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 404, "msg": "无此请求接口，请查阅api文档后再次尝试访问"})
	})
	if Conf.UseTLS {
		err := r.RunTLS(Conf.Port, Conf.TLSPem, Conf.TLSKey)
		if err != nil {
			log.Println(color.FgRed.Render(err.Error()))
		}
	} else {
		err := r.Run(Conf.Port)
		if err != nil {
			log.Println(color.FgRed.Render(err.Error()))
		}
	}
}
