package dto

type RecordRetrieveRequest struct {
	messages []string `json:"messages"`
}

func NewRecordRetrieveRequest(records []string) RecordRetrieveRequest {
	return RecordRetrieveRequest{
		messages: records,
	}
}
