package utils

import (
	"time"

	"github.com/johan9815/aprendizaje_go/src/modules/users/structures"

	"github.com/golang-jwt/jwt"
)

func GetKey(list *structures.User) string {
	claims := &structures.User{
		Id:   list.Id,
		User: list.User,
		Pass: list.Pass,
		//Conocimientos: list.Conocimientos,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err.Error()
	}

	return t
}
