package main

import (
	repository3 "github.com/enchainte/enchainte-sdk-go/internal/anchor/repository"
	service3 "github.com/enchainte/enchainte-sdk-go/internal/anchor/service"
	"github.com/enchainte/enchainte-sdk-go/internal/config/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/config/repository"
	"github.com/enchainte/enchainte-sdk-go/internal/config/service"
	"github.com/enchainte/enchainte-sdk-go/internal/infrastructure/blockchain"
	"github.com/enchainte/enchainte-sdk-go/internal/infrastructure/http"
	repository2 "github.com/enchainte/enchainte-sdk-go/internal/proof/repository"
	service2 "github.com/enchainte/enchainte-sdk-go/internal/proof/service"
	entity2 "github.com/enchainte/enchainte-sdk-go/internal/record/entity"
	"log"
)

func main() {
	configData := repository.NewConfigData()
	configRepo := repository.NewConfigRepository(configData)
	configService := service.NewConfigService(&configRepo)

	arrRecords := make([]entity2.RecordEntity, 0)
	arrRecords = append(arrRecords, entity2.NewRecordEntity("6a83f545cb5693a32b5d56fb4a0530f7054df0c7e2e6b0a9fef36e26a2a96b04"))

	httpData := http.NewDataHttp("test_xculO0olb1Itp-tFMNCjpsLgx4Bik3E7Wd-iUfdL1c2lsgyKvhAZQnd7U8vlPnJX")
	httpClient := http.NewHttp(httpData)

	anchorRepo := repository3.NewAnchorRepository(httpClient, &configService)
	anchorService := service3.NewAnchorService(anchorRepo, &configService)

	bkClient := blockchain.NewWeb3(&configService)

	pr := repository2.NewProofRepository(httpClient, bkClient, &configService)
	ps := service2.NewProofService(pr)

	resp, err := anchorService.GetAnchor(58)
	if err != nil {
		log.Println(err)
	}

	log.Println(resp)

	p, err := ps.VerifyRecords(arrRecords, entity.BloockChain)
	if err != nil {
		log.Println(err)
	}
	log.Println(p)


	/*arrRecords := make([]entity.RecordEntity, 0)
	arrRecords = append(arrRecords, entity.NewRecordEntity("6a83f545cb5693a32b5d56fb4a0530f7054df0c7e2e6b0a9fef36e26a2a96b04"))

	configData := repository3.NewConfigData()
	configRepo := repository3.NewConfigRepository(configData)
	configService := service.NewConfigService(&configRepo)

	httpData := http.NewDataHttp("C1vfvhN2mPUeX0KikgGHVIUSofZIfX6Q4bx0kf7DuAHMt3cuELO2UGdYLUw9bS29")
	httpClient := http.NewHttp(httpData)

	bkClient := blockchain.NewWeb3(configService)

	pr := repository2.NewProofRepository(httpClient, bkClient, configService)
	ps := service2.NewProofService(pr)

	p, err := ps.VerifyRecords(arrRecords, entity3.EthereumRinkeby)
	if err != nil {
		log.Println(err)
	}

	log.Println(p)*/

	/*leaves := []string{
		"02aae7e86eb50f61a62083a320475d9d60cbd52749dbf08fa942b1b97f50aee5",
		"5e1712aca5f3925fc0ce628e7da2e1e407e2cc7b358e83a7152b1958f7982dab",
	}
	nodes := []string{
		"1ca0e9d9a206f08d38a4e2cf485351674ffc9b0f3175e0cb6dbd8e0e19829b97",
		"1ca0e9d9a206f08d38a4e2cf485351674ffc9b0f3175e0cb6dbd8e0e19829b97",
		"54944fcea707a57048c17ca7453fa5078a031143b44629776750e7f0ff7940f0",
		"d6f9bcd042be70b39b65dc2a8168858606b0a2fcf6d02c0a1812b1804efc0c37",
		"e663ec001b81b96eceabd1b766d49ec5d99adedc3e5f03d245b0d90f603f66d3",
	}
	depth := "0004000400030004000400030001"
	bitmap := "7600"

	re, err := pr.VerifyProof(entity3.NewProof(leaves, nodes, depth, bitmap))
	log.Println(re.GetHash())

	r, err := bkClient.ValidateRoot(entity2.EthereumRinkeby, "ad2e170610dafa6bc90a8b1a19210a75e940f23f7ee932280d691b896cf1199f")
	if err != nil {
		log.Println(err)
	}
	log.Println(r)*/
}

/*func randHex(length int) string {
	maxlen := 8
	min := math.Pow(16, math.Min(float64(length), float64(maxlen)) - 1)
	max := math.Pow(16, math.Min(float64(length), float64(maxlen))) - 1
	n := int((rand.Float64() * (max - min + 1)) + min)
	r := strconv.Itoa(n)
	for len(r) < length {
		r += randHex(length - maxlen)
	}
	return r
}*/