package cloud

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/config"
	"github.com/enchainte/enchainte-sdk-go/pkg/http"
	netHttp "net/http"
	"strings"
	"time"
)

var (
	signedHeaders = "x-ms-date;host;x-ms-content-sha256"
	credential    = "ihs8-l9-s0:JPRPUeiXJGsAzFiW9WDc"
	secret        = "1UA2dijC0SIVyrPKUKG0gT0oXxkVaMrUfJuXkLr+i0c="
)

var params map[string]interface{}

type Services interface {
	ConfigParameters() error
	Parameters() map[string]interface{}
}

type service struct {
	http http.Client
	envVars config.Constants
}

func Azure(http http.Client, envVars config.Constants) Services {
	return &service{http, envVars}
}

func (s *service) ConfigParameters() error {
	// check if test environment
	var path string
	if s.envVars.Cloud.Azure.Test {
		path = s.envVars.Cloud.Azure.PathTest
	} else {
		path = s.envVars.Cloud.Azure.PathProd
	}

	headers, err := s.authHeaders("GET", path, "")
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s%s", s.envVars.Cloud.Azure.Host, path)
	resp, err := s.http.Request("", "GET", url, headers, "")
	if err != nil {
		return err
	}

	var res map[string]interface{}
	if err := json.Unmarshal(resp, &res); err != nil {
		return err
	}

	items, ok := res["items"].(map[string]interface{})
	if !ok {
		return errors.New("res[items] is not a list")
	}
	for k, v := range items {
		params[k] = v
	}
	return nil
}

func (s *service) Parameters() map[string]interface{} {
	return params
}


func (s *service) authHeaders(httpVerb, url, body string) (map[string]string, error){
	httpVerb = strings.ToUpper(httpVerb)
	gmtDateTime := time.Now().UTC().Format(netHttp.TimeFormat)

	h := sha256.New()
	h.Write([]byte(body))
	hashedContent := base64.StdEncoding.EncodeToString(h.Sum(nil))

	stringToSign := fmt.Sprintf("%s\n%s\n%s;%s;%s", httpVerb, url, gmtDateTime, s.envVars.Cloud.Azure.Host, hashedContent)

	key, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return nil, err
	}
	hmac := hmac.New(sha256.New, key)
	hmac.Write([]byte(stringToSign))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))


	 return map[string]string{
		"x-ms-date": gmtDateTime,
		"x-ms-content-sha256" : hashedContent,
		"Authorization" : fmt.Sprintf("HMAC-SHA256 Credential=%s&SignedHeaders=%s&Signature=%s", credential, signedHeaders, signature),
	}, nil
}

