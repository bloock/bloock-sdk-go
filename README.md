# Bloock SDK - Golang

This SDK offers all the features available in the Bloock Toolset:

- Write records
- Get records proof
- Validate proof
- Get records details

## Installation

The SDK can be installed with the go get command:

```shell
$ go get github.com/bloock/bloock-sdk-go
```

## Usage

The following examples summarize how to access the different functionalities available:

### Prepare data

In order to interact with the SDK, the data should be processed through the Hash module.

There are several ways to generate a Hash:

```go
import (
    "github.com/bloock/bloock-sdk-go"
    "log"
)

// From an object
type Data struct {
    data string
}
record := bloock.NewRecordFromObject(Data{data: "Example Data"})

// From a hash string (hex encoded 64-chars long string)
record2 := bloock.NewRecordFromHash("5ac706bdef87529b22c08646b74cb98baf310a46bd21ee420814b04c71fa42b1")

// From a hex encoded string
record3, err := bloock.NewRecordFromHex("123456789abcdefa")
if err != nil {
    log.Println(err)
}

// From a string
record4 := bloock.NewRecordFromString("Example Data")

// From a Uint8Array with a lenght of 32
record5 := bloock.NewRecordFromUint8Array([]byte{
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
```

### Send records

This example shows how to send data to Bloock

```go
import (
    "github.com/bloock/bloock-sdk-go"
    "log"
    "os"
)

apiKey := os.Getenv("API_KEY")

client := bloock.NewClient(apiKey)

record := bloock.NewRecordFromString("Example Data 1")
records := make([]bloock.Record, 0)
records = append(records, record)

r, err := client.SendRecords(records)
if err != nil {
	log.Println(err)
}
log.Println(r)
```

### Get records status

This example shows how to get all the details and status of records:

```go
import (
    "github.com/bloock/bloock-sdk-go"
    "log"
    "os"
)

apiKey := os.Getenv("API_KEY")

client := bloock.NewClient(apiKey)

record := bloock.NewRecordFromString("Example Data 1")
record2 := bloock.NewRecordFromString("Example Data 2")
record3 := bloock.NewRecordFromString("Example Data 3")
records := make([]bloock.Record, 0)
records = append(records, record)
records = append(records, record2)
records = append(records, record3)

r, err := client.GetRecords(records)
if err != nil {
    log.Println(err)
}
log.Println(r)
```

### Wait for records to process

This example shows how to wait for a record to be processed by Bloock after sending it.

```go
import (
    "github.com/bloock/bloock-sdk-go"
    "log"
    "os"
)

apiKey := os.Getenv("API_KEY")

client := bloock.NewClient(apiKey)

record := bloock.NewRecordFromString("Example Data 1")
records := make([]bloock.Record, 0)
records = append(records, record)

r, err := client.SendRecords(records)
if err != nil {
    log.Println(err)
}

_, err := client.WaitAnchor(r[0].Anchor, bloock.NewAnchorParams())
if err != nil {
    log.Println(err)
}
```

Snippet to set other timeout, by default it's set 120000.

```go
// By default --> timeout = 120000
anchorParams := bloock.NewAnchorParams()
anchorParams.Timeout = 3000
_, err := client.WaitAnchor(r[0].Anchor, anchorParams)

```

### Get and validate records proof

This example shows how to get a proof for an array of records and validate it:

```go
import (
    "github.com/bloock/bloock-sdk-go"
    "log"
    "os"
)

apiKey := os.Getenv("API_KEY")

client := bloock.NewClient(apiKey)

record := bloock.NewRecordFromString("Example Data 1")
record2 := bloock.NewRecordFromString("Example Data 2")
record3 := bloock.NewRecordFromString("Example Data 3")
records := make([]bloock.Record, 0)
records = append(records, record)
records = append(records, record2)
records = append(records, record3)

p, err := client.GetProof(records)
if err != nil {
    log.Println(err)
}

timestamp, err := client.VerifyProof(p, bloock.NewNetworkParams())
if err != nil {
    log.Println(err)
}
log.Println(timestamp)
```

Snippet to set another Network, by default it's set Ethereum Mainnet.

```go
params := bloock.NewNetworkParams()
list := bloock.ListOfNetworks()
params.Network = list.EthereumRinkeby

timestamp, err := sdk.VerifyProof(proof, params)
```

### Full example

This snippet shows a complete data cycle including: write, wait for record confirmation and proof retrieval and validation.

```go
import (
    "github.com/bloock/bloock-sdk-go"
    "log"
    "math"
    "math/rand"
    "os"
    "strconv"
)

// Helper function to get a random string
func randomString(n int) string {
    var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

    s := make([]rune, n)
    for i := range s {
        s[i] = letters[rand.Intn(len(letters))]
    }
    return string(s)
}

func main() {
    apiKey := os.Getenv("API_KEY")
    client := bloock.NewClient(apiKey)

    record := bloock.NewRecordFromString(randomString(64))
    records := make([]bloock.Record, 0)
    records = append(records, record)

    r, err := client.SendRecords(records)
    if err != nil {
        log.Println(err)
    }
    log.Println("Write record - Successful!")
	
    if r[0].Record == "" && r[0].Status == "" {
        os.Exit(1)
    }
	
    // Default timeout: 120000
    _, err = client.WaitAnchor(r[0].Anchor, bloock.NewAnchorParams())
    if err != nil {
        log.Println(err)
    }
    log.Println("Record reached Blockchain!")

    // Retrieving record proof 
    proof, err := client.GetProof(records)
    if err != nil {
        log.Println(err)
    }
	
    // Default: EthereumMainnet 
    timestamp, err := client.VerifyProof(proof, bloock.NewNetworkParams())
    if err != nil {
        log.Println(err)
    }

    if timestamp != 0 {
        log.Printf("Record is valid - Timestamp: %d", timestamp)
    } else {
        log.Println("Record is invalid")
    }
}
```


