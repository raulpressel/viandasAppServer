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
	DeliveryDriver DeliveryByDeliveryDriver `json:"deliveryDriver"`
}

type DeliveryByDeliveryDriver struct {
	ID       int        `json:"id"`
	DNI      int        `json:"dni"`
	Name     string     `json:"name"`
	LastName string     `json:"lastName"`
	Phone    string     `json:"phone"`
	Delivery []Delivery `json:"deliveries"`
}

type Delivery struct {
	Deli Deli `json:"delivery"`
}

type Deli struct {
	ID      int            `json:"id"`
	Date    time.Time      `json:"date"`
	IdOrder int            `json:"idOrder"`
	Client  Client         `json:"client"`
	Price   float32        `json:"price"`
	Address AddressRespone `json:"address"`
}

type ResponseExcel struct {
	IdOrden        int
	DeliveryDriver string
	Client         string
	Address        string
	Deliverires    []DeliveryExcel
}

type DeliveryExcel struct {
	Date  time.Time
	Price float32
}
