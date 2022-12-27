package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/tanda"
)

func GetAllTanda(rw http.ResponseWriter, r *http.Request) {

	responseModel, err := db.GetAllTanda()

	if err != nil {
		http.Error(rw, "no fue posible recuperar las tandas", http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseModel)

}
