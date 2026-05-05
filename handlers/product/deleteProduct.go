package handlers

import (
	"net/http"
	"strconv"
	db "viandasApp/db/product"
)

func DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("idProduct")
	if len(ID) < 1 {
		http.Error(rw, "debe enviar el parametro idProduct", http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(ID)
	model, err := db.GetProductById(id)
	if err != nil {
		http.Error(rw, "no fue posible recuperar el producto por ID", http.StatusInternalServerError)
		return
	}

	hasOrders, err := db.HasProductOrders(model.Title)
	if err != nil {
		http.Error(rw, "no fue posible verificar los pedidos del producto", http.StatusInternalServerError)
		return
	}
	if hasOrders {
		http.Error(rw, "no se puede eliminar un producto que tiene pedidos asociados", http.StatusConflict)
		return
	}

	status, err := db.DeleteProduct(model)
	if err != nil || !status {
		http.Error(rw, "No se pudo eliminar el producto", http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
}
