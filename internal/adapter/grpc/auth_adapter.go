package grpc

import (
	"context"
	"fmt"
	protoAuth "github.com/Ilhamkawe/crowdfunding-proto/protogen/go/auth"
	"github.com/Ilhamkawe/grpc-auth-service/internal/adapter/database"
	app "github.com/Ilhamkawe/grpc-auth-service/internal/application/application"
	"github.com/Ilhamkawe/grpc-auth-service/internal/application/domain/auth"
	"github.com/Ilhamkawe/grpc-auth-service/internal/application/helper/logger"
	"github.com/Ilhamkawe/grpc-auth-service/internal/interceptor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"path/filepath"
	"time"
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

func (a *GrpcAdapter) UpdateUserInfo(ctx context.Context, req *protoAuth.UpdateInfoUserRequest) (*protoAuth.BooleanResponse, error) {
	reqUpdate := auth.UpdateInfoUserInput{
		Name:       req.Name,
		Occupation: req.Occupation,
	}

	_, err := a.authService.UpdateUserInfo(int(req.Id), reqUpdate)

	if err != nil {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			"Update Gagal")
	}

	return &protoAuth.BooleanResponse{
		Status: true,
	}, nil

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

func (a *GrpcAdapter) UploadAvatar(stream protoAuth.AuthService_UploadAvatarServer) error {
	l := logger.Logger{}
	file := app.NewFileService()
	var fileSize uint32
	fileSize = 0
	defer func() {
		if err := file.OutputFile.Close(); err != nil {
			fmt.Errorf(err.Error())
		}
	}()

	for {
		req, err := stream.Recv()
		if file.FilePath == "" {
			file.SetFile(req.GetAvatar(), "client_files")
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			return status.Errorf(
				codes.FailedPrecondition,
				"Login Gagal")
		}

		chunk := req.GetChunk()
		fileSize += uint32(len(chunk))

		l.Debug("received a chunk with size: %d", fileSize)

		if err := file.Write(chunk); err != nil {
			return l.LogError(status.Error(codes.Internal, err.Error()))
		}
	}
	fileName := filepath.Base(file.FilePath)
	l.Debug("saved file: %s, size: %d", fileName, fileSize)
	return stream.SendAndClose(&protoAuth.BooleanResponse{
		Status: true,
	})
}

func (a *GrpcAdapter) FetchUser(ctx context.Context, req *protoAuth.SendID) (*protoAuth.User, error) {

	currentUser, ok := ctx.Value(interceptor.CurrentUserKey).(database.User) // Sesuaikan dengan tipe user Anda\

	fmt.Println(currentUser)
	if !ok {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			"User Not Founds")
	}

	// Menggunakan currentUser untuk membuat response
	return &protoAuth.User{
		Name:           currentUser.Name,
		Occupation:     currentUser.Occupation,
		Email:          currentUser.Email,
		PasswordHash:   currentUser.PasswordHash,
		Role:           currentUser.Role,
		AvatarFileName: currentUser.AvatarFileName,
		Token:          "some_generated_token", // Sesuaikan token jika perlu
	}, nil
}

func (a *GrpcAdapter) Logout(ctx context.Context, _ *emptypb.Empty) (*protoAuth.BooleanResponse, error) {
	currentUser, ok := ctx.Value(interceptor.CurrentUserKey).(database.User) // Sesuaikan dengan tipe user Anda

	if !ok {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			"User not found")
	}

	// Mengatur TTL dalam detik
	durationInSeconds := 28800
	duration := time.Duration(durationInSeconds) * time.Second

	err := a.authService.Logout(currentUser.Token, duration)

	if err != nil {
		if !ok {
			return nil, status.Errorf(
				codes.FailedPrecondition,
				"Failed")
		}
	}

	return &protoAuth.BooleanResponse{
		Status: true,
	}, nil
	//a.authService.Logout()
}

func (a *GrpcAdapter) ChangePassword(ctx context.Context, req *protoAuth.ChangePasswordRequest) (*protoAuth.BooleanResponse, error) {
	type contextKey string

	const currentUserKey = contextKey("currentUser")

	currentUser, ok := ctx.Value(currentUserKey).(*database.User) // Sesuaikan dengan tipe user Anda
	if !ok {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			"User not found")
	}

	req.Id = int64(currentUser.ID)

	_, err := a.authService.ChangePassword(auth.ChangePasswordInput{PasswordHash: req.PasswordHash})

	if err != nil {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			"Failed Change Passsword")
	}

	return &protoAuth.BooleanResponse{
		Status: true,
	}, nil
}

func (a *GrpcAdapter) mustEmbedUnimplementedAuthServiceServer() {
	//TODO implement me
	panic("implement me")
}
