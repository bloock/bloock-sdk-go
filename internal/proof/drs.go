package proof

type ApiProofRequestBody struct {
	Messages []string `json:"messages"`
	Client   string   `json:"client"`
}

type ProofResponse struct {
	Success bool   `json:"success"`
	Data    *Proof `json:"data"`
}
