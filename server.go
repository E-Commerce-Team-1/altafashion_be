package main

import (
	"altafashion_be/config"
	dCart "altafashion_be/feature/carts/delivery"
	rCart "altafashion_be/feature/carts/repository"
	sCart "altafashion_be/feature/carts/services"
	pDelivery "altafashion_be/feature/products/delivery"
	pRepo "altafashion_be/feature/products/repository"
	pServices "altafashion_be/feature/products/services"
	dUser "altafashion_be/feature/users/delivery"
	rUser "altafashion_be/feature/users/repository"
	sUser "altafashion_be/feature/users/services"
	"altafashion_be/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()
	db := database.InitDB(cfg)

	pRepo := pRepo.New(db)
	pServices := pServices.New(pRepo)
	mdlUser := rUser.New(db)
	mdlCart := rCart.New(db)
	serUser := sUser.New(mdlUser)
	serCart := sCart.New(mdlCart)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	// //e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	// }))

	dUser.New(e, serUser)
	pDelivery.New(e, pServices)
	dCart.New(e, serCart)

	e.Logger.Fatal(e.Start(":8000"))
}
