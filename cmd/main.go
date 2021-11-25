package main

import (
	"github.com/bloock/bloock-sdk-go/internal"
	entity2 "github.com/bloock/bloock-sdk-go/internal/config/entity"
	"github.com/bloock/bloock-sdk-go/internal/record/entity"
	"log"
)

func main() {

	type Data struct {
		data string
	}
	re := entity.FromObject(Data{data: "Example Data"})
	log.Println(re)
	re1 := entity.FromHash("5ac706bdef87529b22c08646b74cb98baf310a46bd21ee420814b04c71fa42b1")
	log.Println(re1)
	re2, err := entity.FromHex("123456789abcdefa")
	if err != nil {
		log.Println(err)
	}
	log.Println(re2)
	re3 := entity.FromString("Example Data")
	log.Println(re3)
	re4 := entity.FromUint8Array([]byte{
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1})
	log.Println(re4)


	apiKey := "test_xculO0olb1Itp-tFMNCjpsLgx4Bik3E7Wd-iUfdL1c2lsgyKvhAZQnd7U8vlPnJX"

	client := internal.NewBloockClient(apiKey)

	records := make([]entity.RecordEntity, 0)
	records = append(records, re4)
	records = append(records, re3)

	/*se, err := client.SendRecords(records)
	if err != nil {
		log.Println(err)
	}*/

	/*rr, err := client.WaitAnchor(se[3].Anchor, 60000)
	if err != nil {
		log.Println(err)
	}*/

	r, err := client.GetProof(records)
	if err != nil {
		log.Println(err)
	}
	log.Println(r)

	i, err := client.VerifyRecords(records, entity2.BloockChain)
	if err != nil {
		log.Println(err)
	}
	log.Println(i)
}