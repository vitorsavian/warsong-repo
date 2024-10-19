// The greeting service definition.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/character.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Character_CreateCharacter_FullMethodName = "/proto.Character/CreateCharacter"
)

// CharacterClient is the client API for Character service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CharacterClient interface {
	CreateCharacter(ctx context.Context, in *CreateCharacterRequest, opts ...grpc.CallOption) (*CreateCharacterResponse, error)
}

type characterClient struct {
	cc grpc.ClientConnInterface
}

func NewCharacterClient(cc grpc.ClientConnInterface) CharacterClient {
	return &characterClient{cc}
}

func (c *characterClient) CreateCharacter(ctx context.Context, in *CreateCharacterRequest, opts ...grpc.CallOption) (*CreateCharacterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCharacterResponse)
	err := c.cc.Invoke(ctx, Character_CreateCharacter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CharacterServer is the server API for Character service.
// All implementations must embed UnimplementedCharacterServer
// for forward compatibility.
type CharacterServer interface {
	CreateCharacter(context.Context, *CreateCharacterRequest) (*CreateCharacterResponse, error)
	mustEmbedUnimplementedCharacterServer()
}

// UnimplementedCharacterServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCharacterServer struct{}

func (UnimplementedCharacterServer) CreateCharacter(context.Context, *CreateCharacterRequest) (*CreateCharacterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCharacter not implemented")
}
func (UnimplementedCharacterServer) mustEmbedUnimplementedCharacterServer() {}
func (UnimplementedCharacterServer) testEmbeddedByValue()                   {}

// UnsafeCharacterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CharacterServer will
// result in compilation errors.
type UnsafeCharacterServer interface {
	mustEmbedUnimplementedCharacterServer()
}

func RegisterCharacterServer(s grpc.ServiceRegistrar, srv CharacterServer) {
	// If the following call pancis, it indicates UnimplementedCharacterServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Character_ServiceDesc, srv)
}

func _Character_CreateCharacter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCharacterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterServer).CreateCharacter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Character_CreateCharacter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterServer).CreateCharacter(ctx, req.(*CreateCharacterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Character_ServiceDesc is the grpc.ServiceDesc for Character service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Character_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Character",
	HandlerType: (*CharacterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCharacter",
			Handler:    _Character_CreateCharacter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/character.proto",
}
