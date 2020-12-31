// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/domo.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Domo post request message
type PostMsgReq struct {
	Domoid               string                 `protobuf:"bytes,1,opt,name=Domoid,proto3" json:"Domoid,omitempty"`
	Payload              string                 `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Timestamp            *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *PostMsgReq) Reset()         { *m = PostMsgReq{} }
func (m *PostMsgReq) String() string { return proto.CompactTextString(m) }
func (*PostMsgReq) ProtoMessage()    {}
func (*PostMsgReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6df02ec791369241, []int{0}
}
func (m *PostMsgReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PostMsgReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PostMsgReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PostMsgReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostMsgReq.Merge(m, src)
}
func (m *PostMsgReq) XXX_Size() int {
	return m.Size()
}
func (m *PostMsgReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PostMsgReq.DiscardUnknown(m)
}

var xxx_messageInfo_PostMsgReq proto.InternalMessageInfo

func (m *PostMsgReq) GetDomoid() string {
	if m != nil {
		return m.Domoid
	}
	return ""
}

func (m *PostMsgReq) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *PostMsgReq) GetTimestamp() *timestamppb.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

// Domo post response message
type PostDomoResp struct {
	Msg                  string                 `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Timestamp            *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *PostDomoResp) Reset()         { *m = PostDomoResp{} }
func (m *PostDomoResp) String() string { return proto.CompactTextString(m) }
func (*PostDomoResp) ProtoMessage()    {}
func (*PostDomoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6df02ec791369241, []int{1}
}
func (m *PostDomoResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PostDomoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PostDomoResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PostDomoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostDomoResp.Merge(m, src)
}
func (m *PostDomoResp) XXX_Size() int {
	return m.Size()
}
func (m *PostDomoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PostDomoResp.DiscardUnknown(m)
}

var xxx_messageInfo_PostDomoResp proto.InternalMessageInfo

func (m *PostDomoResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *PostDomoResp) GetTimestamp() *timestamppb.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*PostMsgReq)(nil), "pb.PostMsgReq")
	proto.RegisterType((*PostDomoResp)(nil), "pb.PostDomoResp")
}

func init() { proto.RegisterFile("pb/domo.proto", fileDescriptor_6df02ec791369241) }

var fileDescriptor_6df02ec791369241 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x48, 0xd2, 0x4f,
	0xc9, 0xcf, 0xcd, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0x92, 0x4f,
	0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x24, 0x95, 0xa6, 0xe9, 0x97, 0x64, 0xe6, 0xa6,
	0x16, 0x97, 0x24, 0xe6, 0x16, 0x40, 0x14, 0x29, 0x55, 0x70, 0x71, 0x05, 0xe4, 0x17, 0x97, 0xf8,
	0x16, 0xa7, 0x07, 0xa5, 0x16, 0x0a, 0x89, 0x71, 0xb1, 0xb9, 0xe4, 0xe7, 0xe6, 0x67, 0xa6, 0x48,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x79, 0x42, 0x12, 0x5c, 0xec, 0x05, 0x89, 0x95, 0x39,
	0xf9, 0x89, 0x29, 0x12, 0x4c, 0x60, 0x09, 0x18, 0x57, 0xc8, 0x82, 0x8b, 0x13, 0x6e, 0xa4, 0x04,
	0xb3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x94, 0x1e, 0xc4, 0x52, 0x3d, 0x98, 0xa5, 0x7a, 0x21, 0x30,
	0x15, 0x41, 0x08, 0xc5, 0x4a, 0x51, 0x5c, 0x3c, 0x20, 0x9b, 0x41, 0x36, 0x04, 0xa5, 0x16, 0x17,
	0x08, 0x09, 0x70, 0x31, 0xe7, 0x16, 0xa7, 0x43, 0x2d, 0x06, 0x31, 0x51, 0xcd, 0x66, 0x22, 0xc1,
	0x6c, 0x23, 0x1b, 0x2e, 0x6e, 0x90, 0xb9, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0xba,
	0x5c, 0xec, 0x50, 0x4f, 0x0a, 0xf1, 0xe9, 0x15, 0x24, 0xe9, 0x21, 0x7c, 0x2c, 0x25, 0x00, 0xe3,
	0xc3, 0xdc, 0xa1, 0xc4, 0xe0, 0x24, 0x70, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f,
	0x1e, 0xc9, 0x31, 0xce, 0x78, 0x2c, 0xc7, 0x90, 0xc4, 0x06, 0xb6, 0xce, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0xd6, 0x19, 0x86, 0x40, 0x62, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DomoServiceClient is the client API for DomoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DomoServiceClient interface {
	PostMsg(ctx context.Context, in *PostMsgReq, opts ...grpc.CallOption) (*PostDomoResp, error)
}

