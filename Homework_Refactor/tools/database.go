package tools

import (
	"github.com/gookit/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var _db *gorm.DB

type Articles struct {
	ID       int    `gorm:"primaryKey;column:id" json:"id"`
	Title    string `gorm:"column:title" json:"title"`
	Category string `gorm:"column:category" json:"category"`
	Content  string `gorm:"column:content" json:"content"`
	Author   string `gorm:"column:author" json:"author"`
	Time     string `gorm:"column:time" json:"time"`
	Views    string `gorm:"column:views" json:"views"`
	Href     string `gorm:"column:href" json:"href"`
}
type Login struct {
	ID       int    `gorm:"primaryKey;column:id" json:"id"`
	Phone    string `gorm:"column:phone" json:"phone"`
	Email    string `gorm:"column:email" json:"email"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Usertype string `gorm:"column:usertype" json:"usertype"`
}
type LoginNoPassword struct {
	ID       int    `gorm:"primaryKey;column:id" json:"id"`
	Phone    string `gorm:"column:phone" json:"phone"`
	Email    string `gorm:"column:email" json:"email"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"-"`
	Usertype string `gorm:"column:usertype" json:"usertype"`
}
type Comment struct {
	ID          int    `gorm:"primaryKey;column:id" json:"id"`
	ArticleID   int    `gorm:"column:article_id" json:"articleId"`
	User        string `gorm:"column:user" json:"user"`
	CommentText string `gorm:"column:comment_text" json:"commentText"`
	Time        string `gorm:"column:time" json:"time"`
	ThumbUp     int    `gorm:"column:thumb_up" json:"thumbUp"`
}

func init() {
	var (
		err        error
		DBUsername = "root"
		DBPassword = "e89r245z"
		DBLocation = "127.0.0.1:3306"
		schemaName = "homework"
	)
	dsn := DBUsername + ":" + DBPassword + "@tcp(" + DBLocation + ")/" + schemaName + "?charset=utf8mb4&parseTime=True&loc=Local"
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(color.FgRed.Render("??????????????? " + schemaName + "?????????????????????:"))
		log.Println(color.FgRed.Render(err.Error()))
	}
	log.Println(color.FgGreen.Render("??????????????????" + schemaName + "????????????"))
	sqlDB, _ := _db.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)
}

func GetDB() *gorm.DB {
	return _db
}

func CheckIfExist(db *gorm.DB, qKey int, qValue string) bool {
	// qKey 1:phone 2:email 3:username
	if qKey > 3 {
		return false
	}
	var lList Login
	var qType string
	switch qKey {
	case 1:
		qType = "phone"
		break
	case 2:
		qType = "email"
		break
	case 3:
		qType = "username"
		break
	default:
		return true
	}
	db.Table("login").Select("*").Where(qType+" = ?", qValue).Scan(&lList)
	if lList.ID != 0 {
		return true
	} else {
		return false
	}
}

func GetFromPhone(db *gorm.DB, phone string) (results []*Login, err error) {
	err = db.Table("login").Model(Login{}).Where("`phone` = ?", phone).Find(&results).Error
	return
}
func GetFromEmail(db *gorm.DB, email string) (results []*Login, err error) {
	err = db.Table("login").Model(Login{}).Where("`email` = ?", email).Find(&results).Error
	return
}
