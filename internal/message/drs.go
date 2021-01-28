package message


type ApiFetchRequestBody struct {
	Messages []string `json:"messages"`
	Client   string   `json:"client"`
}

