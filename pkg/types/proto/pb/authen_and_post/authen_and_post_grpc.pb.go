// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: authen_and_post.proto

package authen_and_post

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
	AuthenticateAndPost_CheckUserAuthentication_FullMethodName = "/authen_and_post.AuthenticateAndPost/CheckUserAuthentication"
	AuthenticateAndPost_CreateUser_FullMethodName              = "/authen_and_post.AuthenticateAndPost/CreateUser"
	AuthenticateAndPost_EditUser_FullMethodName                = "/authen_and_post.AuthenticateAndPost/EditUser"
	AuthenticateAndPost_GetUserFollower_FullMethodName         = "/authen_and_post.AuthenticateAndPost/GetUserFollower"
	AuthenticateAndPost_GetPostDetail_FullMethodName           = "/authen_and_post.AuthenticateAndPost/GetPostDetail"
	AuthenticateAndPost_UnFollowUser_FullMethodName            = "/authen_and_post.AuthenticateAndPost/UnFollowUser"
	AuthenticateAndPost_FollowUser_FullMethodName              = "/authen_and_post.AuthenticateAndPost/FollowUser"
	AuthenticateAndPost_CreatePost_FullMethodName              = "/authen_and_post.AuthenticateAndPost/CreatePost"
	AuthenticateAndPost_GetPostFriend_FullMethodName           = "/authen_and_post.AuthenticateAndPost/GetPostFriend"
	AuthenticateAndPost_DeletePost_FullMethodName              = "/authen_and_post.AuthenticateAndPost/DeletePost"
	AuthenticateAndPost_UpdatePost_FullMethodName              = "/authen_and_post.AuthenticateAndPost/UpdatePost"
	AuthenticateAndPost_CommentPost_FullMethodName             = "/authen_and_post.AuthenticateAndPost/CommentPost"
	AuthenticateAndPost_LikePost_FullMethodName                = "/authen_and_post.AuthenticateAndPost/LikePost"
)

// AuthenticateAndPostClient is the client API for AuthenticateAndPost service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticateAndPostClient interface {
	CheckUserAuthentication(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserResult, error)
	CreateUser(ctx context.Context, in *UserDetailInfo, opts ...grpc.CallOption) (*UserResult, error)
	EditUser(ctx context.Context, in *UserDetailInfo, opts ...grpc.CallOption) (*UserResult, error)
	GetUserFollower(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserFollower, error)
	GetPostDetail(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*Post, error)
	UnFollowUser(ctx context.Context, in *UserFollowRequest, opts ...grpc.CallOption) (*BoolResp, error)
	FollowUser(ctx context.Context, in *UserFollowRequest, opts ...grpc.CallOption) (*BoolResp, error)
	CreatePost(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*Post, error)
	GetPostFriend(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*PostFriend, error)
	DeletePost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*TextResponse, error)
	UpdatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Post, error)
	CommentPost(ctx context.Context, in *CommentPostRequest, opts ...grpc.CallOption) (*BoolResp, error)
	LikePost(ctx context.Context, in *LikePostRequest, opts ...grpc.CallOption) (*BoolResp, error)
}

type authenticateAndPostClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticateAndPostClient(cc grpc.ClientConnInterface) AuthenticateAndPostClient {
	return &authenticateAndPostClient{cc}
}

