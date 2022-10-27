package delivery

import (
	"altafashion_be/config"
	"altafashion_be/feature/carts/domain"
	"altafashion_be/utils/jwt"
	"errors"
	"net/http"
	"strconv"

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
	e.POST("/cart", handler.AddCart(), middleware.JWT([]byte(key)))
	e.GET("/cart", handler.GetMyCart(), middleware.JWT([]byte(key)))
	e.PUT("/cart/:id", handler.UpdateQty(), middleware.JWT([]byte(key)))
	e.DELETE("/cart/:id", handler.DeleteonCart(), middleware.JWT([]byte(key)))

}

func (ch *cartHandler) AddCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input AddFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(errors.New("an invalid client request")))
		}
		input.UserID = jwt.ExtractIdToken(c)
		cnv := ToDomain(input)
		res, err := ch.srv.AddCart(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("success add cart", ToResponse(res, "add")))

	}
}

func (ch *cartHandler) GetMyCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := jwt.ExtractIdToken(c)
		res, err := ch.srv.GetMyCart(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("An invalid client request"))
		}
		return c.JSON(http.StatusOK, SuccessResponse("Success show all data", ToResponseProduct(res, "sukses")))

	}
}

func (ch *cartHandler) UpdateQty() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateFormat
		id, err := strconv.Atoi(c.Param("id"))
		input.ID = uint(id)
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(errors.New("cannot bind data")))
		}

		UserID := jwt.ExtractTokenProd(c)
		input.UserID = uint(UserID)
		cnv := ToDomain(input)
		res, err := ch.srv.UpdateQty(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("Success update cart", ToResponse(res, "update")))
	}
}

func (ch *cartHandler) DeleteonCart() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return errors.New("cannot convert id")
		}
		_, err = ch.srv.DeleteonCart(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("An invalid client request."))
		}
		return c.JSON(http.StatusOK, SuccessResponseNoData("Success delete data."))
	}
}
