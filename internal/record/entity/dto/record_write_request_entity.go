package dto

type RecordWriteRequest struct {
	Messages []string `json:"messages"`
}

func NewRecordWriteRequest(records []string) RecordWriteRequest {
	return RecordWriteRequest{
		Messages: records,
	}
}
