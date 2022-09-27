package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"viandasApp/db"
	dbfood "viandasApp/db/food"
	"viandasApp/handlers"

	"viandasApp/models"
)

/*subir el avatar al servidor*/
func UploadFood(w http.ResponseWriter, r *http.Request) {

	if err := os.MkdirAll("/var/www/default/htdocs/public/food", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	var locationModel models.LocationImg

	w.Header().Add("content-type", "application/json")

	file, handle, err := r.FormFile("image")
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
		locationModel.Location = ""
	default:
		log.Println(err)
	}

	var foodModel models.Food

	foodModel.ID, _ = strconv.Atoi(r.FormValue("id"))

	foodModel.Title = r.FormValue("title")
	foodModel.Description = r.FormValue("description")
	foodModel.CategoryID, _ = strconv.Atoi(r.FormValue("category"))
	foodModel.Active = true

	db.ExistTable(foodModel)
	db.ExistTable(locationModel)

	status, err := dbfood.UploadFood(foodModel, locationModel)
	if err != nil {
		http.Error(w, "No se pudo guardar el mensaje en la base de datos "+err.Error(), 400)
		return
	}

	if !status { //esto es igual a !status == false
		http.Error(w, "no se ha logrado insertar el registro  // status = false ", 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}