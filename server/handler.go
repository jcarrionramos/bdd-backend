package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcarrionramos/bdd-backend/models"
	"github.com/jcarrionramos/bdd-backend/structures"
)

func pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func createDetective(c *gin.Context) {
	var detective structures.Detective

	err := c.ShouldBindJSON(&detective)

	if err != nil {
		c.JSON(http.StatusBadRequest, structures.Response{
			Status: http.StatusBadRequest,
			Meta:   err,
			Data:   err.Error(),
		})
		return
	}

	err = models.InsertDetective(detective)

	if err != nil {
		c.JSON(http.StatusBadRequest, structures.Response{
			Status: http.StatusBadRequest,
			Meta:   err,
			Data:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, structures.Response{
		Status: http.StatusOK,
		Data:   detective,
	})
}

func createRequest(c *gin.Context) {
	var request structures.Request

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, structures.Response{
			Status: http.StatusBadRequest,
			Meta:   err,
			Data:   err.Error(),
		})
		return
	}

	request.Status = structures.StandBy

	err = models.InsertRequest(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, structures.Response{
			Status: http.StatusBadRequest,
			Meta:   err,
			Data:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, structures.Response{
		Status: http.StatusOK,
		Data:   request,
	})
}

func createPosition(c *gin.Context) {
	var position structures.Position

	err := c.ShouldBindJSON(&position)

	if err != nil {
		c.JSON(http.StatusBadRequest, structures.Response{
			Status: http.StatusBadRequest,
			Meta:   err,
			Data:   err.Error(),
		})
		return
	}

	err = models.InsertPosition(position)

	if err != nil {
		c.JSON(http.StatusBadRequest, structures.Response{
			Status: http.StatusBadRequest,
			Meta:   err,
			Data:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, structures.Response{
		Status: http.StatusOK,
		Data:   position,
	})
}

func changePosition(c *gin.Context) {
	var newPosition structures.Position

	err := c.ShouldBindJSON(&newPosition)

	if err != nil {
		c.JSON(http.StatusBadRequest, structures.Response{
			Status: http.StatusBadRequest,
			Meta:   err,
			Data:   err.Error(),
		})
		return
	}

	err = models.ChangePrice(newPosition)

	if err != nil {
		c.JSON(http.StatusBadRequest, structures.Response{
			Status: http.StatusBadRequest,
			Meta:   err,
			Data:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, structures.Response{
		Status: http.StatusOK,
		Data:   "Position changed",
	})
}

func allDetectives(c *gin.Context) {
	detectives, err := models.SelectAllDetectives()

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, structures.Response{
		Status: http.StatusOK,
		Data:   detectives,
	})
}
