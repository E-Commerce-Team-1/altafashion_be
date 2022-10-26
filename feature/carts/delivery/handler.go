package delivery

import (
	"altafashion_be/config"
	"altafashion_be/feature/carts/domain"
	"altafashion_be/utils/jwt"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var key string

type cartHandler struct {
	srv domain.Service
}

func InitJWT(c *config.AppConfig) {
	key = c.JWSecret
}

func New(e *echo.Echo, srv domain.Service) {
	handler := cartHandler{srv: srv}
	e.POST("/cart", handler.AddCart(), middleware.JWT([]byte(key))) // TAMBAH CART

}

func (ch *cartHandler) AddCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input AddFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(errors.New("an invalid client request")))
		}
		input.IdUser = jwt.ExtractIdToken(c)
		cnv := ToDomain(input)
		res, err := ch.srv.AddCart(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("success add cart", ToResponse(res, "register")))

	}
}
