package exception

type InvalidArgumentException struct {
}

func NewInvalidArgumentException() InvalidArgumentException{
	return InvalidArgumentException{
	}
}

func(h InvalidArgumentException) Error() string {
	return "Invalid argument provided"
}
