package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"viandasApp/db"
	"viandasApp/dtos"
)

/*Registro es la funcion para crear en la bd el registro de usuario*/
func Register(w http.ResponseWriter, r *http.Request) {
	var userDto dtos.UserRegister

	err := json.NewDecoder(r.Body).Decode(&userDto) //body es un obj string de solo lectura, una vez q se utiliza body se destruye
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if userDto.Validate() != nil {
		fmt.Println("algo falta amigo") //se puede mejorar mucho esta funcion ver DESP que es lo que se va a validar realmente
		return
	}

	user := userDto.ToModelUser()

	db.ExistTable(user)

	_, find, _ := db.CheckExistUser(user.Email)
	if find {
		http.Error(w, "Email ya registrado ", 400)
		return
	}

	status, err := db.InsertRegistry(*user)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}

	if !status { //esto es igual a !status == false
		http.Error(w, "no se ha logrado insertar el registro  // status = false ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
