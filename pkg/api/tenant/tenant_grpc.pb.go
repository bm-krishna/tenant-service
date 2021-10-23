// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package tenant

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

// TenantClient is the client API for Tenant service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TenantClient interface {
	API(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*TenantResponse, error)
	Create(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*TenantResponse, error)
	Fetch(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*TenantResponse, error)
	GetByID(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*TenantResponse, error)
	Update(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*TenantResponse, error)
	Delete(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*TenantResponse, error)
}

type tenantClient struct {
	cc grpc.ClientConnInterface
}

func NewTenantClient(cc grpc.ClientConnInterface) TenantClient {
	return &tenantClient{cc}
}

func (c *tenantClient) API(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*TenantResponse, error) {
	out := new(TenantResponse)
	err := c.cc.Invoke(ctx, "/tenant.Tenant/API", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) Create(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*TenantResponse, error) {
	out := new(TenantResponse)
	err := c.cc.Invoke(ctx, "/tenant.Tenant/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) Fetch(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*TenantResponse, error) {
	out := new(TenantResponse)
	err := c.cc.Invoke(ctx, "/tenant.Tenant/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) GetByID(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*TenantResponse, error) {
	out := new(TenantResponse)
	err := c.cc.Invoke(ctx, "/tenant.Tenant/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) Update(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*TenantResponse, error) {
	out := new(TenantResponse)
	err := c.cc.Invoke(ctx, "/tenant.Tenant/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) Delete(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*TenantResponse, error) {
	out := new(TenantResponse)
	err := c.cc.Invoke(ctx, "/tenant.Tenant/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantServer is the server API for Tenant service.
// All implementations must embed UnimplementedTenantServer
// for forward compatibility
type TenantServer interface {
	API(context.Context, *TenantRequest) (*TenantResponse, error)
	Create(context.Context, *TenantRequest) (*TenantResponse, error)
	Fetch(context.Context, *TenantRequest) (*TenantResponse, error)
	GetByID(context.Context, *TenantRequest) (*TenantResponse, error)
	Update(context.Context, *TenantRequest) (*TenantResponse, error)
	Delete(context.Context, *TenantRequest) (*TenantResponse, error)
	mustEmbedUnimplementedTenantServer()
}

// UnimplementedTenantServer must be embedded to have forward compatible implementations.
type UnimplementedTenantServer struct {
}

func (UnimplementedTenantServer) API(context.Context, *TenantRequest) (*TenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method API not implemented")
}
func (UnimplementedTenantServer) Create(context.Context, *TenantRequest) (*TenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTenantServer) Fetch(context.Context, *TenantRequest) (*TenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedTenantServer) GetByID(context.Context, *TenantRequest) (*TenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedTenantServer) Update(context.Context, *TenantRequest) (*TenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTenantServer) Delete(context.Context, *TenantRequest) (*TenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTenantServer) mustEmbedUnimplementedTenantServer() {}

// UnsafeTenantServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TenantServer will
// result in compilation errors.
type UnsafeTenantServer interface {
	mustEmbedUnimplementedTenantServer()
}

func RegisterTenantServer(s grpc.ServiceRegistrar, srv TenantServer) {
	s.RegisterService(&Tenant_ServiceDesc, srv)
}

func _Tenant_API_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).API(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tenant.Tenant/API",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).API(ctx, req.(*TenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tenant.Tenant/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).Create(ctx, req.(*TenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tenant.Tenant/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).Fetch(ctx, req.(*TenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tenant.Tenant/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).GetByID(ctx, req.(*TenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tenant.Tenant/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).Update(ctx, req.(*TenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tenant.Tenant/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).Delete(ctx, req.(*TenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Tenant_ServiceDesc is the grpc.ServiceDesc for Tenant service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tenant_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tenant.Tenant",
	HandlerType: (*TenantServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "API",
			Handler:    _Tenant_API_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Tenant_Create_Handler,
		},
		{
			MethodName: "Fetch",
			Handler:    _Tenant_Fetch_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _Tenant_GetByID_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Tenant_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Tenant_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/tenant.proto",
}
