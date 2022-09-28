package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	dbMenu "viandasApp/db/menu"
	"viandasApp/dtos"
)

/*subir el imagen comida al servidor*/
func UpdateMenu(w http.ResponseWriter, r *http.Request) {

	var dayMenuEdit dtos.DayMenuEditRequest

	err := json.NewDecoder(r.Body).Decode(&dayMenuEdit)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	Date, _ := time.Parse(time.RFC3339, dayMenuEdit.Date)

	dayMenu, err := dbMenu.GetDayMenuByDate(Date)

	if err != nil {
		http.Error(w, "No se pudo obtener los platos en la fecha solicitada "+err.Error(), 400)
		return
	}

	for _, valor := range dayMenu {

		dayMenuModel, _ := dbMenu.GetDayMenuById(valor.ID)

		if valor.Categoryid == dayMenuEdit.Category {
			dayMenuModel.FoodID = dayMenuEdit.IdFood
			status, err := dbMenu.UpdateDayMenu(dayMenuModel)

			if err != nil {
				http.Error(w, "Ocurrio un error al intentar modificar el menu "+err.Error(), 400)
				return
			}

			if !status {
				http.Error(w, "no se ha logrado moficar el registro  ", 400)
				return
			}

		}
	}

	w.WriteHeader(http.StatusCreated)

	/* 	var locationModel models.LocationImg

	   	var foodModel models.Food

	   	foodModel.ID, _ = strconv.Atoi(r.FormValue("id"))

	   	w.Header().Add("content-type", "application/json")

	   	file, handle, err := r.FormFile("image")

	   	foodModel, _ = db.GetFoodById(foodModel.ID)

	   	switch err {
	   	case nil:
	   		locationModel.Location = "/var/www/default/htdocs/public/food/" + handle.Filename

	   		f, err := os.OpenFile(locationModel.Location, os.O_WRONLY|os.O_CREATE, 0666)
	   		if err != nil {
	   			http.Error(w, "error al subir imagen comida "+err.Error(), http.StatusBadRequest)
	   			return
	   		}

	   		_, err = io.Copy(f, file)

	   		if err != nil {
	   			http.Error(w, "error al copiar  imagen comida "+err.Error(), http.StatusBadRequest)
	   			return
	   		}

	   		locationModel.Location = handlers.GetHash(locationModel.Location)
	   		file.Close()
	   	case http.ErrMissingFile:
	   		locationModel, _ = imgdb.GetLocationImgById(foodModel.LocationID)
	   	default:
	   		log.Println(err)
	   	}

	   	foodModel.Title = r.FormValue("title")
	   	foodModel.Description = r.FormValue("description")
	   	foodModel.CategoryID, _ = strconv.Atoi(r.FormValue("category"))
	   	foodModel.Active = true

	   	locationModel.ID = foodModel.LocationID

	   	status, err := db.UpdateFood(foodModel, locationModel)
	   	if err != nil {
	   		http.Error(w, "No se pudo guardar el mensaje en la base de datos "+err.Error(), 400)
	   		return
	   	}

	   	if !status {
	   		http.Error(w, "no se ha logrado insertar el registro  // status = false ", 400)
	   		return
	   	}

	   	w.Header().Set("Content-type", "application/json")
	   	w.WriteHeader(http.StatusCreated)
	*/
}
