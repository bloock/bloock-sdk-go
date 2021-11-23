package internal

import (
	entity2 "github.com/enchainte/enchainte-sdk-go/internal/anchor/entity"
	repository2 "github.com/enchainte/enchainte-sdk-go/internal/anchor/repository"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor/service"
	entity4 "github.com/enchainte/enchainte-sdk-go/internal/config/entity"
	repository5 "github.com/enchainte/enchainte-sdk-go/internal/config/repository"
	service5 "github.com/enchainte/enchainte-sdk-go/internal/config/service"
	"github.com/enchainte/enchainte-sdk-go/internal/infrastructure"
	"github.com/enchainte/enchainte-sdk-go/internal/infrastructure/blockchain"
	"github.com/enchainte/enchainte-sdk-go/internal/infrastructure/http"
	entity3 "github.com/enchainte/enchainte-sdk-go/internal/proof/entity"
	repository4 "github.com/enchainte/enchainte-sdk-go/internal/proof/repository"
	service4 "github.com/enchainte/enchainte-sdk-go/internal/proof/service"
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity"
	repository3 "github.com/enchainte/enchainte-sdk-go/internal/record/repository"
	service3 "github.com/enchainte/enchainte-sdk-go/internal/record/service"
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
	configService service5.ConfigurerService
	recordService service3.RecorderService
	proofService  service4.ProoferService
	httpClient    infrastructure.HttpClient
}

/*
NewBloockClient
Constructor with API Key that enables accessing to Bloock's functionalities.
Parameters:
	{string} apiKey Client API Key.
*/
func NewBloockClient(apiKey string) BloockClient {
	c := repository5.NewConfigData()
	cr := repository5.NewConfigRepository(c)
	cs := service5.NewConfigService(&cr)

	d := http.NewDataHttp(apiKey)
	h := http.NewHttp(d)

	ar := repository2.NewAnchorRepository(h, &cs)
	as := service.NewAnchorService(ar, &cs)

	rr := repository3.NewRecordRepository(h, &cs)
	rs := service3.NewRecordService(rr)

	b := blockchain.NewWeb3(&cs)
	pr := repository4.NewProofRepository(h, b, &cs)
	ps := service4.NewProofService(pr)

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
func(b BloockClient) SetApiHost(host string) {
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
func (b BloockClient) SetNetworkConfiguration(network string, configuration entity4.NetworkConfiguration) {
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
func (b BloockClient) GetAnchor(anchor int) (entity2.Anchor, error) {
	return b.anchorService.GetAnchor(anchor)
}

/*
WaitAnchor
Waits until the anchor specified is confirmed in Bloock.
Parameters:
	{int} anchor Id of the Anchor to wait for.
	{int} [timeout=120000] Timeout time in miliseconds. After exceeding this time returns an exception.
Returns:
	{Anchor} Anchor object matching the id.
	{error} Error if something went wrong.
Errors:
	{AnchorNotFoundException} The anchor provided could not be found.
	{InvalidRecordException} At least one of the records sent was not well formated
	{HttpRequestException} Error return by Bloock's API.
	{error} Native Golang error
*/
func (b BloockClient) WaitAnchor(anchor int, timeout int) (entity2.Anchor, error) {
	return b.anchorService.WaitAnchor(anchor, timeout)
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
func (b BloockClient) GetProof(records []entity.RecordEntity) (entity3.Proof, error) {
	return b.proofService.RetrieveProof(records)
}

/*
VerifyProof
Verifies if the specified integrity Proof is valid and checks if it's currently included in the blockchain.
Parameters:
	{Proof} Proof to validate.
	{Network} blockchain network where the proof will be validated
Returns:
	{int} A number representing the timestamp in milliseconds when the anchor was registered in Blockchain
	{error} Error if something went wrong.
Errors:
	{Web3Exception} Error connecting to blockchain.
	{error} Native Golang error
*/
func (b BloockClient) VerifyProof(proof entity3.Proof, network string) (int, error) {
	return b.proofService.VerifyProof(proof, network)
}

/*
VerifyRecords
It retrieves a proof for the specified list of Anchor using getProof and verifies it using verifyProof.
Parameters:
	{[]Record} list of records to validate
	{Network} blockchain network where the records will be validated
Returns:
	{int} A number representing the timestamp in milliseconds when the anchor was registered in Blockchain
	{error} Error if something went wrong.
Errors:
	{InvalidRecordException} At least one of the records sent was not well formated
	{HttpRequestException} Error return by Bloock's API.
	{Web3Exception} Error connecting to blockchain.
	{error} Native Golang error
*/
func (b BloockClient) VerifyRecords(records []entity.RecordEntity, network string) (int, error) {
	return b.proofService.VerifyRecords(records, network)
}
