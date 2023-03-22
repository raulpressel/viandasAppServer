package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func AddClientNote(noteClientModel models.ClientNotes) (bool, error) {

	db := db.GetDB()

	err := db.Save(&noteClientModel)

	if err.Error != nil {
		return false, err.Error
	}
	return true, err.Error
}
