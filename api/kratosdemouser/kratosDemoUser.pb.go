// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: api/kratosdemouser/kratosDemoUser.proto

package kratosdemouser

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type PingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *PingReq) Reset() {
	*x = PingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReq) ProtoMessage() {}

func (x *PingReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReq.ProtoReflect.Descriptor instead.
func (*PingReq) Descriptor() ([]byte, []int) {
	return file_api_kratosdemouser_kratosDemoUser_proto_rawDescGZIP(), []int{0}
}

func (x *PingReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type PingRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *PingRsp) Reset() {
	*x = PingRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRsp) ProtoMessage() {}

func (x *PingRsp) ProtoReflect() protoreflect.Message {
	mi := &file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRsp.ProtoReflect.Descriptor instead.
func (*PingRsp) Descriptor() ([]byte, []int) {
	return file_api_kratosdemouser_kratosDemoUser_proto_rawDescGZIP(), []int{1}
}

func (x *PingRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // 用户名
	Id   int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`    // 用户ID
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_api_kratosdemouser_kratosDemoUser_proto_rawDescGZIP(), []int{2}
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CreateKratosDemoUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // 用户名
}

func (x *CreateKratosDemoUserRequest) Reset() {
	*x = CreateKratosDemoUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateKratosDemoUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateKratosDemoUserRequest) ProtoMessage() {}

func (x *CreateKratosDemoUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateKratosDemoUserRequest.ProtoReflect.Descriptor instead.
func (*CreateKratosDemoUserRequest) Descriptor() ([]byte, []int) {
	return file_api_kratosdemouser_kratosDemoUser_proto_rawDescGZIP(), []int{3}
}

func (x *CreateKratosDemoUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateKratosDemoUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo *UserInfo `protobuf:"bytes,1,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"` // 用户信息
}

func (x *CreateKratosDemoUserReply) Reset() {
	*x = CreateKratosDemoUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateKratosDemoUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateKratosDemoUserReply) ProtoMessage() {}

func (x *CreateKratosDemoUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateKratosDemoUserReply.ProtoReflect.Descriptor instead.
func (*CreateKratosDemoUserReply) Descriptor() ([]byte, []int) {
	return file_api_kratosdemouser_kratosDemoUser_proto_rawDescGZIP(), []int{4}
}

func (x *CreateKratosDemoUserReply) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type UpdateKratosDemoUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo *UserInfo `protobuf:"bytes,1,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"` // 用户信息
}

func (x *UpdateKratosDemoUserRequest) Reset() {
	*x = UpdateKratosDemoUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateKratosDemoUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateKratosDemoUserRequest) ProtoMessage() {}

func (x *UpdateKratosDemoUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateKratosDemoUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateKratosDemoUserRequest) Descriptor() ([]byte, []int) {
	return file_api_kratosdemouser_kratosDemoUser_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateKratosDemoUserRequest) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type UpdateKratosDemoUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateKratosDemoUserReply) Reset() {
	*x = UpdateKratosDemoUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateKratosDemoUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateKratosDemoUserReply) ProtoMessage() {}

func (x *UpdateKratosDemoUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateKratosDemoUserReply.ProtoReflect.Descriptor instead.
func (*UpdateKratosDemoUserReply) Descriptor() ([]byte, []int) {
	return file_api_kratosdemouser_kratosDemoUser_proto_rawDescGZIP(), []int{6}
}

type DeleteKratosDemoUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // 用户ID
}

func (x *DeleteKratosDemoUserRequest) Reset() {
	*x = DeleteKratosDemoUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteKratosDemoUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteKratosDemoUserRequest) ProtoMessage() {}

func (x *DeleteKratosDemoUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteKratosDemoUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteKratosDemoUserRequest) Descriptor() ([]byte, []int) {
	return file_api_kratosdemouser_kratosDemoUser_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteKratosDemoUserRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteKratosDemoUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteKratosDemoUserReply) Reset() {
	*x = DeleteKratosDemoUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteKratosDemoUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteKratosDemoUserReply) ProtoMessage() {}

func (x *DeleteKratosDemoUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteKratosDemoUserReply.ProtoReflect.Descriptor instead.
func (*DeleteKratosDemoUserReply) Descriptor() ([]byte, []int) {
	return file_api_kratosdemouser_kratosDemoUser_proto_rawDescGZIP(), []int{8}
}

type GetKratosDemoUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // 用户ID
}

func (x *GetKratosDemoUserRequest) Reset() {
	*x = GetKratosDemoUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKratosDemoUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKratosDemoUserRequest) ProtoMessage() {}

func (x *GetKratosDemoUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKratosDemoUserRequest.ProtoReflect.Descriptor instead.
func (*GetKratosDemoUserRequest) Descriptor() ([]byte, []int) {
	return file_api_kratosdemouser_kratosDemoUser_proto_rawDescGZIP(), []int{9}
}

func (x *GetKratosDemoUserRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetKratosDemoUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo *UserInfo `protobuf:"bytes,1,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"` // 用户信息
}

func (x *GetKratosDemoUserReply) Reset() {
	*x = GetKratosDemoUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKratosDemoUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKratosDemoUserReply) ProtoMessage() {}

func (x *GetKratosDemoUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKratosDemoUserReply.ProtoReflect.Descriptor instead.
func (*GetKratosDemoUserReply) Descriptor() ([]byte, []int) {
	return file_api_kratosdemouser_kratosDemoUser_proto_rawDescGZIP(), []int{10}
}

func (x *GetKratosDemoUserReply) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

var File_api_kratosdemouser_kratosDemoUser_proto protoreflect.FileDescriptor

var file_api_kratosdemouser_kratosDemoUser_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x64, 0x65, 0x6d, 0x6f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x44, 0x65, 0x6d, 0x6f, 0x55,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x61, 0x70, 0x69, 0x2e, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x64, 0x65, 0x6d, 0x6f, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x07, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x1b, 0x0a, 0x07, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x2e, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x44, 0x65, 0x6d, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x56, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x44, 0x65, 0x6d, 0x6f, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x39, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x64, 0x65, 0x6d, 0x6f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x58, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x44, 0x65, 0x6d, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x39, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x64,
	0x65, 0x6d, 0x6f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x1b, 0x0a, 0x19, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x44, 0x65, 0x6d, 0x6f, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2d, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x44, 0x65, 0x6d, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1b, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x44, 0x65, 0x6d, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x2a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x44, 0x65, 0x6d, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x53, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x44, 0x65, 0x6d, 0x6f,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x39, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x64, 0x65, 0x6d, 0x6f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x32, 0xbe, 0x04, 0x0a, 0x0e, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x44,
	0x65, 0x6d, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x12, 0x76, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x44, 0x65, 0x6d, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x64, 0x65, 0x6d, 0x6f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x44, 0x65, 0x6d, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x64, 0x65, 0x6d,
	0x6f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x72, 0x61, 0x74,
	0x6f, 0x73, 0x44, 0x65, 0x6d, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x76, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x44,
	0x65, 0x6d, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x64, 0x65, 0x6d, 0x6f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x44, 0x65, 0x6d, 0x6f, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x64, 0x65, 0x6d, 0x6f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x44, 0x65, 0x6d, 0x6f, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x76, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x44, 0x65, 0x6d, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x64, 0x65, 0x6d, 0x6f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x44, 0x65, 0x6d, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x64, 0x65, 0x6d,
	0x6f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x72, 0x61, 0x74,
	0x6f, 0x73, 0x44, 0x65, 0x6d, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x6d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x44, 0x65, 0x6d, 0x6f,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x64, 0x65, 0x6d, 0x6f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x44, 0x65, 0x6d, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x64,
	0x65, 0x6d, 0x6f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x44, 0x65, 0x6d, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x55,
	0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x64, 0x65, 0x6d, 0x6f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x64, 0x65, 0x6d, 0x6f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x73, 0x70,
	0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2f,
	0x7b, 0x6d, 0x73, 0x67, 0x7d, 0x42, 0x5e, 0x0a, 0x12, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x64, 0x65, 0x6d, 0x6f, 0x75, 0x73, 0x65, 0x72, 0x50, 0x01, 0x5a, 0x46, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x31, 0x30, 0x30, 0x35, 0x32, 0x38,
	0x31, 0x33, 0x34, 0x32, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x64, 0x65, 0x6d, 0x6f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x64, 0x65,
	0x6d, 0x6f, 0x75, 0x73, 0x65, 0x72, 0x3b, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x64, 0x65, 0x6d,
	0x6f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_kratosdemouser_kratosDemoUser_proto_rawDescOnce sync.Once
	file_api_kratosdemouser_kratosDemoUser_proto_rawDescData = file_api_kratosdemouser_kratosDemoUser_proto_rawDesc
)

func file_api_kratosdemouser_kratosDemoUser_proto_rawDescGZIP() []byte {
	file_api_kratosdemouser_kratosDemoUser_proto_rawDescOnce.Do(func() {
		file_api_kratosdemouser_kratosDemoUser_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_kratosdemouser_kratosDemoUser_proto_rawDescData)
	})
	return file_api_kratosdemouser_kratosDemoUser_proto_rawDescData
}

var file_api_kratosdemouser_kratosDemoUser_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_kratosdemouser_kratosDemoUser_proto_goTypes = []interface{}{
	(*PingReq)(nil),                     // 0: api.kratosdemouser.PingReq
	(*PingRsp)(nil),                     // 1: api.kratosdemouser.PingRsp
	(*UserInfo)(nil),                    // 2: api.kratosdemouser.UserInfo
	(*CreateKratosDemoUserRequest)(nil), // 3: api.kratosdemouser.CreateKratosDemoUserRequest
	(*CreateKratosDemoUserReply)(nil),   // 4: api.kratosdemouser.CreateKratosDemoUserReply
	(*UpdateKratosDemoUserRequest)(nil), // 5: api.kratosdemouser.UpdateKratosDemoUserRequest
	(*UpdateKratosDemoUserReply)(nil),   // 6: api.kratosdemouser.UpdateKratosDemoUserReply
	(*DeleteKratosDemoUserRequest)(nil), // 7: api.kratosdemouser.DeleteKratosDemoUserRequest
	(*DeleteKratosDemoUserReply)(nil),   // 8: api.kratosdemouser.DeleteKratosDemoUserReply
	(*GetKratosDemoUserRequest)(nil),    // 9: api.kratosdemouser.GetKratosDemoUserRequest
	(*GetKratosDemoUserReply)(nil),      // 10: api.kratosdemouser.GetKratosDemoUserReply
}
var file_api_kratosdemouser_kratosDemoUser_proto_depIdxs = []int32{
	2,  // 0: api.kratosdemouser.CreateKratosDemoUserReply.user_info:type_name -> api.kratosdemouser.UserInfo
	2,  // 1: api.kratosdemouser.UpdateKratosDemoUserRequest.user_info:type_name -> api.kratosdemouser.UserInfo
	2,  // 2: api.kratosdemouser.GetKratosDemoUserReply.user_info:type_name -> api.kratosdemouser.UserInfo
	3,  // 3: api.kratosdemouser.KratosDemoUser.CreateKratosDemoUser:input_type -> api.kratosdemouser.CreateKratosDemoUserRequest
	5,  // 4: api.kratosdemouser.KratosDemoUser.UpdateKratosDemoUser:input_type -> api.kratosdemouser.UpdateKratosDemoUserRequest
	7,  // 5: api.kratosdemouser.KratosDemoUser.DeleteKratosDemoUser:input_type -> api.kratosdemouser.DeleteKratosDemoUserRequest
	9,  // 6: api.kratosdemouser.KratosDemoUser.GetKratosDemoUser:input_type -> api.kratosdemouser.GetKratosDemoUserRequest
	0,  // 7: api.kratosdemouser.KratosDemoUser.Ping:input_type -> api.kratosdemouser.PingReq
	4,  // 8: api.kratosdemouser.KratosDemoUser.CreateKratosDemoUser:output_type -> api.kratosdemouser.CreateKratosDemoUserReply
	6,  // 9: api.kratosdemouser.KratosDemoUser.UpdateKratosDemoUser:output_type -> api.kratosdemouser.UpdateKratosDemoUserReply
	8,  // 10: api.kratosdemouser.KratosDemoUser.DeleteKratosDemoUser:output_type -> api.kratosdemouser.DeleteKratosDemoUserReply
	10, // 11: api.kratosdemouser.KratosDemoUser.GetKratosDemoUser:output_type -> api.kratosdemouser.GetKratosDemoUserReply
	1,  // 12: api.kratosdemouser.KratosDemoUser.Ping:output_type -> api.kratosdemouser.PingRsp
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_api_kratosdemouser_kratosDemoUser_proto_init() }
func file_api_kratosdemouser_kratosDemoUser_proto_init() {
	if File_api_kratosdemouser_kratosDemoUser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingReq); i {
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
		file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRsp); i {
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
		file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateKratosDemoUserRequest); i {
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
		file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateKratosDemoUserReply); i {
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
		file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateKratosDemoUserRequest); i {
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
		file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateKratosDemoUserReply); i {
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
		file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteKratosDemoUserRequest); i {
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
		file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteKratosDemoUserReply); i {
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
		file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKratosDemoUserRequest); i {
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
		file_api_kratosdemouser_kratosDemoUser_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKratosDemoUserReply); i {
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
			RawDescriptor: file_api_kratosdemouser_kratosDemoUser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_kratosdemouser_kratosDemoUser_proto_goTypes,
		DependencyIndexes: file_api_kratosdemouser_kratosDemoUser_proto_depIdxs,
		MessageInfos:      file_api_kratosdemouser_kratosDemoUser_proto_msgTypes,
	}.Build()
	File_api_kratosdemouser_kratosDemoUser_proto = out.File
	file_api_kratosdemouser_kratosDemoUser_proto_rawDesc = nil
	file_api_kratosdemouser_kratosDemoUser_proto_goTypes = nil
	file_api_kratosdemouser_kratosDemoUser_proto_depIdxs = nil
}
