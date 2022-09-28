package handlers

import (
	"encoding/json"
	"net/http"
	dbMenu "viandasApp/db/menu"
	"viandasApp/dtos"
)

/*subir el imagen comida al servidor*/
func UpdateMenu(w http.ResponseWriter, r *http.Request) {

	var dayMenuEdit dtos.DayMenuEditRequest

	err := json.NewDecoder(r.Body).Decode(&dayMenuEdit)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	dayMenuModel, _ := dbMenu.GetDayMenuById(dayMenuEdit.IdDayMenu)

	dayMenuModel.FoodID = dayMenuEdit.IdFood

	status, err := dbMenu.UpdateDayMenu(dayMenuModel)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el menu "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "no se ha logrado moficar el registro  ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
