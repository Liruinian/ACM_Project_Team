package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func costTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		nowTime := time.Now()
		c.Next()
		CostTime := time.Since(nowTime)
		url := c.Request.URL.String()
		fmt.Printf("the request URL %s cost %v\n", url, CostTime)
	}
}

func main() {
	r := gin.Default()
	r.Use(costTime())
	r.LoadHTMLGlob("src/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/data", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "error", "message": "此url仅支持POST访问"})
	})
	r.POST("/data", func(c *gin.Context) {
		loginC, _ := c.Cookie("login")
		if loginC == "Success" {
			c.HTML(http.StatusOK, "data.html", gin.H{"loginOnlyData": "您已成功登录！本段文本仅在登录后显示，登录失败将会显示失败原因！"})
		} else {
			if loginC == "" {
				c.HTML(http.StatusOK, "data.html", gin.H{"loginOnlyData": "未登录，请进行登录后访问/data"})

			} else {
				c.HTML(http.StatusOK, "data.html", gin.H{"loginOnlyData": loginC})
			}
		}

	})
	r.POST("/", func(c *gin.Context) {
		id := c.PostForm("idnum")
		password := c.PostForm("pass")

		if id == "A19220064" && password == "ACMlogin2023" {
			c.SetCookie("login", "Success", 1000, "/", "localhost", true, false)
			// c.JSON(200, gin.H{"status": "ok", "message": "Successfully Logged in"})
		} else {
			c.SetCookie("login", "登录失败！学号或密码有误，请重新尝试登录！", 1000, "/", "localhost", true, false)
			//c.JSON(200, gin.H{"status": "error", "message": "Wrong Student ID or Password"})
		}
		c.Redirect(http.StatusTemporaryRedirect, "/data")
	})
	err := r.Run(":80")
	if err != nil {
		log.Println(err.Error())
		return
	}
}
