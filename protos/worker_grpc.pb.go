// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package nemon

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

// WorkerClient is the client API for Worker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkerClient interface {
	GetApps(ctx context.Context, in *GetAppsRequest, opts ...grpc.CallOption) (*GetAppsResponse, error)
	GetSysInfo(ctx context.Context, in *GetSysInfoRequest, opts ...grpc.CallOption) (*GetSysInfoResponse, error)
	DeleteApp(ctx context.Context, in *DeleteAppsRequest, opts ...grpc.CallOption) (*DeleteAppsResponse, error)
	IsEnrolled(ctx context.Context, in *IsEnrolledRequest, opts ...grpc.CallOption) (*IsEnrolledResponse, error)
	SaveEnrollmentInfo(ctx context.Context, in *SaveEnrollmentInfoRequest, opts ...grpc.CallOption) (*SaveEnrollmentInfoResponse, error)
	GetSaltAndSRP(ctx context.Context, in *GetSaltAndSRPRequest, opts ...grpc.CallOption) (*GetSaltAndSRPResponse, error)
	ExchangeEphemeralPublic(ctx context.Context, in *ExchangeEphemeralPublicRequest, opts ...grpc.CallOption) (*ExchangeEphemeralPublicResponse, error)
}

type workerClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerClient(cc grpc.ClientConnInterface) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) GetApps(ctx context.Context, in *GetAppsRequest, opts ...grpc.CallOption) (*GetAppsResponse, error) {
	out := new(GetAppsResponse)
	err := c.cc.Invoke(ctx, "/Worker/GetApps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) GetSysInfo(ctx context.Context, in *GetSysInfoRequest, opts ...grpc.CallOption) (*GetSysInfoResponse, error) {
	out := new(GetSysInfoResponse)
	err := c.cc.Invoke(ctx, "/Worker/GetSysInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) DeleteApp(ctx context.Context, in *DeleteAppsRequest, opts ...grpc.CallOption) (*DeleteAppsResponse, error) {
	out := new(DeleteAppsResponse)
	err := c.cc.Invoke(ctx, "/Worker/DeleteApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) IsEnrolled(ctx context.Context, in *IsEnrolledRequest, opts ...grpc.CallOption) (*IsEnrolledResponse, error) {
	out := new(IsEnrolledResponse)
	err := c.cc.Invoke(ctx, "/Worker/IsEnrolled", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) SaveEnrollmentInfo(ctx context.Context, in *SaveEnrollmentInfoRequest, opts ...grpc.CallOption) (*SaveEnrollmentInfoResponse, error) {
	out := new(SaveEnrollmentInfoResponse)
	err := c.cc.Invoke(ctx, "/Worker/SaveEnrollmentInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) GetSaltAndSRP(ctx context.Context, in *GetSaltAndSRPRequest, opts ...grpc.CallOption) (*GetSaltAndSRPResponse, error) {
	out := new(GetSaltAndSRPResponse)
	err := c.cc.Invoke(ctx, "/Worker/GetSaltAndSRP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) ExchangeEphemeralPublic(ctx context.Context, in *ExchangeEphemeralPublicRequest, opts ...grpc.CallOption) (*ExchangeEphemeralPublicResponse, error) {
	out := new(ExchangeEphemeralPublicResponse)
	err := c.cc.Invoke(ctx, "/Worker/ExchangeEphemeralPublic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServer is the server API for Worker service.
// All implementations must embed UnimplementedWorkerServer
// for forward compatibility
type WorkerServer interface {
	GetApps(context.Context, *GetAppsRequest) (*GetAppsResponse, error)
	GetSysInfo(context.Context, *GetSysInfoRequest) (*GetSysInfoResponse, error)
	DeleteApp(context.Context, *DeleteAppsRequest) (*DeleteAppsResponse, error)
	IsEnrolled(context.Context, *IsEnrolledRequest) (*IsEnrolledResponse, error)
	SaveEnrollmentInfo(context.Context, *SaveEnrollmentInfoRequest) (*SaveEnrollmentInfoResponse, error)
	GetSaltAndSRP(context.Context, *GetSaltAndSRPRequest) (*GetSaltAndSRPResponse, error)
	ExchangeEphemeralPublic(context.Context, *ExchangeEphemeralPublicRequest) (*ExchangeEphemeralPublicResponse, error)
	mustEmbedUnimplementedWorkerServer()
}

// UnimplementedWorkerServer must be embedded to have forward compatible implementations.
type UnimplementedWorkerServer struct {
}

func (UnimplementedWorkerServer) GetApps(context.Context, *GetAppsRequest) (*GetAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApps not implemented")
}
func (UnimplementedWorkerServer) GetSysInfo(context.Context, *GetSysInfoRequest) (*GetSysInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSysInfo not implemented")
}
func (UnimplementedWorkerServer) DeleteApp(context.Context, *DeleteAppsRequest) (*DeleteAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApp not implemented")
}
func (UnimplementedWorkerServer) IsEnrolled(context.Context, *IsEnrolledRequest) (*IsEnrolledResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsEnrolled not implemented")
}
func (UnimplementedWorkerServer) SaveEnrollmentInfo(context.Context, *SaveEnrollmentInfoRequest) (*SaveEnrollmentInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveEnrollmentInfo not implemented")
}
func (UnimplementedWorkerServer) GetSaltAndSRP(context.Context, *GetSaltAndSRPRequest) (*GetSaltAndSRPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSaltAndSRP not implemented")
}
func (UnimplementedWorkerServer) ExchangeEphemeralPublic(context.Context, *ExchangeEphemeralPublicRequest) (*ExchangeEphemeralPublicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExchangeEphemeralPublic not implemented")
}
func (UnimplementedWorkerServer) mustEmbedUnimplementedWorkerServer() {}

// UnsafeWorkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkerServer will
// result in compilation errors.
type UnsafeWorkerServer interface {
	mustEmbedUnimplementedWorkerServer()
}

func RegisterWorkerServer(s grpc.ServiceRegistrar, srv WorkerServer) {
	s.RegisterService(&Worker_ServiceDesc, srv)
}

func _Worker_GetApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).GetApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Worker/GetApps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).GetApps(ctx, req.(*GetAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_GetSysInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSysInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).GetSysInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Worker/GetSysInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).GetSysInfo(ctx, req.(*GetSysInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_DeleteApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).DeleteApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Worker/DeleteApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).DeleteApp(ctx, req.(*DeleteAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_IsEnrolled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsEnrolledRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).IsEnrolled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Worker/IsEnrolled",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).IsEnrolled(ctx, req.(*IsEnrolledRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_SaveEnrollmentInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveEnrollmentInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).SaveEnrollmentInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Worker/SaveEnrollmentInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).SaveEnrollmentInfo(ctx, req.(*SaveEnrollmentInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_GetSaltAndSRP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSaltAndSRPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).GetSaltAndSRP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Worker/GetSaltAndSRP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).GetSaltAndSRP(ctx, req.(*GetSaltAndSRPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_ExchangeEphemeralPublic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangeEphemeralPublicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).ExchangeEphemeralPublic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Worker/ExchangeEphemeralPublic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).ExchangeEphemeralPublic(ctx, req.(*ExchangeEphemeralPublicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Worker_ServiceDesc is the grpc.ServiceDesc for Worker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Worker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetApps",
			Handler:    _Worker_GetApps_Handler,
		},
		{
			MethodName: "GetSysInfo",
			Handler:    _Worker_GetSysInfo_Handler,
		},
		{
			MethodName: "DeleteApp",
			Handler:    _Worker_DeleteApp_Handler,
		},
		{
			MethodName: "IsEnrolled",
			Handler:    _Worker_IsEnrolled_Handler,
		},
		{
			MethodName: "SaveEnrollmentInfo",
			Handler:    _Worker_SaveEnrollmentInfo_Handler,
		},
		{
			MethodName: "GetSaltAndSRP",
			Handler:    _Worker_GetSaltAndSRP_Handler,
		},
		{
			MethodName: "ExchangeEphemeralPublic",
			Handler:    _Worker_ExchangeEphemeralPublic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/worker.proto",
}
