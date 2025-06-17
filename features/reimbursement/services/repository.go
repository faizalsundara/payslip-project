package services

import "github.com/google/uuid"

type ReimbursementInterfaces interface {
	SubmitReimburs(amount float64, userID uuid.UUID, description string, ip string) (int, error)
}
