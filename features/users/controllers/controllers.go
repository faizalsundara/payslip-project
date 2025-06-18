package controllers

import (
	"net/http"
	"salaries-payslip/features/users/services"
	"salaries-payslip/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService services.UserInterface
}

func NewUserController(us services.UserInterface) *userController {
	return &userController{
		userService: us,
	}
}

func (us *userController) GetAllUser(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")
	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)
	result, totalPage, err := us.userService.GetAllUsers(limitInt, offsetInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed("failed to get all data"))
		return
	}
	var resp = map[string]interface{}{
		"users":      result,
		"total_page": totalPage,
	}

	c.JSON(http.StatusOK, utils.ResponseSuccesWithData("success to get all data", resp))
	return
}

func (us *userController) LoginUser(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	errLog := c.ShouldBindJSON(&body)
	if errLog != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed("invalid input data"))
		return
	}
	token, username, e := us.userService.LoginUser(body.Username, body.Password)
	if e != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed("username or password incorrect"))
		return
	}
	data := map[string]interface{}{
		"username": username,
		"token":    token,
	}
	c.JSON(http.StatusOK, utils.ResponseSuccesWithData("LOGIN SUCCES", data))
	return
}
