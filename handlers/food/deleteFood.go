package handlers

import (
	"net/http"
	"strconv"
	db "viandasApp/db/food"
	"viandasApp/models"
)

/*subir el avatar al servidor*/
func DeleteFood(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("idFood")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var foodModel models.Food

	_ID, _ := strconv.Atoi(ID)

	foodModel, _ = db.GetFoodById(_ID)

	foodModel.Active = false

	status, err := db.DeleteFood(foodModel)
	if err != nil {
		http.Error(w, "No se pudo guardar el mensaje en la base de datos "+err.Error(), 400)
		return
	}

	if !status { //esto es igual a !status == false
		http.Error(w, "no se ha logrado insertar el registro  // status = false ", 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
