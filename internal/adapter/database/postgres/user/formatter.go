package user

import (
	"github.com/Ilhamkawe/grpc-auth-service/internal/application/domain/auth"
)

func ToDomainUser(dbUser User) auth.User {
	return auth.User{
		ID:             dbUser.ID,
		Name:           dbUser.Name,
		Occupation:     dbUser.Occupation,
		Email:          dbUser.Email,
		AvatarFileName: dbUser.AvatarFileName,
		Role:           dbUser.Role,
	}
}

func ToDBUser(domainUser auth.User) User {
	return User{
		ID:             domainUser.ID,
		Name:           domainUser.Name,
		Occupation:     domainUser.Occupation,
		Email:          domainUser.Email,
		AvatarFileName: domainUser.AvatarFileName,
		Role:           domainUser.Role,
	}
}

func ToDomainUsers(users []User) []auth.User {
	result := make([]auth.User, len(users))
	for i, u := range users {
		result[i] = ToDomainUser(u)
	}
	return result
}

func ToDBUsers(users []auth.User) []User {
	result := make([]User, len(users))
	for i, u := range users {
		result[i] = ToDBUser(u)
	}

	return result
}
