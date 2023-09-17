package service

import (
	"errors"
	"latihan-api-startup/model/domain"
	"latihan-api-startup/model/web"
	"latihan-api-startup/repository"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (s *UserServiceImpl) RegisterUser(input web.RegisterUserInput) (domain.User, error) {
	user := domain.User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	passwordHas, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHas)
	user.Role = "user"

	newUser := s.UserRepository.Save(user)

	return newUser, nil

}

func (s *UserServiceImpl) Login(input web.LoginInput) (domain.User, error) {
	email := input.Email
	password := input.Password

	user := s.UserRepository.FindByEmail(email)

	if user.ID == 0 {
		return user, errors.New("No user found on that email")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s UserServiceImpl) IsEmailAvailable(input web.CheckEmailInput) bool {
	email := input.Email

	user := s.UserRepository.FindByEmail(email)

	// if user.ID == 0 is mean no email in data users
	return user.ID == 0
}

func (s UserServiceImpl) SaveAvatar(ID int, fileLocation string) (domain.User, error) {
	user := s.UserRepository.FindByID(ID)

	user.AvatarFileName = fileLocation

	updatedUser, err := s.UserRepository.Update(user)

	if err != nil {
		return updatedUser, err
	}

	if updatedUser.ID == 0 {
		return updatedUser, errors.New("No user found on with that ID")
	}

	return updatedUser, nil
}

func (s UserServiceImpl) GetUserByID(ID int) (domain.User, error) {
	user := s.UserRepository.FindByID(ID)
	if user.ID == 0 {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}

func (s UserServiceImpl) GetAllUsers() ([]domain.User, error) {
	users := s.UserRepository.FindAll()

	return users, nil
}

func (s UserServiceImpl) UpdateUser(input web.FormUpdateUserInput) (domain.User, error) {
	user := s.UserRepository.FindByID(input.ID)

	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	updatedUser, err := s.UserRepository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}
