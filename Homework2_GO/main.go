package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gookit/color"
	"github.com/unrolled/secure"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	time2 "time"
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

var Conf = Config{
	Port:    ":8880", // 网站访问端口
	UseCors: true,    // 是否允许跨域访问
	Domain:  "api.liruinian.top",
	UseTLS:  false, //是否使用TLS加密（https）*使用加密需要填写以下字段
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
		if err != nil {
			return
		}

		c.Next()
	}
}
func UserVerifyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := c.Cookie("username")
		token, err := c.Cookie("login_token")
		if err != nil {
			return
		}
		if VerifyToken(token, "usertoken", username) {
			c.Set("username", username)
			c.Next()
		} else {
			c.JSON(200, gin.H{"code": 201, "msg": "访问受限，请重新登录后再进行操作"})
			c.Abort()
			return
		}
	}
}
func main() {
	r := gin.New()
	r.Use(Cors())
	if Conf.UseTLS {
		r.Use(TlsHandler())
	}

	Db = DbConn()
	user := r.Group("/user")
	{
		user.POST("/login", Login)
		user.POST("/logout", Logout)
		user.POST("/register", Signup)
		user.POST("/userinfo", GetUserInfo)
	}
	article := r.Group("/article")
	article.Use(UserVerifyMiddleware())
	{
		article.POST("/list", GetArticles)
		article.POST("/create", UploadArticle)
		article.DELETE("/delete", RemoveArticle)

		article.POST("/:id", GetArticle)
	}

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

func GetPwd(pwd string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return hash, err
}

func ComparePwd(pwd1 string, pwd2 string) bool {
	// Returns true on success, pwd1 is for the database.
	err := bcrypt.CompareHashAndPassword([]byte(pwd1), []byte(pwd2))
	if err != nil {
		return false
	} else {
		return true
	}
}

func GrantPermission(username string, isAdmin bool, c *gin.Context) {
	c.SetCookie("username", username, int(3*time2.Hour), "/", Conf.Domain, true, false)
	c.SetCookie("login_token", CreateToken("usertoken", username), int(3*time2.Hour), "/", Conf.Domain, true, false)
	if isAdmin == true {
		c.SetCookie("admin_token", CreateToken("admintoken", username), int(3*time2.Hour), "/", Conf.Domain, true, false)
	}
}

func VerifyUserIfAdmin(username string, lToken string, aToken string) bool {
	if VerifyToken(lToken, "usertoken", username) && VerifyToken(aToken, "admintoken", username+" is admin") {
		return true
	}
	return false
}

func CreateToken(tokenName string, tokenCtx string) string {
	token, err := JwtEncoder(tokenName, tokenCtx, int64(3*time2.Hour))
	if err != nil {
		log.Println("Create Token Failed! name:" + tokenName + " ctx:" + tokenCtx)
		return ""
	}
	return token
}

func VerifyToken(token string, tokenName string, tokenCtx string) bool {
	name, ctx, err := JwtDecoder(token)
	if err != nil {
		log.Println("Verify Token Failed! Decoder Error!")
		return false
	}
	if name != tokenName {
		log.Println("Verify Token Failed! name Not Valid!")
	} else {
		if ctx == tokenCtx {
			return true
		}
	}

	return false
}
