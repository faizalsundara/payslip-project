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
	adminID := uuid.MustParse("af47a9c9-3cad-4ee0-8c6a-e3f2eadaaccb")

	res, errorPaySlip := Pys.payslipService.CreatePaySlip(adminID, periodID)
	if errorPaySlip != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(errorPaySlip.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseSuccesWithData("Success", res))
	return
}
