package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcarrionramos/mdd-backend/models"
	"github.com/jcarrionramos/mdd-backend/structures"
)

func pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func createDetective(c *gin.Context) {
	detective := structures.Detective{
		ID:         c.Query("rut"),
		Name:       c.Query("name"),
		Lastname:   c.Query("lastname"),
		Address:    c.Query("address"),
		City:       c.Query("city"),
		PostalCode: c.Query("postalcode"),
		Phone:      c.Query("phone"),
		Device:     c.Query("device"),
		Level:      0,
	}

	models.InsertDetective(detective)

	c.JSON(http.StatusOK, detective)
}

func createRequest(c *gin.Context) {
	request := structures.Request{
		ID:          c.Query("id"),
		DetectiveID: c.Query("detective_id"),
		Date:        c.Query("date"),
		Description: c.Query("description"),
		Curriculum:  c.Query("curriculum"),
		Status:      structures.Accepted,
	}

	c.JSON(http.StatusOK, request)

}

func manageRequest(c *gin.Context) {
	//TODO: ALL!
}

func changeLevel(c *gin.Context) {
	//TODO: ALL!
}
