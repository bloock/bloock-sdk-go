// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: availability.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigData *ConfigData `protobuf:"bytes,1,opt,name=config_data,json=configData,proto3" json:"config_data,omitempty"`
	Record     *Record     `protobuf:"bytes,2,opt,name=record,proto3" json:"record,omitempty"`
	Publisher  *Publisher  `protobuf:"bytes,3,opt,name=publisher,proto3" json:"publisher,omitempty"`
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_availability_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_availability_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_availability_proto_rawDescGZIP(), []int{0}
}

func (x *PublishRequest) GetConfigData() *ConfigData {
	if x != nil {
		return x.ConfigData
	}
	return nil
}

func (x *PublishRequest) GetRecord() *Record {
	if x != nil {
		return x.Record
	}
	return nil
}

func (x *PublishRequest) GetPublisher() *Publisher {
	if x != nil {
		return x.Publisher
	}
	return nil
}

type PublishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Error *Error `protobuf:"bytes,2,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *PublishResponse) Reset() {
	*x = PublishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_availability_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishResponse) ProtoMessage() {}

func (x *PublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_availability_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishResponse.ProtoReflect.Descriptor instead.
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return file_availability_proto_rawDescGZIP(), []int{1}
}

func (x *PublishResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PublishResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type RetrieveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigData *ConfigData `protobuf:"bytes,1,opt,name=config_data,json=configData,proto3" json:"config_data,omitempty"`
	Loader     *Loader     `protobuf:"bytes,3,opt,name=loader,proto3" json:"loader,omitempty"`
}

func (x *RetrieveRequest) Reset() {
	*x = RetrieveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_availability_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveRequest) ProtoMessage() {}

func (x *RetrieveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_availability_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveRequest.ProtoReflect.Descriptor instead.
func (*RetrieveRequest) Descriptor() ([]byte, []int) {
	return file_availability_proto_rawDescGZIP(), []int{2}
}

func (x *RetrieveRequest) GetConfigData() *ConfigData {
	if x != nil {
		return x.ConfigData
	}
	return nil
}

func (x *RetrieveRequest) GetLoader() *Loader {
	if x != nil {
		return x.Loader
	}
	return nil
}

type RetrieveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record *Record `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	Error  *Error  `protobuf:"bytes,2,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *RetrieveResponse) Reset() {
	*x = RetrieveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_availability_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveResponse) ProtoMessage() {}

func (x *RetrieveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_availability_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveResponse.ProtoReflect.Descriptor instead.
func (*RetrieveResponse) Descriptor() ([]byte, []int) {
	return file_availability_proto_rawDescGZIP(), []int{3}
}

func (x *RetrieveResponse) GetRecord() *Record {
	if x != nil {
		return x.Record
	}
	return nil
}

func (x *RetrieveResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_availability_proto protoreflect.FileDescriptor

var file_availability_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x62, 0x6c, 0x6f, 0x6f, 0x63, 0x6b, 0x1a, 0x1b, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a,
	0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x33, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6c, 0x6f, 0x6f, 0x63, 0x6b, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x6c, 0x6f, 0x6f, 0x63, 0x6b, 0x2e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x2f, 0x0a, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x62, 0x6c, 0x6f, 0x6f, 0x63, 0x6b, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x72, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x22, 0x55, 0x0a,
	0x0f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x28, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x62, 0x6c, 0x6f, 0x6f, 0x63, 0x6b, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x6e, 0x0a, 0x0f, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62,
	0x6c, 0x6f, 0x6f, 0x63, 0x6b, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x06,
	0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62,
	0x6c, 0x6f, 0x6f, 0x63, 0x6b, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6c, 0x6f,
	0x61, 0x64, 0x65, 0x72, 0x22, 0x6e, 0x0a, 0x10, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x6c, 0x6f, 0x6f, 0x63,
	0x6b, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x28, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x62, 0x6c, 0x6f, 0x6f, 0x63, 0x6b, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x32, 0x90, 0x01, 0x0a, 0x13, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x07,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x16, 0x2e, 0x62, 0x6c, 0x6f, 0x6f, 0x63, 0x6b,
	0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x6f, 0x63, 0x6b, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x52, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x76, 0x65, 0x12, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x6f, 0x63, 0x6b, 0x2e, 0x52, 0x65,
	0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x62, 0x6c, 0x6f, 0x6f, 0x63, 0x6b, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x57, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x62,
	0x6c, 0x6f, 0x6f, 0x63, 0x6b, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x62, 0x6c, 0x6f, 0x6f, 0x63, 0x6b, 0x2f, 0x62, 0x6c, 0x6f, 0x6f, 0x63, 0x6b,
	0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_availability_proto_rawDescOnce sync.Once
	file_availability_proto_rawDescData = file_availability_proto_rawDesc
)

func file_availability_proto_rawDescGZIP() []byte {
	file_availability_proto_rawDescOnce.Do(func() {
		file_availability_proto_rawDescData = protoimpl.X.CompressGZIP(file_availability_proto_rawDescData)
	})
	return file_availability_proto_rawDescData
}

var file_availability_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_availability_proto_goTypes = []interface{}{
	(*PublishRequest)(nil),   // 0: bloock.PublishRequest
	(*PublishResponse)(nil),  // 1: bloock.PublishResponse
	(*RetrieveRequest)(nil),  // 2: bloock.RetrieveRequest
	(*RetrieveResponse)(nil), // 3: bloock.RetrieveResponse
	(*ConfigData)(nil),       // 4: bloock.ConfigData
	(*Record)(nil),           // 5: bloock.Record
	(*Publisher)(nil),        // 6: bloock.Publisher
	(*Error)(nil),            // 7: bloock.Error
	(*Loader)(nil),           // 8: bloock.Loader
}
var file_availability_proto_depIdxs = []int32{
	4,  // 0: bloock.PublishRequest.config_data:type_name -> bloock.ConfigData
	5,  // 1: bloock.PublishRequest.record:type_name -> bloock.Record
	6,  // 2: bloock.PublishRequest.publisher:type_name -> bloock.Publisher
	7,  // 3: bloock.PublishResponse.error:type_name -> bloock.Error
	4,  // 4: bloock.RetrieveRequest.config_data:type_name -> bloock.ConfigData
	8,  // 5: bloock.RetrieveRequest.loader:type_name -> bloock.Loader
	5,  // 6: bloock.RetrieveResponse.record:type_name -> bloock.Record
	7,  // 7: bloock.RetrieveResponse.error:type_name -> bloock.Error
	0,  // 8: bloock.AvailabilityService.Publish:input_type -> bloock.PublishRequest
	2,  // 9: bloock.AvailabilityService.Retrieve:input_type -> bloock.RetrieveRequest
	1,  // 10: bloock.AvailabilityService.Publish:output_type -> bloock.PublishResponse
	3,  // 11: bloock.AvailabilityService.Retrieve:output_type -> bloock.RetrieveResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_availability_proto_init() }
func file_availability_proto_init() {
	if File_availability_proto != nil {
		return
	}
	file_availability_entities_proto_init()
	file_record_entities_proto_init()
	file_shared_proto_init()
	file_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_availability_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_availability_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_availability_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_availability_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_availability_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_availability_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_availability_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_availability_proto_goTypes,
		DependencyIndexes: file_availability_proto_depIdxs,
		MessageInfos:      file_availability_proto_msgTypes,
	}.Build()
	File_availability_proto = out.File
	file_availability_proto_rawDesc = nil
	file_availability_proto_goTypes = nil
	file_availability_proto_depIdxs = nil
}
