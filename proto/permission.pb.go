// Code generated by protoc-gen-go. DO NOT EDIT.
// source: permission.proto

package permission

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreatePermissionRequest struct {
	// The ID of the file which is being permitted.
	FileID string `protobuf:"bytes,1,opt,name=fileID,proto3" json:"fileID,omitempty"`
	// The ID of the user that's given the permission.
	UserID string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	// The inheritors of the file which is being permitted to inherit the permission.
	Inheritors           []string `protobuf:"bytes,3,rep,name=inheritors,proto3" json:"inheritors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePermissionRequest) Reset()         { *m = CreatePermissionRequest{} }
func (m *CreatePermissionRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePermissionRequest) ProtoMessage()    {}
func (*CreatePermissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{0}
}

func (m *CreatePermissionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePermissionRequest.Unmarshal(m, b)
}
func (m *CreatePermissionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePermissionRequest.Marshal(b, m, deterministic)
}
func (m *CreatePermissionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePermissionRequest.Merge(m, src)
}
func (m *CreatePermissionRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePermissionRequest.Size(m)
}
func (m *CreatePermissionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePermissionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePermissionRequest proto.InternalMessageInfo

func (m *CreatePermissionRequest) GetFileID() string {
	if m != nil {
		return m.FileID
	}
	return ""
}

func (m *CreatePermissionRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *CreatePermissionRequest) GetInheritors() []string {
	if m != nil {
		return m.Inheritors
	}
	return nil
}

type CreatePermissionResponse struct {
	// The ID of the permission that's been created.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePermissionResponse) Reset()         { *m = CreatePermissionResponse{} }
func (m *CreatePermissionResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePermissionResponse) ProtoMessage()    {}
func (*CreatePermissionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{1}
}

func (m *CreatePermissionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePermissionResponse.Unmarshal(m, b)
}
func (m *CreatePermissionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePermissionResponse.Marshal(b, m, deterministic)
}
func (m *CreatePermissionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePermissionResponse.Merge(m, src)
}
func (m *CreatePermissionResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePermissionResponse.Size(m)
}
func (m *CreatePermissionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePermissionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePermissionResponse proto.InternalMessageInfo

func (m *CreatePermissionResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetPermissionRequest struct {
	// The ID of the permission to get.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPermissionRequest) Reset()         { *m = GetPermissionRequest{} }
func (m *GetPermissionRequest) String() string { return proto.CompactTextString(m) }
func (*GetPermissionRequest) ProtoMessage()    {}
func (*GetPermissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{2}
}

func (m *GetPermissionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPermissionRequest.Unmarshal(m, b)
}
func (m *GetPermissionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPermissionRequest.Marshal(b, m, deterministic)
}
func (m *GetPermissionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPermissionRequest.Merge(m, src)
}
func (m *GetPermissionRequest) XXX_Size() int {
	return xxx_messageInfo_GetPermissionRequest.Size(m)
}
func (m *GetPermissionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPermissionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPermissionRequest proto.InternalMessageInfo

func (m *GetPermissionRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeletePermissionRequest struct {
	// The ID of the permission to delete.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePermissionRequest) Reset()         { *m = DeletePermissionRequest{} }
func (m *DeletePermissionRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePermissionRequest) ProtoMessage()    {}
func (*DeletePermissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{3}
}

func (m *DeletePermissionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePermissionRequest.Unmarshal(m, b)
}
func (m *DeletePermissionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePermissionRequest.Marshal(b, m, deterministic)
}
func (m *DeletePermissionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePermissionRequest.Merge(m, src)
}
func (m *DeletePermissionRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePermissionRequest.Size(m)
}
func (m *DeletePermissionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePermissionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePermissionRequest proto.InternalMessageInfo

func (m *DeletePermissionRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type PermissionObject struct {
	// The ID of the permission.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The ID of the file which is being pemitted.
	FileID string `protobuf:"bytes,2,opt,name=fileID,proto3" json:"fileID,omitempty"`
	// The ID of the user that's given the permission.
	UserID string `protobuf:"bytes,3,opt,name=userID,proto3" json:"userID,omitempty"`
	// The ID of the file that this permission was inherited from.
	Inherited            string   `protobuf:"bytes,4,opt,name=inherited,proto3" json:"inherited,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PermissionObject) Reset()         { *m = PermissionObject{} }
func (m *PermissionObject) String() string { return proto.CompactTextString(m) }
func (*PermissionObject) ProtoMessage()    {}
func (*PermissionObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{4}
}

func (m *PermissionObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PermissionObject.Unmarshal(m, b)
}
func (m *PermissionObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PermissionObject.Marshal(b, m, deterministic)
}
func (m *PermissionObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermissionObject.Merge(m, src)
}
func (m *PermissionObject) XXX_Size() int {
	return xxx_messageInfo_PermissionObject.Size(m)
}
func (m *PermissionObject) XXX_DiscardUnknown() {
	xxx_messageInfo_PermissionObject.DiscardUnknown(m)
}

var xxx_messageInfo_PermissionObject proto.InternalMessageInfo

func (m *PermissionObject) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PermissionObject) GetFileID() string {
	if m != nil {
		return m.FileID
	}
	return ""
}

