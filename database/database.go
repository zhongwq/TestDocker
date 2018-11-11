package database

import (
	"github.com/zhongwq/TestDocker/Model"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var err error
var stmt *sql.Stmt

func init() {
	db, err = sql.Open("sqlite3", "./user.db")
	if err != nil {
		fmt.Println("Fail to open database!")
	}

	//创建表
	sql_table := `
    CREATE TABLE IF NOT EXISTS userinfo(
        uid INTEGER PRIMARY KEY AUTOINCREMENT,
        username VARCHAR(64) UNIQUE  NOT NULL,
        password VARCHAR(64) NOT NULL,
        phone VARCHAR(64) NOT NULL,
		email VARCHAR(64) NOT NULL
    );
    `
	db.Exec(sql_table)
}

func InsertUser(username string, password string, phone string, email string) (bool, string) {
	stmt, err := db.Prepare("INSERT INTO userinfo(username, password, phone, email) values(?, ?, ?, ?)")

	res, err := stmt.Exec(username, password, phone, email)
	if err != nil {
		fmt.Println(err.Error())
		return false, "User is exist"
	}

	_ , err = res.LastInsertId()
	if err !=  nil {
		return false, "Fail to insert user"
	}
	return true, "User " + username + " created"
}

func GetUserInfo(username string, password string) (Model.User, string) {
	rows, err := db.Query("SELECT * FROM userinfo WHERE username = ? and password = ?", username, password)
	if err != nil {
		fmt.Println("Error when query data")
		return Model.User{"", "", "", ""}, "Error when query data"
	}

	if rows.Next() {
		var uid int
		var username string
		var password string
		var phone string
		var email string
		err = rows.Scan(&uid, &username, &password, &phone, &email)
		if err == nil {
			fmt.Println(username, password, email, phone)
			return Model.User{username, password, email, phone}, "Create successfully"
		}

		return Model.User{"","","",""}, "Error when get userinfo"
	}


	return Model.User{"","","",""}, "No user match"
}