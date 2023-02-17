package service

import (
	"Homework_Refactor/tools"
	json2 "encoding/json"
	"github.com/gin-gonic/gin"
)

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

	json, err := json2.Marshal(status)

	if err != nil {
		return
	}
	c.String(200, string(json))
}
func CreateComment(c *gin.Context) {
	id := c.Param("id")
	var cList []tools.Comment
	DB.Table("comment").Select("*").Where("article_id = ?", id).Order("thumb_up DESC").Scan(&cList)
	type Status struct {
		Code     int    `json:"code"`
		Msg      string `json:"msg"`
		Comments *[]tools.Comment
	}
	status := Status{
		Code: 200,
		Msg:  "success",
	}

	json, err := json2.Marshal(status)

	if err != nil {
		return
	}
	c.String(200, string(json))
}
