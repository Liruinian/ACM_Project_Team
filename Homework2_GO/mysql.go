package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gookit/color"
	"log"
	"time"
)

var (
	Dba *sql.DB
	Dbl *sql.DB
)

func MysqlConn(DbName string) *sql.DB {
	db, _ := sql.Open("mysql", "root:@a20040207@tcp(127.0.0.1:3306)/"+DbName)
	db.SetConnMaxLifetime(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	if err := db.Ping(); err != nil {
		log.Println(color.FgRed.Render("Open " + DbName + " Database Fail"))
		return nil
	}
	log.Println(color.FgGreen.Render("Database " + DbName + " Connect Success"))
	return db
}

func CheckIfExist(db *sql.DB, qKey int, qValue string) bool {
	// qKey 1:phone 2:email 3:username
	if qKey > 3 {
		return false
	}
	rows, err := db.Query("SELECT * FROM login")
	if err != nil {
		log.Println(color.FgRed.Render(err.Error()))
	}
	defer rows.Close()
	for rows.Next() {
		var id, phone, email, username, password, usertype string
		if err = rows.Scan(&id, &phone, &email, &username, &password, &usertype); err != nil {
			log.Println(color.FgRed.Render(err.Error()))
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

func getJSON(db *sql.DB, sqlString string) (string, error) {
	rows, err := db.Query(sqlString)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrS := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrS[i] = &values[i]
		}
		err = rows.Scan(valuePtrS...)
		if err != nil {
			return "", err
		}
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
