package handler

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"bwastartup/user"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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

	token, err := h.authService.GenerateToken(newUser.ID)

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


	formatter := user.FormatUser(newUser, token)

	response := helper.APIResponse(
		"Account has been registered",
		http.StatusOK,
		"Succes",
		formatter,
	)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) Login(c *gin.Context) {

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

	token, err := h.authService.GenerateToken(loggedInUser.ID)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse(
			"Login Failed",
			http.StatusBadRequest,
			"Failed",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedInUser, token)

	response := helper.APIResponse(
		"Login is successful",
		http.StatusOK,
		"Succes",
		formatter,
	)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailibility(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(
			"Email Validation Failed",
			http.StatusBadRequest,
			"Failed",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)

	if err != nil {
		errorMessage := gin.H{"errors": "Server Error"}

		response := helper.APIResponse(
			"Email Validation Failed",
			http.StatusUnprocessableEntity,
			"Failed",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	responseMessage := "Email has been registered"

	if isEmailAvailable {
		responseMessage = "Email is available"
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	response := helper.APIResponse(
		responseMessage,
		http.StatusOK,
		"Success",
		data,
	)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")

	responseMessage := "Failed to upload avatar"
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse(responseMessage,http.StatusBadRequest,"Failed",data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID


	path := fmt.Sprintf("images/%d-%s",userID, file.Filename)
	err = c.SaveUploadedFile(file,path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse(responseMessage,http.StatusBadRequest,"Failed",data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID,path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse(responseMessage,http.StatusBadRequest,"Failed",data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	responseMessage = "Succes upload avatar"
	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse(responseMessage,http.StatusOK,"Success",data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) FetchUser(c *gin.Context) {
	currentUser:= c.MustGet("currentUser").(user.User)

	formatter := user.FormatUser(currentUser,"")

	response := helper.APIResponse("Succesfuly fetch user data", http.StatusOK,"success",formatter)

	c.JSON(http.StatusOK,response)
}
