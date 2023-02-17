package service

import (
	"Homework_Refactor/tools"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"log"
	"time"
)

type RequestLoginForm struct {
	Usr       string `json:"username"`
	Pass      string `json:"password"`
	LoginType int    `json:"login_type"`
}
type RequestSignupForm struct {
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Usertype string `json:"usertype"`
}

func Login(c *gin.Context) {
	var lForm RequestLoginForm
	err := c.ShouldBindJSON(&lForm)
	if err != nil {
		log.Println(c.ClientIP() + " User Login Failed: Wrong JSON Format " + err.Error())
		c.JSON(200, gin.H{"code": 2001, "msg": "登录失败：输入格式不正确"})
		return
	}
	if lForm.Usr == "" || lForm.Pass == "" {
		c.JSON(200, gin.H{"code": 2002, "msg": "登录失败：用户名或密码不能为空"})
		return
	}
	var res []*tools.Login
	if lForm.LoginType == 1 {
		// Use phone
		res, err = tools.GetFromPhone(DB, lForm.Usr)
		if err != nil {

			c.JSON(200, gin.H{"code": 2003, "msg": "登录失败：请检查用户名或密码是否正确"})
			return
		}
	} else {
		// Use Email
		res, err = tools.GetFromEmail(DB, lForm.Usr)
		if err != nil {
			c.JSON(200, gin.H{"code": 2003, "msg": "登录失败：请检查用户名或密码是否正确"})
			return
		}
	}
	if len(res) > 1 {
		c.JSON(200, gin.H{"code": 2004, "msg": "登录失败：存在多个对应账号，请尝试使用其他方式登录"})
		return
	}
	lDBForm := res[0]
	if err != nil {
		return
	}
	if tools.ComparePwd(lDBForm.Password, lForm.Pass) {
		log.Println(lDBForm.Username + time.Now().Format("2023-02-01"))
		loginToken := tools.CreateToken("user_token", lDBForm.Username+time.Now().Format("2023-02-01"))
		adminToken := "NOT ADMIN"
		if lDBForm.Usertype == "admin" {
			adminToken = tools.CreateToken("admin_token", lDBForm.Username+" is admin")
		}
		lJSON := gin.H{"username": lDBForm.Username, "login_token": loginToken, "admin_token": adminToken}
		log.Printf(color.FgGreen.Render(lDBForm.Username + " 用户成功登录！"))
		c.JSON(200, gin.H{"code": 2000, "msg": "success", "data": lJSON})
		return
	} else {
		c.JSON(200, gin.H{"code": 2003, "msg": "登录失败：请检查用户名或密码是否正确"})
		return
	}
}

func Signup(c *gin.Context) {
	ip := c.ClientIP()
	var sForm RequestSignupForm
	err := c.ShouldBindJSON(&sForm)
	if err != nil {
		log.Println(ip + " User Signup Failed: Wrong JSON Format " + err.Error())
		c.JSON(200, gin.H{"code": 2011, "msg": "注册失败：输入格式不正确"})
		return
	}
	sForm.Usertype = "user"

	if sForm.Username == "" || sForm.Email == "" || sForm.Phone == "" || sForm.Password == "" {
		c.JSON(200, gin.H{"code": 2012, "msg": "注册失败：注册部分不能有空值"})
		return
	}
	// TODO 校验post得到的格式是否符合规范：增加安全性
	if tools.CheckIfExist(DB, 1, sForm.Phone) || tools.CheckIfExist(DB, 2, sForm.Email) || tools.CheckIfExist(DB, 3, sForm.Username) {
		log.Println(ip + " User Signup Failed: Existed phone or email")
		c.JSON(200, gin.H{"code": 2013, "msg": "注册失败：用户已存在 请尝试更改邮箱、手机号或用户名"})
		return
	}
	encrypted, _ := tools.GetPwd(sForm.Password)
	sForm.Password = string(encrypted)

	DB.Table("login").Create(sForm)
	log.Println(color.FgGreen.Render(ip + " User Signup Method Called!"))
	c.JSON(200, gin.H{"code": 2010, "msg": "success"})
	return
}

func Logout(c *gin.Context) {
	usrName, err := c.Cookie("username")
	if err != nil {
		c.JSON(200, gin.H{"code": 2021, "msg": "登出失败：可能登录已过期"})
		return
	}
	c.SetCookie("username", "", 0, "/", Conf.Domain, true, false)
	c.SetCookie("login_token", "", 0, "/", Conf.Domain, true, false)
	c.SetCookie("admin_token", "", 0, "/", Conf.Domain, true, false)
	log.Printf(color.FgGreen.Render(usrName + " 用户成功登出！"))
	c.JSON(200, gin.H{"code": 2020, "msg": "success"})
}

func GetUserInfo(c *gin.Context) {
	username := c.GetHeader("username")
	lToken := c.GetHeader("authorization")
	if tools.VerifyUser(username, lToken) {
		var lList []tools.LoginNoPassword
		DB.Table("login").Select("*").Where("username = ?", username).Scan(&lList)
		type Status struct {
			Code  int    `json:"code"`
			Msg   string `json:"msg"`
			AForm *[]tools.LoginNoPassword
		}

		status := Status{
			Code:  200,
			Msg:   "success",
			AForm: &lList,
		}

		jsonR, err := json.Marshal(status)

		if err != nil {
			return
		}
		c.String(200, string(jsonR))
	} else {
		c.JSON(200, gin.H{"code": 2091, "msg": "登录已过期或异常登录：请重新登录！"})
		return
	}
}
func EditUserInfo(c *gin.Context) {
	ip := c.ClientIP()
	username := c.GetHeader("username")
	lToken := c.GetHeader("authorization")
	if tools.VerifyUser(username, lToken) {
		var sForm RequestSignupForm
		err := c.ShouldBindJSON(&sForm)
		if err != nil {
			log.Println(ip + " User Signup Failed: Wrong JSON Format " + err.Error())
			c.JSON(200, gin.H{"code": 2102, "msg": "输入格式不正确"})
			return
		}
		if sForm.Username == "" || sForm.Email == "" || sForm.Phone == "" || sForm.Password == "" {
			c.JSON(200, gin.H{"code": 2013, "msg": "注册部分不能有空值"})
			return
		}
		encrypted, _ := tools.GetPwd(sForm.Password)
		sForm.Password = string(encrypted)

		if tools.CheckIfExist(DB, 1, sForm.Phone) && tools.CheckIfExist(DB, 2, sForm.Email) && tools.CheckIfExist(DB, 3, sForm.Username) {
			DB.Table("login").Where("username = ?", username).Updates(sForm)
			log.Println(color.FgGreen.Render(ip + " User Edit Info Method Called!"))
			c.JSON(200, gin.H{"code": 2100, "msg": "success"})
			return
		} else {
			c.JSON(200, gin.H{"code": 2104, "msg": "用户不存在"})
			return
		}
	} else {
		c.JSON(200, gin.H{"code": 2101, "msg": "登录已过期或异常登录：请重新登录！"})
		return
	}
}
