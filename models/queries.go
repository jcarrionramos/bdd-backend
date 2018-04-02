package models

import (
	"log"

	"github.com/jcarrionramos/mdd-backend/structures"
	_ "github.com/mattn/go-sqlite3"
)

// db is a global variable defined in /models/db.go

func InsertDetective(detective structures.Detective) error {
	sqlstmt := `
		insert into detectives (id, name, lastname, address,
		city, postalcode, phone, device, level) values
		($1,$2,$3,$4,$5,$6,$7,$8,$9)
	`
	_, err := db.Exec(sqlstmt, detective.ID, detective.Name,
		detective.Lastname, detective.Address, detective.City,
		detective.PostalCode, detective.Phone, detective.Device, detective.Level,
	)
	if err != nil {
		log.Println(err)
	}
	return nil
}
