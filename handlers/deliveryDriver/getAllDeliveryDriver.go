package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/deliveryDriver"
)

/*Obtenerbanner envia el banner al http*/

func GetAllDeliveryDriver(rw http.ResponseWriter, r *http.Request) {

	responseModelDeliveryDriver, err := db.GetAllDeliveryDriver()

	if err != nil {
		http.Error(rw, "no fue posible recuperar los platos", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseModelDeliveryDriver)

}
