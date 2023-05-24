// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: api/v1/worker.proto

package v1

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

type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerAddr string `protobuf:"bytes,1,opt,name=server_addr,json=serverAddr,proto3" json:"server_addr,omitempty"`
	ServerId   string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_worker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_worker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_worker_proto_rawDescGZIP(), []int{0}
}

func (x *JoinRequest) GetServerAddr() string {
	if x != nil {
		return x.ServerAddr
	}
	return ""
}

func (x *JoinRequest) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

type JoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JoinResponse) Reset() {
	*x = JoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_worker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinResponse) ProtoMessage() {}

func (x *JoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_worker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinResponse.ProtoReflect.Descriptor instead.
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_worker_proto_rawDescGZIP(), []int{1}
}

type InsertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *InsertRequest) Reset() {
	*x = InsertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_worker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertRequest) ProtoMessage() {}

func (x *InsertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_worker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertRequest.ProtoReflect.Descriptor instead.
func (*InsertRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_worker_proto_rawDescGZIP(), []int{2}
}

func (x *InsertRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *InsertRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type InsertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InsertResponse) Reset() {
	*x = InsertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_worker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertResponse) ProtoMessage() {}

func (x *InsertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_worker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertResponse.ProtoReflect.Descriptor instead.
func (*InsertResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_worker_proto_rawDescGZIP(), []int{3}
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_worker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_worker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_worker_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_worker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_worker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_worker_proto_rawDescGZIP(), []int{5}
}

type FetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *FetchRequest) Reset() {
	*x = FetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_worker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRequest) ProtoMessage() {}

func (x *FetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_worker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRequest.ProtoReflect.Descriptor instead.
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_worker_proto_rawDescGZIP(), []int{6}
}

func (x *FetchRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type FetchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *FetchResponse) Reset() {
	*x = FetchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_worker_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchResponse) ProtoMessage() {}

func (x *FetchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_worker_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchResponse.ProtoReflect.Descriptor instead.
func (*FetchResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_worker_proto_rawDescGZIP(), []int{7}
}

func (x *FetchResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *FetchResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type RaftStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RaftStateRequest) Reset() {
	*x = RaftStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_worker_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftStateRequest) ProtoMessage() {}

func (x *RaftStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_worker_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftStateRequest.ProtoReflect.Descriptor instead.
func (*RaftStateRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_worker_proto_rawDescGZIP(), []int{8}
}

type RaftStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State string `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *RaftStateResponse) Reset() {
	*x = RaftStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_worker_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftStateResponse) ProtoMessage() {}

func (x *RaftStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_worker_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftStateResponse.ProtoReflect.Descriptor instead.
func (*RaftStateResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_worker_proto_rawDescGZIP(), []int{9}
}

func (x *RaftStateResponse) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type LeaderHeaderbeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RaftState string `protobuf:"bytes,1,opt,name=raft_state,json=raftState,proto3" json:"raft_state,omitempty"`
	GrpcAddr  string `protobuf:"bytes,2,opt,name=grpc_addr,json=grpcAddr,proto3" json:"grpc_addr,omitempty"`
	RaftAddr  string `protobuf:"bytes,3,opt,name=raft_addr,json=raftAddr,proto3" json:"raft_addr,omitempty"`
}

func (x *LeaderHeaderbeat) Reset() {
	*x = LeaderHeaderbeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_worker_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderHeaderbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderHeaderbeat) ProtoMessage() {}

func (x *LeaderHeaderbeat) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_worker_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderHeaderbeat.ProtoReflect.Descriptor instead.
func (*LeaderHeaderbeat) Descriptor() ([]byte, []int) {
	return file_api_v1_worker_proto_rawDescGZIP(), []int{10}
}

func (x *LeaderHeaderbeat) GetRaftState() string {
	if x != nil {
		return x.RaftState
	}
	return ""
}

func (x *LeaderHeaderbeat) GetGrpcAddr() string {
	if x != nil {
		return x.GrpcAddr
	}
	return ""
}

func (x *LeaderHeaderbeat) GetRaftAddr() string {
	if x != nil {
		return x.RaftAddr
	}
	return ""
}

var File_api_v1_worker_proto protoreflect.FileDescriptor

var file_api_v1_worker_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x22, 0x4b, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x0e, 0x0a,
	0x0c, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x37, 0x0a,
	0x0d, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x10, 0x0a, 0x0e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x0a,
	0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22,
	0x37, 0x0a, 0x0d, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x52, 0x61, 0x66, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x29, 0x0a, 0x11,
	0x52, 0x61, 0x66, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x6b, 0x0a, 0x10, 0x4c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x65, 0x61, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x61, 0x66, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x61, 0x66, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67,
	0x72, 0x70, 0x63, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x61, 0x66, 0x74, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x61, 0x66, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x32, 0xc3, 0x02, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12,
	0x37, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x12, 0x18, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x18, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12,
	0x17, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x52, 0x61, 0x66, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x1b, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x66, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x66, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x7a, 0x61, 0x61, 0x6b, 0x64, 0x61,
	0x6c, 0x65, 0x2f, 0x64, 0x69, 0x6e, 0x67, 0x68, 0x79, 0x2d, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_worker_proto_rawDescOnce sync.Once
	file_api_v1_worker_proto_rawDescData = file_api_v1_worker_proto_rawDesc
)

func file_api_v1_worker_proto_rawDescGZIP() []byte {
	file_api_v1_worker_proto_rawDescOnce.Do(func() {
		file_api_v1_worker_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_worker_proto_rawDescData)
	})
	return file_api_v1_worker_proto_rawDescData
}

var file_api_v1_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_v1_worker_proto_goTypes = []interface{}{
	(*JoinRequest)(nil),       // 0: worker.v1.JoinRequest
	(*JoinResponse)(nil),      // 1: worker.v1.JoinResponse
	(*InsertRequest)(nil),     // 2: worker.v1.InsertRequest
	(*InsertResponse)(nil),    // 3: worker.v1.InsertResponse
	(*DeleteRequest)(nil),     // 4: worker.v1.DeleteRequest
	(*DeleteResponse)(nil),    // 5: worker.v1.DeleteResponse
	(*FetchRequest)(nil),      // 6: worker.v1.FetchRequest
	(*FetchResponse)(nil),     // 7: worker.v1.FetchResponse
	(*RaftStateRequest)(nil),  // 8: worker.v1.RaftStateRequest
	(*RaftStateResponse)(nil), // 9: worker.v1.RaftStateResponse
	(*LeaderHeaderbeat)(nil),  // 10: worker.v1.LeaderHeaderbeat
}
var file_api_v1_worker_proto_depIdxs = []int32{
	0, // 0: worker.v1.Worker.Join:input_type -> worker.v1.JoinRequest
	2, // 1: worker.v1.Worker.Insert:input_type -> worker.v1.InsertRequest
	4, // 2: worker.v1.Worker.Delete:input_type -> worker.v1.DeleteRequest
	6, // 3: worker.v1.Worker.Fetch:input_type -> worker.v1.FetchRequest
	8, // 4: worker.v1.Worker.RaftState:input_type -> worker.v1.RaftStateRequest
	1, // 5: worker.v1.Worker.Join:output_type -> worker.v1.JoinResponse
	3, // 6: worker.v1.Worker.Insert:output_type -> worker.v1.InsertResponse
	5, // 7: worker.v1.Worker.Delete:output_type -> worker.v1.DeleteResponse
	7, // 8: worker.v1.Worker.Fetch:output_type -> worker.v1.FetchResponse
	9, // 9: worker.v1.Worker.RaftState:output_type -> worker.v1.RaftStateResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_worker_proto_init() }
func file_api_v1_worker_proto_init() {
	if File_api_v1_worker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_worker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRequest); i {
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
		file_api_v1_worker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinResponse); i {
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
		file_api_v1_worker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertRequest); i {
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
		file_api_v1_worker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertResponse); i {
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
		file_api_v1_worker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_api_v1_worker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_api_v1_worker_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRequest); i {
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
		file_api_v1_worker_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchResponse); i {
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
		file_api_v1_worker_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftStateRequest); i {
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
		file_api_v1_worker_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftStateResponse); i {
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
		file_api_v1_worker_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderHeaderbeat); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_worker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_worker_proto_goTypes,
		DependencyIndexes: file_api_v1_worker_proto_depIdxs,
		MessageInfos:      file_api_v1_worker_proto_msgTypes,
	}.Build()
	File_api_v1_worker_proto = out.File
	file_api_v1_worker_proto_rawDesc = nil
	file_api_v1_worker_proto_goTypes = nil
	file_api_v1_worker_proto_depIdxs = nil
}
