package dto

type RecordWriteRequest struct {
	messages []string `json:"messages"`
}

func NewRecordWriteRequest(records []string) RecordWriteRequest {
	return RecordWriteRequest{
		messages: records,
	}
}
