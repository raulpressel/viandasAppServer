package db

import (
	user "viandasApp/db/user"
	"viandasApp/models"

	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin realiza el chequeo de login a la bd*/

func GetLogin(email string, password string) (models.User, bool) {

	user, find, _ := user.CheckExistUser(email)

	if !find {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {

		return user, false
	}

	return user, true

}
