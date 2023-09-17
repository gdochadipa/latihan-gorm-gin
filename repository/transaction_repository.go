package repository

import "latihan-api-startup/model/domain"

type TransactionRepository interface {
	GetByCampaignID(campaignID int) ([]domain.Transaction, error)
	GetByUserID(userID int) ([]domain.Transaction, error)
	GetByID(ID int) (domain.Transaction, error)
	Save(transaction domain.Transaction) (domain.Transaction, error)
	Update(transaction domain.Transaction) (domain.Transaction, error)
	FindAll() ([]domain.Transaction, error)
}
