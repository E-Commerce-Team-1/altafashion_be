package services

import (
	"altafashion_be/config"
	"altafashion_be/feature/carts/domain"
	"errors"
	"strings"

	"github.com/labstack/gommon/log"
)

type cartServices struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &cartServices{
		qry: repo,
	}
}

// AddCart implements domain.Service
func (cs *cartServices) AddCart(newCart domain.Core) (domain.Core, error) {
	res, err := cs.qry.Insert(newCart)
	if err != nil {
		if strings.Contains(err.Error(), "cannot") {
			return domain.Core{}, errors.New("cannot buy own product")
		} else if strings.Contains(err.Error(), "stock") {
			return domain.Core{}, errors.New("stock product tidak cukup")
		}
		return domain.Core{}, errors.New("some problem on database")
	}

	return res, nil

}

// DeleteonCart implements domain.Service
func (cs *cartServices) DeleteonCart(id uint) (domain.Core, error) {
	res, err := cs.qry.Delete(id)
	if err != nil {
		return domain.Core{}, err
	}
	return res, err
}

// GetMyCart implements domain.Service
func (cs *cartServices) GetMyCart(id uint) ([]domain.Core, error) {
	res, err := cs.qry.Get(id)
	if err != nil {
		log.Error(err.Error())
		return []domain.Core{}, errors.New("no data")
	}

	return res, nil
}

// UpdateQty implements domain.Service
func (cs *cartServices) UpdateQty(NewCart domain.Core) (domain.Core, error) {
	res, err := cs.qry.Update(NewCart)
	if err != nil {
		if strings.Contains(err.Error(), "coloumn") {
			return domain.Core{}, errors.New(config.REJECTED_DATA)
		}

		return domain.Core{}, errors.New(config.DATABASE_ERROR)
	}

	return res, nil
}
