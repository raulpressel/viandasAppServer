package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"viandasApp/db"
	dbfood "viandasApp/db/food"
	"viandasApp/handlers"

	"viandasApp/models"
)

/*subir el avatar al servidor*/
func UploadFood(rw http.ResponseWriter, r *http.Request) {

	if err := os.MkdirAll("/var/www/default/htdocs/public/food", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	var locationModel models.LocationImg

	rw.Header().Add("content-type", "application/json")

	file, handle, err := r.FormFile("image")
	switch err {
	case nil:
		locationModel.Location = "/var/www/default/htdocs/public/food/" + handle.Filename

		f, err := os.OpenFile(locationModel.Location, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			http.Error(rw, "error al subir imagen comida "+err.Error(), http.StatusBadRequest)
			return
		}

		_, err = io.Copy(f, file)

		if err != nil {
			http.Error(rw, "error al copiar  imagen comida "+err.Error(), http.StatusBadRequest)
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

	/* 	foodModel.ID, err = strconv.Atoi(r.FormValue("id"))

	   	if err != nil {
	   		http.Error(rw, "Error al convertir el ID", http.StatusInternalServerError)
	   		return
	   	} */

	foodModel.Title = r.FormValue("title")
	foodModel.Description = r.FormValue("description")
	foodModel.Active = true

	cat := (r.FormValue("categories"))

	categoryArray := strings.Split(cat, ", ")

	db.ExistTable(foodModel)
	db.ExistTable(locationModel)

	var foodCategoryModel models.FoodCategory

	var foodCategoriesModel []models.FoodCategory

	db.ExistTable(foodCategoryModel)

	for _, value := range categoryArray {
		foodCategoryModel.CategoryID, _ = strconv.Atoi(value)
		foodCategoryModel.FoodID = foodModel.ID
		foodCategoriesModel = append(foodCategoriesModel, foodCategoryModel)
	}

	status, err := dbfood.UploadFood(foodModel, locationModel, foodCategoriesModel)
	if err != nil {
		http.Error(rw, "No se pudo guardar el mensaje en la base de datos "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado insertar el registro  // status = false ", 400)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
