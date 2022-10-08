package handlers

import (
	"net/http"
	"strconv"
	db "viandasApp/db/menu"
)

func DeleteMenu(w http.ResponseWriter, r *http.Request) {

	idMenu := r.URL.Query().Get("idMenu")
	if len(idMenu) < 1 {
		http.Error(w, "El parametro IDMENU es obligatorio", http.StatusBadRequest)
		return
	}

	idTurn := r.URL.Query().Get("idTurn")
	if len(idTurn) < 1 {
		http.Error(w, "El parametro IDTURN es obligatorio", http.StatusBadRequest)
		return
	}

	_IDMenu, _ := strconv.Atoi(idMenu)

	_IDTurn, _ := strconv.Atoi(idTurn)

	status, err := db.DeleteTurnMenu(_IDMenu, _IDTurn)

	if err != nil {
		http.Error(w, "No se pudo borrar el TURN MENU de la base de datos "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(w, "No se pudo borrar el TURN MENU de la base de datos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)

}
