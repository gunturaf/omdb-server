// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpcstub

import (
	context "context"
	entity "github.com/gunturaf/omdb-server/entity"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// OmdbClient is the client API for Omdb service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OmdbClient interface {
	Search(ctx context.Context, in *entity.SearchRequest, opts ...grpc.CallOption) (*entity.SearchReply, error)
	Single(ctx context.Context, in *entity.SingleRequest, opts ...grpc.CallOption) (*entity.SingleReply, error)
}

type omdbClient struct {
	cc grpc.ClientConnInterface
}

func NewOmdbClient(cc grpc.ClientConnInterface) OmdbClient {
	return &omdbClient{cc}
}

func (c *omdbClient) Search(ctx context.Context, in *entity.SearchRequest, opts ...grpc.CallOption) (*entity.SearchReply, error) {
	out := new(entity.SearchReply)
	err := c.cc.Invoke(ctx, "/grpcstub.Omdb/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omdbClient) Single(ctx context.Context, in *entity.SingleRequest, opts ...grpc.CallOption) (*entity.SingleReply, error) {
	out := new(entity.SingleReply)
	err := c.cc.Invoke(ctx, "/grpcstub.Omdb/Single", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OmdbServer is the server API for Omdb service.
// All implementations must embed UnimplementedOmdbServer
// for forward compatibility
type OmdbServer interface {
	Search(context.Context, *entity.SearchRequest) (*entity.SearchReply, error)
	Single(context.Context, *entity.SingleRequest) (*entity.SingleReply, error)
	mustEmbedUnimplementedOmdbServer()
}

// UnimplementedOmdbServer must be embedded to have forward compatible implementations.
type UnimplementedOmdbServer struct {
}

func (UnimplementedOmdbServer) Search(context.Context, *entity.SearchRequest) (*entity.SearchReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedOmdbServer) Single(context.Context, *entity.SingleRequest) (*entity.SingleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Single not implemented")
}
func (UnimplementedOmdbServer) mustEmbedUnimplementedOmdbServer() {}

// UnsafeOmdbServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OmdbServer will
// result in compilation errors.
type UnsafeOmdbServer interface {
	mustEmbedUnimplementedOmdbServer()
}

func RegisterOmdbServer(s grpc.ServiceRegistrar, srv OmdbServer) {
	s.RegisterService(&_Omdb_serviceDesc, srv)
}

func _Omdb_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(entity.SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmdbServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcstub.Omdb/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmdbServer).Search(ctx, req.(*entity.SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Omdb_Single_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(entity.SingleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmdbServer).Single(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcstub.Omdb/Single",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmdbServer).Single(ctx, req.(*entity.SingleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Omdb_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcstub.Omdb",
	HandlerType: (*OmdbServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Omdb_Search_Handler,
		},
		{
			MethodName: "Single",
			Handler:    _Omdb_Single_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "omdbService.proto",
}
