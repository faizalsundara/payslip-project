package models

import (
	"time"

	"github.com/google/uuid"
	// "gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	Username     string
	PasswordHash string `gorm:"not null"`
	Role         string `gorm:"type:varchar(10);not null"`
	Salary       float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type AttendancePeriod struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	StartDate time.Time `gorm:"type:date"`
	EndDate   time.Time `gorm:"type:date"`
	IsPaid    string
	IPAddress string
	CreatedBy uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Attendance struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID
	Date      time.Time `gorm:"type:date"`
	StartHour string
	EndHour   string
	Duration  string
	CreatedAt time.Time
	CreatedBy uuid.UUID
	IPAddress string
}

type Overtime struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID
	Date      time.Time `gorm:"uniqueIndex:idx_user_date;type:date"`
	Hours     float64
	CreatedAt time.Time
	CreatedBy uuid.UUID
	IPAddress string
}

type Reimbursement struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID      uuid.UUID
	Date        time.Time `gorm:"type:date"`
	Amount      float64
	Description string
	CreatedBy   uuid.UUID
	CreatedAt   time.Time
	IPAddress   string
}

type Payroll struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	PeriodID  uuid.UUID
	IsPaid    string
	CreatedBy uuid.UUID
	CreatedAt time.Time
	StartDate time.Time `gorm:"type:date"`
	EndDate   time.Time `gorm:"type:date"`
}

type Payslip struct {
	ID                  uuid.UUID `gorm:"type:uuid;primaryKey"`
	PayrollID           uuid.UUID
	UserID              uuid.UUID
	StartDate           time.Time `gorm:"type:date"`
	EndDate             time.Time `gorm:"type:date"`
	BaseSalaryDay       float64
	BaseSalaryOvertime  float64
	SalaryThisMonth     float64
	SalaryOvertimeTotal float64
	TotalAttendanceDays int64
	TotalOvertimeHours  float64
	TotalReimbursements float64
	TakeHomePay         float64
	CreatedAt           time.Time
}

type AuditLog struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	TableName string
	RecordID  uuid.UUID
	Action    string
	UserID    uuid.UUID
	RequestID uuid.UUID
	IPAddress string
	CreatedAt time.Time
}

// func AutoMigrate(db *gorm.DB) error {
// 	return db.AutoMigrate(
// 		&User{},
// 		&AttendancePeriod{},
// 		&Attendance{},
// 		&Overtime{},
// 		&Reimbursement{},
// 		&Payroll{},
// 		&Payslip{},
// 		&AuditLog{},
// 	)
// }

type SummarySalariesRes struct {
	UserID                 string
	Username               string
	TakeHomePay            string
	TakeHomePayAllEmployee string
}

type PayslipRes struct {
	ID                  uuid.UUID
	PayrollID           uuid.UUID
	UserID              uuid.UUID
	StartDate           time.Time
	EndDate             time.Time
	BaseSalaryDay       string
	BaseSalaryOvertime  string
	SalaryThisMonth     string
	SalaryOvertimeTotal string
	TotalAttendanceDays int64
	TotalOvertimeHours  float64
	TotalReimbursements float64
	TakeHomePay         string
	CreatedAt           time.Time
}
