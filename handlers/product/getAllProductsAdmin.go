package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/product"
)

func GetAllProductsAdmin(rw http.ResponseWriter, r *http.Request) {
	responseModel, err := db.GetAllProductsAdmin()
	if err != nil {
		http.Error(rw, "no fue posible recuperar los productos", http.StatusBadRequest)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseModel)
}
