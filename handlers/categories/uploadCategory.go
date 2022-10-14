package handlers

import (
	"net/http"
	"viandasApp/db"

	"io"
	"log"
	"os"
	"strconv"
	"viandasApp/handlers"

	dbcategory "viandasApp/db/categories"


	"viandasApp/models"
)

func UploadCategory(rw http.ResponseWriter, r *http.Request) {

	if err := os.MkdirAll("/var/www/default/htdocs/public/category", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	var locationModel models.LocationImg

	rw.Header().Add("content-type", "application/json")

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
		locationModel.Location = ""
	default:
		log.Println(err)
	}

	var categoryModel models.Category

	/* foodModel.ID, _ = strconv.Atoi(r.FormValue("id")) */

	categoryModel.Title = r.FormValue("title")
	categoryModel.Description = r.FormValue("description")

	if _Price, err := strconv.ParseFloat(r.FormValue("price"), 32); err != nil {
		http.Error(rw, "Error al convertir el Precio", http.StatusInternalServerError)
		return
	} else {
		categoryModel.Price = float32(_Price)
	}

	categoryModel.Active = true

	
	db.ExistTable(categoryModel)
	db.ExistTable(locationModel)




	status, err := dbcategory.UploadCategory(categoryModel, locationModel)
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

	/* 	var categoryDto dtos.Category

	   	err := json.NewDecoder(r.Body).Decode(&categoryDto)

	   	if err != nil {
	   		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
	   		return
	   	}

	   	category := categoryDto.ToModelCategory()

	   	category.Active = true

	   	db.ExistTable(category)

	   	status, err := dbCategories.UploadCategory(*category)
	   	if err != nil {
	   		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario "+err.Error(), 400)
	   		return
	   	}

	   	if !status { //esto es igual a !status == false
	   		http.Error(w, "no se ha logrado insertar el registro  // status = false ", 400)
	   		return
	   	}

	   	w.WriteHeader(http.StatusCreated) */

}
