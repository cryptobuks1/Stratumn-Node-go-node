// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/stratumn/go-node/core/app/monitoring/grpc/monitoring.proto

package grpc // import "github.com/stratumn/go-node/core/app/monitoring/grpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/stratumn/go-node/cli/grpc/ext"

import context "context"
import grpc "google.golang.org/grpc"

import encoding_binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SamplingRatio struct {
	Value                float32  `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SamplingRatio) Reset()         { *m = SamplingRatio{} }
func (m *SamplingRatio) String() string { return proto.CompactTextString(m) }
func (*SamplingRatio) ProtoMessage()    {}
func (*SamplingRatio) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitoring_bd197238a14d2b19, []int{0}
}
func (m *SamplingRatio) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SamplingRatio) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SamplingRatio.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SamplingRatio) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SamplingRatio.Merge(dst, src)
}
func (m *SamplingRatio) XXX_Size() int {
	return m.Size()
}
func (m *SamplingRatio) XXX_DiscardUnknown() {
	xxx_messageInfo_SamplingRatio.DiscardUnknown(m)
}

var xxx_messageInfo_SamplingRatio proto.InternalMessageInfo

func (m *SamplingRatio) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Ack struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ack) Reset()         { *m = Ack{} }
func (m *Ack) String() string { return proto.CompactTextString(m) }
func (*Ack) ProtoMessage()    {}
func (*Ack) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitoring_bd197238a14d2b19, []int{1}
}
func (m *Ack) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ack.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Ack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ack.Merge(dst, src)
}
func (m *Ack) XXX_Size() int {
	return m.Size()
}
func (m *Ack) XXX_DiscardUnknown() {
	xxx_messageInfo_Ack.DiscardUnknown(m)
}

var xxx_messageInfo_Ack proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SamplingRatio)(nil), "stratumn.node.core.app.monitoring.grpc.SamplingRatio")
	proto.RegisterType((*Ack)(nil), "stratumn.node.core.app.monitoring.grpc.Ack")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Monitoring service

type MonitoringClient interface {
	SetSamplingRatio(ctx context.Context, in *SamplingRatio, opts ...grpc.CallOption) (*Ack, error)
}

type monitoringClient struct {
	cc *grpc.ClientConn
}

func NewMonitoringClient(cc *grpc.ClientConn) MonitoringClient {
	return &monitoringClient{cc}
}

func (c *monitoringClient) SetSamplingRatio(ctx context.Context, in *SamplingRatio, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/stratumn.node.core.app.monitoring.grpc.Monitoring/SetSamplingRatio", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Monitoring service

type MonitoringServer interface {
	SetSamplingRatio(context.Context, *SamplingRatio) (*Ack, error)
}

func RegisterMonitoringServer(s *grpc.Server, srv MonitoringServer) {
	s.RegisterService(&_Monitoring_serviceDesc, srv)
}

func _Monitoring_SetSamplingRatio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SamplingRatio)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).SetSamplingRatio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stratumn.node.core.app.monitoring.grpc.Monitoring/SetSamplingRatio",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).SetSamplingRatio(ctx, req.(*SamplingRatio))
	}
	return interceptor(ctx, in, info, handler)
}

var _Monitoring_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stratumn.node.core.app.monitoring.grpc.Monitoring",
	HandlerType: (*MonitoringServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetSamplingRatio",
			Handler:    _Monitoring_SetSamplingRatio_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/stratumn/go-node/core/app/monitoring/grpc/monitoring.proto",
}

func (m *SamplingRatio) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SamplingRatio) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		dAtA[i] = 0xd
		i++
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Value))))
		i += 4
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Ack) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ack) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintMonitoring(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SamplingRatio) Size() (n int) {
	var l int
	_ = l
	if m.Value != 0 {
		n += 5
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Ack) Size() (n int) {
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMonitoring(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMonitoring(x uint64) (n int) {
	return sovMonitoring(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SamplingRatio) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMonitoring
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SamplingRatio: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SamplingRatio: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Value = float32(math.Float32frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipMonitoring(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMonitoring
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
func (m *Ack) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMonitoring
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Ack: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ack: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMonitoring(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMonitoring
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
func skipMonitoring(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMonitoring
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
					return 0, ErrIntOverflowMonitoring
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMonitoring
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMonitoring
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMonitoring
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMonitoring(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMonitoring = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMonitoring   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/stratumn/go-node/core/app/monitoring/grpc/monitoring.proto", fileDescriptor_monitoring_bd197238a14d2b19)
}

var fileDescriptor_monitoring_bd197238a14d2b19 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0x2e, 0x29, 0x4a, 0x2c, 0x29, 0xcd, 0xcd, 0xd3,
	0x4f, 0xcf, 0xd7, 0xcd, 0xcb, 0x4f, 0x49, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0x2c, 0x28,
	0xd0, 0xcf, 0xcd, 0xcf, 0xcb, 0x2c, 0xc9, 0x2f, 0xca, 0xcc, 0x4b, 0xd7, 0x4f, 0x2f, 0x2a, 0x48,
	0x46, 0xe2, 0xeb, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xa9, 0xc1, 0xf4, 0xea, 0x81, 0x34, 0xea,
	0x81, 0x34, 0xea, 0x25, 0x16, 0x14, 0xe8, 0x21, 0x29, 0x04, 0x69, 0x94, 0x32, 0xc2, 0x6b, 0x5d,
	0x4e, 0x26, 0xc4, 0xf8, 0xd4, 0x8a, 0x12, 0x10, 0x86, 0x98, 0xad, 0x54, 0xcc, 0xc5, 0x1b, 0x9c,
	0x98, 0x5b, 0x90, 0x93, 0x99, 0x97, 0x1e, 0x94, 0x58, 0x92, 0x99, 0x2f, 0x94, 0xc4, 0xc5, 0x5a,
	0x96, 0x98, 0x53, 0x9a, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xe4, 0xe4, 0xb3, 0x68, 0xb7, 0x84,
	0x4b, 0x48, 0x51, 0x62, 0x72, 0xaa, 0x42, 0x31, 0x54, 0x9d, 0x42, 0x11, 0x48, 0xa1, 0x9e, 0x42,
	0x68, 0x71, 0xaa, 0x82, 0x81, 0x9e, 0x81, 0x42, 0x49, 0xbe, 0x42, 0x5e, 0x6a, 0x59, 0x6a, 0x11,
	0x44, 0x41, 0xaa, 0x8e, 0x82, 0x21, 0x44, 0x30, 0x31, 0xa7, 0x3c, 0xb1, 0xb2, 0x18, 0x2a, 0xaa,
	0xb7, 0x62, 0xb7, 0x04, 0x63, 0x10, 0xc4, 0x68, 0x25, 0x56, 0x2e, 0x66, 0xc7, 0xe4, 0x6c, 0xa3,
	0xa5, 0x8c, 0x5c, 0x5c, 0xbe, 0x70, 0x3f, 0x08, 0x4d, 0x63, 0xe4, 0x12, 0x08, 0x4e, 0x2d, 0x41,
	0x75, 0x8e, 0xa9, 0x1e, 0x71, 0x9e, 0xd7, 0x43, 0xd1, 0x26, 0xa5, 0x4d, 0xac, 0x36, 0xc7, 0xe4,
	0x6c, 0x25, 0xb9, 0xa6, 0xad, 0x12, 0x52, 0xc1, 0xa9, 0x25, 0x0a, 0x25, 0x45, 0x89, 0xc9, 0x20,
	0x0f, 0xa2, 0xfa, 0xd4, 0xc9, 0xfb, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c,
	0x92, 0x63, 0x9c, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x92, 0x9c, 0x88, 0xb5, 0x06, 0x11, 0x49, 0x6c,
	0xe0, 0x70, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x03, 0xac, 0xb9, 0x75, 0x1c, 0x02, 0x00,
	0x00,
}