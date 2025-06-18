package services

import (
	"time"

	"github.com/google/uuid"
)

type ReimbursementInterfaces interface {
	SubmitReimburs(amount float64, userID uuid.UUID, description string, ip string, date time.Time) (int, error)
}
