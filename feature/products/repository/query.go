package repository

import (
	"altafashion_be/config"
	"altafashion_be/feature/products/domain"
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

func (rq *repoQuery) ShowAll(category, name string, page int) ([]domain.Core, error) {
	var resQry []Product
	if page != 0 {
		ofst := (page - 1) * 10
		if err := rq.db.Offset(ofst).Limit(10).Order("created_at desc").
			Find(&resQry).Error; err != nil {
			return nil, errors.New(config.DATABASE_ERROR)
		}
	} else if name != "" {
		ofst := (page - 1) * 10
		if err := rq.db.Where("name like ?", "%"+name+"%").
			Offset(ofst).Limit(10).Order("created_at desc").
			Find(&resQry).Error; err != nil {
			return nil, errors.New(config.DATABASE_ERROR)
		}
	} else if category != "" {
		ofst := (page - 1) * 10
		if err := rq.db.Where("category = ?", category).
			Offset(ofst).Limit(10).Order("created_at desc").
			Find(&resQry).Error; err != nil {
			return nil, errors.New(config.DATABASE_ERROR)
		}
	} else {
		if err := rq.db.Find(&resQry).Error; err != nil {
			return nil, errors.New(config.DATABASE_ERROR)
		}
	}

	return ToDomainArray(resQry), nil
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
	var resQry Product
	err := rq.db.Where("id = ?", ID).First(&resQry).Error
	if err != nil {
		return domain.Core{}, errors.New(config.DATABASE_ERROR)
	}

	res := ToDomainDetail(resQry)

	return res, nil
}

func (rq *repoQuery) Update(updateData domain.Core) (domain.Core, error) {
	resQry := FromDomain(updateData)

	err := rq.db.Where("id = ?", resQry.ID).Updates(&resQry).Error
	if err != nil {
		log.Error(config.DATABASE_ERROR)
		return domain.Core{}, err
	}

	updateData = ToDomain(resQry)

	return updateData, nil
}

func (rq *repoQuery) Delete(ID uint) error {
	var resQry Product
	err := rq.db.Where("id = ?", ID).Delete(&resQry)
	if err != nil {
		return errors.New("cant delete data")
	}

	if err.RowsAffected < 1 {
		return errors.New("row isnt affected")
	}

	return nil
}

func (rq *repoQuery) ShowMyProduct(ID uint) ([]domain.Core, error) {
	var resQry []Product
	if err := rq.db.Model(&[]Product{}).Where("products.user_id=?", ID).
		Joins("left join users on users.id = products.user_id").
		Select("products.image", "products.name", "products.description", "products.category", "products.qty", "products.price").
		Scan(&resQry).Error; err != nil {
		return nil, err
	}
	res := ToDomainArray(resQry)
	return res, nil
}
