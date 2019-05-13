package server

import (
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

// New creates a new rest server
func New() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", pong)

	r.POST("/createdetective", createDetective)
	r.POST("/createrequest", createRequest)
	r.POST("/createposition", createPosition)

	r.GET("/changeposition", changePosition)
	r.GET("/alldetectives", allDetectives)

	return r
}
