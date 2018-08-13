package server

import (
	"net/http"
	"strconv"

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

	err := models.InsertDetective(detective)

	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, detective)
}

func createRequest(c *gin.Context) {
	request := structures.Request{
		ID:          c.Query("id"),
		DetectiveID: c.Query("detective_id"),
		Date:        c.Query("date"),
		Description: c.Query("description"),
		Curriculum:  c.Query("curriculum"),
		Status:      structures.StandBy,
	}

	// if !models.IsDetective(request.DetectiveID) {
	// 	c.String(401, "Error: there is no detective with that id, please create one")
	// }

	err := models.InsertRequest(request)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, request)
}

func changeLevel(c *gin.Context) {
	level, _ := strconv.Atoi(c.Query("level"))
	price, _ := strconv.Atoi(c.Query("price"))

	err := models.ChangePrice(level, price)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.Status(http.StatusOK)

}

func manageRequest(c *gin.Context) {
	//TODO: ALL!
}
