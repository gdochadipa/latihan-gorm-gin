package service

import (
	"errors"
	"latihan-api-startup/model/domain"
	"latihan-api-startup/model/web"
	"latihan-api-startup/repository"
	"strconv"
)

type TransactionServiceImp struct {
	transactionRepository repository.TransactionRepository
	campaignRepository    repository.CampaignRepository
	paymentService        PaymentService
}

func NewTransactionService(repository repository.TransactionRepository, campaign repository.CampaignRepository, payment PaymentService) TransactionService {
	return &TransactionServiceImp{transactionRepository: repository, campaignRepository: campaign, paymentService: payment}
}

func (s *TransactionServiceImp) GetTransactionsByCampaignID(input web.GetCampaignTransactionsInput) ([]domain.Transaction, error) {
	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []domain.Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []domain.Transaction{}, errors.New("Not an owner of the campaign")
	}

	transaction, err := s.transactionRepository.GetByCampaignID(input.ID)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (s *TransactionServiceImp) GetTransactionsByUserID(userID int) ([]domain.Transaction, error) {
	transactions, err := s.transactionRepository.GetByUserID(userID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *TransactionServiceImp) CreateTransaction(input web.CreateTransactionInput) (domain.Transaction, error) {
	transaction := domain.Transaction{}
	transaction.CampaignID = input.CampaignID
	transaction.Amount = input.Amount
	transaction.UserID = input.User.ID
	transaction.Status = "pending"

	newTransaction, err := s.transactionRepository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	paymentTransaction := domain.PaymentTransaction{
		ID:     newTransaction.ID,
		Amount: newTransaction.Amount,
	}

	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentURL
	newTransaction, err = s.transactionRepository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

func (s *TransactionServiceImp) ProcessPayment(input web.TransactionNotificationInput) error {
	transaction_id, _ := strconv.Atoi(input.OrderID)

	transaction, err := s.transactionRepository.GetByID(transaction_id)

	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updateTransaction, err := s.transactionRepository.Update(transaction)
	if err != nil {
		return err
	}

	campaign, err := s.campaignRepository.FindByID(updateTransaction.CampaignID)
	if err != nil {
		return err
	}

	if updateTransaction.Status == "paid" {
		campaign.BackerCount = campaign.BackerCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + updateTransaction.Amount

		_, err := s.campaignRepository.Update(campaign)
		if err != nil {
			return err
		}
	}
	return nil

}

func (s *TransactionServiceImp) GetAllTransactions() ([]domain.Transaction, error) {
	transactions, err := s.transactionRepository.FindAll()
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
