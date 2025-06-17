package services

import (
	"salaries-payslip/config"
	"salaries-payslip/models"

	"github.com/google/uuid"
)

type SummarySalariesService struct{}

func NewSummaryService() SummarySalariesInterface {
	return &SummarySalariesService{}
}

func (SS *SummarySalariesService) SummaryTHP(userID uuid.UUID, payrollID uuid.UUID) (models.SummarySalariesRes, error) {
	var SummaryPayslip models.Payslip

	type SummaryPayslipWithUser struct {
		Payslip  models.Payslip
		Username string `json:"username"`
	}

	var result SummaryPayslipWithUser

	var CountTHP int64
	err := config.DB.
		Table("summary_playslips").
		Select("summary_playslips.*, users.username").
		Joins("left join users on users.id = summary_playslips.user_id").
		Where("summary_playslips.user_id = ?", userID).
		First(&result).Error

	if err != nil {
		return models.SummarySalariesRes{}, err
	}

	if err := config.DB.Model(&SummaryPayslip).Where("payroll_id = ?", payrollID).Count(&CountTHP).Error; err != nil {
		return models.SummarySalariesRes{}, err
	}

	var Result = models.SummarySalariesRes{
		UserID:                 result.Payslip.UserID.String(),
		Username:               result.Username,
		TakeHomePay:            result.Payslip.TakeHomePay,
		TakeHomePayAllEmployee: float64(CountTHP),
	}

	return Result, nil
}
