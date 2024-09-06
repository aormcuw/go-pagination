package helpers

import (
	"fmt"
	"math"

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
	ThreeAfter   int
	Offset       int
	BaseUrl      string
}

func GetPaginationData(page int, perPage int, model interface{}, baseUrl string) PaginationData {
	// Calculate total pages
	var totalRows int64
	initializers.DB.Model(&models.Person{}).Count(&totalRows)
	totalPages := math.Ceil((float64(totalRows) / float64(perPage)))
	fmt.Println(totalPages)

	offset := (page - 1) * perPage

	return PaginationData{
		PreviousPage: page - 1,
		CurrentPage:  page,
		NextPage:     page + 1,
		TotalPages:   int(totalPages),
		TwoBelow:     page - 2,
		TwoAfter:     page + 2,
		ThreeAfter:   page + 3,
		Offset:       offset,
		BaseUrl:      baseUrl,
	}
}
