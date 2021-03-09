package main

import (
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/config"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor"
	"github.com/enchainte/enchainte-sdk-go/internal/cloud"
	"github.com/enchainte/enchainte-sdk-go/internal/credential"
	"github.com/enchainte/enchainte-sdk-go/internal/message"
	"github.com/enchainte/enchainte-sdk-go/internal/proof"
	"github.com/enchainte/enchainte-sdk-go/pkg/blockchain"
	"github.com/enchainte/enchainte-sdk-go/pkg/crypto"
	"github.com/enchainte/enchainte-sdk-go/pkg/http"
	"log"
	"math/rand"
	"time"
)

func main() {
	cli := EnchainteClient("")
	//cli.Message.Write([]byte("helloo"))

	go dataCollector(cli.Message.Write)
	go func() {
		var count int
		for {
			fmt.Println(message.Receive())
			count++
			if count == 4 {
				message.Done()
				return
			}
		}
	}()
	fmt.Scanln()
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

// TODO remove
func dataCollector(write func([]byte) error) {
	for {
		c := time.Tick(time.Second)
		for range c {
			write([]byte(words[rand.Intn(len(words)-1)]))
		}
	}
}

var words = []string{"Lebuffic", "Caming", "Unizans", "Nantilien", "Losychle", "Deping", "Subsce", "Shemon", "Unhyle", "Reighthes"}
