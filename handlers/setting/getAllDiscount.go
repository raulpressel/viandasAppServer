package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/setting"
)

func GetAllDiscount(rw http.ResponseWriter, r *http.Request) {

	responseModel, err := db.GetAllDiscount()

	if err != nil {
		http.Error(rw, "no fue posible recuperar los descuentos", http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseModel)

}
