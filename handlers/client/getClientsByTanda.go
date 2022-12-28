package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/menu"
	"viandasApp/dtos"
)

func GetClientsByTandas(rw http.ResponseWriter, r *http.Request) {

	var cateogories dtos.MenuByCategoriesRequest

	err := json.NewDecoder(r.Body).Decode(&cateogories)
	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	responseMenuFood, err := db.GetMenuByCategories(cateogories.IDCategories)

	if err != nil {
		http.Error(rw, "Menu no encontrado", http.StatusBadRequest)
		return
	}

	if responseMenuFood.Menu.ID == 0 {
		http.Error(rw, "No hay menus en la BD", http.StatusNotFound)
		return
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseMenuFood)

}
