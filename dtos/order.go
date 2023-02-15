package dtos

import (
	"time"
)

type ResCantDB struct {
	ID          int
	Description string
	Title       string
	Color       string
	Price       float32
	Cant        int
}

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

type SaveOrderResponse struct {
	IDOrder int `json:"idOrder"`
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

type OrdersResponse struct {
	TandasTable TandaTable `json:"getOrdersResponse"`
}

type TandaTable struct {
	TandaTable    []Tanda         `json:"tandaTable"`
	CategoryTable []CategoryTable `json:"categoryTable"`
}

type Tanda struct {
	Tanda         TandaRes        `json:"tanda"`
	CategoryTable []CategoryTable `json:"categoryTable"`
	OrderRes      []OrdersRes     `json:"order"`
}

type CategoryTable struct {
	Category CategoryResponse `json:"category"`
	Cant     int              `json:"cant"`
}

type OrdersRes struct {
	ID            int             `json:"id"`
	OrderDate     time.Time       `json:"date"`
	Observation   string          `json:"observation"`
	Total         float32         `json:"total"`
	Status        string          `json:"status"`
	Client        Client          `json:"client"`
	CategoryTable []CategoryTable `json:"categoryTable"`
	Address       AddressRespone  `json:"address"`
	Observations  []string        `json:"observations"`
}
