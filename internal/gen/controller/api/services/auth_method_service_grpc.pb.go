// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

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

// AuthMethodServiceClient is the client API for AuthMethodService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthMethodServiceClient interface {
	// GetAuthMethod returns a stored Auth Method if present.  The provided request
	// must include the Auth Method id. If missing, malformed or referencing a
	// non existing resource an error is returned.
	GetAuthMethod(ctx context.Context, in *GetAuthMethodRequest, opts ...grpc.CallOption) (*GetAuthMethodResponse, error)
	// ListAuthMethods returns a list of stored Auth Methods which are in the
	// provided scope. The request must include the scope id and if missing,
	// malformed, or referencing a non existing scope, an error is returned.
	ListAuthMethods(ctx context.Context, in *ListAuthMethodsRequest, opts ...grpc.CallOption) (*ListAuthMethodsResponse, error)
	// CreateAuthMethod creates and stores an Auth Method in boundary.  The
	// provided request must include the scope in which the Auth Method will be
	// created. If the scope id is missing, malformed or referencing a
	// non existing resource an error is returned.  If a name is provided that is
	// in use in another Auth Method in the same scope, an error is returned.
	CreateAuthMethod(ctx context.Context, in *CreateAuthMethodRequest, opts ...grpc.CallOption) (*CreateAuthMethodResponse, error)
	// UpdateAuthMethod updates an existing Auth Method in boundary.  The provided
	// Auth Method must not have any read only fields set.  The update mask must be
	// included in the request and contain at least 1 mutable field.  To unset
	// a field's value, include the field in the update mask and don't set it
	// in the provided user. An error is returned if the Auth Method id is missing
	// or reference a non existing resource.  An error is also returned if the
	// request attempts to update the name to one that is already in use by
	// another Auth Method in the parent scope.
	UpdateAuthMethod(ctx context.Context, in *UpdateAuthMethodRequest, opts ...grpc.CallOption) (*UpdateAuthMethodResponse, error)
	// DeleteAuthMethod removes an Auth Method from Boundary. If the Auth Method id
	// is malformed or not provided an error is returned.
	DeleteAuthMethod(ctx context.Context, in *DeleteAuthMethodRequest, opts ...grpc.CallOption) (*DeleteAuthMethodResponse, error)
	// ChangeState changes the state of an Auth Method from Boundary.
	ChangeState(ctx context.Context, in *ChangeStateRequest, opts ...grpc.CallOption) (*ChangeStateResponse, error)
	// Deprecated: Do not use.
	// Authenticate validates credentials provided and returns an Auth Token.
	// Deprecated: use AuthenticateLogin instead.
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	// AuthenticateLogin validates credentials provided for a single-step login
	// action and returns an auth token.
	AuthenticateLogin(ctx context.Context, in *AuthenticateLoginRequest, opts ...grpc.CallOption) (*AuthenticateLoginResponse, error)
}

type authMethodServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthMethodServiceClient(cc grpc.ClientConnInterface) AuthMethodServiceClient {
	return &authMethodServiceClient{cc}
}

