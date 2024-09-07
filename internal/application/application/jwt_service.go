package application

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type JwtService struct {
}

func NewJwtService() *JwtService {
	return &JwtService{}
}

var SECRET_KEY = []byte("test123")

func (s *JwtService) GenerateToken(UserID int) (string, error) {
	claim := jwt.MapClaims{}

	claim["user_id"] = UserID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *JwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Token Tidak Valid")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
