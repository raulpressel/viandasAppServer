package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/food"
)

/*Obtenerbanner envia el banner al http*/

func GetAllFood(rw http.ResponseWriter, r *http.Request) {

	responseModelFood, err := db.GetAllFood()

	if err != nil {
		http.Error(rw, "no fue posible recuperar los platos", http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseModelFood)

}
