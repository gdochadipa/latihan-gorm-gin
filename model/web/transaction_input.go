package web

import "latihan-api-startup/model/domain"

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User domain.User
}

type CreateTransactionInput struct {
	Amount     int `json:"amount" binding:"required"`
	CampaignID int `json:"campaign_id" binding:"required"`
	User       domain.User
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
