package models

import (
	"database/sql"
	"log"
)

var db *sql.DB

func InitDB(str string) {
	var err error
	db, err = sql.Open("sqlite3", str)
	if err != nil {
		log.Println(err)
	}
	if err = db.Ping(); err != nil {
		log.Println(err)
	}
}
