// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: anchor.proto

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

type Anchor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BlockRoots []string         `protobuf:"bytes,2,rep,name=block_roots,json=blockRoots,proto3" json:"block_roots,omitempty"`
	Networks   []*AnchorNetwork `protobuf:"bytes,3,rep,name=networks,proto3" json:"networks,omitempty"`
	Root       string           `protobuf:"bytes,4,opt,name=root,proto3" json:"root,omitempty"`
	Status     string           `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Anchor) Reset() {
	*x = Anchor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anchor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Anchor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Anchor) ProtoMessage() {}

func (x *Anchor) ProtoReflect() protoreflect.Message {
	mi := &file_anchor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Anchor.ProtoReflect.Descriptor instead.
func (*Anchor) Descriptor() ([]byte, []int) {
	return file_anchor_proto_rawDescGZIP(), []int{0}
}

func (x *Anchor) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Anchor) GetBlockRoots() []string {
	if x != nil {
		return x.BlockRoots
	}
	return nil
}

func (x *Anchor) GetNetworks() []*AnchorNetwork {
	if x != nil {
		return x.Networks
	}
	return nil
}

func (x *Anchor) GetRoot() string {
	if x != nil {
		return x.Root
	}
	return ""
}

func (x *Anchor) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type AnchorNetwork struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	State  string `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	TxHash string `protobuf:"bytes,3,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
}

func (x *AnchorNetwork) Reset() {
	*x = AnchorNetwork{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anchor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnchorNetwork) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnchorNetwork) ProtoMessage() {}

func (x *AnchorNetwork) ProtoReflect() protoreflect.Message {
	mi := &file_anchor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnchorNetwork.ProtoReflect.Descriptor instead.
func (*AnchorNetwork) Descriptor() ([]byte, []int) {
	return file_anchor_proto_rawDescGZIP(), []int{1}
}

func (x *AnchorNetwork) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AnchorNetwork) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *AnchorNetwork) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

type GetAnchorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigData *ConfigData `protobuf:"bytes,1,opt,name=config_data,json=configData,proto3" json:"config_data,omitempty"`
	AnchorId   int64       `protobuf:"varint,2,opt,name=anchor_id,json=anchorId,proto3" json:"anchor_id,omitempty"`
}

func (x *GetAnchorRequest) Reset() {
	*x = GetAnchorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anchor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnchorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnchorRequest) ProtoMessage() {}

func (x *GetAnchorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_anchor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnchorRequest.ProtoReflect.Descriptor instead.
func (*GetAnchorRequest) Descriptor() ([]byte, []int) {
	return file_anchor_proto_rawDescGZIP(), []int{2}
}

func (x *GetAnchorRequest) GetConfigData() *ConfigData {
	if x != nil {
		return x.ConfigData
	}
	return nil
}

func (x *GetAnchorRequest) GetAnchorId() int64 {
	if x != nil {
		return x.AnchorId
	}
	return 0
}

type GetAnchorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Anchor *Anchor `protobuf:"bytes,1,opt,name=anchor,proto3,oneof" json:"anchor,omitempty"`
	Error  *Error  `protobuf:"bytes,2,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *GetAnchorResponse) Reset() {
	*x = GetAnchorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anchor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnchorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnchorResponse) ProtoMessage() {}

