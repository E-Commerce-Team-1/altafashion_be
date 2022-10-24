package delivery

import (
	"altafashion_be/feature/users/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := userHandler{srv: srv}

	e.POST("/register", handler.Register())
	e.POST("/login", handler.Login())
	// e.GET("/users", handler.ShowAllUser())
	// e.GET("/users/:email", handler.Profile(), middleware.JWT([]byte(key)))
	// e.PUT("/users", handler.EditProfile(), middleware.JWT([]byte(key)))
	// e.DELETE("/users", handler.DeleteUser(), middleware.JWT([]byte(key)))
}

func (us *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {

		var input RegisterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponses("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := us.srv.Register(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponses(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponses("berhasil register", ToResponse(res, "reg", "")))
	}

}

func (us *userHandler) Login() echo.HandlerFunc {
	//autentikasi user login
	return func(c echo.Context) error {
		var resQry LoginFormat
		if err := c.Bind(&resQry); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponses("cannot bind input"))
		}

		cnv := ToDomain(resQry)
		res, token, err := us.srv.Login(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponses(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponses("berhasil login", ToResponse(res, "login", token)))
	}
}
