package services

import (
	"errors"
	"salaries-payslip/config"
	"salaries-payslip/models"
	"time"

	"github.com/google/uuid"
)

type payrollService struct{}

func NewPayrollService() PayrollsInterface {
	return &payrollService{}
}

func (pyl *payrollService) PayrollPeriod(periodId uuid.UUID, adminID uuid.UUID) (int, error) {
	var attendPeriode models.AttendancePeriod
	if err := config.DB.Where("is_paid = ? AND id = ?", "O", periodId).First(&attendPeriode).Error; err != nil {
		return 0, err
	}

	if attendPeriode.IsPaid == "I" {
		return 0, errors.New("Payroll was done")
	}
	payroll := models.Payroll{
		ID:        uuid.New(),
		StartDate: attendPeriode.StartDate,
		EndDate:   attendPeriode.EndDate,
		PeriodID:  periodId,
		IsPaid:    "I",
		CreatedBy: adminID,
		CreatedAt: time.Now(),
	}
	errCreate := config.DB.Create(&payroll).Error
	if errCreate != nil {
		return 0, errCreate
	}

	res := config.DB.Model(&attendPeriode).Where("is_paid = ? AND id = ?", "O", periodId).Update("is_paid", "I")
	if res.Error != nil {
		return 0, res.Error
	}

	return int(res.RowsAffected), nil
}
