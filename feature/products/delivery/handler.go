package delivery

import (
	"altafashion_be/config"
	"altafashion_be/feature/products/domain"
	"altafashion_be/utils/jwt"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"altafashion_be/utils/aws"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var key string

func InitJWT(c *config.AppConfig) {
	key = c.JWSecret
}

type productHandler struct {
	srv domain.Services
}

func New(e *echo.Echo, srv domain.Services) {
	handler := productHandler{
		srv: srv,
	}

	/* Routing endpoints: */
	e.GET("/products", handler.GetAll())
	e.GET("/products/:id", handler.GetDetail())
	e.POST("/products", handler.AddProduct(), middleware.JWT([]byte(key)))
	e.PUT("/products/:id", handler.EditProduct(), middleware.JWT([]byte(key)))
	e.DELETE("/products/:id", handler.Destroy(), middleware.JWT([]byte(key)))
}

func (ph *productHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil {
			page = 0
		}
		category := c.QueryParam("category")
		name := c.QueryParam("name")

		res, err := ph.srv.GetAll(category, name, page)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailedResponse(config.DATA_NOT_FOUND))
		}

		return c.JSON(http.StatusOK, SuccessResponse("Success get all product", ToResponseList(res)))
	}
}

func (ph *productHandler) AddProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input AddProductFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailedResponse(errors.New("an invalid client request")))
		}
		input.UserID = jwt.ExtractTokenProd(c)
		file, _ := c.FormFile("image")
		if file != nil {
			res, err := aws.UploadProfileProduct(c)
			if err != nil {
				return err
			}
			log.Print(res)
			input.Image = res
		}
		cnv := ToDomain(input)
		res, err := ph.srv.AddProduct(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailedResponse("There is problem on server."))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("Success add product", ToResponse(res, "add")))
	}
}

func (ph *productHandler) GetDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		cnv, errCnv := strconv.Atoi(c.Param("id"))
		if errCnv != nil {
			return c.JSON(http.StatusInternalServerError, FailedResponse("cant convert id"))
		}

		res, err := ph.srv.GetDetail(uint(cnv))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailedResponse(config.DATA_NOT_FOUND))
		}

		return c.JSON(http.StatusOK, SuccessResponse("Success get detail product", ToResponse(res, "edit")))
	}
}

func (ph *productHandler) EditProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input EditProductFormat
		id, _ := strconv.Atoi(c.Param("id"))
		input.ID = uint(id)
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailedResponse(errors.New("cant bind data")))
		}

		file, _ := c.FormFile("image")
		if file != nil {
			res, err := aws.UploadProfileProduct(c)
			if err != nil {
				return err
			}
			log.Print(res)
			input.Image = res
		}

		UserID := jwt.ExtractTokenProd(c)
		input.UserID = uint(UserID)
		cnv := ToDomain(input)
		res, err := ph.srv.EditProduct(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailedResponse("There is problem on server."))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("Success update product", ToResponse(res, "edit")))
	}
}

func (ph *productHandler) Destroy() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID, _ := strconv.Atoi(c.Param("id"))

		err := ph.srv.Destroy(uint(ID))
		if err != nil {
			if strings.Contains(err.Error(), "table") {
				return c.JSON(http.StatusInternalServerError, FailedResponse(err.Error()))
			} else if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusInternalServerError, FailedResponse(err.Error()))
			}
		}

		return c.JSON(http.StatusOK, SuccessNoDataResponse("Success delete product."))
	}
}
