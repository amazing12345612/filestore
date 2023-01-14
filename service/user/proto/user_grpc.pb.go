// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service/user/proto/user.proto

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// 用户注册
	Signup(ctx context.Context, in *ReqSignup, opts ...grpc.CallOption) (*RespSignup, error)
	// 用户登录
	Signin(ctx context.Context, in *ReqSignin, opts ...grpc.CallOption) (*RespSignin, error)
	// 获取用户信息
	UserInfo(ctx context.Context, in *ReqUserInfo, opts ...grpc.CallOption) (*RespUserInfo, error)
	// 获取用户文件
	UserFiles(ctx context.Context, in *ReqUserFile, opts ...grpc.CallOption) (*RespUserFile, error)
	// 获取用户文件
	UserFileRename(ctx context.Context, in *ReqUserFileRename, opts ...grpc.CallOption) (*RespUserFileRename, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Signup(ctx context.Context, in *ReqSignup, opts ...grpc.CallOption) (*RespSignup, error) {
	out := new(RespSignup)
	err := c.cc.Invoke(ctx, "/user.UserService/Signup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Signin(ctx context.Context, in *ReqSignin, opts ...grpc.CallOption) (*RespSignin, error) {
	out := new(RespSignin)
	err := c.cc.Invoke(ctx, "/user.UserService/Signin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserInfo(ctx context.Context, in *ReqUserInfo, opts ...grpc.CallOption) (*RespUserInfo, error) {
	out := new(RespUserInfo)
	err := c.cc.Invoke(ctx, "/user.UserService/UserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserFiles(ctx context.Context, in *ReqUserFile, opts ...grpc.CallOption) (*RespUserFile, error) {
	out := new(RespUserFile)
	err := c.cc.Invoke(ctx, "/user.UserService/UserFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserFileRename(ctx context.Context, in *ReqUserFileRename, opts ...grpc.CallOption) (*RespUserFileRename, error) {
	out := new(RespUserFileRename)
	err := c.cc.Invoke(ctx, "/user.UserService/UserFileRename", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// 用户注册
	Signup(context.Context, *ReqSignup) (*RespSignup, error)
	// 用户登录
	Signin(context.Context, *ReqSignin) (*RespSignin, error)
	// 获取用户信息
	UserInfo(context.Context, *ReqUserInfo) (*RespUserInfo, error)
	// 获取用户文件
	UserFiles(context.Context, *ReqUserFile) (*RespUserFile, error)
	// 获取用户文件
	UserFileRename(context.Context, *ReqUserFileRename) (*RespUserFileRename, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Signup(context.Context, *ReqSignup) (*RespSignup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedUserServiceServer) Signin(context.Context, *ReqSignin) (*RespSignin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signin not implemented")
}
func (UnimplementedUserServiceServer) UserInfo(context.Context, *ReqUserInfo) (*RespUserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedUserServiceServer) UserFiles(context.Context, *ReqUserFile) (*RespUserFile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserFiles not implemented")
}
func (UnimplementedUserServiceServer) UserFileRename(context.Context, *ReqUserFileRename) (*RespUserFileRename, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserFileRename not implemented")
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

func _UserService_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqSignup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Signup(ctx, req.(*ReqSignup))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Signin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqSignin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Signin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Signin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Signin(ctx, req.(*ReqSignin))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqUserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserInfo(ctx, req.(*ReqUserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqUserFile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UserFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserFiles(ctx, req.(*ReqUserFile))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserFileRename_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqUserFileRename)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserFileRename(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UserFileRename",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserFileRename(ctx, req.(*ReqUserFileRename))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _UserService_Signup_Handler,
		},
		{
			MethodName: "Signin",
			Handler:    _UserService_Signin_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _UserService_UserInfo_Handler,
		},
		{
			MethodName: "UserFiles",
			Handler:    _UserService_UserFiles_Handler,
		},
		{
			MethodName: "UserFileRename",
			Handler:    _UserService_UserFileRename_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/user/proto/user.proto",
}
