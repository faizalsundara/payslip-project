package models

import (
	"log"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	// Seed admin
	adminID := uuid.New()
	adminPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	admin := User{
		ID:           adminID,
		Username:     "admin",
		PasswordHash: string(adminPassword),
		Role:         "admin",
	}
	db.Create(&admin)

	// Seed 100 employees
	for i := 1; i <= 100; i++ {
		id := uuid.New()
		username := gofakeit.Username() + gofakeit.DigitN(3)
		password, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		salary := float64(gofakeit.Number(5000000, 15000000))

		user := User{
			ID:           id,
			Username:     username,
			PasswordHash: string(password),
			Role:         "employee",
			Salary:       salary,
		}
		if err := db.Create(&user).Error; err != nil {
			log.Println("Error seeding user:", err)
		}
	}
	log.Println("Seeding done.")
}
