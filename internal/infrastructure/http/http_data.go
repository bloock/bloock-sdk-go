package http

type DataHttp struct {
	apiKey string
}

func NewDataHttp(apiKey string) DataHttp {
	return DataHttp{
		apiKey: apiKey,
	}
}

func (d DataHttp) GetApiKey() string {
	return d.apiKey
}
