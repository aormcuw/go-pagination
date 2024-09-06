package controllers

import (
	"fmt"
	"math"
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
	TotalPages   int
	TwoBelow     int
	TwoAfter     int
}

func PeopleIndexGET(c *gin.Context) {
	// Get the page number
	perPage := 10
	page := 1
	pageString := c.Param("page")
	if pageString != "" {
		page, _ = strconv.Atoi(pageString)
	}
	// Calculate total pages
	var totalRows int64
	initializers.DB.Model(&models.Person{}).Count(&totalRows)
	totalPages := math.Ceil((float64(totalRows) / float64(perPage)))
	fmt.Println(totalPages)

	// Calculate the offset
	offset := (page - 1) * perPage

	// Get the people
	var people []models.Person
	initializers.DB.Limit(perPage).Offset(offset).Find(&people)

	// Render the page
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"people": people,
		"pagination": PaginationData{
			PreviousPage: page - 1,
			CurrentPage:  page,
			NextPage:     page + 1,
			TotalPages:   int(totalPages),
			TwoBelow:     page - 2,
			TwoAfter:     page + 2,
		},
	})
}
