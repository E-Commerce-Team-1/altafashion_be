package repository

import (
	"altafashion_be/feature/users/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname string
	Email    string `gorm:"unique"`
	Password string
	Profile  string
	Location string
	Products []Product `gorm:"foreignKey:UserID"`
	Carts    []Cart    `gorm:"foreignKey:UserID"`
}

type Product struct {
	gorm.Model
	Image       string
	Name        string
	Description string
	Category    string
	Qty         uint
	Price       int
	UserID      uint
	Carts       []Cart `gorm:"foreignKey:IdProduct"`
}

type Cart struct {
	gorm.Model
	IdProduct uint
	UserID    uint
	Name      string `gorm:"-:migration" gorm:"->"`
	Qty       uint
	Price     int    `gorm:"-:migration" gorm:"->"`
	Image     string `gorm:"-:migration" gorm:"->"`
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
