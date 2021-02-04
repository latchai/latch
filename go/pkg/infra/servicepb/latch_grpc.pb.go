// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package servicepb

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

// DaemonClient is the client API for Daemon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DaemonClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	LaunchJob(ctx context.Context, in *LaunchJobRequest, opts ...grpc.CallOption) (*LaunchJobReply, error)
}

type daemonClient struct {
	cc grpc.ClientConnInterface
}

func NewDaemonClient(cc grpc.ClientConnInterface) DaemonClient {
	return &daemonClient{cc}
}

func (c *daemonClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/service.Daemon/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) LaunchJob(ctx context.Context, in *LaunchJobRequest, opts ...grpc.CallOption) (*LaunchJobReply, error) {
	out := new(LaunchJobReply)
	err := c.cc.Invoke(ctx, "/service.Daemon/LaunchJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DaemonServer is the server API for Daemon service.
// All implementations must embed UnimplementedDaemonServer
// for forward compatibility
type DaemonServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	LaunchJob(context.Context, *LaunchJobRequest) (*LaunchJobReply, error)
	mustEmbedUnimplementedDaemonServer()
}

// UnimplementedDaemonServer must be embedded to have forward compatible implementations.
type UnimplementedDaemonServer struct {
}

func (UnimplementedDaemonServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedDaemonServer) LaunchJob(context.Context, *LaunchJobRequest) (*LaunchJobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LaunchJob not implemented")
}
func (UnimplementedDaemonServer) mustEmbedUnimplementedDaemonServer() {}

// UnsafeDaemonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DaemonServer will
// result in compilation errors.
type UnsafeDaemonServer interface {
	mustEmbedUnimplementedDaemonServer()
}

func RegisterDaemonServer(s grpc.ServiceRegistrar, srv DaemonServer) {
	s.RegisterService(&Daemon_ServiceDesc, srv)
}

func _Daemon_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Daemon/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_LaunchJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LaunchJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).LaunchJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Daemon/LaunchJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).LaunchJob(ctx, req.(*LaunchJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Daemon_ServiceDesc is the grpc.ServiceDesc for Daemon service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Daemon_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Daemon",
	HandlerType: (*DaemonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Daemon_Ping_Handler,
		},
		{
			MethodName: "LaunchJob",
			Handler:    _Daemon_LaunchJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "latch.proto",
}
