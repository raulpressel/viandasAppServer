package dtos

type DiscountRequest struct {
	Discount Discount `json:"discount"`
}

type Discount struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Cant        int     `json:"cant"`
	Percentage  float32 `json:"percentage"`
}

type DiscountResponse struct {
	Discount []Discount `json:"discount"`
}

type ZoneRequest struct {
	Zone Zone `json:"zone"`
}

type Zone struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

type ZoneResponse struct {
	Zone []Zone `json:"zone"`
}

type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type ZonePolygon struct {
	IDZone int      `json:"idZone"`
	Paths  []LatLng `json:"paths"`
}

type ZonePolygonResponse struct {
	Zones []ZonePolygon `json:"zones"`
}

type ZoneAddressChange struct {
	ID        int    `json:"id"`
	Street    string `json:"street"`
	Number    string `json:"number"`
	OldIDZone int    `json:"oldIdZone"`
	NewIDZone int    `json:"newIdZone"`
}

type UpdateZoneAddressResponse struct {
	DryRun      bool                `json:"dryRun"`
	Total       int                 `json:"total"`
	Changed     []ZoneAddressChange `json:"changed"`
	WithoutZone []ZoneAddressChange `json:"withoutZone"`
}
