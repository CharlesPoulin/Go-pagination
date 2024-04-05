package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/robbyklein/pages/initializers"
	"github.com/robbyklein/pages/models"
	"net/http"
	"strconv"
)

type PaginationData struct {
	NextPage     int
	PreviousPage int
	CurrentPage  int
}

func PeopleIndexGET(c *gin.Context) {
	// GET PAGE number
	page := 1
	pageStr := c.Param("page")

	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr) // dont care about error here
	}
	offset := (page - 1) * 10
	// Get the people
	var people []models.Person
	initializers.DB.Limit(10).Offset(offset).Find(&people)

	// Render the page
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"people": people,
		"pagination": PaginationData{
			NextPage:     page + 1,
			PreviousPage: page - 1,
			CurrentPage:  page,
		},
	})
}
