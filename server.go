package main

import (
	"altafashion_be/config"
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

	mdlUser := rUser.New(db)

	serUser := sUser.New(mdlUser)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	// //e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	// }))

	dUser.New(e, serUser)

	e.Logger.Fatal(e.Start(":8000"))
}