func (m *PermissionObject) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *PermissionObject) GetInherited() string {
	if m != nil {
		return m.Inherited
	}
	return ""
}

func init() {
	proto.RegisterType((*CreatePermissionRequest)(nil), "permission.CreatePermissionRequest")
	proto.RegisterType((*CreatePermissionResponse)(nil), "permission.CreatePermissionResponse")
	proto.RegisterType((*GetPermissionRequest)(nil), "permission.GetPermissionRequest")
	proto.RegisterType((*DeletePermissionRequest)(nil), "permission.DeletePermissionRequest")
	proto.RegisterType((*PermissionObject)(nil), "permission.PermissionObject")
}

func init() { proto.RegisterFile("permission.proto", fileDescriptor_c837ef01cbda0ad8) }

var fileDescriptor_c837ef01cbda0ad8 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x51, 0x4b, 0xc3, 0x30,
	0x10, 0xc7, 0xb7, 0x54, 0x06, 0x3d, 0x50, 0xca, 0x21, 0x2e, 0x8c, 0x21, 0x25, 0x8a, 0x4c, 0x1f,
	0xf6, 0xa0, 0x1f, 0xc1, 0x82, 0xec, 0x49, 0xed, 0x8b, 0x8f, 0xe2, 0xec, 0x89, 0x91, 0xd9, 0xd4,
	0x24, 0xfb, 0x30, 0x7e, 0x5b, 0x71, 0x2d, 0x26, 0x4d, 0x5b, 0xf7, 0x78, 0xff, 0xfc, 0x93, 0xbb,
	0xff, 0xef, 0x02, 0x49, 0x45, 0xfa, 0x53, 0x1a, 0x23, 0x55, 0xb9, 0xac, 0xb4, 0xb2, 0x0a, 0xc1,
	0x29, 0x42, 0xc2, 0xf4, 0x56, 0xd3, 0x8b, 0xa5, 0x87, 0x3f, 0x2d, 0xa7, 0xaf, 0x2d, 0x19, 0x8b,
	0x27, 0x30, 0x79, 0x93, 0x1b, 0x5a, 0x65, 0x7c, 0x9c, 0x8e, 0x17, 0x71, 0xde, 0x54, 0xbf, 0xfa,
	0xd6, 0x90, 0x5e, 0x65, 0x9c, 0xd5, 0x7a, 0x5d, 0xe1, 0x29, 0x80, 0x2c, 0xdf, 0x49, 0x4b, 0xab,
	0xb4, 0xe1, 0x51, 0x1a, 0x2d, 0xe2, 0xdc, 0x53, 0xc4, 0x15, 0xf0, 0x6e, 0x2b, 0x53, 0xa9, 0xd2,
	0x10, 0x1e, 0x01, 0x93, 0x45, 0xd3, 0x87, 0xc9, 0x42, 0x5c, 0xc0, 0xf1, 0x1d, 0xd9, 0xee, 0x4c,
	0xa1, 0xef, 0x12, 0xa6, 0x19, 0x6d, 0xa8, 0x6f, 0xfc, 0xd0, 0x5a, 0x41, 0xe2, 0x4c, 0xf7, 0xeb,
	0x0f, 0x7a, 0xed, 0x78, 0xbc, 0xc8, 0x6c, 0x20, 0x72, 0xd4, 0x8a, 0x3c, 0x87, 0xb8, 0x09, 0x48,
	0x05, 0x3f, 0xd8, 0x1d, 0x39, 0xe1, 0xfa, 0x9b, 0x01, 0xb8, 0x96, 0xf8, 0x0c, 0x49, 0x98, 0x1f,
	0xcf, 0x96, 0xde, 0x76, 0x06, 0x16, 0x31, 0x3b, 0xff, 0xdf, 0x54, 0x23, 0x14, 0x23, 0x7c, 0x84,
	0xc3, 0x16, 0x34, 0x4c, 0xfd, 0x8b, 0x7d, 0x3c, 0x67, 0x73, 0xdf, 0x11, 0xe2, 0x11, 0x23, 0x7c,
	0x82, 0x24, 0xe4, 0xdb, 0x9e, 0x79, 0x80, 0xfe, 0xbe, 0x87, 0xd7, 0x93, 0xdd, 0x57, 0xbc, 0xf9,
	0x09, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x89, 0x5a, 0x60, 0x9e, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PermissionClient is the client API for Permission service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PermissionClient interface {
	CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...grpc.CallOption) (*CreatePermissionResponse, error)
	GetPermission(ctx context.Context, in *GetPermissionRequest, opts ...grpc.CallOption) (*PermissionObject, error)
	DeletePermission(ctx context.Context, in *DeletePermissionRequest, opts ...grpc.CallOption) (*PermissionObject, error)
}

