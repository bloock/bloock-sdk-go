package client

import (
	"context"
	"errors"

	"github.com/bloock/bloock-sdk-go/v2/entity/authenticity"
	"github.com/bloock/bloock-sdk-go/v2/entity/availability"
	"github.com/bloock/bloock-sdk-go/v2/entity/encryption"
	"github.com/bloock/bloock-sdk-go/v2/entity/record"
	"github.com/bloock/bloock-sdk-go/v2/internal/bridge"
	"github.com/bloock/bloock-sdk-go/v2/internal/bridge/proto"
	"github.com/bloock/bloock-sdk-go/v2/internal/config"
)

type RecordClient struct {
	bridgeClient bridge.BloockBridge
	configData   *proto.ConfigData
}

func NewRecordClient() RecordClient {
	return RecordClient{
		bridgeClient: bridge.NewBloockBridge(),
		configData:   config.NewConfigDataDefault(),
	}
}

func NewRecordClientWithConfig(configData *proto.ConfigData) RecordClient {
	return RecordClient{
		bridgeClient: bridge.NewBloockBridge(),
		configData:   configData,
	}
}

func (c RecordClient) FromRecord(record record.Record) RecordBuilder {
	return newRecordBuilder(record.ToProto(), proto.RecordTypes_RECORD, c.configData)
}

func (c RecordClient) FromLoader(loader availability.Loader) RecordBuilder {
	return newRecordBuilder(loader.ToProto(), proto.RecordTypes_LOADER, c.configData)
}

func (c RecordClient) FromString(str string) RecordBuilder {
	return newRecordBuilder(str, proto.RecordTypes_STRING, c.configData)
}

func (c RecordClient) FromHex(hex string) RecordBuilder {
	return newRecordBuilder(hex, proto.RecordTypes_HEX, c.configData)
}

func (c RecordClient) FromJSON(json string) RecordBuilder {
	return newRecordBuilder(json, proto.RecordTypes_JSON, c.configData)
}

func (c RecordClient) FromFile(file_bytes []byte) RecordBuilder {
	return newRecordBuilder(file_bytes, proto.RecordTypes_FILE, c.configData)
}

func (c RecordClient) FromBytes(bytes []byte) RecordBuilder {
	return newRecordBuilder(bytes, proto.RecordTypes_BYTES, c.configData)
}

type RecordBuilder struct {
	payload     interface{}
	payloadType proto.RecordTypes
	configData  *proto.ConfigData
	signer      *proto.Signer
	encrypter   *proto.Encrypter
	decrypter   *proto.Encrypter
}

func newRecordBuilder(payload interface{}, payloadType proto.RecordTypes, configData *proto.ConfigData) RecordBuilder {
	return RecordBuilder{
		payload:     payload,
		payloadType: payloadType,
		configData:  configData,
	}
}

func (b RecordBuilder) WithSigner(signer authenticity.Signer) RecordBuilder {
	b.signer = signer.ToProto()
	return b
}

func (b RecordBuilder) WithEncrypter(encrypter encryption.Encrypter) RecordBuilder {
	b.encrypter = encrypter.ToProto()
	return b
}

func (b RecordBuilder) WithDecrypter(decrypter encryption.Encrypter) RecordBuilder {
	b.decrypter = decrypter.ToProto()
	return b
}

func (b RecordBuilder) Build() (record.Record, error) {
	bridgeClient := bridge.NewBloockBridge()

	var res *proto.RecordBuilderResponse
	var err error

	switch b.payloadType {
	case proto.RecordTypes_STRING:
		res, err = bridgeClient.Record().BuildRecordFromString(context.Background(), &proto.RecordBuilderFromStringRequest{
			ConfigData: b.configData,
			Payload:    b.payload.(string),
			Signer:     b.signer,
			Encrypter:  b.encrypter,
			Decrypter:  b.decrypter,
		})
	case proto.RecordTypes_BYTES:
		res, err = bridgeClient.Record().BuildRecordFromBytes(context.Background(), &proto.RecordBuilderFromBytesRequest{
			ConfigData: b.configData,
			Payload:    b.payload.([]byte),
			Signer:     b.signer,
			Encrypter:  b.encrypter,
			Decrypter:  b.decrypter,
		})
	case proto.RecordTypes_FILE:
		res, err = bridgeClient.Record().BuildRecordFromFile(context.Background(), &proto.RecordBuilderFromFileRequest{
			ConfigData: b.configData,
			Payload:    b.payload.([]byte),
			Signer:     b.signer,
			Encrypter:  b.encrypter,
			Decrypter:  b.decrypter,
		})
	case proto.RecordTypes_JSON:
		res, err = bridgeClient.Record().BuildRecordFromJson(context.Background(), &proto.RecordBuilderFromJSONRequest{
			ConfigData: b.configData,
			Payload:    b.payload.(string),
			Signer:     b.signer,
			Encrypter:  b.encrypter,
			Decrypter:  b.decrypter,
		})
	case proto.RecordTypes_HEX:
		res, err = bridgeClient.Record().BuildRecordFromHex(context.Background(), &proto.RecordBuilderFromHexRequest{
			ConfigData: b.configData,
			Payload:    b.payload.(string),
			Signer:     b.signer,
			Encrypter:  b.encrypter,
			Decrypter:  b.decrypter,
		})
	case proto.RecordTypes_RECORD:
		res, err = bridgeClient.Record().BuildRecordFromRecord(context.Background(), &proto.RecordBuilderFromRecordRequest{
			ConfigData: b.configData,
			Payload:    b.payload.(*proto.Record),
			Signer:     b.signer,
			Encrypter:  b.encrypter,
			Decrypter:  b.decrypter,
		})

	case proto.RecordTypes_LOADER:
		res, err = bridgeClient.Record().BuildRecordFromLoader(context.Background(), &proto.RecordBuilderFromLoaderRequest{
			ConfigData: b.configData,
			Loader:     b.payload.(*proto.Loader),
			Signer:     b.signer,
			Encrypter:  b.encrypter,
			Decrypter:  b.decrypter,
		})
	}

	if err != nil {
		return record.Record{}, err
	}

	if res.Error != nil {
		return record.Record{}, errors.New(res.Error.Message)
	}

	return record.NewRecordFromProto(res.Record, b.configData), nil
}

func (b RecordBuilder) GetDetails() (record.RecordDetails, error) {
	bridgeClient := bridge.NewBloockBridge()

	res, err := bridgeClient.Record().GetDetails(context.Background(), &proto.GetDetailsRequest{
		ConfigData: b.configData,
		Payload:    b.payload.([]byte),
	})

	if err != nil {
		return record.RecordDetails{}, err
	}

	if res.Error != nil {
		return record.RecordDetails{}, errors.New(res.Error.Message)
	}

	return record.NewRecordDetailsFromProto(res.Details, b.configData), nil
}
