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
	Usr       string `json:"username"`
	Pass      string `json:"password"`
	LoginType int    `json:"login_type"`
}
type SignupForm struct {
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Pass     string `json:"password"`
	Usertype string `json:"usertype"`
}
type UserInfo struct {
	Username   string `json:"username"`
	Usertype   string `json:"usertype"`
	ExpireTime int64  `json:"expiration"`
}
type ArticleForm struct {
	Edit     bool   `json:"edit"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Content  string `json:"content"`
	Author   string `json:"author"`
	Time     string `json:"time"`
	Views    string `json:"views"`
	Href     string `json:"href"`
}

func Signup(ginServer *gin.Engine) {
	ginServer.POST("/signup", func(context *gin.Context) {
		ip := context.ClientIP()
		var sForm SignupForm
		err := context.ShouldBindJSON(&sForm)
		if err != nil {
			log.Println(ip + " User Signup Failed: Wrong JSON Format " + err.Error())
			context.JSON(200, gin.H{"status": "输入格式不正确"})
			return
		}
		sForm.Usertype = "user"
		if sForm.Username == "" || sForm.Email == "" || sForm.Phone == "" || sForm.Pass == "" {
			context.JSON(200, gin.H{"status": "注册部分不能有空值"})
			return
		}
		if CheckIfExist(Dbl, 1, sForm.Phone) || CheckIfExist(Dbl, 2, sForm.Email) || CheckIfExist(Dbl, 3, sForm.Username) {
			log.Println(ip + " User Signup Failed: Existed phone or email")
			context.JSON(200, gin.H{"status": "注册失败：用户已存在 请尝试更改邮箱、手机号或用户名"})
			return
		}
		encrypted, _ := GetPwd(sForm.Pass)
		sForm.Pass = string(encrypted)

		result, err := Dbl.Exec("INSERT INTO login(phone,email,username,password,usertype) VALUES (?,?,?,?,?)", sForm.Phone, sForm.Email, sForm.Username, sForm.Pass, sForm.Usertype)
		id, err := result.LastInsertId()
		if err != nil {
			log.Println(color.FgRed.Render(err.Error()))
		}
		log.Println(color.FgGreen.Render(ip + " User Signup Success"))
		fmt.Printf("Successfully Registered with ID: %d\n", id)
		context.JSON(200, gin.H{"status": "success"})
		return
	})
}

func Login(ginServer *gin.Engine) {
	ginServer.POST("/login", func(context *gin.Context) {
		ip := context.ClientIP()
		var lForm LoginForm
		var sqlStr, username, password, usertype string
		err := context.ShouldBindJSON(&lForm)
		if err != nil {
			log.Println(context.ClientIP() + " User Login Failed: Wrong JSON Format " + err.Error())
			context.JSON(200, gin.H{"status": "输入格式不正确"})
			return
		}
		if lForm.Usr == "" || lForm.Pass == "" {
			context.JSON(200, gin.H{"status": "用户名或密码不能为空"})
			return
		}
		if lForm.LoginType == 1 {
			// Use phone
			sqlStr = "SELECT username,password,usertype FROM login WHERE phone=?"

		} else {
			// Use Email
			sqlStr = "SELECT username,password,usertype FROM login WHERE email=?"

		}

		rows := Dbl.QueryRow(sqlStr, lForm.Usr)
		err = rows.Scan(&username, &password, &usertype)
		if err != nil {
			log.Println(ip + " User Login Failed " + err.Error())
			context.JSON(200, gin.H{"status": "登录失败：请检查用户名或密码是否正确"})
			return
		}

		if ComparePwd(password, lForm.Pass) {
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
func Logout(ginServer *gin.Engine) {
	ginServer.POST("/logout", func(context *gin.Context) {
		ip := context.ClientIP()
		_, err := Dbl.Exec("DELETE FROM currentlogins WHERE ip LIKE ?", ip)
		if err != nil {
			context.JSON(200, gin.H{"status": "Logout Failed"})
			log.Println(color.FgRed.Render(err.Error()))
		}
		log.Println(ip + " User Logout")
		context.JSON(200, gin.H{"status": "success"})
	})
}
func GetIdentity(ginServer *gin.Engine) {
	ginServer.GET("/userinfo", func(context *gin.Context) {
		ip := context.ClientIP()
		if IsPermitted(ip, false) {
			var uInfo UserInfo
			rows := Dbl.QueryRow("SELECT username,usertype,expiretime FROM currentlogins WHERE ip=?", ip)
			err := rows.Scan(&uInfo.Username, &uInfo.Usertype, &uInfo.ExpireTime)
			if err != nil {
				log.Println(color.FgRed.Render(err.Error()))
			}
			bJson, _ := json.Marshal(uInfo)
			log.Println(ip + " User Get Identity: " + string(bJson))
			context.JSON(200, string(bJson))
		} else {
			context.JSON(200, gin.H{"status": "登录已过期：请重新登录！"})
		}
	})
}
func GetArticles(ginServer *gin.Engine) {
	ginServer.GET("/articles", func(context *gin.Context) {
		ip := context.ClientIP()
		if IsPermitted(ip, false) {
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
func GetArticleList(ginServer *gin.Engine) {
	ginServer.GET("/article-list", func(context *gin.Context) {
		ip := context.ClientIP()
		if IsPermitted(ip, false) {
			art, err := getJSON(Dba, "SELECT id, title FROM articles")
			if err != nil {
				context.JSON(200, gin.H{"status": "获取文章失败，请稍后重试"})
				return
			}
			log.Println(ip + " User Get Article List")
			context.JSON(200, art)
		} else {
			context.JSON(200, gin.H{"status": "登录已过期：请重新登录！"})
		}
	})
}
func GetArticle(ginServer *gin.Engine) {
	ginServer.GET("/articles/:id", func(context *gin.Context) {
		ip := context.ClientIP()
		if IsPermitted(ip, false) {
			id := context.Param("id")
			sqlStr := "SELECT * FROM articles WHERE id = " + id + ";"
			art, err := getJSON(Dba, sqlStr)
			if err != nil {
				context.JSON(200, gin.H{"status": "获取文章失败，请稍后重试"})
				return
			}
			log.Println(ip + " User Get Articles id = " + id)
			context.JSON(200, art)
		} else {
			context.JSON(200, gin.H{"status": "登录已过期：请重新登录！"})
		}
	})
}
func RemoveArticle(ginServer *gin.Engine) {
	ginServer.DELETE("/articles/:id", func(context *gin.Context) {
		ip := context.ClientIP()
		aId := context.Param("id")
		if IsPermitted(ip, true) {
			_, err := Dba.Exec("DELETE FROM articles WHERE id = ?", aId)
			if err != nil {
				context.JSON(200, gin.H{"status": "删除文章出错：" + err.Error()})
				log.Println(color.FgRed.Render(err.Error()))
			}
			log.Println(color.FgYellow.Render(ip + " Admin Remove Articles, id=" + aId))
			context.JSON(200, gin.H{"status": "成功删除文章，文章id:" + aId})
		} else {
			context.JSON(200, gin.H{"status": "请以管理员身份登录系统"})
		}
	})
}

func UploadArticle(ginServer *gin.Engine) {
	ginServer.PUT("/articles", func(context *gin.Context) {
		ip := context.ClientIP()
		if IsPermitted(ip, true) {
			var aForm ArticleForm
			err := context.ShouldBindJSON(&aForm)
			if err != nil {
				log.Println(context.ClientIP() + " Admin Upload Article Failed: Wrong JSON Format " + err.Error())
				context.JSON(200, gin.H{"status": "提交json格式不正确"})
				return
			}
			if aForm.Edit == false {
				_, err = Dba.Exec("INSERT INTO articles (title,category,content,author,time,views,href) VALUES (?,?,?,?,?,?,?)", aForm.Title, aForm.Category, aForm.Content, aForm.Author, aForm.Time, aForm.Views, aForm.Href)
				if err != nil {
					log.Println(color.FgRed.Render(err.Error()))
				}
				context.JSON(200, gin.H{"status": "成功创建新文章"})
			} else {
				var Aid = context.PostForm("id")
				_, err = Dba.Exec("UPDATE articles SET title=?,category=?,content=?,author=?,time=?,views=?,href=? WHERE id=?", aForm.Title, aForm.Category, aForm.Content, aForm.Author, aForm.Time, aForm.Views, aForm.Href, Aid)

				if err != nil {
					log.Println(color.FgRed.Render(err.Error()))
				}
				log.Println(color.FgYellow.Render(ip + " User Modify Articles"))
				context.JSON(200, gin.H{"status": "成功修改文章"})
			}
		} else {
			context.JSON(200, gin.H{"status": "请以管理员身份登录系统"})
		}
	})
}
func SearchArticle(ginServer *gin.Engine) {
	ginServer.GET("/search-articles", func(context *gin.Context) {
		ip := context.ClientIP()
		if IsPermitted(ip, false) {
			ser := context.GetString("text")
			log.Println(color.FgGreen.Render(ip + " User Try Search " + ser))
			str, err := getJSON(Dba, "SELECT title FROM articles WHERE title LIKE "+ser)
			if err != nil {
				context.JSON(200, gin.H{"status": "Not Found"})
				log.Println(color.FgYellow.Render(ip + " User Search " + ser + " Failed " + err.Error()))
				return
			}
			context.JSON(200, str)
		} else {
			context.JSON(200, gin.H{"status": "登录已过期：请重新登录！"})
		}
	})
}
func grantPermission(ip, username, usertype string) int64 {
	time := time2.Now().Unix() + 3600

	result, err := Dbl.Exec("INSERT INTO currentlogins(ip,username,usertype,expiretime) VALUES (?,?,?,?)", ip, username, usertype, time)
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(color.FgRed.Render(err.Error()))
	}
	return id
}

func IsPermitted(ip string, isAdminOnly bool) bool {

	rows, err := Dbl.Query("SELECT ip,usertype,expiretime FROM currentlogins")
	if err != nil {
		log.Println(color.FgRed.Render(err.Error()))
	}
	defer rows.Close()
	for rows.Next() {
		var DbIp, usertype string
		var exTime int64
		if err = rows.Scan(&DbIp, &usertype, &exTime); err != nil {
			return false
		}
		if DbIp == ip {
			if exTime > time2.Now().Unix() {
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
				fmt.Printf("timestamp %d is bigger than %d, signing off\n", time2.Now().Unix(), exTime)
				_, err = Dbl.Exec("DELETE FROM currentlogins WHERE ip LIKE ?", ip)
				if err != nil {
					log.Println(color.FgRed.Render(err.Error()))
				}
			}
		}
	}
	return false
}
