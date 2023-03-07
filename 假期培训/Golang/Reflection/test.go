package main

import (
	"fmt"
	"reflect"
)

var name string

type Cal struct {
	Num1 int `json:"num_1"`
	Num2 int `json:"num_2"`
}

func (this Cal) GetSub(name string) {
	fmt.Println(name, "完成了", this.Num1, "-", this.Num2, "=", this.Num1-this.Num2)
}
func test(re interface{}) {
	rType := reflect.TypeOf(re)
	rValue := reflect.ValueOf(re)

	for i := 0; i < rType.NumField(); i++ {
		field := rType.Field(i)
		fmt.Printf("结构体 %s 字段值为 %d json: %v\n", field.Name, rValue.Field(i), field.Tag.Get("json"))
	}

	for i := 0; i < rValue.NumMethod(); i++ {
		vOfName := reflect.ValueOf(name)
		var args = []reflect.Value{vOfName}
		rValue.Method(i).Call(args)
	}

}
func main() {
	var a, b int
	_, _ = fmt.Scanf("%d %d %s", &a, &b, &name)
	cal := Cal{a, b}
	test(cal)
}
