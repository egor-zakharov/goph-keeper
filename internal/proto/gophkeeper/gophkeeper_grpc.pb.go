// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0--rc3
// source: pkg/proto/gophkeeper/gophkeeper.proto

package gophkeeper

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
	GophKeeper_SignUp_FullMethodName             = "/proto.gophkeeper.GophKeeper/SignUp"
	GophKeeper_SignIn_FullMethodName             = "/proto.gophkeeper.GophKeeper/SignIn"
	GophKeeper_CreateAuthData_FullMethodName     = "/proto.gophkeeper.GophKeeper/CreateAuthData"
	GophKeeper_GetAuthData_FullMethodName        = "/proto.gophkeeper.GophKeeper/GetAuthData"
	GophKeeper_UpdateAuthData_FullMethodName     = "/proto.gophkeeper.GophKeeper/UpdateAuthData"
	GophKeeper_DeleteAuthData_FullMethodName     = "/proto.gophkeeper.GophKeeper/DeleteAuthData"
	GophKeeper_CreateCard_FullMethodName         = "/proto.gophkeeper.GophKeeper/CreateCard"
	GophKeeper_GetCards_FullMethodName           = "/proto.gophkeeper.GophKeeper/GetCards"
	GophKeeper_UpdateCard_FullMethodName         = "/proto.gophkeeper.GophKeeper/UpdateCard"
	GophKeeper_DeleteCard_FullMethodName         = "/proto.gophkeeper.GophKeeper/DeleteCard"
	GophKeeper_CreateConfTextData_FullMethodName = "/proto.gophkeeper.GophKeeper/CreateConfTextData"
	GophKeeper_GetConfTextData_FullMethodName    = "/proto.gophkeeper.GophKeeper/GetConfTextData"
	GophKeeper_UpdateConfTextData_FullMethodName = "/proto.gophkeeper.GophKeeper/UpdateConfTextData"
	GophKeeper_DeleteConfTextData_FullMethodName = "/proto.gophkeeper.GophKeeper/DeleteConfTextData"
	GophKeeper_SubscribeToChanges_FullMethodName = "/proto.gophkeeper.GophKeeper/SubscribeToChanges"
	GophKeeper_GetFiles_FullMethodName           = "/proto.gophkeeper.GophKeeper/GetFiles"
	GophKeeper_DeleteFile_FullMethodName         = "/proto.gophkeeper.GophKeeper/DeleteFile"
	GophKeeper_UploadFile_FullMethodName         = "/proto.gophkeeper.GophKeeper/UploadFile"
	GophKeeper_DownloadFile_FullMethodName       = "/proto.gophkeeper.GophKeeper/DownloadFile"
)

// GophKeeperClient is the client API for GophKeeper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// GophKeeperServer service provides ability to store date securely
type GophKeeperClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	CreateAuthData(ctx context.Context, in *CreateAuthDataRequest, opts ...grpc.CallOption) (*CreateAuthDataResponse, error)
	GetAuthData(ctx context.Context, in *GetAuthDataRequest, opts ...grpc.CallOption) (*GetAuthDataResponse, error)
	UpdateAuthData(ctx context.Context, in *UpdateAuthDataRequest, opts ...grpc.CallOption) (*UpdateAuthDataResponse, error)
	DeleteAuthData(ctx context.Context, in *DeleteAuthDataRequest, opts ...grpc.CallOption) (*DeleteAuthDataResponse, error)
	CreateCard(ctx context.Context, in *CreateCardRequest, opts ...grpc.CallOption) (*CreateCardResponse, error)
	GetCards(ctx context.Context, in *GetCardsRequest, opts ...grpc.CallOption) (*GetCardsResponse, error)
	UpdateCard(ctx context.Context, in *UpdateCardRequest, opts ...grpc.CallOption) (*UpdateCardResponse, error)
	DeleteCard(ctx context.Context, in *DeleteCardRequest, opts ...grpc.CallOption) (*DeleteCardResponse, error)
	CreateConfTextData(ctx context.Context, in *CreateConfTextDataRequest, opts ...grpc.CallOption) (*CreateConfTextDataResponse, error)
	GetConfTextData(ctx context.Context, in *GetConfTextDataRequest, opts ...grpc.CallOption) (*GetConfTextDataResponse, error)
	UpdateConfTextData(ctx context.Context, in *UpdateConfTextDataRequest, opts ...grpc.CallOption) (*UpdateConfTextDataResponse, error)
	DeleteConfTextData(ctx context.Context, in *DeleteConfTextDataRequest, opts ...grpc.CallOption) (*DeleteConfTextDataResponse, error)
	SubscribeToChanges(ctx context.Context, in *SubscribeToChangesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[SubscribeToChangesResponse], error)
	GetFiles(ctx context.Context, in *GetFilesRequest, opts ...grpc.CallOption) (*GetFilesResponse, error)
	DeleteFile(ctx context.Context, in *DeleteFileRequest, opts ...grpc.CallOption) (*DeleteFileResponse, error)
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadFileRequest, UploadFileResponse], error)
	DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DownloadFileResponse], error)
}

