package repository

import (
	"altafashion_be/feature/users/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
	Profile  string
	Phone    string
	Location string
}

func FromDomain(du domain.Core) User {
	return User{
		Model:    gorm.Model{ID: du.ID},
		Username: du.Username,
		Email:    du.Email,
		Password: du.Password,
		Profile:  du.Profile,
		Phone:    du.Phone,
		Location: du.Location,
	}
}

func ToDomain(u User) domain.Core {
	return domain.Core{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		Profile:  u.Profile,
		Phone:    u.Phone,
		Location: u.Location,
	}
}
