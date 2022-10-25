package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

/* Email valor de email usado en todos los EndPoint*/
//var Email string

type Rol struct {
	Roles []string
}

/* IDUsuario es el ID devuelto del modelo, que se usara en todos los EndPoints*/
//var IDUsuario string

/*ProcesoToken proceso token para extraer sus valores*/
func ProcessToken(tk string) (*jwt.MapClaims, bool, error) {

	var SecretKey, _ = os.ReadFile("C:/Users/Rulo/github.com/raulpressel/viandasAppServer/cert.pem")

	//reqToken := strings.Split((c.Header.Get("Authorization")), "Bearer")

	//reqToken := c.Header.Get("Authorization")

	var admin bool

	claims := &jwt.MapClaims{}

	splitToken := strings.Replace(tk, "Bearer ", "", -1)

	key, er := jwt.ParseRSAPublicKeyFromPEM([]byte(SecretKey))
	if er != nil {
		fmt.Println(er)
		/* c.Abort()
		   c.Writer.WriteHeader(http.StatusUnauthorized)
		   c.Writer.Write([]byte("Unauthorized")) */
		return claims, false, errors.New("formato de token invalido")
	}

	token, err := jwt.ParseWithClaims(splitToken, claims, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		fmt.Println(err)
		/* c.Abort()
		   c.Writer.WriteHeader(http.StatusUnauthorized)
		   c.Writer.Write([]byte("Unauthorized"))
		*/return claims, admin, errors.New("formato de token invalido")
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("token is valid")
	}

	for key, val := range *claims {
		if key == "sub" {
			fmt.Printf("Key: %v, value: %v\n", key, val)
		}
		if key == "realm_access" {

			fmt.Printf("Key: %v, value: %v\n", key, val)

			b, err := json.Marshal(val)
			if err != nil {
				log.Fatal(err)
			}
			test := string(b)

			fmt.Println(test)

			var r Rol

			if err := json.Unmarshal(b, &r); err != nil {
				fmt.Println(err)
			}

			for _, v := range r.Roles {
				if v == "admin" {
					admin = true
				}
			}

		}

	}

	/* 	if err == nil {
		_, encontrado, _ := db.CheckExistUser(claims.)
		if encontrado {
			Email = claims.Email
			//IDUsuario = strconv.FormatInt(claims.ID, 10)
		}
		return claims, encontrado, nil
	} */

	/* 	miClave := []byte("Master")
	   	claims := &models.Claim{}

	   	splitToken := strings.Split(tk, "Bearer")

	   	if len(splitToken) != 2 {
	   		return claims, false, errors.New("formato de token invalido")
	   	}
	   	tk = strings.TrimSpace(splitToken[1])

	   	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
	   		return miClave, nil
	   	})
	   	if err == nil {
	   		_, encontrado, _ := db.CheckExistUser(claims.Email)
	   		if encontrado {
	   			Email = claims.Email
	   			//IDUsuario = strconv.FormatInt(claims.ID, 10)
	   		}
	   		return claims, encontrado, nil
	   	}
	   	if !tkn.Valid {
	   		return claims, false, errors.New("token invalido")
	   	}
	*/
	return claims, admin, err

}
