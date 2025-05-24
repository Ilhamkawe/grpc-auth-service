package formatter

import (
	userModel "github.com/Ilhamkawe/grpc-auth-service/internal/adapter/database/postgres/user"
	"github.com/Ilhamkawe/grpc-auth-service/internal/application/domain/auth"
)

func ToDomainUser(dbUser userModel.User) auth.User {
	return auth.User{
		ID:             dbUser.ID,
		Name:           dbUser.Name,
		Occupation:     dbUser.Occupation,
		Email:          dbUser.Email,
		AvatarFileName: dbUser.AvatarFileName,
		Role:           dbUser.Role,
	}
}

func ToDBUser(domainUser auth.User) userModel.User {
	return userModel.User{
		ID:             domainUser.ID,
		Name:           domainUser.Name,
		Occupation:     domainUser.Occupation,
		Email:          domainUser.Email,
		AvatarFileName: domainUser.AvatarFileName,
		Role:           domainUser.Role,
	}
}

func ToDomainUsers(users []userModel.User) []auth.User {
	result := make([]auth.User, len(users))
	for i, u := range users {
		result[i] = ToDomainUser(u)
	}
	return result
}

func ToDBUsers(users []auth.User) []userModel.User {
	result := make([]userModel.User, len(users))
	for i, u := range users {
		result[i] = ToDBUser(u)
	}

	return result
}
