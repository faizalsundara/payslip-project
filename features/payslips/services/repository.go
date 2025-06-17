package services

import (
	"salaries-payslip/models"

	"github.com/google/uuid"
)

type PayslipInterface interface {
	CreatePaySlip(userId uuid.UUID, periodID uuid.UUID) (models.Payslip, error)
}
