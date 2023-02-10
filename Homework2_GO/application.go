package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"log"
	time2 "time"
)

type LoginForm struct {
	usr       string
	pass      string
	loginType string
}
type SignupForm struct {
	phone    string
	email    string
	username string
	pass     string
	usertype string
}
type UserInfo struct {
	Username   string `json:"username"`
	Usertype   string `json:"usertype"`
	ExpireTime int64  `json:"expiration"`
}
type ArticleForm struct {
	title    string
	category string
	content  string
	author   string
	time     string
	views    string
	href     string
}

func signup(ginServer *gin.Engine) {
	ginServer.POST("/signup", func(context *gin.Context) {

		var sForm SignupForm
		context.DefaultPostForm("type", "post")
		sForm.phone = context.PostForm("phone")
		sForm.email = context.PostForm("email")
		sForm.username = context.PostForm("username")
		sForm.pass = context.PostForm("password")
		sForm.usertype = "user"
		if sForm.username == "" || sForm.email == "" || sForm.phone == "" || sForm.pass == "" {
			context.JSON(200, gin.H{"status": "注册部分不能有空值"})
			return
		}
		if CheckIfExist(Dbl, 1, sForm.phone) || CheckIfExist(Dbl, 2, sForm.email) || CheckIfExist(Dbl, 3, sForm.username) {
			log.Println(context.ClientIP() + "User Signup Failed: Existed phone or email")
			context.JSON(200, gin.H{"status": "注册失败：用户已存在 请尝试更改邮箱、手机号或用户名"})
			return
		}
		result, err := Dbl.Exec("INSERT INTO login(phone,email,username,password,usertype) VALUES (?,?,?,?,?)", sForm.phone, sForm.email, sForm.username, sForm.pass, sForm.usertype)
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		log.Println(color.FgGreen.Render(context.ClientIP() + " User Signup Success"))
		fmt.Printf("Successfully Registered with ID: %d\n", id)
		context.JSON(200, gin.H{"status": "success"})
		return
	})
}

func login(ginServer *gin.Engine) {
	ginServer.POST("/login", func(context *gin.Context) {
		ip := context.ClientIP()
		var lForm LoginForm
		var sqlStr, username, password, usertype string
		context.DefaultPostForm("type", "post")
		lForm.usr = context.PostForm("username")
		lForm.pass = context.PostForm("password")
		lForm.loginType = context.PostForm("login_type")
		if lForm.usr == "" || lForm.pass == "" {
			context.JSON(200, gin.H{"status": "用户名或密码不能为空"})
			return
		}
		if lForm.loginType == "1" {
			// Use phone
			sqlStr = "SELECT username,password,usertype FROM login WHERE phone=?"

		} else {
			// Use Email
			sqlStr = "SELECT username,password,usertype FROM login WHERE email=?"

		}

		rows := Dbl.QueryRow(sqlStr, lForm.usr)

		err := rows.Scan(&username, &password, &usertype)
		if err != nil {
			log.Println(ip + " User Login Failed ")
			context.JSON(200, gin.H{"status": "登录失败：请检查用户名或密码是否正确"})
			return
		}
		log.Println(username, password, usertype)
		if lForm.pass == password {
			grantPermission(ip, username, usertype)
			log.Printf(color.FgGreen.Render(ip + " Successfully Logged in with username: " + username))
			context.JSON(200, gin.H{"status": "success"})
			return
		} else {
			log.Println(ip + " User Login Failed")
			context.JSON(200, gin.H{"status": "登录失败：请检查用户名或密码是否正确"})
			return
		}

	})
}
func logout(ginServer *gin.Engine) {
	ginServer.POST("/logout", func(context *gin.Context) {
		ip := context.ClientIP()
		_, err := Dbl.Exec("DELETE FROM currentlogins WHERE ip LIKE ?", ip)
		if err != nil {
			context.JSON(200, gin.H{"status": "Logout Failed"})
			log.Fatal(err)
		}
		log.Println(ip + " User Logout")
		context.JSON(200, gin.H{"status": "success"})
	})
}
func getIdentity(ginServer *gin.Engine) {
	ginServer.POST("/userinfo", func(context *gin.Context) {
		userip := context.ClientIP()
		if isPermitted(userip, false) {
			var uInfo UserInfo
			rows := Dbl.QueryRow("SELECT username,usertype,expiretime FROM currentlogins WHERE ip=?", userip)
			err := rows.Scan(&uInfo.Username, &uInfo.Usertype, &uInfo.ExpireTime)
			if err != nil {
				log.Fatal(err)
			}
			bjson, _ := json.Marshal(uInfo)
			log.Println(userip + " User Get Identity: " + string(bjson))
			context.JSON(200, string(bjson))
		} else {
			context.JSON(200, gin.H{"status": "登录已过期：请重新登录！"})
		}
	})
}
func getArticles(ginServer *gin.Engine) {
	ginServer.POST("/get-articles", func(context *gin.Context) {
		ip := context.ClientIP()
		if isPermitted(ip, false) {

			art, err := getJSON(Dba, "SELECT * FROM articles")
			if err != nil {
				context.JSON(200, gin.H{"status": "获取文章失败，请稍后重试"})
				return
			}
			log.Println(ip + " User Get Articles")
			context.JSON(200, art)
		} else {
			context.JSON(200, gin.H{"status": "登录已过期：请重新登录！"})
		}
	})
}
func removeArticle(ginServer *gin.Engine) {
	ginServer.POST("/delete-art", func(context *gin.Context) {
		userip := context.ClientIP()
		aId := context.PostForm("id")
		if isPermitted(userip, true) {
			_, err := Dba.Exec("DELETE FROM articles WHERE id = ?", aId)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(color.FgYellow.Render(userip + " Admin Remove Articles"))
			context.JSON(200, gin.H{"status": "成功删除文章，文章id:" + aId})
		} else {
			context.JSON(200, gin.H{"status": "请以管理员身份登录系统"})
		}
	})
}

