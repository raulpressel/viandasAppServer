package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/productCategory"
	"viandasApp/dtos"
)

func UpdateProductCategory(rw http.ResponseWriter, r *http.Request) {
	var req dtos.ProductCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	if req.ProductCategory.ID < 1 {
		http.Error(rw, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	model, err := db.GetProductCategoryById(int(req.ProductCategory.ID))
	if err != nil {
		http.Error(rw, "no fue posible recuperar la categoria por ID", http.StatusInternalServerError)
		return
	}

	model.Title = req.ProductCategory.Title
	model.Description = req.ProductCategory.Description

	status, err := db.UpdateProductCategory(model)
	if err != nil || !status {
		http.Error(rw, "No se pudo actualizar la categoria de producto", http.StatusInternalServerError)
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
