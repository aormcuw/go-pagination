package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/robbyklein/pages/helpers"
	"github.com/robbyklein/pages/initializers"
	"github.com/robbyklein/pages/models"
)

func PeopleIndexGET(c *gin.Context) {
	// Get the page number
	perPage := 10
	page := 1
	pageString := c.Param("page")
	if pageString != "" {
		page, _ = strconv.Atoi(pageString)
	}

	pagination := helpers.GetPaginationData(page, perPage, models.Person{}, "/people")

	// Get the people
	var people []models.Person
	initializers.DB.Limit(perPage).Offset(pagination.Offset).Find(&people)

	// Render the page
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"people":     people,
		"pagination": pagination,
	})
}
