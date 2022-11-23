package dtos

type Client struct {
	Client RegisterRequest `json:"client"`
}

type RegisterRequest struct {
	ID             int                `json:"id"`
	PhonePrimary   string             `json:"phonePrimary"`
	PhoneSecondary string             `json:"phoneSecondary"`
	ObsClient      string             `json:"observation"`
	BornDate       string             `json:"bornDate"`
	Pathologies    []PathologyRequest `json:"pathologies"`
	Address        []AddressRequest   `json:"addresses"`
}
