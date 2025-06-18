package controllers

import (
	"net/http"
	"salaries-payslip/features/reimbursement/services"
	"salaries-payslip/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type reimbursControllers struct {
	reimbursService services.ReimbursementInterfaces
}

func NewReimbursementControllers(reimburs services.ReimbursementInterfaces) *reimbursControllers {
	return &reimbursControllers{
		reimbursService: reimburs,
	}
}

func (reimburs *reimbursControllers) SubmitReimburs(c *gin.Context) {
	var body struct {
		Amount      float64 `json:"amount"`
		Description string  `json:"description"`
		Date        string  `json:"date"`
	}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed("invalid input"))
		return
	}
	userID := uuid.MustParse(c.GetString("user_id"))
	// userID := uuid.MustParse("af47a9c9-3cad-4ee0-8c6a-e3f2eadaaccb")
	date, errStart := time.Parse("2006-01-02", body.Date)
	if errStart != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed("invalid date format"))
		return
	}
	// userID := uuid.MustParse("6b4c795f-4bc5-4b28-96a8-0b245ac69e23")
	ip := c.ClientIP()
	_, errReimburs := reimburs.reimbursService.SubmitReimburs(body.Amount, userID, body.Description, ip, date)
	if errReimburs != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(errReimburs.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseSuccesNoData("submit reimburs success"))
	return
}
