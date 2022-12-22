package dtos

import "time"

type DeliveryDriverRequest struct {
	DeliveryDriver DeliveryDriverReq `json:"deliveryDriver"`
}

type DeliveryDriverReq struct {
	ID       int            `json:"id"`
	DNI      int            `json:"dni"`
	Name     string         `json:"name"`
	LastName string         `json:"lastName"`
	Phone    string         `json:"phone"`
	BornDate string         `json:"bornDate"`
	Vehicle  Vehicle        `json:"vehicle"`
	Address  AddressRequest `json:"address"`
}

type Vehicle struct {
	ID     int    `json:"id"`
	Brand  string `json:"brand"`
	Models string `json:"model"`
	Patent string `json:"patent"`
	Year   int    `json:"year"`
}

type DeliveryDriverResponse struct {
	DeliveryDriver []DeliveryDriverRes `json:"deliveryDriver"`
}

type DeliveryDriverRes struct {
	ID       int            `json:"id"`
	DNI      int            `json:"dni"`
	Name     string         `json:"name"`
	LastName string         `json:"lastName"`
	Phone    string         `json:"phone"`
	BornDate time.Time      `json:"bornDate"`
	Vehicle  Vehicle        `json:"vehicle"`
	Address  AddressRespone `json:"address"`
}
