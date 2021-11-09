package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/config"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor"
	"github.com/enchainte/enchainte-sdk-go/internal/cloud"
	"github.com/enchainte/enchainte-sdk-go/internal/credential"
	http2 "github.com/enchainte/enchainte-sdk-go/internal/infrastructure/http"
	"github.com/enchainte/enchainte-sdk-go/internal/infrastructure/http/exception"
	"github.com/enchainte/enchainte-sdk-go/internal/message"
	"github.com/enchainte/enchainte-sdk-go/internal/proof"
	"github.com/enchainte/enchainte-sdk-go/pkg/blockchain"
	"github.com/enchainte/enchainte-sdk-go/pkg/crypto"
	"github.com/enchainte/enchainte-sdk-go/pkg/http"
	"log"
)

func main() {
	/*configData := repository.NewConfigData()
	configRepo := repository.NewConfigRepository(configData)
	configService := service.NewConfigService(configRepo)
	web3 := blockchain2.NewWeb3(configService)

	timestamp, err := web3.ValidateRoot(entity.EthereumRinkeby, "c7afe76d6dabae68c10c32e5673ed20535ebb00436e615eccc208f14c0993744")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(timestamp)*/

	httpData := http2.NewDataHttp("0oVmAgtI2sWLOkh9Y2P-cuSAd5Qkd3kNTxDiED76J5otLBeRE31DHDpWBYCg1xkU")
	httpClient := http2.NewHttp(httpData)

	var header = make(map[string]string)
	header["Content-Type"] = "application/json"
	response, err := httpClient.Get("https://api.bloock.dev/core/anchor/1", nil)
	//values := map[string][]string{"messages": {"00f0a8794e3f93fe4ac828e0f80aa4ddaa7ee0ca345af5ca130f9c3699a6b84d", "010604029843ade2399779ecacecd316180f6f790c77ba230e1fa1a25568f133"}}

	//response, err := httpClient.Post("https://api.bloock.dev/core/proof", values, header)
	if err != nil {
		if errors.Is(err, exception.HttpRequestException{}) {
			log.Fatal(err.Error())
		} else {
			log.Fatal(err.Error())
		}
	}
	fmt.Println(response)



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