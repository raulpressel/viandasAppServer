package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	db "viandasApp/db/menu"
	"viandasApp/dtos"
)

func GetDayMenuByDate(rw http.ResponseWriter, r *http.Request) {

	var dayDateMenuDto dtos.DayDateMenuRequest

	err := json.NewDecoder(r.Body).Decode(&dayDateMenuDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	date, err := time.Parse(time.RFC3339, dayDateMenuDto.Date)
	if err != nil {
		http.Error(rw, "Error en el formato de fecha recibido "+err.Error(), http.StatusBadRequest)
		return
	}

	responseMenuFood, err := db.GetDayMenuByDate(date)

	if err != nil {
		http.Error(rw, "Menu no encontrado", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseMenuFood)

}
