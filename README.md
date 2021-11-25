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
    "github.com/bloock/bloock-sdk-go/internal/record/entity"
    "log"
)

// From an object
type Data struct {
    data string
}
record := entity.FromObject(Data{data: "Example Data"})

// From a hash string (hex encoded 64-chars long string)
record2 := entity.FromHash("5ac706bdef87529b22c08646b74cb98baf310a46bd21ee420814b04c71fa42b1")

// From a hex encoded string
record2, err := entity.FromHex("123456789abcdefa")
if err != nil {
    log.Println(err)
}

// From a string
record3 := entity.FromString("Example Data")

// From a Uint8Array with a lenght of 32
record4 := entity.FromUint8Array([]byte{
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
    "github.com/bloock/bloock-sdk-go/internal"
    "github.com/bloock/bloock-sdk-go/internal/record/entity"
    "log"
    "os"
)


apiKey := os.Getenv("API_KEY")

client := internal.NewBloockClient(apiKey)

record := entity.FromString("Example Data 1")
records := make([]entity.RecordEntity, 0)
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
    "github.com/bloock/bloock-sdk-go/internal"
    "github.com/bloock/bloock-sdk-go/internal/record/entity"
    "log"
    "os"
)


apiKey := os.Getenv("API_KEY")

client := internal.NewBloockClient(apiKey)

r := entity.FromString("Example Data 1")
r2 := entity.FromString("Example Data 2")
r3 := entity.FromString("Example Data 3")
records := make([]entity.RecordEntity, 0)
records = append(records, r)
records = append(records, r2)
records = append(records, r3)

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
    "github.com/bloock/bloock-sdk-go/internal"
    "github.com/bloock/bloock-sdk-go/internal/record/entity"
    "log"
    "os"
)


apiKey := os.Getenv("API_KEY")

client := internal.NewBloockClient(apiKey)

r := entity.FromString("Example Data 1")

r, err := client.SendRecords(records)
if err != nil {
    log.Println(err)
}

_, err := client.WaitAnchor(r[0].Anchor, 120000)
if err != nil {
    log.Println(err)
}
```

### Get and validate records proof

This example shows how to get a proof for an array of records and validate it:

```go
import (
    "github.com/bloock/bloock-sdk-go/internal"
    "github.com/bloock/bloock-sdk-go/internal/record/entity"
    configEntity "github.com/bloock/bloock-sdk-go/internal/config/entity"
    "log"
    "os"
)


apiKey := os.Getenv("API_KEY")

client := internal.NewBloockClient(apiKey)

r := entity.FromString("Example Data 1")
r2 := entity.FromString("Example Data 2")
r3 := entity.FromString("Example Data 3")
records := make([]entity.RecordEntity, 0)
records = append(records, r)
records = append(records, r2)
records = append(records, r3)

p, err := client.GetProof(records)
if err != nil {
    log.Println(err)
}

timestamp, err := client.VerifyProof(p, configEntity.EthereumMainnet)
if err != nil {
    log.Println(err)
}
log.Println(timestamp)
```

### Full example

This snippet shows a complete data cycle including: write, wait for record confirmation and proof retrieval and validation.

```go

// Helper function to get a random hex string
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

func main() {
    apiKey := os.Getenv("API_KEY")
    client := internal.NewBloockClient(apiKey)

    r := entity.FromString(randHex(64))
    records := make([]entity.RecordEntity, 0)
    records = append(records, r)
	
	
	
}

```


