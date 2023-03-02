package service

import (
	"Homework_Refactor/tools"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/gookit/color"
	"log"
	"strconv"
)

type RequestComment struct {
	CommentText string `gorm:"column:comment_text" json:"commentText"`
	Time        string `gorm:"column:time" json:"time"`
}

func GetComments(c *gin.Context) {
	id := c.Param("id")
	var cList []tools.Comment
	DB.Table("comment").Select("*").Where("article_id = ?", id).Order("thumb_up DESC").Scan(&cList)
	type Status struct {
		Code     int    `json:"code"`
		Msg      string `json:"msg"`
		Comments *[]tools.Comment
	}
	status := Status{
		Code:     200,
		Msg:      "success",
		Comments: &cList,
	}

	jsonR, err := json.Marshal(status)

	if err != nil {
		return
	}
	c.String(200, string(jsonR))
}

func CreateComment(c *gin.Context) {
	id := c.Param("id")
	ip := c.ClientIP()
	var cForm RequestComment
	username := c.GetHeader("username")
	err := c.ShouldBindJSON(&cForm)
	Aid, err := strconv.Atoi(id)
	if err != nil {
		log.Println(ip + " 用户对文章" + id + "评价失败，接收到的JSON格式不正确： " + err.Error())
		c.JSON(200, gin.H{"code": 2071, "msg": "输入格式不正确，请稍后重试"})
		return
	}
	UpdateForm := &tools.Comment{
		ArticleID:   Aid,
		User:        username,
		CommentText: cForm.CommentText,
		Time:        cForm.Time,
		ThumbUp:     0,
	}
	DB.Table("comment").Create(UpdateForm)
	log.Println(color.FgYellow.Render("用户 " + username + " 对文章：\"" + id + "\" 评价：" + cForm.CommentText))
	c.JSON(200, gin.H{"code": 2070, "msg": "success"})
}

func ThumbUp(c *gin.Context) {
	cid := c.Param("id")
	username := c.GetHeader("username")
	var cForm tools.Comment
	DB.Table("comment").Where("id = ?", cid).Find(&cForm)
	nThumbUp := cForm.ThumbUp + 1
	DB.Table("comment").Where("id = ?", cid).Update("thumb_up", nThumbUp)
	log.Println(color.FgGreen.Render("用户 " + username + " 对评论：\"" + cid + "\" 点赞了"))
	c.JSON(200, gin.H{"code": 2080, "msg": "success"})
}

func RemoveComment(c *gin.Context) {
	cid := c.Param("id")
	username := c.GetHeader("username")
	lToken := c.GetHeader("authorization")
	aToken := c.GetHeader("adminauth")
	var cForm tools.Comment
	DB.Table("comment").Where("id = ?", cid).Find(&cForm)
	if tools.VerifyUserIfAdmin(username, lToken, aToken) {
		DB.Table("comment").Where("id = ?", cid).Delete(cForm)
		log.Println(color.FgGreen.Render("管理员 " + username + " 删除评论\"" + cid + "\" "))
		c.JSON(200, gin.H{"code": 2110, "msg": "success"})
		return
	} else {
		if cForm.User != username {
			log.Println(color.FgYellow.Render("用户 " + username + " 试图删除评论\"" + cid + "\" 失败"))
			c.JSON(200, gin.H{"code": 2111, "msg": "删除评论失败：非本人评论"})
			return
		} else {
			DB.Table("comment").Where("id = ?", cid).Delete(cForm)
			log.Println(color.FgGreen.Render("用户 " + username + " 删除评论\"" + cid + "\" "))
			c.JSON(200, gin.H{"code": 2110, "msg": "success"})
			return
		}
	}
}
