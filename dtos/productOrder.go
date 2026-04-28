package dtos

import "time"

type ProductOrderItemResponse struct {
	ProductTitle         string `json:"productTitle"`
	ProductCategoryTitle string `json:"productCategoryTitle"`
	Cant                 int    `json:"cant"`
}

type ProductOrderResponse struct {
	ID             uint                       `json:"id"`
	ClientName     string                     `json:"clientName"`
	ClientLastName string                     `json:"clientLastName"`
	Date           time.Time                  `json:"date"`
	Status         string                     `json:"status"`
	Products       []ProductOrderItemResponse `json:"products"`
}

type ProductOrderItemRequest struct {
	ProductTitle         string `json:"productTitle"`
	ProductCategoryTitle string `json:"productCategoryTitle"`
	Cant                 int    `json:"cant"`
}

type ProductOrderRequest struct {
	ClientName     string                    `json:"clientName"`
	ClientLastName string                    `json:"clientLastName"`
	Date           time.Time                 `json:"date"`
	Status         string                    `json:"status"`
	Products       []ProductOrderItemRequest `json:"products"`
}
