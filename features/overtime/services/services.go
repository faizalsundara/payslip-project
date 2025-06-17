package services

import (
	"errors"
	"salaries-payslip/config"
	"salaries-payslip/models"
	"salaries-payslip/utils"
	"time"

	"github.com/google/uuid"
)

type overtimeServices struct{}

func NewOvertimeService() OvertimeInterfaces {
	return &overtimeServices{}
}

func (OS *overtimeServices) SubmitOvertime(userID uuid.UUID, date time.Time, hours float64, ipAddress string) (int, error) {
	var existing models.Overtime
	if err := config.DB.Where("user_id = ? AND date = ?", userID, date).First(&existing).Error; err == nil {
		return 0, errors.New("Overtime already submitted")
	}

	var attendance models.Attendance
	errSelect := config.DB.Where("user_id = ? AND date = ?", userID, date).First(&attendance).Error
	if errSelect != nil {
		return 0, errSelect
	}
	dur, errDuration := utils.ConvertToDuration(attendance.Duration)

	if errDuration != nil {
		return 0, errDuration
	}
	threshold := 8 * time.Hour
	if dur < threshold {
		return 0, errors.New("job duration under 8 hours")
	}
	var overtime = models.Overtime{
		UserID:    userID,
		Date:      date,
		IPAddress: ipAddress,
		Hours:     hours,
		CreatedBy: userID,
	}

	res := config.DB.Create(&overtime)
	if res.Error != nil {
		return 0, errors.New(res.Error.Error())
	}
	return int(res.RowsAffected), nil
}
