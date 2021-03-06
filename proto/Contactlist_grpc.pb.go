// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// ContactRequestClient is the client API for ContactRequest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContactRequestClient interface {
	Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	Get(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	Delete(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
}

type contactRequestClient struct {
	cc grpc.ClientConnInterface
}

func NewContactRequestClient(cc grpc.ClientConnInterface) ContactRequestClient {
	return &contactRequestClient{cc}
}

func (c *contactRequestClient) Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/contactList.ContactRequest/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactRequestClient) Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/contactList.ContactRequest/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactRequestClient) Get(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/contactList.ContactRequest/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactRequestClient) Delete(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/contactList.ContactRequest/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactRequestClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/contactList.ContactRequest/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactRequestServer is the server API for ContactRequest service.
// All implementations must embed UnimplementedContactRequestServer
// for forward compatibility
type ContactRequestServer interface {
	Create(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	Update(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	Get(context.Context, *GetUserRequest) (*GetUserResponse, error)
	Delete(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	mustEmbedUnimplementedContactRequestServer()
}

// UnimplementedContactRequestServer must be embedded to have forward compatible implementations.
type UnimplementedContactRequestServer struct {
}

func (UnimplementedContactRequestServer) Create(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedContactRequestServer) Update(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedContactRequestServer) Get(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedContactRequestServer) Delete(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedContactRequestServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedContactRequestServer) mustEmbedUnimplementedContactRequestServer() {}

// UnsafeContactRequestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContactRequestServer will
// result in compilation errors.
type UnsafeContactRequestServer interface {
	mustEmbedUnimplementedContactRequestServer()
}

func RegisterContactRequestServer(s grpc.ServiceRegistrar, srv ContactRequestServer) {
	s.RegisterService(&ContactRequest_ServiceDesc, srv)
}

func _ContactRequest_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactRequestServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contactList.ContactRequest/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactRequestServer).Create(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactRequest_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactRequestServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contactList.ContactRequest/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactRequestServer).Update(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactRequest_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactRequestServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contactList.ContactRequest/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactRequestServer).Get(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactRequest_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactRequestServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contactList.ContactRequest/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactRequestServer).Delete(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactRequest_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactRequestServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contactList.ContactRequest/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactRequestServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContactRequest_ServiceDesc is the grpc.ServiceDesc for ContactRequest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContactRequest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "contactList.ContactRequest",
	HandlerType: (*ContactRequestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ContactRequest_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ContactRequest_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ContactRequest_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ContactRequest_Delete_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ContactRequest_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/Contactlist.proto",
}
