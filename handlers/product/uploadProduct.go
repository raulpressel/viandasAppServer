package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"viandasApp/db"
	db_product "viandasApp/db/product"
	"viandasApp/dtos"
	"viandasApp/handlers"
	"viandasApp/models"
)

func UploadProduct(rw http.ResponseWriter, r *http.Request) {
	if err := os.MkdirAll("/var/www/default/htdocs/public/product", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	var locationModel models.LocationImg

	rw.Header().Add("content-type", "application/json")

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
		file.Close()
	case http.ErrMissingFile:
		locationModel.Location = ""
	default:
		log.Println(err)
	}

	var model models.Product

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

	db.ExistTable(locationModel)

	status, err := db_product.UploadProduct(model, locationModel)
	if err != nil || !status {
		http.Error(rw, "No se pudo guardar el producto", http.StatusBadRequest)
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
