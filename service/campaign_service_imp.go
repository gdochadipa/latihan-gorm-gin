package service

import (
	"errors"
	"fmt"
	"latihan-api-startup/model/domain"
	"latihan-api-startup/model/web"
	"latihan-api-startup/repository"

	"github.com/gosimple/slug"
)

type CampaignServiceImpl struct {
	campaignRepository repository.CampaignRepository
}

func NewCampaignService(repository repository.CampaignRepository) CampaignService {
	return &CampaignServiceImpl{campaignRepository: repository}
}

func (s *CampaignServiceImpl) GetCampaigns(userID int) ([]domain.Campaign, error) {
	if userID != 0 {
		campaigns, err := s.campaignRepository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.campaignRepository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *CampaignServiceImpl) GetCampaignByID(input web.GetCampaignDetailInput) (domain.Campaign, error) {
	campaign, err := s.campaignRepository.FindByID(input.ID)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *CampaignServiceImpl) CreateCampaign(input web.CreateCampaignInput) (domain.Campaign, error) {
	campaign := domain.Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.UserID = input.User.ID

	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugCandidate)

	newCampaign, err := s.campaignRepository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}

func (s *CampaignServiceImpl) UpdateCampaign(inputID web.GetCampaignDetailInput, inputData web.CreateCampaignInput) (domain.Campaign, error) {
	campaign, err := s.campaignRepository.FindByID(inputID.ID)
	if err != nil {
		return campaign, err
	}

	if campaign.UserID != inputData.User.ID {
		return campaign, errors.New("Not an owner of the campaign")
	}

	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount

	updatedCampaign, err := s.campaignRepository.Update(campaign)
	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil
}

func (s *CampaignServiceImpl) SaveCampaignImage(input web.CreateCampaignImageInput, fileLocation string) (domain.CampaignImage, error) {
	campaign, err := s.campaignRepository.FindByID(input.CampaignID)
	if err != nil {
		return domain.CampaignImage{}, err
	}

	if campaign.UserID != input.User.ID {
		return domain.CampaignImage{}, errors.New("Not an owner of the campaign")
	}

	isPrimary := 0
	if input.IsPrimary {
		isPrimary = 1

		_, err := s.campaignRepository.MarkAllImagesAsNonPrimary(input.CampaignID)
		if err != nil {
			return domain.CampaignImage{}, err
		}
	}

	campaignImage := domain.CampaignImage{}
	campaignImage.CampaignID = input.CampaignID
	campaignImage.IsPrimary = isPrimary
	campaignImage.FileName = fileLocation

	newCampaignImage, err := s.campaignRepository.CreateImage(campaignImage)
	if err != nil {
		return newCampaignImage, err
	}

	return newCampaignImage, nil
}
