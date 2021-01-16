package handler

import (
	"bwastartup/campaign"
	"bwastartup/user"
	"bwastartup/helper"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)

	if err != nil {
		response := helper.APIResponse(
			"Error to get Campaigns",
			http.StatusBadRequest,
			"error",
			nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignFormaters := campaign.FormatCampaigns(campaigns)

	response := helper.APIResponse(
		"List of Campaigns",
		http.StatusOK,
		"success",
		campaignFormaters,
	)
	c.JSON(http.StatusOK, response)
	return

}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of Campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.service.GetCampaignByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of Campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if campaignDetail.ID == 0 {
		response := helper.APIResponse("No Campaign Found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	campaignDetailFormatter := campaign.FormatCampaignDetail(campaignDetail)

	response := helper.APIResponse("Campaign Detail", http.StatusOK, "success", campaignDetailFormatter)
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) CreateCampaign (c *gin.Context){
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)

	if err != nil{
			errors := helper.FormatError(err)
			errorMessage := gin.H{"errors": errors}
			response := helper.APIResponse("Failed to create campaign",http.StatusUnprocessableEntity,"error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity,response)
			return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)

	if err != nil {
		response := helper.APIResponse("Failed to create campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Succes to create campaign", http.StatusOK, "success", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) UpdateCampaign (c *gin.Context){
	var inputID campaign.GetCampaignDetailInput


	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.APIResponse("Failed to update of Campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}


	var inputData campaign.CreateCampaignInput

	err = c.ShouldBindJSON(&inputData)

	if err != nil{
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update campaign",http.StatusUnprocessableEntity,"error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity,response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	updateCampaign, err := h.service.UpdateCampaign(inputID,inputData)

	if err != nil{
		response := helper.APIResponse("Faile to update campaign", http.StatusBadRequest,"error",nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Succes to update campaign", http.StatusOK, "success", campaign.FormatCampaign(updateCampaign))
	c.JSON(http.StatusOK, response)
}

