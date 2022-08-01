package routes

import (
	"errors"
	"strconv"
	"strings"

	"viandasApp/db"
	"viandasApp/models"

	"github.com/golang-jwt/jwt"
)

/* Email valor de email usado en todos los EndPoint*/
var Email string

/* IDUsuario es el ID devuelto del modelo, que se usara en todos los EndPoints*/
var IDUsuario string

/*ProcesoToken proceso token para extraer sus valores*/
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("Master")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := db.CheckExistUser(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = strconv.FormatInt(claims.ID, 10)
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err

}
