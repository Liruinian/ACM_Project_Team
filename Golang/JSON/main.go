package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type RegForm struct {
	Status  string  `json:"status"`
	Version float32 `json:"version"`
	Name    string  `json:"name"`
	Users   []User  `json:"users"`
}
type User struct {
	// 使用 json 标签
	Name   string   `json:"name"`
	Gender string   `json:"gender"`
	Age    int      `json:"age"`
	Tags   []string `json:"tags"`
}

func main() {
	var rForm RegForm
	ReadJson("./regform.json", &rForm)
	newUser := User{
		Name:   "QuestionMark",
		Gender: "helicopter", //What?
		Age:    32767,        //How?
		Tags:   []string{"I Have", "Sooo Many", "Questions!"},
	}
	rForm.Users = append(rForm.Users, newUser)
	//增加一个名为QuestionMark的用户
	WriteJson("./regform.json", rForm)
	fmt.Println("用户列表：")
	for _, i := range rForm.Users {
		fmt.Println(i.Name)
	}
}

func ReadJson(filePath string, ref interface{}) {
	f, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) || f.IsDir() {
		log.Println("文件不存在或存在同名文件夹")
		return
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("io 错误")
		return
	}
	err = json.Unmarshal(data, ref)
	if err != nil {
		log.Println("反序列化失败")
		return
	}
}

func WriteJson(filePath string, ref interface{}) {
	f, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) || f.IsDir() {
		log.Println("文件不存在或存在同名文件夹")
		return
	}
	file, err := os.OpenFile(filePath, os.O_RDWR, 0755)
	defer file.Close()
	if err != nil && os.IsNotExist(err) || f.IsDir() {
		log.Println("io 错误")
		return
	}
	json, err := json.MarshalIndent(ref, "", "\t")
	if err != nil {
		log.Println("序列化失败")
	}
	file.Write(json)

}
