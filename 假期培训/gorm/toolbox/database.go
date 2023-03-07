package toolbox

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var _db *gorm.DB

func init() {
	var err error
	username := "root"
	password := "e89r245z"
	host := "127.0.0.1"
	port := 3306
	DBName := "gorm_homework"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, host, port, DBName)
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("连接数据库失败, error=" + err.Error())
	}

	sqlDB, _ := _db.DB()
	sqlDB.SetMaxOpenConns(100)
	// 最大连接数100
	sqlDB.SetMaxIdleConns(20)
	// 最大空闲连接数20
}

func GetDB() *gorm.DB {
	return _db
}
