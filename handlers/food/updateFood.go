package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	db "viandasApp/db/food"
	imgdb "viandasApp/db/img"
	"viandasApp/handlers"
	"viandasApp/models"
)

/*subir el imagen comida al servidor*/
func UpdateFood(w http.ResponseWriter, r *http.Request) {

	var locationModel models.LocationImg

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

}
