// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: celestiadomain/celestiadomain/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgRegisterDomain struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Domain  string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (m *MsgRegisterDomain) Reset()         { *m = MsgRegisterDomain{} }
func (m *MsgRegisterDomain) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterDomain) ProtoMessage()    {}
func (*MsgRegisterDomain) Descriptor() ([]byte, []int) {
	return fileDescriptor_194f5c85b2e28281, []int{0}
}
func (m *MsgRegisterDomain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterDomain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterDomain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterDomain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterDomain.Merge(m, src)
}
func (m *MsgRegisterDomain) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterDomain) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterDomain.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterDomain proto.InternalMessageInfo

func (m *MsgRegisterDomain) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgRegisterDomain) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

type MsgRegisterDomainResponse struct {
	Owner      string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Domain     string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Expiration uint64 `protobuf:"varint,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
}

func (m *MsgRegisterDomainResponse) Reset()         { *m = MsgRegisterDomainResponse{} }
func (m *MsgRegisterDomainResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterDomainResponse) ProtoMessage()    {}
func (*MsgRegisterDomainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_194f5c85b2e28281, []int{1}
}
func (m *MsgRegisterDomainResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterDomainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterDomainResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterDomainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterDomainResponse.Merge(m, src)
}
func (m *MsgRegisterDomainResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterDomainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterDomainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterDomainResponse proto.InternalMessageInfo

func (m *MsgRegisterDomainResponse) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgRegisterDomainResponse) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *MsgRegisterDomainResponse) GetExpiration() uint64 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

type MsgSetPrimaryDomain struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Domain  string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (m *MsgSetPrimaryDomain) Reset()         { *m = MsgSetPrimaryDomain{} }
func (m *MsgSetPrimaryDomain) String() string { return proto.CompactTextString(m) }
func (*MsgSetPrimaryDomain) ProtoMessage()    {}
func (*MsgSetPrimaryDomain) Descriptor() ([]byte, []int) {
	return fileDescriptor_194f5c85b2e28281, []int{2}
}
func (m *MsgSetPrimaryDomain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetPrimaryDomain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetPrimaryDomain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetPrimaryDomain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetPrimaryDomain.Merge(m, src)
}
func (m *MsgSetPrimaryDomain) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetPrimaryDomain) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetPrimaryDomain.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetPrimaryDomain proto.InternalMessageInfo

func (m *MsgSetPrimaryDomain) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSetPrimaryDomain) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

type MsgSetPrimaryDomainResponse struct {
}

func (m *MsgSetPrimaryDomainResponse) Reset()         { *m = MsgSetPrimaryDomainResponse{} }
func (m *MsgSetPrimaryDomainResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetPrimaryDomainResponse) ProtoMessage()    {}
func (*MsgSetPrimaryDomainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_194f5c85b2e28281, []int{3}
}
func (m *MsgSetPrimaryDomainResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetPrimaryDomainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetPrimaryDomainResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetPrimaryDomainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetPrimaryDomainResponse.Merge(m, src)
}
func (m *MsgSetPrimaryDomainResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetPrimaryDomainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetPrimaryDomainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetPrimaryDomainResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRegisterDomain)(nil), "celestiadomain.celestiadomain.MsgRegisterDomain")
	proto.RegisterType((*MsgRegisterDomainResponse)(nil), "celestiadomain.celestiadomain.MsgRegisterDomainResponse")
	proto.RegisterType((*MsgSetPrimaryDomain)(nil), "celestiadomain.celestiadomain.MsgSetPrimaryDomain")
	proto.RegisterType((*MsgSetPrimaryDomainResponse)(nil), "celestiadomain.celestiadomain.MsgSetPrimaryDomainResponse")
}

func init() {
	proto.RegisterFile("celestiadomain/celestiadomain/tx.proto", fileDescriptor_194f5c85b2e28281)
}

var fileDescriptor_194f5c85b2e28281 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4b, 0x4e, 0xcd, 0x49,
	0x2d, 0x2e, 0xc9, 0x4c, 0x4c, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0xd3, 0x47, 0xe3, 0x96, 0x54, 0xe8,
	0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xc9, 0xa2, 0x4a, 0xe8, 0xa1, 0x72, 0x95, 0x5c, 0xb9, 0x04,
	0x7d, 0x8b, 0xd3, 0x83, 0x52, 0xd3, 0x33, 0x8b, 0x4b, 0x52, 0x8b, 0x5c, 0xc0, 0x82, 0x42, 0x12,
	0x5c, 0xec, 0xc9, 0x45, 0xa9, 0x89, 0x25, 0xf9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x30, 0xae, 0x90, 0x18, 0x17, 0x1b, 0x44, 0xa3, 0x04, 0x13, 0x58, 0x02, 0xca, 0x53, 0xca, 0xe4,
	0x92, 0xc4, 0x30, 0x26, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x84, 0x8b, 0x35,
	0xbf, 0x3c, 0x2f, 0x15, 0x66, 0x18, 0x84, 0x83, 0xcb, 0x28, 0x21, 0x39, 0x2e, 0xae, 0xd4, 0x8a,
	0x82, 0xcc, 0xa2, 0xc4, 0x92, 0xcc, 0xfc, 0x3c, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x24,
	0x11, 0x25, 0x77, 0x2e, 0x61, 0xdf, 0xe2, 0xf4, 0xe0, 0xd4, 0x92, 0x80, 0xa2, 0xcc, 0xdc, 0xc4,
	0xa2, 0x4a, 0xb2, 0xdd, 0x2c, 0xcb, 0x25, 0x8d, 0xc5, 0x20, 0x98, 0xab, 0x8d, 0x3a, 0x98, 0xb8,
	0x98, 0x7d, 0x8b, 0xd3, 0x85, 0x6a, 0xb8, 0xf8, 0xd0, 0x82, 0xc7, 0x40, 0x0f, 0x6f, 0x98, 0xea,
	0x61, 0x84, 0x84, 0x94, 0x05, 0xa9, 0x3a, 0xe0, 0x61, 0xd7, 0xc4, 0xc8, 0x25, 0x80, 0xe1, 0x57,
	0x23, 0xc2, 0xc6, 0xa1, 0xeb, 0x91, 0xb2, 0x22, 0x5d, 0x0f, 0xcc, 0x11, 0x4e, 0x0e, 0x27, 0x1e,
	0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17,
	0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x05, 0x4f, 0x85, 0xba, 0xd0, 0x74, 0x57, 0x81,
	0x91, 0x10, 0x2b, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x89, 0xd1, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x8d, 0x80, 0x27, 0xe9, 0xb6, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	RegisterDomain(ctx context.Context, in *MsgRegisterDomain, opts ...grpc.CallOption) (*MsgRegisterDomainResponse, error)
	SetPrimaryDomain(ctx context.Context, in *MsgSetPrimaryDomain, opts ...grpc.CallOption) (*MsgSetPrimaryDomainResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RegisterDomain(ctx context.Context, in *MsgRegisterDomain, opts ...grpc.CallOption) (*MsgRegisterDomainResponse, error) {
	out := new(MsgRegisterDomainResponse)
	err := c.cc.Invoke(ctx, "/celestiadomain.celestiadomain.Msg/RegisterDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetPrimaryDomain(ctx context.Context, in *MsgSetPrimaryDomain, opts ...grpc.CallOption) (*MsgSetPrimaryDomainResponse, error) {
	out := new(MsgSetPrimaryDomainResponse)
	err := c.cc.Invoke(ctx, "/celestiadomain.celestiadomain.Msg/SetPrimaryDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	RegisterDomain(context.Context, *MsgRegisterDomain) (*MsgRegisterDomainResponse, error)
	SetPrimaryDomain(context.Context, *MsgSetPrimaryDomain) (*MsgSetPrimaryDomainResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RegisterDomain(ctx context.Context, req *MsgRegisterDomain) (*MsgRegisterDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDomain not implemented")
}
func (*UnimplementedMsgServer) SetPrimaryDomain(ctx context.Context, req *MsgSetPrimaryDomain) (*MsgSetPrimaryDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPrimaryDomain not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RegisterDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterDomain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/celestiadomain.celestiadomain.Msg/RegisterDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterDomain(ctx, req.(*MsgRegisterDomain))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetPrimaryDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetPrimaryDomain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetPrimaryDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/celestiadomain.celestiadomain.Msg/SetPrimaryDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetPrimaryDomain(ctx, req.(*MsgSetPrimaryDomain))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "celestiadomain.celestiadomain.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDomain",
			Handler:    _Msg_RegisterDomain_Handler,
		},
		{
			MethodName: "SetPrimaryDomain",
			Handler:    _Msg_SetPrimaryDomain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "celestiadomain/celestiadomain/tx.proto",
}

func (m *MsgRegisterDomain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterDomain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterDomain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Domain) > 0 {
		i -= len(m.Domain)
		copy(dAtA[i:], m.Domain)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Domain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRegisterDomainResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterDomainResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterDomainResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Expiration != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Expiration))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Domain) > 0 {
		i -= len(m.Domain)
		copy(dAtA[i:], m.Domain)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Domain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetPrimaryDomain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetPrimaryDomain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetPrimaryDomain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Domain) > 0 {
		i -= len(m.Domain)
		copy(dAtA[i:], m.Domain)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Domain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetPrimaryDomainResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetPrimaryDomainResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetPrimaryDomainResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRegisterDomain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRegisterDomainResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Expiration != 0 {
		n += 1 + sovTx(uint64(m.Expiration))
	}
	return n
}

func (m *MsgSetPrimaryDomain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSetPrimaryDomainResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRegisterDomain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRegisterDomain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterDomain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRegisterDomainResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRegisterDomainResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterDomainResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			m.Expiration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expiration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSetPrimaryDomain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSetPrimaryDomain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetPrimaryDomain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSetPrimaryDomainResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSetPrimaryDomainResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetPrimaryDomainResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
