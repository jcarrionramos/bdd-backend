package models

import (
	"log"

	"github.com/jcarrionramos/bdd-backend/structures"
)

// db is a global variable defined in /models/db.go

func InsertDetective(detective structures.Detective) error {
	sqlstmt := `
		insert into detectives (id, firstname, lastname,
		city, postalcode, phone, device, position) values
		($1,$2,$3,$4,$5,$6,$7,$8)
	`
	_, err := db.Exec(sqlstmt, detective.ID, detective.FirstName,
		detective.Lastname, detective.City, detective.PostalCode,
		detective.Phone, detective.Device, detective.Position,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func InsertRequest(request structures.Request) error {
	sqlstmt := `
		insert into requests (id, detective_id, created_at, info,
		curriculum, req_status) values ($1,$2,$3,$4,$5,$6)
	`
	_, err := db.Exec(sqlstmt, request.ID, request.DetectiveID, request.CreatedAt,
		request.Info, request.Curriculum, structures.StandBy)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func InsertPosition(position structures.Position) error {
	sqlstmt := `insert into positions (position, price) values ($1,$2)`

	_, err := db.Exec(sqlstmt, position.Position, position.Price)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func ChangePrice(position structures.Position) error {
	sqlstmt := `update levels set price=$1 where position=$2`

	_, err := db.Exec(sqlstmt, position.Price, position.Position)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func SelectAllDetectives() (detectives []structures.Detective, err error) {
	rows, err := db.Query("select id, firstname from detectives")

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		detectives = append(detectives, structures.Detective{
			ID:        id,
			FirstName: name,
		})
	}

	return detectives, nil
}
