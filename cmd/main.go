package main

import (
	"encoding/json"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/config"
	"github.com/enchainte/enchainte-sdk-go/config/repository"
	"github.com/enchainte/enchainte-sdk-go/config/service"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor"
	"github.com/enchainte/enchainte-sdk-go/internal/cloud"
	"github.com/enchainte/enchainte-sdk-go/internal/credential"
	http2 "github.com/enchainte/enchainte-sdk-go/internal/infrastructure/http"
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity"
	repository2 "github.com/enchainte/enchainte-sdk-go/internal/record/repository"
	service2 "github.com/enchainte/enchainte-sdk-go/internal/record/service"
	"github.com/enchainte/enchainte-sdk-go/pkg/http"
	"log"
)

func main() {
	configData := repository.NewConfigData()
	configRepo := repository.NewConfigRepository(configData)
	configService := service.NewConfigService(configRepo)

	arrRecords := make([]entity.RecordEntity, 0)
	arrRecords = append(arrRecords, entity.NewRecordEntity("6a83f545cb5693a32b5d56fb4a0530f7054df0c7e2e6b0a9fef36e26a2a96b04"))
	arrRecords = append(arrRecords, entity.NewRecordEntity("2d9130eb0900a08f22dee5e0330672861e6035eb858e1d1ac0d0d5d98a676800"))
	arrRecords = append(arrRecords, entity.NewRecordEntity("cadc5a160b48bde5727b08e1f8d1b8fe08704ff3cc730bf4919a2ef10ae6e291"))
	arrRecords = append(arrRecords, entity.NewRecordEntity("db6d0af6e743ca02954f1feb7dec3033fe4f86d429b8dd5b7dd654b794d71dee"))

	httpData := http2.NewDataHttp("C1vfvhN2mPUeX0KikgGHVIUSofZIfX6Q4bx0kf7DuAHMt3cuELO2UGdYLUw9bS29")
	httpClient := http2.NewHttp(httpData)
	recordRepo := repository2.NewRecordRepository(httpClient, configService)
	recordService := service2.NewRecordService(recordRepo)

	resp, err := recordService.GetRecords(arrRecords)
	if err != nil {
		log.Println(err)
	}

	log.Println(resp)
}

type Services struct {
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

	return Services{
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