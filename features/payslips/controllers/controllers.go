package controllers

import (
	"net/http"
	"salaries-payslip/features/payslips/services"
	"salaries-payslip/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type payslipController struct {
	payslipService services.PayslipInterface
}

func NewPayslipController(pys services.PayslipInterface) *payslipController {
	return &payslipController{
		payslipService: pys,
	}
}

func (Pys *payslipController) CreatePayslip(c *gin.Context) {
	var body struct {
		PeriodID string `json:"periodID"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed("Invalid input"))
		return
	}

	periodID := uuid.MustParse(body.PeriodID)
	userID := uuid.MustParse(c.GetString("user_id"))
	// userID := uuid.MustParse("6b4c795f-4bc5-4b28-96a8-0b245ac69e23")

	res, errorPaySlip := Pys.payslipService.CreatePaySlip(userID, periodID)
	if errorPaySlip != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(errorPaySlip.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseSuccesWithData("generate payslip success", res))
	return
}
