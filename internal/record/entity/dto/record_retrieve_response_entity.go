package dto

type RecordRetrieveResponse struct {
	Anchor int `json:"anchor" default:"0"`
	Client string `json:"client" default:""`
	Message string `json:"message" default:""`
	Status string `json:"status" default:"Pending"`
}
