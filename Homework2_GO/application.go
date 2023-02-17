package main

import (
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

func Signup(c *gin.Context) {
	ip := c.ClientIP()
	var sForm SignupForm
	err := c.ShouldBindJSON(&sForm)
	if err != nil {
		log.Println(ip + " User Signup Failed: Wrong JSON Format " + err.Error())
		c.JSON(200, gin.H{"status": "输入格式不正确"})
		return
	}
	sForm.Usertype = "user"
	if sForm.Username == "" || sForm.Email == "" || sForm.Phone == "" || sForm.Pass == "" {
		c.JSON(200, gin.H{"status": "注册部分不能有空值"})
		return
	}
	if CheckIfExist(Db, 1, sForm.Phone) || CheckIfExist(Db, 2, sForm.Email) || CheckIfExist(Dbl, 3, sForm.Username) {
		log.Println(ip + " User Signup Failed: Existed phone or email")
		c.JSON(200, gin.H{"status": "注册失败：用户已存在 请尝试更改邮箱、手机号或用户名"})
		return
	}
	encrypted, _ := GetPwd(sForm.Pass)
	sForm.Pass = string(encrypted)

	result, err := Dbl.Exec("INSERT INTO login.login(phone,email,username,password,usertype) VALUES (?,?,?,?,?)", sForm.Phone, sForm.Email, sForm.Username, sForm.Pass, sForm.Usertype)
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(color.FgRed.Render(err.Error()))
	}
	log.Println(color.FgGreen.Render(ip + " User Signup Success"))
	fmt.Printf("Successfully Registered with ID: %d\n", id)
	c.JSON(200, gin.H{"status": "success"})
	return
}

func Login(c *gin.Context) {
	ip := c.ClientIP()
	var lForm LoginForm
	var sqlStr, username, password, usertype string
	err := c.ShouldBindJSON(&lForm)
	if err != nil {
		log.Println(c.ClientIP() + " User Login Failed: Wrong JSON Format " + err.Error())
		c.JSON(200, gin.H{"status": "输入格式不正确"})
		return
	}
	if lForm.Usr == "" || lForm.Pass == "" {
		c.JSON(200, gin.H{"status": "用户名或密码不能为空"})
		return
	}
	if lForm.LoginType == 1 {
		// Use phone
		sqlStr = "SELECT username,password,usertype FROM login WHERE phone=?"

	} else {
		// Use Email
		sqlStr = "SELECT username,password,usertype FROM login WHERE email=?"

	}

	rows := Db.
	err = rows.Scan(&username, &password, &usertype)
	if err != nil {
		log.Println(ip + " User Login Failed " + err.Error())
		c.JSON(200, gin.H{"status": "登录失败：请检查用户名或密码是否正确"})
		return
	}

	if ComparePwd(password, lForm.Pass) {
		GrantPermission( username,false ,c)
		if usertype == "admin" {
			GrantPermission( username,true ,c)
		}
		log.Printf(color.FgGreen.Render(ip + " Successfully Logged in with username: " + username))
		c.JSON(200, gin.H{"status": "success"})
		return
	} else {
		log.Println(ip + " User Login Failed")
		c.JSON(200, gin.H{"status": "登录失败：请检查用户名或密码是否正确"})
		return
	}

}
func Logout(c *gin.Context) {
	c.SetCookie("username", "", 0, "/", Conf.Domain, true, false)
	c.SetCookie("login_token", "", 0, "/", Conf.Domain, true, false)
	c.SetCookie("admin_token", "", 0, "/", Conf.Domain, true, false)
	ip := c.ClientIP()
	log.Println(ip + " User Logout")
	c.JSON(200, gin.H{"code": 200, "msg": "success"})
}
func GetUserInfo(c *gin.Context) {
	username, err := c.Cookie("username")
	lToken, err := c.Cookie("login_token")
	aToken, err := c.Cookie("admin_token")
	if err != nil {
		return
	}
	if VerifyUserIfAdmin(username,lToken, aToken) {
	} else {
		c.JSON(200, gin.H{"status": "登录已过期：请重新登录！"})
	}
}
func GetArticles(c *gin.Context) {
	ip := c.ClientIP()
		art, err := getJSON(Dba, "SELECT * FROM articles")
		if err != nil {
			c.JSON(200, gin.H{"status": "获取文章失败，请稍后重试"})
			return
		}
		log.Println(ip + " User Get Articles")
		c.JSON(200, art)
	}


func GetArticle(c *gin.Context) {

	ip := c.ClientIP()
		id := c.Param("id")
		sqlStr := "SELECT * FROM articles WHERE id = " + id + ";"
		art, err := getJSON(Dba, sqlStr)
		if err != nil {
			c.JSON(200, gin.H{"status": "获取文章失败，请稍后重试"})
			return
		}
		log.Println(ip + " User Get Articles id = " + id)
		c.JSON(200, art)
	}
}
func RemoveArticle(c *gin.Context) {
	ip := c.ClientIP()
	aId := c.Param("id")
	if VerifyUserIfAdmin(ip, true) {
		_, err := Dba.Exec("DELETE FROM articles WHERE id = ?", aId)
		if err != nil {
			c.JSON(200, gin.H{"status": "删除文章出错：" + err.Error()})
			log.Println(color.FgRed.Render(err.Error()))
		}
		log.Println(color.FgYellow.Render(ip + " Admin Remove Articles, id=" + aId))
		c.JSON(200, gin.H{"status": "成功删除文章，文章id:" + aId})
	} else {
		c.JSON(200, gin.H{"status": "请以管理员身份登录系统"})
	}
}

func UploadArticle(c *gin.Context) {
	ip := c.ClientIP()
	if VerifyUserIfAdmin(ip, true) {
		var aForm ArticleForm
		err := c.ShouldBindJSON(&aForm)
		if err != nil {
			log.Println(c.ClientIP() + " Admin Upload Article Failed: Wrong JSON Format " + err.Error())
			c.JSON(200, gin.H{"status": "提交json格式不正确"})
			return
		}
		if aForm.Edit == false {
			_, err = Dba.Exec("INSERT INTO articles (title,category,content,author,time,views,href) VALUES (?,?,?,?,?,?,?)", aForm.Title, aForm.Category, aForm.Content, aForm.Author, aForm.Time, aForm.Views, aForm.Href)
			if err != nil {
				log.Println(color.FgRed.Render(err.Error()))
			}
			c.JSON(200, gin.H{"status": "成功创建新文章"})
		} else {
			var Aid = c.PostForm("id")
			_, err = Dba.Exec("UPDATE articles SET title=?,category=?,content=?,author=?,time=?,views=?,href=? WHERE id=?", aForm.Title, aForm.Category, aForm.Content, aForm.Author, aForm.Time, aForm.Views, aForm.Href, Aid)

			if err != nil {
				log.Println(color.FgRed.Render(err.Error()))
			}
			log.Println(color.FgYellow.Render(ip + " User Modify Articles"))
			c.JSON(200, gin.H{"status": "成功修改文章"})
		}
	} else {
		c.JSON(200, gin.H{"status": "请以管理员身份登录系统"})
	}
}

