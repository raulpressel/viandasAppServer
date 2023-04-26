package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	db "viandasApp/db/menu"
	"viandasApp/dtos"
)

func GetMenuViewer(rw http.ResponseWriter, r *http.Request) {

	var filterMenu dtos.FilterDateMenu

	err := json.NewDecoder(r.Body).Decode(&filterMenu)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	dateStart, err := time.Parse(time.RFC3339, filterMenu.DateStart)
	if err != nil {
		http.Error(rw, "Error en el formato de fecha recibido "+err.Error(), http.StatusBadRequest)
		return
	}

	dateEnd, err := time.Parse(time.RFC3339, filterMenu.DateEnd)
	if err != nil {
		http.Error(rw, "Error en el formato de fecha recibido "+err.Error(), http.StatusBadRequest)
		return
	}

	responseMenuFood, err := db.GetMenuActive(dateStart, dateEnd)

	if err != nil {
		http.Error(rw, "Error al recuperar el Menu del servidor", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseMenuFood)

}
