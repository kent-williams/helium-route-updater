// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: service/poc_lora.proto

package poc_lora

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PocLoraClient is the client API for PocLora service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PocLoraClient interface {
	SubmitLoraBeacon(ctx context.Context, in *LoraBeaconReportReqV1, opts ...grpc.CallOption) (*LoraBeaconReportRespV1, error)
	SubmitLoraWitness(ctx context.Context, in *LoraWitnessReportReqV1, opts ...grpc.CallOption) (*LoraWitnessReportRespV1, error)
}

type pocLoraClient struct {
	cc grpc.ClientConnInterface
}

func NewPocLoraClient(cc grpc.ClientConnInterface) PocLoraClient {
	return &pocLoraClient{cc}
}

func (c *pocLoraClient) SubmitLoraBeacon(ctx context.Context, in *LoraBeaconReportReqV1, opts ...grpc.CallOption) (*LoraBeaconReportRespV1, error) {
	out := new(LoraBeaconReportRespV1)
	err := c.cc.Invoke(ctx, "/helium.poc_lora.poc_lora/submit_lora_beacon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pocLoraClient) SubmitLoraWitness(ctx context.Context, in *LoraWitnessReportReqV1, opts ...grpc.CallOption) (*LoraWitnessReportRespV1, error) {
	out := new(LoraWitnessReportRespV1)
	err := c.cc.Invoke(ctx, "/helium.poc_lora.poc_lora/submit_lora_witness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PocLoraServer is the server API for PocLora service.
// All implementations must embed UnimplementedPocLoraServer
// for forward compatibility
type PocLoraServer interface {
	SubmitLoraBeacon(context.Context, *LoraBeaconReportReqV1) (*LoraBeaconReportRespV1, error)
	SubmitLoraWitness(context.Context, *LoraWitnessReportReqV1) (*LoraWitnessReportRespV1, error)
	mustEmbedUnimplementedPocLoraServer()
}

// UnimplementedPocLoraServer must be embedded to have forward compatible implementations.
type UnimplementedPocLoraServer struct {
}

func (UnimplementedPocLoraServer) SubmitLoraBeacon(context.Context, *LoraBeaconReportReqV1) (*LoraBeaconReportRespV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitLoraBeacon not implemented")
}
func (UnimplementedPocLoraServer) SubmitLoraWitness(context.Context, *LoraWitnessReportReqV1) (*LoraWitnessReportRespV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitLoraWitness not implemented")
}
func (UnimplementedPocLoraServer) mustEmbedUnimplementedPocLoraServer() {}

// UnsafePocLoraServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PocLoraServer will
// result in compilation errors.
type UnsafePocLoraServer interface {
	mustEmbedUnimplementedPocLoraServer()
}

func RegisterPocLoraServer(s grpc.ServiceRegistrar, srv PocLoraServer) {
	s.RegisterService(&PocLora_ServiceDesc, srv)
}

func _PocLora_SubmitLoraBeacon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoraBeaconReportReqV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PocLoraServer).SubmitLoraBeacon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helium.poc_lora.poc_lora/submit_lora_beacon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PocLoraServer).SubmitLoraBeacon(ctx, req.(*LoraBeaconReportReqV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _PocLora_SubmitLoraWitness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoraWitnessReportReqV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PocLoraServer).SubmitLoraWitness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helium.poc_lora.poc_lora/submit_lora_witness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PocLoraServer).SubmitLoraWitness(ctx, req.(*LoraWitnessReportReqV1))
	}
	return interceptor(ctx, in, info, handler)
}

// PocLora_ServiceDesc is the grpc.ServiceDesc for PocLora service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PocLora_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helium.poc_lora.poc_lora",
	HandlerType: (*PocLoraServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "submit_lora_beacon",
			Handler:    _PocLora_SubmitLoraBeacon_Handler,
		},
		{
			MethodName: "submit_lora_witness",
			Handler:    _PocLora_SubmitLoraWitness_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/poc_lora.proto",
}