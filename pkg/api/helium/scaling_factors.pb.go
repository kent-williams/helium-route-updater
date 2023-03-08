// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: scaling_factors.proto

package helium

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

// map of string h3 res 12 hex address to reward scaling factor
// as an integer representation of the 4-point precision decimal multiplier
// Ex: 5015 = 0.5015
type ScalingFactors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HexToSf map[string]uint32 `protobuf:"bytes,1,rep,name=hex_to_sf,json=hexToSf,proto3" json:"hex_to_sf,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *ScalingFactors) Reset() {
	*x = ScalingFactors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scaling_factors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScalingFactors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScalingFactors) ProtoMessage() {}

func (x *ScalingFactors) ProtoReflect() protoreflect.Message {
	mi := &file_scaling_factors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScalingFactors.ProtoReflect.Descriptor instead.
func (*ScalingFactors) Descriptor() ([]byte, []int) {
	return file_scaling_factors_proto_rawDescGZIP(), []int{0}
}

func (x *ScalingFactors) GetHexToSf() map[string]uint32 {
	if x != nil {
		return x.HexToSf
	}
	return nil
}

var File_scaling_factors_proto protoreflect.FileDescriptor

var file_scaling_factors_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x22,
	0x8f, 0x01, 0x0a, 0x0f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x73, 0x12, 0x40, 0x0a, 0x09, 0x68, 0x65, 0x78, 0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x66,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e,
	0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e,
	0x48, 0x65, 0x78, 0x54, 0x6f, 0x53, 0x66, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65,
	0x78, 0x54, 0x6f, 0x53, 0x66, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x78, 0x54, 0x6f, 0x53, 0x66,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scaling_factors_proto_rawDescOnce sync.Once
	file_scaling_factors_proto_rawDescData = file_scaling_factors_proto_rawDesc
)

func file_scaling_factors_proto_rawDescGZIP() []byte {
	file_scaling_factors_proto_rawDescOnce.Do(func() {
		file_scaling_factors_proto_rawDescData = protoimpl.X.CompressGZIP(file_scaling_factors_proto_rawDescData)
	})
	return file_scaling_factors_proto_rawDescData
}

var file_scaling_factors_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_scaling_factors_proto_goTypes = []interface{}{
	(*ScalingFactors)(nil), // 0: helium.scaling_factors
	nil,                    // 1: helium.scaling_factors.HexToSfEntry
}
var file_scaling_factors_proto_depIdxs = []int32{
	1, // 0: helium.scaling_factors.hex_to_sf:type_name -> helium.scaling_factors.HexToSfEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_scaling_factors_proto_init() }
func file_scaling_factors_proto_init() {
	if File_scaling_factors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scaling_factors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScalingFactors); i {
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
			RawDescriptor: file_scaling_factors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_scaling_factors_proto_goTypes,
		DependencyIndexes: file_scaling_factors_proto_depIdxs,
		MessageInfos:      file_scaling_factors_proto_msgTypes,
	}.Build()
	File_scaling_factors_proto = out.File
	file_scaling_factors_proto_rawDesc = nil
	file_scaling_factors_proto_goTypes = nil
	file_scaling_factors_proto_depIdxs = nil
}
