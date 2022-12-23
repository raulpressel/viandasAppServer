package handlers

import (
	"net/http"
	"strconv"
	dbpathology "viandasApp/db/pathology"
)

func DeletePathology(rw http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("idPathology")

	if len(ID) < 1 {
		http.Error(rw, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	idPathology, err := strconv.Atoi(ID)

	if err != nil {
		http.Error(rw, "Error al convertir el ID", http.StatusInternalServerError)
		return
	}

	patholohyModel, err := dbpathology.GetPathologyById(idPathology)

	if err != nil {
		http.Error(rw, "Error al recuperar la patologia de la BD "+err.Error(), http.StatusInternalServerError)
		return
	}

	patholohyModel.Active = false

	rw.Header().Add("content-type", "application/json")

	status, err := dbpathology.DeletePathology(patholohyModel)

	if err != nil {
		http.Error(rw, "Ocurrio un error al eliminar la patologia "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado eliminar la patologia en la BD", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
