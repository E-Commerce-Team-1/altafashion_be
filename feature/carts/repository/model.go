package repository

import (
	"altafashion_be/feature/carts/domain"

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
	Qty         int
	Price       int
	UserID      uint
	Carts       []Cart `gorm:"foreignKey:IdProduct"`
}

type Cart struct {
	gorm.Model
	IdProduct     uint
	UserID        uint
	Name          string `gorm:"-:migration" gorm:"->"`
	Qty           int
	Price         int    `gorm:"-:migration" gorm:"->"`
	Image         string `gorm:"-:migration" gorm:"->"`
	DetailProduct string `gorm:"-:migration" gorm:"->"`
}

func FromDomain(dc domain.Core) Cart {
	return Cart{
		Model:     gorm.Model{ID: dc.ID},
		IdProduct: dc.IdProduct,
		UserID:    dc.UserID,
		Name:      dc.Name,
		Qty:       dc.Qty,
		Price:     dc.Price,
		Image:     dc.Image,
	}
}

func ToDomain(dp Cart) domain.Core {
	return domain.Core{
		ID:        dp.ID,
		IdProduct: dp.IdProduct,
		UserID:    dp.UserID,
		Name:      dp.Name,
		Qty:       dp.Qty,
		Price:     dp.Price,
		Image:     dp.Image,
	}
}

func ToDomainArray(dp []Cart) []domain.Core {
	var res []domain.Core
	for _, val := range dp {
		res = append(res, domain.Core{ID: val.ID, IdProduct: val.IdProduct, UserID: val.UserID, Name: val.Name, Qty: val.Qty, Price: val.Price, Image: val.Image})
	}
	return res
}
