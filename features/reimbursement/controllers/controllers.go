package controllers

import (
	"net/http"
	"salaries-payslip/features/reimbursement/services"
	"salaries-payslip/utils"

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
	}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed("Invalid input"))
		return
	}
	// userID := uuid.MustParse(c.GetString("user_id"))
	userID := uuid.MustParse("af47a9c9-3cad-4ee0-8c6a-e3f2eadaaccb")
	ip := c.ClientIP()
	_, errReimburs := reimburs.reimbursService.SubmitReimburs(body.Amount, userID, body.Description, ip)
	if errReimburs != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(errReimburs.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseSuccesNoData("SubmitReimburs Success"))
	return
}
