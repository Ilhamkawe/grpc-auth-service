package port

import (
	"github.com/Ilhamkawe/grpc-auth-service/internal/adapter/database"
	"github.com/Ilhamkawe/grpc-auth-service/internal/application/domain/auth"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type AuthServicePort interface {
	RegisterUser(req auth.RegisterInputUser) (database.User, error)
	Login(req auth.LoginInput) (database.User, error)
	UpdateUserInfo(id int, req auth.UpdateInfoUserInput) (database.User, error)
	IsEmailAvailable(req auth.CheckEmailInput) (bool, error)
	GetUserByID(ID int) (database.User, error)
	ChangePassword(req auth.ChangePasswordInput) (database.User, error)
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
