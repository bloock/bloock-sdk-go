# Bloock SDK - Golang

This SDK offers all the features available in the Bloock Toolset:

- Write records
- Get records proof
- Validate proof
- Get records details

## Installation

The SDK can be installed with the go get command:

```shell
$ go get github.com/bloock-go-sdk
```

## Usage

The following examples summarize how to access the different functionalities available:

### Prepare data

In order to interact with the SDK, the data should be processed through the Hash module.

There are several ways to generate a Hash:

```go
anchor, err := client.GetAnchor(1)
	if err != nil {
		log.Println(err)
	}
```

### Send records

This example shows how to send data to Bloock

```golang
```

### Get records status

This example shows how to get all the details and status of records:

```golang
```

### Wait for records to process

This example shows how to wait for a record to be processed by Bloock after sending it.

```golang
```

### Get and validate records proof

This example shows how to get a proof for an array of records and validate it:

```golang
```

### Full example

This snippet shows a complete data cycle including: write, wait for record confirmation and proof retrieval and validation.

```golang
```


