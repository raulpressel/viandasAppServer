package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	db "viandasApp/db/food"
	imgdb "viandasApp/db/img"
	"viandasApp/handlers"
	"viandasApp/models"
)

/*subir el imagen comida al servidor*/
func UpdateFood(rw http.ResponseWriter, r *http.Request) {

	var locationModel models.LocationImg

	var foodModel models.Food

	_ID, err := strconv.Atoi(r.FormValue("id"))

	if _ID < 1 {
		http.Error(rw, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(rw, "Error al convertir el ID", http.StatusInternalServerError)
		return
	}

	rw.Header().Add("content-type", "application/json")

	foodModel, err = db.GetFoodById(_ID)
	if err != nil {
		http.Error(rw, "no fue posible recuperar el plato por ID", http.StatusInternalServerError)
		return
	}

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
		locationModel, _ = imgdb.GetLocationImgById(*foodModel.LocationID)
	default:
		log.Println(err)
	}

	foodModel.Title = r.FormValue("title")
	foodModel.Description = r.FormValue("description")

	foodModel.Active = true

	locationModel.ID = *foodModel.LocationID

	cat := (r.FormValue("categories"))

	categoryArray := strings.Split(cat, ", ")

	var foodCategoryModel models.FoodCategory

	var foodCategoriesModel []models.FoodCategory

	for _, value := range categoryArray {
		foodCategoryModel.CategoryID, _ = strconv.Atoi(value)
		foodCategoryModel.FoodID = foodModel.ID
		foodCategoriesModel = append(foodCategoriesModel, foodCategoryModel)
	}

	status, err := db.UpdateFood(foodModel, locationModel, foodCategoriesModel)
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
