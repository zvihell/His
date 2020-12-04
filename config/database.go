package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DBClient *sql.DB

func InitialiseDBConnection() {
	db, err := sql.Open("mysql", "root:88888888@tcp(localhost:3306)/test_his?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	DBClient = db
}
