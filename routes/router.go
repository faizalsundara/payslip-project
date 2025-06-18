package routes

import (
	_attendanceCtl "salaries-payslip/features/attendances/controllers"
	_oveetimeCtl "salaries-payslip/features/overtime/controllers"
	_payrollCtl "salaries-payslip/features/payrolls/controllers"
	_payslipCtl "salaries-payslip/features/payslips/controllers"
	_reimbursCtl "salaries-payslip/features/reimbursement/controllers"
	_summaryCtl "salaries-payslip/features/summarySalaries/controllers"
	_userCtl "salaries-payslip/features/users/controllers"
	"salaries-payslip/middlewares"

	_attendanceSvc "salaries-payslip/features/attendances/services"
	_overtimeSvc "salaries-payslip/features/overtime/services"
	_payrollSvc "salaries-payslip/features/payrolls/services"
	_payslipSvc "salaries-payslip/features/payslips/services"
	_reimbursSvc "salaries-payslip/features/reimbursement/services"
	_summarySvc "salaries-payslip/features/summarySalaries/services"
	_userSvc "salaries-payslip/features/users/services"

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
	userSvc := _userSvc.NewUserService()
	userCtl := _userCtl.NewUserController(userSvc)

	employee := r.Group("/employee")
	admin := r.Group("/admin")
	// employee.Use(middlewares.AuthMiddleware("employee"))

	// Attendance
	employee.POST("login-time", middlewares.AuthMiddleware("employee"), attendanceController.AttendanceLoginTime)
	employee.POST("logout-time", middlewares.AuthMiddleware("employee"), attendanceController.AttendanceLogoutTime)
	admin.POST("attendance-period", middlewares.AuthMiddleware("admin"), attendanceController.AttendancePeriode)
	// Overtime
	employee.POST("overtime", middlewares.AuthMiddleware("employee"), overtimeController.SubmitOvertime)
	// Reimburse
	employee.POST("reimbursement", middlewares.AuthMiddleware("employee"), reimbursController.SubmitReimburs)
	//Payroll
	admin.POST("payroll", middlewares.AuthMiddleware("admin"), payrollCtl.PayrollPeriod)
	//Payslip
	employee.POST("payslip", middlewares.AuthMiddleware("employee"), payslipCtl.CreatePayslip)
	//SummarySalaries
	admin.GET("summary-salaries", middlewares.AuthMiddleware("admin"), summaryCtl.SummaryTHP)
	//user
	r.GET("users", userCtl.GetAllUser)
	r.POST("users/login", userCtl.LoginUser)
}
