package dto

type RecordWriteResponse struct {
	Anchor int `json:"anchor"`
	Client string `json:"client"`
	Messages []string `json:"messages"`
	Status string `json:"status"`
}


