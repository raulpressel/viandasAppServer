package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/client"
)

func GetClientByIDUser(rw http.ResponseWriter, r *http.Request) {

	idUserKL := r.URL.Query().Get("idUser")

	if len(idUserKL) < 1 {
		http.Error(rw, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	responseClient, err := db.GetClientByIDUser(idUserKL)

	if err != nil {
		http.Error(rw, "Cliente no encontrado", http.StatusBadRequest)
		return
	}

	if responseClient.ID == 0 {
		http.Error(rw, "No hay clientes en la BD", http.StatusNotFound)
		return
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseClient)

}
