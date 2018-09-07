package main

import (
	"log"

	"github.com/jcarrionramos/bdd-backend/models"
	"github.com/jcarrionramos/bdd-backend/server"
)

func main() {
	log.SetFlags(log.Lshortfile)
	log.Println("Starting API-MDD REST on port 4242")

	server := server.New()
	err := models.InitDB("./models/database.db")

	if err != nil {
		log.Println(err)
		return
	}

	server.Run(":4242")
}
