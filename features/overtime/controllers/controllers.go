package controllers

import (
	"net/http"
	"salaries-payslip/features/overtime/services"
	"salaries-payslip/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type overtimeController struct {
	overtimeServices services.OvertimeInterfaces
}

func NewOvertimeController(ovtSvc services.OvertimeInterfaces) *overtimeController {
	return &overtimeController{
		overtimeServices: ovtSvc,
	}
}

func (ovt *overtimeController) SubmitOvertime(c *gin.Context) {
	var body struct {
		Date  string  `json:"date"`
		Hours float64 `json:"hours"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed("Invalid input"))
		return
	}

	date, err := time.Parse("2006-01-02", body.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed("invalid input date"))
		return
	}

	hours := body.Hours
	if hours > 3 {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed("invalid input hours"))
		return
	}

	// userID := uuid.MustParse(c.GetString("user_id"))
	userID := uuid.MustParse("af47a9c9-3cad-4ee0-8c6a-e3f2eadaaccb")
	ip := c.ClientIP()

	_, errSubmit := ovt.overtimeServices.SubmitOvertime(userID, date, hours, ip)

	if errSubmit != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(errSubmit.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccesNoData("Overtime Submit Success"))
	return
}
