package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gookit/color"
	"github.com/unrolled/secure"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type Config struct {
	Port    string
	UseCors bool
	UseTLS  bool
	TLSPem  string
	TLSKey  string
	SSLHost string
}

var Conf = Config{
	Port:    ":8880", // 网站访问端口
	UseCors: true,    // 是否允许跨域访问
	UseTLS:  false,   //是否使用TLS加密（https）*使用加密需要填写以下字段
	TLSPem:  "api.liruinian.top.pem",
	TLSKey:  "api.liruinian.top.key",
	SSLHost: "api.liruinian.top:8880",
}

func Cors() gin.HandlerFunc {
	// 跨域设置： 前后端分离在不同服务器上，需要进行跨域处理
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

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}

func main() {

	r := gin.New()
	r.Use(Cors())
	if Conf.UseTLS {
		r.Use(TlsHandler())
	}

	Dba = MysqlConn("articles")
	Dbl = MysqlConn("login")

	Login(r)
	Signup(r)
	GetIdentity(r)
	Logout(r)
	GetArticles(r)
	GetArticle(r)
	GetArticleList(r)
	UploadArticle(r)
	RemoveArticle(r)
	SearchArticle(r)

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "hello,gin"})
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

// GetPwd 给密码加密
func GetPwd(pwd string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return hash, err
}

// ComparePwd 比对密码
func ComparePwd(pwd1 string, pwd2 string) bool {
	// Returns true on success, pwd1 is for the database.
	err := bcrypt.CompareHashAndPassword([]byte(pwd1), []byte(pwd2))
	if err != nil {
		return false
	} else {
		return true
	}
}
