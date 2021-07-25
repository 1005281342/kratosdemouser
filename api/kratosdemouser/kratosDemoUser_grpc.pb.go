// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package kratosdemouser

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

// KratosDemoUserClient is the client API for KratosDemoUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KratosDemoUserClient interface {
	CreateKratosDemoUser(ctx context.Context, in *CreateKratosDemoUserRequest, opts ...grpc.CallOption) (*CreateKratosDemoUserReply, error)
	UpdateKratosDemoUser(ctx context.Context, in *UpdateKratosDemoUserRequest, opts ...grpc.CallOption) (*UpdateKratosDemoUserReply, error)
	DeleteKratosDemoUser(ctx context.Context, in *DeleteKratosDemoUserRequest, opts ...grpc.CallOption) (*DeleteKratosDemoUserReply, error)
	GetKratosDemoUser(ctx context.Context, in *GetKratosDemoUserRequest, opts ...grpc.CallOption) (*GetKratosDemoUserReply, error)
	Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRsp, error)
}

type kratosDemoUserClient struct {
	cc grpc.ClientConnInterface
}

func NewKratosDemoUserClient(cc grpc.ClientConnInterface) KratosDemoUserClient {
	return &kratosDemoUserClient{cc}
}

func (c *kratosDemoUserClient) CreateKratosDemoUser(ctx context.Context, in *CreateKratosDemoUserRequest, opts ...grpc.CallOption) (*CreateKratosDemoUserReply, error) {
	out := new(CreateKratosDemoUserReply)
	err := c.cc.Invoke(ctx, "/api.kratosdemouser.KratosDemoUser/CreateKratosDemoUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kratosDemoUserClient) UpdateKratosDemoUser(ctx context.Context, in *UpdateKratosDemoUserRequest, opts ...grpc.CallOption) (*UpdateKratosDemoUserReply, error) {
	out := new(UpdateKratosDemoUserReply)
	err := c.cc.Invoke(ctx, "/api.kratosdemouser.KratosDemoUser/UpdateKratosDemoUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kratosDemoUserClient) DeleteKratosDemoUser(ctx context.Context, in *DeleteKratosDemoUserRequest, opts ...grpc.CallOption) (*DeleteKratosDemoUserReply, error) {
	out := new(DeleteKratosDemoUserReply)
	err := c.cc.Invoke(ctx, "/api.kratosdemouser.KratosDemoUser/DeleteKratosDemoUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kratosDemoUserClient) GetKratosDemoUser(ctx context.Context, in *GetKratosDemoUserRequest, opts ...grpc.CallOption) (*GetKratosDemoUserReply, error) {
	out := new(GetKratosDemoUserReply)
	err := c.cc.Invoke(ctx, "/api.kratosdemouser.KratosDemoUser/GetKratosDemoUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kratosDemoUserClient) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRsp, error) {
	out := new(PingRsp)
	err := c.cc.Invoke(ctx, "/api.kratosdemouser.KratosDemoUser/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KratosDemoUserServer is the server API for KratosDemoUser service.
// All implementations must embed UnimplementedKratosDemoUserServer
// for forward compatibility
type KratosDemoUserServer interface {
	CreateKratosDemoUser(context.Context, *CreateKratosDemoUserRequest) (*CreateKratosDemoUserReply, error)
	UpdateKratosDemoUser(context.Context, *UpdateKratosDemoUserRequest) (*UpdateKratosDemoUserReply, error)
	DeleteKratosDemoUser(context.Context, *DeleteKratosDemoUserRequest) (*DeleteKratosDemoUserReply, error)
	GetKratosDemoUser(context.Context, *GetKratosDemoUserRequest) (*GetKratosDemoUserReply, error)
	Ping(context.Context, *PingReq) (*PingRsp, error)
	mustEmbedUnimplementedKratosDemoUserServer()
}

// UnimplementedKratosDemoUserServer must be embedded to have forward compatible implementations.
type UnimplementedKratosDemoUserServer struct {
}

func (UnimplementedKratosDemoUserServer) CreateKratosDemoUser(context.Context, *CreateKratosDemoUserRequest) (*CreateKratosDemoUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKratosDemoUser not implemented")
}
func (UnimplementedKratosDemoUserServer) UpdateKratosDemoUser(context.Context, *UpdateKratosDemoUserRequest) (*UpdateKratosDemoUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateKratosDemoUser not implemented")
}
func (UnimplementedKratosDemoUserServer) DeleteKratosDemoUser(context.Context, *DeleteKratosDemoUserRequest) (*DeleteKratosDemoUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKratosDemoUser not implemented")
}
func (UnimplementedKratosDemoUserServer) GetKratosDemoUser(context.Context, *GetKratosDemoUserRequest) (*GetKratosDemoUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKratosDemoUser not implemented")
}
func (UnimplementedKratosDemoUserServer) Ping(context.Context, *PingReq) (*PingRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedKratosDemoUserServer) mustEmbedUnimplementedKratosDemoUserServer() {}

// UnsafeKratosDemoUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KratosDemoUserServer will
// result in compilation errors.
type UnsafeKratosDemoUserServer interface {
	mustEmbedUnimplementedKratosDemoUserServer()
}

func RegisterKratosDemoUserServer(s grpc.ServiceRegistrar, srv KratosDemoUserServer) {
	s.RegisterService(&KratosDemoUser_ServiceDesc, srv)
}

func _KratosDemoUser_CreateKratosDemoUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateKratosDemoUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KratosDemoUserServer).CreateKratosDemoUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.kratosdemouser.KratosDemoUser/CreateKratosDemoUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KratosDemoUserServer).CreateKratosDemoUser(ctx, req.(*CreateKratosDemoUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KratosDemoUser_UpdateKratosDemoUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateKratosDemoUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KratosDemoUserServer).UpdateKratosDemoUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.kratosdemouser.KratosDemoUser/UpdateKratosDemoUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KratosDemoUserServer).UpdateKratosDemoUser(ctx, req.(*UpdateKratosDemoUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KratosDemoUser_DeleteKratosDemoUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKratosDemoUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KratosDemoUserServer).DeleteKratosDemoUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.kratosdemouser.KratosDemoUser/DeleteKratosDemoUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KratosDemoUserServer).DeleteKratosDemoUser(ctx, req.(*DeleteKratosDemoUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KratosDemoUser_GetKratosDemoUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKratosDemoUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KratosDemoUserServer).GetKratosDemoUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.kratosdemouser.KratosDemoUser/GetKratosDemoUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KratosDemoUserServer).GetKratosDemoUser(ctx, req.(*GetKratosDemoUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KratosDemoUser_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KratosDemoUserServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.kratosdemouser.KratosDemoUser/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KratosDemoUserServer).Ping(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

// KratosDemoUser_ServiceDesc is the grpc.ServiceDesc for KratosDemoUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KratosDemoUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.kratosdemouser.KratosDemoUser",
	HandlerType: (*KratosDemoUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateKratosDemoUser",
			Handler:    _KratosDemoUser_CreateKratosDemoUser_Handler,
		},
		{
			MethodName: "UpdateKratosDemoUser",
			Handler:    _KratosDemoUser_UpdateKratosDemoUser_Handler,
		},
		{
			MethodName: "DeleteKratosDemoUser",
			Handler:    _KratosDemoUser_DeleteKratosDemoUser_Handler,
		},
		{
			MethodName: "GetKratosDemoUser",
			Handler:    _KratosDemoUser_GetKratosDemoUser_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _KratosDemoUser_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/kratosdemouser/kratosDemoUser.proto",
}
