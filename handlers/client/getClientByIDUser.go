package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/client"
	"viandasApp/handlers"
)

func GetClientByIDUser(rw http.ResponseWriter, r *http.Request) {

	idUserKL := r.URL.Query().Get("idUser")

	if len(idUserKL) < 1 {
		http.Error(rw, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	usr := handlers.GetUser()

	if usr.ID != idUserKL {
		if !usr.Admin {
			http.Error(rw, "No tienes los permisos para ver esta información", http.StatusBadRequest)
			return
		}
	}

	responseClient, err := db.GetClientByIDUser(idUserKL)

	if err != nil {
		http.Error(rw, "Cliente no encontrado", http.StatusBadRequest)
		return
	}

	/* if responseClient.Client.ID == 0 {
		http.Error(rw, "No hay clientes en la BD", http.StatusNotFound)
		return
	} */

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseClient)

}
