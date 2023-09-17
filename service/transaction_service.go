package service

import (
	"latihan-api-startup/model/domain"
	"latihan-api-startup/model/web"
)

type TransactionService interface {
	GetTransactionsByCampaignID(input web.GetCampaignTransactionsInput) ([]domain.Transaction, error)
	GetTransactionsByUserID(userID int) ([]domain.Transaction, error)
	CreateTransaction(input web.CreateTransactionInput) (domain.Transaction, error)
	ProcessPayment(input web.TransactionNotificationInput) error
	GetAllTransactions() ([]domain.Transaction, error)
}
