package handlers

import (
	"encoding/json"
	"net/http"
	"viandasApp/db"
	dbCategories "viandasApp/db/categories"
	"viandasApp/dtos"
)

func UploadCategory(w http.ResponseWriter, r *http.Request) {

	var categoryDto dtos.CategoryRequest

	err := json.NewDecoder(r.Body).Decode(&categoryDto)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	category := categoryDto.ToModelCategory()

	category.Active = true

	db.ExistTable(category)

	status, err := dbCategories.UploadCategory(*category)
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
