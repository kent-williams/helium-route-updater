// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: blockchain_txn_consensus_group_v1.proto

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

type BlockchainTxnConsensusGroupV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Members [][]byte `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	Proof   []byte   `protobuf:"bytes,2,opt,name=proof,proto3" json:"proof,omitempty"`
	Height  uint64   `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Delay   uint32   `protobuf:"varint,4,opt,name=delay,proto3" json:"delay,omitempty"`
}

func (x *BlockchainTxnConsensusGroupV1) Reset() {
	*x = BlockchainTxnConsensusGroupV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_txn_consensus_group_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainTxnConsensusGroupV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainTxnConsensusGroupV1) ProtoMessage() {}

func (x *BlockchainTxnConsensusGroupV1) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_txn_consensus_group_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainTxnConsensusGroupV1.ProtoReflect.Descriptor instead.
func (*BlockchainTxnConsensusGroupV1) Descriptor() ([]byte, []int) {
	return file_blockchain_txn_consensus_group_v1_proto_rawDescGZIP(), []int{0}
}

func (x *BlockchainTxnConsensusGroupV1) GetMembers() [][]byte {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *BlockchainTxnConsensusGroupV1) GetProof() []byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

func (x *BlockchainTxnConsensusGroupV1) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *BlockchainTxnConsensusGroupV1) GetDelay() uint32 {
	if x != nil {
		return x.Delay
	}
	return 0
}

var File_blockchain_txn_consensus_group_v1_proto protoreflect.FileDescriptor

var file_blockchain_txn_consensus_group_v1_proto_rawDesc = []byte{
	0x0a, 0x27, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x78, 0x6e,
	0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x68, 0x65, 0x6c, 0x69, 0x75,
	0x6d, 0x22, 0x81, 0x01, 0x0a, 0x21, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x5f, 0x74, 0x78, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x76, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x64, 0x65, 0x6c, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockchain_txn_consensus_group_v1_proto_rawDescOnce sync.Once
	file_blockchain_txn_consensus_group_v1_proto_rawDescData = file_blockchain_txn_consensus_group_v1_proto_rawDesc
)

func file_blockchain_txn_consensus_group_v1_proto_rawDescGZIP() []byte {
	file_blockchain_txn_consensus_group_v1_proto_rawDescOnce.Do(func() {
		file_blockchain_txn_consensus_group_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchain_txn_consensus_group_v1_proto_rawDescData)
	})
	return file_blockchain_txn_consensus_group_v1_proto_rawDescData
}

var file_blockchain_txn_consensus_group_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_blockchain_txn_consensus_group_v1_proto_goTypes = []interface{}{
	(*BlockchainTxnConsensusGroupV1)(nil), // 0: helium.blockchain_txn_consensus_group_v1
}
var file_blockchain_txn_consensus_group_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_blockchain_txn_consensus_group_v1_proto_init() }
func file_blockchain_txn_consensus_group_v1_proto_init() {
	if File_blockchain_txn_consensus_group_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockchain_txn_consensus_group_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainTxnConsensusGroupV1); i {
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
			RawDescriptor: file_blockchain_txn_consensus_group_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blockchain_txn_consensus_group_v1_proto_goTypes,
		DependencyIndexes: file_blockchain_txn_consensus_group_v1_proto_depIdxs,
		MessageInfos:      file_blockchain_txn_consensus_group_v1_proto_msgTypes,
	}.Build()
	File_blockchain_txn_consensus_group_v1_proto = out.File
	file_blockchain_txn_consensus_group_v1_proto_rawDesc = nil
	file_blockchain_txn_consensus_group_v1_proto_goTypes = nil
	file_blockchain_txn_consensus_group_v1_proto_depIdxs = nil
}
