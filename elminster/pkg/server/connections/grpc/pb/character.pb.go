// The greeting service definition.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: pkg/server/connections/grpc/pb/character.proto

package pb

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

type Stats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Str int32 `protobuf:"varint,1,opt,name=str,proto3" json:"str,omitempty"`
	Dex int32 `protobuf:"varint,2,opt,name=dex,proto3" json:"dex,omitempty"`
	Con int32 `protobuf:"varint,3,opt,name=con,proto3" json:"con,omitempty"`
	Int int32 `protobuf:"varint,4,opt,name=int,proto3" json:"int,omitempty"`
	Cha int32 `protobuf:"varint,5,opt,name=cha,proto3" json:"cha,omitempty"`
	Wil int32 `protobuf:"varint,6,opt,name=wil,proto3" json:"wil,omitempty"`
}

func (x *Stats) Reset() {
	*x = Stats{}
	mi := &file_pkg_server_connections_grpc_pb_character_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stats) ProtoMessage() {}

func (x *Stats) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_server_connections_grpc_pb_character_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stats.ProtoReflect.Descriptor instead.
func (*Stats) Descriptor() ([]byte, []int) {
	return file_pkg_server_connections_grpc_pb_character_proto_rawDescGZIP(), []int{0}
}

func (x *Stats) GetStr() int32 {
	if x != nil {
		return x.Str
	}
	return 0
}

func (x *Stats) GetDex() int32 {
	if x != nil {
		return x.Dex
	}
	return 0
}

func (x *Stats) GetCon() int32 {
	if x != nil {
		return x.Con
	}
	return 0
}

func (x *Stats) GetInt() int32 {
	if x != nil {
		return x.Int
	}
	return 0
}

func (x *Stats) GetCha() int32 {
	if x != nil {
		return x.Cha
	}
	return 0
}

func (x *Stats) GetWil() int32 {
	if x != nil {
		return x.Wil
	}
	return 0
}

// The request message containing the user's name.
type CreateCharacterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Stats *Stats `protobuf:"bytes,2,opt,name=stats,proto3" json:"stats,omitempty"`
	Level int32  `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *CreateCharacterRequest) Reset() {
	*x = CreateCharacterRequest{}
	mi := &file_pkg_server_connections_grpc_pb_character_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCharacterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCharacterRequest) ProtoMessage() {}

func (x *CreateCharacterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_server_connections_grpc_pb_character_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCharacterRequest.ProtoReflect.Descriptor instead.
func (*CreateCharacterRequest) Descriptor() ([]byte, []int) {
	return file_pkg_server_connections_grpc_pb_character_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCharacterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCharacterRequest) GetStats() *Stats {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *CreateCharacterRequest) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

// The response message containing the greetings
type CreateCharacterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *CreateCharacterResponse) Reset() {
	*x = CreateCharacterResponse{}
	mi := &file_pkg_server_connections_grpc_pb_character_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCharacterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCharacterResponse) ProtoMessage() {}

func (x *CreateCharacterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_server_connections_grpc_pb_character_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCharacterResponse.ProtoReflect.Descriptor instead.
func (*CreateCharacterResponse) Descriptor() ([]byte, []int) {
	return file_pkg_server_connections_grpc_pb_character_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCharacterResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

// The request message containing the user's name.
type UpdateCharacterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Stats *Stats `protobuf:"bytes,3,opt,name=stats,proto3" json:"stats,omitempty"`
	Level int32  `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *UpdateCharacterRequest) Reset() {
	*x = UpdateCharacterRequest{}
	mi := &file_pkg_server_connections_grpc_pb_character_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCharacterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCharacterRequest) ProtoMessage() {}

func (x *UpdateCharacterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_server_connections_grpc_pb_character_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCharacterRequest.ProtoReflect.Descriptor instead.
func (*UpdateCharacterRequest) Descriptor() ([]byte, []int) {
	return file_pkg_server_connections_grpc_pb_character_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateCharacterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateCharacterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateCharacterRequest) GetStats() *Stats {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *UpdateCharacterRequest) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

// The response message containing the greetings
type UpdateCharacterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *UpdateCharacterResponse) Reset() {
	*x = UpdateCharacterResponse{}
	mi := &file_pkg_server_connections_grpc_pb_character_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCharacterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCharacterResponse) ProtoMessage() {}

