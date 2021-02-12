package cloud

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
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

var sdkParams SdkParams

type Service interface {
	ConfigParameters() error
	SdkParameters() SdkParams
}

type service struct {
	http    http.Client
	envVars config.Constants
}

func Azure(http http.Client, envVars config.Constants) Service {
	return &service{http, envVars}
}

type SdkParams struct {
	Host                string `json:"SDK_HOST"`
	MessageWrite        string `json:"SDK_WRITE_ENDPOINT"`
	MessageFetch        string `json:"SDK_FETCH_ENDPOINT"`
	MessageProof        string `json:"SDK_PROOF_ENDPOINT"`
	SmartContract       string `json:"SDK_CONTRACT_ADDRESS"`
	ContractAbi         string `json:"SDK_CONTRACT_ABI"`
	Provider            string `json:"SDK_PROVIDER"`
	WriteInterval       int    `json:"SDK_WRITE_INTERVAL"`
	WaitIntervalFactor  int    `json:"SDK_WAIT_MESSAGE_INTERVAL_FACTOR"`
	WaitIntervalDefault int    `json:"SDK_WAIT_MESSAGE_INTERVAL_DEFAULT"`
	ConfigInterval      int `json:"SDK_CONFIG_INTERVAL"`
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
	resp, err := s.http.Request("", "GET", url, headers, nil)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(resp, &sdkParams); err != nil {
		return err
	}

	return nil
}

func (s *service) SdkParameters() SdkParams {
	return sdkParams
}

func (s *service) authHeaders(httpVerb, url, body string) (map[string]string, error) {
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
		"x-ms-date":           gmtDateTime,
		"x-ms-content-sha256": hashedContent,
		"Authorization":       fmt.Sprintf("HMAC-SHA256 Credential=%s&SignedHeaders=%s&Signature=%s", credential, signedHeaders, signature),
	}, nil
}
