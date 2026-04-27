package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/productOrder"
	"viandasApp/dtos"
	"viandasApp/models"
)

func UploadProductOrders(rw http.ResponseWriter, r *http.Request) {
	var req []dtos.ProductOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	var orders []models.ProductOrder
	for _, item := range req {
		orders = append(orders, models.ProductOrder{
			ClientName:           item.ClientName,
			ClientLastName:       item.ClientLastName,
			ProductCategoryTitle: item.ProductCategoryTitle,
			ProductTitle:         item.ProductTitle,
			Cant:                 item.Cant,
			Date:                 item.Date,
			Status:               item.Status,
		})
	}

	saved, err := db.UploadProductOrders(orders)
	if err != nil {
		http.Error(rw, "No se pudo guardar las ordenes de productos", http.StatusBadRequest)
		return
	}

	var responseModel []dtos.AllProductOrderResponse
	for _, o := range saved {
		responseModel = append(responseModel, dtos.AllProductOrderResponse{
			ID:                   o.ID,
			ClientName:           o.ClientName,
			ClientLastName:       o.ClientLastName,
			ProductCategoryTitle: o.ProductCategoryTitle,
			ProductTitle:         o.ProductTitle,
			Cant:                 o.Cant,
			Date:                 o.Date,
			Status:               o.Status,
		})
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(map[string]interface{}{"productOrders": responseModel})
}
