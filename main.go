package main

import (
	"log"
	"salaries-payslip/config"
	"salaries-payslip/models"
	"salaries-payslip/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or failed to load")
	}

	config.ConnectDatabase()

	if err := models.AutoMigrate(config.DB); err != nil {
		log.Fatal("Migration failed:", err)
	}
	models.Seed(config.DB)

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
