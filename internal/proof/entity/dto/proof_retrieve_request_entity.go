package dto

type ProofRetrieveRequest struct {
	Messages []string `json:"messages"`
}

func NewProofRetrieveRequest(messages []string) ProofRetrieveRequest {
	return ProofRetrieveRequest{
		Messages: messages,
	}
}