func (x *GetAnchorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_anchor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnchorResponse.ProtoReflect.Descriptor instead.
func (*GetAnchorResponse) Descriptor() ([]byte, []int) {
	return file_anchor_proto_rawDescGZIP(), []int{3}
}

func (x *GetAnchorResponse) GetAnchor() *Anchor {
	if x != nil {
		return x.Anchor
	}
	return nil
}

func (x *GetAnchorResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type WaitAnchorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigData *ConfigData `protobuf:"bytes,1,opt,name=config_data,json=configData,proto3" json:"config_data,omitempty"`
	AnchorId   int64       `protobuf:"varint,2,opt,name=anchor_id,json=anchorId,proto3" json:"anchor_id,omitempty"`
	Timeout    int64       `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *WaitAnchorRequest) Reset() {
	*x = WaitAnchorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anchor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WaitAnchorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaitAnchorRequest) ProtoMessage() {}

func (x *WaitAnchorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_anchor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaitAnchorRequest.ProtoReflect.Descriptor instead.
func (*WaitAnchorRequest) Descriptor() ([]byte, []int) {
	return file_anchor_proto_rawDescGZIP(), []int{4}
}

func (x *WaitAnchorRequest) GetConfigData() *ConfigData {
	if x != nil {
		return x.ConfigData
	}
	return nil
}

func (x *WaitAnchorRequest) GetAnchorId() int64 {
	if x != nil {
		return x.AnchorId
	}
	return 0
}

func (x *WaitAnchorRequest) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type WaitAnchorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Anchor *Anchor `protobuf:"bytes,1,opt,name=anchor,proto3,oneof" json:"anchor,omitempty"`
	Error  *Error  `protobuf:"bytes,2,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *WaitAnchorResponse) Reset() {
	*x = WaitAnchorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anchor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WaitAnchorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaitAnchorResponse) ProtoMessage() {}

func (x *WaitAnchorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_anchor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaitAnchorResponse.ProtoReflect.Descriptor instead.
func (*WaitAnchorResponse) Descriptor() ([]byte, []int) {
	return file_anchor_proto_rawDescGZIP(), []int{5}
}

func (x *WaitAnchorResponse) GetAnchor() *Anchor {
	if x != nil {
		return x.Anchor
	}
	return nil
}

func (x *WaitAnchorResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_anchor_proto protoreflect.FileDescriptor

var file_anchor_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x62, 0x6c, 0x6f, 0x6f, 0x63, 0x6b, 0x1a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x06, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x6f, 0x6f, 0x74, 0x73, 0x12, 0x31,
	0x0a, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x62, 0x6c, 0x6f, 0x6f, 0x63, 0x6b, 0x2e, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x72, 0x6f, 0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x52, 0x0a,
	0x0d, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73,
	0x68, 0x22, 0x64, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6c, 0x6f,
	0x6f, 0x63, 0x6b, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6e,
	0x63, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61,
	0x6e, 0x63, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x7f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6e,
	0x63, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06,
	0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62,
	0x6c, 0x6f, 0x6f, 0x63, 0x6b, 0x2e, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x06,
	0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x6c, 0x6f, 0x6f, 0x63,
	0x6b, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x01, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x7f, 0x0a, 0x11, 0x57, 0x61, 0x69, 0x74,
	0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a,
	0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6c, 0x6f, 0x6f, 0x63, 0x6b, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x80, 0x01, 0x0a, 0x12, 0x57, 0x61,
	0x69, 0x74, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2b, 0x0a, 0x06, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x62, 0x6c, 0x6f, 0x6f, 0x63, 0x6b, 0x2e, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72,
	0x48, 0x00, 0x52, 0x06, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62,
	0x6c, 0x6f, 0x6f, 0x63, 0x6b, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x01, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x6e, 0x63, 0x68,
	0x6f, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x96, 0x01, 0x0a,
	0x0d, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x12, 0x18, 0x2e, 0x62, 0x6c,
	0x6f, 0x6f, 0x63, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x6c, 0x6f, 0x6f, 0x63, 0x6b, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x43, 0x0a, 0x0a, 0x57, 0x61, 0x69, 0x74, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x12, 0x19,
	0x2e, 0x62, 0x6c, 0x6f, 0x6f, 0x63, 0x6b, 0x2e, 0x57, 0x61, 0x69, 0x74, 0x41, 0x6e, 0x63, 0x68,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x6c, 0x6f, 0x6f,
	0x63, 0x6b, 0x2e, 0x57, 0x61, 0x69, 0x74, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x5c, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x6c, 0x6f,
	0x6f, 0x63, 0x6b, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x2e, 0x62, 0x72, 0x69,
	0x64, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6c, 0x6f, 0x6f, 0x63, 0x6b, 0x2f, 0x62, 0x6c, 0x6f,
	0x6f, 0x63, 0x6b, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_anchor_proto_rawDescOnce sync.Once
	file_anchor_proto_rawDescData = file_anchor_proto_rawDesc
)

func file_anchor_proto_rawDescGZIP() []byte {
	file_anchor_proto_rawDescOnce.Do(func() {
		file_anchor_proto_rawDescData = protoimpl.X.CompressGZIP(file_anchor_proto_rawDescData)
	})
	return file_anchor_proto_rawDescData
}

var file_anchor_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_anchor_proto_goTypes = []interface{}{
	(*Anchor)(nil),             // 0: bloock.Anchor
	(*AnchorNetwork)(nil),      // 1: bloock.AnchorNetwork
	(*GetAnchorRequest)(nil),   // 2: bloock.GetAnchorRequest
	(*GetAnchorResponse)(nil),  // 3: bloock.GetAnchorResponse
	(*WaitAnchorRequest)(nil),  // 4: bloock.WaitAnchorRequest
	(*WaitAnchorResponse)(nil), // 5: bloock.WaitAnchorResponse
	(*ConfigData)(nil),         // 6: bloock.ConfigData
	(*Error)(nil),              // 7: bloock.Error
}
var file_anchor_proto_depIdxs = []int32{
	1, // 0: bloock.Anchor.networks:type_name -> bloock.AnchorNetwork
	6, // 1: bloock.GetAnchorRequest.config_data:type_name -> bloock.ConfigData
	0, // 2: bloock.GetAnchorResponse.anchor:type_name -> bloock.Anchor
	7, // 3: bloock.GetAnchorResponse.error:type_name -> bloock.Error
	6, // 4: bloock.WaitAnchorRequest.config_data:type_name -> bloock.ConfigData
	0, // 5: bloock.WaitAnchorResponse.anchor:type_name -> bloock.Anchor
	7, // 6: bloock.WaitAnchorResponse.error:type_name -> bloock.Error
	2, // 7: bloock.AnchorService.GetAnchor:input_type -> bloock.GetAnchorRequest
	4, // 8: bloock.AnchorService.WaitAnchor:input_type -> bloock.WaitAnchorRequest
	3, // 9: bloock.AnchorService.GetAnchor:output_type -> bloock.GetAnchorResponse
	5, // 10: bloock.AnchorService.WaitAnchor:output_type -> bloock.WaitAnchorResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_anchor_proto_init() }
func file_anchor_proto_init() {
	if File_anchor_proto != nil {
		return
	}
	file_shared_proto_init()
	file_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_anchor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Anchor); i {
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
		file_anchor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnchorNetwork); i {
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
		file_anchor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnchorRequest); i {
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
		file_anchor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnchorResponse); i {
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
		file_anchor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WaitAnchorRequest); i {
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
		file_anchor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WaitAnchorResponse); i {
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
	file_anchor_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_anchor_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_anchor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_anchor_proto_goTypes,
		DependencyIndexes: file_anchor_proto_depIdxs,
		MessageInfos:      file_anchor_proto_msgTypes,
	}.Build()
	File_anchor_proto = out.File
	file_anchor_proto_rawDesc = nil
	file_anchor_proto_goTypes = nil
	file_anchor_proto_depIdxs = nil
}
