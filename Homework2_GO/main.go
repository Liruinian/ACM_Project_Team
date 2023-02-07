package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	ginServer = gin.Default()
)

type LoginForm struct {
	usr     any
	pass    any
	usrType any
}

func mysqlConn() *sql.DB {
	Db, _ := sql.Open("mysql", "root:@a20040207@tcp(127.0.0.1:3306)/articles")
	Db.SetConnMaxLifetime(100)
	Db.SetMaxIdleConns(10)
	if err := Db.Ping(); err != nil {
		fmt.Println("open database fail")
		return nil
	}
	fmt.Println("connect success")
	return Db
}
func DbOperation(Db *sql.DB, method string, sqlStr string, arg1 string, arg2 string) {
	switch method {
	case "Exec":
		_, err := Db.Exec(sqlStr, arg1, arg2)
		if err != nil {
			fmt.Printf("sqlOperation failed , err:%v\n", err)
			return
		}
	case "Query":
		ctx := context.Background()
		rows, err := Db.QueryContext(ctx, sqlStr, arg1, arg2)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("id: %v  name:%v\n", id, name)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Printf("Unknown Operation")
	}
}

func login() {
	ginServer.POST("/login", func(context *gin.Context) {
		var lForm LoginForm

		lForm.usr, _ = context.Get("username")
		lForm.pass, _ = context.Get("password")
		lForm.usrType, _ = context.Get("login_type")
		context.JSON(200, nil)
	})
}

func main() {
	Db := mysqlConn()
	login()
	DbOperation(Db)
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello,gin"})
	})

	ginServer.Run(":80")
}
