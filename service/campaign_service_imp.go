package service

import (
	"latihan-api-startup/model/domain"
	"latihan-api-startup/model/web"
	"latihan-api-startup/repository"
)

type CampaignServiceImpl struct {
	campaignRepository repository.CampaignRepository
}

func NewCampaignService(repository repository.CampaignRepository) CampaignService {
	return &CampaignServiceImpl{campaignRepository: repository}
}

func (s CampaignServiceImpl) GetCampaigns(userID int) ([]domain.Campaign, error) {
	panic("not implemented") // TODO: Implement
}

func (s CampaignServiceImpl) GetCampaignByID(input web.GetCampaignDetailInput) (domain.Campaign, error) {
	panic("not implemented") // TODO: Implement
}

func (s CampaignServiceImpl) CreateCampaign(input web.CreateCampaignInput) (domain.Campaign, error) {
	panic("not implemented") // TODO: Implement
}

func (s CampaignServiceImpl) UpdateCampaign(inputID web.GetCampaignDetailInput, inputData web.CreateCampaignInput) (domain.Campaign, error) {
	panic("not implemented") // TODO: Implement
}

func (s CampaignServiceImpl) SaveCampaignImage(input web.CreateCampaignImageInput, fileLocation string) (domain.CampaignImage, error) {
	panic("not implemented") // TODO: Implement
}
