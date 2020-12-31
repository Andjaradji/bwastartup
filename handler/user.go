package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(
			"Registration Failed",
			http.StatusBadRequest,
			"Failed",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse(
			"Registration Failed",
			http.StatusBadRequest,
			"Failed",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "tokentokentokentoken")

	response := helper.APIResponse(
		"Accoun has been registered",
		http.StatusOK,
		"Succes",
		formatter,
	)

	c.JSON(http.StatusOK,response)

}

func (h *userHandler) Login (c *gin.Context) {
	
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(
			"Login Failed",
			http.StatusBadRequest,
			"Failed",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loggedInUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse(
			"Registration Failed",
			http.StatusBadRequest,
			"Failed",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedInUser, "tokentokentokentoken")

	response := helper.APIResponse(
		"Login is successful",
		http.StatusOK,
		"Succes",
		formatter,
	)

	c.JSON(http.StatusOK,response)
}
