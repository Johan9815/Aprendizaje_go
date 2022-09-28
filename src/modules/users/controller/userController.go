package controller

import (
	"fmt"
	"net/http"

	"github.com/johan9815/aprendizaje_go/src/modules/users/models"

	. "github.com/johan9815/aprendizaje_go/src/modules/users/structures"
	"github.com/johan9815/aprendizaje_go/src/utils"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {

	u := new(User)

	if err := c.Bind(u); err != nil {
		return err
	}

	list, err := models.GetUser(u.User, u.Pass)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "fail", "message": "Ha ocurrido un error en el servidor"})

	}

	if list.Id == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": "El usuario no se encuentra registrado"})

	}

	t := utils.GetKey(list)

	return c.JSON(http.StatusOK, echo.Map{
		"token":   t,
		"status":  "OK",
		"message": "Usuario encontrado con exito",
		"data":    list,
	})

}

func GetAll(c echo.Context) error {

	u := new(User)

	if err := c.Bind(u); err != nil {
		return err
	}

	list, err := models.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "fail", "message": "Ha ocurrido un error en el servidor"})

	}

	if len(list) == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": "No hay usuarios registrados"})

	}

	return c.JSON(http.StatusOK, echo.Map{

		"status":  "OK",
		"message": "Usuarios",
		"data":    list,
	})

}

func Add(c echo.Context) error {

	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	user, err := models.Add(*u)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"status": "OK", "message": "Usuario creado con exito", "data": user})

}

func UpdateUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	user, err := models.UpdateUser(*u)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "fail", "message": "Ha ocurrido un error en el servidor"})
	}

	return c.JSON(http.StatusOK, echo.Map{"status": "OK", "message": "Usuario actualizado con exito", "data": user})

}

func DeleteUser(c echo.Context) error {

	u := new(User)

	if err := c.Bind(u); err != nil {
		return err
	}

	id, err := models.DeleteUser(*u)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "fail", "message": "Ha ocurrido un error en el servidor"})
	}

	return c.JSON(http.StatusOK, echo.Map{"status": "OK", "message": "Usuario eliminado con exito", "data": id})

}

func UploadImages(c echo.Context) error {

	name := c.FormValue("name")
	email := c.FormValue("email")

	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return err
	}

	filename := utils.UploadImage(file)
	//-----------
	// Read file
	//-----------

	// Source

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s and email=%s.</p>", filename, name, email))
}

func GetToken(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*User)
	//name := claims.Conocimientos
	return c.JSON(http.StatusOK, echo.Map{
		"user": claims.User,
	})
}

func Cripto(c echo.Context) error {

	ch := make(chan interface{}, 1)

	go utils.EtheriumApi(ch, "0x52f1098db91257030bf10fd848e2da9974a3e2b3e021261d5a3da7146dd4b0ab")

	resp := <-ch
	close(ch)

	return c.JSON(http.StatusOK, echo.Map{
		"data": resp,
	})

}
