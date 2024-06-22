// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: services/model_base_service.proto

package grpc

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

// ModelBaseServiceClient is the client API for ModelBaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelBaseServiceClient interface {
	GetById(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetList(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	DeleteById(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	DeleteList(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	ForceDeleteList(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	UpdateById(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Update(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	UpdateDoc(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Insert(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Count(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type modelBaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewModelBaseServiceClient(cc grpc.ClientConnInterface) ModelBaseServiceClient {
	return &modelBaseServiceClient{cc}
}

func (c *modelBaseServiceClient) GetById(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelBaseServiceClient) Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelBaseServiceClient) GetList(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelBaseServiceClient) DeleteById(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseService/DeleteById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelBaseServiceClient) Delete(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelBaseServiceClient) DeleteList(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseService/DeleteList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelBaseServiceClient) ForceDeleteList(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseService/ForceDeleteList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelBaseServiceClient) UpdateById(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseService/UpdateById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelBaseServiceClient) Update(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelBaseServiceClient) UpdateDoc(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseService/UpdateDoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelBaseServiceClient) Insert(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseService/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelBaseServiceClient) Count(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseService/Count", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelBaseServiceServer is the server API for ModelBaseService service.
// All implementations must embed UnimplementedModelBaseServiceServer
// for forward compatibility
type ModelBaseServiceServer interface {
	GetById(context.Context, *Request) (*Response, error)
	Get(context.Context, *Request) (*Response, error)
	GetList(context.Context, *Request) (*Response, error)
	DeleteById(context.Context, *Request) (*Response, error)
	Delete(context.Context, *Request) (*Response, error)
	DeleteList(context.Context, *Request) (*Response, error)
	ForceDeleteList(context.Context, *Request) (*Response, error)
	UpdateById(context.Context, *Request) (*Response, error)
	Update(context.Context, *Request) (*Response, error)
	UpdateDoc(context.Context, *Request) (*Response, error)
	Insert(context.Context, *Request) (*Response, error)
	Count(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedModelBaseServiceServer()
}

// UnimplementedModelBaseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedModelBaseServiceServer struct {
}

func (UnimplementedModelBaseServiceServer) GetById(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedModelBaseServiceServer) Get(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedModelBaseServiceServer) GetList(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedModelBaseServiceServer) DeleteById(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteById not implemented")
}
func (UnimplementedModelBaseServiceServer) Delete(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedModelBaseServiceServer) DeleteList(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteList not implemented")
}
func (UnimplementedModelBaseServiceServer) ForceDeleteList(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForceDeleteList not implemented")
}
func (UnimplementedModelBaseServiceServer) UpdateById(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateById not implemented")
}
func (UnimplementedModelBaseServiceServer) Update(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedModelBaseServiceServer) UpdateDoc(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDoc not implemented")
}
func (UnimplementedModelBaseServiceServer) Insert(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedModelBaseServiceServer) Count(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Count not implemented")
}
func (UnimplementedModelBaseServiceServer) mustEmbedUnimplementedModelBaseServiceServer() {}

// UnsafeModelBaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelBaseServiceServer will
// result in compilation errors.
type UnsafeModelBaseServiceServer interface {
	mustEmbedUnimplementedModelBaseServiceServer()
}

func RegisterModelBaseServiceServer(s grpc.ServiceRegistrar, srv ModelBaseServiceServer) {
	s.RegisterService(&ModelBaseService_ServiceDesc, srv)
}

func _ModelBaseService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceServer).GetById(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelBaseService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceServer).Get(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelBaseService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceServer).GetList(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelBaseService_DeleteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceServer).DeleteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseService/DeleteById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceServer).DeleteById(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelBaseService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceServer).Delete(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelBaseService_DeleteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceServer).DeleteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseService/DeleteList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceServer).DeleteList(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelBaseService_ForceDeleteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceServer).ForceDeleteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseService/ForceDeleteList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceServer).ForceDeleteList(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelBaseService_UpdateById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceServer).UpdateById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseService/UpdateById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceServer).UpdateById(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelBaseService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceServer).Update(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelBaseService_UpdateDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceServer).UpdateDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseService/UpdateDoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceServer).UpdateDoc(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelBaseService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseService/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceServer).Insert(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelBaseService_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseService/Count",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceServer).Count(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ModelBaseService_ServiceDesc is the grpc.ServiceDesc for ModelBaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModelBaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.ModelBaseService",
	HandlerType: (*ModelBaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetById",
			Handler:    _ModelBaseService_GetById_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ModelBaseService_Get_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _ModelBaseService_GetList_Handler,
		},
		{
			MethodName: "DeleteById",
			Handler:    _ModelBaseService_DeleteById_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ModelBaseService_Delete_Handler,
		},
		{
			MethodName: "DeleteList",
			Handler:    _ModelBaseService_DeleteList_Handler,
		},
		{
			MethodName: "ForceDeleteList",
			Handler:    _ModelBaseService_ForceDeleteList_Handler,
		},
		{
			MethodName: "UpdateById",
			Handler:    _ModelBaseService_UpdateById_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ModelBaseService_Update_Handler,
		},
		{
			MethodName: "UpdateDoc",
			Handler:    _ModelBaseService_UpdateDoc_Handler,
		},
		{
			MethodName: "Insert",
			Handler:    _ModelBaseService_Insert_Handler,
		},
		{
			MethodName: "Count",
			Handler:    _ModelBaseService_Count_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/model_base_service.proto",
}