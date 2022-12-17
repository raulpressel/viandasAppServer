package dtos

import "time"

type OrderRequest struct {
	IDClient         int                `json:"idClient"`
	Observation      string             `json:"observation"`
	Total            float32            `json:"total"`
	Date             string             `json:"date"`
	DaysOrderRequest []DaysOrderRequest `json:"daysOrderRequest"`
}

type DaysOrderRequest struct {
	Amount      int    `json:"cant"`
	IDDayFood   int    `json:"idDayFood"`
	IDAddress   int    `json:"idAddress"`
	Observation string `json:"observation"`
}

type OrderViewerResponse struct {
	Order []OrderResponse `json:"orderViewer"`
}

type OrderResponse struct {
	ID          int       `json:"id"`
	OrderDate   time.Time `json:"date"`
	Observation string    `json:"observation"`
	Total       float32   `json:"total"`
	Status      string    `json:"status"`
	DateStart   time.Time `json:"dateStart"`
	DateEnd     time.Time `json:"dateEnd"`
}

type FullOrderResponse struct {
	ID          int                `json:"id"`
	OrderDate   time.Time          `json:"date"`
	Observation string             `json:"observation"`
	Total       float32            `json:"total"`
	Status      string             `json:"status"`
	DayOrder    []DayOrderResponse `json:"daysOrder"`
}

type DayOrderResponse struct {
	ID          int              `json:"id"`
	Date        time.Time        `json:"date"`
	Food        FoodResponse     `json:"food"`
	Category    CategoryResponse `json:"category"`
	Amount      int              `json:"cant"`
	Observation string           `json:"observation"`
	Address     AddressRespone   `json:"address"`
	Status      string           `json:"status"`
}
