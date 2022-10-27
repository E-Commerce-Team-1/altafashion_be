package services

import (
	"altafashion_be/config"
	"altafashion_be/feature/products/domain"
	"errors"
	"log"
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

func (ps *productServices) GetAll(category, name string, page int) ([]domain.Core, error) {
	res, err := ps.qry.ShowAll(category, name, page)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New(config.DATA_NOT_FOUND)
	}

	return res, nil
}

func (ps *productServices) AddProduct(newProduct domain.Core) (domain.Core, error) {
	res, err := ps.qry.Insert(newProduct)
	if err != nil {
		return domain.Core{}, errors.New(config.DUPLICATED_DATA)
	}

	return res, nil
}

func (ps *productServices) GetDetail(ID uint) (domain.Core, error) {
	res, err := ps.qry.ShowDetail(ID)
	if err != nil {
		return domain.Core{}, errors.New(config.DATA_NOT_FOUND)
	}

	return res, nil
}

func (ps *productServices) EditProduct(updateData domain.Core) (domain.Core, error) {
	res, err := ps.qry.Update(updateData)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Core{}, errors.New(config.REJECTED_DATA)
		}

		return domain.Core{}, errors.New(config.DATABASE_ERROR)
	}

	return res, nil
}

func (ps *productServices) Destroy(ID uint) error {
	err := ps.qry.Delete(ID)
	if err != nil {
		log.Println(err.Error())
		if strings.Contains(err.Error(), "table") {
			return errors.New(config.DATABASE_ERROR)
		} else if strings.Contains(err.Error(), "found") {
			return errors.New(config.DATA_NOT_FOUND)
		}
	}
	return nil
}

func (ps *productServices) GetMyProduct(ID uint) ([]domain.Core, error) {
	res, err := ps.qry.ShowMyProduct(ID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New(config.DATA_NOT_FOUND)
	}

	return res, nil
}
