package entity

type Network struct {
	name   string
	state  string
	txHash string
}

func NewNetwork(name, state, txHash string) Network {
	return Network{
		name:   name,
		state:  state,
		txHash: txHash,
	}
}
