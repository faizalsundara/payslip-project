package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbUsername := os.Getenv("DB_Username")
	dbPassword := os.Getenv("DB_Password")
	dbPort := os.Getenv("DB_Port")
	dbHost := os.Getenv("DB_Host")
	dbName := os.Getenv("DB_Name")

	fmt.Println("dbUsername--", dbUsername)
	dbPostgres := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost,
		dbUsername,
		dbPassword,
		dbName,
		dbPort,
	)
	db, err := gorm.Open(postgres.Open(dbPostgres), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to Connect to database:", err)
	}

	DB = db
}
