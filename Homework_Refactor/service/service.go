package service

import (
	"Homework_Refactor/conf"
	"Homework_Refactor/tools"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	Conf conf.Config
)

func init() {
	DB = tools.GetDB()
	Conf = conf.Conf
}