func (c *authMethodServiceClient) GetAuthMethod(ctx context.Context, in *GetAuthMethodRequest, opts ...grpc.CallOption) (*GetAuthMethodResponse, error) {
	out := new(GetAuthMethodResponse)
	err := c.cc.Invoke(ctx, "/controller.api.services.v1.AuthMethodService/GetAuthMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authMethodServiceClient) ListAuthMethods(ctx context.Context, in *ListAuthMethodsRequest, opts ...grpc.CallOption) (*ListAuthMethodsResponse, error) {
	out := new(ListAuthMethodsResponse)
	err := c.cc.Invoke(ctx, "/controller.api.services.v1.AuthMethodService/ListAuthMethods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authMethodServiceClient) CreateAuthMethod(ctx context.Context, in *CreateAuthMethodRequest, opts ...grpc.CallOption) (*CreateAuthMethodResponse, error) {
	out := new(CreateAuthMethodResponse)
	err := c.cc.Invoke(ctx, "/controller.api.services.v1.AuthMethodService/CreateAuthMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authMethodServiceClient) UpdateAuthMethod(ctx context.Context, in *UpdateAuthMethodRequest, opts ...grpc.CallOption) (*UpdateAuthMethodResponse, error) {
	out := new(UpdateAuthMethodResponse)
	err := c.cc.Invoke(ctx, "/controller.api.services.v1.AuthMethodService/UpdateAuthMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authMethodServiceClient) DeleteAuthMethod(ctx context.Context, in *DeleteAuthMethodRequest, opts ...grpc.CallOption) (*DeleteAuthMethodResponse, error) {
	out := new(DeleteAuthMethodResponse)
	err := c.cc.Invoke(ctx, "/controller.api.services.v1.AuthMethodService/DeleteAuthMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authMethodServiceClient) ChangeState(ctx context.Context, in *ChangeStateRequest, opts ...grpc.CallOption) (*ChangeStateResponse, error) {
	out := new(ChangeStateResponse)
	err := c.cc.Invoke(ctx, "/controller.api.services.v1.AuthMethodService/ChangeState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *authMethodServiceClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, "/controller.api.services.v1.AuthMethodService/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authMethodServiceClient) AuthenticateLogin(ctx context.Context, in *AuthenticateLoginRequest, opts ...grpc.CallOption) (*AuthenticateLoginResponse, error) {
	out := new(AuthenticateLoginResponse)
	err := c.cc.Invoke(ctx, "/controller.api.services.v1.AuthMethodService/AuthenticateLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthMethodServiceServer is the server API for AuthMethodService service.
// All implementations must embed UnimplementedAuthMethodServiceServer
// for forward compatibility
type AuthMethodServiceServer interface {
	// GetAuthMethod returns a stored Auth Method if present.  The provided request
	// must include the Auth Method id. If missing, malformed or referencing a
	// non existing resource an error is returned.
	GetAuthMethod(context.Context, *GetAuthMethodRequest) (*GetAuthMethodResponse, error)
	// ListAuthMethods returns a list of stored Auth Methods which are in the
	// provided scope. The request must include the scope id and if missing,
	// malformed, or referencing a non existing scope, an error is returned.
	ListAuthMethods(context.Context, *ListAuthMethodsRequest) (*ListAuthMethodsResponse, error)
	// CreateAuthMethod creates and stores an Auth Method in boundary.  The
	// provided request must include the scope in which the Auth Method will be
	// created. If the scope id is missing, malformed or referencing a
	// non existing resource an error is returned.  If a name is provided that is
	// in use in another Auth Method in the same scope, an error is returned.
	CreateAuthMethod(context.Context, *CreateAuthMethodRequest) (*CreateAuthMethodResponse, error)
	// UpdateAuthMethod updates an existing Auth Method in boundary.  The provided
	// Auth Method must not have any read only fields set.  The update mask must be
	// included in the request and contain at least 1 mutable field.  To unset
	// a field's value, include the field in the update mask and don't set it
	// in the provided user. An error is returned if the Auth Method id is missing
	// or reference a non existing resource.  An error is also returned if the
	// request attempts to update the name to one that is already in use by
	// another Auth Method in the parent scope.
	UpdateAuthMethod(context.Context, *UpdateAuthMethodRequest) (*UpdateAuthMethodResponse, error)
	// DeleteAuthMethod removes an Auth Method from Boundary. If the Auth Method id
	// is malformed or not provided an error is returned.
	DeleteAuthMethod(context.Context, *DeleteAuthMethodRequest) (*DeleteAuthMethodResponse, error)
	// ChangeState changes the state of an Auth Method from Boundary.
	ChangeState(context.Context, *ChangeStateRequest) (*ChangeStateResponse, error)
	// Deprecated: Do not use.
	// Authenticate validates credentials provided and returns an Auth Token.
	// Deprecated: use AuthenticateLogin instead.
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	// AuthenticateLogin validates credentials provided for a single-step login
	// action and returns an auth token.
	AuthenticateLogin(context.Context, *AuthenticateLoginRequest) (*AuthenticateLoginResponse, error)
	mustEmbedUnimplementedAuthMethodServiceServer()
}

// UnimplementedAuthMethodServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthMethodServiceServer struct {
}

func (UnimplementedAuthMethodServiceServer) GetAuthMethod(context.Context, *GetAuthMethodRequest) (*GetAuthMethodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthMethod not implemented")
}
func (UnimplementedAuthMethodServiceServer) ListAuthMethods(context.Context, *ListAuthMethodsRequest) (*ListAuthMethodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAuthMethods not implemented")
}
func (UnimplementedAuthMethodServiceServer) CreateAuthMethod(context.Context, *CreateAuthMethodRequest) (*CreateAuthMethodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuthMethod not implemented")
}
func (UnimplementedAuthMethodServiceServer) UpdateAuthMethod(context.Context, *UpdateAuthMethodRequest) (*UpdateAuthMethodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuthMethod not implemented")
}
func (UnimplementedAuthMethodServiceServer) DeleteAuthMethod(context.Context, *DeleteAuthMethodRequest) (*DeleteAuthMethodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAuthMethod not implemented")
}
func (UnimplementedAuthMethodServiceServer) ChangeState(context.Context, *ChangeStateRequest) (*ChangeStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeState not implemented")
}
func (UnimplementedAuthMethodServiceServer) Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedAuthMethodServiceServer) AuthenticateLogin(context.Context, *AuthenticateLoginRequest) (*AuthenticateLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticateLogin not implemented")
}
func (UnimplementedAuthMethodServiceServer) mustEmbedUnimplementedAuthMethodServiceServer() {}

// UnsafeAuthMethodServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthMethodServiceServer will
// result in compilation errors.
type UnsafeAuthMethodServiceServer interface {
	mustEmbedUnimplementedAuthMethodServiceServer()
}

func RegisterAuthMethodServiceServer(s grpc.ServiceRegistrar, srv AuthMethodServiceServer) {
	s.RegisterService(&AuthMethodService_ServiceDesc, srv)
}

func _AuthMethodService_GetAuthMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthMethodServiceServer).GetAuthMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.api.services.v1.AuthMethodService/GetAuthMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthMethodServiceServer).GetAuthMethod(ctx, req.(*GetAuthMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthMethodService_ListAuthMethods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAuthMethodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthMethodServiceServer).ListAuthMethods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.api.services.v1.AuthMethodService/ListAuthMethods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthMethodServiceServer).ListAuthMethods(ctx, req.(*ListAuthMethodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthMethodService_CreateAuthMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAuthMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthMethodServiceServer).CreateAuthMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.api.services.v1.AuthMethodService/CreateAuthMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthMethodServiceServer).CreateAuthMethod(ctx, req.(*CreateAuthMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthMethodService_UpdateAuthMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAuthMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthMethodServiceServer).UpdateAuthMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.api.services.v1.AuthMethodService/UpdateAuthMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthMethodServiceServer).UpdateAuthMethod(ctx, req.(*UpdateAuthMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthMethodService_DeleteAuthMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAuthMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthMethodServiceServer).DeleteAuthMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.api.services.v1.AuthMethodService/DeleteAuthMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthMethodServiceServer).DeleteAuthMethod(ctx, req.(*DeleteAuthMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthMethodService_ChangeState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthMethodServiceServer).ChangeState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.api.services.v1.AuthMethodService/ChangeState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthMethodServiceServer).ChangeState(ctx, req.(*ChangeStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthMethodService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthMethodServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.api.services.v1.AuthMethodService/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthMethodServiceServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthMethodService_AuthenticateLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthMethodServiceServer).AuthenticateLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.api.services.v1.AuthMethodService/AuthenticateLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthMethodServiceServer).AuthenticateLogin(ctx, req.(*AuthenticateLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthMethodService_ServiceDesc is the grpc.ServiceDesc for AuthMethodService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthMethodService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "controller.api.services.v1.AuthMethodService",
	HandlerType: (*AuthMethodServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuthMethod",
			Handler:    _AuthMethodService_GetAuthMethod_Handler,
		},
		{
			MethodName: "ListAuthMethods",
			Handler:    _AuthMethodService_ListAuthMethods_Handler,
		},
		{
			MethodName: "CreateAuthMethod",
			Handler:    _AuthMethodService_CreateAuthMethod_Handler,
		},
		{
			MethodName: "UpdateAuthMethod",
			Handler:    _AuthMethodService_UpdateAuthMethod_Handler,
		},
		{
			MethodName: "DeleteAuthMethod",
			Handler:    _AuthMethodService_DeleteAuthMethod_Handler,
		},
		{
			MethodName: "ChangeState",
			Handler:    _AuthMethodService_ChangeState_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _AuthMethodService_Authenticate_Handler,
		},
		{
			MethodName: "AuthenticateLogin",
			Handler:    _AuthMethodService_AuthenticateLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controller/api/services/v1/auth_method_service.proto",
}
