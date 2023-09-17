package service

import (
	"latihan-api-startup/model/domain"
	"latihan-api-startup/model/web"
	"latihan-api-startup/repository"
)

type TransactionServiceImp struct {
	repository repository.TransactionRepository
}

func NewTransactionService(repository repository.TransactionRepository) TransactionService {
	return &TransactionServiceImp{repository: repository}
}

func (s TransactionServiceImp) GetTransactionsByCampaignID(input web.GetCampaignTransactionsInput) ([]domain.Transaction, error) {
	panic("not implemented") // TODO: Implement
}

func (s TransactionServiceImp) GetTransactionsByUserID(userID int) ([]domain.Transaction, error) {
	panic("not implemented") // TODO: Implement
}

func (s TransactionServiceImp) CreateTransaction(input web.CreateTransactionInput) (domain.Transaction, error) {
	panic("not implemented") // TODO: Implement
}

func (s TransactionServiceImp) ProcessPayment(input web.TransactionNotificationInput) error {
	panic("not implemented") // TODO: Implement
}

func (s TransactionServiceImp) GetAllTransactions() ([]domain.Transaction, error) {
	panic("not implemented") // TODO: Implement
}
