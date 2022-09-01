package models

/*usuario es el modelo de usuario de la base de mysql*/

type User struct {
	ID       int64
	Name     string
	LastName string
	Password string
	Email    string
	//FechaNacimiento time.Time `json:"fechaNacimiento"`
}
