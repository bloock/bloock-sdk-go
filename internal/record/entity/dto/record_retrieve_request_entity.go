package dto

type RecordRetrieveRequest struct {
	Messages []string `json:"messages"`
}

func NewRecordRetrieveRequest(records []string) RecordRetrieveRequest {
	return RecordRetrieveRequest{
		Messages: records,
	}
}
