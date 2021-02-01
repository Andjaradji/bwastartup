package transaction

import (
	"bwastartup/campaign"
	"errors"
)

type service struct {
	repository Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetTransactionsByCampaignID (input GetCampaignTransactionInput) ([]Transaction, error)
	GetTransactionsByUserID (userID int)([]Transaction, error)
}

func NewService (repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionsByCampaignID (input GetCampaignTransactionInput) ([]Transaction, error) {

	thisCampaign,err:= s.campaignRepository.FindByID(input.ID)

	if err != nil{
		return []Transaction{}, err
	}

	if thisCampaign.User.ID != input.User.ID {
		return []Transaction{}, errors.New("Not an owner of the campaign")
	}

	transaction,err := s.repository.GetByCampaignID(input.ID)

	if err != nil {
		return transaction,err
	}

	return transaction,nil
}

func (s *service) 	GetTransactionsByUserID (userID int)([]Transaction, error){
	transaction, err := s.repository.GetByUserID(userID)
	
	if err != nil {
		return transaction,err
	}

	return transaction,nil
}
