package repository

import (
	"latihan-api-startup/model/domain"

	"gorm.io/gorm"
)

type CampaignRepositoryImp struct {
	DB *gorm.DB
}

func NewCampaignRepository(DB *gorm.DB) CampaignRepository {
	return &CampaignRepositoryImp{DB: DB}
}

func (s CampaignRepositoryImp) FindAll() ([]domain.Campaign, error) {
	panic("not implemented") // TODO: Implement
}

func (s CampaignRepositoryImp) FindByUserID(userID int) ([]domain.Campaign, error) {
	panic("not implemented") // TODO: Implement
}

func (s CampaignRepositoryImp) FindByID(ID int) (domain.Campaign, error) {
	panic("not implemented") // TODO: Implement
}

func (s CampaignRepositoryImp) Save(campaign domain.Campaign) (domain.Campaign, error) {
	panic("not implemented") // TODO: Implement
}

func (s CampaignRepositoryImp) Update(campaign domain.Campaign) (domain.Campaign, error) {
	panic("not implemented") // TODO: Implement
}

func (s CampaignRepositoryImp) CreateImage(campaignImage domain.CampaignImage) (domain.CampaignImage, error) {
	panic("not implemented") // TODO: Implement
}

func (s CampaignRepositoryImp) MarkAllImagesAsNonPrimary(campaignID int) (bool, error) {
	panic("not implemented") // TODO: Implement
}
