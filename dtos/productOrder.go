package dtos

import "time"

type AllProductOrderResponse struct {
	ID                   uint      `json:"id"`
	ClientName           string    `json:"clientName"`
	ClientLastName       string    `json:"clientLastName"`
	ProductCategoryTitle string    `json:"productCategoryTitle"`
	ProductTitle         string    `json:"productTitle"`
	Cant                 int       `json:"cant"`
	Date                 time.Time `json:"date"`
	Status               string    `json:"status"`
}

type ProductOrderRequest struct {
	ClientName           string    `json:"clientName"`
	ClientLastName       string    `json:"clientLastName"`
	ProductCategoryTitle string    `json:"productCategoryTitle"`
	ProductTitle         string    `json:"productTitle"`
	Cant                 int       `json:"cant"`
	Date                 time.Time `json:"date"`
	Status               string    `json:"status"`
}
