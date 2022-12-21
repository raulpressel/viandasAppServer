package dtos

type DeliveryDriverRequest struct {
	DeliveryDriver DeliveryDriver `json:"deliveryDriver"`
}

type DeliveryDriver struct {
	ID       int            `json:"id"`
	DNI      int            `json:"dni"`
	Name     string         `json:"name"`
	LastName string         `json:"lastName"`
	Phone    string         `json:"phone"`
	BornDate string         `json:"bornDate"`
	Vehicle  VehicleRequest `json:"vehicle"`
	Address  AddressRequest `json:"address"`
}

type VehicleRequest struct {
	ID     int    `json:"id"`
	Brand  string `json:"brand"`
	Models string `json:"model"`
	Patent string `json:"patent"`
	Year   int    `json:"year"`
}
