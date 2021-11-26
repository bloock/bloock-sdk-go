package main

import (
	"github.com/bloock/bloock-sdk-go"
	"github.com/bloock/bloock-sdk-go/internal/record/entity"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	type Data struct {
		data string
	}
	re := bloock.NewRecordFromObject(Data{data: "Example Data"})
	log.Println(re)
	re1 := bloock.NewRecordFromHash("5ac706bdef87529b22c08646b74cb98baf310a46bd21ee420814b04c71fa42b1")
	log.Println(re1)
	re2, err := bloock.NewRecordFromHex("123456789abcdefa")
	if err != nil {
		log.Println(err)
	}
	log.Println(re2)
	re3 := bloock.NewRecordFromString("Example Data")
	log.Println(re3)
	re4 := bloock.NewRecordFromUint8Array([]byte{
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


	/*apiKey := "test_xculO0olb1Itp-tFMNCjpsLgx4Bik3E7Wd-iUfdL1c2lsgyKvhAZQnd7U8vlPnJX"

	sdk := internal.NewBloockClient(apiKey)

	records := make([]entity.RecordEntity, 0)
	records = append(records, re4)
	records = append(records, re3)*/



	apiKey := os.Getenv("API_KEY")
	sdk := bloock.NewBloockClient(apiKey)

	record := entity.FromString(randHex(64))
	records := make([]bloock.Record, 0)
	records = append(records, record)

	r, err := sdk.SendRecords(records)
	if err != nil {
		log.Println(err)
	}
	log.Println("Write record - Successful!")

	if r[0].Record == "" && r[0].Status == "" {
		os.Exit(1)
	}

	a := bloock.NewAnchorParams()
	a.Timeout = 3000
	_, err = sdk.WaitAnchor(r[0].Anchor, a)
	if err != nil {
		log.Println(err)
	}
	log.Println("Record reached Blockchain!")

	// Retrieving record proof
	proof, err := sdk.GetProof(records)
	if err != nil {
		log.Println(err)
	}

	n := bloock.NewNetworkParams()
	n.Network = bloock.BloockChain
	timestamp, err := sdk.VerifyProof(proof, n)
	if err != nil {
		log.Println(err)
	}

	if timestamp != 0 {
		log.Printf("Record is valid - Timestamp: %d", timestamp)
	} else {
		log.Println("Record is invalid")
	}
}

func randHex(length int) string {
	maxlength := 8
	min := math.Pow(16, math.Min(float64(length), float64(maxlength))-1)
	max := math.Pow(16, math.Min(float64(length), float64(maxlength))) - 1
	n := int((rand.Float64() * (max - min + 1)) + min)
	r := strconv.Itoa(n)
	for len(r) < length {
		r += randHex(length - maxlength)
	}
	return r
}