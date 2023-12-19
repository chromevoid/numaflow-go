// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: pkg/apis/proto/reduce/v1/reduce.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReduceRequest_WindowOperation_Event int32

const (
	ReduceRequest_WindowOperation_OPEN   ReduceRequest_WindowOperation_Event = 0
	ReduceRequest_WindowOperation_CLOSE  ReduceRequest_WindowOperation_Event = 1
	ReduceRequest_WindowOperation_APPEND ReduceRequest_WindowOperation_Event = 4
)

// Enum value maps for ReduceRequest_WindowOperation_Event.
var (
	ReduceRequest_WindowOperation_Event_name = map[int32]string{
		0: "OPEN",
		1: "CLOSE",
		4: "APPEND",
	}
	ReduceRequest_WindowOperation_Event_value = map[string]int32{
		"OPEN":   0,
		"CLOSE":  1,
		"APPEND": 4,
	}
)

func (x ReduceRequest_WindowOperation_Event) Enum() *ReduceRequest_WindowOperation_Event {
	p := new(ReduceRequest_WindowOperation_Event)
	*p = x
	return p
}

func (x ReduceRequest_WindowOperation_Event) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReduceRequest_WindowOperation_Event) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_apis_proto_reduce_v1_reduce_proto_enumTypes[0].Descriptor()
}

func (ReduceRequest_WindowOperation_Event) Type() protoreflect.EnumType {
	return &file_pkg_apis_proto_reduce_v1_reduce_proto_enumTypes[0]
}

func (x ReduceRequest_WindowOperation_Event) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReduceRequest_WindowOperation_Event.Descriptor instead.
func (ReduceRequest_WindowOperation_Event) EnumDescriptor() ([]byte, []int) {
	return file_pkg_apis_proto_reduce_v1_reduce_proto_rawDescGZIP(), []int{0, 0, 0}
}

// *
// ReduceRequest represents a request element.
type ReduceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload   *ReduceRequest_Payload         `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Operation *ReduceRequest_WindowOperation `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *ReduceRequest) Reset() {
	*x = ReduceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_reduce_v1_reduce_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReduceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReduceRequest) ProtoMessage() {}

func (x *ReduceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_reduce_v1_reduce_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReduceRequest.ProtoReflect.Descriptor instead.
func (*ReduceRequest) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_reduce_v1_reduce_proto_rawDescGZIP(), []int{0}
}

func (x *ReduceRequest) GetPayload() *ReduceRequest_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *ReduceRequest) GetOperation() *ReduceRequest_WindowOperation {
	if x != nil {
		return x.Operation
	}
	return nil
}

// Window represents a window.
// Since the client doesn't track keys, window doesn't have a keys field.
type Window struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	Slot  string                 `protobuf:"bytes,3,opt,name=slot,proto3" json:"slot,omitempty"`
}

func (x *Window) Reset() {
	*x = Window{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_reduce_v1_reduce_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Window) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Window) ProtoMessage() {}

func (x *Window) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_reduce_v1_reduce_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Window.ProtoReflect.Descriptor instead.
func (*Window) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_reduce_v1_reduce_proto_rawDescGZIP(), []int{1}
}

func (x *Window) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *Window) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *Window) GetSlot() string {
	if x != nil {
		return x.Slot
	}
	return ""
}

// *
// ReduceResponse represents a response element.
type ReduceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *ReduceResponse_Result `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	// window represents a window to which the result belongs.
	Window *Window `protobuf:"bytes,2,opt,name=window,proto3" json:"window,omitempty"`
	// EOF represents the end of the response for a window.
	EOF bool `protobuf:"varint,3,opt,name=EOF,proto3" json:"EOF,omitempty"`
}

func (x *ReduceResponse) Reset() {
	*x = ReduceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_reduce_v1_reduce_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReduceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReduceResponse) ProtoMessage() {}

func (x *ReduceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_reduce_v1_reduce_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReduceResponse.ProtoReflect.Descriptor instead.
func (*ReduceResponse) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_reduce_v1_reduce_proto_rawDescGZIP(), []int{2}
}

func (x *ReduceResponse) GetResult() *ReduceResponse_Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *ReduceResponse) GetWindow() *Window {
	if x != nil {
		return x.Window
	}
	return nil
}

func (x *ReduceResponse) GetEOF() bool {
	if x != nil {
		return x.EOF
	}
	return false
}

// *
// ReadyResponse is the health check result.
type ReadyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ready bool `protobuf:"varint,1,opt,name=ready,proto3" json:"ready,omitempty"`
}

func (x *ReadyResponse) Reset() {
	*x = ReadyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_reduce_v1_reduce_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadyResponse) ProtoMessage() {}

func (x *ReadyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_reduce_v1_reduce_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadyResponse.ProtoReflect.Descriptor instead.
func (*ReadyResponse) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_reduce_v1_reduce_proto_rawDescGZIP(), []int{3}
}

func (x *ReadyResponse) GetReady() bool {
	if x != nil {
		return x.Ready
	}
	return false
}

// WindowOperation represents a window operation.
// For Aligned windows, OPEN, APPEND and CLOSE events are sent.
type ReduceRequest_WindowOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event   ReduceRequest_WindowOperation_Event `protobuf:"varint,1,opt,name=event,proto3,enum=reduce.v1.ReduceRequest_WindowOperation_Event" json:"event,omitempty"`
	Windows []*Window                           `protobuf:"bytes,2,rep,name=windows,proto3" json:"windows,omitempty"`
}

func (x *ReduceRequest_WindowOperation) Reset() {
	*x = ReduceRequest_WindowOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_reduce_v1_reduce_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReduceRequest_WindowOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReduceRequest_WindowOperation) ProtoMessage() {}

func (x *ReduceRequest_WindowOperation) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_reduce_v1_reduce_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReduceRequest_WindowOperation.ProtoReflect.Descriptor instead.
func (*ReduceRequest_WindowOperation) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_reduce_v1_reduce_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ReduceRequest_WindowOperation) GetEvent() ReduceRequest_WindowOperation_Event {
	if x != nil {
		return x.Event
	}
	return ReduceRequest_WindowOperation_OPEN
}

func (x *ReduceRequest_WindowOperation) GetWindows() []*Window {
	if x != nil {
		return x.Windows
	}
	return nil
}

// Payload represents a payload element.
type ReduceRequest_Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys      []string               `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	Value     []byte                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	EventTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	Watermark *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=watermark,proto3" json:"watermark,omitempty"`
}

