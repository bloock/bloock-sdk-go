package entity

type Network struct {
	Name   string `json:"name"`
	State  string `json:"state"`
	TxHash string `json:"tx_hash"`
}
