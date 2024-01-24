package database

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"noteify-api/models"
	"os"
)

var DB *gorm.DB
var err error

func ConnectDB() (*gorm.DB, error) {
	err = godotenv.Load()
	dsn := os.Getenv("DSN")
	log.Print("DSN: ", dsn)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB.AutoMigrate(&models.User{}, &models.Note{})

	return DB, nil
}
