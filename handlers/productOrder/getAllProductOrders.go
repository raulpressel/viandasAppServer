package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/productOrder"
)

func GetAllProductOrders(rw http.ResponseWriter, r *http.Request) {
	responseModel, err := db.GetAllProductOrders()
	if err != nil {
		http.Error(rw, "no fue posible recuperar las ordenes de productos", http.StatusBadRequest)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(map[string]interface{}{"productOrders": responseModel})
}