func (x *UpdateCharacterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_server_connections_grpc_pb_character_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCharacterResponse.ProtoReflect.Descriptor instead.
func (*UpdateCharacterResponse) Descriptor() ([]byte, []int) {
	return file_pkg_server_connections_grpc_pb_character_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCharacterResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type GetCharacterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCharacterRequest) Reset() {
	*x = GetCharacterRequest{}
	mi := &file_pkg_server_connections_grpc_pb_character_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCharacterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCharacterRequest) ProtoMessage() {}

func (x *GetCharacterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_server_connections_grpc_pb_character_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCharacterRequest.ProtoReflect.Descriptor instead.
func (*GetCharacterRequest) Descriptor() ([]byte, []int) {
	return file_pkg_server_connections_grpc_pb_character_proto_rawDescGZIP(), []int{5}
}

func (x *GetCharacterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCharacterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Stats *Stats `protobuf:"bytes,2,opt,name=stats,proto3" json:"stats,omitempty"`
	Level int32  `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *GetCharacterResponse) Reset() {
	*x = GetCharacterResponse{}
	mi := &file_pkg_server_connections_grpc_pb_character_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCharacterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCharacterResponse) ProtoMessage() {}

func (x *GetCharacterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_server_connections_grpc_pb_character_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCharacterResponse.ProtoReflect.Descriptor instead.
func (*GetCharacterResponse) Descriptor() ([]byte, []int) {
	return file_pkg_server_connections_grpc_pb_character_proto_rawDescGZIP(), []int{6}
}

func (x *GetCharacterResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetCharacterResponse) GetStats() *Stats {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *GetCharacterResponse) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type DeleteCharacterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCharacterRequest) Reset() {
	*x = DeleteCharacterRequest{}
	mi := &file_pkg_server_connections_grpc_pb_character_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCharacterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCharacterRequest) ProtoMessage() {}

func (x *DeleteCharacterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_server_connections_grpc_pb_character_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCharacterRequest.ProtoReflect.Descriptor instead.
func (*DeleteCharacterRequest) Descriptor() ([]byte, []int) {
	return file_pkg_server_connections_grpc_pb_character_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteCharacterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteCharacterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *DeleteCharacterResponse) Reset() {
	*x = DeleteCharacterResponse{}
	mi := &file_pkg_server_connections_grpc_pb_character_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCharacterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCharacterResponse) ProtoMessage() {}

func (x *DeleteCharacterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_server_connections_grpc_pb_character_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCharacterResponse.ProtoReflect.Descriptor instead.
func (*DeleteCharacterResponse) Descriptor() ([]byte, []int) {
	return file_pkg_server_connections_grpc_pb_character_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteCharacterResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_pkg_server_connections_grpc_pb_character_proto protoreflect.FileDescriptor

var file_pkg_server_connections_grpc_pb_character_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62,
	0x2f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x73, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x74, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x64, 0x65,
	0x78, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x63, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x69, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x68, 0x61, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x63, 0x68, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x69, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x77, 0x69, 0x6c, 0x22, 0x63, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x35,
	0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x73, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x35, 0x0a, 0x17, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x61, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x28, 0x0a, 0x16, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xba, 0x02, 0x0a,
	0x09, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x0f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x70, 0x62,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0f, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x1a,
	0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x69, 0x74, 0x6f, 0x72, 0x73, 0x61, 0x76,
	0x69, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x72, 0x73, 0x6f, 0x6e, 0x67, 0x2d, 0x72, 0x65, 0x70, 0x6f,
	0x2f, 0x65, 0x6c, 0x6d, 0x69, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_server_connections_grpc_pb_character_proto_rawDescOnce sync.Once
	file_pkg_server_connections_grpc_pb_character_proto_rawDescData = file_pkg_server_connections_grpc_pb_character_proto_rawDesc
)

func file_pkg_server_connections_grpc_pb_character_proto_rawDescGZIP() []byte {
	file_pkg_server_connections_grpc_pb_character_proto_rawDescOnce.Do(func() {
		file_pkg_server_connections_grpc_pb_character_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_server_connections_grpc_pb_character_proto_rawDescData)
	})
	return file_pkg_server_connections_grpc_pb_character_proto_rawDescData
}

var file_pkg_server_connections_grpc_pb_character_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_server_connections_grpc_pb_character_proto_goTypes = []any{
	(*Stats)(nil),                   // 0: pb.Stats
	(*CreateCharacterRequest)(nil),  // 1: pb.CreateCharacterRequest
	(*CreateCharacterResponse)(nil), // 2: pb.CreateCharacterResponse
	(*UpdateCharacterRequest)(nil),  // 3: pb.UpdateCharacterRequest
	(*UpdateCharacterResponse)(nil), // 4: pb.UpdateCharacterResponse
	(*GetCharacterRequest)(nil),     // 5: pb.GetCharacterRequest
	(*GetCharacterResponse)(nil),    // 6: pb.GetCharacterResponse
	(*DeleteCharacterRequest)(nil),  // 7: pb.DeleteCharacterRequest
	(*DeleteCharacterResponse)(nil), // 8: pb.DeleteCharacterResponse
}
var file_pkg_server_connections_grpc_pb_character_proto_depIdxs = []int32{
	0, // 0: pb.CreateCharacterRequest.stats:type_name -> pb.Stats
	0, // 1: pb.UpdateCharacterRequest.stats:type_name -> pb.Stats
	0, // 2: pb.GetCharacterResponse.stats:type_name -> pb.Stats
	1, // 3: pb.Character.CreateCharacter:input_type -> pb.CreateCharacterRequest
	3, // 4: pb.Character.UpdateCharacter:input_type -> pb.UpdateCharacterRequest
	5, // 5: pb.Character.GetCharacter:input_type -> pb.GetCharacterRequest
	7, // 6: pb.Character.DeleteCharacter:input_type -> pb.DeleteCharacterRequest
	2, // 7: pb.Character.CreateCharacter:output_type -> pb.CreateCharacterResponse
	4, // 8: pb.Character.UpdateCharacter:output_type -> pb.UpdateCharacterResponse
	6, // 9: pb.Character.GetCharacter:output_type -> pb.GetCharacterResponse
	8, // 10: pb.Character.DeleteCharacter:output_type -> pb.DeleteCharacterResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_server_connections_grpc_pb_character_proto_init() }
func file_pkg_server_connections_grpc_pb_character_proto_init() {
	if File_pkg_server_connections_grpc_pb_character_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_server_connections_grpc_pb_character_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_server_connections_grpc_pb_character_proto_goTypes,
		DependencyIndexes: file_pkg_server_connections_grpc_pb_character_proto_depIdxs,
		MessageInfos:      file_pkg_server_connections_grpc_pb_character_proto_msgTypes,
	}.Build()
	File_pkg_server_connections_grpc_pb_character_proto = out.File
	file_pkg_server_connections_grpc_pb_character_proto_rawDesc = nil
	file_pkg_server_connections_grpc_pb_character_proto_goTypes = nil
	file_pkg_server_connections_grpc_pb_character_proto_depIdxs = nil
}
