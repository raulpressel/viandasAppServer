package jwt

import (
	"strconv"
	"time"

	"viandasApp/models"

	"github.com/golang-jwt/jwt"
)

/* GeneroJWT genera el enciptado con JWT */

func GenerateJWT(user models.User) (string, error) {
	miClave := []byte("Master")

	payload := jwt.MapClaims{
		"email":     user.Email,
		"nombre":    user.Name,
		"apellidos": user.LastName,
		/* "fecha_nacimiento": t.FechaNacimiento, */
		"_id": strconv.FormatInt(user.ID, 10),
		"exp": time.Now().Add(time.Hour * 168).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil

}
