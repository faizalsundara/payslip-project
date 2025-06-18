package services

import (
	"salaries-payslip/config"
	"salaries-payslip/models"
	"salaries-payslip/utils"

	"github.com/google/uuid"
)

type SummarySalariesService struct{}

func NewSummaryService() SummarySalariesInterface {
	return &SummarySalariesService{}
}

func (SS *SummarySalariesService) SummaryTHP(userID uuid.UUID, periodID uuid.UUID) (models.SummarySalariesRes, error) {
	var SummaryPayslip models.Payslip

	type SummaryPayslipWithUser struct {
		models.Payslip
		Username string `json:"username"`
	}

	var result SummaryPayslipWithUser

	var SumTHP float64
	err := config.DB.
		Table("payslips").
		Select("payslips.*, users.username").
		Joins("left join users on users.id = payslips.user_id").
		Where("payslips.user_id = ?", userID).
		First(&result).Error

	if err != nil {
		return models.SummarySalariesRes{}, err
	}

	if err := config.DB.Model(&SummaryPayslip).Select("SUM(take_home_pay)").Where("payroll_id = ?", result.PayrollID).Scan(&SumTHP).Error; err != nil {
		return models.SummarySalariesRes{}, err
	}

	payTHP := utils.ConvertFloatToString(result.Payslip.TakeHomePay)
	payAllTHP := utils.ConvertFloatToString(SumTHP)

	var Result = models.SummarySalariesRes{
		UserID:                 result.Payslip.UserID.String(),
		Username:               result.Username,
		TakeHomePay:            payTHP + "IDR",
		TakeHomePayAllEmployee: payAllTHP + "IDR",
	}

	return Result, nil
}
