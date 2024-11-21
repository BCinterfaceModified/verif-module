// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: blockchain_interface.proto

package bcinterface

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

const (
	BlockchainInterface_EnrollNodeInfo_FullMethodName = "/commu_module.BlockchainInterface/EnrollNodeInfo"
	BlockchainInterface_SetupCommittee_FullMethodName = "/commu_module.BlockchainInterface/SetupCommittee"
	BlockchainInterface_LeaveRequest_FullMethodName   = "/commu_module.BlockchainInterface/LeaveRequest"
)

// BlockchainInterfaceClient is the client API for BlockchainInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlockchainInterfaceClient interface {
	EnrollNodeInfo(ctx context.Context, in *NodeData, opts ...grpc.CallOption) (*EnrollAccountResponse, error)
	SetupCommittee(ctx context.Context, in *SetupCommitteeRequest, opts ...grpc.CallOption) (*SetupCommitteeResponse, error)
	LeaveRequest(ctx context.Context, in *NodeData, opts ...grpc.CallOption) (*Empty, error)
}

type blockchainInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockchainInterfaceClient(cc grpc.ClientConnInterface) BlockchainInterfaceClient {
	return &blockchainInterfaceClient{cc}
}

func (c *blockchainInterfaceClient) EnrollNodeInfo(ctx context.Context, in *NodeData, opts ...grpc.CallOption) (*EnrollAccountResponse, error) {
	out := new(EnrollAccountResponse)
	err := c.cc.Invoke(ctx, BlockchainInterface_EnrollNodeInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainInterfaceClient) SetupCommittee(ctx context.Context, in *SetupCommitteeRequest, opts ...grpc.CallOption) (*SetupCommitteeResponse, error) {
	out := new(SetupCommitteeResponse)
	err := c.cc.Invoke(ctx, BlockchainInterface_SetupCommittee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainInterfaceClient) LeaveRequest(ctx context.Context, in *NodeData, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, BlockchainInterface_LeaveRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockchainInterfaceServer is the server API for BlockchainInterface service.
// All implementations must embed UnimplementedBlockchainInterfaceServer
// for forward compatibility
type BlockchainInterfaceServer interface {
	EnrollNodeInfo(context.Context, *NodeData) (*EnrollAccountResponse, error)
	SetupCommittee(context.Context, *SetupCommitteeRequest) (*SetupCommitteeResponse, error)
	LeaveRequest(context.Context, *NodeData) (*Empty, error)
	mustEmbedUnimplementedBlockchainInterfaceServer()
}

// UnimplementedBlockchainInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedBlockchainInterfaceServer struct {
}

func (UnimplementedBlockchainInterfaceServer) EnrollNodeInfo(context.Context, *NodeData) (*EnrollAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnrollNodeInfo not implemented")
}
func (UnimplementedBlockchainInterfaceServer) SetupCommittee(context.Context, *SetupCommitteeRequest) (*SetupCommitteeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetupCommittee not implemented")
}
func (UnimplementedBlockchainInterfaceServer) LeaveRequest(context.Context, *NodeData) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveRequest not implemented")
}
func (UnimplementedBlockchainInterfaceServer) mustEmbedUnimplementedBlockchainInterfaceServer() {}

// UnsafeBlockchainInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockchainInterfaceServer will
// result in compilation errors.
type UnsafeBlockchainInterfaceServer interface {
	mustEmbedUnimplementedBlockchainInterfaceServer()
}

func RegisterBlockchainInterfaceServer(s grpc.ServiceRegistrar, srv BlockchainInterfaceServer) {
	s.RegisterService(&BlockchainInterface_ServiceDesc, srv)
}

func _BlockchainInterface_EnrollNodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainInterfaceServer).EnrollNodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlockchainInterface_EnrollNodeInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainInterfaceServer).EnrollNodeInfo(ctx, req.(*NodeData))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainInterface_SetupCommittee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetupCommitteeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainInterfaceServer).SetupCommittee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlockchainInterface_SetupCommittee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainInterfaceServer).SetupCommittee(ctx, req.(*SetupCommitteeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainInterface_LeaveRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainInterfaceServer).LeaveRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlockchainInterface_LeaveRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainInterfaceServer).LeaveRequest(ctx, req.(*NodeData))
	}
	return interceptor(ctx, in, info, handler)
}

// BlockchainInterface_ServiceDesc is the grpc.ServiceDesc for BlockchainInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlockchainInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "commu_module.BlockchainInterface",
	HandlerType: (*BlockchainInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnrollNodeInfo",
			Handler:    _BlockchainInterface_EnrollNodeInfo_Handler,
		},
		{
			MethodName: "SetupCommittee",
			Handler:    _BlockchainInterface_SetupCommittee_Handler,
		},
		{
			MethodName: "LeaveRequest",
			Handler:    _BlockchainInterface_LeaveRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blockchain_interface.proto",
}

const (
	NetworkInterface_RequestLeader_FullMethodName = "/commu_module.NetworkInterface/RequestLeader"
)

// NetworkInterfaceClient is the client API for NetworkInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetworkInterfaceClient interface {
	RequestLeader(ctx context.Context, in *RequestLeaderRequest, opts ...grpc.CallOption) (*RequestLeaderResponse, error)
}

type networkInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkInterfaceClient(cc grpc.ClientConnInterface) NetworkInterfaceClient {
	return &networkInterfaceClient{cc}
}

func (c *networkInterfaceClient) RequestLeader(ctx context.Context, in *RequestLeaderRequest, opts ...grpc.CallOption) (*RequestLeaderResponse, error) {
	out := new(RequestLeaderResponse)
	err := c.cc.Invoke(ctx, NetworkInterface_RequestLeader_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkInterfaceServer is the server API for NetworkInterface service.
// All implementations must embed UnimplementedNetworkInterfaceServer
// for forward compatibility
type NetworkInterfaceServer interface {
	RequestLeader(context.Context, *RequestLeaderRequest) (*RequestLeaderResponse, error)
	mustEmbedUnimplementedNetworkInterfaceServer()
}

// UnimplementedNetworkInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedNetworkInterfaceServer struct {
}

func (UnimplementedNetworkInterfaceServer) RequestLeader(context.Context, *RequestLeaderRequest) (*RequestLeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestLeader not implemented")
}
func (UnimplementedNetworkInterfaceServer) mustEmbedUnimplementedNetworkInterfaceServer() {}

// UnsafeNetworkInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkInterfaceServer will
// result in compilation errors.
type UnsafeNetworkInterfaceServer interface {
	mustEmbedUnimplementedNetworkInterfaceServer()
}

func RegisterNetworkInterfaceServer(s grpc.ServiceRegistrar, srv NetworkInterfaceServer) {
	s.RegisterService(&NetworkInterface_ServiceDesc, srv)
}

func _NetworkInterface_RequestLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestLeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkInterfaceServer).RequestLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkInterface_RequestLeader_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkInterfaceServer).RequestLeader(ctx, req.(*RequestLeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NetworkInterface_ServiceDesc is the grpc.ServiceDesc for NetworkInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetworkInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "commu_module.NetworkInterface",
	HandlerType: (*NetworkInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestLeader",
			Handler:    _NetworkInterface_RequestLeader_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blockchain_interface.proto",
}