func (x *ReduceRequest_Payload) Reset() {
	*x = ReduceRequest_Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_reduce_v1_reduce_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReduceRequest_Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReduceRequest_Payload) ProtoMessage() {}

func (x *ReduceRequest_Payload) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_reduce_v1_reduce_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReduceRequest_Payload.ProtoReflect.Descriptor instead.
func (*ReduceRequest_Payload) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_reduce_v1_reduce_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ReduceRequest_Payload) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *ReduceRequest_Payload) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *ReduceRequest_Payload) GetEventTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EventTime
	}
	return nil
}

func (x *ReduceRequest_Payload) GetWatermark() *timestamppb.Timestamp {
	if x != nil {
		return x.Watermark
	}
	return nil
}

// Result represents a result element. It contains the result of the reduce function.
type ReduceResponse_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys  []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	Value []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Tags  []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *ReduceResponse_Result) Reset() {
	*x = ReduceResponse_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_reduce_v1_reduce_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReduceResponse_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReduceResponse_Result) ProtoMessage() {}

func (x *ReduceResponse_Result) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_reduce_v1_reduce_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReduceResponse_Result.ProtoReflect.Descriptor instead.
func (*ReduceResponse_Result) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_reduce_v1_reduce_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ReduceResponse_Result) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *ReduceResponse_Result) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *ReduceResponse_Result) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_pkg_apis_proto_reduce_v1_reduce_proto protoreflect.FileDescriptor

var file_pkg_apis_proto_reduce_v1_reduce_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x64, 0x75, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xef, 0x03, 0x0a, 0x0d, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3a, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x46,
	0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x64, 0x75, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x57, 0x69, 0x6e, 0x64,
	0x6f, 0x77, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0xae, 0x01, 0x0a, 0x0f, 0x57, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x05, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x72, 0x65, 0x64, 0x75,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x2b, 0x0a, 0x07, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69,
	0x6e, 0x64, 0x6f, 0x77, 0x52, 0x07, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x22, 0x28, 0x0a,
	0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41,
	0x50, 0x50, 0x45, 0x4e, 0x44, 0x10, 0x04, 0x1a, 0xa8, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x39, 0x0a,
	0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x77, 0x61, 0x74, 0x65,
	0x72, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x77, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61,
	0x72, 0x6b, 0x22, 0x7c, 0x0a, 0x06, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x30, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2c,
	0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x6c, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74,
	0x22, 0xcf, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x29, 0x0a,
	0x06, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77,
	0x52, 0x06, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x4f, 0x46, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x45, 0x4f, 0x46, 0x1a, 0x46, 0x0a, 0x06, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x22, 0x25, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x32, 0x8a, 0x01, 0x0a, 0x06, 0x52, 0x65,
	0x64, 0x75, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x46, 0x6e,
	0x12, 0x18, 0x2e, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x64,
	0x75, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x65, 0x64,
	0x75, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x3b, 0x0a, 0x07, 0x49, 0x73, 0x52,
	0x65, 0x61, 0x64, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x72,
	0x65, 0x64, 0x75, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x75, 0x6d, 0x61, 0x70, 0x72, 0x6f, 0x6a, 0x2f, 0x6e, 0x75,
	0x6d, 0x61, 0x66, 0x6c, 0x6f, 0x77, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_apis_proto_reduce_v1_reduce_proto_rawDescOnce sync.Once
	file_pkg_apis_proto_reduce_v1_reduce_proto_rawDescData = file_pkg_apis_proto_reduce_v1_reduce_proto_rawDesc
)

