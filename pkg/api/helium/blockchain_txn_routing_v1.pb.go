// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: blockchain_txn_routing_v1.proto

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

type UpdateRouters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouterAddresses [][]byte `protobuf:"bytes,1,rep,name=router_addresses,json=routerAddresses,proto3" json:"router_addresses,omitempty"` // max of 3
}

func (x *UpdateRouters) Reset() {
	*x = UpdateRouters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_txn_routing_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRouters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRouters) ProtoMessage() {}

func (x *UpdateRouters) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_txn_routing_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRouters.ProtoReflect.Descriptor instead.
func (*UpdateRouters) Descriptor() ([]byte, []int) {
	return file_blockchain_txn_routing_v1_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateRouters) GetRouterAddresses() [][]byte {
	if x != nil {
		return x.RouterAddresses
	}
	return nil
}

type UpdateXor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index  uint32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`  // we allow up to 5
	Filter []byte `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"` // 100kb or less
}

func (x *UpdateXor) Reset() {
	*x = UpdateXor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_txn_routing_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateXor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateXor) ProtoMessage() {}

func (x *UpdateXor) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_txn_routing_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateXor.ProtoReflect.Descriptor instead.
func (*UpdateXor) Descriptor() ([]byte, []int) {
	return file_blockchain_txn_routing_v1_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateXor) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *UpdateXor) GetFilter() []byte {
	if x != nil {
		return x.Filter
	}
	return nil
}

type BlockchainTxnRoutingV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Oui   uint32 `protobuf:"varint,1,opt,name=oui,proto3" json:"oui,omitempty"`
	Owner []byte `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// Types that are assignable to Update:
	//	*BlockchainTxnRoutingV1_UpdateRouters
	//	*BlockchainTxnRoutingV1_NewXor
	//	*BlockchainTxnRoutingV1_UpdateXor
	//	*BlockchainTxnRoutingV1_RequestSubnet
	Update     isBlockchainTxnRoutingV1_Update `protobuf_oneof:"update"`
	Fee        uint64                          `protobuf:"varint,7,opt,name=fee,proto3" json:"fee,omitempty"`
	Nonce      uint64                          `protobuf:"varint,8,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Signature  []byte                          `protobuf:"bytes,9,opt,name=signature,proto3" json:"signature,omitempty"`
	StakingFee uint64                          `protobuf:"varint,10,opt,name=staking_fee,json=stakingFee,proto3" json:"staking_fee,omitempty"`
}

func (x *BlockchainTxnRoutingV1) Reset() {
	*x = BlockchainTxnRoutingV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_txn_routing_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainTxnRoutingV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainTxnRoutingV1) ProtoMessage() {}

func (x *BlockchainTxnRoutingV1) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_txn_routing_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainTxnRoutingV1.ProtoReflect.Descriptor instead.
func (*BlockchainTxnRoutingV1) Descriptor() ([]byte, []int) {
	return file_blockchain_txn_routing_v1_proto_rawDescGZIP(), []int{2}
}

func (x *BlockchainTxnRoutingV1) GetOui() uint32 {
	if x != nil {
		return x.Oui
	}
	return 0
}

func (x *BlockchainTxnRoutingV1) GetOwner() []byte {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (m *BlockchainTxnRoutingV1) GetUpdate() isBlockchainTxnRoutingV1_Update {
	if m != nil {
		return m.Update
	}
	return nil
}

func (x *BlockchainTxnRoutingV1) GetUpdateRouters() *UpdateRouters {
	if x, ok := x.GetUpdate().(*BlockchainTxnRoutingV1_UpdateRouters); ok {
		return x.UpdateRouters
	}
	return nil
}

func (x *BlockchainTxnRoutingV1) GetNewXor() []byte {
	if x, ok := x.GetUpdate().(*BlockchainTxnRoutingV1_NewXor); ok {
		return x.NewXor
	}
	return nil
}

func (x *BlockchainTxnRoutingV1) GetUpdateXor() *UpdateXor {
	if x, ok := x.GetUpdate().(*BlockchainTxnRoutingV1_UpdateXor); ok {
		return x.UpdateXor
	}
	return nil
}

func (x *BlockchainTxnRoutingV1) GetRequestSubnet() uint32 {
	if x, ok := x.GetUpdate().(*BlockchainTxnRoutingV1_RequestSubnet); ok {
		return x.RequestSubnet
	}
	return 0
}

func (x *BlockchainTxnRoutingV1) GetFee() uint64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *BlockchainTxnRoutingV1) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *BlockchainTxnRoutingV1) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *BlockchainTxnRoutingV1) GetStakingFee() uint64 {
	if x != nil {
		return x.StakingFee
	}
	return 0
}

type isBlockchainTxnRoutingV1_Update interface {
	isBlockchainTxnRoutingV1_Update()
}

