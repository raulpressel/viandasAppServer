package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/menu"
)

func GetMenu(rw http.ResponseWriter, r *http.Request) {

	responseMenuFood, err := db.GetMenuActive()

	if err != nil {
		http.Error(rw, "Menu no encontrado", http.StatusBadRequest)
		return
	}

	if responseMenuFood.ID == 0 {
		http.Error(rw, "No hay menus en la BD", http.StatusNotFound)
		return
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseMenuFood)

}
