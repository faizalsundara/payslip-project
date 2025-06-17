package services

import "github.com/google/uuid"

type PayrollsInterface interface {
	PayrollPeriod(IsPaid bool, periodId uuid.UUID, adminID uuid.UUID) (int, error)
}