func (c *authenticateAndPostClient) CheckUserAuthentication(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserResult, error) {
	out := new(UserResult)
	err := c.cc.Invoke(ctx, AuthenticateAndPost_CheckUserAuthentication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateAndPostClient) CreateUser(ctx context.Context, in *UserDetailInfo, opts ...grpc.CallOption) (*UserResult, error) {
	out := new(UserResult)
	err := c.cc.Invoke(ctx, AuthenticateAndPost_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateAndPostClient) EditUser(ctx context.Context, in *UserDetailInfo, opts ...grpc.CallOption) (*UserResult, error) {
	out := new(UserResult)
	err := c.cc.Invoke(ctx, AuthenticateAndPost_EditUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateAndPostClient) GetUserFollower(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserFollower, error) {
	out := new(UserFollower)
	err := c.cc.Invoke(ctx, AuthenticateAndPost_GetUserFollower_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateAndPostClient) GetPostDetail(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, AuthenticateAndPost_GetPostDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateAndPostClient) UnFollowUser(ctx context.Context, in *UserFollowRequest, opts ...grpc.CallOption) (*BoolResp, error) {
	out := new(BoolResp)
	err := c.cc.Invoke(ctx, AuthenticateAndPost_UnFollowUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateAndPostClient) FollowUser(ctx context.Context, in *UserFollowRequest, opts ...grpc.CallOption) (*BoolResp, error) {
	out := new(BoolResp)
	err := c.cc.Invoke(ctx, AuthenticateAndPost_FollowUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateAndPostClient) CreatePost(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, AuthenticateAndPost_CreatePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateAndPostClient) GetPostFriend(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*PostFriend, error) {
	out := new(PostFriend)
	err := c.cc.Invoke(ctx, AuthenticateAndPost_GetPostFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateAndPostClient) DeletePost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*TextResponse, error) {
	out := new(TextResponse)
	err := c.cc.Invoke(ctx, AuthenticateAndPost_DeletePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateAndPostClient) UpdatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, AuthenticateAndPost_UpdatePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateAndPostClient) CommentPost(ctx context.Context, in *CommentPostRequest, opts ...grpc.CallOption) (*BoolResp, error) {
	out := new(BoolResp)
	err := c.cc.Invoke(ctx, AuthenticateAndPost_CommentPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateAndPostClient) LikePost(ctx context.Context, in *LikePostRequest, opts ...grpc.CallOption) (*BoolResp, error) {
	out := new(BoolResp)
	err := c.cc.Invoke(ctx, AuthenticateAndPost_LikePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticateAndPostServer is the server API for AuthenticateAndPost service.
// All implementations must embed UnimplementedAuthenticateAndPostServer
// for forward compatibility
type AuthenticateAndPostServer interface {
	CheckUserAuthentication(context.Context, *UserInfo) (*UserResult, error)
	CreateUser(context.Context, *UserDetailInfo) (*UserResult, error)
	EditUser(context.Context, *UserDetailInfo) (*UserResult, error)
	GetUserFollower(context.Context, *UserInfo) (*UserFollower, error)
	GetPostDetail(context.Context, *GetPostRequest) (*Post, error)
	UnFollowUser(context.Context, *UserFollowRequest) (*BoolResp, error)
	FollowUser(context.Context, *UserFollowRequest) (*BoolResp, error)
	CreatePost(context.Context, *PostRequest) (*Post, error)
	GetPostFriend(context.Context, *UserInfo) (*PostFriend, error)
	DeletePost(context.Context, *GetPostRequest) (*TextResponse, error)
	UpdatePost(context.Context, *Post) (*Post, error)
	CommentPost(context.Context, *CommentPostRequest) (*BoolResp, error)
	LikePost(context.Context, *LikePostRequest) (*BoolResp, error)
	mustEmbedUnimplementedAuthenticateAndPostServer()
}

// UnimplementedAuthenticateAndPostServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticateAndPostServer struct {
}

func (UnimplementedAuthenticateAndPostServer) CheckUserAuthentication(context.Context, *UserInfo) (*UserResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserAuthentication not implemented")
}
func (UnimplementedAuthenticateAndPostServer) CreateUser(context.Context, *UserDetailInfo) (*UserResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedAuthenticateAndPostServer) EditUser(context.Context, *UserDetailInfo) (*UserResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUser not implemented")
}
func (UnimplementedAuthenticateAndPostServer) GetUserFollower(context.Context, *UserInfo) (*UserFollower, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFollower not implemented")
}
func (UnimplementedAuthenticateAndPostServer) GetPostDetail(context.Context, *GetPostRequest) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostDetail not implemented")
}
func (UnimplementedAuthenticateAndPostServer) UnFollowUser(context.Context, *UserFollowRequest) (*BoolResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFollowUser not implemented")
}
func (UnimplementedAuthenticateAndPostServer) FollowUser(context.Context, *UserFollowRequest) (*BoolResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowUser not implemented")
}
func (UnimplementedAuthenticateAndPostServer) CreatePost(context.Context, *PostRequest) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedAuthenticateAndPostServer) GetPostFriend(context.Context, *UserInfo) (*PostFriend, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostFriend not implemented")
}
func (UnimplementedAuthenticateAndPostServer) DeletePost(context.Context, *GetPostRequest) (*TextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedAuthenticateAndPostServer) UpdatePost(context.Context, *Post) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (UnimplementedAuthenticateAndPostServer) CommentPost(context.Context, *CommentPostRequest) (*BoolResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentPost not implemented")
}
func (UnimplementedAuthenticateAndPostServer) LikePost(context.Context, *LikePostRequest) (*BoolResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikePost not implemented")
}
func (UnimplementedAuthenticateAndPostServer) mustEmbedUnimplementedAuthenticateAndPostServer() {}

// UnsafeAuthenticateAndPostServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticateAndPostServer will
// result in compilation errors.
type UnsafeAuthenticateAndPostServer interface {
	mustEmbedUnimplementedAuthenticateAndPostServer()
}

func RegisterAuthenticateAndPostServer(s grpc.ServiceRegistrar, srv AuthenticateAndPostServer) {
	s.RegisterService(&AuthenticateAndPost_ServiceDesc, srv)
}

func _AuthenticateAndPost_CheckUserAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateAndPostServer).CheckUserAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticateAndPost_CheckUserAuthentication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateAndPostServer).CheckUserAuthentication(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticateAndPost_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDetailInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateAndPostServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticateAndPost_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateAndPostServer).CreateUser(ctx, req.(*UserDetailInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticateAndPost_EditUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDetailInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateAndPostServer).EditUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticateAndPost_EditUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateAndPostServer).EditUser(ctx, req.(*UserDetailInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticateAndPost_GetUserFollower_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateAndPostServer).GetUserFollower(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticateAndPost_GetUserFollower_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateAndPostServer).GetUserFollower(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticateAndPost_GetPostDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateAndPostServer).GetPostDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticateAndPost_GetPostDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateAndPostServer).GetPostDetail(ctx, req.(*GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticateAndPost_UnFollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateAndPostServer).UnFollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticateAndPost_UnFollowUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateAndPostServer).UnFollowUser(ctx, req.(*UserFollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticateAndPost_FollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateAndPostServer).FollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticateAndPost_FollowUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateAndPostServer).FollowUser(ctx, req.(*UserFollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticateAndPost_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateAndPostServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticateAndPost_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateAndPostServer).CreatePost(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticateAndPost_GetPostFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateAndPostServer).GetPostFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticateAndPost_GetPostFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateAndPostServer).GetPostFriend(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticateAndPost_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateAndPostServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticateAndPost_DeletePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateAndPostServer).DeletePost(ctx, req.(*GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticateAndPost_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateAndPostServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticateAndPost_UpdatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateAndPostServer).UpdatePost(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticateAndPost_CommentPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateAndPostServer).CommentPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticateAndPost_CommentPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateAndPostServer).CommentPost(ctx, req.(*CommentPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticateAndPost_LikePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateAndPostServer).LikePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticateAndPost_LikePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateAndPostServer).LikePost(ctx, req.(*LikePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenticateAndPost_ServiceDesc is the grpc.ServiceDesc for AuthenticateAndPost service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenticateAndPost_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authen_and_post.AuthenticateAndPost",
	HandlerType: (*AuthenticateAndPostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckUserAuthentication",
			Handler:    _AuthenticateAndPost_CheckUserAuthentication_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _AuthenticateAndPost_CreateUser_Handler,
		},
		{
			MethodName: "EditUser",
			Handler:    _AuthenticateAndPost_EditUser_Handler,
		},
		{
			MethodName: "GetUserFollower",
			Handler:    _AuthenticateAndPost_GetUserFollower_Handler,
		},
		{
			MethodName: "GetPostDetail",
			Handler:    _AuthenticateAndPost_GetPostDetail_Handler,
		},
		{
			MethodName: "UnFollowUser",
			Handler:    _AuthenticateAndPost_UnFollowUser_Handler,
		},
		{
			MethodName: "FollowUser",
			Handler:    _AuthenticateAndPost_FollowUser_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _AuthenticateAndPost_CreatePost_Handler,
		},
		{
			MethodName: "GetPostFriend",
			Handler:    _AuthenticateAndPost_GetPostFriend_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _AuthenticateAndPost_DeletePost_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _AuthenticateAndPost_UpdatePost_Handler,
		},
		{
			MethodName: "CommentPost",
			Handler:    _AuthenticateAndPost_CommentPost_Handler,
		},
		{
			MethodName: "LikePost",
			Handler:    _AuthenticateAndPost_LikePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authen_and_post.proto",
}