type BlockchainTxnRoutingV1_UpdateRouters struct {
	UpdateRouters *UpdateRouters `protobuf:"bytes,3,opt,name=update_routers,json=updateRouters,proto3,oneof"`
}

type BlockchainTxnRoutingV1_NewXor struct {
	NewXor []byte `protobuf:"bytes,4,opt,name=new_xor,json=newXor,proto3,oneof"` // 100kb or less
}

type BlockchainTxnRoutingV1_UpdateXor struct {
	UpdateXor *UpdateXor `protobuf:"bytes,5,opt,name=update_xor,json=updateXor,proto3,oneof"`
}

type BlockchainTxnRoutingV1_RequestSubnet struct {
	RequestSubnet uint32 `protobuf:"varint,6,opt,name=request_subnet,json=requestSubnet,proto3,oneof"` // number of addresses to request
}

func (*BlockchainTxnRoutingV1_UpdateRouters) isBlockchainTxnRoutingV1_Update() {}

func (*BlockchainTxnRoutingV1_NewXor) isBlockchainTxnRoutingV1_Update() {}

func (*BlockchainTxnRoutingV1_UpdateXor) isBlockchainTxnRoutingV1_Update() {}

func (*BlockchainTxnRoutingV1_RequestSubnet) isBlockchainTxnRoutingV1_Update() {}

var File_blockchain_txn_routing_v1_proto protoreflect.FileDescriptor

var file_blockchain_txn_routing_v1_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x78, 0x6e,
	0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x22, 0x3b, 0x0a, 0x0e, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x3a, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x78, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x22, 0xee, 0x02, 0x0a, 0x19, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x5f, 0x74, 0x78, 0x6e, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x31,
	0x12, 0x10, 0x0a, 0x03, 0x6f, 0x75, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6f,
	0x75, 0x69, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0e, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x48, 0x00, 0x52, 0x0d, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x12, 0x19, 0x0a, 0x07, 0x6e, 0x65, 0x77,
	0x5f, 0x78, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x06, 0x6e, 0x65,
	0x77, 0x58, 0x6f, 0x72, 0x12, 0x33, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x78,
	0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75,
	0x6d, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x78, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x58, 0x6f, 0x72, 0x12, 0x27, 0x0a, 0x0e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x48, 0x00, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6e,
	0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x66, 0x65, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x6b,
	0x69, 0x6e, 0x67, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x46, 0x65, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockchain_txn_routing_v1_proto_rawDescOnce sync.Once
	file_blockchain_txn_routing_v1_proto_rawDescData = file_blockchain_txn_routing_v1_proto_rawDesc
)

func file_blockchain_txn_routing_v1_proto_rawDescGZIP() []byte {
	file_blockchain_txn_routing_v1_proto_rawDescOnce.Do(func() {
		file_blockchain_txn_routing_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchain_txn_routing_v1_proto_rawDescData)
	})
	return file_blockchain_txn_routing_v1_proto_rawDescData
}

var file_blockchain_txn_routing_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_blockchain_txn_routing_v1_proto_goTypes = []interface{}{
	(*UpdateRouters)(nil),          // 0: helium.update_routers
	(*UpdateXor)(nil),              // 1: helium.update_xor
	(*BlockchainTxnRoutingV1)(nil), // 2: helium.blockchain_txn_routing_v1
}
var file_blockchain_txn_routing_v1_proto_depIdxs = []int32{
	0, // 0: helium.blockchain_txn_routing_v1.update_routers:type_name -> helium.update_routers
	1, // 1: helium.blockchain_txn_routing_v1.update_xor:type_name -> helium.update_xor
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_blockchain_txn_routing_v1_proto_init() }
func file_blockchain_txn_routing_v1_proto_init() {
	if File_blockchain_txn_routing_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockchain_txn_routing_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRouters); i {
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
		file_blockchain_txn_routing_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateXor); i {
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
		file_blockchain_txn_routing_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainTxnRoutingV1); i {
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
	file_blockchain_txn_routing_v1_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*BlockchainTxnRoutingV1_UpdateRouters)(nil),
		(*BlockchainTxnRoutingV1_NewXor)(nil),
		(*BlockchainTxnRoutingV1_UpdateXor)(nil),
		(*BlockchainTxnRoutingV1_RequestSubnet)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blockchain_txn_routing_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blockchain_txn_routing_v1_proto_goTypes,
		DependencyIndexes: file_blockchain_txn_routing_v1_proto_depIdxs,
		MessageInfos:      file_blockchain_txn_routing_v1_proto_msgTypes,
	}.Build()
	File_blockchain_txn_routing_v1_proto = out.File
	file_blockchain_txn_routing_v1_proto_rawDesc = nil
	file_blockchain_txn_routing_v1_proto_goTypes = nil
	file_blockchain_txn_routing_v1_proto_depIdxs = nil
}