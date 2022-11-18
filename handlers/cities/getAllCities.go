package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/city"
)

func GetAllCities(rw http.ResponseWriter, r *http.Request) {

	responseModel, err := db.GetCities()

	if err != nil {
		http.Error(rw, "no fue posible recuperar las ciudades", http.StatusBadRequest)
		return
	}
	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseModel)

}
