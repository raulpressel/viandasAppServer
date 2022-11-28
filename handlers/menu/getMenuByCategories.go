package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/menu"
	"viandasApp/dtos"
)

func GetMenuByCategories(rw http.ResponseWriter, r *http.Request) {

	/* ID := r.URL.Query().Get("idCategory")

	if len(ID) < 1 {
		http.Error(rw, "debe enviar el parametro id", http.StatusBadRequest)
		return
	} */

	/* //idCategory, err := strconv.Atoi(ID)

	if err != nil {
		http.Error(rw, "Error al convertir el ID", http.StatusInternalServerError)
		return
	} */

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

	/* if responseMenuFood.ID == 0 {
		http.Error(rw, "No hay menus en la BD", http.StatusNotFound)
		return
	} */

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseMenuFood)

}
