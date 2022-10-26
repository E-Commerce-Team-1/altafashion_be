package services

import (
	"altafashion_be/config"
	"altafashion_be/feature/users/domain"
	"altafashion_be/utils/jwt"
	"errors"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
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
		return domain.Core{}, errors.New(config.ENCRYPT_ERROR)
	}
	newUser.Password = string(generate)

	res, err := us.qry.Insert(newUser)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Core{}, errors.New(config.DUPLICATED_DATA)
		}
		return domain.Core{}, errors.New(config.DATABASE_ERROR)
	}
	return res, nil
}

// Login implements domain.Service
func (us *userService) Login(existUser domain.Core) (domain.Core, string, error) {
	if strings.TrimSpace(existUser.Email) == "" || strings.TrimSpace(existUser.Password) == "" {
		return domain.Core{}, "", errors.New("email or password empty")
	}
	res, err := us.qry.GetUser(existUser)
	if err != nil {
		if strings.Contains(err.Error(), "table") {
			return domain.Core{}, "", errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.Core{}, "", errors.New("no data")
		}
	}
	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(existUser.Password))
	if err != nil {
		return domain.Core{}, "", errors.New("password not match")
	}
	token, err := jwt.GenerateJWTToken(res.ID)
	if err != nil {
		return domain.Core{}, "", err
	}

	return res, token, nil
}

// UpdateProfile implements domain.Service
func (us *userService) UpdateProfile(updateData domain.Core, c echo.Context) (domain.Core, error) {
	id, _ := jwt.ExtractToken(c)
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
func (us *userService) Deactivate(c echo.Context) (domain.Core, error) {
	id := jwt.ExtractIdToken(c)
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

// ShowByFullname implements domain.Service
func (us *userService) ShowByEmail(Email string) (domain.Core, error) {
	res, err := us.qry.GetByEmail(Email)
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
func (us *userService) MyProfile(c echo.Context) (domain.Core, error) {
	id := jwt.ExtractIdToken(c)
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

func (rs *userService) IsAuthorized(c echo.Context) error {
	id, exp := jwt.ExtractToken(c)
	// loggo.Println("id dr tken = ", id)
	// loggo.Println("exp dr tken = ", exp)
	if id == 0 {
		return errors.New("request not authorized, please check token, user not found.")
	} else if time.Now().Unix() > exp {
		return errors.New("request not authorized, please check token, expired token.")
	} else {
		return nil
	}
}
