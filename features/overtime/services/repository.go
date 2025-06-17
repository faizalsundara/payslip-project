package services

import (
	"time"

	"github.com/google/uuid"
)

type OvertimeInterfaces interface {
	SubmitOvertime(userID uuid.UUID, date time.Time, hours float64, ipAddress string) (int, error)
}
