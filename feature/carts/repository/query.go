package repository

import (
	"altafashion_be/feature/carts/domain"
	"errors"
	"log"

	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) domain.Repository {
	return &repoQuery{
		db: dbConn,
	}
}

// Insert implements domain.Repository
func (rq *repoQuery) Insert(newCart domain.Core) (domain.Core, error) {
	var cnv Cart = FromDomain(newCart)
	var compare Product
	if err := rq.db.Where("id_user = ? AND id = ?", cnv.IdUser, cnv.IdProduct).First(&compare).Error; err == nil {
		log.Print(errors.New("cannot buy own product"))
		return domain.Core{}, errors.New("cannot buy own product")
	}

	if err := rq.db.Where("id = ? AND product_qty>=?", cnv.IdProduct, cnv.ProductQty).First(&compare).Error; err != nil {
		log.Print(errors.New("stock product tidak cukup"))
		return domain.Core{}, errors.New("stock product tidak cukup")
	}

	if err := rq.db.Select("id_product", "id_user", "carts.product_qty").Create(&cnv).Error; err != nil {
		return domain.Core{}, err
	}

	newCart = ToDomain(cnv)
	return newCart, nil

}

// Get implements domain.Repository
func (rq *repoQuery) Get(id uint) ([]domain.Core, error) {
	panic("unimplemented")
}

// Update implements domain.Repository
func (rq *repoQuery) Update(NewCart domain.Core) (domain.Core, error) {
	panic("unimplemented")
}

// Delete implements domain.Repository
func (rq *repoQuery) Delete(id uint) error {
	panic("unimplemented")
}
