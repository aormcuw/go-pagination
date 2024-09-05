package initializers

import (
	"fmt"
	"os"

	"github.com/robbyklein/pages/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SyncDB() {
	DB.AutoMigrate(&models.Person{})
}

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to db")
	}
}
