// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: security/v1beta1/mtls.proto

package v1beta1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	_ "istio.io/api/type/v1beta1"
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

// Defines the acceptable connection TLS mode.
type MTLSPolicy_Mode int32

const (
	// Connection in plaintext.
	MTLSPolicy_NONE MTLSPolicy_Mode = 0
	// Connection can be either plaintext or TLS, and client cert can be omitted.
	MTLSPolicy_PERMISSIVE MTLSPolicy_Mode = 1
	// Client cert must be presented, connection is in TLS.
	MTLSPolicy_STRICT MTLSPolicy_Mode = 2
)

var MTLSPolicy_Mode_name = map[int32]string{
	0: "NONE",
	1: "PERMISSIVE",
	2: "STRICT",
}

var MTLSPolicy_Mode_value = map[string]int32{
	"NONE":       0,
	"PERMISSIVE": 1,
	"STRICT":     2,
}

func (x MTLSPolicy_Mode) String() string {
	return proto.EnumName(MTLSPolicy_Mode_name, int32(x))
}

func (MTLSPolicy_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5732f9202bce5de0, []int{0, 0}
}

type MTLSPolicy struct {
	// Defines the mode of mTLS authentication.
	Mode                 MTLSPolicy_Mode `protobuf:"varint,1,opt,name=mode,proto3,enum=istio.security.v1beta1.MTLSPolicy_Mode" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MTLSPolicy) Reset()         { *m = MTLSPolicy{} }
func (m *MTLSPolicy) String() string { return proto.CompactTextString(m) }
func (*MTLSPolicy) ProtoMessage()    {}
func (*MTLSPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_5732f9202bce5de0, []int{0}
}
func (m *MTLSPolicy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MTLSPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MTLSPolicy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MTLSPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MTLSPolicy.Merge(m, src)
}
func (m *MTLSPolicy) XXX_Size() int {
	return m.Size()
}
func (m *MTLSPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_MTLSPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_MTLSPolicy proto.InternalMessageInfo

func (m *MTLSPolicy) GetMode() MTLSPolicy_Mode {
	if m != nil {
		return m.Mode
	}
	return MTLSPolicy_NONE
}

func init() {
	proto.RegisterEnum("istio.security.v1beta1.MTLSPolicy_Mode", MTLSPolicy_Mode_name, MTLSPolicy_Mode_value)
	proto.RegisterType((*MTLSPolicy)(nil), "istio.security.v1beta1.MTLSPolicy")
}

func init() { proto.RegisterFile("security/v1beta1/mtls.proto", fileDescriptor_5732f9202bce5de0) }

var fileDescriptor_5732f9202bce5de0 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x4e, 0x4d, 0x2e,
	0x2d, 0xca, 0x2c, 0xa9, 0xd4, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0xcf, 0x2d, 0xc9,
	0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcb, 0x2c, 0x2e, 0xc9, 0xcc, 0xd7, 0x83,
	0x29, 0xd1, 0x83, 0x2a, 0x91, 0x92, 0x2e, 0xa9, 0x2c, 0x48, 0x85, 0x6b, 0x28, 0x4e, 0xcd, 0x49,
	0x4d, 0x2e, 0xc9, 0x2f, 0x82, 0x68, 0x52, 0x2a, 0xe7, 0xe2, 0xf2, 0x0d, 0xf1, 0x09, 0x0e, 0xc8,
	0xcf, 0xc9, 0x4c, 0xae, 0x14, 0xb2, 0xe6, 0x62, 0xc9, 0xcd, 0x4f, 0x49, 0x95, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x33, 0x52, 0xd7, 0xc3, 0x6e, 0xa2, 0x1e, 0x42, 0x87, 0x9e, 0x6f, 0x7e, 0x4a, 0x6a,
	0x10, 0x58, 0x93, 0x92, 0x0e, 0x17, 0x0b, 0x88, 0x27, 0xc4, 0xc1, 0xc5, 0xe2, 0xe7, 0xef, 0xe7,
	0x2a, 0xc0, 0x20, 0xc4, 0xc7, 0xc5, 0x15, 0xe0, 0x1a, 0xe4, 0xeb, 0x19, 0x1c, 0xec, 0x19, 0xe6,
	0x2a, 0xc0, 0x28, 0xc4, 0xc5, 0xc5, 0x16, 0x1c, 0x12, 0xe4, 0xe9, 0x1c, 0x22, 0xc0, 0xe4, 0xa4,
	0x7d, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x46, 0xc9, 0x42,
	0x6c, 0xca, 0xcc, 0xd7, 0x4f, 0x2c, 0xc8, 0xd4, 0x47, 0xf7, 0x65, 0x12, 0x1b, 0xd8, 0xb1, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0xc5, 0xe1, 0xa1, 0x00, 0x01, 0x00, 0x00,
}

func (m *MTLSPolicy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MTLSPolicy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MTLSPolicy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Mode != 0 {
		i = encodeVarintMtls(dAtA, i, uint64(m.Mode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMtls(dAtA []byte, offset int, v uint64) int {
	offset -= sovMtls(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MTLSPolicy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Mode != 0 {
		n += 1 + sovMtls(uint64(m.Mode))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMtls(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMtls(x uint64) (n int) {
	return sovMtls(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MTLSPolicy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMtls
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
			return fmt.Errorf("proto: MTLSPolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MTLSPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mode", wireType)
			}
			m.Mode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMtls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mode |= MTLSPolicy_Mode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMtls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMtls
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMtls
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
func skipMtls(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMtls
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
					return 0, ErrIntOverflowMtls
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
					return 0, ErrIntOverflowMtls
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
				return 0, ErrInvalidLengthMtls
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthMtls
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMtls
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
				next, err := skipMtls(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthMtls
				}
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
	ErrInvalidLengthMtls = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMtls   = fmt.Errorf("proto: integer overflow")
)
