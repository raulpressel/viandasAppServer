package handlers

import (
	"encoding/json"
	"net/http"
	dbClient "viandasApp/db/client"
	"viandasApp/dtos"
)

func EditClientNote(rw http.ResponseWriter, r *http.Request) {

	var editNote dtos.EditNoteRequest

	err := json.NewDecoder(r.Body).Decode(&editNote)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	notesClientModel, err := dbClient.GetNoteClientById(editNote.Note.ID)

	if err != nil {
		http.Error(rw, "Nota de cliente no encontrada", http.StatusBadRequest)
		return
	}

	notesClientModel.Note = editNote.Note.Note

	status, err := dbClient.EditClientNote(*notesClientModel)

	if err != nil {
		http.Error(rw, "Ocurrio un error al editar la nota al cliente "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no fue posible editar la nota al cliente en la BD", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
