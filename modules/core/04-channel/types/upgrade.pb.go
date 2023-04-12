// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/core/channel/v1/upgrade.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
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

// UpgradeTimeout defines a type which encapsulates the upgrade timeout values at which the counterparty
// must no longer proceed with the upgrade handshake.
type UpgradeTimeout struct {
	// block height after which the upgrade times out
	TimeoutHeight types.Height `protobuf:"bytes,1,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height"`
	// block timestamp (in nanoseconds) after which the upgrade times out
	TimeoutTimestamp uint64 `protobuf:"varint,2,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty"`
}

func (m *UpgradeTimeout) Reset()         { *m = UpgradeTimeout{} }
func (m *UpgradeTimeout) String() string { return proto.CompactTextString(m) }
func (*UpgradeTimeout) ProtoMessage()    {}
func (*UpgradeTimeout) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb1cef68588848b2, []int{0}
}
func (m *UpgradeTimeout) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpgradeTimeout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpgradeTimeout.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpgradeTimeout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpgradeTimeout.Merge(m, src)
}
func (m *UpgradeTimeout) XXX_Size() int {
	return m.Size()
}
func (m *UpgradeTimeout) XXX_DiscardUnknown() {
	xxx_messageInfo_UpgradeTimeout.DiscardUnknown(m)
}

var xxx_messageInfo_UpgradeTimeout proto.InternalMessageInfo

func (m *UpgradeTimeout) GetTimeoutHeight() types.Height {
	if m != nil {
		return m.TimeoutHeight
	}
	return types.Height{}
}

func (m *UpgradeTimeout) GetTimeoutTimestamp() uint64 {
	if m != nil {
		return m.TimeoutTimestamp
	}
	return 0
}

// ErrorReceipt defines a type which encapsulates the upgrade sequence and error associated with the
// upgrade handshake failure. When a channel upgrade handshake is aborted both chains are expected to increment to the
// next sequence.
type ErrorReceipt struct {
	// the channel upgrade sequence
	Sequence uint64 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// the error message detailing the cause of failure
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *ErrorReceipt) Reset()         { *m = ErrorReceipt{} }
func (m *ErrorReceipt) String() string { return proto.CompactTextString(m) }
func (*ErrorReceipt) ProtoMessage()    {}
func (*ErrorReceipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb1cef68588848b2, []int{1}
}
func (m *ErrorReceipt) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ErrorReceipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ErrorReceipt.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ErrorReceipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorReceipt.Merge(m, src)
}
func (m *ErrorReceipt) XXX_Size() int {
	return m.Size()
}
func (m *ErrorReceipt) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorReceipt.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorReceipt proto.InternalMessageInfo

func (m *ErrorReceipt) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *ErrorReceipt) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*UpgradeTimeout)(nil), "ibc.core.channel.v1.UpgradeTimeout")
	proto.RegisterType((*ErrorReceipt)(nil), "ibc.core.channel.v1.ErrorReceipt")
}

func init() { proto.RegisterFile("ibc/core/channel/v1/upgrade.proto", fileDescriptor_fb1cef68588848b2) }

