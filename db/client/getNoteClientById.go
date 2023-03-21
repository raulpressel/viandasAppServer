package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetNoteClientById(id int) (models.ClientNotes, error) {

	db := db.GetDB()

	var clientNotesModel models.ClientNotes

	err := db.First(&clientNotesModel, id).Error

	return clientNotesModel, err

}
