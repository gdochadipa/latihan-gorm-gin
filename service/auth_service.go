package service

import "github.com/dgrijalva/jwt-go"

type AuthService interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(encodeToken string) (*jwt.Token, error)
}
