package controllers

import (
	"fmt"
	"net/http"
	"salaries-payslip/features/attendances/services"
	"salaries-payslip/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AttendanceController struct {
	attendanceService services.AttendanceInterface
}

func NewAttendanceController(attSvc services.AttendanceInterface) *AttendanceController {
	return &AttendanceController{
		attendanceService: attSvc,
	}
}

func (ec *AttendanceController) AttendanceLoginTime(c *gin.Context) {

	var body struct {
		Date string `json:"date"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed("invalid input"))
		return
	}

	date, err := time.Parse("2006-01-02", body.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed("invalid input"))
		return
	}

	userID := uuid.MustParse(c.GetString("user_id"))
	fmt.Println("User---ID", userID)
	// userID := uuid.MustParse("af47a9c9-3cad-4ee0-8c6a-e3f2eadaaccb")
	// userID := uuid.MustParse("6b4c795f-4bc5-4b28-96a8-0b245ac69e23")
	ip := c.ClientIP()

	_, errAttend := ec.attendanceService.AttendanceLoginTime(userID, date, ip)

	if errAttend != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(errAttend.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccesNoData("login time Submitted"))
	return
}

func (ec *AttendanceController) AttendanceLogoutTime(c *gin.Context) {
	var body struct {
		Date string `json:"date"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed("invalid input"))
		return
	}

	date, err := time.Parse("2006-01-02", body.Date)
	fmt.Println("DATE---12309", date)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed("invalid input"))
		return
	}

	userID := uuid.MustParse(c.GetString("user_id"))
	// userID := uuid.MustParse("af47a9c9-3cad-4ee0-8c6a-e3f2eadaaccb")
	// userID := uuid.MustParse("6b4c795f-4bc5-4b28-96a8-0b245ac69e23")
	ip := c.ClientIP()

	_, errAttend := ec.attendanceService.AttendanceLogoutTime(userID, date, ip)
	if errAttend != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(errAttend.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccesNoData("logout time submitted"))
	return
}

func (ec *AttendanceController) AttendancePeriode(c *gin.Context) {
	var body struct {
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed("Invalid Input"))
		return
	}

	startDate, errStart := time.Parse("2006-01-02", body.StartDate)
	endDate, errEnd := time.Parse("2006-01-02", body.EndDate)
	if errStart != nil || errEnd != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed("Invalid date format"))
		return
	}

	if endDate.Before(startDate) {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed("End date must be after start date"))
		return
	}

	ip := c.ClientIP()
	adminID := uuid.MustParse(c.GetString("user_id"))
	// adminID := uuid.MustParse("af47a9c9-3cad-4ee0-8c6a-e3f2eadaaccb")
	// adminID := uuid.MustParse("6b4c795f-4bc5-4b28-96a8-0b245ac69e23")
	_, errAttend := ec.attendanceService.AttendancePeriod(startDate, endDate, adminID, ip)
	if errAttend != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(errAttend.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseSuccesNoData("Attendance period created successfully"))
	return
}
