package bloock

import (
	anchorEntity "github.com/bloock/bloock-sdk-go/internal/anchor/entity"
	anchorRepository "github.com/bloock/bloock-sdk-go/internal/anchor/repository"
	"github.com/bloock/bloock-sdk-go/internal/anchor/service"
	configEntity "github.com/bloock/bloock-sdk-go/internal/config/entity"
	configRepository "github.com/bloock/bloock-sdk-go/internal/config/repository"
	configService "github.com/bloock/bloock-sdk-go/internal/config/service"
	"github.com/bloock/bloock-sdk-go/internal/infrastructure"
	"github.com/bloock/bloock-sdk-go/internal/infrastructure/blockchain"
	"github.com/bloock/bloock-sdk-go/internal/infrastructure/http"
	proofEntity "github.com/bloock/bloock-sdk-go/internal/proof/entity"
	proofRepository "github.com/bloock/bloock-sdk-go/internal/proof/repository"
	proofService "github.com/bloock/bloock-sdk-go/internal/proof/service"
	"github.com/bloock/bloock-sdk-go/internal/record/entity"
	recordRepository "github.com/bloock/bloock-sdk-go/internal/record/repository"
	recordService "github.com/bloock/bloock-sdk-go/internal/record/service"
)

/*
BloockClient
Entrypoint to the Bloock SDK:
	This SDK offers all the features available in the Bloock Toolset:
		- Write records
		- Get records proof
		- Validate proof
		- Get records details
*/
type BloockClient struct {
	anchorService service.AnchorerService
	configService configService.ConfigurerService
	recordService recordService.RecorderService
	proofService  proofService.ProoferService
	httpClient    infrastructure.HttpClient
}

/*
NewBloockClient
Constructor with API Key that enables accessing to Bloock's functionalities.
Parameters:
	{string} apiKey Client API Key.
*/
func NewBloockClient(apiKey string) BloockClient {
	c := configRepository.NewConfigData()
	cr := configRepository.NewConfigRepository(c)
	cs := configService.NewConfigService(&cr)

	d := http.NewDataHttp(apiKey)
	h := http.NewHttp(d)

	ar := anchorRepository.NewAnchorRepository(h, &cs)
	as := service.NewAnchorService(ar, &cs)

	rr := recordRepository.NewRecordRepository(h, &cs)
	rs := recordService.NewRecordService(rr)

	b := blockchain.NewWeb3(&cs)
	pr := proofRepository.NewProofRepository(h, b, &cs)
	ps := proofService.NewProofService(pr)

	return BloockClient{
		configService: &cs,
		httpClient:    h,
		anchorService: as,
		recordService: rs,
		proofService:  ps,
	}
}

/*
SetApiHost
Overrides the API host.
Parameters:
	{string} The API host to apply
Returns:
	{void}
*/
func (b BloockClient) SetApiHost(host string) {
	b.configService.SetApiHost(host)
}

/*
SetNetworkConfiguration
Overrides the Network configuration.
Parameters:
	{string} The Network to apply
	{NetworkConfiguration} The contract address, contract ABI and http provider to override the network specified
Returns:
	{void}
*/
func (b BloockClient) SetNetworkConfiguration(network string, configuration configEntity.NetworkConfiguration) {
	b.configService.SetNetworkConfiguration(network, configuration)
}

/*
SendRecords
Sends a list of Record to Bloock.
Parameters:
	{[]Record} records List of Record to send.
Returns:
	{[]RecordReceipt} List of RecordReceipt of each Record sent.
	{error} Error if something went wrong.
Errors:
	{InvalidRecordException} At least one of the records sent was not well formated
	{HttpRequestException} Error return by Bloock's API.
	{error} Native Golang error
*/
func (b BloockClient) SendRecords(records []entity.RecordEntity) ([]entity.RecordReceipt, error) {
	return b.recordService.SendRecords(records)
}

/*
GetRecords
Retrieves all RecordReceipt for the specified Anchor.
Parameters:
	{[]Record} records List of Record to fetch.
Returns:
	{[]RecordReceipt} List with the RecordReceipt of each record requested.
	{error} Error if something went wrong.
Errors:
	{InvalidRecordException} At least one of the records sent was not well formated
	{HttpRequestException} Error return by Bloock's API.
	{error} Native Golang error
*/
func (b BloockClient) GetRecords(records []entity.RecordEntity) ([]entity.RecordReceipt, error) {
	return b.recordService.GetRecords(records)
}

/*
GetAnchor
Gets an specific anchor id details.
Parameters:
	{int} anchor Id of the Anchor to look for.
Returns:
	{Anchor} Anchor object matching the id.
	{error} Error if something went wrong.
Errors:
	{InvalidRecordException} At least one of the records sent was not well formated
	{HttpRequestException} Error return by Bloock's API.
	{error} Native Golang error
*/
func (b BloockClient) GetAnchor(anchor int) (anchorEntity.Anchor, error) {
	return b.anchorService.GetAnchor(anchor)
}

