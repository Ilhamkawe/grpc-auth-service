package database

import "time"

type User struct {
	ID             int `gorm:"primaryKey"`
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	Token          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (User) TableName() string {
	return "User"
}
