package delivery

import (
	"altafashion_be/feature/products/domain"
)

type productHandler struct {
	srv domain.Services
}

// func New(e *echo.Echo, srv domain.Services) {
// 	handler := productHandler{
// 		srv: srv,
// 	}

// 	/* Routing endpoints: */
// 	e.POST("/products", handler.AddProduct(), middleware.JWT([]byte(src.SECRET_JWT)))
// }

// func (ph *productHandler) AddProduct() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 	}
// }
