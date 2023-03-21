package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetNoteByClientId(idClient int) (*models.ClientNotes, error) {

	db := db.GetDB()

	var clientNotesModel models.ClientNotes

	err := db.First(&clientNotesModel, "client_id = ?", idClient).Error

	return &clientNotesModel, err

}
