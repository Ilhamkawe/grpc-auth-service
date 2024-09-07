package grpc

import (
	"context"
	protoAuth "github.com/Ilhamkawe/crowdfunding-proto/protogen/go/auth"
	"github.com/Ilhamkawe/grpc-auth-service/internal/application/domain/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *GrpcAdapter) RegisterUser(ctx context.Context, req *protoAuth.RegisterUserRequest) (*protoAuth.User, error) {

	reqUser := auth.RegisterInputUser{
		Email:      req.Email,
		Name:       req.Name,
		Occupation: req.Occupation,
		Password:   req.Password,
	}
	newUser, err := a.authService.RegisterUser(reqUser)

	if err != nil {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			"Register Gagal")
	}

	token, err := a.jwtService.GenerateToken(newUser.ID)
	if err != nil {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			"Register Gagal")
	}

	return &protoAuth.User{
		Name:           newUser.Name,
		Occupation:     newUser.Occupation,
		Email:          newUser.Email,
		PasswordHash:   newUser.PasswordHash,
		Role:           newUser.Role,
		AvatarFileName: newUser.AvatarFileName,
		Token:          token,
	}, nil
}

func (a *GrpcAdapter) UpdateUserInfo(ctx context.Context, req *protoAuth.UpdateInfoUserRequest) (*protoAuth.User, error) {
	reqUpdate := auth.UpdateInfoUserInput{
		: req.Email
	}

	_, err := a.authService.UpdateUserInfo(reqUpdate)
}

func (a *GrpcAdapter) Login(ctx context.Context, req *protoAuth.LoginRequest) (*protoAuth.User, error) {

	reqLogin := auth.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	}

	loggedInUser, err := a.authService.Login(reqLogin)

	if err != nil {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			"Login Gagal")
	}

	token, err := a.jwtService.GenerateToken(loggedInUser.ID)

	if err != nil {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			"Login Gagal")
	}

	return &protoAuth.User{
		Name:           loggedInUser.Name,
		Occupation:     loggedInUser.Occupation,
		Email:          loggedInUser.Email,
		PasswordHash:   loggedInUser.PasswordHash,
		Role:           loggedInUser.Role,
		AvatarFileName: loggedInUser.AvatarFileName,
		Token:          token,
	}, nil
}
