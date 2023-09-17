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

	router := app.NewRouter(userController)

	router.Run()

}
