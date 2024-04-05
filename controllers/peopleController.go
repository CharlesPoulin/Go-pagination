package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/robbyklein/pages/initializers"
	"github.com/robbyklein/pages/models"
	"math"
	"net/http"
	"strconv"
)

type PaginationData struct {
	NextPage     int
	PreviousPage int
	CurrentPage  int
	TotalPages   int
}

func PeopleIndexGET(c *gin.Context) {
	// GET PAGE number
	perPage := 10
	page := 1
	pageStr := c.Param("page")

	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr) // dont care about error here
	}

	// calculate total pages for next
	var totalRows int64
	initializers.DB.Model(&models.Person{}).Count(&totalRows)
	totalpages := math.Ceil(float64(totalRows / int64(perPage)))

	offset := (page - 1) * perPage
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
			TotalPages:   int(totalpages),
		},
	})
}
