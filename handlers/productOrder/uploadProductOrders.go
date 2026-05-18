package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/productOrder"
	"viandasApp/dtos"
	"viandasApp/models"
)

func UploadProductOrder(rw http.ResponseWriter, r *http.Request) {
	var req dtos.ProductOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	var items []models.ProductOrderItem
	for _, p := range req.Products {
		items = append(items, models.ProductOrderItem{
			ProductTitle:         p.ProductTitle,
			ProductCategoryTitle: p.ProductCategoryTitle,
			Cant:                 p.Cant,
			Status:               "pending",
		})
	}

	order := models.ProductOrder{
		ClientName:     req.ClientName,
		ClientLastName: req.ClientLastName,
		Date:           req.Date,
		Status:         "pending",
		Products:       items,
	}

	saved, err := db.UploadProductOrder(order)
	if err != nil {
		http.Error(rw, "No se pudo guardar la orden de productos", http.StatusBadRequest)
		return
	}

	var itemsResponse []dtos.ProductOrderItemResponse
	for _, p := range saved.Products {
		itemsResponse = append(itemsResponse, dtos.ProductOrderItemResponse{
			ProductTitle:         p.ProductTitle,
			ProductCategoryTitle: p.ProductCategoryTitle,
			Cant:                 p.Cant,
		})
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(map[string]interface{}{
		"productOrder": dtos.ProductOrderResponse{
			ID:             saved.ID,
			ClientName:     saved.ClientName,
			ClientLastName: saved.ClientLastName,
			Date:           saved.Date,
			Status:         saved.Status,
			Products:       itemsResponse,
		},
	})
}
