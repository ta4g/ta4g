// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.11.2
// source: interval/bar/bar.proto

package bar

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type StandardBar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Time interval of this bar
	Time *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	// Open price for the current bar
	Open float64 `protobuf:"fixed64,2,opt,name=open,proto3" json:"open,omitempty"`
	// High price for the current bar
	High float64 `protobuf:"fixed64,3,opt,name=high,proto3" json:"high,omitempty"`
	// Low price for the current bar
	Low float64 `protobuf:"fixed64,4,opt,name=low,proto3" json:"low,omitempty"`
	// Close price for the current bar
	Close float64 `protobuf:"fixed64,5,opt,name=close,proto3" json:"close,omitempty"`
	// Volume of shares, options, coins, etc traded during this bar
	Volume float64 `protobuf:"fixed64,6,opt,name=volume,proto3" json:"volume,omitempty"`
	// OpenInterest (Optional) amount of derivatives currently outstanding for this bar
	// If there is no open interest then this will be -1, to indicate no data
	OpenInterest int64 `protobuf:"varint,7,opt,name=open_interest,json=openInterest,proto3" json:"open_interest,omitempty"`
}

func (x *StandardBar) Reset() {
	*x = StandardBar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interval_bar_bar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StandardBar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StandardBar) ProtoMessage() {}

func (x *StandardBar) ProtoReflect() protoreflect.Message {
	mi := &file_interval_bar_bar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StandardBar.ProtoReflect.Descriptor instead.
func (*StandardBar) Descriptor() ([]byte, []int) {
	return file_interval_bar_bar_proto_rawDescGZIP(), []int{0}
}

func (x *StandardBar) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *StandardBar) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *StandardBar) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *StandardBar) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *StandardBar) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *StandardBar) GetVolume() float64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *StandardBar) GetOpenInterest() int64 {
	if x != nil {
		return x.OpenInterest
	}
	return 0
}

type StandardBars struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bars []*StandardBar `protobuf:"bytes,1,rep,name=bars,proto3" json:"bars,omitempty"`
}

func (x *StandardBars) Reset() {
	*x = StandardBars{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interval_bar_bar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StandardBars) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StandardBars) ProtoMessage() {}

func (x *StandardBars) ProtoReflect() protoreflect.Message {
	mi := &file_interval_bar_bar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StandardBars.ProtoReflect.Descriptor instead.
func (*StandardBars) Descriptor() ([]byte, []int) {
	return file_interval_bar_bar_proto_rawDescGZIP(), []int{1}
}

func (x *StandardBars) GetBars() []*StandardBar {
	if x != nil {
		return x.Bars
	}
	return nil
}

var File_interval_bar_bar_proto protoreflect.FileDescriptor

var file_interval_bar_bar_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x2f, 0x62, 0x61, 0x72, 0x2f, 0x62,
	0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x62, 0x61, 0x72, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca,
	0x01, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x42, 0x61, 0x72, 0x12, 0x2e,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6f, 0x70,
	0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6f,
	0x70, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x0c, 0x53,
	0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x42, 0x61, 0x72, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x62,
	0x61, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x61, 0x72, 0x2e,
	0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x42, 0x61, 0x72, 0x52, 0x04, 0x62, 0x61, 0x72,
	0x73, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x61, 0x34, 0x67, 0x2f, 0x74, 0x61, 0x34, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x2f, 0x62, 0x61, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interval_bar_bar_proto_rawDescOnce sync.Once
	file_interval_bar_bar_proto_rawDescData = file_interval_bar_bar_proto_rawDesc
)

func file_interval_bar_bar_proto_rawDescGZIP() []byte {
	file_interval_bar_bar_proto_rawDescOnce.Do(func() {
		file_interval_bar_bar_proto_rawDescData = protoimpl.X.CompressGZIP(file_interval_bar_bar_proto_rawDescData)
	})
	return file_interval_bar_bar_proto_rawDescData
}

var file_interval_bar_bar_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_interval_bar_bar_proto_goTypes = []interface{}{
	(*StandardBar)(nil),         // 0: bar.StandardBar
	(*StandardBars)(nil),        // 1: bar.StandardBars
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_interval_bar_bar_proto_depIdxs = []int32{
	2, // 0: bar.StandardBar.time:type_name -> google.protobuf.Timestamp
	0, // 1: bar.StandardBars.bars:type_name -> bar.StandardBar
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_interval_bar_bar_proto_init() }
func file_interval_bar_bar_proto_init() {
	if File_interval_bar_bar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interval_bar_bar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StandardBar); i {
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
		file_interval_bar_bar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StandardBars); i {
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
			RawDescriptor: file_interval_bar_bar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_interval_bar_bar_proto_goTypes,
		DependencyIndexes: file_interval_bar_bar_proto_depIdxs,
		MessageInfos:      file_interval_bar_bar_proto_msgTypes,
	}.Build()
	File_interval_bar_bar_proto = out.File
	file_interval_bar_bar_proto_rawDesc = nil
	file_interval_bar_bar_proto_goTypes = nil
	file_interval_bar_bar_proto_depIdxs = nil
}
