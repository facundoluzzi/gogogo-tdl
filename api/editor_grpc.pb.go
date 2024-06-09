// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: api/editor.proto

package api

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
	TextEditor_ReadFile_FullMethodName   = "/editor.TextEditor/ReadFile"
	TextEditor_SaveFile_FullMethodName = "/editor.TextEditor/SaveFile"
)

// TextEditorClient is the client API for TextEditor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TextEditorClient interface {
	ReadFile(ctx context.Context, in *ReadFileRequest, opts ...grpc.CallOption) (*ReadFileResponse, error)
	SaveFile(ctx context.Context, in *SaveFileRequest, opts ...grpc.CallOption) (*SaveFileResponse, error)
}

type textEditorClient struct {
	cc grpc.ClientConnInterface
}

func NewTextEditorClient(cc grpc.ClientConnInterface) TextEditorClient {
	return &textEditorClient{cc}
}

func (c *textEditorClient) ReadFile(ctx context.Context, in *ReadFileRequest, opts ...grpc.CallOption) (*ReadFileResponse, error) {
	out := new(ReadFileResponse)
	err := c.cc.Invoke(ctx, TextEditor_ReadFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textEditorClient) SaveFile(ctx context.Context, in *SaveFileRequest, opts ...grpc.CallOption) (*SaveFileResponse, error) {
	out := new(SaveFileResponse)
	err := c.cc.Invoke(ctx, TextEditor_SaveFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TextEditorServer is the server API for TextEditor service.
// All implementations must embed UnimplementedTextEditorServer
// for forward compatibility
type TextEditorServer interface {
	ReadFile(context.Context, *ReadFileRequest) (*ReadFileResponse, error)
	SaveFile(context.Context, *SaveFileRequest) (*SaveFileResponse, error)
	mustEmbedUnimplementedTextEditorServer()
}

// UnimplementedTextEditorServer must be embedded to have forward compatible implementations.
type UnimplementedTextEditorServer struct {
}

func (UnimplementedTextEditorServer) ReadFile(context.Context, *ReadFileRequest) (*ReadFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadFile not implemented")
}
func (UnimplementedTextEditorServer) SaveFile(context.Context, *SaveFileRequest) (*SaveFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveFile not implemented")
}
func (UnimplementedTextEditorServer) mustEmbedUnimplementedTextEditorServer() {}

// UnsafeTextEditorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TextEditorServer will
// result in compilation errors.
type UnsafeTextEditorServer interface {
	mustEmbedUnimplementedTextEditorServer()
}

func RegisterTextEditorServer(s grpc.ServiceRegistrar, srv TextEditorServer) {
	s.RegisterService(&TextEditor_ServiceDesc, srv)
}

func _TextEditor_ReadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextEditorServer).ReadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TextEditor_ReadFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextEditorServer).ReadFile(ctx, req.(*ReadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TextEditor_SaveFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextEditorServer).SaveFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TextEditor_SaveFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextEditorServer).SaveFile(ctx, req.(*SaveFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TextEditor_ServiceDesc is the grpc.ServiceDesc for TextEditor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TextEditor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "editor.TextEditor",
	HandlerType: (*TextEditorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadFile",
			Handler:    _TextEditor_ReadFile_Handler,
		},
		{
			MethodName: "SaveFile",
			Handler:    _TextEditor_SaveFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/editor.proto",
}
