package main

import (
	"encoding/json"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/config"
	"github.com/enchainte/enchainte-sdk-go/config/entity"
	"github.com/enchainte/enchainte-sdk-go/config/repository"
	"github.com/enchainte/enchainte-sdk-go/config/service"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor"
	"github.com/enchainte/enchainte-sdk-go/internal/cloud"
	"github.com/enchainte/enchainte-sdk-go/internal/credential"
	blockchain2 "github.com/enchainte/enchainte-sdk-go/internal/infrastructure/blockchain"
	"github.com/enchainte/enchainte-sdk-go/internal/message"
	"github.com/enchainte/enchainte-sdk-go/internal/proof"
	"github.com/enchainte/enchainte-sdk-go/pkg/blockchain"
	"github.com/enchainte/enchainte-sdk-go/pkg/crypto"
	"github.com/enchainte/enchainte-sdk-go/pkg/http"
	"log"
)

func main() {
	configData := repository.NewConfigData()
	configRepo := repository.NewConfigRepositoryImpl(configData)
	configService := service.NewConfigServiceImpl(configRepo)
	web3 := blockchain2.NewWeb3(configService)

	timestamp, err := web3.ValidateRoot(entity.EthereumRinkeby, "c7afe76d6dabae68c10c32e5673ed20535ebb00436e615eccc208f14c0993744")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(timestamp)

}

type Services struct {
	Message    message.Service
	Proof      proof.Service
	Credential credential.Service
	Anchor     anchor.Service
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
		Anchor:     anchor.NewService(apiKey, http, azureService.SdkParameters()),
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