package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	db "viandasApp/db/productOrder"
)

func GetProductOrdersByDate(rw http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	if dateStr == "" {
		http.Error(rw, "debe enviar el parametro date", http.StatusBadRequest)
		return
	}

	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		http.Error(rw, "formato de fecha invalido, use ISO 8601", http.StatusBadRequest)
		return
	}

	responseModel, err := db.GetProductOrdersByDate(date)
	if err != nil {
		http.Error(rw, "no fue posible recuperar las ordenes por fecha", http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(map[string]interface{}{"productOrders": responseModel})
}
