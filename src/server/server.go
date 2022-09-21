package server

import (
	indexrouter "github.com/johan9815/aprendizaje_go/src/index_router"
	"github.com/johan9815/aprendizaje_go/src/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Server() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	__dirname := utils.Path(2)

	e.Static("/public", __dirname+"/public/images")
	indexrouter.Indexrouter(e)
	e.Logger.Fatal(e.Start(":3000"))
}
