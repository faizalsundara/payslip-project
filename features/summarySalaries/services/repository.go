package services

import (
	"salaries-payslip/models"

	"github.com/google/uuid"
)

type SummarySalariesInterface interface {
	SummaryTHP(userID uuid.UUID) (models.SummarySalariesRes, error)
}
