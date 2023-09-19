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

func (s *CampaignRepositoryImp) FindAll() ([]domain.Campaign, error) {
	var campaigns []domain.Campaign
	err := s.DB.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *CampaignRepositoryImp) FindByUserID(userID int) ([]domain.Campaign, error) {
	var campaigns []domain.Campaign

	err := s.DB.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *CampaignRepositoryImp) FindByID(ID int) (domain.Campaign, error) {
	var campaign domain.Campaign

	err := s.DB.Preload("User").Preload("CampaignImages").Where("id = ?", ID).Find(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *CampaignRepositoryImp) Save(campaign domain.Campaign) (domain.Campaign, error) {
	err := s.DB.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *CampaignRepositoryImp) Update(campaign domain.Campaign) (domain.Campaign, error) {
	err := s.DB.Save(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *CampaignRepositoryImp) CreateImage(campaignImage domain.CampaignImage) (domain.CampaignImage, error) {
	err := s.DB.Create(&campaignImage).Error
	if err != nil {
		return campaignImage, err
	}

	return campaignImage, nil
}

func (s *CampaignRepositoryImp) MarkAllImagesAsNonPrimary(campaignID int) (bool, error) {
	err := s.DB.Model(&domain.CampaignImage{}).Where("campaign_id = ?", campaignID).Update("is_primary", false).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
