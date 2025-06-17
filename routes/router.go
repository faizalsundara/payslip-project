package routes

import (
	_attendanceCtl "salaries-payslip/features/attendances/controllers"
	_oveetimeCtl "salaries-payslip/features/overtime/controllers"
	_payrollCtl "salaries-payslip/features/payrolls/controllers"
	_payslipCtl "salaries-payslip/features/payslips/controllers"
	_reimbursCtl "salaries-payslip/features/reimbursement/controllers"
	_summaryCtl "salaries-payslip/features/summarySalaries/controllers"

	// "salaries-payslip/middlewares"
	_attendanceSvc "salaries-payslip/features/attendances/services"
	_overtimeSvc "salaries-payslip/features/overtime/services"
	_payrollSvc "salaries-payslip/features/payrolls/services"
	_payslipSvc "salaries-payslip/features/payslips/services"
	_reimbursSvc "salaries-payslip/features/reimbursement/services"
	_summarySvc "salaries-payslip/features/summarySalaries/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	attendanceService := _attendanceSvc.NewAttendanceService()
	attendanceController := _attendanceCtl.NewAttendanceController(attendanceService)
	overtimeService := _overtimeSvc.NewOvertimeService()
	overtimeController := _oveetimeCtl.NewOvertimeController(overtimeService)
	reimbursService := _reimbursSvc.NewReimbersementService()
	reimbursController := _reimbursCtl.NewReimbursementControllers(reimbursService)
	payrollSvc := _payrollSvc.NewPayrollService()
	payrollCtl := _payrollCtl.NewPayrollController(payrollSvc)
	payslipSvc := _payslipSvc.NewPayslipService()
	payslipCtl := _payslipCtl.NewPayslipController(payslipSvc)
	summarySvc := _summarySvc.NewSummaryService()
	summaryCtl := _summaryCtl.NewSummarySalariesController(summarySvc)

	employee := r.Group("/employee")
	admin := r.Group("/admin")
	// employee.Use(middlewares.AuthMiddleware("employee"))

	// Attendance
	employee.POST("login-time", attendanceController.AttendanceLoginTime)
	employee.POST("logout-time", attendanceController.AttendanceLogoutTime)
	admin.POST("attendance-period", attendanceController.AttendancePeriode)
	// Overtime
	employee.POST("overtime", overtimeController.SubmitOvertime)
	// Reimburse
	employee.POST("reimbursement", reimbursController.SubmitReimburs)
	//Payroll
	admin.POST("payroll", payrollCtl.PayrollPeriod)
	//Payslip
	employee.POST("payslip", payslipCtl.CreatePayslip)
	//SummarySalaries
	admin.GET("summary-salaries", summaryCtl.SummaryTHP)
}
