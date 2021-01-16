package campaign

import (
	"strings"
)

type CampaignFormater struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Slug             string `json:"slug"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

func FormatCampaign(campaign Campaign) CampaignFormater {
	campaignFormatter := CampaignFormater{}

	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.Slug = campaign.Slug
	campaignFormatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		for _, campaignImage := range campaign.CampaignImages {
			if campaignImage.IsPrimary == 1 {
				campaignFormatter.ImageURL = campaignImage.FileName
				break
			}
			campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
		}
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormater {

	campaignFormaters := []CampaignFormater{}

	for _, c := range campaigns {
		campaignFormater := FormatCampaign(c)
		campaignFormaters = append(campaignFormaters, campaignFormater)
	}

	return campaignFormaters
}

type CampaignDetailFormater struct {
	ID               int                      `json:"id"`
	UserID           int                      `json:"user_id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_description"`
	Description      string                   `json:"description"`
	Slug             string                   `json:"slug"`
	ImageURL         string                   `json:"image_url"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	Perks            []string                 `json:"perks"`
	Images           []CampaignImageFormatter `json:"images"`
	User             CampaignUserFormatter    `json:"user"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormater {

	campaignDetailFormater := CampaignDetailFormater{}

	campaignDetailFormater.ID = campaign.ID
	campaignDetailFormater.UserID = campaign.UserID
	campaignDetailFormater.Name = campaign.Name
	campaignDetailFormater.ShortDescription = campaign.ShortDescription
	campaignDetailFormater.Description = campaign.Description
	campaignDetailFormater.GoalAmount = campaign.GoalAmount
	campaignDetailFormater.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormater.Slug = campaign.Slug
	campaignDetailFormater.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		for _, campaignImage := range campaign.CampaignImages {
			if campaignImage.IsPrimary == 1 {
				campaignDetailFormater.ImageURL = campaignImage.FileName
				break
			}
			campaignDetailFormater.ImageURL = campaign.CampaignImages[0].FileName
		}
	}

	user := campaign.User
	campaignUserFormatter := CampaignUserFormatter{}
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.ImageURL = user.AvatarFileName

	campaignDetailFormater.User = campaignUserFormatter

	campaignImagesFormatters := []CampaignImageFormatter{}

	if len(campaign.CampaignImages) > 0 {
		for _, campaignImage := range campaign.CampaignImages {
			campaignImageFormatter := CampaignImageFormatter{}
			campaignImageFormatter.ImageURL = campaignImage.FileName
			isPrimary := false
			if campaignImage.IsPrimary == 1 {
				isPrimary = true
			}
			campaignImageFormatter.IsPrimary = isPrimary

			campaignImagesFormatters = append(campaignImagesFormatters, campaignImageFormatter)
		}
	}

	campaignDetailFormater.Images = campaignImagesFormatters

	campaignDetailFormater.Perks = strings.Split(campaign.Perks, ", ")

	return campaignDetailFormater
}
