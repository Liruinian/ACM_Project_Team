package main

import (
	"database/sql"
	"fmt"
)

func main() {
	DB, _ := sql.Open("mysql", "root:@a20040207@tcp(127.0.0.1:3306)/login")
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
}
