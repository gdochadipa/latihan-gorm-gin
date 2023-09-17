package repository

import (
	"latihan-api-startup/model/domain"

	"gorm.io/gorm"
)

type TransactionRepositoryImp struct {
	DB *gorm.DB
}

func NewTransactionRepository(DB *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImp{DB: DB}
}

func (s *TransactionRepositoryImp) GetByCampaignID(campaignID int) ([]domain.Transaction, error) {
	var transactions []domain.Transaction

	err := s.DB.Preload("User").Where("campaign_id = ?", campaignID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *TransactionRepositoryImp) GetByUserID(userID int) ([]domain.Transaction, error) {
	var transactions []domain.Transaction

	err := s.DB.Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id = ?", userID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s TransactionRepositoryImp) GetByID(ID int) (domain.Transaction, error) {
	var transaction domain.Transaction

	err := s.DB.Where("id = ?", ID).Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (s *TransactionRepositoryImp) Save(transaction domain.Transaction) (domain.Transaction, error) {
	err := s.DB.Create(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (s *TransactionRepositoryImp) Update(transaction domain.Transaction) (domain.Transaction, error) {
	err := s.DB.Save(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (s *TransactionRepositoryImp) FindAll() ([]domain.Transaction, error) {
	var transactions []domain.Transaction

	err := s.DB.Preload("Campaign").Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
