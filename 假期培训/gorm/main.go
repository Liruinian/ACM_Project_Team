package main

import (
	"gorm.io/gorm"
	"gorm/toolbox"
	"log"
)

type Test struct {
	ID    int    `gorm:"primaryKey;column:id"`
	UName string `gorm:"column:uname"`
	Data  string `gorm:"column:data"`
}

type User struct {
	ID      int    `gorm:"primaryKey;column:id"`
	Message string `gorm:"column:msg"`
	Name    string `gorm:"column:name"`
	Test    Test   `gorm:"foreignKey:name;references:uname"`
}

func (Test) TableName() string {
	return "test"
}
func (User) TableName() string {
	return "user"
}

// 表结构与初始数据：
// User:
//
//	id(int) name(varchar(20)) msg(text)
//
// Test:
//
//	id(int) name(varchar(20)) data(text)
//
// 其中两个id 为主键
// 外键 user.name -> test.name
// user中存在一条数据： 1,hello,Success
// test中存在一条数据： 1,hello,world

func main() {
	db := toolbox.GetDB()
	t := Test{}
	u := User{}

	newT := &Test{
		UName: "wow",
		Data:  "gorm",
	}
	newU := &User{
		Test:    *newT,
		Message: "Hello! NewT",
	}
	// newU.Name -> newT.Name

	if err := db.Create(newT).Error; err != nil { // 向test插入newT
		log.Println("插入失败 " + err.Error())
		return
	} else {
		if err = db.Table("user").Create(newU).Error; err != nil { // 向user插入newU
			log.Println("插入失败 " + err.Error())
			return
		} else {
			log.Println("1. 插入数据: 插入成功")
		}
	}

	log.Println("2. 查询数据")
	db.Where("uname = ?", "hello").First(&t) // 查询第一条test中名字为hello的数据
	log.Println(t.Data)

	db.First(&t) // 查询test中的第一个数据 (id=1)
	log.Println(t)
	t = Test{}
	u = User{}
	var UList []User
	db.Table("user").Select("*").Scan(&UList) // 查询user中的所有数据
	for i, _ := range UList {
		db.Table("test").Select("*").Where("uname = ?", UList[i].Name).Scan(&(UList[i].Test))
	}
	log.Println(UList)

	log.Println("3. 更新数据")
	oriT := &Test{
		UName: "hello",
		Data:  "world",
	}
	db.Where("id = ?", 1).Take(&oriT) // 先查询一条数据, 保存在模型变量oriT中
	db.Model(&oriT).Update("data", "gorm")
	db.Model(&oriT).Update("data", "world")                 // 更新数据
	db.First(&u, "name = ?", "wow").Update("msg", "Modify") // 通过condition表达式更新数据

	log.Println("4. 删除数据")
	db.Delete(u, "name = ?", "wow") // 因为存在外键关系，需要先删除User中的数据才能删除Test中的
	db.Delete(t, "uname = ?", "wow")

	log.Println("5. 事务处理")
	err := db.Transaction(func(tx *gorm.DB) error {
		// 测试事务处理
		if err := tx.Create(&Test{UName: "Amazing", Data: "Yes"}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}
	db.Delete(t, "uname = ?", "Amazing")
}
