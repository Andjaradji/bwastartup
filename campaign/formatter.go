package campaign

type CampaignFormater struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

func formatCampaign(campaign Campaign) CampaignFormater {
	campaignFormatter := CampaignFormater{}

	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
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

func FormatCampaigns (campaigns []Campaign) []CampaignFormater {

	campaignFormaters := []CampaignFormater{}

	for _, c := range campaigns {
		campaignFormater := formatCampaign(c)
		campaignFormaters = append(campaignFormaters, campaignFormater)
	}

	return campaignFormaters
}
