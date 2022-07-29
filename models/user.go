package models

/*usuario es el modelo de usuario de la base de mysql*/

type User struct {
	ID        int64  `json:"id"`
	Nombre    string `json:"nombre"`
	Apellidos string `json:"apellidos"`
	Password  string `json:"password"`
	Email     string `json:"email"`

	/* FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"` */

}

/* func MigrateUser() {
	_db := db.ConnectDB()
	if _db.Migrator().HasTable(User{}) {
		fmt.Println("ya existe la tabla")
	} else {
		_db.AutoMigrate(User{})
	}

}
*/
