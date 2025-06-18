package services

import "salaries-payslip/models"

type UserInterface interface {
	GetAllUsers(limit, offset int) ([]models.User, int, error)
	LoginUser(username string, password string) (usname string, token string, e error)
}
