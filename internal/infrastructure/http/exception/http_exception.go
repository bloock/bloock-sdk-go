package exception

type HttpRequestException struct {
	message string
}

func NewHttpRequestException(message string) HttpRequestException{
	return HttpRequestException{
		message: message,
	}
}

func(h HttpRequestException) Error() string {
	return "HttpClient was not successful: " + h.message
}
