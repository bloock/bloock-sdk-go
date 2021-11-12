package dto

type MessageWriteRequest struct {
	messages []string `json:"messages"`
}

func NewMessageWriteRequest(messages []string) MessageWriteRequest{
	return MessageWriteRequest{
		messages: messages,
	}
}
