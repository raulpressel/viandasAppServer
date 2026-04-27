package handlers

import (
	"net/http"
	"strconv"
	db "viandasApp/db/productCategory"
)

func DeleteProductCategory(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("idProductCategory")
	if len(ID) < 1 {
		http.Error(rw, "debe enviar el parametro idProductCategory", http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(ID)
	model, err := db.GetProductCategoryById(id)
	if err != nil {
		http.Error(rw, "no fue posible recuperar la categoria por ID", http.StatusInternalServerError)
		return
	}

	model.Active = false

	status, err := db.DeleteProductCategory(model)
	if err != nil || !status {
		http.Error(rw, "No se pudo eliminar la categoria de producto", http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
}
