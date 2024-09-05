package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/robbyklein/pages/initializers"
	"github.com/robbyklein/pages/models"
)

type PaginationData struct {
	PreviousPage int
	CurrentPage  int
	NextPage     int
}

func PeopleIndexGET(c *gin.Context) {
	// Get the page number
	pageString := c.Param("page")
	page, _ := strconv.Atoi(pageString)
	offset := (page - 1) * 10
	// Get the people
	var people []models.Person
	initializers.DB.Limit(10).Offset(offset).Find(&people)

	// Render the page
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"people": people,
		"pagination": PaginationData{
			PreviousPage: page - 1,
			CurrentPage:  page,
			NextPage:     page + 1,
		},
	})
}
