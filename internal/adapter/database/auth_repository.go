package database

func (a *DatabaseAdapter) Save(user User) (User, error) {
	err := a.db.Create(&user).Error
	if err != nil {
		return user, nil
	}

	return user, nil
}

func (a *DatabaseAdapter) FindByEmail(email string) (User, error) {
	var user User

	err := a.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (a *DatabaseAdapter) FindByID(id int) (User, error) {
	var user User

	err := a.db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (a *DatabaseAdapter) Update(user User) (User, error) {
	err := a.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (a *DatabaseAdapter) ChangePassword(user User) (User, error) {
	err := a.db.Updates(User{PasswordHash: user.PasswordHash}).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (a *DatabaseAdapter) FindAll() ([]User, error) {
	var user []User
	// select * from user
	err := a.db.Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
