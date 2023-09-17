package service

import "latihan-api-startup/model/domain"

type PaymentServiceImp struct {
}

type PaymentService interface {
	GetPaymentURL(transaction domain.PaymentTransaction, user domain.User) (string, error)
}

func NewService() PaymentService {
	return &PaymentServiceImp{}
}

func (s *PaymentServiceImp) GetPaymentURL(transaction domain.PaymentTransaction, user domain.User) (string, error) {
	// Client Key = "SB-Mid-client-xOxsSVG3Cd40HYG-"
	// Server Key = "SB-Mid-server-NDhkQWAGKMo5F2uWQlKIhbXw"4
	// ID Merchant = "G324368064"
	return "", nil
}
