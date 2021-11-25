package main

import (
	"github.com/enchainte/enchainte-sdk-go/internal"
	entity2 "github.com/enchainte/enchainte-sdk-go/internal/config/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity"
	"log"
)

func main() {

	type Data struct {
		data string
	}

	re := entity.FromObject(Data{data: "Example Data"})

	apiKey := "test_xculO0olb1Itp-tFMNCjpsLgx4Bik3E7Wd-iUfdL1c2lsgyKvhAZQnd7U8vlPnJX"

	client := internal.NewBloockClient(apiKey)

	records := make([]entity.RecordEntity, 0)
	records = append(records, re)

	rr, err := client.SendRecords(records)
	if err != nil {
		log.Println(err)
	}
	log.Println(rr)

	a, err := client.WaitAnchor(rr[0].Anchor, 120000)
	if err != nil {
		log.Println(err)
	}
	log.Println(a)

	r, err := client.GetAnchor(rr[0].Anchor)
	if err != nil {
		log.Println(err)
	}
	log.Println(r)

	p, err := client.GetProof(records)
	if err != nil {
		log.Println(err)
	}
	i, err := client.VerifyProof(p, entity2.BloockChain)
	if err != nil {
		log.Println(err)
	}
	log.Println(i)
}