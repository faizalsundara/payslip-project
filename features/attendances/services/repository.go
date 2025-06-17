package services

import (
	"time"

	"github.com/google/uuid"
)

type AttendanceInterface interface {
	AttendanceLoginTime(userID uuid.UUID, date time.Time, ip string) (int, error)
	AttendanceLogoutTime(userID uuid.UUID, date time.Time, ip string) (int, error)
	AttendancePeriod(startDate time.Time, endDate time.Time, adminID uuid.UUID, ip string) (int, error)
}
