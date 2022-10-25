package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID       uint
	Fullname string
	Email    string
	Password string
	Profile  string
	Location string
}

type Repository interface {
	Insert(newUser Core) (Core, error)
	GetUser(existUser Core) (Core, error)
	Update(updateData Core, id uint) (Core, error)
	Delete(id uint) (Core, error)
	GetByEmail(Email string) (Core, error)
	GetMyUser(id uint) (Core, error)
}

type Service interface {
	Register(newUser Core) (Core, error)
	Login(existUser Core) (Core, string, error)
	UpdateProfile(updateData Core, c echo.Context) (Core, error)
	Deactivate(c echo.Context) (Core, error)
	ShowByEmail(Email string) (Core, error)
	MyProfile(c echo.Context) (Core, error)
	IsAuthorized(c echo.Context) error
}
