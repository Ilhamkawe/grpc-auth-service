package port

import (
	"context"
	"time"

	"github.com/Ilhamkawe/grpc-auth-service/internal/adapter/database/postgres/user"
	"github.com/Ilhamkawe/grpc-auth-service/internal/application/domain/auth"
)

type AuthDatabasePort interface {
	Save(user user.User) (auth.User, error)
	FindByEmail(email string) (auth.User, error)
	FindByID(id int) (auth.User, error)
	Update(user user.User) (auth.User, error)
	ChangePassword(user user.User) (auth.User, error)
	//GetUserByID(ID int) (db.User, error)
}
type JwtRedisPort interface {
	SetKey(ctx context.Context, key, value string, ttl time.Duration) error
	GetKey(ctx context.Context, key string) (string, error)
}
