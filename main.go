package main

import (
	"log"

	"github.com/jcarrionramos/mdd-backend/models"
	"github.com/jcarrionramos/mdd-backend/server"
)

func main() {
	log.SetFlags(log.Lshortfile)
	log.Println("Starting API-MDD REST on port 4242")
	server := server.New()
	models.InitDB(".???.db")
	server.Run(":4242")
}
