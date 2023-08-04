// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: download.proto

package pb

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

// DownFileClient is the client API for DownFile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DownFileClient interface {
	//获取下载命令
	DownloadFile(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error)
	//根据包名删除远程机器的下载包
	DeleteFile(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type downFileClient struct {
	cc grpc.ClientConnInterface
}

func NewDownFileClient(cc grpc.ClientConnInterface) DownFileClient {
	return &downFileClient{cc}
}

func (c *downFileClient) DownloadFile(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error) {
	out := new(DownloadResponse)
	err := c.cc.Invoke(ctx, "/pb.DownFile/DownloadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *downFileClient) DeleteFile(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/pb.DownFile/DeleteFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DownFileServer is the server API for DownFile service.
// All implementations must embed UnimplementedDownFileServer
// for forward compatibility
type DownFileServer interface {
	//获取下载命令
	DownloadFile(context.Context, *DownloadRequest) (*DownloadResponse, error)
	//根据包名删除远程机器的下载包
	DeleteFile(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedDownFileServer()
}

// UnimplementedDownFileServer must be embedded to have forward compatible implementations.
type UnimplementedDownFileServer struct {
}

func (UnimplementedDownFileServer) DownloadFile(context.Context, *DownloadRequest) (*DownloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedDownFileServer) DeleteFile(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}
func (UnimplementedDownFileServer) mustEmbedUnimplementedDownFileServer() {}

// UnsafeDownFileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DownFileServer will
// result in compilation errors.
type UnsafeDownFileServer interface {
	mustEmbedUnimplementedDownFileServer()
}

func RegisterDownFileServer(s grpc.ServiceRegistrar, srv DownFileServer) {
	s.RegisterService(&DownFile_ServiceDesc, srv)
}

func _DownFile_DownloadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownFileServer).DownloadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DownFile/DownloadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownFileServer).DownloadFile(ctx, req.(*DownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DownFile_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownFileServer).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DownFile/DeleteFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownFileServer).DeleteFile(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DownFile_ServiceDesc is the grpc.ServiceDesc for DownFile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DownFile_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DownFile",
	HandlerType: (*DownFileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DownloadFile",
			Handler:    _DownFile_DownloadFile_Handler,
		},
		{
			MethodName: "DeleteFile",
			Handler:    _DownFile_DeleteFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "download.proto",
}
