package indexrouter

import (
	userroutes "github.com/johan9815/aprendizaje_go/src/modules/users/routes"
	"github.com/labstack/echo/v4"
)

func Indexrouter(e *echo.Echo) {
	userroutes.UserRoutes(e)
}
