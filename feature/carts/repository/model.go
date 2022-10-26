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
	Products []Product `gorm:"foreignKey:IdUser"`
	Carts    []Cart    `gorm:"foreignKey:IdUser"`
}

type Product struct {
	gorm.Model
	IdUser      uint
	ProductName string
	ProductQty  uint
	Price       int
	Image       string
	Carts       []Cart `gorm:"foreignKey:IdProduct"`
}

type Cart struct {
	gorm.Model
	IdProduct   uint
	IdUser      uint
	ProductName string `gorm:"-:migration" gorm:"->"`
	ProductQty  uint
	Price       int    `gorm:"-:migration" gorm:"->"`
	Image       string `gorm:"-:migration" gorm:"->"`
}

func FromDomain(dc domain.Core) Cart {
	return Cart{
		Model:       gorm.Model{ID: dc.ID},
		IdProduct:   dc.IdProduct,
		IdUser:      dc.IdUser,
		ProductName: dc.ProductName,
		ProductQty:  dc.ProductQty,
		Price:       dc.Price,
		Image:       dc.Image,
	}
}

func ToDomain(dp Cart) domain.Core {
	return domain.Core{
		ID:          dp.ID,
		IdProduct:   dp.IdProduct,
		IdUser:      dp.IdUser,
		ProductName: dp.ProductName,
		ProductQty:  dp.ProductQty,
		Price:       dp.Price,
		Image:       dp.Image,
	}
}

func ToDomainArray(dp []Cart) []domain.Core {
	var res []domain.Core
	for _, val := range dp {
		res = append(res, domain.Core{ID: val.ID, IdProduct: val.IdProduct, IdUser: val.IdUser, ProductName: val.ProductName, ProductQty: val.ProductQty, Price: val.Price, Image: val.Image})
	}
	return res
}
