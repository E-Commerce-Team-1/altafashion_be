package repository

import (
	"altafashion_be/feature/products/domain"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Image       string
	Name        string
	Description string
	Category    string
	Qty         uint
	Price       int
	UserID      uint
}

type User struct {
	gorm.Model
	Fullname string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
	Profile  string
	Location string
	Products []Product `gorm:"foreignKey:UserID"`
}

func FromDomain(dc domain.Core) Product {
	return Product{
		Model:       gorm.Model{ID: dc.ID},
		Image:       dc.Image,
		Name:        dc.Name,
		Description: dc.Description,
		Category:    dc.Category,
		Qty:         dc.Qty,
		Price:       dc.Price,
		UserID:      dc.UserID,
	}
}

func ToDomain(p Product) domain.Core {
	return domain.Core{
		ID:          p.ID,
		Image:       p.Image,
		Name:        p.Name,
		Description: p.Description,
		Category:    p.Category,
		Qty:         p.Qty,
		Price:       p.Price,
		UserID:      p.UserID,
	}
}

func ToDomainArray(listProduct []Product) []domain.Core {
	var res []domain.Core
	for _, val := range listProduct {
		res = append(res, domain.Core{
			ID:          val.ID,
			Image:       val.Image,
			Name:        val.Name,
			Description: val.Description,
			Category:    val.Category,
			Qty:         val.Qty,
			Price:       val.Price,
			UserID:      val.UserID,
		})
	}

	return res
}

func ToDomainDetail(p Product) domain.Core {
	return domain.Core{
		ID:          p.ID,
		Image:       p.Image,
		Name:        p.Name,
		Description: p.Description,
		Category:    p.Category,
		Qty:         p.Qty,
		Price:       p.Price,
		UserID:      p.UserID,
	}
}
