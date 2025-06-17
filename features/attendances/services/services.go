package services

import (
	"fmt"
	"time"

	"errors"
	"salaries-payslip/config"
	"salaries-payslip/models"
	"salaries-payslip/utils"

	"github.com/google/uuid"
)

type attendanceService struct{}

func NewAttendanceService() AttendanceInterface {
	return &attendanceService{}
}

func (s *attendanceService) AttendanceLoginTime(userID uuid.UUID, date time.Time, ip string) (int, error) {
	if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		return 0, errors.New("cannot submit on weekends")
	}

	// Check if already submitted
	var existing models.Attendance
	if err := config.DB.Where("user_id = ? AND date = ?", userID, date).First(&existing).Error; err == nil {
		return 0, errors.New("already submitted")
	}

	attendance := models.Attendance{
		ID:        uuid.New(),
		UserID:    userID,
		Date:      date,
		StartHour: time.Now().Format("15:04:05"),
		CreatedBy: userID,
		IPAddress: ip,
	}

	Result := config.DB.Create(&attendance)
	return int(Result.RowsAffected), nil
}
func (s *attendanceService) AttendanceLogoutTime(userID uuid.UUID, date time.Time, ip string) (result int, err error) {
	if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		return 0, errors.New("cannot submit on weekends")
	}

	var Attend = models.Attendance{}
	errSelect := config.DB.Model(&Attend).Where("user_id = ? AND date = ?", userID.String(), date).First(&Attend).Error
	if errSelect != nil {
		return 0, errSelect
	}

	fmt.Print("HASIL---", Attend.StartHour)

	timeStart, errTime := utils.ConvertToTime(Attend.StartHour)
	if errTime != nil {
		return 0, errTime
	}

	var Endhour = time.Now().Format("15:04:05")

	timeEnd, errTimeEnd := utils.ConvertToTime(Endhour)
	if errTimeEnd != nil {
		return 0, errTimeEnd
	}

	duration := timeEnd.Sub(timeStart)

	formatted := fmt.Sprintf("%02d:%02d:%02d",
		int(duration.Hours()),
		int(duration.Minutes())%60,
		int(duration.Seconds())%60,
	)

	res := config.DB.Model(&models.Attendance{}).Where("user_id = ? AND date = ?", userID.String(), date).Updates(models.Attendance{EndHour: Endhour, Duration: formatted})
	if res.Error != nil {
		return 0, res.Error
	}
	return int(res.RowsAffected), nil
}

func (s *attendanceService) AttendancePeriod(startDate time.Time, endDate time.Time, adminID uuid.UUID, ip string) (int, error) {
	var existing models.AttendancePeriod

	if err := config.DB.Where("start_date <= ? AND end_date >= ?", endDate, startDate).First(&existing).Error; err == nil {
		return 0, errors.New("Attendance period overlaps with existing period")
	}

	newPeriod := models.AttendancePeriod{
		ID:        uuid.New(),
		StartDate: startDate,
		EndDate:   endDate,
		CreatedBy: adminID,
		IsPaid:    false,
		IPAddress: ip,
	}

	res := config.DB.Create(&newPeriod)
	if res.Error != nil {
		return 0, res.Error
	}
	return int(res.RowsAffected), nil
}
