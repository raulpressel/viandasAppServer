package handlers

import (
	"net/http"

	"io"
	"log"
	"os"
	"strconv"
	imgdb "viandasApp/db/img"
	"viandasApp/handlers"

	dbcategory "viandasApp/db/categories"

	"viandasApp/models"
)

/*subir el avatar al servidor*/
func UpdateCategory(rw http.ResponseWriter, r *http.Request) {

	var locationModel models.LocationImg

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

	categoryModel, err := dbcategory.GetCategoryById(_ID)
	if err != nil {
		http.Error(rw, "no fue posible recuperar el plato por ID", http.StatusInternalServerError)
		return
	}

	if categoryModel.LocationID != nil {

		locationModel, err = imgdb.GetLocationImgById(*categoryModel.LocationID)
		if err != nil {
			http.Error(rw, "no fue posible recuperar la imagen por ID", http.StatusInternalServerError)
			return
		}
	}

	file, handle, err := r.FormFile("image")
	switch err {
	case nil:
		locationModel.Location = "/var/www/default/htdocs/public/category/" + handle.Filename

		f, err := os.OpenFile(locationModel.Location, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			http.Error(rw, "error al subir imagen de categoria "+err.Error(), http.StatusBadRequest)
			return
		}

		_, err = io.Copy(f, file)

		if err != nil {
			http.Error(rw, "error al copiar imagen de categoria "+err.Error(), http.StatusBadRequest)
			return
		}

		locationModel.Location = handlers.GetHash(locationModel.Location)

		if categoryModel.LocationID != nil {

			locationModel.ID = *categoryModel.LocationID
		} else {
			numero := 0
			categoryModel.LocationID = &numero
		}

		file.Close()
	case http.ErrMissingFile:
		if locationModel.Location != "" {

			categoryModel.LocationID = nil

		}

	default:
		log.Println(err)
	}

	categoryModel.Title = r.FormValue("title")
	categoryModel.Description = r.FormValue("description")

	if _Price, err := strconv.ParseFloat(r.FormValue("price"), 32); err != nil {
		http.Error(rw, "Error al convertir el Precio", http.StatusInternalServerError)
		return
	} else {
		categoryModel.Price = float32(_Price)
	}

	categoryModel.Active = true

	status, err := dbcategory.UpdateCategory(categoryModel, locationModel)
	if err != nil {
		http.Error(rw, "No se pudo guardar actualizar la categoria en la base de datos "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "No se pudo guardar actualizar la categoria en la base de datos ", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
