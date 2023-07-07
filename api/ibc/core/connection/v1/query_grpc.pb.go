// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: ibc/core/connection/v1/query.proto

package connectionv1

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
	Query_Connection_FullMethodName               = "/ibc.core.connection.v1.Query/Connection"
	Query_Connections_FullMethodName              = "/ibc.core.connection.v1.Query/Connections"
	Query_ClientConnections_FullMethodName        = "/ibc.core.connection.v1.Query/ClientConnections"
	Query_ConnectionClientState_FullMethodName    = "/ibc.core.connection.v1.Query/ConnectionClientState"
	Query_ConnectionConsensusState_FullMethodName = "/ibc.core.connection.v1.Query/ConnectionConsensusState"
	Query_ConnectionParams_FullMethodName         = "/ibc.core.connection.v1.Query/ConnectionParams"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Connection queries an IBC connection end.
	Connection(ctx context.Context, in *QueryConnectionRequest, opts ...grpc.CallOption) (*QueryConnectionResponse, error)
	// Connections queries all the IBC connections of a chain.
	Connections(ctx context.Context, in *QueryConnectionsRequest, opts ...grpc.CallOption) (*QueryConnectionsResponse, error)
	// ClientConnections queries the connection paths associated with a client
	// state.
	ClientConnections(ctx context.Context, in *QueryClientConnectionsRequest, opts ...grpc.CallOption) (*QueryClientConnectionsResponse, error)
	// ConnectionClientState queries the client state associated with the
	// connection.
	ConnectionClientState(ctx context.Context, in *QueryConnectionClientStateRequest, opts ...grpc.CallOption) (*QueryConnectionClientStateResponse, error)
	// ConnectionConsensusState queries the consensus state associated with the
	// connection.
	ConnectionConsensusState(ctx context.Context, in *QueryConnectionConsensusStateRequest, opts ...grpc.CallOption) (*QueryConnectionConsensusStateResponse, error)
	// ConnectionParams queries all parameters of the ibc connection submodule.
	ConnectionParams(ctx context.Context, in *QueryConnectionParamsRequest, opts ...grpc.CallOption) (*QueryConnectionParamsResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Connection(ctx context.Context, in *QueryConnectionRequest, opts ...grpc.CallOption) (*QueryConnectionResponse, error) {
	out := new(QueryConnectionResponse)
	err := c.cc.Invoke(ctx, Query_Connection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Connections(ctx context.Context, in *QueryConnectionsRequest, opts ...grpc.CallOption) (*QueryConnectionsResponse, error) {
	out := new(QueryConnectionsResponse)
	err := c.cc.Invoke(ctx, Query_Connections_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ClientConnections(ctx context.Context, in *QueryClientConnectionsRequest, opts ...grpc.CallOption) (*QueryClientConnectionsResponse, error) {
	out := new(QueryClientConnectionsResponse)
	err := c.cc.Invoke(ctx, Query_ClientConnections_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ConnectionClientState(ctx context.Context, in *QueryConnectionClientStateRequest, opts ...grpc.CallOption) (*QueryConnectionClientStateResponse, error) {
	out := new(QueryConnectionClientStateResponse)
	err := c.cc.Invoke(ctx, Query_ConnectionClientState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ConnectionConsensusState(ctx context.Context, in *QueryConnectionConsensusStateRequest, opts ...grpc.CallOption) (*QueryConnectionConsensusStateResponse, error) {
	out := new(QueryConnectionConsensusStateResponse)
	err := c.cc.Invoke(ctx, Query_ConnectionConsensusState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ConnectionParams(ctx context.Context, in *QueryConnectionParamsRequest, opts ...grpc.CallOption) (*QueryConnectionParamsResponse, error) {
	out := new(QueryConnectionParamsResponse)
	err := c.cc.Invoke(ctx, Query_ConnectionParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Connection queries an IBC connection end.
	Connection(context.Context, *QueryConnectionRequest) (*QueryConnectionResponse, error)
	// Connections queries all the IBC connections of a chain.
	Connections(context.Context, *QueryConnectionsRequest) (*QueryConnectionsResponse, error)
	// ClientConnections queries the connection paths associated with a client
	// state.
	ClientConnections(context.Context, *QueryClientConnectionsRequest) (*QueryClientConnectionsResponse, error)
	// ConnectionClientState queries the client state associated with the
	// connection.
	ConnectionClientState(context.Context, *QueryConnectionClientStateRequest) (*QueryConnectionClientStateResponse, error)
	// ConnectionConsensusState queries the consensus state associated with the
	// connection.
	ConnectionConsensusState(context.Context, *QueryConnectionConsensusStateRequest) (*QueryConnectionConsensusStateResponse, error)
	// ConnectionParams queries all parameters of the ibc connection submodule.
	ConnectionParams(context.Context, *QueryConnectionParamsRequest) (*QueryConnectionParamsResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Connection(context.Context, *QueryConnectionRequest) (*QueryConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connection not implemented")
}
func (UnimplementedQueryServer) Connections(context.Context, *QueryConnectionsRequest) (*QueryConnectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connections not implemented")
}
func (UnimplementedQueryServer) ClientConnections(context.Context, *QueryClientConnectionsRequest) (*QueryClientConnectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientConnections not implemented")
}
func (UnimplementedQueryServer) ConnectionClientState(context.Context, *QueryConnectionClientStateRequest) (*QueryConnectionClientStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectionClientState not implemented")
}
func (UnimplementedQueryServer) ConnectionConsensusState(context.Context, *QueryConnectionConsensusStateRequest) (*QueryConnectionConsensusStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectionConsensusState not implemented")
}
func (UnimplementedQueryServer) ConnectionParams(context.Context, *QueryConnectionParamsRequest) (*QueryConnectionParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectionParams not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Connection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Connection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Connection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Connection(ctx, req.(*QueryConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Connections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryConnectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Connections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Connections_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Connections(ctx, req.(*QueryConnectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ClientConnections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryClientConnectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ClientConnections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ClientConnections_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ClientConnections(ctx, req.(*QueryClientConnectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ConnectionClientState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryConnectionClientStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ConnectionClientState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ConnectionClientState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ConnectionClientState(ctx, req.(*QueryConnectionClientStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ConnectionConsensusState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryConnectionConsensusStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ConnectionConsensusState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ConnectionConsensusState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ConnectionConsensusState(ctx, req.(*QueryConnectionConsensusStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ConnectionParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryConnectionParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ConnectionParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ConnectionParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ConnectionParams(ctx, req.(*QueryConnectionParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.core.connection.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connection",
			Handler:    _Query_Connection_Handler,
		},
		{
			MethodName: "Connections",
			Handler:    _Query_Connections_Handler,
		},
		{
			MethodName: "ClientConnections",
			Handler:    _Query_ClientConnections_Handler,
		},
		{
			MethodName: "ConnectionClientState",
			Handler:    _Query_ConnectionClientState_Handler,
		},
		{
			MethodName: "ConnectionConsensusState",
			Handler:    _Query_ConnectionConsensusState_Handler,
		},
		{
			MethodName: "ConnectionParams",
			Handler:    _Query_ConnectionParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/core/connection/v1/query.proto",
}
