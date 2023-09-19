package main

import (
	"latihan-api-startup/app"
	"latihan-api-startup/controller"
	"latihan-api-startup/repository"
	"latihan-api-startup/service"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()

	validate := validator.New()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, db, validate)
	authService := service.NewAuthService()
	userController := controller.NewUserController(userService, authService)

	campaignRepository := repository.NewCampaignRepository(db)
	campaignService := service.NewCampaignService(campaignRepository)
	campaignController := controller.NewCampaignController(campaignService)

	paymentService := service.NewPaymentService()

	transRepository := repository.NewTransactionRepository(db)
	transService := service.NewTransactionService(transRepository, campaignRepository, paymentService)
	transController := controller.NewTransactionHandler(transService)

	router := app.NewRouter(userController, campaignController, transController, userService, authService)

	router.Run()

}
