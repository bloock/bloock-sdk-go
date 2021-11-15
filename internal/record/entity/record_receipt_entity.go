package entity

type RecordReceipt struct {
	anchor int
	client string
	record string
	status string
}

func NewRecordReceipt(anchor int, client, record, status string) RecordReceipt {
	return RecordReceipt{
		anchor: anchor,
		client: client,
		record: record,
		status: status,
	}
}


