package infrastructure

//go:generate mockgen -source=internal/infrastructure/http_client.go -destination internal/infrastructure/http/mockhttp/mocks_http.go -package=mockhttp
type HttpClient interface {
	Get(url string, headers map[string]string) (interface{}, error)
	Post(url string, body interface{}, headers map[string]string) (interface{}, error)
}
