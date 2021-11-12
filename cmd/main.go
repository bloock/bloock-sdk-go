package main

import (
	"encoding/json"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/config"
	repository2 "github.com/enchainte/enchainte-sdk-go/config/repository"
	"github.com/enchainte/enchainte-sdk-go/config/service"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor/repository"
	service2 "github.com/enchainte/enchainte-sdk-go/internal/anchor/service"
	"github.com/enchainte/enchainte-sdk-go/internal/cloud"
	"github.com/enchainte/enchainte-sdk-go/internal/credential"
	http2 "github.com/enchainte/enchainte-sdk-go/internal/infrastructure/http"
	"github.com/enchainte/enchainte-sdk-go/internal/message"
	"github.com/enchainte/enchainte-sdk-go/internal/proof"
	"github.com/enchainte/enchainte-sdk-go/pkg/blockchain"
	"github.com/enchainte/enchainte-sdk-go/pkg/crypto"
	"github.com/enchainte/enchainte-sdk-go/pkg/http"
	"log"
)

func main() {
	configData := repository2.NewConfigData()
	configRepo := repository2.NewConfigRepository(configData)
	configService := service.NewConfigService(configRepo)


	api := http2.NewDataHttp("C1vfvhN2mPUeX0KikgGHVIUSofZIfX6Q4bx0kf7DuAHMt3cuELO2UGdYLUw9bS29")
	httpClient := http2.NewHttp(api)
	anchorRepo := repository.NewAnchorRepository(httpClient, configService)
	anchorService := service2.NewAnchorService(anchorRepo, configService)

	//anchor, err := anchorService.WaitAnchor(1, 100)
	anchor2, err := anchorService.WaitAnchor(1111111, 100)
	if err != nil {
		log.Println(err)
	}
	//log.Println(anchor)
	log.Println(anchor2)


}

type Services struct {
	Message    message.Service
	Proof      proof.Service
	Credential credential.Service
}

func EnchainteClient(apiKey string) Services {
	http := http.NewClient()
	constants := config.EnvVariables()
	azureService := cloud.Azure(http, constants)
	if err := azureService.RequestSdkParameters(); err != nil {
		log.Fatalf(fmt.Sprintf("fatal error: something went wrong when requesting the SDK paramaters to Azure:: %s", err.Error()))
	}

	hasher := crypto.Blake2b()
	bc := blockchain.Web3(constants, azureService.SdkParameters())

	return Services{
		Message:    message.NewService(apiKey, http, azureService.SdkParameters()),
		Proof:      proof.NewService(apiKey, http, azureService.SdkParameters(), hasher, bc),
		Credential: credential.NewService(apiKey, http, azureService.SdkParameters()),
	}
}

// JsonRemarshal takes a json string and returns the byte representation of that json with the fields alphabetically sorted by key
func JsonRemarshal(data string) ([]byte, error) {
	var ifce interface{}
	err := json.Unmarshal([]byte(data), &ifce)
	if err != nil {
		return []byte{}, err
	}
	output, err := json.Marshal(ifce)
	if err != nil {
		return []byte{}, err
	}
	return output, nil
}

// ToSortedJson takes a map or struct as parameter and returns its byte representation with the fields alphabetically sorted by key
func ToSortedJson(data interface{}) ([]byte, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var ifce interface{}
	err = json.Unmarshal(bytes, &ifce)
	if err != nil {
		return []byte{}, err
	}
	output, err := json.Marshal(ifce)
	if err != nil {
		return []byte{}, err
	}
	return output, nil
}