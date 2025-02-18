// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: user.proto

package pb

import (
	context "context"
	transport "github.com/erda-project/erda-infra/pkg/transport"
	grpc1 "github.com/erda-project/erda-infra/pkg/transport/grpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// get user
	// +publish path:"/user/{id}"
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	// update user
	// +private
	UpdateUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
}

type userServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewUserServiceClient(cc grpc1.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/erda.example.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/erda.example.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations should embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// get user
	// +publish path:"/user/{id}"
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	// update user
	// +private
	UpdateUser(context.Context, *GetUserRequest) (*UpdateUserResponse, error)
}

// UnimplementedUserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserServiceServer) UpdateUser(context.Context, *GetUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

func RegisterUserServiceServer(s grpc1.ServiceRegistrar, srv UserServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_UserService_serviceDesc(srv, opts...), srv)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.example.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "user.proto",
}

func _get_UserService_serviceDesc(srv UserServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_UserService_GetUser_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetUser(ctx, req.(*GetUserRequest))
	}
	var _UserService_GetUser_info transport.ServiceInfo
	if h.Interceptor != nil {
		_UserService_GetUser_info = transport.NewServiceInfo("erda.example.UserService", "GetUser", srv)
		_UserService_GetUser_Handler = h.Interceptor(_UserService_GetUser_Handler)
	}

	_UserService_UpdateUser_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.UpdateUser(ctx, req.(*GetUserRequest))
	}
	var _UserService_UpdateUser_info transport.ServiceInfo
	if h.Interceptor != nil {
		_UserService_UpdateUser_info = transport.NewServiceInfo("erda.example.UserService", "UpdateUser", srv)
		_UserService_UpdateUser_Handler = h.Interceptor(_UserService_UpdateUser_Handler)
	}

	var serviceDesc = _UserService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetUserRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(UserServiceServer).GetUser(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _UserService_GetUser_info)
				}
				if interceptor == nil {
					return _UserService_GetUser_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.example.UserService/GetUser",
				}
				return interceptor(ctx, in, info, _UserService_GetUser_Handler)
			},
		},
		{
			MethodName: "UpdateUser",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetUserRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(UserServiceServer).UpdateUser(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _UserService_UpdateUser_info)
				}
				if interceptor == nil {
					return _UserService_UpdateUser_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.example.UserService/UpdateUser",
				}
				return interceptor(ctx, in, info, _UserService_UpdateUser_Handler)
			},
		},
	}
	return &serviceDesc
}
