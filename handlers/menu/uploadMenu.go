package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"viandasApp/db"
	dbMenu "viandasApp/db/menu"
	"viandasApp/dtos"
	"viandasApp/models"
)

/*subir el avatar al servidor*/
func UploadMenu(rw http.ResponseWriter, r *http.Request) {

	var turnDto dtos.TurnMenuRequest

	var menuModel models.Menu

	var turnMenuModel models.TurnMenu

	var dayModel []models.DayMenu

	var dModel models.DayMenu

	err := json.NewDecoder(r.Body).Decode(&turnDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	db.ExistTable(menuModel)

	/* 	first := menuDto[0].DayMenu[0]

	   	last := menuDto[len(menuDto)-1].DayMenu[len(menuDto[len(menuDto)-1].DayMenu)-1] */

	for _, menu := range turnDto.Menu {

		id, err := dbMenu.ValidateTurnId(menu.TurnId)
		if err != nil {
			http.Error(rw, "Ocurrio un error al intentar modificar el menu "+err.Error(), http.StatusInternalServerError)
			return
		}
		if id == 0 {
			http.Error(rw, "El turno no existe ", http.StatusBadRequest)
			return
		}

		turnMenuModel.TurnId = menu.TurnId

		menuModel.DateStart, err = time.Parse(time.RFC3339, menu.DateStart)
		if err != nil {
			http.Error(rw, "Error en el formato de fecha recibido "+err.Error(), http.StatusBadRequest)
			return
		}
		menuModel.DateEnd, err = time.Parse(time.RFC3339, menu.DateEnd)
		if err != nil {
			http.Error(rw, "Error en el formato de fecha recibido "+err.Error(), http.StatusBadRequest)
			return
		}

		menuModel.Active = true

		for _, day := range menu.DayMenu {
			//dModel.TurnID = menu.TurnId
			dModel.FoodID = day.Food
			dModel.Date, _ = time.Parse(time.RFC3339, day.Date)
			dayModel = append(dayModel, dModel)
		}
	}
	status, err := dbMenu.UploadMenu(dayModel, menuModel, turnMenuModel)

	if err != nil {
		http.Error(rw, "Ocurrio un error al intentar modificar el menu "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado cargar el menu en la BD", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
