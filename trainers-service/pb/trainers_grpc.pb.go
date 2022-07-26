// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: protos/trainers.proto

package pb

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

// TrainersServiceClient is the client API for TrainersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrainersServiceClient interface {
	Create(ctx context.Context, in *CreateTrainerRequest, opts ...grpc.CallOption) (*CreateTrainerResponse, error)
	Get(ctx context.Context, in *GetTrainerRequest, opts ...grpc.CallOption) (*GetTrainerResponse, error)
	List(ctx context.Context, in *ListTrainersRequest, opts ...grpc.CallOption) (*ListTrainersResponse, error)
	Update(ctx context.Context, in *UpdateTrainerRequest, opts ...grpc.CallOption) (*UpdateTrainerResponse, error)
}

type trainersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrainersServiceClient(cc grpc.ClientConnInterface) TrainersServiceClient {
	return &trainersServiceClient{cc}
}

func (c *trainersServiceClient) Create(ctx context.Context, in *CreateTrainerRequest, opts ...grpc.CallOption) (*CreateTrainerResponse, error) {
	out := new(CreateTrainerResponse)
	err := c.cc.Invoke(ctx, "/protos.TrainersService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainersServiceClient) Get(ctx context.Context, in *GetTrainerRequest, opts ...grpc.CallOption) (*GetTrainerResponse, error) {
	out := new(GetTrainerResponse)
	err := c.cc.Invoke(ctx, "/protos.TrainersService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainersServiceClient) List(ctx context.Context, in *ListTrainersRequest, opts ...grpc.CallOption) (*ListTrainersResponse, error) {
	out := new(ListTrainersResponse)
	err := c.cc.Invoke(ctx, "/protos.TrainersService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainersServiceClient) Update(ctx context.Context, in *UpdateTrainerRequest, opts ...grpc.CallOption) (*UpdateTrainerResponse, error) {
	out := new(UpdateTrainerResponse)
	err := c.cc.Invoke(ctx, "/protos.TrainersService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrainersServiceServer is the server API for TrainersService service.
// All implementations must embed UnimplementedTrainersServiceServer
// for forward compatibility
type TrainersServiceServer interface {
	Create(context.Context, *CreateTrainerRequest) (*CreateTrainerResponse, error)
	Get(context.Context, *GetTrainerRequest) (*GetTrainerResponse, error)
	List(context.Context, *ListTrainersRequest) (*ListTrainersResponse, error)
	Update(context.Context, *UpdateTrainerRequest) (*UpdateTrainerResponse, error)
	mustEmbedUnimplementedTrainersServiceServer()
}

// UnimplementedTrainersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTrainersServiceServer struct {
}

func (UnimplementedTrainersServiceServer) Create(context.Context, *CreateTrainerRequest) (*CreateTrainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTrainersServiceServer) Get(context.Context, *GetTrainerRequest) (*GetTrainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTrainersServiceServer) List(context.Context, *ListTrainersRequest) (*ListTrainersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTrainersServiceServer) Update(context.Context, *UpdateTrainerRequest) (*UpdateTrainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTrainersServiceServer) mustEmbedUnimplementedTrainersServiceServer() {}

// UnsafeTrainersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrainersServiceServer will
// result in compilation errors.
type UnsafeTrainersServiceServer interface {
	mustEmbedUnimplementedTrainersServiceServer()
}

func RegisterTrainersServiceServer(s grpc.ServiceRegistrar, srv TrainersServiceServer) {
	s.RegisterService(&TrainersService_ServiceDesc, srv)
}

func _TrainersService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTrainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainersServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TrainersService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainersServiceServer).Create(ctx, req.(*CreateTrainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainersService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTrainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainersServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TrainersService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainersServiceServer).Get(ctx, req.(*GetTrainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainersService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTrainersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainersServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TrainersService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainersServiceServer).List(ctx, req.(*ListTrainersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainersService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTrainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainersServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TrainersService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainersServiceServer).Update(ctx, req.(*UpdateTrainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TrainersService_ServiceDesc is the grpc.ServiceDesc for TrainersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrainersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.TrainersService",
	HandlerType: (*TrainersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TrainersService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TrainersService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _TrainersService_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TrainersService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/trainers.proto",
}
