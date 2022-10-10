package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/menu"
)

func GetAllMenu(rw http.ResponseWriter, r *http.Request) {

	responseAllMenu, err := db.GetAllMenu()

	if err != nil {
		http.Error(rw, "Error al recuperar Menu de la BD", http.StatusInternalServerError)
		return
	}

	if responseAllMenu == nil {
		http.Error(rw, "No hay menus en la BD", http.StatusNotFound)
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseAllMenu)

}
