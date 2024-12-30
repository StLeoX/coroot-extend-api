// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: coroot/event/v1/server_span.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// server-side L7 span
type ServerSpan struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// timestamp when reading the request
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// duration of the span
	Duration int64 `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
	// server-side container ID
	ContainerId string `protobuf:"bytes,3,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	// thread who reads the request
	TgidRead string `protobuf:"bytes,4,opt,name=tgid_read,json=tgidRead,proto3" json:"tgid_read,omitempty"`
	// thread who writes the response
	TgidWrite string `protobuf:"bytes,5,opt,name=tgid_write,json=tgidWrite,proto3" json:"tgid_write,omitempty"`
	// Some L7 protocols support.
	RequestId     string `protobuf:"bytes,6,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ServerSpan) Reset() {
	*x = ServerSpan{}
	mi := &file_coroot_event_v1_server_span_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServerSpan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerSpan) ProtoMessage() {}

func (x *ServerSpan) ProtoReflect() protoreflect.Message {
	mi := &file_coroot_event_v1_server_span_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerSpan.ProtoReflect.Descriptor instead.
func (*ServerSpan) Descriptor() ([]byte, []int) {
	return file_coroot_event_v1_server_span_proto_rawDescGZIP(), []int{0}
}

func (x *ServerSpan) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *ServerSpan) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *ServerSpan) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

func (x *ServerSpan) GetTgidRead() string {
	if x != nil {
		return x.TgidRead
	}
	return ""
}

func (x *ServerSpan) GetTgidWrite() string {
	if x != nil {
		return x.TgidWrite
	}
	return ""
}

func (x *ServerSpan) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

var File_coroot_event_v1_server_span_proto protoreflect.FileDescriptor

var file_coroot_event_v1_server_span_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x6f, 0x72, 0x6f, 0x6f, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x61, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6f, 0x72, 0x6f, 0x6f, 0x74, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x01, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x53, 0x70, 0x61, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x67, 0x69, 0x64, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x67, 0x69, 0x64, 0x52, 0x65, 0x61, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x67,
	0x69, 0x64, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x67, 0x69, 0x64, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x74, 0x4c, 0x65, 0x6f, 0x58, 0x2f, 0x63, 0x6f,
	0x72, 0x6f, 0x6f, 0x74, 0x2d, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x6f, 0x6f, 0x74,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_coroot_event_v1_server_span_proto_rawDescOnce sync.Once
	file_coroot_event_v1_server_span_proto_rawDescData = file_coroot_event_v1_server_span_proto_rawDesc
)

func file_coroot_event_v1_server_span_proto_rawDescGZIP() []byte {
	file_coroot_event_v1_server_span_proto_rawDescOnce.Do(func() {
		file_coroot_event_v1_server_span_proto_rawDescData = protoimpl.X.CompressGZIP(file_coroot_event_v1_server_span_proto_rawDescData)
	})
	return file_coroot_event_v1_server_span_proto_rawDescData
}

var file_coroot_event_v1_server_span_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_coroot_event_v1_server_span_proto_goTypes = []any{
	(*ServerSpan)(nil),            // 0: coroot.event.v1.ServerSpan
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_coroot_event_v1_server_span_proto_depIdxs = []int32{
	1, // 0: coroot.event.v1.ServerSpan.timestamp:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_coroot_event_v1_server_span_proto_init() }
func file_coroot_event_v1_server_span_proto_init() {
	if File_coroot_event_v1_server_span_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coroot_event_v1_server_span_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coroot_event_v1_server_span_proto_goTypes,
		DependencyIndexes: file_coroot_event_v1_server_span_proto_depIdxs,
		MessageInfos:      file_coroot_event_v1_server_span_proto_msgTypes,
	}.Build()
	File_coroot_event_v1_server_span_proto = out.File
	file_coroot_event_v1_server_span_proto_rawDesc = nil
	file_coroot_event_v1_server_span_proto_goTypes = nil
	file_coroot_event_v1_server_span_proto_depIdxs = nil
}
