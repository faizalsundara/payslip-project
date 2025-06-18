package services

import (
	"fmt"
	"salaries-payslip/config"
	"salaries-payslip/models"
	"salaries-payslip/utils"

	_bcrypt "golang.org/x/crypto/bcrypt"
)

type userService struct{}

func NewUserService() UserInterface {
	return &userService{}
}

func (Us *userService) GetAllUsers(limit, offset int) ([]models.User, int, error) {
	var dataUsers []models.User
	result := config.DB.Table("users").Limit(limit).Offset(offset).Find(&dataUsers)
	totalPage := len(dataUsers)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return dataUsers, totalPage, nil
}

func (Us *userService) LoginUser(username string, password string) (usname string, token string, e error) {
	userData := models.User{}
	config.DB.Where("username = ?", username).First(&userData)
	errLogin := _bcrypt.CompareHashAndPassword([]byte(userData.PasswordHash), []byte(password))

	if errLogin != nil {
		return "", "", fmt.Errorf("error")
	}
	token, errToken := utils.GenerateToken(userData.ID.String(), userData.Role)
	if errToken != nil {
		return "", "", errToken
	}
	return token, userData.Username, nil
}
