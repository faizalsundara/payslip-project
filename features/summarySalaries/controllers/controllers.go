package controllers

import (
	"net/http"
	"salaries-payslip/features/summarySalaries/services"
	"salaries-payslip/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SummarySalariesController struct {
	SummarySalaries services.SummarySalariesInterface
}

func NewSummarySalariesController(SS services.SummarySalariesInterface) *SummarySalariesController {
	return &SummarySalariesController{
		SummarySalaries: SS,
	}
}

func (SS *SummarySalariesController) SummaryTHP(c *gin.Context) {
	var body struct {
		PeriodID string `json:"periodID"`
		UserID   string `json:"userID"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed("Invalid input"))
		return
	}

	userID := uuid.MustParse(body.UserID)
	// userID := uuid.MustParse("6b4c795f-4bc5-4b28-96a8-0b245ac69e23")
	periodID := uuid.MustParse(body.PeriodID)
	res, errSvc := SS.SummarySalaries.SummaryTHP(userID, periodID)
	if errSvc != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(errSvc.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseSuccesWithData("success get summary", res))
	return
}
