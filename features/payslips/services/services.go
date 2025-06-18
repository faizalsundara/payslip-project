package services

import (
	"fmt"
	"salaries-payslip/config"
	"salaries-payslip/models"
	"salaries-payslip/utils"
	"time"

	"github.com/google/uuid"
)

type PayslipService struct{}

func NewPayslipService() PayslipInterface {
	return &PayslipService{}
}

func (Ps *PayslipService) CreatePaySlip(userId uuid.UUID, periodID uuid.UUID) (models.PayslipRes, error) {
	var attend models.Attendance
	var user models.User
	var ovTime models.Overtime
	var reimburs models.Reimbursement
	var payroll models.Payroll
	var countAttend int64
	var totalOvertimeHours float64
	var totalReimburs float64

	if errUser := config.DB.Where("id = ?", userId).First(&user).Error; errUser != nil {
		return models.PayslipRes{}, errUser
	}

	if errPayroll := config.DB.Where("period_id = ? AND is_paid = ?", periodID, "I").First(&payroll).Error; errPayroll != nil {
		return models.PayslipRes{}, errPayroll
	}

	if err := config.DB.Model(&attend).Where("user_id = ?", userId).Where("date >= ? AND date <= ?", payroll.StartDate, payroll.EndDate).Count(&countAttend).Error; err != nil {
		return models.PayslipRes{}, err
	}

	config.DB.Model(&ovTime).
		Select("SUM(hours)").
		Where("user_id = ?", userId).
		Where("date >= ? AND date <= ?", payroll.StartDate, payroll.EndDate).Scan(&totalOvertimeHours)

	// if errOvt != nil {
	// 	return models.PayslipRes{}, errOvt
	// }

	config.DB.Model(&reimburs).
		Select("SUM(amount)").
		Where("user_id = ?", userId).
		Where("date >= ? AND date <= ?", payroll.StartDate, payroll.EndDate).
		Scan(&totalReimburs)

	// if errReimburs != nil {
	// 	return models.PayslipRes{}, errReimburs
	// }

	totalDays := payroll.EndDate.Sub(payroll.StartDate).Hours() / 24
	salaryPerDay := (user.Salary / totalDays)
	salaryPerHour := (user.Salary / totalDays) / 8
	salaryPerAttend := salaryPerDay * float64(countAttend)
	payOvertime := salaryPerHour * totalOvertimeHours * 2
	fmt.Println("total_days", totalDays)
	// thp := (float64(countAttend) * baseSalary) + totalReimburs + (2 * payOvertime)
	thp := salaryPerAttend + totalReimburs + payOvertime
	idUUID := uuid.New()
	timeCreated := time.Now()
	var InsertSlip = models.Payslip{
		ID:                  idUUID,
		PayrollID:           payroll.ID,
		BaseSalaryDay:       salaryPerDay,
		BaseSalaryOvertime:  salaryPerHour,
		SalaryThisMonth:     salaryPerAttend,
		SalaryOvertimeTotal: payOvertime,
		UserID:              userId,
		TotalAttendanceDays: countAttend,
		TotalOvertimeHours:  totalOvertimeHours,
		TotalReimbursements: totalReimburs,
		TakeHomePay:         thp,
		StartDate:           payroll.StartDate,
		EndDate:             payroll.EndDate,
		CreatedAt:           timeCreated,
	}
	res := config.DB.Create(&InsertSlip)
	if res.Error != nil {
		return models.PayslipRes{}, res.Error
	}

	salaryDay := utils.ConvertFloatToString(salaryPerDay)
	salaryHour := utils.ConvertFloatToString(salaryPerHour)
	salaryAttend := utils.ConvertFloatToString(salaryPerAttend)
	payOvt := utils.ConvertFloatToString(payOvertime)
	payTHP := utils.ConvertFloatToString(thp)

	return models.PayslipRes{
		ID:                  idUUID,
		PayrollID:           payroll.ID,
		BaseSalaryDay:       salaryDay + "IDR",
		BaseSalaryOvertime:  salaryHour + "IDR",
		SalaryThisMonth:     salaryAttend + "IDR",
		SalaryOvertimeTotal: payOvt + "IDR",
		UserID:              userId,
		TotalAttendanceDays: countAttend,
		TotalOvertimeHours:  totalOvertimeHours,
		TotalReimbursements: totalReimburs,
		TakeHomePay:         payTHP + "IDR",
		StartDate:           payroll.StartDate,
		EndDate:             payroll.EndDate,
		CreatedAt:           timeCreated,
	}, nil
}
