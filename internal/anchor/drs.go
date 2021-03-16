package anchor

type GetAnchorResponse struct {
	Success bool                   `json:"success"`
	Data    *getAnchorDataResponse `json:"data"`
}

type getAnchorDataResponse struct {
	AnchorId   int       `json:"anchor_id"`
	BlockRoots []string  `json:"block_roots"`
	Networks   []network `json:"networks"`
	Root       string    `json:"root"`
	Status     string    `json:"status"`
}

type network struct {
	Name   string `json:"name"`
	State  string `json:"state"`
	TxHash string `json:"tx_hash"`
}
