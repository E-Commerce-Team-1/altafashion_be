package services

import (
	"altafashion_be/config"
	"altafashion_be/feature/users/domain"
	"errors"
	"strings"

	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &userService{
		qry: repo,
	}
}

// Register implements domain.Service
func (us *userService) Register(newUser domain.Core) (domain.Core, error) {

	if strings.TrimSpace(newUser.Email) == "" || strings.TrimSpace(newUser.Password) == "" {
		return domain.Core{}, errors.New("email or password empty")
	}
	generate, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("error on bcrypt", err.Error())
		return domain.Core{}, errors.New("cannot encrypt password")
	}
	newUser.Password = string(generate)

	res, err := us.qry.Insert(newUser)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Core{}, errors.New("rejected from database")
		}
		return domain.Core{}, errors.New("some problem on database")
	}
	return res, nil
}

// Login implements domain.Service
func (us *userService) Login(existUser domain.Core) (domain.Core, error) {
	if strings.TrimSpace(existUser.Email) == "" || strings.TrimSpace(existUser.Password) == "" {
		return domain.Core{}, errors.New("email or password empty")
	}
	res, err := us.qry.GetUser(existUser)
	if err != nil {
		if strings.Contains(err.Error(), "table") {
			return domain.Core{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.Core{}, errors.New("no data")
		}
	}
	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(existUser.Password))
	if err != nil {
		return domain.Core{}, errors.New("password not match")
	}
	return res, nil
}

// UpdateProfile implements domain.Service
func (us *userService) UpdateProfile(updateData domain.Core, id uint) (domain.Core, error) {
	if updateData.Password != "" {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(updateData.Password), bcrypt.DefaultCost)

		updateData.Password = string(hashed)
	}

	res, err := us.qry.Update(updateData, id)
	if err != nil {
		if strings.Contains(err.Error(), config.DUPLICATED_DATA) {
			return domain.Core{}, errors.New(config.REJECTED_DATA)
		}
	}

	return res, nil
}

// Deactivate implements domain.Service
func (us *userService) Deactivate(id uint) (domain.Core, error) {
	res, err := us.qry.Delete(id)
	if err != nil {
		if strings.Contains(err.Error(), "table") {
			return domain.Core{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.Core{}, errors.New("no data")
		}
	}
	return res, nil
}

// ShowByUsername implements domain.Service
func (us *userService) ShowByUsername(username string) (domain.Core, error) {
	res, err := us.qry.GetByUsername(username)
	if err != nil {
		if strings.Contains(err.Error(), "table") {
			return domain.Core{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.Core{}, errors.New("no data")
		}
	}
	return res, nil
}

// MyProfile implements domain.Service
func (us *userService) MyProfile(id uint) (domain.Core, error) {
	res, err := us.qry.GetMyUser(id)
	if err != nil {
		if strings.Contains(err.Error(), "table") {
			return domain.Core{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.Core{}, errors.New("no data")
		}
	}
	return res, nil
}
