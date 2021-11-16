package entity

type RecordReceipt struct {
	Anchor int
	Client string
	Record string
	Status string
}

func NewRecordReceipt(anchor int, client, record, status string) RecordReceipt {
	return RecordReceipt{
		Anchor: anchor,
		Client: client,
		Record: record,
		Status: status,
	}
}


