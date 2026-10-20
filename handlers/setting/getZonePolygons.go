package handlers

import (
	"encoding/json"
	"net/http"
	"viandasApp/dtos"
	"viandasApp/handlers"
)

func GetZonePolygons(rw http.ResponseWriter, r *http.Request) {

	response := dtos.ZonePolygonResponse{
		Zones: handlers.ZonePolygons,
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(response)

}
