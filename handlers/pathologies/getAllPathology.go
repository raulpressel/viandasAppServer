package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/pathology"
)

/*Obtenerbanner envia el banner al http*/

func GetAllPathology(rw http.ResponseWriter, r *http.Request) {

	responseModel, err := db.GetAllPathology()

	if err != nil {
		http.Error(rw, "no fue posible recuperar las patologias", http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseModel)

}
