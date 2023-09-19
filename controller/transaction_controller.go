package controller

import (
	"latihan-api-startup/helper"
	"latihan-api-startup/model/domain"
	"latihan-api-startup/model/web"
	"latihan-api-startup/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionControllerImp struct {
	service service.TransactionService
}

func NewTransactionHandler(service service.TransactionService) TransactionController {
	return &TransactionControllerImp{service: service}
}

type TransactionController interface {
	GetCampaignTransactions(context *gin.Context)
	GetUserTransactions(context *gin.Context)
	CreateTransaction(context *gin.Context)
	GetNotification(context *gin.Context)
}

func (s *TransactionControllerImp) GetCampaignTransactions(context *gin.Context) {
	var input web.GetCampaignTransactionsInput

	err := context.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := context.MustGet("currentUser").(domain.User)

	input.User = currentUser

	transactions, err := s.service.GetTransactionsByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Campaign's transactions", http.StatusOK, "success", web.FormatCampaignTransactions(transactions))
	context.JSON(http.StatusOK, response)
}

func (s *TransactionControllerImp) GetUserTransactions(context *gin.Context) {
	currentUser := context.MustGet("currentUser").(domain.User)
	userID := currentUser.ID

	transactions, err := s.service.GetTransactionsByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get user's transactions", http.StatusBadRequest, "error", nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("User's transactions", http.StatusOK, "success", web.FormatUserTransactions(transactions))
	context.JSON(http.StatusOK, response)
}

func (s *TransactionControllerImp) CreateTransaction(context *gin.Context) {
	var input web.CreateTransactionInput

	err := context.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create transaction", http.StatusUnprocessableEntity, "error", errorMessage)
		context.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	currentUser := context.MustGet("currentUser").(domain.User)

	input.User = currentUser

	newTransaction, err := s.service.CreateTransaction(input)

	if err != nil {
		response := helper.APIResponse("Failed to create transaction", http.StatusBadRequest, "error", nil)
		context.JSON(http.StatusBadRequest, response)

		return
	}

	response := helper.APIResponse("Success to create transaction", http.StatusOK, "success", web.FormatTransaction(newTransaction))
	context.JSON(http.StatusOK, response)
}

func (s *TransactionControllerImp) GetNotification(context *gin.Context) {
	var input web.TransactionNotificationInput

	err := context.ShouldBind(&input)
	if err != nil {
		response := helper.APIResponse("Failed to process notification", http.StatusBadRequest, "error", nil)
		context.JSON(http.StatusBadRequest, response)

		return
	}

	err = s.service.ProcessPayment(input)
	if err != nil {
		response := helper.APIResponse("Failed to process notification", http.StatusBadRequest, "error", nil)
		context.JSON(http.StatusBadRequest, response)

		return
	}

	context.JSON(http.StatusOK, input)
}
