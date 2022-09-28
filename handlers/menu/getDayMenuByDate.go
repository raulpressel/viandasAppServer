package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	db "viandasApp/db/menu"
	"viandasApp/dtos"
)

func GetDayMenuByDate(rw http.ResponseWriter, r *http.Request) {

	var dayDateMenuDto dtos.DayDateMenuRequest

	err := json.NewDecoder(r.Body).Decode(&dayDateMenuDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	date, err := time.Parse(time.RFC3339, dayDateMenuDto.Date)
	if err != nil {
		fmt.Println(err)
		return
	}

	responseMenuFood, err := db.GetDayMenuByDate(date)

	if err != nil {
		http.Error(rw, "Menu no encontrado", http.StatusBadRequest)
		return
	}

	var dayMenuResponse []dtos.DayMenuResponse

	for _, valor := range responseMenuFood {
		dayMenuResponse = append(dayMenuResponse, *valor.ToDayMenuDateResponse())
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(dayMenuResponse)

}
