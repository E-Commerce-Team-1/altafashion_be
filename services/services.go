package services

import (
	"altafashion_be/config"
	"altafashion_be/features/product/domain"
	"errors"
	"strings"
)

type productServices struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Services {
	return &productServices{
		qry: repo,
	}
}

func (ps *productServices) GetAll(category, key string, page int) ([]domain.Core, error) {
	return nil, nil
}

func (ps *productServices) AddProduct(newProduct domain.Core) (domain.Core, error) {
	res, err := ps.qry.Insert(newProduct)
	if err != nil {
		return domain.Core{}, errors.New(config.DUPLICATED_DATA)
	}

	return res, nil
}

func (ps *productServices) GetDetail(ID uint) (domain.Core, error) {
	return domain.Core{}, nil
}

func (ps *productServices) EditProduct(updateData domain.Core, ID uint) (domain.Core, error) {
	res, err := ps.qry.Update(updateData, ID)
	if err != nil {
		if strings.Contains(err.Error(), config.DUPLICATED_DATA) {
			return domain.Core{}, errors.New(config.DATABASE_ERROR)
		}
	}

	return res, nil
}

func (ps *productServices) Destroy(ID uint) error {
	err := ps.qry.Delete(ID)
	if err != nil {
		return errors.New(config.DATA_NOT_FOUND)
	}
	return nil
}

func (ps *productServices) GetMyProduct() ([]domain.Core, error) {
	return nil, nil
}
