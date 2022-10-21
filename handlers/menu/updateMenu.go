package handlers

import (
	"encoding/json"
	"net/http"
	dbFood "viandasApp/db/food"
	dbMenu "viandasApp/db/menu"
	"viandasApp/dtos"
)

/*subir el imagen comida al servidor*/
func UpdateMenu(rw http.ResponseWriter, r *http.Request) {

	var dayMenuEdit dtos.DayMenuEditRequest

	err := json.NewDecoder(r.Body).Decode(&dayMenuEdit)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	dayMenuModel, err := dbMenu.GetDayMenuById(dayMenuEdit.IdDayMenu)

	if err != nil {
		http.Error(rw, "Ocurrio un error al obtener el ID del menu "+err.Error(), http.StatusInternalServerError)
		return
	}

	_idFoodCategory, err := dbFood.GetIdFoodCategory(dayMenuEdit.IdFood, dayMenuEdit.IdCategory)
	if _idFoodCategory < 1 {
		http.Error(rw, "El plato con no existe "+err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(rw, "no se ha recuperar el plato de la BD", http.StatusInternalServerError)
		return
	}

	dayMenuModel.FoodCategoryID = _idFoodCategory

	status, err := dbMenu.UpdateDayMenu(dayMenuModel)

	if err != nil {
		http.Error(rw, "Ocurrio un error al intentar modificar el menu "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado moficar el registro  ", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusCreated)

}
