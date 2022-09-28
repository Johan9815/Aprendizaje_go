package userroutes

import (
	jwtmiddleware "github.com/johan9815/aprendizaje_go/src/middleware"
	"github.com/johan9815/aprendizaje_go/src/modules/users/controller"
	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	userGrpup := e.Group("/user")

	userGrpup.POST("/getUser", controller.GetUser)
	userGrpup.POST("/AddUser", controller.Add)
	userGrpup.GET("/getAlluser", controller.GetAll /*jwtmiddleware.JWTmiddleware()*/)
	userGrpup.PATCH("/updateUser", controller.UpdateUser, jwtmiddleware.JWTmiddleware())
	userGrpup.DELETE("/deleteUser", controller.DeleteUser, jwtmiddleware.JWTmiddleware())
	userGrpup.POST("/upload", controller.UploadImages, jwtmiddleware.JWTmiddleware())
	userGrpup.POST("/getToken", controller.GetToken, jwtmiddleware.JWTmiddleware())
	userGrpup.GET("/getCripto", controller.Cripto)

}
