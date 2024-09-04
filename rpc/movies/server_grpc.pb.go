// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.19.4
// source: server.proto

package movies

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	MoviesServer_Ping_FullMethodName = "/server.MoviesServer/Ping"
)

// MoviesServerClient is the client API for MoviesServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MoviesServerClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type moviesServerClient struct {
	cc grpc.ClientConnInterface
}

func NewMoviesServerClient(cc grpc.ClientConnInterface) MoviesServerClient {
	return &moviesServerClient{cc}
}

func (c *moviesServerClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, MoviesServer_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MoviesServerServer is the server API for MoviesServer service.
// All implementations must embed UnimplementedMoviesServerServer
// for forward compatibility
type MoviesServerServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedMoviesServerServer()
}

// UnimplementedMoviesServerServer must be embedded to have forward compatible implementations.
type UnimplementedMoviesServerServer struct {
}

func (UnimplementedMoviesServerServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedMoviesServerServer) mustEmbedUnimplementedMoviesServerServer() {}

// UnsafeMoviesServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MoviesServerServer will
// result in compilation errors.
type UnsafeMoviesServerServer interface {
	mustEmbedUnimplementedMoviesServerServer()
}

func RegisterMoviesServerServer(s grpc.ServiceRegistrar, srv MoviesServerServer) {
	s.RegisterService(&MoviesServer_ServiceDesc, srv)
}

func _MoviesServer_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesServerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MoviesServer_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesServerServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// MoviesServer_ServiceDesc is the grpc.ServiceDesc for MoviesServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MoviesServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.MoviesServer",
	HandlerType: (*MoviesServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _MoviesServer_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
