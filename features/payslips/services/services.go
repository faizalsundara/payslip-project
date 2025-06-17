package services

import (
	"salaries-payslip/config"
	"salaries-payslip/models"

	"github.com/google/uuid"
)

type PayslipService struct{}

func NewPayslipService() PayslipInterface {
	return &PayslipService{}
}

func (Ps *PayslipService) CreatePaySlip(userId uuid.UUID, periodID uuid.UUID) (models.Payslip, error) {
	var attend models.Attendance
	var user models.User
	var ovTime models.Overtime
	var reimburs models.Reimbursement
	var payroll models.Payroll
	var countAttend int64
	var totalOvertimeHours float64
	var totalReimburs float64

	if errUser := config.DB.Where("userd_id = ?", userId).First(&user).Error; errUser != nil {
		return models.Payslip{}, errUser
	}

	if errPayroll := config.DB.Where("userd_id = ? AND ID = ? AND is_paid = ?", userId, periodID, true).First(&payroll).Error; errPayroll != nil {
		return models.Payslip{}, errPayroll
	}

	if err := config.DB.Model(&attend).Where("userd_id = ?", userId).Where("date >= ? AND date <= ?", payroll.StartDate, payroll.EndDate).Count(&countAttend).Error; err != nil {
		return models.Payslip{}, err
	}

	errOvt := config.DB.Model(&ovTime).
		Select("SUM(hours)").
		Where("user_id = ?", userId).
		Where("date >= ? AND date <= ?", payroll.StartDate, payroll.EndDate).
		Scan(&totalOvertimeHours).Error

	if errOvt != nil {
		return models.Payslip{}, errOvt
	}

	errReimburs := config.DB.Model(&reimburs).
		Select("SUM(amount)").
		Where("user_id = ?", userId).
		Where("date >= ? AND date <= ?", payroll.StartDate, payroll.EndDate).
		Scan(&totalReimburs).Error

	if errReimburs != nil {
		return models.Payslip{}, errReimburs
	}

	totalDays := payroll.EndDate.Sub(payroll.StartDate).Hours() / 24
	baseSalary := (user.Salary / totalDays) * float64(countAttend)
	payOvertime := baseSalary / 8
	thp := (float64(countAttend) * baseSalary) + totalReimburs + (2 * payOvertime)

	var InsertSlip = models.Payslip{
		PayrollID:           payroll.ID,
		BaseSalary:          baseSalary,
		UserID:              userId,
		TotalAttendanceDays: countAttend,
		TotalOvertimeHours:  totalOvertimeHours,
		TotalReimbursements: totalReimburs,
		TakeHomePay:         thp,
		StartDate:           payroll.StartDate,
		EndDate:             payroll.EndDate,
	}
	res := config.DB.Create(&InsertSlip)
	if res.Error != nil {
		return models.Payslip{}, res.Error
	}

	return models.Payslip{
		PayrollID:           payroll.ID,
		BaseSalary:          baseSalary,
		UserID:              userId,
		TotalAttendanceDays: countAttend,
		TotalOvertimeHours:  totalOvertimeHours,
		TotalReimbursements: totalReimburs,
		TakeHomePay:         thp,
		StartDate:           payroll.StartDate,
		EndDate:             payroll.EndDate,
	}, nil
}
