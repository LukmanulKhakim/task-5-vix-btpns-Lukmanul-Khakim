package main

import (
	"bookapi/config"
	dUser "bookapi/feature/user/delivery"
	rUser "bookapi/feature/user/repository"
	sUser "bookapi/feature/user/services"
	"bookapi/utils/database"
	"log"

	dBook "bookapi/feature/book/delivery"
	rBook "bookapi/feature/book/repository"
	sBook "bookapi/feature/book/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	//pemanggilan config
	cfg := config.NewConfig()
	db := database.InitDB(cfg)

	mdl := rUser.New(db)
	mbl := rBook.New(db)

	serUser := sUser.New(mdl)
	serBook := sBook.New(mbl)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	// }))

	dUser.New(e, serUser)
	dBook.New(e, serBook)

	log.Fatal(e.Start(":8000"))

}
