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

const (
	signedHeaders = "X-Ms-Date;Host;X-Ms-Content-Sha256"
	credential    = "ihs8-l9-s0:JPRPUeiXJGsAzFiW9WDc"
	secret        = "1UA2dijC0SIVyrPKUKG0gT0oXxkVaMrUfJuXkLr+i0c="
)

var sdkParams SdkParams

type Service interface {
	RequestSdkParameters() error
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
	Host                 string `json:"SDK_HOST"`
	MessageWrite         string `json:"SDK_WRITE_ENDPOINT"`
	MessageFetch         string `json:"SDK_FETCH_ENDPOINT"`
	MessageProof         string `json:"SDK_PROOF_ENDPOINT"`
	SmartContractAddress string `json:"SDK_CONTRACT_ADDRESS"`
	ContractAbi          string `json:"SDK_CONTRACT_ABI"`
	Web3Provider         string `json:"SDK_PROVIDER"`
	WriteInterval        string `json:"SDK_WRITE_INTERVAL"`
	WaitIntervalFactor   string `json:"SDK_WAIT_MESSAGE_INTERVAL_FACTOR"`
	WaitIntervalDefault  string `json:"SDK_WAIT_MESSAGE_INTERVAL_DEFAULT"`
	ConfigInterval       string `json:"SDK_CONFIG_INTERVAL"`
}

type AzureParameters struct {
	Items []Item `json:"items"`
}

type Item struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (s *service) RequestSdkParameters() error {
	// check if test environment
	var endpoint string
	if s.envVars.Cloud.Azure.Test {
		endpoint = s.envVars.Cloud.Azure.PathTest
	} else {
		endpoint = s.envVars.Cloud.Azure.PathProd
	}

	headers, err := s.authHeaders("GET", endpoint, "")
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s%s", s.envVars.Cloud.Azure.Host, endpoint)

	resp, err := s.http.GetRequest(fmt.Sprintf("https://%s", url), headers)
	if err != nil {
		return err
	}

	var res AzureParameters
	if err := json.Unmarshal(resp, &res); err != nil {
		return errors.New(fmt.Sprintf("error unmarshaling: %s", err.Error()))
	}

	items := make(map[string]string)
	for _, item := range res.Items {
		items[item.Key] = item.Value
	}

	bytes, err := json.Marshal(items)
	if err != nil {
		return errors.New(fmt.Sprintf("error marshaling: %s", err.Error()))
	}

	if err := json.Unmarshal(bytes, &sdkParams); err != nil {
		return errors.New(fmt.Sprintf("error unmarshaling: %s", err.Error()))
	}

	return nil
}

func (s *service) SdkParameters() SdkParams {
	return sdkParams
}

func (s *service) authHeaders(httpVerb, endpoint, body string) (map[string]string, error) {

	gmtDateTime := time.Now().UTC().Format(netHttp.TimeFormat)

	hashedContent := contentHashBase64([]byte(body))
	stringToSign := fmt.Sprintf("%s\n%s\n%s;%s;%s", strings.ToUpper(httpVerb), endpoint, gmtDateTime, s.envVars.Cloud.Azure.Host, hashedContent)

	key, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return nil, err
	}

	signature := getHmac(stringToSign, key)

	return map[string]string{
		"x-ms-date":           gmtDateTime,
		"x-ms-content-sha256": hashedContent,
		"Authorization":       fmt.Sprintf("HMAC-SHA256 Credential=%s, SignedHeaders=%s, Signature=%s", credential, signedHeaders, signature),
	}, nil
}

func contentHashBase64(content []byte) string {
	hasher := sha256.New()
	hasher.Write(content)
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}

func getHmac(content string, key []byte) string {
	hmac := hmac.New(sha256.New, key)
	hmac.Write([]byte(content))
	return base64.StdEncoding.EncodeToString(hmac.Sum(nil))
}