func uploadArticle(ginServer *gin.Engine) {
	ginServer.POST("/upload-art", func(context *gin.Context) {
		userip := context.ClientIP()
		if isPermitted(userip, true) {
			var aForm ArticleForm
			aForm.title = context.PostForm("title")
			aForm.category = context.PostForm("category")
			aForm.content = context.PostForm("content")
			aForm.author = context.PostForm("author")
			aForm.time = context.PostForm("time")
			aForm.views = context.PostForm("views")
			aForm.href = context.PostForm("href")
			if context.PostForm("edit") == "false" {
				_, err := Dba.Exec("INSERT INTO articles (title,category,content,author,time,views,href) VALUES (?,?,?,?,?,?,?)", aForm.title, aForm.category, aForm.content, aForm.author, aForm.time, aForm.views, aForm.href)
				if err != nil {
					log.Fatal(err)
				}
				context.JSON(200, gin.H{"status": "成功创建新文章"})
			} else {
				var Aid = context.PostForm("id")
				_, err := Dba.Exec("UPDATE articles SET title=?,category=?,content=?,author=?,time=?,views=?,href=? WHERE id=?", aForm.title, aForm.category, aForm.content, aForm.author, aForm.time, aForm.views, aForm.href, Aid)

				if err != nil {
					log.Fatal(err)
				}
				log.Println(color.FgYellow.Render(userip + " User Modify Articles"))
				context.JSON(200, gin.H{"status": "成功修改文章"})
			}
		} else {
			context.JSON(200, gin.H{"status": "请以管理员身份登录系统"})
		}
	})
}
func grantPermission(ip, username, usertype string) int64 {
	time := time2.Now().Unix() + 3600

	result, err := Dbl.Exec("INSERT INTO currentlogins(ip,username,usertype,expiretime) VALUES (?,?,?,?)", ip, username, usertype, time)
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	return id
}

func isPermitted(ip string, isAdminOnly bool) bool {

	rows, err := Dbl.Query("SELECT ip,usertype,expiretime FROM currentlogins")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var dbip, usertype string
		var extime int64
		if err = rows.Scan(&dbip, &usertype, &extime); err != nil {
			return false
		}
		if dbip == ip {
			if extime > time2.Now().Unix() {
				if isAdminOnly {
					if usertype == "admin" {
						return true
					} else {
						return false
					}
				} else {
					return true
				}
			} else {
				fmt.Printf("timestamp %d is smaller than %d, signing off\n", time2.Now().Unix(), extime)
				_, err = Dbl.Exec("DELETE FROM currentlogins WHERE ip LIKE ?", ip)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	return false
}
