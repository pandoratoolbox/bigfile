// Code generated by protoc-gen-go. DO NOT EDIT.
// source: token_update.proto

package rpc

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// TokenUpdateRequest is used to describe the input format
// for updating token
type TokenUpdateRequest struct {
	AppUid               string                `protobuf:"bytes,1,opt,name=app_uid,json=appUid,proto3" json:"app_uid,omitempty"`
	AppSecret            string                `protobuf:"bytes,2,opt,name=app_secret,json=appSecret,proto3" json:"app_secret,omitempty"`
	Token                string                `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Path                 *wrappers.StringValue `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Ip                   *wrappers.StringValue `protobuf:"bytes,5,opt,name=ip,proto3" json:"ip,omitempty"`
	Secret               *wrappers.StringValue `protobuf:"bytes,6,opt,name=secret,proto3" json:"secret,omitempty"`
	AvailableTimes       *wrappers.UInt32Value `protobuf:"bytes,7,opt,name=available_times,json=availableTimes,proto3" json:"available_times,omitempty"`
	ReadOnly             *wrappers.BoolValue   `protobuf:"bytes,8,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
	ExpiredAt            *timestamp.Timestamp  `protobuf:"bytes,9,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TokenUpdateRequest) Reset()         { *m = TokenUpdateRequest{} }
func (m *TokenUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*TokenUpdateRequest) ProtoMessage()    {}
func (*TokenUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30424e7b5fab9ce2, []int{0}
}

func (m *TokenUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenUpdateRequest.Unmarshal(m, b)
}
func (m *TokenUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenUpdateRequest.Marshal(b, m, deterministic)
}
func (m *TokenUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenUpdateRequest.Merge(m, src)
}
func (m *TokenUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_TokenUpdateRequest.Size(m)
}
func (m *TokenUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TokenUpdateRequest proto.InternalMessageInfo

func (m *TokenUpdateRequest) GetAppUid() string {
	if m != nil {
		return m.AppUid
	}
	return ""
}

func (m *TokenUpdateRequest) GetAppSecret() string {
	if m != nil {
		return m.AppSecret
	}
	return ""
}

func (m *TokenUpdateRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *TokenUpdateRequest) GetPath() *wrappers.StringValue {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *TokenUpdateRequest) GetIp() *wrappers.StringValue {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *TokenUpdateRequest) GetSecret() *wrappers.StringValue {
	if m != nil {
		return m.Secret
	}
	return nil
}

func (m *TokenUpdateRequest) GetAvailableTimes() *wrappers.UInt32Value {
	if m != nil {
		return m.AvailableTimes
	}
	return nil
}

func (m *TokenUpdateRequest) GetReadOnly() *wrappers.BoolValue {
	if m != nil {
		return m.ReadOnly
	}
	return nil
}

func (m *TokenUpdateRequest) GetExpiredAt() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiredAt
	}
	return nil
}

// TokenUpdateResponse
type TokenUpdateResponse struct {
	RequestId            uint64   `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Token                *Token   `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenUpdateResponse) Reset()         { *m = TokenUpdateResponse{} }
func (m *TokenUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*TokenUpdateResponse) ProtoMessage()    {}
func (*TokenUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_30424e7b5fab9ce2, []int{1}
}

