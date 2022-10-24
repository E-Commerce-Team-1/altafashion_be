package repository

import (
	"altafashion_be/config"
	"altafashion_be/feature/users/domain"

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

func (rq *repoQuery) Insert(newUser domain.Core) (domain.Core, error) {
	var cnv User = FromDomain(newUser)
	if err := rq.db.Create(&cnv).Error; err != nil {
		log.Error("error on adding user", err.Error())
		return domain.Core{}, err
	}
	newUser = ToDomain(cnv)
	return newUser, nil
}

func (rq *repoQuery) GetUser(existUser domain.Core) (domain.Core, error) {
	var resQuery User
	if err := rq.db.First(&resQuery, "email = ?", existUser.Email).Error; err != nil {
		log.Error("error on get user login", err.Error())
		return domain.Core{}, nil
	}
	res := ToDomain(resQuery)
	return res, nil
}

func (rq *repoQuery) Update(updateData domain.Core, id uint) (domain.Core, error) {
	var resQry User
	resQry = FromDomain(updateData)

	err := rq.db.Where("id = ?", id).Updates(resQry).Error
	if err != nil {
		log.Error(config.DATABASE_ERROR)
		return domain.Core{}, err
	}
	updateData = ToDomain(resQry)

	return updateData, nil
}

func (rq *repoQuery) Delete(id uint) (domain.Core, error) {
	var data User
	if err := rq.db.Delete(&data, "id = ?", id).Error; err != nil {
		log.Error("error on deleting user", err.Error())
		return domain.Core{}, err
	}
	res := ToDomain(data)
	return res, nil
}

func (rq *repoQuery) GetByUsername(username string) (domain.Core, error) {
	var resQuery User
	if err := rq.db.First(&resQuery, "username = ?", username).Error; err != nil {
		log.Error("error on get by email", err.Error())
		return domain.Core{}, err
	}
	res := ToDomain(resQuery)
	return res, nil
}

func (rq *repoQuery) GetMyUser(userID uint) (domain.Core, error) {
	var resQuery User
	if err := rq.db.First(&resQuery, "id = ?", userID).Error; err != nil {
		log.Error("error on get my user", err.Error())
		return domain.Core{}, err
	}
	res := ToDomain(resQuery)
	return res, nil
}
