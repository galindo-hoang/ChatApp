// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: app/chat_app/v1/relationship.proto

package v1

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
	RelationshipService_CreateNode_FullMethodName        = "/chat_app.v1.RelationshipService/CreateNode"
	RelationshipService_FollowingPerson_FullMethodName   = "/chat_app.v1.RelationshipService/FollowingPerson"
	RelationshipService_UnFollowingPerson_FullMethodName = "/chat_app.v1.RelationshipService/UnFollowingPerson"
)

// RelationshipServiceClient is the client API for RelationshipService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelationshipServiceClient interface {
	CreateNode(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	FollowingPerson(ctx context.Context, in *MakeRelationshipRequest, opts ...grpc.CallOption) (*MakeRelationshipResponse, error)
	UnFollowingPerson(ctx context.Context, in *MakeRelationshipRequest, opts ...grpc.CallOption) (*MakeRelationshipResponse, error)
}

type relationshipServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRelationshipServiceClient(cc grpc.ClientConnInterface) RelationshipServiceClient {
	return &relationshipServiceClient{cc}
}

func (c *relationshipServiceClient) CreateNode(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, RelationshipService_CreateNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipServiceClient) FollowingPerson(ctx context.Context, in *MakeRelationshipRequest, opts ...grpc.CallOption) (*MakeRelationshipResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MakeRelationshipResponse)
	err := c.cc.Invoke(ctx, RelationshipService_FollowingPerson_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipServiceClient) UnFollowingPerson(ctx context.Context, in *MakeRelationshipRequest, opts ...grpc.CallOption) (*MakeRelationshipResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MakeRelationshipResponse)
	err := c.cc.Invoke(ctx, RelationshipService_UnFollowingPerson_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelationshipServiceServer is the server API for RelationshipService service.
// All implementations must embed UnimplementedRelationshipServiceServer
// for forward compatibility.
type RelationshipServiceServer interface {
	CreateNode(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	FollowingPerson(context.Context, *MakeRelationshipRequest) (*MakeRelationshipResponse, error)
	UnFollowingPerson(context.Context, *MakeRelationshipRequest) (*MakeRelationshipResponse, error)
	mustEmbedUnimplementedRelationshipServiceServer()
}

// UnimplementedRelationshipServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRelationshipServiceServer struct{}

func (UnimplementedRelationshipServiceServer) CreateNode(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNode not implemented")
}
func (UnimplementedRelationshipServiceServer) FollowingPerson(context.Context, *MakeRelationshipRequest) (*MakeRelationshipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowingPerson not implemented")
}
func (UnimplementedRelationshipServiceServer) UnFollowingPerson(context.Context, *MakeRelationshipRequest) (*MakeRelationshipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFollowingPerson not implemented")
}
func (UnimplementedRelationshipServiceServer) mustEmbedUnimplementedRelationshipServiceServer() {}
func (UnimplementedRelationshipServiceServer) testEmbeddedByValue()                             {}

// UnsafeRelationshipServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelationshipServiceServer will
// result in compilation errors.
type UnsafeRelationshipServiceServer interface {
	mustEmbedUnimplementedRelationshipServiceServer()
}

func RegisterRelationshipServiceServer(s grpc.ServiceRegistrar, srv RelationshipServiceServer) {
	// If the following call pancis, it indicates UnimplementedRelationshipServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RelationshipService_ServiceDesc, srv)
}

func _RelationshipService_CreateNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServiceServer).CreateNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationshipService_CreateNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServiceServer).CreateNode(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationshipService_FollowingPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakeRelationshipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServiceServer).FollowingPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationshipService_FollowingPerson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServiceServer).FollowingPerson(ctx, req.(*MakeRelationshipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationshipService_UnFollowingPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakeRelationshipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServiceServer).UnFollowingPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationshipService_UnFollowingPerson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServiceServer).UnFollowingPerson(ctx, req.(*MakeRelationshipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RelationshipService_ServiceDesc is the grpc.ServiceDesc for RelationshipService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RelationshipService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat_app.v1.RelationshipService",
	HandlerType: (*RelationshipServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNode",
			Handler:    _RelationshipService_CreateNode_Handler,
		},
		{
			MethodName: "FollowingPerson",
			Handler:    _RelationshipService_FollowingPerson_Handler,
		},
		{
			MethodName: "UnFollowingPerson",
			Handler:    _RelationshipService_UnFollowingPerson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/chat_app/v1/relationship.proto",
}