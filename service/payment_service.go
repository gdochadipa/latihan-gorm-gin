package service

import (
	"latihan-api-startup/model/domain"
	"strconv"

	"github.com/veritrans/go-midtrans"
)

type PaymentServiceImp struct {
}

type PaymentService interface {
	GetPaymentURL(transaction domain.PaymentTransaction, user domain.User) (string, error)
}

func NewPaymentService() PaymentService {
	return &PaymentServiceImp{}
}

func (s *PaymentServiceImp) GetPaymentURL(transaction domain.PaymentTransaction, user domain.User) (string, error) {
	// Client Key = "SB-Mid-client-xOxsSVG3Cd40HYG-"
	// Server Key = "SB-Mid-server-NDhkQWAGKMo5F2uWQlKIhbXw"
	// ID Merchant = "G324368064"

	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-NDhkQWAGKMo5F2uWQlKIhbXw"
	midclient.ClientKey = "SB-Mid-client-xOxsSVG3Cd40HYG-"
	midclient.APIEnvType = midtrans.Sandbox

	snap := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokRes, err := snap.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokRes.RedirectURL, nil
}
