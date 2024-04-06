package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/robbyklein/pages/helpers"
	"github.com/robbyklein/pages/initializers"
	"github.com/robbyklein/pages/models"
	"net/http"
	"strconv"
)

type PaginationData struct {
	NextPage     int
	PreviousPage int
	CurrentPage  int
	TotalPages   int
	TwoAfter     int
	TwoBelow     int
	ThreeAfter   int
}

func PeopleIndexGET(c *gin.Context) {
	// GET PAGE number
	perPage := 10
	page := 1
	pageStr := c.Param("page")

	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr) // dont care about error here
	}

	pagination := helpers.GetPaginationData(page, perPage, &models.Person{}, "/people")

	// Get the people
	var people []models.Person
	initializers.DB.Limit(10).Offset(pagination.Offset).Find(&people)

	// Render the page
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"people":     people,
		"pagination": pagination,
	})
}
