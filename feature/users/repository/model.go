package repository

import (
	"altafashion_be/feature/users/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
	Profile  string
	Location string
}

func FromDomain(du domain.Core) User {
	return User{
		Model:    gorm.Model{ID: du.ID},
		Fullname: du.Fullname,
		Email:    du.Email,
		Password: du.Password,
		Profile:  du.Profile,
		Location: du.Location,
	}
}

func ToDomain(u User) domain.Core {
	return domain.Core{
		ID:       u.ID,
		Fullname: u.Fullname,
		Email:    u.Email,
		Password: u.Password,
		Profile:  u.Profile,
		Location: u.Location,
	}
}
