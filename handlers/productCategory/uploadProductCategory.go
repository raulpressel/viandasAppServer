package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/productCategory"
	"viandasApp/dtos"
	"viandasApp/models"
)

func UploadProductCategory(rw http.ResponseWriter, r *http.Request) {
	var req dtos.ProductCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	var model models.ProductCategory
	model.Title = req.ProductCategory.Title
	model.Description = req.ProductCategory.Description
	model.Active = true

	status, err := db.UploadProductCategory(model)
	if err != nil || !status {
		http.Error(rw, "No se pudo guardar la categoria de producto", http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(map[string]interface{}{
		"productCategory": dtos.AllProductCategoryResponse{
			ID:          model.ID,
			Title:       model.Title,
			Description: model.Description,
		},
	})
}
