package structures

import "github.com/golang-jwt/jwt"

type User struct {
	Id   int    `json:"id"`
	User string `json:"user"`
	Pass string `json:"pass"`
	jwt.StandardClaims
}

type Saberes struct {
	Id         int    `json:"id"`
	Id_usuario int    `json:"id_usuario"`
	Lenguaje   string `json:"lenguaje"`
	Exp        int    `json:"exp"`
}
