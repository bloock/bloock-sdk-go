package proof

type ApiProofRequestBody struct {
	Messages []string `json:"messages"`
	Client   string   `json:"client"`
}