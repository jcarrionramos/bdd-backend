package models

import (
	"database/sql"
	"log"
)

var db *sql.DB

func InitDB(str string) error {
	var err error
	db, err = sql.Open("sqlite3", str)
	if err != nil {
		log.Println(err)
		return err
	}
	if err = db.Ping(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
