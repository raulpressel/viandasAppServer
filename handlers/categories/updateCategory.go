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
		file.Close()
	case http.ErrMissingFile:
		locationModel, _ = imgdb.GetLocationImgById(*categoryModel.LocationID)
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
		http.Error(rw, "No se pudo guardar el mensaje en la base de datos "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado insertar el registro  // status = false ", 400)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

	/* 	w.Header().Add("content-type", "application/json")

	   	var categoryDto dtos.Category

	   	err := json.NewDecoder(r.Body).Decode(&categoryDto)

	   	if err != nil {
	   		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
	   		return
	   	}

	   	var categoryModel models.Category

	   	categoryModel, _ = dbCategories.GetCategoryById(categoryDto.Category.ID)

	   	if categoryDto.Category.Description != categoryModel.Description {
	   		categoryModel.Description = categoryDto.Category.Description
	   	}

	   	if categoryDto.Category.Title != categoryModel.Title {
	   		categoryModel.Title = categoryDto.Category.Title
	   	}

	   	if categoryDto.Category.Price != categoryModel.Price {
	   		categoryModel.Price = categoryDto.Category.Price
	   	}

	   	categoryModel.Active = true

	   	status, err := dbCategories.UpdateCategory(categoryModel)

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
