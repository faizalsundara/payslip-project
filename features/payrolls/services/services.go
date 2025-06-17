package services

import (
	"errors"
	"salaries-payslip/config"
	"salaries-payslip/models"

	"github.com/google/uuid"
)

type payrollService struct{}

func NewPayrollService() PayrollsInterface {
	return &payrollService{}
}

func (pyl *payrollService) PayrollPeriod(IsPaid bool, periodId uuid.UUID, adminID uuid.UUID) (int, error) {
	var attendPeriode models.AttendancePeriod
	if err := config.DB.Where("is_paid = ? AND period_id = ?", false, periodId).First(&attendPeriode).Error; err != nil {
		return 0, err
	}

	if attendPeriode.IsPaid {
		return 0, errors.New("Payroll was done")
	}
	payroll := models.Payroll{
		PeriodID:  periodId,
		IsPaid:    IsPaid,
		CreatedBy: adminID,
	}
	res := config.DB.Create(&payroll)
	if res.Error != nil {
		return 0, res.Error
	}
	return int(res.RowsAffected), nil
}
