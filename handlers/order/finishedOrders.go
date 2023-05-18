package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/order"
)

type responseFinishedOrders struct {
	Status bool  `json:"status"`
	Error  error `json:"error"`
}

func FinishedOrders(rw http.ResponseWriter, r *http.Request) {

	var response responseFinishedOrders

	response.Status, response.Error = db.FinishedOrder()

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(response)

}
