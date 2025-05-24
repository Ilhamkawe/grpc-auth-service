package application

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Ilhamkawe/grpc-auth-service/internal/adapter/database/postgres/user"
	"github.com/Ilhamkawe/grpc-auth-service/internal/application/domain/auth"
	"github.com/Ilhamkawe/grpc-auth-service/internal/application/helper/formatter"
	"github.com/Ilhamkawe/grpc-auth-service/internal/port"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	db  port.AuthDatabasePort
	rdb port.JwtRedisPort
}

func NewAuthService(dbPort port.AuthDatabasePort, rdbPort port.JwtRedisPort) *AuthService {
	return &AuthService{
		db:  dbPort,
		rdb: rdbPort,
	}
}

func (s *AuthService) RegisterUser(req auth.RegisterInputUser) (auth.User, error) {
	newUser := auth.User{}
	user := user.User{}
	user.Name = req.Name
	user.Occupation = req.Occupation
	user.AvatarFileName = "images/default-user.jpg"
	user.Email = req.Email

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return newUser, err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "User"

	newUser, err = s.db.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *AuthService) Login(req auth.LoginInput) (auth.User, error) {
	email := req.Email
	password := req.Password

	user, err := s.db.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User tidak ditemukan")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *AuthService) UpdateUserInfo(id int, req auth.UpdateInfoUserInput) (auth.User, error) {
	user, err := s.db.FindByID(id)

	if err != nil {
		return user, nil
	}

	user.Name = req.Name
	user.Occupation = req.Occupation

	updatedUser, err := s.db.Update(formatter.ToDBUser(user))
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, err
}

func (s *AuthService) IsEmailAvailable(req auth.CheckEmailInput) (bool, error) {
	email := req.Email

	user, err := s.db.FindByEmail(email)

	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *AuthService) GetUserByID(id int) (auth.User, error) {
	user, err := s.db.FindByID(id)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User Tidak Ditemukan")
	}

	return user, nil
}

//func (s *AuthService) GetAllUsers() ([]database.User, error) {
//	users, err := s.db.FindAll()
//	if err != nil {
//		return users, err
//	}
//
//	return users, nil
//}

func (s *AuthService) ChangePassword(input auth.ChangePasswordInput) (auth.User, error) {
	user, err := s.db.FindByID(input.ID)

	if err != nil {
		return user, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.PasswordHash), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)

	newPassword, err := s.db.ChangePassword(formatter.ToDBUser(user))
	if err != nil {
		return newPassword, err
	}
	return newPassword, nil
}

func (s *AuthService) Logout(tokenString string, exp time.Duration) error {
	ctx := context.Background()
	return s.rdb.SetKey(ctx, fmt.Sprintf("blacklist_token:%s", tokenString), "blacklisted", exp)
}