type gophKeeperClient struct {
	cc grpc.ClientConnInterface
}

func NewGophKeeperClient(cc grpc.ClientConnInterface) GophKeeperClient {
	return &gophKeeperClient{cc}
}

func (c *gophKeeperClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, GophKeeper_SignUp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, GophKeeper_SignIn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) CreateAuthData(ctx context.Context, in *CreateAuthDataRequest, opts ...grpc.CallOption) (*CreateAuthDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAuthDataResponse)
	err := c.cc.Invoke(ctx, GophKeeper_CreateAuthData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) GetAuthData(ctx context.Context, in *GetAuthDataRequest, opts ...grpc.CallOption) (*GetAuthDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAuthDataResponse)
	err := c.cc.Invoke(ctx, GophKeeper_GetAuthData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) UpdateAuthData(ctx context.Context, in *UpdateAuthDataRequest, opts ...grpc.CallOption) (*UpdateAuthDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateAuthDataResponse)
	err := c.cc.Invoke(ctx, GophKeeper_UpdateAuthData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) DeleteAuthData(ctx context.Context, in *DeleteAuthDataRequest, opts ...grpc.CallOption) (*DeleteAuthDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteAuthDataResponse)
	err := c.cc.Invoke(ctx, GophKeeper_DeleteAuthData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) CreateCard(ctx context.Context, in *CreateCardRequest, opts ...grpc.CallOption) (*CreateCardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCardResponse)
	err := c.cc.Invoke(ctx, GophKeeper_CreateCard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) GetCards(ctx context.Context, in *GetCardsRequest, opts ...grpc.CallOption) (*GetCardsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCardsResponse)
	err := c.cc.Invoke(ctx, GophKeeper_GetCards_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) UpdateCard(ctx context.Context, in *UpdateCardRequest, opts ...grpc.CallOption) (*UpdateCardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCardResponse)
	err := c.cc.Invoke(ctx, GophKeeper_UpdateCard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) DeleteCard(ctx context.Context, in *DeleteCardRequest, opts ...grpc.CallOption) (*DeleteCardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCardResponse)
	err := c.cc.Invoke(ctx, GophKeeper_DeleteCard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) CreateConfTextData(ctx context.Context, in *CreateConfTextDataRequest, opts ...grpc.CallOption) (*CreateConfTextDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateConfTextDataResponse)
	err := c.cc.Invoke(ctx, GophKeeper_CreateConfTextData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) GetConfTextData(ctx context.Context, in *GetConfTextDataRequest, opts ...grpc.CallOption) (*GetConfTextDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetConfTextDataResponse)
	err := c.cc.Invoke(ctx, GophKeeper_GetConfTextData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) UpdateConfTextData(ctx context.Context, in *UpdateConfTextDataRequest, opts ...grpc.CallOption) (*UpdateConfTextDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateConfTextDataResponse)
	err := c.cc.Invoke(ctx, GophKeeper_UpdateConfTextData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) DeleteConfTextData(ctx context.Context, in *DeleteConfTextDataRequest, opts ...grpc.CallOption) (*DeleteConfTextDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteConfTextDataResponse)
	err := c.cc.Invoke(ctx, GophKeeper_DeleteConfTextData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) SubscribeToChanges(ctx context.Context, in *SubscribeToChangesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[SubscribeToChangesResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &GophKeeper_ServiceDesc.Streams[0], GophKeeper_SubscribeToChanges_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SubscribeToChangesRequest, SubscribeToChangesResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GophKeeper_SubscribeToChangesClient = grpc.ServerStreamingClient[SubscribeToChangesResponse]

func (c *gophKeeperClient) GetFiles(ctx context.Context, in *GetFilesRequest, opts ...grpc.CallOption) (*GetFilesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFilesResponse)
	err := c.cc.Invoke(ctx, GophKeeper_GetFiles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) DeleteFile(ctx context.Context, in *DeleteFileRequest, opts ...grpc.CallOption) (*DeleteFileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteFileResponse)
	err := c.cc.Invoke(ctx, GophKeeper_DeleteFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadFileRequest, UploadFileResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &GophKeeper_ServiceDesc.Streams[1], GophKeeper_UploadFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[UploadFileRequest, UploadFileResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GophKeeper_UploadFileClient = grpc.ClientStreamingClient[UploadFileRequest, UploadFileResponse]

func (c *gophKeeperClient) DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DownloadFileResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &GophKeeper_ServiceDesc.Streams[2], GophKeeper_DownloadFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[DownloadFileRequest, DownloadFileResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GophKeeper_DownloadFileClient = grpc.ServerStreamingClient[DownloadFileResponse]

// GophKeeperServer is the server API for GophKeeper service.
// All implementations must embed UnimplementedGophKeeperServer
// for forward compatibility.
//
// GophKeeperServer service provides ability to store date securely
type GophKeeperServer interface {
	SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	CreateAuthData(context.Context, *CreateAuthDataRequest) (*CreateAuthDataResponse, error)
	GetAuthData(context.Context, *GetAuthDataRequest) (*GetAuthDataResponse, error)
	UpdateAuthData(context.Context, *UpdateAuthDataRequest) (*UpdateAuthDataResponse, error)
	DeleteAuthData(context.Context, *DeleteAuthDataRequest) (*DeleteAuthDataResponse, error)
	CreateCard(context.Context, *CreateCardRequest) (*CreateCardResponse, error)
	GetCards(context.Context, *GetCardsRequest) (*GetCardsResponse, error)
	UpdateCard(context.Context, *UpdateCardRequest) (*UpdateCardResponse, error)
	DeleteCard(context.Context, *DeleteCardRequest) (*DeleteCardResponse, error)
	CreateConfTextData(context.Context, *CreateConfTextDataRequest) (*CreateConfTextDataResponse, error)
	GetConfTextData(context.Context, *GetConfTextDataRequest) (*GetConfTextDataResponse, error)
	UpdateConfTextData(context.Context, *UpdateConfTextDataRequest) (*UpdateConfTextDataResponse, error)
	DeleteConfTextData(context.Context, *DeleteConfTextDataRequest) (*DeleteConfTextDataResponse, error)
	SubscribeToChanges(*SubscribeToChangesRequest, grpc.ServerStreamingServer[SubscribeToChangesResponse]) error
	GetFiles(context.Context, *GetFilesRequest) (*GetFilesResponse, error)
	DeleteFile(context.Context, *DeleteFileRequest) (*DeleteFileResponse, error)
	UploadFile(grpc.ClientStreamingServer[UploadFileRequest, UploadFileResponse]) error
	DownloadFile(*DownloadFileRequest, grpc.ServerStreamingServer[DownloadFileResponse]) error
	mustEmbedUnimplementedGophKeeperServer()
}

// UnimplementedGophKeeperServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGophKeeperServer struct{}

func (UnimplementedGophKeeperServer) SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedGophKeeperServer) SignIn(context.Context, *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedGophKeeperServer) CreateAuthData(context.Context, *CreateAuthDataRequest) (*CreateAuthDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuthData not implemented")
}
func (UnimplementedGophKeeperServer) GetAuthData(context.Context, *GetAuthDataRequest) (*GetAuthDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthData not implemented")
}
func (UnimplementedGophKeeperServer) UpdateAuthData(context.Context, *UpdateAuthDataRequest) (*UpdateAuthDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuthData not implemented")
}
func (UnimplementedGophKeeperServer) DeleteAuthData(context.Context, *DeleteAuthDataRequest) (*DeleteAuthDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAuthData not implemented")
}
func (UnimplementedGophKeeperServer) CreateCard(context.Context, *CreateCardRequest) (*CreateCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCard not implemented")
}
func (UnimplementedGophKeeperServer) GetCards(context.Context, *GetCardsRequest) (*GetCardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCards not implemented")
}
func (UnimplementedGophKeeperServer) UpdateCard(context.Context, *UpdateCardRequest) (*UpdateCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCard not implemented")
}
func (UnimplementedGophKeeperServer) DeleteCard(context.Context, *DeleteCardRequest) (*DeleteCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCard not implemented")
}
func (UnimplementedGophKeeperServer) CreateConfTextData(context.Context, *CreateConfTextDataRequest) (*CreateConfTextDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConfTextData not implemented")
}
func (UnimplementedGophKeeperServer) GetConfTextData(context.Context, *GetConfTextDataRequest) (*GetConfTextDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfTextData not implemented")
}
func (UnimplementedGophKeeperServer) UpdateConfTextData(context.Context, *UpdateConfTextDataRequest) (*UpdateConfTextDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfTextData not implemented")
}
func (UnimplementedGophKeeperServer) DeleteConfTextData(context.Context, *DeleteConfTextDataRequest) (*DeleteConfTextDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConfTextData not implemented")
}
func (UnimplementedGophKeeperServer) SubscribeToChanges(*SubscribeToChangesRequest, grpc.ServerStreamingServer[SubscribeToChangesResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToChanges not implemented")
}
func (UnimplementedGophKeeperServer) GetFiles(context.Context, *GetFilesRequest) (*GetFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFiles not implemented")
}
func (UnimplementedGophKeeperServer) DeleteFile(context.Context, *DeleteFileRequest) (*DeleteFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}
func (UnimplementedGophKeeperServer) UploadFile(grpc.ClientStreamingServer[UploadFileRequest, UploadFileResponse]) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedGophKeeperServer) DownloadFile(*DownloadFileRequest, grpc.ServerStreamingServer[DownloadFileResponse]) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedGophKeeperServer) mustEmbedUnimplementedGophKeeperServer() {}
func (UnimplementedGophKeeperServer) testEmbeddedByValue()                    {}

// UnsafeGophKeeperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GophKeeperServer will
// result in compilation errors.
type UnsafeGophKeeperServer interface {
	mustEmbedUnimplementedGophKeeperServer()
}

func RegisterGophKeeperServer(s grpc.ServiceRegistrar, srv GophKeeperServer) {
	// If the following call pancis, it indicates UnimplementedGophKeeperServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GophKeeper_ServiceDesc, srv)
}

func _GophKeeper_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_SignUp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_SignIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_CreateAuthData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAuthDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).CreateAuthData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_CreateAuthData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).CreateAuthData(ctx, req.(*CreateAuthDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_GetAuthData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).GetAuthData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_GetAuthData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).GetAuthData(ctx, req.(*GetAuthDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_UpdateAuthData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAuthDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).UpdateAuthData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_UpdateAuthData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).UpdateAuthData(ctx, req.(*UpdateAuthDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_DeleteAuthData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAuthDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).DeleteAuthData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_DeleteAuthData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).DeleteAuthData(ctx, req.(*DeleteAuthDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_CreateCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).CreateCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_CreateCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).CreateCard(ctx, req.(*CreateCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_GetCards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).GetCards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_GetCards_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).GetCards(ctx, req.(*GetCardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_UpdateCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).UpdateCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_UpdateCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).UpdateCard(ctx, req.(*UpdateCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_DeleteCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).DeleteCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_DeleteCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).DeleteCard(ctx, req.(*DeleteCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_CreateConfTextData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConfTextDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).CreateConfTextData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_CreateConfTextData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).CreateConfTextData(ctx, req.(*CreateConfTextDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_GetConfTextData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfTextDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).GetConfTextData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_GetConfTextData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).GetConfTextData(ctx, req.(*GetConfTextDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_UpdateConfTextData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConfTextDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).UpdateConfTextData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_UpdateConfTextData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).UpdateConfTextData(ctx, req.(*UpdateConfTextDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_DeleteConfTextData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConfTextDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).DeleteConfTextData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_DeleteConfTextData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).DeleteConfTextData(ctx, req.(*DeleteConfTextDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_SubscribeToChanges_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeToChangesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GophKeeperServer).SubscribeToChanges(m, &grpc.GenericServerStream[SubscribeToChangesRequest, SubscribeToChangesResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GophKeeper_SubscribeToChangesServer = grpc.ServerStreamingServer[SubscribeToChangesResponse]

func _GophKeeper_GetFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).GetFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_GetFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).GetFiles(ctx, req.(*GetFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_DeleteFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).DeleteFile(ctx, req.(*DeleteFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GophKeeperServer).UploadFile(&grpc.GenericServerStream[UploadFileRequest, UploadFileResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GophKeeper_UploadFileServer = grpc.ClientStreamingServer[UploadFileRequest, UploadFileResponse]

func _GophKeeper_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GophKeeperServer).DownloadFile(m, &grpc.GenericServerStream[DownloadFileRequest, DownloadFileResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GophKeeper_DownloadFileServer = grpc.ServerStreamingServer[DownloadFileResponse]

// GophKeeper_ServiceDesc is the grpc.ServiceDesc for GophKeeper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GophKeeper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.gophkeeper.GophKeeper",
	HandlerType: (*GophKeeperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _GophKeeper_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _GophKeeper_SignIn_Handler,
		},
		{
			MethodName: "CreateAuthData",
			Handler:    _GophKeeper_CreateAuthData_Handler,
		},
		{
			MethodName: "GetAuthData",
			Handler:    _GophKeeper_GetAuthData_Handler,
		},
		{
			MethodName: "UpdateAuthData",
			Handler:    _GophKeeper_UpdateAuthData_Handler,
		},
		{
			MethodName: "DeleteAuthData",
			Handler:    _GophKeeper_DeleteAuthData_Handler,
		},
		{
			MethodName: "CreateCard",
			Handler:    _GophKeeper_CreateCard_Handler,
		},
		{
			MethodName: "GetCards",
			Handler:    _GophKeeper_GetCards_Handler,
		},
		{
			MethodName: "UpdateCard",
			Handler:    _GophKeeper_UpdateCard_Handler,
		},
		{
			MethodName: "DeleteCard",
			Handler:    _GophKeeper_DeleteCard_Handler,
		},
		{
			MethodName: "CreateConfTextData",
			Handler:    _GophKeeper_CreateConfTextData_Handler,
		},
		{
			MethodName: "GetConfTextData",
			Handler:    _GophKeeper_GetConfTextData_Handler,
		},
		{
			MethodName: "UpdateConfTextData",
			Handler:    _GophKeeper_UpdateConfTextData_Handler,
		},
		{
			MethodName: "DeleteConfTextData",
			Handler:    _GophKeeper_DeleteConfTextData_Handler,
		},
		{
			MethodName: "GetFiles",
			Handler:    _GophKeeper_GetFiles_Handler,
		},
		{
			MethodName: "DeleteFile",
			Handler:    _GophKeeper_DeleteFile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToChanges",
			Handler:       _GophKeeper_SubscribeToChanges_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UploadFile",
			Handler:       _GophKeeper_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _GophKeeper_DownloadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/proto/gophkeeper/gophkeeper.proto",
}
