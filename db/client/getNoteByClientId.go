package db

import (
	"errors"
	"viandasApp/db"
	"viandasApp/models"

	"gorm.io/gorm"
)

func GetNoteByClientId(idClient int) (*models.ClientNotes, error) {

	db := db.GetDB()

	var clientNotesModel models.ClientNotes

	err := db.First(&clientNotesModel, "client_id = ?", idClient).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		clientNotesModel.ClientID = idClient
		AddClientNote(clientNotesModel)
	}

	return &clientNotesModel, nil

}
