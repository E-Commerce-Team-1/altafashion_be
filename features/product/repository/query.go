package repository

import (
	"altafashion_be/config"
	"altafashion_be/features/product/domain"
	"errors"

	"github.com/labstack/gommon/log"
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

func (rq *repoQuery) ShowAll(category, key string, page int) ([]domain.Core, error) {
	return nil, nil
}

func (rq *repoQuery) Insert(newProduct domain.Core) (domain.Core, error) {
	input := FromDomain(newProduct)

	if err := rq.db.Create(&input).Error; err != nil {
		return domain.Core{}, nil
	}

	newProduct = ToDomain(input)

	return newProduct, nil
}

func (rq *repoQuery) ShowDetail(ID uint) (domain.Core, error) {
	return domain.Core{}, nil
}

func (rq *repoQuery) Update(updateData domain.Core, ID uint) (domain.Core, error) {
	resQry := FromDomain(updateData)

	err := rq.db.Where("id = ?", ID).Updates(&resQry).Error
	if err != nil {
		log.Error(config.DATABASE_ERROR)
		return domain.Core{}, err
	}

	updateData = ToDomain(resQry)

	return updateData, nil
}

func (rq *repoQuery) Delete(ID uint) error {
	err := rq.db.Where("id = ?", ID).Delete(&Product{})
	if err != nil {
		return errors.New("cant delete data")
	}

	if err.RowsAffected < 1 {
		return errors.New("row isnt affected")
	}

	return nil
}

func (rq *repoQuery) ShowMyProduct() ([]domain.Core, error) {
	return nil, nil
}
