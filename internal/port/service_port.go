package port

import (
	"time"

	"github.com/Ilhamkawe/grpc-auth-service/internal/application/domain/auth"
	"github.com/dgrijalva/jwt-go"
)

type AuthServicePort interface {
	RegisterUser(req auth.RegisterInputUser) (auth.User, error)
	Login(req auth.LoginInput) (auth.User, error)
	UpdateUserInfo(id int, req auth.UpdateInfoUserInput) (auth.User, error)
	IsEmailAvailable(req auth.CheckEmailInput) (bool, error)
	GetUserByID(ID int) (auth.User, error)
	ChangePassword(req auth.ChangePasswordInput) (auth.User, error)
	Logout(tokenString string, exp time.Duration) error
}

type JwtServicePort interface {
	GenerateToken(UserID int) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type FileServicePort interface {
	SetFile(fileName, path string) error
	Write(chunk []byte) error
	Close() error
}
