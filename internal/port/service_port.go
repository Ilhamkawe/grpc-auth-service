package port

import (
	"github.com/Ilhamkawe/grpc-auth-service/internal/adapter/database"
	"github.com/Ilhamkawe/grpc-auth-service/internal/application/domain/auth"
	"github.com/dgrijalva/jwt-go"
)

type AuthServicePort interface {
	RegisterUser(req auth.RegisterInputUser) (database.User, error)
	Login(req auth.LoginInput) (database.User, error)
	UpdateUserInfo(id int, req auth.UpdateInfoUserInput) (database.User, error)
	//UploadAvatar(context.Context, *UploadAvatarRequest) (*BooleanResponse, error)
	IsEmailAvailable(req auth.CheckEmailInput) (bool, error)
}

type JwtServicePort interface {
	GenerateToken(UserID int) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}
