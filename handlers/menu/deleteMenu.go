package handlers

import (
	"net/http"
	"strconv"
	db "viandasApp/db/menu"
)

func DeleteMenu(rw http.ResponseWriter, r *http.Request) {

	idMenu := r.URL.Query().Get("idMenu")
	if len(idMenu) < 1 {
		http.Error(rw, "El parametro IDMENU es obligatorio", http.StatusBadRequest)
		return
	}

	idTurn := r.URL.Query().Get("idTurn")
	if len(idTurn) < 1 {
		http.Error(rw, "El parametro IDTURN es obligatorio", http.StatusBadRequest)
		return
	}

	_IDMenu, err := strconv.Atoi(idMenu)

	if err != nil {
		http.Error(rw, "Error al convertir el ID", http.StatusInternalServerError)
		return
	}

	_IDTurn, err := strconv.Atoi(idTurn)

	if err != nil {
		http.Error(rw, "Error al convertir el ID", http.StatusInternalServerError)
		return
	}

	//status, err := db.DeleteTurnMenu(_IDMenu, _IDTurn)

	status, err := db.DeleteMenu(_IDMenu, _IDTurn)

	if err != nil {
		http.Error(rw, "No se pudo borrar el TURN MENU de la base de datos "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "No se pudo borrar el TURN MENU de la base de datos", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusAccepted)

}
