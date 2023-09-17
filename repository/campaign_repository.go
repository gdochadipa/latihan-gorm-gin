package repository

import "latihan-api-startup/model/domain"

type CampaignRepository interface {
	FindAll() ([]domain.Campaign, error)
	FindByUserID(userID int) ([]domain.Campaign, error)
	FindByID(ID int) (domain.Campaign, error)
	Save(campaign domain.Campaign) (domain.Campaign, error)
	Update(campaign domain.Campaign) (domain.Campaign, error)
	CreateImage(campaignImage domain.CampaignImage) (domain.CampaignImage, error)
	MarkAllImagesAsNonPrimary(campaignID int) (bool, error)
}
