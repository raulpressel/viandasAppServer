package handlers

import (
	"net/http"
	"strconv"
	db "viandasApp/db/food"
	"viandasApp/models"
)

func DeleteFood(rw http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("idFood")
	if len(ID) < 1 {
		http.Error(rw, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var foodModel models.Food

	_ID, err := strconv.Atoi(ID)
	if err != nil {
		http.Error(rw, "Error al convertir el ID", http.StatusInternalServerError)
		return
	}

	foodModel, err = db.GetFoodById(_ID)
	if err != nil {
		http.Error(rw, "No se pudo recuperar el plato de la base de datos "+err.Error(), http.StatusInternalServerError)
		return
	}

	foodModel.Active = false

	status, err := db.DeleteFood(foodModel)
	if err != nil {
		http.Error(rw, "No se pudo borrar el plato de la base de datos "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "No se pudo borrar el plato de la base de datos ", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
