package repository

type ProoferRepository interface {
	RetrieveProof()
	VerifyProof()
	ValidateRoot()
}
