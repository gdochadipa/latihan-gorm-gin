package service

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type AuthServiceImp struct {
}

var SECRET_KEY = []byte("BWASTARTUP_s3cr3T_k3Y")

func NewAuthService() AuthService {
	return &AuthServiceImp{}
}

func (s AuthServiceImp) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil

}

func (s AuthServiceImp) ValidateToken(encodeToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodeToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token.")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
