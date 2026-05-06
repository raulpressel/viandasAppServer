package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/product"
	"viandasApp/dtos"
	"viandasApp/models"
)

func UploadProduct(rw http.ResponseWriter, r *http.Request) {
	var req dtos.ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	var model models.Product
	model.Title = req.Product.Title
	model.Description = req.Product.Description
	model.Price = req.Product.Price
	model.Available = req.Product.Available
	model.ProductCategoryID = req.Product.ProductCategoryID

	status, err := db.UploadProduct(model)
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
	})
}
