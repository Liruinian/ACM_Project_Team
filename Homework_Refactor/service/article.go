package service

import (
	"Homework_Refactor/tools"
	json2 "encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"log"
)

type RequestArticleForm struct {
	Edit     bool   `json:"edit"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Content  string `json:"content"`
	Author   string `json:"author"`
	Time     string `json:"time"`
	Views    string `json:"views"`
	Href     string `json:"href"`
	Id       int    `json:"id"`
}

func GetArticles(c *gin.Context) {
	var aList []tools.Articles
	DB.Table("articles").Select("*").Scan(&aList)
	type Status struct {
		Code     int    `json:"code"`
		Msg      string `json:"msg"`
		Articles *[]tools.Articles
	}
	status := Status{
		Code:     200,
		Msg:      "success",
		Articles: &aList,
	}

	json, err := json2.Marshal(status)

	if err != nil {
		return
	}
	c.String(200, string(json))
}

func GetArticle(c *gin.Context) {
	id := c.Param("id")
	var aList []tools.Articles
	DB.Table("articles").Select("*").Where("id = ?", id).Scan(&aList)
	type Status struct {
		Code  int    `json:"code"`
		Msg   string `json:"msg"`
		AForm *[]tools.Articles
	}
	status := Status{
		Code:  200,
		Msg:   "success",
		AForm: &aList,
	}

	json, err := json2.Marshal(status)

	if err != nil {
		return
	}
	c.String(200, string(json))
}

func CreateArticle(c *gin.Context) {
	ip := c.ClientIP()
	username := c.GetHeader("username")
	lToken := c.GetHeader("authorization")
	aToken := c.GetHeader("adminauth")
	if tools.VerifyUserIfAdmin(username, lToken, aToken) {
		var aForm RequestArticleForm
		err := c.ShouldBindJSON(&aForm)
		if err != nil {
			log.Println(ip + " 用户新建文章失败，接收到的JSON格式不正确： " + err.Error())
			c.JSON(200, gin.H{"code": 2041, "msg": "输入格式不正确: 请登录管理员账号后再进行操作"})
			return
		}
		if aForm.Edit != false {
			log.Println(ip + " 用户新建文章失败，属性不正确 ")
			c.JSON(200, gin.H{"code": 2042, "msg": "edit属性不正确"})
			return
		}
		UpdateForm := &tools.Articles{
			Title:    aForm.Title,
			Category: aForm.Category,
			Content:  aForm.Content,
			Author:   aForm.Author,
			Time:     aForm.Time,
			Views:    aForm.Views,
			Href:     aForm.Href,
		}
		DB.Table("articles").Create(UpdateForm)
		log.Println(color.FgRed.Render("[高危操作警告] ") + (color.FgYellow.Render("管理员 " + username + " 增加新文章：\"" + aForm.Title + "\"")))
		c.JSON(200, gin.H{"code": 2040, "msg": "success"})
	} else {
		log.Println(color.FgRed.Render("[高危操作警告] 不允许ip为" + ip + "的游客访问管理员文章接口"))
		c.JSON(200, gin.H{"code": 2049, "msg": "访问受限，请检查是否为管理员账号"})
		return
	}
}

func EditArticle(c *gin.Context) {
	ip := c.ClientIP()
	username := c.GetHeader("username")
	lToken := c.GetHeader("authorization")
	aToken := c.GetHeader("adminauth")
	if tools.VerifyUserIfAdmin(username, lToken, aToken) {
		var aForm RequestArticleForm
		err := c.ShouldBindJSON(&aForm)
		if err != nil {
			log.Println(ip + " 用户修改文章失败，接收到的JSON格式不正确： " + err.Error())
			c.JSON(200, gin.H{"code": 2051, "msg": "输入格式不正确: 请登录管理员账号后再进行操作"})
			return
		}
		if aForm.Edit != true || aForm.Id == 0 {
			log.Println(ip + " 用户修改文章失败，属性不正确 ")
			c.JSON(200, gin.H{"code": 2052, "msg": "edit属性不正确"})
			return
		}
		UpdateForm := &tools.Articles{
			ID:       aForm.Id,
			Title:    aForm.Title,
			Category: aForm.Category,
			Content:  aForm.Content,
			Author:   aForm.Author,
			Time:     aForm.Time,
			Views:    aForm.Views,
			Href:     aForm.Href,
		}
		DB.Table("articles").Updates(UpdateForm)
		log.Println(color.FgRed.Render("[高危操作警告] ") + (color.FgYellow.Render("管理员 " + username + " 修改文章：\"" + aForm.Title + "\"")))
		log.Println(color.FgYellow.Render(UpdateForm))

		c.JSON(200, gin.H{"code": 2050, "msg": "success"})
	} else {
		log.Println(color.FgRed.Render("[高危操作警告] 不允许ip为" + ip + "的游客访问管理员文章接口"))
		c.JSON(200, gin.H{"code": 2059, "msg": "访问受限，请检查是否为管理员账号"})
		return
	}
}

func RemoveArticle(c *gin.Context) {
	ip := c.ClientIP()
	username := c.GetHeader("username")
	lToken := c.GetHeader("authorization")
	aToken := c.GetHeader("adminauth")
	if tools.VerifyUserIfAdmin(username, lToken, aToken) {
		id := c.Param("id")
		if id != "" {
			DB.Table("articles").Delete("*", "id = ?", id)
			log.Println(color.FgRed.Render("[高危操作警告] ") + (color.FgYellow.Render("管理员 " + username + " 删除文章：\"" + id + "\"")))

			c.JSON(200, gin.H{"code": 2060, "msg": "success"})
		} else {
			log.Println(color.FgRed.Render("[高危操作警告] ") + (color.FgYellow.Render("管理员 " + username + " 删除文章：\"" + id + "\" 失败：id不能为空")))

			c.JSON(200, gin.H{"code": 2061, "msg": "删除文章失败：id不能为空"})
		}
	} else {
		log.Println(color.FgRed.Render("[高危操作警告] 不允许ip为" + ip + "的游客访问管理员文章接口"))
		c.JSON(200, gin.H{"code": 2069, "msg": "访问受限，请检查是否为管理员账号"})
		return
	}
}