/*
WaitAnchor
Waits until the anchor specified is confirmed in Bloock.
Parameters:
	{int} anchor Id of the Anchor to wait for.
	{AnchorParams} Timeout time in miliseconds. After exceeding this time returns an exception. Default = 120000
Returns:
	{Anchor} Anchor object matching the id.
	{error} Error if something went wrong.
Errors:
	{AnchorNotFoundException} The anchor provided could not be found.
	{InvalidRecordException} At least one of the records sent was not well formated
	{HttpRequestException} Error return by Bloock's API.
	{error} Native Golang error
*/
func (b BloockClient) WaitAnchor(anchor int, params anchorEntity.AnchorParams) (anchorEntity.Anchor, error) {
	return b.anchorService.WaitAnchor(anchor, params)
}

/*
GetProof
Retrieves an integrity Proof for the specified list of Record.
Parameters:
	{[]Record} List of records to validate.
Returns:
	{Proof} The Proof object containing the elements necessary to verify the integrity of the records in the input list
	If no record was requested, then returns None.
	{error} Error if something went wrong.
Errors:
	{InvalidRecordException} At least one of the records sent was not well formated
	{HttpRequestException} Error return by Bloock's API.
	{error} Native Golang error
*/
func (b BloockClient) GetProof(records []entity.RecordEntity) (proofEntity.Proof, error) {
	return b.proofService.RetrieveProof(records)
}

/*
VerifyProof
Verifies if the specified integrity Proof is valid and checks if it's currently included in the blockchain.
Parameters:
	{Proof} Proof to validate.
	{NetworkParams} blockchain network where the proof will be validated. Default: EthereumMainnet
Returns:
	{int} A number representing the timestamp in milliseconds when the anchor was registered in Blockchain
	{error} Error if something went wrong.
Errors:
	{Web3Exception} Error connecting to blockchain.
	{error} Native Golang error
*/
func (b BloockClient) VerifyProof(proof proofEntity.Proof, params configEntity.NetworkParams) (int, error) {
	return b.proofService.VerifyProof(proof, params)
}

/*
VerifyRecords
It retrieves a proof for the specified list of Anchor using getProof and verifies it using verifyProof.
Parameters:
	{[]Record} list of records to validate
	{NetworkParams} blockchain network where the records will be validated. Default: EthereumMainnet
Returns:
	{int} A number representing the timestamp in milliseconds when the anchor was registered in Blockchain
	{error} Error if something went wrong.
Errors:
	{InvalidRecordException} At least one of the records sent was not well formated
	{HttpRequestException} Error return by Bloock's API.
	{Web3Exception} Error connecting to blockchain.
	{error} Native Golang error
*/
func (b BloockClient) VerifyRecords(records []entity.RecordEntity, params configEntity.NetworkParams) (int, error) {
	return b.proofService.VerifyRecords(records, params)
}

/*
NewRecordFromObject
Given an JSON object, returns a Record with its value hashed.
Parameters:
	{interface{}} any type of data
Returns:
	{RecordEntity} RecordEntity object of the hashed input.
*/
func NewRecordFromObject(data interface{}) entity.RecordEntity {
	return entity.FromObject(data)
}

/*
NewRecordFromHash
It converts string to a Record hash.
Parameters:
	{string} Hexadecimal string without prefix and length 64.
Returns:
	{RecordEntity} RecordEntity object of the hashed input.
*/
func NewRecordFromHash(hash string) entity.RecordEntity {
	return entity.FromHash(hash)
}

/*
NewRecordFromHex
Given a hexadecimal string (with no 0x prefix) returns a Record with its value hashed.
Parameters:
	{string} Hexadecimal string without prefix.
Returns:
	{RecordEntity} RecordEntity object of the hashed input.
	{error} any type of error when hashing or converting
*/
func NewRecordFromHex(hex string) (entity.RecordEntity, error) {
	return entity.FromHex(hex)
}

/*
NewRecordFromString
Given a string returns a Record with its value hashed.
Parameters:
	{string} String object.
Returns:
	{RecordEntity} RecordEntity object of the hashed input.
*/
func NewRecordFromString(string string) entity.RecordEntity {
	return entity.FromString(string)
}

/*
NewRecordFromUint8Array
Given a bytes object returns a Record with its value hashed.
Parameters:
	{[]byte} Bytes object.
Returns:
	{RecordEntity} RecordEntity object of the hashed input.
*/
func NewRecordFromUint8Array(array []byte) entity.RecordEntity {
	return entity.FromUint8Array(array)
}
