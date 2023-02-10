package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func mysqlConn() *sql.DB {
	Db, _ := sql.Open("mysql", "root:@a20040207@tcp(127.0.0.1:3306)/login")
	Db.SetConnMaxLifetime(100)
	Db.SetMaxIdleConns(10)
	Db.SetConnMaxLifetime(10 * time.Minute)
	Db.SetConnMaxIdleTime(10 * time.Minute)
	if err := Db.Ping(); err != nil {
		fmt.Println("open database fail")
		return nil
	}
	fmt.Println("connect success")
	return Db
}

func CheckIfExist(db *sql.DB, qKey int, qValue string) bool {
	// qKey 1:phone 2:email 3:username
	if qKey > 3 {
		return false
	}
	rows, err := db.Query("SELECT * FROM login")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id, phone, email, username, password, usertype string
		if err := rows.Scan(&id, &phone, &email, &username, &password, &usertype); err != nil {
			log.Fatal(err)
		}
		if qKey == 1 {
			if phone == qValue {
				return true
			}
		} else if qKey == 2 {
			if email == qValue {
				return true
			}
		} else if qKey == 3 {
			if username == qValue {
				return true
			}
		}
	}
	return false
}
