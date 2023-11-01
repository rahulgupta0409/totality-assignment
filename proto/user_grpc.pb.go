// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: user.proto

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

const (
	UserService_GetUserById_FullMethodName      = "/UserService.UserService/GetUserById"
	UserService_GetUserListByIds_FullMethodName = "/UserService.UserService/GetUserListByIds"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUserById(ctx context.Context, opts ...grpc.CallOption) (UserService_GetUserByIdClient, error)
	GetUserListByIds(ctx context.Context, opts ...grpc.CallOption) (UserService_GetUserListByIdsClient, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserById(ctx context.Context, opts ...grpc.CallOption) (UserService_GetUserByIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[0], UserService_GetUserById_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetUserByIdClient{stream}
	return x, nil
}

type UserService_GetUserByIdClient interface {
	Send(*UserRequest) error
	Recv() (*UserResponse, error)
	grpc.ClientStream
}

type userServiceGetUserByIdClient struct {
	grpc.ClientStream
}

func (x *userServiceGetUserByIdClient) Send(m *UserRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceGetUserByIdClient) Recv() (*UserResponse, error) {
	m := new(UserResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) GetUserListByIds(ctx context.Context, opts ...grpc.CallOption) (UserService_GetUserListByIdsClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[1], UserService_GetUserListByIds_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetUserListByIdsClient{stream}
	return x, nil
}

type UserService_GetUserListByIdsClient interface {
	Send(*UserList) error
	Recv() (*UserResponseList, error)
	grpc.ClientStream
}

type userServiceGetUserListByIdsClient struct {
	grpc.ClientStream
}

func (x *userServiceGetUserListByIdsClient) Send(m *UserList) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceGetUserListByIdsClient) Recv() (*UserResponseList, error) {
	m := new(UserResponseList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetUserById(UserService_GetUserByIdServer) error
	GetUserListByIds(UserService_GetUserListByIdsServer) error
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetUserById(UserService_GetUserByIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUserServiceServer) GetUserListByIds(UserService_GetUserListByIdsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUserListByIds not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetUserById_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServiceServer).GetUserById(&userServiceGetUserByIdServer{stream})
}

type UserService_GetUserByIdServer interface {
	Send(*UserResponse) error
	Recv() (*UserRequest, error)
	grpc.ServerStream
}

type userServiceGetUserByIdServer struct {
	grpc.ServerStream
}

func (x *userServiceGetUserByIdServer) Send(m *UserResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceGetUserByIdServer) Recv() (*UserRequest, error) {
	m := new(UserRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UserService_GetUserListByIds_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServiceServer).GetUserListByIds(&userServiceGetUserListByIdsServer{stream})
}

type UserService_GetUserListByIdsServer interface {
	Send(*UserResponseList) error
	Recv() (*UserList, error)
	grpc.ServerStream
}

type userServiceGetUserListByIdsServer struct {
	grpc.ServerStream
}

func (x *userServiceGetUserListByIdsServer) Send(m *UserResponseList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceGetUserListByIdsServer) Recv() (*UserList, error) {
	m := new(UserList)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserService.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUserById",
			Handler:       _UserService_GetUserById_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetUserListByIds",
			Handler:       _UserService_GetUserListByIds_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "user.proto",
}
