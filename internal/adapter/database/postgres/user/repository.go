package user

import "github.com/Ilhamkawe/grpc-auth-service/internal/application/domain/auth"

func (a *DatabaseAdapter) Save(user User) (auth.User, error) {
	var result auth.User
	err := a.DB.Create(&user).Error
	if err != nil {
		return result, nil
	}

	result = ToDomainUser(user)
	return result, nil
}

func (a *DatabaseAdapter) FindByEmail(email string) (auth.User, error) {
	var user User
	var result auth.User

	err := a.DB.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return result, err
	}

	result = ToDomainUser(user)
	return result, nil
}

func (a *DatabaseAdapter) FindByID(id int) (auth.User, error) {
	var user User
	var result auth.User

	err := a.DB.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return result, err
	}

	result = ToDomainUser(user)
	return result, nil
}

func (a *DatabaseAdapter) Update(user User) (auth.User, error) {
	var result auth.User
	err := a.DB.Save(&user).Error

	if err != nil {
		return result, err
	}

	result = ToDomainUser(user)
	return result, nil

}

func (a *DatabaseAdapter) ChangePassword(user User) (auth.User, error) {
	var result auth.User
	err := a.DB.Updates(User{PasswordHash: user.PasswordHash}).Error

	if err != nil {
		return result, err
	}

	result = ToDomainUser(user)
	return result, nil
}

func (a *DatabaseAdapter) FindAll() ([]auth.User, error) {
	var user []User
	var result []auth.User
	// select * from user
	err := a.DB.Find(&user).Error

	if err != nil {
		return result, err
	}

	result = ToDomainUsers(user)
	return result, nil
}