func file_pkg_apis_proto_reduce_v1_reduce_proto_rawDescGZIP() []byte {
	file_pkg_apis_proto_reduce_v1_reduce_proto_rawDescOnce.Do(func() {
		file_pkg_apis_proto_reduce_v1_reduce_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_apis_proto_reduce_v1_reduce_proto_rawDescData)
	})
	return file_pkg_apis_proto_reduce_v1_reduce_proto_rawDescData
}

var file_pkg_apis_proto_reduce_v1_reduce_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_apis_proto_reduce_v1_reduce_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_apis_proto_reduce_v1_reduce_proto_goTypes = []interface{}{
	(ReduceRequest_WindowOperation_Event)(0), // 0: reduce.v1.ReduceRequest.WindowOperation.Event
	(*ReduceRequest)(nil),                    // 1: reduce.v1.ReduceRequest
	(*Window)(nil),                           // 2: reduce.v1.Window
	(*ReduceResponse)(nil),                   // 3: reduce.v1.ReduceResponse
	(*ReadyResponse)(nil),                    // 4: reduce.v1.ReadyResponse
	(*ReduceRequest_WindowOperation)(nil),    // 5: reduce.v1.ReduceRequest.WindowOperation
	(*ReduceRequest_Payload)(nil),            // 6: reduce.v1.ReduceRequest.Payload
	(*ReduceResponse_Result)(nil),            // 7: reduce.v1.ReduceResponse.Result
	(*timestamppb.Timestamp)(nil),            // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                    // 9: google.protobuf.Empty
}
var file_pkg_apis_proto_reduce_v1_reduce_proto_depIdxs = []int32{
	6,  // 0: reduce.v1.ReduceRequest.payload:type_name -> reduce.v1.ReduceRequest.Payload
	5,  // 1: reduce.v1.ReduceRequest.operation:type_name -> reduce.v1.ReduceRequest.WindowOperation
	8,  // 2: reduce.v1.Window.start:type_name -> google.protobuf.Timestamp
	8,  // 3: reduce.v1.Window.end:type_name -> google.protobuf.Timestamp
	7,  // 4: reduce.v1.ReduceResponse.result:type_name -> reduce.v1.ReduceResponse.Result
	2,  // 5: reduce.v1.ReduceResponse.window:type_name -> reduce.v1.Window
	0,  // 6: reduce.v1.ReduceRequest.WindowOperation.event:type_name -> reduce.v1.ReduceRequest.WindowOperation.Event
	2,  // 7: reduce.v1.ReduceRequest.WindowOperation.windows:type_name -> reduce.v1.Window
	8,  // 8: reduce.v1.ReduceRequest.Payload.event_time:type_name -> google.protobuf.Timestamp
	8,  // 9: reduce.v1.ReduceRequest.Payload.watermark:type_name -> google.protobuf.Timestamp
	1,  // 10: reduce.v1.Reduce.ReduceFn:input_type -> reduce.v1.ReduceRequest
	9,  // 11: reduce.v1.Reduce.IsReady:input_type -> google.protobuf.Empty
	3,  // 12: reduce.v1.Reduce.ReduceFn:output_type -> reduce.v1.ReduceResponse
	4,  // 13: reduce.v1.Reduce.IsReady:output_type -> reduce.v1.ReadyResponse
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_pkg_apis_proto_reduce_v1_reduce_proto_init() }
func file_pkg_apis_proto_reduce_v1_reduce_proto_init() {
	if File_pkg_apis_proto_reduce_v1_reduce_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_apis_proto_reduce_v1_reduce_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReduceRequest); i {
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
		file_pkg_apis_proto_reduce_v1_reduce_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Window); i {
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
		file_pkg_apis_proto_reduce_v1_reduce_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReduceResponse); i {
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
		file_pkg_apis_proto_reduce_v1_reduce_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadyResponse); i {
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
		file_pkg_apis_proto_reduce_v1_reduce_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReduceRequest_WindowOperation); i {
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
		file_pkg_apis_proto_reduce_v1_reduce_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReduceRequest_Payload); i {
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
		file_pkg_apis_proto_reduce_v1_reduce_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReduceResponse_Result); i {
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
			RawDescriptor: file_pkg_apis_proto_reduce_v1_reduce_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_apis_proto_reduce_v1_reduce_proto_goTypes,
		DependencyIndexes: file_pkg_apis_proto_reduce_v1_reduce_proto_depIdxs,
		EnumInfos:         file_pkg_apis_proto_reduce_v1_reduce_proto_enumTypes,
		MessageInfos:      file_pkg_apis_proto_reduce_v1_reduce_proto_msgTypes,
	}.Build()
	File_pkg_apis_proto_reduce_v1_reduce_proto = out.File
	file_pkg_apis_proto_reduce_v1_reduce_proto_rawDesc = nil
	file_pkg_apis_proto_reduce_v1_reduce_proto_goTypes = nil
	file_pkg_apis_proto_reduce_v1_reduce_proto_depIdxs = nil
}
