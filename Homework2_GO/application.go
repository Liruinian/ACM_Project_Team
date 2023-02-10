package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
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
	Expiretime int64  `json:"expiration"`
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
		if CheckIfExist(Dbl, 1, sForm.phone) || CheckIfExist(Dbl, 2, sForm.email) || CheckIfExist(Dbl, 3, sForm.username) {
			context.JSON(200, "注册失败：用户已存在 请尝试更改邮箱、手机号或用户名")
			return
		}

		if sForm.username == "" || sForm.email == "" || sForm.phone == "" || sForm.pass == "" {
			context.JSON(200, "注册部分不能有空值")
			return
		}
		result, err := Dbl.Exec("INSERT INTO login(phone,email,username,password,usertype) VALUES (?,?,?,?,?)", sForm.phone, sForm.email, sForm.username, sForm.pass, sForm.usertype)
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Successfully Registered with ID: %d\n", id)
		context.JSON(200, "success")
		return
	})
}

func login(ginServer *gin.Engine) {
	ginServer.POST("/login", func(context *gin.Context) {
		var lForm LoginForm
		var sqlStr, username, password, usertype string
		context.DefaultPostForm("type", "post")
		lForm.usr = context.PostForm("username")
		lForm.pass = context.PostForm("password")
		lForm.loginType = context.PostForm("login_type")
		if lForm.usr == "" || lForm.pass == "" {
			context.JSON(200, "用户名或密码不能为空")
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
			context.JSON(200, "登录失败：请检查用户名或密码是否正确")
			return
		}
		log.Println(username, password, usertype)
		if lForm.pass == password {
			grantPermission(context.ClientIP(), username, usertype)
			context.JSON(200, "success")
			return
		} else {
			context.JSON(200, "登录失败：请检查用户名或密码是否正确")
			return
		}

	})
}
func logout(ginServer *gin.Engine) {
	ginServer.POST("/logout", func(context *gin.Context) {
		ip := context.ClientIP()
		rows, err := Dbl.Query("SELECT ip,expiretime FROM currentlogins")
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			var dbip string
			var extime int64
			if err = rows.Scan(&dbip, &extime); err != nil {
				context.JSON(200, "登出失败")
				log.Fatal(err)
			}
			_, err = Dbl.Exec("DELETE FROM currentlogins WHERE ip LIKE ?", ip)
			context.JSON(200, "success")
		}
	})
}
func getIdentity(ginServer *gin.Engine) {
	ginServer.POST("/userinfo", func(context *gin.Context) {
		userip := context.ClientIP()
		if isPermitted(userip, false) {
			var uInfo UserInfo
			rows := Dbl.QueryRow("SELECT username,usertype,expiretime FROM currentlogins WHERE ip=?", userip)
			err := rows.Scan(&uInfo.Username, &uInfo.Usertype, &uInfo.Expiretime)
			if err != nil {
				log.Fatal(err)
			}
			bjson, _ := json.Marshal(uInfo)
			log.Println(string(bjson))
			context.JSON(200, string(bjson))
		} else {
			context.JSON(200, "Please Login")
		}
	})
}
func getArticles(ginServer *gin.Engine) {
	ginServer.POST("/get-articles", func(context *gin.Context) {
		if isPermitted(context.ClientIP(), false) {

			art, err := getJSON(Dba, "SELECT * FROM articles")
			if err != nil {
				context.JSON(200, err)
				return
			}
			context.JSON(200, art)
		} else {
			context.JSON(200, "Please Login")
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
			context.JSON(200, "成功删除文章，文章id:"+aId)
		} else {
			context.JSON(200, "Please Login With Admin Account")
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
				context.JSON(200, "成功创建新文章")
			} else {
				var Aid = context.PostForm("id")
				_, err := Dba.Exec("UPDATE articles SET title=?,category=?,content=?,author=?,time=?,views=?,href=? WHERE id=?", aForm.title, aForm.category, aForm.content, aForm.author, aForm.time, aForm.views, aForm.href, Aid)

				if err != nil {
					context.JSON(200, "Modify Article FAILED")
					log.Fatal(err)
				}
				context.JSON(200, "成功修改文章")
			}
		} else {
			context.JSON(200, "Please Login With Admin Account")
		}
	})
}
func grantPermission(ip, username, usertype string) {
	log.Println("Permission Added For ", ip, username, usertype)

	time := time2.Now().Unix() + 3600

	result, err := Dbl.Exec("INSERT INTO currentlogins(ip,username,usertype,expiretime) VALUES (?,?,?,?)", ip, username, usertype, time)
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully Logged in with ID: %d\n", id)
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
