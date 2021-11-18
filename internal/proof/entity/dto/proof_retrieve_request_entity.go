package dto

type ProofRetrieveRequest struct {
	messages []string
}

func NewProofRetrieveRequest(messages []string) ProofRetrieveRequest {
	return ProofRetrieveRequest{
		messages: messages,
	}
}