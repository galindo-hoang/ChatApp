// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: app/chat_app/v1/relationship.proto

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

type CreateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId   uint64 `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	AccountName string `protobuf:"bytes,2,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	mi := &file_app_chat_app_v1_relationship_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_chat_app_v1_relationship_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_app_chat_app_v1_relationship_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAccountRequest) GetAccountId() uint64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *CreateAccountRequest) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *CreateAccountRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type MakeRelationshipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From uint64 `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	To   uint64 `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *MakeRelationshipRequest) Reset() {
	*x = MakeRelationshipRequest{}
	mi := &file_app_chat_app_v1_relationship_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MakeRelationshipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeRelationshipRequest) ProtoMessage() {}

func (x *MakeRelationshipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_chat_app_v1_relationship_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeRelationshipRequest.ProtoReflect.Descriptor instead.
func (*MakeRelationshipRequest) Descriptor() ([]byte, []int) {
	return file_app_chat_app_v1_relationship_proto_rawDescGZIP(), []int{1}
}

func (x *MakeRelationshipRequest) GetFrom() uint64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *MakeRelationshipRequest) GetTo() uint64 {
	if x != nil {
		return x.To
	}
	return 0
}

type MakeRelationshipResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool   `protobuf:"varint,1,opt,name=isSuccess,proto3" json:"isSuccess,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MakeRelationshipResponse) Reset() {
	*x = MakeRelationshipResponse{}
	mi := &file_app_chat_app_v1_relationship_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MakeRelationshipResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeRelationshipResponse) ProtoMessage() {}

func (x *MakeRelationshipResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_chat_app_v1_relationship_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeRelationshipResponse.ProtoReflect.Descriptor instead.
func (*MakeRelationshipResponse) Descriptor() ([]byte, []int) {
	return file_app_chat_app_v1_relationship_proto_rawDescGZIP(), []int{2}
}

func (x *MakeRelationshipResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *MakeRelationshipResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool   `protobuf:"varint,1,opt,name=isSuccess,proto3" json:"isSuccess,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateAccountResponse) Reset() {
	*x = CreateAccountResponse{}
	mi := &file_app_chat_app_v1_relationship_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountResponse) ProtoMessage() {}

func (x *CreateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_chat_app_v1_relationship_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountResponse.ProtoReflect.Descriptor instead.
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return file_app_chat_app_v1_relationship_proto_rawDescGZIP(), []int{3}
}

func (x *CreateAccountResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *CreateAccountResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_app_chat_app_v1_relationship_proto protoreflect.FileDescriptor

var file_app_chat_app_v1_relationship_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x76,
	0x31, 0x22, 0x6e, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x22, 0x3d, 0x0a, 0x17, 0x4d, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x74, 0x6f,
	0x22, 0x52, 0x0a, 0x18, 0x4d, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x4f, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xb2, 0x02, 0x0a, 0x13, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x0f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e,
	0x67, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x61,
	0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x6b, 0x65,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x11, 0x55, 0x6e, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x69, 0x6e, 0x67, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_chat_app_v1_relationship_proto_rawDescOnce sync.Once
	file_app_chat_app_v1_relationship_proto_rawDescData = file_app_chat_app_v1_relationship_proto_rawDesc
)

func file_app_chat_app_v1_relationship_proto_rawDescGZIP() []byte {
	file_app_chat_app_v1_relationship_proto_rawDescOnce.Do(func() {
		file_app_chat_app_v1_relationship_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_chat_app_v1_relationship_proto_rawDescData)
	})
	return file_app_chat_app_v1_relationship_proto_rawDescData
}

var file_app_chat_app_v1_relationship_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_app_chat_app_v1_relationship_proto_goTypes = []any{
	(*CreateAccountRequest)(nil),     // 0: chat_app.v1.CreateAccountRequest
	(*MakeRelationshipRequest)(nil),  // 1: chat_app.v1.MakeRelationshipRequest
	(*MakeRelationshipResponse)(nil), // 2: chat_app.v1.MakeRelationshipResponse
	(*CreateAccountResponse)(nil),    // 3: chat_app.v1.CreateAccountResponse
}
var file_app_chat_app_v1_relationship_proto_depIdxs = []int32{
	0, // 0: chat_app.v1.RelationshipService.CreateNode:input_type -> chat_app.v1.CreateAccountRequest
	1, // 1: chat_app.v1.RelationshipService.FollowingPerson:input_type -> chat_app.v1.MakeRelationshipRequest
	1, // 2: chat_app.v1.RelationshipService.UnFollowingPerson:input_type -> chat_app.v1.MakeRelationshipRequest
	3, // 3: chat_app.v1.RelationshipService.CreateNode:output_type -> chat_app.v1.CreateAccountResponse
	2, // 4: chat_app.v1.RelationshipService.FollowingPerson:output_type -> chat_app.v1.MakeRelationshipResponse
	2, // 5: chat_app.v1.RelationshipService.UnFollowingPerson:output_type -> chat_app.v1.MakeRelationshipResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_chat_app_v1_relationship_proto_init() }
func file_app_chat_app_v1_relationship_proto_init() {
	if File_app_chat_app_v1_relationship_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_chat_app_v1_relationship_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_chat_app_v1_relationship_proto_goTypes,
		DependencyIndexes: file_app_chat_app_v1_relationship_proto_depIdxs,
		MessageInfos:      file_app_chat_app_v1_relationship_proto_msgTypes,
	}.Build()
	File_app_chat_app_v1_relationship_proto = out.File
	file_app_chat_app_v1_relationship_proto_rawDesc = nil
	file_app_chat_app_v1_relationship_proto_goTypes = nil
	file_app_chat_app_v1_relationship_proto_depIdxs = nil
}
