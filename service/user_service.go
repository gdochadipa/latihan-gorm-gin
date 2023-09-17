package service

import (
	"latihan-api-startup/model/domain"
	"latihan-api-startup/model/web"
)

type UserService interface {
	RegisterUser(input web.RegisterUserInput) (domain.User, error)
	Login(input web.LoginInput) (domain.User, error)
	IsEmailAvailable(input web.CheckEmailInput) bool
	SaveAvatar(ID int, fileLocation string) (domain.User, error)
	GetUserByID(ID int) (domain.User, error)
	GetAllUsers() ([]domain.User, error)
	UpdateUser(input web.FormUpdateUserInput) (domain.User, error)
}
