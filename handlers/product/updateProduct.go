package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	imgdb "viandasApp/db/img"
	db "viandasApp/db/product"
	"viandasApp/dtos"
	"viandasApp/handlers"
	"viandasApp/models"
)

func UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	var locationModel models.LocationImg
	var model models.Product

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

	model, err = db.GetProductById(_ID)
	if err != nil {
		http.Error(rw, "no fue posible recuperar el producto por ID", http.StatusInternalServerError)
		return
	}

	if model.LocationID != nil {
		locationModel, err = imgdb.GetLocationImgById(*model.LocationID)
		if err != nil {
			http.Error(rw, "no fue posible recuperar la imagen por ID", http.StatusInternalServerError)
			return
		}
	}

	file, handle, err := r.FormFile("image")
	switch err {
	case nil:
		locationModel.Location = "/var/www/default/htdocs/public/product/" + handle.Filename

		f, err := os.OpenFile(locationModel.Location, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			http.Error(rw, "error al subir imagen producto "+err.Error(), http.StatusBadRequest)
			return
		}

		_, err = io.Copy(f, file)
		if err != nil {
			http.Error(rw, "error al copiar imagen producto "+err.Error(), http.StatusBadRequest)
			return
		}

		locationModel.Location = handlers.GetHash(locationModel.Location)

		if model.LocationID != nil {
			locationModel.ID = *model.LocationID
		} else {
			numero := 0
			model.LocationID = &numero
		}

		file.Close()
	case http.ErrMissingFile:
		if locationModel.Location == "" {
			model.LocationID = nil
		}
	default:
		log.Println(err)
	}

	model.Title = r.FormValue("title")
	model.Description = r.FormValue("description")

	price, err := strconv.ParseFloat(r.FormValue("price"), 32)
	if err != nil {
		http.Error(rw, "precio inválido "+err.Error(), http.StatusBadRequest)
		return
	}
	model.Price = float32(price)

	available, err := strconv.ParseBool(r.FormValue("available"))
	if err != nil {
		model.Available = true
	} else {
		model.Available = available
	}

	productCategoryID, err := strconv.Atoi(r.FormValue("productCategoryId"))
	if err != nil {
		http.Error(rw, "productCategoryId inválido "+err.Error(), http.StatusBadRequest)
		return
	}
	model.ProductCategoryID = productCategoryID

	status, err := db.UpdateProduct(model, locationModel)
	if err != nil || !status {
		http.Error(rw, "No se pudo actualizar el producto", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(dtos.AllProductResponse{
		ID:                model.ID,
		Title:             model.Title,
		Description:       model.Description,
		Price:             model.Price,
		Available:         model.Available,
		ProductCategoryID: model.ProductCategoryID,
		UrlImage:          locationModel.Location,
	})
}