func (m *TokenUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenUpdateResponse.Unmarshal(m, b)
}
func (m *TokenUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenUpdateResponse.Marshal(b, m, deterministic)
}
func (m *TokenUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenUpdateResponse.Merge(m, src)
}
func (m *TokenUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_TokenUpdateResponse.Size(m)
}
func (m *TokenUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TokenUpdateResponse proto.InternalMessageInfo

func (m *TokenUpdateResponse) GetRequestId() uint64 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *TokenUpdateResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func init() {
	proto.RegisterType((*TokenUpdateRequest)(nil), "bigfile.token_update.TokenUpdateRequest")
	proto.RegisterType((*TokenUpdateResponse)(nil), "bigfile.token_update.TokenUpdateResponse")
}

func init() { proto.RegisterFile("token_update.proto", fileDescriptor_30424e7b5fab9ce2) }

var fileDescriptor_30424e7b5fab9ce2 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0x36, 0x69, 0xb7, 0xbb, 0x99, 0x82, 0xca, 0x58, 0x30, 0x14, 0x77, 0x5d, 0x7a, 0xaa, 0x22,
	0x53, 0xe9, 0x0a, 0xe2, 0xd1, 0x80, 0xc2, 0x9e, 0x2c, 0xd9, 0xd6, 0x83, 0x97, 0x38, 0x69, 0xde,
	0x66, 0x07, 0xd3, 0xcc, 0x73, 0x66, 0xa2, 0xf6, 0xdf, 0xf1, 0xe8, 0xd5, 0x7f, 0xce, 0xa3, 0x64,
	0x66, 0x5a, 0xba, 0xae, 0x85, 0x9e, 0xc2, 0x7b, 0xdf, 0x8f, 0x7c, 0xcc, 0xf7, 0x08, 0x35, 0xf2,
	0x0b, 0xd4, 0x59, 0x83, 0x05, 0x37, 0xc0, 0x50, 0x49, 0x23, 0xe9, 0x20, 0x17, 0xe5, 0xb5, 0xa8,
	0x80, 0xed, 0x62, 0xc3, 0xbe, 0x9d, 0x1c, 0x65, 0x78, 0x56, 0x4a, 0x59, 0x56, 0x30, 0xb1, 0x53,
	0xde, 0x5c, 0x4f, 0xbe, 0x2b, 0x8e, 0x08, 0x4a, 0x7b, 0xfc, 0xe9, 0xbf, 0xb8, 0x11, 0x2b, 0xd0,
	0x86, 0xaf, 0xd0, 0x11, 0x46, 0xbf, 0x3b, 0x84, 0xce, 0x5b, 0xc3, 0x85, 0x75, 0x4f, 0xe1, 0x6b,
	0x03, 0xda, 0xd0, 0xc7, 0xe4, 0x98, 0x23, 0x66, 0x8d, 0x28, 0xe2, 0xe0, 0x3c, 0x18, 0x47, 0x69,
	0x8f, 0x23, 0x2e, 0x44, 0x41, 0x4f, 0x09, 0x69, 0x01, 0x0d, 0x4b, 0x05, 0x26, 0x0e, 0x2d, 0x16,
	0x71, 0xc4, 0x2b, 0xbb, 0xa0, 0x03, 0x72, 0x64, 0xe3, 0xc5, 0x1d, 0x8b, 0xb8, 0x81, 0xbe, 0x24,
	0x5d, 0xe4, 0xe6, 0x26, 0xee, 0x9e, 0x07, 0xe3, 0xfe, 0xf4, 0x09, 0x73, 0xa1, 0xd8, 0x26, 0x14,
	0xbb, 0x32, 0x4a, 0xd4, 0xe5, 0x47, 0x5e, 0x35, 0x90, 0x5a, 0x26, 0x7d, 0x41, 0x42, 0x81, 0xf1,
	0xd1, 0x01, 0xfc, 0x50, 0x20, 0x7d, 0x45, 0x7a, 0x3e, 0x50, 0xef, 0x00, 0x85, 0xe7, 0xd2, 0x77,
	0xe4, 0x01, 0xff, 0xc6, 0x45, 0xc5, 0xf3, 0x0a, 0x32, 0xfb, 0x2e, 0xf1, 0xf1, 0x1e, 0xf9, 0xe2,
	0xb2, 0x36, 0x17, 0x53, 0x27, 0xbf, 0xbf, 0x15, 0xcd, 0x5b, 0x0d, 0x7d, 0x4d, 0x22, 0x05, 0xbc,
	0xc8, 0x64, 0x5d, 0xad, 0xe3, 0x13, 0x6b, 0x30, 0xbc, 0x63, 0x90, 0x48, 0x59, 0x39, 0xf9, 0x49,
	0x4b, 0xfe, 0x50, 0x57, 0x6b, 0xfa, 0x86, 0x10, 0xf8, 0x81, 0x42, 0x41, 0x91, 0x71, 0x13, 0x47,
	0x7b, 0x94, 0xf3, 0x4d, 0x61, 0x69, 0xe4, 0xd9, 0x6f, 0xcd, 0xe8, 0x33, 0x79, 0x74, 0xab, 0x34,
	0x8d, 0xb2, 0xd6, 0xd0, 0x96, 0xa3, 0x5c, 0x81, 0x99, 0x2f, 0xae, 0x9b, 0x46, 0x7e, 0x73, 0x59,
	0xd0, 0xe7, 0x9b, 0x72, 0x5c, 0x0f, 0x03, 0x76, 0xeb, 0xbe, 0x98, 0x75, 0xf4, 0x95, 0x4d, 0x35,
	0xe9, 0xef, 0xfc, 0x81, 0x16, 0xc4, 0x9d, 0x9d, 0x1f, 0xc7, 0xec, 0x7f, 0xa7, 0xc9, 0xee, 0x1e,
	0xd2, 0xf0, 0xd9, 0x01, 0x4c, 0x97, 0x7e, 0x74, 0x2f, 0xd1, 0x64, 0xb0, 0x94, 0xab, 0xad, 0x62,
	0xf3, 0x06, 0xc9, 0xc3, 0x1d, 0xfa, 0xac, 0x5d, 0xce, 0x82, 0x4f, 0x67, 0xa5, 0x30, 0x37, 0x4d,
	0xce, 0x96, 0x72, 0x35, 0xf1, 0x82, 0xed, 0x57, 0xe1, 0xf2, 0x4f, 0x10, 0xfc, 0x0c, 0x3b, 0xc9,
	0x2c, 0xfd, 0x15, 0x9e, 0x26, 0xde, 0x6f, 0xb6, 0x6d, 0x43, 0x94, 0xef, 0x45, 0x05, 0xf3, 0x35,
	0x82, 0xce, 0x7b, 0xf6, 0x37, 0x17, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xb8, 0xbf, 0xde,
	0x82, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TokenUpdateClient is the client API for TokenUpdate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TokenUpdateClient interface {
	TokenUpdate(ctx context.Context, in *TokenUpdateRequest, opts ...grpc.CallOption) (*TokenUpdateResponse, error)
}

type tokenUpdateClient struct {
	cc *grpc.ClientConn
}

func NewTokenUpdateClient(cc *grpc.ClientConn) TokenUpdateClient {
	return &tokenUpdateClient{cc}
}

func (c *tokenUpdateClient) TokenUpdate(ctx context.Context, in *TokenUpdateRequest, opts ...grpc.CallOption) (*TokenUpdateResponse, error) {
	out := new(TokenUpdateResponse)
	err := c.cc.Invoke(ctx, "/bigfile.token_update.TokenUpdate/tokenUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenUpdateServer is the server API for TokenUpdate service.
type TokenUpdateServer interface {
	TokenUpdate(context.Context, *TokenUpdateRequest) (*TokenUpdateResponse, error)
}

// UnimplementedTokenUpdateServer can be embedded to have forward compatible implementations.
type UnimplementedTokenUpdateServer struct {
}

func (*UnimplementedTokenUpdateServer) TokenUpdate(ctx context.Context, req *TokenUpdateRequest) (*TokenUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenUpdate not implemented")
}

func RegisterTokenUpdateServer(s *grpc.Server, srv TokenUpdateServer) {
	s.RegisterService(&_TokenUpdate_serviceDesc, srv)
}

func _TokenUpdate_TokenUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenUpdateServer).TokenUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bigfile.token_update.TokenUpdate/TokenUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenUpdateServer).TokenUpdate(ctx, req.(*TokenUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TokenUpdate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bigfile.token_update.TokenUpdate",
	HandlerType: (*TokenUpdateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "tokenUpdate",
			Handler:    _TokenUpdate_TokenUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "token_update.proto",
}
