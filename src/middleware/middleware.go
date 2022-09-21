package jwtmiddleware

import (
	"net/http"

	"github.com/johan9815/aprendizaje_go/src/modules/users/structures"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTmiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &structures.User{},
		SigningKey: []byte("secret"),
		ErrorHandler: func(err error) error {

			if _, ok := err.(*echo.HTTPError); ok {

				if ok {

					return echo.NewHTTPError(http.StatusBadRequest, "No tienes autorizacion")
				}

			}

			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  "Token invalido o expirado",
				Internal: err,
			}
		},
	}
	return middleware.JWTWithConfig(config)
}
