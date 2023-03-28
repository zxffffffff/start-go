// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: MsgType.proto

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

type PBMsgType int32

const (
	PBMsgType_InvalidMsgType PBMsgType = 0
	// socket连接心跳
	PBMsgType_HeartbeatMsgType PBMsgType = 1
	// 行情推送订阅
	PBMsgType_QuoteSubscribeTickMsgType PBMsgType = 2
	// 行情推送取消订阅
	PBMsgType_QuoteUnsubscribeTickMsgType PBMsgType = 3
	// 行情推送
	PBMsgType_QuotePushTickMsgType PBMsgType = 4
)

// Enum value maps for PBMsgType.
var (
	PBMsgType_name = map[int32]string{
		0: "InvalidMsgType",
		1: "HeartbeatMsgType",
		2: "QuoteSubscribeTickMsgType",
		3: "QuoteUnsubscribeTickMsgType",
		4: "QuotePushTickMsgType",
	}
	PBMsgType_value = map[string]int32{
		"InvalidMsgType":              0,
		"HeartbeatMsgType":            1,
		"QuoteSubscribeTickMsgType":   2,
		"QuoteUnsubscribeTickMsgType": 3,
		"QuotePushTickMsgType":        4,
	}
)

func (x PBMsgType) Enum() *PBMsgType {
	p := new(PBMsgType)
	*p = x
	return p
}

func (x PBMsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PBMsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_MsgType_proto_enumTypes[0].Descriptor()
}

func (PBMsgType) Type() protoreflect.EnumType {
	return &file_MsgType_proto_enumTypes[0]
}

func (x PBMsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PBMsgType.Descriptor instead.
func (PBMsgType) EnumDescriptor() ([]byte, []int) {
	return file_MsgType_proto_rawDescGZIP(), []int{0}
}

var File_MsgType_proto protoreflect.FileDescriptor

var file_MsgType_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a,
	0x8f, 0x01, 0x0a, 0x09, 0x50, 0x42, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x0e, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x4d, 0x73,
	0x67, 0x54, 0x79, 0x70, 0x65, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x4d, 0x73, 0x67,
	0x54, 0x79, 0x70, 0x65, 0x10, 0x02, 0x12, 0x1f, 0x0a, 0x1b, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x55,
	0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x4d, 0x73,
	0x67, 0x54, 0x79, 0x70, 0x65, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x50, 0x75, 0x73, 0x68, 0x54, 0x69, 0x63, 0x6b, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x10,
	0x04, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_MsgType_proto_rawDescOnce sync.Once
	file_MsgType_proto_rawDescData = file_MsgType_proto_rawDesc
)

func file_MsgType_proto_rawDescGZIP() []byte {
	file_MsgType_proto_rawDescOnce.Do(func() {
		file_MsgType_proto_rawDescData = protoimpl.X.CompressGZIP(file_MsgType_proto_rawDescData)
	})
	return file_MsgType_proto_rawDescData
}

var file_MsgType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_MsgType_proto_goTypes = []interface{}{
	(PBMsgType)(0), // 0: PBMsgType
}
var file_MsgType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MsgType_proto_init() }
func file_MsgType_proto_init() {
	if File_MsgType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_MsgType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MsgType_proto_goTypes,
		DependencyIndexes: file_MsgType_proto_depIdxs,
		EnumInfos:         file_MsgType_proto_enumTypes,
	}.Build()
	File_MsgType_proto = out.File
	file_MsgType_proto_rawDesc = nil
	file_MsgType_proto_goTypes = nil
	file_MsgType_proto_depIdxs = nil
}
