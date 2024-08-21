// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: api/thanos/demo.proto

package thanos

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

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thanos_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_thanos_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_api_thanos_demo_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thanos_demo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_thanos_demo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_api_thanos_demo_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateDemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Comment string `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CreateDemoRequest) Reset() {
	*x = CreateDemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thanos_demo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDemoRequest) ProtoMessage() {}

func (x *CreateDemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_thanos_demo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDemoRequest.ProtoReflect.Descriptor instead.
func (*CreateDemoRequest) Descriptor() ([]byte, []int) {
	return file_api_thanos_demo_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDemoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateDemoRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateDemoRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type CreateDemoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *CreateDemoReply) Reset() {
	*x = CreateDemoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thanos_demo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDemoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDemoReply) ProtoMessage() {}

func (x *CreateDemoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_thanos_demo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDemoReply.ProtoReflect.Descriptor instead.
func (*CreateDemoReply) Descriptor() ([]byte, []int) {
	return file_api_thanos_demo_proto_rawDescGZIP(), []int{3}
}

func (x *CreateDemoReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type UpdateDemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Comment string `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *UpdateDemoRequest) Reset() {
	*x = UpdateDemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thanos_demo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDemoRequest) ProtoMessage() {}

func (x *UpdateDemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_thanos_demo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDemoRequest.ProtoReflect.Descriptor instead.
func (*UpdateDemoRequest) Descriptor() ([]byte, []int) {
	return file_api_thanos_demo_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateDemoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateDemoRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpdateDemoRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type UpdateDemoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *UpdateDemoReply) Reset() {
	*x = UpdateDemoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thanos_demo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDemoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDemoReply) ProtoMessage() {}

func (x *UpdateDemoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_thanos_demo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDemoReply.ProtoReflect.Descriptor instead.
func (*UpdateDemoReply) Descriptor() ([]byte, []int) {
	return file_api_thanos_demo_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateDemoReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type DeleteDemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteDemoRequest) Reset() {
	*x = DeleteDemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thanos_demo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDemoRequest) ProtoMessage() {}

func (x *DeleteDemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_thanos_demo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDemoRequest.ProtoReflect.Descriptor instead.
func (*DeleteDemoRequest) Descriptor() ([]byte, []int) {
	return file_api_thanos_demo_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteDemoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteDemoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *DeleteDemoReply) Reset() {
	*x = DeleteDemoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thanos_demo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDemoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDemoReply) ProtoMessage() {}

func (x *DeleteDemoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_thanos_demo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDemoReply.ProtoReflect.Descriptor instead.
func (*DeleteDemoReply) Descriptor() ([]byte, []int) {
	return file_api_thanos_demo_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteDemoReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type GetDemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDemoRequest) Reset() {
	*x = GetDemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thanos_demo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDemoRequest) ProtoMessage() {}

func (x *GetDemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_thanos_demo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDemoRequest.ProtoReflect.Descriptor instead.
func (*GetDemoRequest) Descriptor() ([]byte, []int) {
	return file_api_thanos_demo_proto_rawDescGZIP(), []int{8}
}

type GetDemoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *GetDemoReply) Reset() {
	*x = GetDemoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thanos_demo_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDemoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDemoReply) ProtoMessage() {}

func (x *GetDemoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_thanos_demo_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDemoReply.ProtoReflect.Descriptor instead.
func (*GetDemoReply) Descriptor() ([]byte, []int) {
	return file_api_thanos_demo_proto_rawDescGZIP(), []int{9}
}

func (x *GetDemoReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type ListDemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ListDemoRequest) Reset() {
	*x = ListDemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thanos_demo_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDemoRequest) ProtoMessage() {}

func (x *ListDemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_thanos_demo_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDemoRequest.ProtoReflect.Descriptor instead.
func (*ListDemoRequest) Descriptor() ([]byte, []int) {
	return file_api_thanos_demo_proto_rawDescGZIP(), []int{10}
}

func (x *ListDemoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListDemoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *ListDemoReply) Reset() {
	*x = ListDemoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thanos_demo_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDemoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDemoReply) ProtoMessage() {}

func (x *ListDemoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_thanos_demo_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDemoReply.ProtoReflect.Descriptor instead.
func (*ListDemoReply) Descriptor() ([]byte, []int) {
	return file_api_thanos_demo_proto_rawDescGZIP(), []int{11}
}

func (x *ListDemoReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_api_thanos_demo_proto protoreflect.FileDescriptor

var file_api_thanos_demo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2f, 0x64, 0x65, 0x6d,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x68, 0x61,
	0x6e, 0x6f, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x57, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x21, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x57, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x21, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44,
	0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x0f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x10, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x1e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22,
	0x21, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x1f, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x02, 0x6f, 0x6b, 0x32, 0xb8, 0x04, 0x0a, 0x04, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x64, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65,
	0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6d,
	0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01,
	0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x64, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f,
	0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x61, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x68, 0x61,
	0x6e, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x68, 0x61, 0x6e,
	0x6f, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x50, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x68, 0x61,
	0x6e, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x59, 0x0a,
	0x08, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x68, 0x61,
	0x6e, 0x6f, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x54, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f,
	0x73, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e,
	0x2f, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x42, 0x27,
	0x0a, 0x0a, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x17,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73,
	0x3b, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_thanos_demo_proto_rawDescOnce sync.Once
	file_api_thanos_demo_proto_rawDescData = file_api_thanos_demo_proto_rawDesc
)

func file_api_thanos_demo_proto_rawDescGZIP() []byte {
	file_api_thanos_demo_proto_rawDescOnce.Do(func() {
		file_api_thanos_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_thanos_demo_proto_rawDescData)
	})
	return file_api_thanos_demo_proto_rawDescData
}

var file_api_thanos_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_api_thanos_demo_proto_goTypes = []any{
	(*HelloRequest)(nil),      // 0: api.thanos.HelloRequest
	(*HelloReply)(nil),        // 1: api.thanos.HelloReply
	(*CreateDemoRequest)(nil), // 2: api.thanos.CreateDemoRequest
	(*CreateDemoReply)(nil),   // 3: api.thanos.CreateDemoReply
	(*UpdateDemoRequest)(nil), // 4: api.thanos.UpdateDemoRequest
	(*UpdateDemoReply)(nil),   // 5: api.thanos.UpdateDemoReply
	(*DeleteDemoRequest)(nil), // 6: api.thanos.DeleteDemoRequest
	(*DeleteDemoReply)(nil),   // 7: api.thanos.DeleteDemoReply
	(*GetDemoRequest)(nil),    // 8: api.thanos.GetDemoRequest
	(*GetDemoReply)(nil),      // 9: api.thanos.GetDemoReply
	(*ListDemoRequest)(nil),   // 10: api.thanos.ListDemoRequest
	(*ListDemoReply)(nil),     // 11: api.thanos.ListDemoReply
}
var file_api_thanos_demo_proto_depIdxs = []int32{
	2,  // 0: api.thanos.Demo.CreateDemo:input_type -> api.thanos.CreateDemoRequest
	4,  // 1: api.thanos.Demo.UpdateDemo:input_type -> api.thanos.UpdateDemoRequest
	6,  // 2: api.thanos.Demo.DeleteDemo:input_type -> api.thanos.DeleteDemoRequest
	8,  // 3: api.thanos.Demo.GetDemo:input_type -> api.thanos.GetDemoRequest
	10, // 4: api.thanos.Demo.ListDemo:input_type -> api.thanos.ListDemoRequest
	0,  // 5: api.thanos.Demo.SayHello:input_type -> api.thanos.HelloRequest
	3,  // 6: api.thanos.Demo.CreateDemo:output_type -> api.thanos.CreateDemoReply
	5,  // 7: api.thanos.Demo.UpdateDemo:output_type -> api.thanos.UpdateDemoReply
	7,  // 8: api.thanos.Demo.DeleteDemo:output_type -> api.thanos.DeleteDemoReply
	9,  // 9: api.thanos.Demo.GetDemo:output_type -> api.thanos.GetDemoReply
	11, // 10: api.thanos.Demo.ListDemo:output_type -> api.thanos.ListDemoReply
	1,  // 11: api.thanos.Demo.SayHello:output_type -> api.thanos.HelloReply
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_thanos_demo_proto_init() }
func file_api_thanos_demo_proto_init() {
	if File_api_thanos_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_thanos_demo_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*HelloRequest); i {
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
		file_api_thanos_demo_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*HelloReply); i {
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
		file_api_thanos_demo_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateDemoRequest); i {
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
		file_api_thanos_demo_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateDemoReply); i {
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
		file_api_thanos_demo_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateDemoRequest); i {
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
		file_api_thanos_demo_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateDemoReply); i {
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
		file_api_thanos_demo_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteDemoRequest); i {
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
		file_api_thanos_demo_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteDemoReply); i {
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
		file_api_thanos_demo_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetDemoRequest); i {
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
		file_api_thanos_demo_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GetDemoReply); i {
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
		file_api_thanos_demo_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*ListDemoRequest); i {
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
		file_api_thanos_demo_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*ListDemoReply); i {
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
			RawDescriptor: file_api_thanos_demo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_thanos_demo_proto_goTypes,
		DependencyIndexes: file_api_thanos_demo_proto_depIdxs,
		MessageInfos:      file_api_thanos_demo_proto_msgTypes,
	}.Build()
	File_api_thanos_demo_proto = out.File
	file_api_thanos_demo_proto_rawDesc = nil
	file_api_thanos_demo_proto_goTypes = nil
	file_api_thanos_demo_proto_depIdxs = nil
}
