package services

import (
	"salaries-payslip/config"
	"salaries-payslip/models"
	"time"

	"github.com/google/uuid"
)

type reimbursementService struct{}

func NewReimbersementService() ReimbursementInterfaces {
	return &reimbursementService{}
}

func (reimburs *reimbursementService) SubmitReimburs(amount float64, userID uuid.UUID, description string, ip string, date time.Time) (int, error) {
	var reimbursInsert = models.Reimbursement{
		ID:          uuid.New(),
		Date:        date,
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
