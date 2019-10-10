// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/protobuf/source_context.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// `SourceContext` represents information about the source of a
// protobuf element, like the file in which it is defined.
type SourceContext struct {
	// The path-qualified name of the .proto file that contained the associated
	// protobuf element.  For example: `"google/protobuf/source_context.proto"`.
	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
}

func (m *SourceContext) Reset()      { *m = SourceContext{} }
func (*SourceContext) ProtoMessage() {}
func (*SourceContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_b686cdb126d509db, []int{0}
}
func (m *SourceContext) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SourceContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SourceContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceContext.Merge(m, src)
}
func (m *SourceContext) XXX_Size() int {
	return m.ProtoSize()
}
func (m *SourceContext) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceContext.DiscardUnknown(m)
}

var xxx_messageInfo_SourceContext proto.InternalMessageInfo

func (*SourceContext) XXX_MessageName() string {
	return "google.protobuf.SourceContext"
}
func init() {
	proto.RegisterType((*SourceContext)(nil), "google.protobuf.SourceContext")
	golang_proto.RegisterType((*SourceContext)(nil), "google.protobuf.SourceContext")
}

func init() {
	proto.RegisterFile("google/protobuf/source_context.proto", fileDescriptor_b686cdb126d509db)
}
func init() {
	golang_proto.RegisterFile("google/protobuf/source_context.proto", fileDescriptor_b686cdb126d509db)
}

var fileDescriptor_b686cdb126d509db = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0xce, 0x2f, 0x2d,
	0x4a, 0x4e, 0x8d, 0x4f, 0xce, 0xcf, 0x2b, 0x49, 0xad, 0x28, 0xd1, 0x03, 0x8b, 0x0b, 0xf1, 0x43,
	0x54, 0xe9, 0xc1, 0x54, 0x29, 0xe9, 0x70, 0xf1, 0x06, 0x83, 0x15, 0x3a, 0x43, 0xd4, 0x09, 0x49,
	0x73, 0x71, 0xa6, 0x65, 0xe6, 0xa4, 0xc6, 0xe7, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x06, 0x71, 0x80, 0x04, 0xfc, 0x12, 0x73, 0x53, 0x9d, 0x8e, 0x32, 0x9e, 0x78, 0x28, 0xc7,
	0x70, 0xe3, 0xa1, 0x1c, 0x43, 0xc3, 0x23, 0x39, 0xc6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92,
	0x63, 0xbc, 0xf1, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x05, 0x8f, 0xe5, 0x18, 0x0f, 0x3c,
	0x96, 0x63, 0x3c, 0xf1, 0x58, 0x8e, 0xf1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xb8,
	0x84, 0x93, 0xf3, 0x73, 0xf5, 0xd0, 0x6c, 0x75, 0x12, 0x42, 0xb1, 0x33, 0x00, 0x24, 0x1c, 0xc0,
	0x18, 0xe5, 0x08, 0x55, 0x96, 0x9e, 0x9f, 0x93, 0x98, 0x97, 0xae, 0x97, 0x5f, 0x94, 0xae, 0x9f,
	0x9e, 0x9a, 0x07, 0xd6, 0x84, 0xcb, 0x5b, 0xd6, 0xa8, 0xdc, 0x45, 0x4c, 0xcc, 0xee, 0x01, 0x4e,
	0xab, 0x98, 0xe4, 0xdc, 0x21, 0x26, 0x05, 0x40, 0x75, 0xe9, 0x85, 0xa7, 0xe6, 0xe4, 0x78, 0xe7,
	0xe5, 0x97, 0xe7, 0x85, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x8d, 0x33, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x3d, 0x1d, 0x12, 0xc8, 0x35, 0x01, 0x00, 0x00,
}

func (m *SourceContext) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SourceContext) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SourceContext) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FileName) > 0 {
		i -= len(m.FileName)
		copy(dAtA[i:], m.FileName)
		i = encodeVarintSourceContext(dAtA, i, uint64(len(m.FileName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSourceContext(dAtA []byte, offset int, v uint64) int {
	offset -= sovSourceContext(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SourceContext) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FileName)
	if l > 0 {
		n += 1 + l + sovSourceContext(uint64(l))
	}
	return n
}

func sovSourceContext(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSourceContext(x uint64) (n int) {
	return sovSourceContext(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SourceContext) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SourceContext{`,
		`FileName:` + fmt.Sprintf("%v", this.FileName) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSourceContext(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SourceContext) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSourceContext
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
			return fmt.Errorf("proto: SourceContext: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SourceContext: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSourceContext
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
				return ErrInvalidLengthSourceContext
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSourceContext
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FileName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSourceContext(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSourceContext
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSourceContext
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
func skipSourceContext(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSourceContext
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
					return 0, ErrIntOverflowSourceContext
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
					return 0, ErrIntOverflowSourceContext
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
				return 0, ErrInvalidLengthSourceContext
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthSourceContext
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSourceContext
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
				next, err := skipSourceContext(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthSourceContext
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
	ErrInvalidLengthSourceContext = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSourceContext   = fmt.Errorf("proto: integer overflow")
)
