package exception

type AnchorNotFoundException struct {
}

func NewAnchorNotFoundException() AnchorNotFoundException{
	return AnchorNotFoundException{
	}
}

func(h AnchorNotFoundException) Error() string {
	return "Anchor not found"
}
