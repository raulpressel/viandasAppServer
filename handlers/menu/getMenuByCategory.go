package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	db "viandasApp/db/menu"
	"viandasApp/dtos"
)

func GetMenuByCategory(rw http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("idCategory")

	if len(ID) < 1 {
		http.Error(rw, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	idCategory, _ := strconv.Atoi(ID)

	responseMenuFood := dtos.MenuViewer{}

	responseMenuFood, err := db.GetMenuByCategory(idCategory)

	if err != nil {
		http.Error(rw, "Menu no encontrado", http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseMenuFood)

}
