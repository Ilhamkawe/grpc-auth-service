package port

import (
	"context"
	db "github.com/Ilhamkawe/grpc-auth-service/internal/adapter/database"
	"time"
)

type AuthDatabasePort interface {
	Save(user db.User) (db.User, error)
	FindByEmail(email string) (db.User, error)
	FindByID(id int) (db.User, error)
	Update(user db.User) (db.User, error)
	ChangePassword(user db.User) (db.User, error)
	//GetUserByID(ID int) (db.User, error)
}
type JwtRedisPort interface {
	SetKey(ctx context.Context, key, value string, ttl time.Duration) error
	GetKey(ctx context.Context, key string) (string, error)
}
