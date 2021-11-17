package exception

type InvalidRecordException struct {
}

func NewInvalidRecordException() InvalidRecordException{
	return InvalidRecordException{
	}
}

func(h InvalidRecordException) Error() string {
	return "Record not valid"
}
