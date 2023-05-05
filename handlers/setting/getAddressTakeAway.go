package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	db "viandasApp/db/address"
)

func GetAddressTakeAway(rw http.ResponseWriter, r *http.Request) {

	responseModel, err := db.GetAddressTakeAway()

	if err != nil && !strings.Contains(err.Error(), "record not found") {
		http.Error(rw, "no fue posible recuperar la dirección de take away", http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseModel)

}
