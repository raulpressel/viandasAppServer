package dtos

import "time"

type DeliveryRequest struct {
	DateStart        string `json:"dateStart"`
	DateEnd          string `json:"dateEnd"`
	DeliveryDriverID *int   `json:"idDeliveryDriver"`
}

type DeliveryResponse struct {
	DeliveryRes []DeliveryRes `json:"deliveryResponse"`
}

type DeliveryRes struct {
	DeliveryDriverRes DeliveryDriverRes `json:"deliveryDriver"`
	Delivery          []Delivery        `json:"deliveries"`
}

type Delivery struct {
	Deli Deli `json:"delivery"`
}

type Deli struct {
	ID      int            `json:"id"`
	Date    time.Time      `json:"date"`
	IdOrden int            `json:"idOrden"`
	Client  Client         `json:"client"`
	Price   float32        `json:"price"`
	Address AddressRespone `json:"address"`
}

/*
DeliveryRequest
{
    "dateStart": "19/06/2022",
    "dateEnd": "23/06/2022",
    "idDeliveryDriver": 4
} */
