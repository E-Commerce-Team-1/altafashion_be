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
	if err := rq.db.Where("user_id = ? AND id = ?", cnv.UserID, cnv.IdProduct).First(&compare).Error; err == nil {
		log.Print(errors.New("cannot buy own product"))
		return domain.Core{}, errors.New("cannot buy own product")
	}

	if err := rq.db.Where("id = ? AND qty >=?", cnv.IdProduct, cnv.Qty).First(&compare).Error; err != nil {
		log.Print(errors.New("stock product tidak cukup"))
		return domain.Core{}, errors.New("stock product tidak cukup")
	}

	if err := rq.db.Select("id_product", "user_id", "carts.qty").Create(&cnv).Error; err != nil {
		return domain.Core{}, err
	}

	newCart = ToDomain(cnv)
	return newCart, nil

}

// Get implements domain.Repository
func (rq *repoQuery) Get(id uint) ([]domain.Core, error) {
	var resQry []Cart
	if err := rq.db.Model(&[]Cart{}).Where("carts.user_id=?", id).
		Joins("left join products on products.id = carts.id_product").
		Joins("left join users on users.id = carts.user_id").
		Select("carts.qty", "carts.id", "id_product", "carts.user_id", "products.name", "products.price", "image").
		Scan(&resQry).Error; err != nil {
		return nil, err
	}
	res := ToDomainArray(resQry)
	return res, nil
}

// Update implements domain.Repository
func (rq *repoQuery) Update(NewCart domain.Core) (domain.Core, error) {
	cnv := FromDomain(NewCart)
	if err := rq.db.Where("id = ?", cnv.ID).Updates(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	NewCart = ToDomain(cnv)
	return NewCart, nil
}

// Delete implements domain.Repository
func (rq *repoQuery) Delete(id uint) (domain.Core, error) {
	if err := rq.db.Where("id = ?", id).Delete(&Cart{}); err != nil {
		return domain.Core{}, err.Error
	}
	return domain.Core{}, nil
}
