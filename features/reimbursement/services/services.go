package services

import (
	"salaries-payslip/config"
	"salaries-payslip/models"

	"github.com/google/uuid"
)

type reimbursementService struct{}

func NewReimbersementService() ReimbursementInterfaces {
	return &reimbursementService{}
}

func (reimburs *reimbursementService) SubmitReimburs(amount float64, userID uuid.UUID, description string, ip string) (int, error) {
	var reimbursInsert = models.Reimbursement{
		UserID:      userID,
		Amount:      amount,
		Description: description,
		CreatedBy:   userID,
		IPAddress:   ip,
	}
	res := config.DB.Create(&reimbursInsert)
	if res.Error != nil {
		return 0, res.Error
	}

	return int(res.RowsAffected), nil
}
