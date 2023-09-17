package repository

import (
	"latihan-api-startup/helper"
	"latihan-api-startup/model/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (s *UserRepositoryImpl) Save(user domain.User) domain.User {
	err := s.DB.Create(&user).Error
	helper.PanicIfError(err)

	return user
}

func (s *UserRepositoryImpl) FindByEmail(email string) domain.User {
	var user domain.User

	err := s.DB.Where("email = ?", email).Find(&user).Error

	helper.PanicIfError(err)
	return user
}

func (s *UserRepositoryImpl) FindByID(ID int) domain.User {
	var user domain.User
	err := s.DB.Where("id = ?", ID).Find(&user).Error

	helper.PanicIfError(err)

	return user
}

func (s *UserRepositoryImpl) Update(user domain.User) (domain.User, error) {
	err := s.DB.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *UserRepositoryImpl) FindAll() []domain.User {
	var users []domain.User

	err := s.DB.Find(&users).Error

	helper.PanicIfError(err)

	return users
}
