package dtos

import (
	"viandasApp/models"

	validator "github.com/go-playground/validator/v10"
)

type UserRegister struct {
	ID        int64  `json:"id"`
	Nombre    string `json:"nombre" validate:"required"`
	Apellidos string `json:"apellidos" validate:"required"`
	Password  string `json:"password" validate:"required,min=6"`
	Email     string `json:"email" validate:"required,email"`
	/* FechaNacimiento time.Time `json:"fechaNacimiento"` */
}

//var validate *validator.Validate

func (userRegister UserRegister) ToModelUser() *models.User {

	modelUser := models.User{
		ID:        userRegister.ID,
		Nombre:    userRegister.Nombre,
		Apellidos: userRegister.Apellidos,
		Password:  userRegister.Password,
		Email:     userRegister.Email,
	}

	return &modelUser
}

func (userRegister UserRegister) Validate() error {
	validate := validator.New()
	return validate.Struct(userRegister)

}
