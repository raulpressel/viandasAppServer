package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/client"
	"viandasApp/dtos"
)

func GetClientsByTandas(rw http.ResponseWriter, r *http.Request) {

	var tandasDto dtos.TandasRequest

	var response *[]dtos.Client

	err := json.NewDecoder(r.Body).Decode(&tandasDto)
	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(tandasDto.IDTanda) > 0 {
		response, _ = db.GetAllClientByTandas(tandasDto.IDTanda)
	} else {
		response, _ = db.GetAllClientNotInTandas()
	}

	if err != nil {
		http.Error(rw, "Menu no encontrado", http.StatusBadRequest)
		return
	}

	/* if responseMenuFood == 0 {
		http.Error(rw, "No hay menus en la BD", http.StatusNotFound)
		return
	} */

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(response)

}
