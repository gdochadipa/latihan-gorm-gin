package repository

import (
	"latihan-api-startup/model/domain"
)

type UserRepository interface {
	Save(user domain.User) domain.User
	FindByEmail(email string) domain.User
	FindByID(ID int) domain.User
	Update(user domain.User) (domain.User, error)
	FindAll() []domain.User
}
