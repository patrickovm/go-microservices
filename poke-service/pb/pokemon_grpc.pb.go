// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: protos/pokemon.proto

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

// PokeServiceClient is the client API for PokeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PokeServiceClient interface {
	Create(ctx context.Context, in *CreatePokemonRequest, opts ...grpc.CallOption) (*CreatePokemonResponse, error)
	GetAllFromOneTrainer(ctx context.Context, in *GetAllFromOneTrainerRequest, opts ...grpc.CallOption) (*GetAllFromOneTrainerResponse, error)
	Update(ctx context.Context, in *UpdatePokemonRequest, opts ...grpc.CallOption) (*UpdatePokemonResponse, error)
}

type pokeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPokeServiceClient(cc grpc.ClientConnInterface) PokeServiceClient {
	return &pokeServiceClient{cc}
}

func (c *pokeServiceClient) Create(ctx context.Context, in *CreatePokemonRequest, opts ...grpc.CallOption) (*CreatePokemonResponse, error) {
	out := new(CreatePokemonResponse)
	err := c.cc.Invoke(ctx, "/protos.PokeService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokeServiceClient) GetAllFromOneTrainer(ctx context.Context, in *GetAllFromOneTrainerRequest, opts ...grpc.CallOption) (*GetAllFromOneTrainerResponse, error) {
	out := new(GetAllFromOneTrainerResponse)
	err := c.cc.Invoke(ctx, "/protos.PokeService/GetAllFromOneTrainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokeServiceClient) Update(ctx context.Context, in *UpdatePokemonRequest, opts ...grpc.CallOption) (*UpdatePokemonResponse, error) {
	out := new(UpdatePokemonResponse)
	err := c.cc.Invoke(ctx, "/protos.PokeService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PokeServiceServer is the server API for PokeService service.
// All implementations must embed UnimplementedPokeServiceServer
// for forward compatibility
type PokeServiceServer interface {
	Create(context.Context, *CreatePokemonRequest) (*CreatePokemonResponse, error)
	GetAllFromOneTrainer(context.Context, *GetAllFromOneTrainerRequest) (*GetAllFromOneTrainerResponse, error)
	Update(context.Context, *UpdatePokemonRequest) (*UpdatePokemonResponse, error)
	mustEmbedUnimplementedPokeServiceServer()
}

// UnimplementedPokeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPokeServiceServer struct {
}

func (UnimplementedPokeServiceServer) Create(context.Context, *CreatePokemonRequest) (*CreatePokemonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPokeServiceServer) GetAllFromOneTrainer(context.Context, *GetAllFromOneTrainerRequest) (*GetAllFromOneTrainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFromOneTrainer not implemented")
}
func (UnimplementedPokeServiceServer) Update(context.Context, *UpdatePokemonRequest) (*UpdatePokemonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPokeServiceServer) mustEmbedUnimplementedPokeServiceServer() {}

// UnsafePokeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PokeServiceServer will
// result in compilation errors.
type UnsafePokeServiceServer interface {
	mustEmbedUnimplementedPokeServiceServer()
}

func RegisterPokeServiceServer(s grpc.ServiceRegistrar, srv PokeServiceServer) {
	s.RegisterService(&PokeService_ServiceDesc, srv)
}

func _PokeService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePokemonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokeServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PokeService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokeServiceServer).Create(ctx, req.(*CreatePokemonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokeService_GetAllFromOneTrainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllFromOneTrainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokeServiceServer).GetAllFromOneTrainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PokeService/GetAllFromOneTrainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokeServiceServer).GetAllFromOneTrainer(ctx, req.(*GetAllFromOneTrainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokeService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePokemonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokeServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PokeService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokeServiceServer).Update(ctx, req.(*UpdatePokemonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PokeService_ServiceDesc is the grpc.ServiceDesc for PokeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PokeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.PokeService",
	HandlerType: (*PokeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PokeService_Create_Handler,
		},
		{
			MethodName: "GetAllFromOneTrainer",
			Handler:    _PokeService_GetAllFromOneTrainer_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PokeService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/pokemon.proto",
}
