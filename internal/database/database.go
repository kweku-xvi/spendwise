package database

import (
	"fmt"
	"log"

	"github.com/kweku-xvi/spendwise/api/v1/models"
	"github.com/kweku-xvi/spendwise/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", config.ENV.DBHost, config.ENV.DBUser, config.ENV.DBPassword, config.ENV.DBName, config.ENV.DBPort)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database:", err)
	}

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("failed to run database migrations:", err)
	}

	fmt.Println("Connected to database and migrations applied!")
}
