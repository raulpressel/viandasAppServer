package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/productCategory"
)

func GetAllProductCategories(rw http.ResponseWriter, r *http.Request) {
	responseModel, err := db.GetAllProductCategories()
	if err != nil {
		http.Error(rw, "no fue posible recuperar las categorias de productos", http.StatusBadRequest)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(map[string]interface{}{"productCategories": responseModel})
}
