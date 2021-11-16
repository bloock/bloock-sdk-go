package main

import (
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity"
	"log"
	"sort"
	"strings"
)

func main() {
	/*configData := repository.NewConfigData()
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

	log.Println(resp)*/
	arrRecords := make([]entity.RecordEntity, 0)
	arrRecords = append(arrRecords, entity.NewRecordEntity("6a83f545cb5693a32b5d56fb4a0530f7054df0c7e2e6b0a9fef36e26a2a96b04"))
	arrRecords = append(arrRecords, entity.NewRecordEntity("2d9130eb0900a08f22dee5e0330672861e6035eb858e1d1ac0d0d5d98a676800"))
	arrRecords = append(arrRecords, entity.NewRecordEntity("cadc5a160b48bde5727b08e1f8d1b8fe08704ff3cc730bf4919a2ef10ae6e291"))
	arrRecords = append(arrRecords, entity.NewRecordEntity("db6d0af6e743ca02954f1feb7dec3033fe4f86d429b8dd5b7dd654b794d71dee"))

	sort.SliceStable(arrRecords, func(i, j int) bool {
		return strings.ToUpper(arrRecords[i].GetHash()) < strings.ToUpper(arrRecords[j].GetHash())
	})

	log.Println(arrRecords)
}