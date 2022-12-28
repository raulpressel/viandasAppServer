package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/client"
)

func GetAllClient(rw http.ResponseWriter, r *http.Request) {

	responseModelFood, err := db.GetAllClient()

	if err != nil {
		http.Error(rw, "no fue posible recuperar los clientes", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseModelFood)

}
