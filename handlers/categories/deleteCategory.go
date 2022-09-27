package handlers

import (
	"encoding/json"
	"net/http"
	dbCategories "viandasApp/db/categories"
	"viandasApp/dtos"
	"viandasApp/models"
)

/*subir el avatar al servidor*/
func DeelteCategory(w http.ResponseWriter, r *http.Request) {

	var categoryDto dtos.CategoryDeleteRequest

	err := json.NewDecoder(r.Body).Decode(&categoryDto)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	var categoryModel models.Category

	categoryModel, _ = dbCategories.GetCategoryById(categoryDto.ID)

	categoryModel.Active = false

	w.Header().Add("content-type", "application/json")

	status, err := dbCategories.DeleteCategory(categoryModel)

	if err != nil {
		http.Error(w, "No se pudo guardar el mensaje en la base de datos "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "no se ha logrado insertar el registro  // status = false ", 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
