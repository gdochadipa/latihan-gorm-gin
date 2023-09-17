package service

import (
	"latihan-api-startup/model/domain"
	"latihan-api-startup/model/web"
)

type CampaignService interface {
	GetCampaigns(userID int) ([]domain.Campaign, error)
	GetCampaignByID(input web.GetCampaignDetailInput) (domain.Campaign, error)
	CreateCampaign(input web.CreateCampaignInput) (domain.Campaign, error)
	UpdateCampaign(inputID web.GetCampaignDetailInput, inputData web.CreateCampaignInput) (domain.Campaign, error)
	SaveCampaignImage(input web.CreateCampaignImageInput, fileLocation string) (domain.CampaignImage, error)
}
