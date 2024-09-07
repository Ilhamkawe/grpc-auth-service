package port

import db "github.com/Ilhamkawe/grpc-auth-service/internal/adapter/database"

type AuthDatabasePort interface {
	Save(user db.User) (db.User, error)
	FindByEmail(email string) (db.User, error)
	FindByID(id int) (db.User, error)
	Update(user db.User) (db.User, error)
	ChangePassword(user db.User) (db.User, error)
	FindAll() ([]db.User, error)
}