var fileDescriptor_fb1cef68588848b2 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0x87, 0x1b, 0x99, 0xa2, 0x51, 0x87, 0xd6, 0x1d, 0x46, 0x0f, 0xdd, 0xdc, 0x69, 0x20, 0x4b,
	0x9c, 0x0a, 0xe2, 0x4d, 0x06, 0xa2, 0xe7, 0x3a, 0x2f, 0x5e, 0x64, 0xcd, 0x5e, 0xd2, 0xc0, 0xda,
	0xb7, 0x26, 0x69, 0xc1, 0x2f, 0xe0, 0xd9, 0x8f, 0xb5, 0xe3, 0x8e, 0x9e, 0x44, 0xb6, 0x2f, 0x22,
	0x6d, 0x36, 0x77, 0xca, 0xfb, 0xe7, 0xc9, 0xf3, 0xc2, 0x8f, 0x9e, 0xab, 0x58, 0x70, 0x81, 0x1a,
	0xb8, 0x48, 0x26, 0x59, 0x06, 0x33, 0x5e, 0x0e, 0x79, 0x91, 0x4b, 0x3d, 0x99, 0x02, 0xcb, 0x35,
	0x5a, 0xf4, 0xcf, 0x54, 0x2c, 0x58, 0x85, 0xb0, 0x35, 0xc2, 0xca, 0x61, 0xd0, 0x92, 0x28, 0xb1,
	0xde, 0xf3, 0xaa, 0x72, 0x68, 0xd0, 0xd9, 0xda, 0x66, 0x0a, 0x32, 0x5b, 0xc9, 0x5c, 0xe5, 0x80,
	0xde, 0x27, 0xa1, 0xcd, 0x17, 0x67, 0x1f, 0xab, 0x14, 0xb0, 0xb0, 0xfe, 0x23, 0x6d, 0x5a, 0x57,
	0xbe, 0x25, 0xa0, 0x64, 0x62, 0xdb, 0xa4, 0x4b, 0xfa, 0x87, 0x57, 0x01, 0xdb, 0xde, 0x75, 0x8a,
	0x72, 0xc8, 0x9e, 0x6a, 0x62, 0xd4, 0x98, 0xff, 0x74, 0xbc, 0xe8, 0x78, 0xfd, 0xcf, 0x0d, 0xfd,
	0x0b, 0x7a, 0xba, 0x11, 0x55, 0xaf, 0xb1, 0x93, 0x34, 0x6f, 0xef, 0x74, 0x49, 0xbf, 0x11, 0x9d,
	0xac, 0x17, 0xe3, 0xcd, 0xbc, 0x77, 0x4f, 0x8f, 0x1e, 0xb4, 0x46, 0x1d, 0x81, 0x00, 0x95, 0x5b,
	0x3f, 0xa0, 0xfb, 0x06, 0xde, 0x0b, 0xc8, 0x04, 0xd4, 0xf7, 0x1b, 0xd1, 0x7f, 0xef, 0xb7, 0xe8,
	0x2e, 0x54, 0x6c, 0x2d, 0x3b, 0x88, 0x5c, 0x33, 0x7a, 0x9e, 0x2f, 0x43, 0xb2, 0x58, 0x86, 0xe4,
	0x77, 0x19, 0x92, 0xaf, 0x55, 0xe8, 0x2d, 0x56, 0xa1, 0xf7, 0xbd, 0x0a, 0xbd, 0xd7, 0x3b, 0xa9,
	0x6c, 0x52, 0xc4, 0x4c, 0x60, 0xca, 0x05, 0x9a, 0x14, 0x0d, 0x57, 0xb1, 0x18, 0x48, 0xe4, 0xe5,
	0x2d, 0x4f, 0x71, 0x5a, 0xcc, 0xc0, 0xb8, 0x94, 0x2e, 0x6f, 0x06, 0x9b, 0xd8, 0xed, 0x47, 0x0e,
	0x26, 0xde, 0xab, 0x63, 0xba, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x11, 0x96, 0x44, 0x6d, 0x97,
	0x01, 0x00, 0x00,
}

func (m *UpgradeTimeout) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpgradeTimeout) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpgradeTimeout) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintUpgrade(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.TimeoutHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintUpgrade(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ErrorReceipt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ErrorReceipt) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ErrorReceipt) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintUpgrade(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x12
	}
	if m.Sequence != 0 {
		i = encodeVarintUpgrade(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintUpgrade(dAtA []byte, offset int, v uint64) int {
	offset -= sovUpgrade(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UpgradeTimeout) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TimeoutHeight.Size()
	n += 1 + l + sovUpgrade(uint64(l))
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovUpgrade(uint64(m.TimeoutTimestamp))
	}
	return n
}

func (m *ErrorReceipt) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sequence != 0 {
		n += 1 + sovUpgrade(uint64(m.Sequence))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovUpgrade(uint64(l))
	}
	return n
}

func sovUpgrade(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUpgrade(x uint64) (n int) {
	return sovUpgrade(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UpgradeTimeout) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpgrade
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
			return fmt.Errorf("proto: UpgradeTimeout: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpgradeTimeout: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
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
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TimeoutHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUpgrade(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUpgrade
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
func (m *ErrorReceipt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpgrade
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
			return fmt.Errorf("proto: ErrorReceipt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ErrorReceipt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
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
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUpgrade(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUpgrade
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
func skipUpgrade(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUpgrade
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
					return 0, ErrIntOverflowUpgrade
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
					return 0, ErrIntOverflowUpgrade
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
				return 0, ErrInvalidLengthUpgrade
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUpgrade
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUpgrade
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUpgrade        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUpgrade          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUpgrade = fmt.Errorf("proto: unexpected end of group")
)