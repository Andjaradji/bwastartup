package campaign

import "time"

type Campaign struct {
	ID               int
	userID           int
	Name             string
	ShortDescription string
	GoalAmount       int
	CurrentAmount    int
	Description      string
	Perks            string
	BackerCount      int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImage
}

type CampaignImage struct {
	ID         int
	CampaignID int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
