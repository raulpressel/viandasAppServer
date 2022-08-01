package db

import (
	"viandasApp/models"

	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin realiza el chequeo de login a la bd*/

func TryLogin(email string, password string) (models.User, bool) {

	usu, encontrado, _ := CheckExistUser(email)

	if !encontrado {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {

		return usu, false
	}

	return usu, true

}
