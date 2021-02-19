package message


type FetchRequest struct {
	Messages []string `json:"messages"`
	Client   string   `json:"client_id"`
}

type WriteRequest struct {
	Messages []string `json:"messages"`
	Client   string   `json:"client"`
}

type WriteResponse struct {
	Anchor int `json:"anchor"`
}