type domoServiceClient struct {
	cc *grpc.ClientConn
}

func NewDomoServiceClient(cc *grpc.ClientConn) DomoServiceClient {
	return &domoServiceClient{cc}
}

func (c *domoServiceClient) PostMsg(ctx context.Context, in *PostMsgReq, opts ...grpc.CallOption) (*PostDomoResp, error) {
	out := new(PostDomoResp)
	err := c.cc.Invoke(ctx, "/pb.DomoService/PostMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DomoServiceServer is the server API for DomoService service.
type DomoServiceServer interface {
	PostMsg(context.Context, *PostMsgReq) (*PostDomoResp, error)
}

// UnimplementedDomoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDomoServiceServer struct {
}

func (*UnimplementedDomoServiceServer) PostMsg(ctx context.Context, req *PostMsgReq) (*PostDomoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostMsg not implemented")
}

func RegisterDomoServiceServer(s *grpc.Server, srv DomoServiceServer) {
	s.RegisterService(&_DomoService_serviceDesc, srv)
}

func _DomoService_PostMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomoServiceServer).PostMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DomoService/PostMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomoServiceServer).PostMsg(ctx, req.(*PostMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _DomoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DomoService",
	HandlerType: (*DomoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostMsg",
			Handler:    _DomoService_PostMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/domo.proto",
}

func (m *PostMsgReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PostMsgReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PostMsgReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Timestamp != nil {
		{
			size, err := m.Timestamp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDomo(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintDomo(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Domoid) > 0 {
		i -= len(m.Domoid)
		copy(dAtA[i:], m.Domoid)
		i = encodeVarintDomo(dAtA, i, uint64(len(m.Domoid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PostDomoResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PostDomoResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PostDomoResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Timestamp != nil {
		{
			size, err := m.Timestamp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDomo(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintDomo(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDomo(dAtA []byte, offset int, v uint64) int {
	offset -= sovDomo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PostMsgReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Domoid)
	if l > 0 {
		n += 1 + l + sovDomo(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovDomo(uint64(l))
	}
	if m.Timestamp != nil {
		l = m.Timestamp.Size()
		n += 1 + l + sovDomo(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PostDomoResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovDomo(uint64(l))
	}
	if m.Timestamp != nil {
		l = m.Timestamp.Size()
		n += 1 + l + sovDomo(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDomo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDomo(x uint64) (n int) {
	return sovDomo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PostMsgReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDomo
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
			return fmt.Errorf("proto: PostMsgReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PostMsgReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domoid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomo
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
				return ErrInvalidLengthDomo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDomo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domoid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomo
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
				return ErrInvalidLengthDomo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDomo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDomo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDomo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = &timestamppb.Timestamp{}
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDomo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDomo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDomo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PostDomoResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDomo
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
			return fmt.Errorf("proto: PostDomoResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PostDomoResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomo
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
				return ErrInvalidLengthDomo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDomo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDomo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDomo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = &timestamppb.Timestamp{}
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDomo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDomo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDomo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDomo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDomo
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
					return 0, ErrIntOverflowDomo
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
					return 0, ErrIntOverflowDomo
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
				return 0, ErrInvalidLengthDomo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDomo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDomo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDomo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDomo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDomo = fmt.Errorf("proto: unexpected end of group")
)
