package models

/*usuario es el modelo de usuario de la base de mysql*/

type User struct {
	ID        int64  `json:"id"`
	Nombre    string `json:"nombre"`
	Apellidos string `json:"apellidos"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	/* FechaNacimiento time.Time `json:"fechaNacimiento"` */
}
