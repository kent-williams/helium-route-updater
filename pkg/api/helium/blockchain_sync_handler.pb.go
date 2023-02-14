// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: blockchain_sync_handler.proto

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

type BlockchainSyncHash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash    []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Heights []uint64 `protobuf:"varint,2,rep,packed,name=heights,proto3" json:"heights,omitempty"`
}

func (x *BlockchainSyncHash) Reset() {
	*x = BlockchainSyncHash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_sync_handler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainSyncHash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainSyncHash) ProtoMessage() {}

func (x *BlockchainSyncHash) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_sync_handler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainSyncHash.ProtoReflect.Descriptor instead.
func (*BlockchainSyncHash) Descriptor() ([]byte, []int) {
	return file_blockchain_sync_handler_proto_rawDescGZIP(), []int{0}
}

func (x *BlockchainSyncHash) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *BlockchainSyncHash) GetHeights() []uint64 {
	if x != nil {
		return x.Heights
	}
	return nil
}

type BlockchainSyncBlocks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks [][]byte `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
	Final  bool     `protobuf:"varint,2,opt,name=final,proto3" json:"final,omitempty"`
}

func (x *BlockchainSyncBlocks) Reset() {
	*x = BlockchainSyncBlocks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_sync_handler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainSyncBlocks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainSyncBlocks) ProtoMessage() {}

func (x *BlockchainSyncBlocks) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_sync_handler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainSyncBlocks.ProtoReflect.Descriptor instead.
func (*BlockchainSyncBlocks) Descriptor() ([]byte, []int) {
	return file_blockchain_sync_handler_proto_rawDescGZIP(), []int{1}
}

func (x *BlockchainSyncBlocks) GetBlocks() [][]byte {
	if x != nil {
		return x.Blocks
	}
	return nil
}

func (x *BlockchainSyncBlocks) GetFinal() bool {
	if x != nil {
		return x.Final
	}
	return false
}

type BlockchainSyncReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//	*BlockchainSyncReq_Hash
	//	*BlockchainSyncReq_Response
	Msg isBlockchainSyncReq_Msg `protobuf_oneof:"msg"`
}

func (x *BlockchainSyncReq) Reset() {
	*x = BlockchainSyncReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_sync_handler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainSyncReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainSyncReq) ProtoMessage() {}

func (x *BlockchainSyncReq) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_sync_handler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainSyncReq.ProtoReflect.Descriptor instead.
func (*BlockchainSyncReq) Descriptor() ([]byte, []int) {
	return file_blockchain_sync_handler_proto_rawDescGZIP(), []int{2}
}

func (m *BlockchainSyncReq) GetMsg() isBlockchainSyncReq_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *BlockchainSyncReq) GetHash() *BlockchainSyncHash {
	if x, ok := x.GetMsg().(*BlockchainSyncReq_Hash); ok {
		return x.Hash
	}
	return nil
}

func (x *BlockchainSyncReq) GetResponse() bool {
	if x, ok := x.GetMsg().(*BlockchainSyncReq_Response); ok {
		return x.Response
	}
	return false
}

type isBlockchainSyncReq_Msg interface {
	isBlockchainSyncReq_Msg()
}

type BlockchainSyncReq_Hash struct {
	Hash *BlockchainSyncHash `protobuf:"bytes,1,opt,name=hash,proto3,oneof"`
}

type BlockchainSyncReq_Response struct {
	Response bool `protobuf:"varint,2,opt,name=response,proto3,oneof"`
}

func (*BlockchainSyncReq_Hash) isBlockchainSyncReq_Msg() {}

func (*BlockchainSyncReq_Response) isBlockchainSyncReq_Msg() {}

var File_blockchain_sync_handler_proto protoreflect.FileDescriptor

var file_blockchain_sync_handler_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x79, 0x6e,
	0x63, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x44, 0x0a, 0x14, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x79,
	0x6e, 0x63, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04, 0x52, 0x07, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x73, 0x22, 0x46, 0x0a, 0x16, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6e, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x22, 0x67, 0x0a,
	0x13, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x79, 0x6e, 0x63,
	0x5f, 0x72, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f,
	0x73, 0x79, 0x6e, 0x63, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x48, 0x00, 0x52, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x12, 0x1c, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockchain_sync_handler_proto_rawDescOnce sync.Once
	file_blockchain_sync_handler_proto_rawDescData = file_blockchain_sync_handler_proto_rawDesc
)

func file_blockchain_sync_handler_proto_rawDescGZIP() []byte {
	file_blockchain_sync_handler_proto_rawDescOnce.Do(func() {
		file_blockchain_sync_handler_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchain_sync_handler_proto_rawDescData)
	})
	return file_blockchain_sync_handler_proto_rawDescData
}

var file_blockchain_sync_handler_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_blockchain_sync_handler_proto_goTypes = []interface{}{
	(*BlockchainSyncHash)(nil),   // 0: blockchain_sync_hash
	(*BlockchainSyncBlocks)(nil), // 1: blockchain_sync_blocks
	(*BlockchainSyncReq)(nil),    // 2: blockchain_sync_req
}
var file_blockchain_sync_handler_proto_depIdxs = []int32{
	0, // 0: blockchain_sync_req.hash:type_name -> blockchain_sync_hash
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_blockchain_sync_handler_proto_init() }
func file_blockchain_sync_handler_proto_init() {
	if File_blockchain_sync_handler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockchain_sync_handler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainSyncHash); i {
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
		file_blockchain_sync_handler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainSyncBlocks); i {
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
		file_blockchain_sync_handler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainSyncReq); i {
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
	file_blockchain_sync_handler_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*BlockchainSyncReq_Hash)(nil),
		(*BlockchainSyncReq_Response)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blockchain_sync_handler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blockchain_sync_handler_proto_goTypes,
		DependencyIndexes: file_blockchain_sync_handler_proto_depIdxs,
		MessageInfos:      file_blockchain_sync_handler_proto_msgTypes,
	}.Build()
	File_blockchain_sync_handler_proto = out.File
	file_blockchain_sync_handler_proto_rawDesc = nil
	file_blockchain_sync_handler_proto_goTypes = nil
	file_blockchain_sync_handler_proto_depIdxs = nil
}