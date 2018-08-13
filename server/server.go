package server

import (
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

// New creates a new rest server
func New() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", pong)

	r.GET("/createdetective", createDetective)
	r.GET("/createrequest", createRequest)

	r.GET("/changelevel", changeLevel)

	// TODO
	r.GET("/managerequest", manageRequest)

	return r
}
