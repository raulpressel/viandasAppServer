package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

type Rol struct {
	Roles []string
}

var secretKey *[]byte

//var claims, _, _ = ProcessToken()

//var claims &jwt.MapClaims{}

//var Route string

func GetCert(key string) (string, error) {

	keyCert := os.Getenv(key)
	if keyCert == "" {
		return keyCert, errors.New("missing key")
	}
	return keyCert, nil
}

func GetSecretKey(route string) (*[]byte, error) {
	secretK, err := os.ReadFile(route)
	if err != nil {
		return &secretK, err
	}
	secretKey = &secretK
	return secretKey, err
}

func GetSK() *[]byte {
	return secretKey
}

type user struct {
	Email    string
	ID       string
	Name     string
	LastName string
}

func GetUser() user {
	return usr
}

var usr user

/*ProcesoToken proceso token para extraer sus valores*/
func ProcessToken(tk string) (*jwt.MapClaims, bool, error) {

	secretk := GetSK()

	var admin bool

	claims := &jwt.MapClaims{}

	splitToken := strings.Replace(tk, "Bearer ", "", -1)

	key, er := jwt.ParseRSAPublicKeyFromPEM([]byte(*secretk))
	if er != nil {
		return claims, false, er
	}

	token, err := jwt.ParseWithClaims(splitToken, claims, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {

		return claims, admin, err
	}

	if !token.Valid {
		return claims, admin, err
	}

	for key, val := range *claims {
		if key == "sub" {
			usr.ID = fmt.Sprintf("%v", val)

		}

		if key == "email" {
			usr.Email = fmt.Sprintf("%v", val)
		}

		if key == "given_name" {
			usr.Name = fmt.Sprintf("%v", val)
		}

		if key == "family_name" {
			usr.LastName = fmt.Sprintf("%v", val)
		}

		if key == "realm_access" {

			b, err := json.Marshal(val)
			if err != nil {
				return claims, false, err
			}

			var rol Rol

			if err := json.Unmarshal(b, &rol); err != nil {
				fmt.Println(err)
			}

			for _, v := range rol.Roles {
				if v == "admin" {
					admin = true
				}
			}

		}

	}

	return claims, admin, err

}
