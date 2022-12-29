package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	db "viandasApp/db/menu"
)

func UpdateDayOrderAddress(rw http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("idMenu")

	if len(ID) < 1 {
		http.Error(rw, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	idMenu, err := strconv.Atoi(ID)

	if err != nil {
		http.Error(rw, "Error al convertir el ID", http.StatusInternalServerError)
		return
	}

	responseMenuFood, err := db.GetMenuById(idMenu)

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
