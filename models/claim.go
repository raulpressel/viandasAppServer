package models

import (
	"github.com/golang-jwt/jwt"
)

/* Claim es la escructura usada para procesar el JWT*/
type Claim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
