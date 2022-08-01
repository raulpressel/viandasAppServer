package jwt

import (
	"strconv"
	"time"

	"viandasApp/models"

	"github.com/golang-jwt/jwt"
)

/* GeneroJWT genera el enciptado con JWT */

func GeneroJWT(t models.User) (string, error) {
	miClave := []byte("Master")

	payload := jwt.MapClaims{
		"email":     t.Email,
		"nombre":    t.Nombre,
		"apellidos": t.Apellidos,
		/* "fecha_nacimiento": t.FechaNacimiento, */
		"_id": strconv.FormatInt(t.ID, 10),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil

}
