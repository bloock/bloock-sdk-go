package entity

type Network struct {
	name   string `json:"name"`
	state  string `json:"state"`
	txHash string `json:"tx_hash"`
}

func NewNetwork(name, state, txHash string) Network {
	return Network{
		name:   name,
		state:  state,
		txHash: txHash,
	}
}
