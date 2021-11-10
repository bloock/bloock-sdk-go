package entity

type Anchor struct {
	id         int       `json:"id"`
	blockRoots []string  `json:"block_roots"`
	networks   []Network `json:"networks"`
	root       string    `json:"root"`
	status     string    `json:"status"`
}

func NewAnchor(id int, blockRoots []string, networks []Network, root, status string) Anchor {
	return Anchor{
		id:         id,
		blockRoots: blockRoots,
		networks:   networks,
		root:       root,
		status:     status,
	}
}
