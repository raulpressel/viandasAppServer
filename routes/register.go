package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"viandasApp/db"
	"viandasApp/models"
)

/*Registro es la funcion para crear en la bd el registro de usuario*/
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t) //body es un obj string de solo lectura, una vez q se utiliza body se destruye
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	_db := db.ConnectDB()

	if _db.Migrator().HasTable(t) {
		fmt.Println("ya existe la tabla", t)

	} else {

		_db.AutoMigrate(t)
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido "+err.Error(), 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Contraseña no valida < 6 "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
