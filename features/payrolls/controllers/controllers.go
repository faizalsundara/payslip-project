package controllers

import (
	"net/http"
	"salaries-payslip/features/payrolls/services"
	"salaries-payslip/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type payrollController struct {
	payrollService services.PayrollsInterface
}

func NewPayrollController(pyl services.PayrollsInterface) *payrollController {
	return &payrollController{
		payrollService: pyl,
	}
}

func (pyl *payrollController) PayrollPeriod(c *gin.Context) {
	var body struct {
		IsPaid   bool   `json:"isPaid"`
		PeriodId string `json:"periodID"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed("Invalid input"))
		return
	}
	// adminID := uuid.MustParse(c.GetString("user_id"))
	adminID := uuid.MustParse("af47a9c9-3cad-4ee0-8c6a-e3f2eadaaccb")
	periodID := uuid.MustParse(body.PeriodId)

	_, errPayroll := pyl.payrollService.PayrollPeriod(body.IsPaid, periodID, adminID)
	if errPayroll != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(errPayroll.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseSuccesNoData("Payroll Success"))
	return
}
