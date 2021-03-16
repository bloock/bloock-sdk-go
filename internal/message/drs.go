package message

type FetchRequest struct {
	Messages []string `json:"messages"`
	Client   string   `json:"client_id"`
}

type WriteRequest struct {
	Messages []string `json:"messages"`
	Client   string   `json:"client"`
}

type WriteResponse struct {
	Success bool               `json:"success"`
	Data    *WriteResponseData `json:"data"`
}

type WriteResponseData struct {
	Anchor int `json:"anchor"`
}

type SearchMessageResponse struct {
	Success bool               `json:"success"`
	Data    *[]Receipt `json:"data"`
}

type Receipts struct {
	Messages []Receipt `json:"messages"`
}

// Receipt contains the information related to a message:
// - Message is the hexadecimal representation of a message hash
// - Anchor is the value of the anchor in which the message is contained
// - Client is the client uuid that sent the message
// - Status shows the current message status
type Receipt struct {
	Message string `json:"message"`
	Anchor  int    `json:"anchor"`
	Client  string `json:"client"`
	Status  string `json:"status"`
}