type permissionClient struct {
	cc *grpc.ClientConn
}

func NewPermissionClient(cc *grpc.ClientConn) PermissionClient {
	return &permissionClient{cc}
}

func (c *permissionClient) CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...grpc.CallOption) (*CreatePermissionResponse, error) {
	out := new(CreatePermissionResponse)
	err := c.cc.Invoke(ctx, "/permission.Permission/CreatePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionClient) GetPermission(ctx context.Context, in *GetPermissionRequest, opts ...grpc.CallOption) (*PermissionObject, error) {
	out := new(PermissionObject)
	err := c.cc.Invoke(ctx, "/permission.Permission/GetPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionClient) DeletePermission(ctx context.Context, in *DeletePermissionRequest, opts ...grpc.CallOption) (*PermissionObject, error) {
	out := new(PermissionObject)
	err := c.cc.Invoke(ctx, "/permission.Permission/DeletePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionServer is the server API for Permission service.
type PermissionServer interface {
	CreatePermission(context.Context, *CreatePermissionRequest) (*CreatePermissionResponse, error)
	GetPermission(context.Context, *GetPermissionRequest) (*PermissionObject, error)
	DeletePermission(context.Context, *DeletePermissionRequest) (*PermissionObject, error)
}

// UnimplementedPermissionServer can be embedded to have forward compatible implementations.
type UnimplementedPermissionServer struct {
}

func (*UnimplementedPermissionServer) CreatePermission(ctx context.Context, req *CreatePermissionRequest) (*CreatePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePermission not implemented")
}
func (*UnimplementedPermissionServer) GetPermission(ctx context.Context, req *GetPermissionRequest) (*PermissionObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermission not implemented")
}
func (*UnimplementedPermissionServer) DeletePermission(ctx context.Context, req *DeletePermissionRequest) (*PermissionObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePermission not implemented")
}

func RegisterPermissionServer(s *grpc.Server, srv PermissionServer) {
	s.RegisterService(&_Permission_serviceDesc, srv)
}

func _Permission_CreatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServer).CreatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permission.Permission/CreatePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServer).CreatePermission(ctx, req.(*CreatePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Permission_GetPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServer).GetPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permission.Permission/GetPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServer).GetPermission(ctx, req.(*GetPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Permission_DeletePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServer).DeletePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permission.Permission/DeletePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServer).DeletePermission(ctx, req.(*DeletePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Permission_serviceDesc = grpc.ServiceDesc{
	ServiceName: "permission.Permission",
	HandlerType: (*PermissionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePermission",
			Handler:    _Permission_CreatePermission_Handler,
		},
		{
			MethodName: "GetPermission",
			Handler:    _Permission_GetPermission_Handler,
		},
		{
			MethodName: "DeletePermission",
			Handler:    _Permission_DeletePermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "permission.proto",
}
