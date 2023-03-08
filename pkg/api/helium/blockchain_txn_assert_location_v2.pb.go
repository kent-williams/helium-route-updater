// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: blockchain_txn_assert_location_v2.proto

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

type BlockchainTxnAssertLocationV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gateway        []byte `protobuf:"bytes,1,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Owner          []byte `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Payer          []byte `protobuf:"bytes,3,opt,name=payer,proto3" json:"payer,omitempty"`
	OwnerSignature []byte `protobuf:"bytes,4,opt,name=owner_signature,json=ownerSignature,proto3" json:"owner_signature,omitempty"`
	PayerSignature []byte `protobuf:"bytes,5,opt,name=payer_signature,json=payerSignature,proto3" json:"payer_signature,omitempty"`
	// Use the h3 string representation since it will use all 64 bits
	// (and JS doesn't support that)
	Location string `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	Nonce    uint64 `protobuf:"varint,7,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// It is assumed that this gain applies to both rx and tx
	Gain       int32  `protobuf:"varint,8,opt,name=gain,proto3" json:"gain,omitempty"`
	Elevation  int32  `protobuf:"varint,9,opt,name=elevation,proto3" json:"elevation,omitempty"`
	StakingFee uint64 `protobuf:"varint,10,opt,name=staking_fee,json=stakingFee,proto3" json:"staking_fee,omitempty"`
	Fee        uint64 `protobuf:"varint,11,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (x *BlockchainTxnAssertLocationV2) Reset() {
	*x = BlockchainTxnAssertLocationV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_txn_assert_location_v2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainTxnAssertLocationV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainTxnAssertLocationV2) ProtoMessage() {}

func (x *BlockchainTxnAssertLocationV2) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_txn_assert_location_v2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainTxnAssertLocationV2.ProtoReflect.Descriptor instead.
func (*BlockchainTxnAssertLocationV2) Descriptor() ([]byte, []int) {
	return file_blockchain_txn_assert_location_v2_proto_rawDescGZIP(), []int{0}
}

func (x *BlockchainTxnAssertLocationV2) GetGateway() []byte {
	if x != nil {
		return x.Gateway
	}
	return nil
}

func (x *BlockchainTxnAssertLocationV2) GetOwner() []byte {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *BlockchainTxnAssertLocationV2) GetPayer() []byte {
	if x != nil {
		return x.Payer
	}
	return nil
}

func (x *BlockchainTxnAssertLocationV2) GetOwnerSignature() []byte {
	if x != nil {
		return x.OwnerSignature
	}
	return nil
}

func (x *BlockchainTxnAssertLocationV2) GetPayerSignature() []byte {
	if x != nil {
		return x.PayerSignature
	}
	return nil
}

func (x *BlockchainTxnAssertLocationV2) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *BlockchainTxnAssertLocationV2) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *BlockchainTxnAssertLocationV2) GetGain() int32 {
	if x != nil {
		return x.Gain
	}
	return 0
}

func (x *BlockchainTxnAssertLocationV2) GetElevation() int32 {
	if x != nil {
		return x.Elevation
	}
	return 0
}

func (x *BlockchainTxnAssertLocationV2) GetStakingFee() uint64 {
	if x != nil {
		return x.StakingFee
	}
	return 0
}

func (x *BlockchainTxnAssertLocationV2) GetFee() uint64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

var File_blockchain_txn_assert_location_v2_proto protoreflect.FileDescriptor

var file_blockchain_txn_assert_location_v2_proto_rawDesc = []byte{
	0x0a, 0x27, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x78, 0x6e,
	0x5f, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x68, 0x65, 0x6c, 0x69, 0x75,
	0x6d, 0x22, 0xd2, 0x02, 0x0a, 0x21, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x5f, 0x74, 0x78, 0x6e, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x79, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x61, 0x79, 0x65, 0x72, 0x12, 0x27, 0x0a,
	0x0f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0e, 0x70, 0x61, 0x79, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x6f, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x61, 0x69, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x67, 0x61, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x66,
	0x65, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e,
	0x67, 0x46, 0x65, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x66, 0x65, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockchain_txn_assert_location_v2_proto_rawDescOnce sync.Once
	file_blockchain_txn_assert_location_v2_proto_rawDescData = file_blockchain_txn_assert_location_v2_proto_rawDesc
)

func file_blockchain_txn_assert_location_v2_proto_rawDescGZIP() []byte {
	file_blockchain_txn_assert_location_v2_proto_rawDescOnce.Do(func() {
		file_blockchain_txn_assert_location_v2_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchain_txn_assert_location_v2_proto_rawDescData)
	})
	return file_blockchain_txn_assert_location_v2_proto_rawDescData
}

var file_blockchain_txn_assert_location_v2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_blockchain_txn_assert_location_v2_proto_goTypes = []interface{}{
	(*BlockchainTxnAssertLocationV2)(nil), // 0: helium.blockchain_txn_assert_location_v2
}
var file_blockchain_txn_assert_location_v2_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_blockchain_txn_assert_location_v2_proto_init() }
func file_blockchain_txn_assert_location_v2_proto_init() {
	if File_blockchain_txn_assert_location_v2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockchain_txn_assert_location_v2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainTxnAssertLocationV2); i {
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
			RawDescriptor: file_blockchain_txn_assert_location_v2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blockchain_txn_assert_location_v2_proto_goTypes,
		DependencyIndexes: file_blockchain_txn_assert_location_v2_proto_depIdxs,
		MessageInfos:      file_blockchain_txn_assert_location_v2_proto_msgTypes,
	}.Build()
	File_blockchain_txn_assert_location_v2_proto = out.File
	file_blockchain_txn_assert_location_v2_proto_rawDesc = nil
	file_blockchain_txn_assert_location_v2_proto_goTypes = nil
	file_blockchain_txn_assert_location_v2_proto_depIdxs = nil
}
