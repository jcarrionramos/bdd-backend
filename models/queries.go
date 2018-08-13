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
		detective.PostalCode, detective.Phone, detective.Device, structures.One,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func InsertRequest(request structures.Request) error {
	sqlstmt := `
		insert into requests (id, detective_id, date, description,
		curriculum, status) values ($1,$2,$3,$4,$5,$6)
	`
	_, err := db.Exec(sqlstmt, request.ID, request.DetectiveID, request.Date,
		request.Description, request.Curriculum, structures.StandBy,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func ChangePrice(level, price int) error {
	sqlstmt := `update levels set price=$1 where level=$2`
	_, err := db.Exec(sqlstmt, price, level)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
