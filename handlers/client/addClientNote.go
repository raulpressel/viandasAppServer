package handlers

import (
	"encoding/json"
	"net/http"
	dbClient "viandasApp/db/client"
	"viandasApp/dtos"
	"viandasApp/models"
)

func AddClientNote(rw http.ResponseWriter, r *http.Request) {

	var addNote dtos.AddNoteRequest

	var notesClientModel models.ClientNotes

	err := json.NewDecoder(r.Body).Decode(&addNote)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	notesClientModel.ClientID = addNote.IDClient

	notesClientModel.Note = addNote.Note.Note

	status, err := dbClient.AddClientNote(notesClientModel)

	if err != nil {
		http.Error(rw, "Ocurrio un error al cargar la nota al cliente "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no fue posible cargar la nota al cliente en la BD", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
