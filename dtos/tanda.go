package dtos

type TandaRequest struct {
	Tanda            TandaReq `json:"tanda"`
	IdDeliveryDriver int      `json:"idDeliveryDriver"`
}

type TandaReq struct {
	ID             int                   `json:"id"`
	Description    string                `json:"description"`
	HourStart      string                `json:"hourStart"`
	HourEnd        string                `json:"hourEnd"`
	DeliveryDriver DeliveryDriverRequest `json:"deliveryDriver"`
}

type TandaResponse struct {
	Tanda []TandaRes `json:"tanda"`
}

type TandaRes struct {
	ID             int               `json:"id"`
	Description    string            `json:"description"`
	HourStart      string            `json:"hourStart"`
	HourEnd        string            `json:"hourEnd"`
	DeliveryDriver DeliveryDriverRes `json:"deliveryDriver"`
}

type TandaAddressRequest struct {
	IDTanda   int   `json:"idTanda"`
	IDAddress []int `json:"idAddress"`
}

type TandasRequest struct {
	IDTanda []int `json:"idTanda"`
}
