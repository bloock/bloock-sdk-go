package infrastructure

type HttpClient interface {
	Get(url string, headers map[string]string) (interface{}, error)
	Post(url string, body interface{}, headers map[string]string) (interface{}, error)
}
