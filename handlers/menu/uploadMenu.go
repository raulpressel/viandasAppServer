package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"viandasApp/db"
	dbMenu "viandasApp/db/menu"
	"viandasApp/dtos"
	"viandasApp/models"
)

/*subir el avatar al servidor*/
func UploadMenu(w http.ResponseWriter, r *http.Request) {

	var turnDto dtos.TurnMenuRequest

	var menuModel models.Menu

	var dayModel []models.DayMenu

	var dModel models.DayMenu

	err := json.NewDecoder(r.Body).Decode(&turnDto)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	db.ExistTable(menuModel)

	/* 	first := menuDto[0].DayMenu[0]

	   	last := menuDto[len(menuDto)-1].DayMenu[len(menuDto[len(menuDto)-1].DayMenu)-1] */

	for _, menu := range turnDto.Menu {
		menuModel.TurnId = menu.TurnId

		menuModel.DateStart, err = time.Parse(time.RFC3339, menu.DateStart)
		if err != nil {
			fmt.Println(err)
			return
		}
		menuModel.DateEnd, err = time.Parse(time.RFC3339, menu.DateEnd)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, day := range menu.DayMenu {
			dModel.MenuID = menuModel.ID
			dModel.FoodID = day.Food
			dModel.Date, _ = time.Parse(time.RFC3339, day.Date)
			dayModel = append(dayModel, dModel)
		}
	}
	dbMenu.UploadMenu(dayModel, menuModel)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
