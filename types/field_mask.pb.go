// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/protobuf/field_mask.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// `FieldMask` represents a set of symbolic field paths, for example:
//
//     paths: "f.a"
//     paths: "f.b.d"
//
// Here `f` represents a field in some root message, `a` and `b`
// fields in the message found in `f`, and `d` a field found in the
// message in `f.b`.
//
// Field masks are used to specify a subset of fields that should be
// returned by a get operation or modified by an update operation.
// Field masks also have a custom JSON encoding (see below).
//
// # Field Masks in Projections
//
// When used in the context of a projection, a response message or
// sub-message is filtered by the API to only contain those fields as
// specified in the mask. For example, if the mask in the previous
// example is applied to a response message as follows:
//
//     f {
//       a : 22
//       b {
//         d : 1
//         x : 2
//       }
//       y : 13
//     }
//     z: 8
//
// The result will not contain specific values for fields x,y and z
// (their value will be set to the default, and omitted in proto text
// output):
//
//
//     f {
//       a : 22
//       b {
//         d : 1
//       }
//     }
//
// A repeated field is not allowed except at the last position of a
// paths string.
//
// If a FieldMask object is not present in a get operation, the
// operation applies to all fields (as if a FieldMask of all fields
// had been specified).
//
// Note that a field mask does not necessarily apply to the
// top-level response message. In case of a REST get operation, the
// field mask applies directly to the response, but in case of a REST
// list operation, the mask instead applies to each individual message
// in the returned resource list. In case of a REST custom method,
// other definitions may be used. Where the mask applies will be
// clearly documented together with its declaration in the API.  In
// any case, the effect on the returned resource/resources is required
// behavior for APIs.
//
// # Field Masks in Update Operations
//
// A field mask in update operations specifies which fields of the
// targeted resource are going to be updated. The API is required
// to only change the values of the fields as specified in the mask
// and leave the others untouched. If a resource is passed in to
// describe the updated values, the API ignores the values of all
// fields not covered by the mask.
//
// If a repeated field is specified for an update operation, new values will
// be appended to the existing repeated field in the target resource. Note that
// a repeated field is only allowed in the last position of a `paths` string.
//
// If a sub-message is specified in the last position of the field mask for an
// update operation, then new value will be merged into the existing sub-message
// in the target resource.
//
// For example, given the target message:
//
//     f {
//       b {
//         d: 1
//         x: 2
//       }
//       c: [1]
//     }
//
// And an update message:
//
//     f {
//       b {
//         d: 10
//       }
//       c: [2]
//     }
//
// then if the field mask is:
//
//  paths: ["f.b", "f.c"]
//
// then the result will be:
//
//     f {
//       b {
//         d: 10
//         x: 2
//       }
//       c: [1, 2]
//     }
//
// An implementation may provide options to override this default behavior for
// repeated and message fields.
//
// In order to reset a field's value to the default, the field must
// be in the mask and set to the default value in the provided resource.
// Hence, in order to reset all fields of a resource, provide a default
// instance of the resource and set all fields in the mask, or do
// not provide a mask as described below.
//
// If a field mask is not present on update, the operation applies to
// all fields (as if a field mask of all fields has been specified).
// Note that in the presence of schema evolution, this may mean that
// fields the client does not know and has therefore not filled into
// the request will be reset to their default. If this is unwanted
// behavior, a specific service may require a client to always specify
// a field mask, producing an error if not.
//
// As with get operations, the location of the resource which
// describes the updated values in the request message depends on the
// operation kind. In any case, the effect of the field mask is
// required to be honored by the API.
//
// ## Considerations for HTTP REST
//
// The HTTP kind of an update operation which uses a field mask must
// be set to PATCH instead of PUT in order to satisfy HTTP semantics
// (PUT must only be used for full updates).
//
// # JSON Encoding of Field Masks
//
// In JSON, a field mask is encoded as a single string where paths are
// separated by a comma. Fields name in each path are converted
// to/from lower-camel naming conventions.
//
// As an example, consider the following message declarations:
//
//     message Profile {
//       User user = 1;
//       Photo photo = 2;
//     }
//     message User {
//       string display_name = 1;
//       string address = 2;
//     }
//
// In proto a field mask for `Profile` may look as such:
//
//     mask {
//       paths: "user.display_name"
//       paths: "photo"
//     }
//
// In JSON, the same mask is represented as below:
//
//     {
//       mask: "user.displayName,photo"
//     }
//
// # Field Masks and Oneof Fields
//
// Field masks treat fields in oneofs just as regular fields. Consider the
// following message:
//
//     message SampleMessage {
//       oneof test_oneof {
//         string name = 4;
//         SubMessage sub_message = 9;
//       }
//     }
//
// The field mask can be:
//
//     mask {
//       paths: "name"
//     }
//
// Or:
//
//     mask {
//       paths: "sub_message"
//     }
//
// Note that oneof type names ("test_oneof" in this case) cannot be used in
// paths.
//
// ## Field Mask Verification
//
// The implementation of any API method which has a FieldMask type field in the
// request should verify the included field paths, and return an
// `INVALID_ARGUMENT` error if any path is duplicated or unmappable.
type FieldMask struct {
	// The set of field mask paths.
	Paths []string `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
}

func (m *FieldMask) Reset()      { *m = FieldMask{} }
func (*FieldMask) ProtoMessage() {}
func (*FieldMask) Descriptor() ([]byte, []int) {
	return fileDescriptor_5158202634f0da48, []int{0}
}
func (m *FieldMask) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FieldMask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *FieldMask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldMask.Merge(m, src)
}
func (m *FieldMask) XXX_Size() int {
	return m.ProtoSize()
}
func (m *FieldMask) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldMask.DiscardUnknown(m)
}

var xxx_messageInfo_FieldMask proto.InternalMessageInfo

func (*FieldMask) XXX_MessageName() string {
	return "google.protobuf.FieldMask"
}
func init() {
	proto.RegisterType((*FieldMask)(nil), "google.protobuf.FieldMask")
	golang_proto.RegisterType((*FieldMask)(nil), "google.protobuf.FieldMask")
}

func init() { proto.RegisterFile("google/protobuf/field_mask.proto", fileDescriptor_5158202634f0da48) }
func init() {
	golang_proto.RegisterFile("google/protobuf/field_mask.proto", fileDescriptor_5158202634f0da48)
}

var fileDescriptor_5158202634f0da48 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x4f, 0xcb, 0x4c, 0xcd,
	0x49, 0x89, 0xcf, 0x4d, 0x2c, 0xce, 0xd6, 0x03, 0x8b, 0x09, 0xf1, 0x43, 0x54, 0xe8, 0xc1, 0x54,
	0x28, 0x29, 0x72, 0x71, 0xba, 0x81, 0x14, 0xf9, 0x26, 0x16, 0x67, 0x0b, 0x89, 0x70, 0xb1, 0x16,
	0x24, 0x96, 0x64, 0x14, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x06, 0x41, 0x38, 0x4e, 0x1d, 0x8c,
	0x27, 0x1e, 0xca, 0x31, 0xdc, 0x78, 0x28, 0xc7, 0xd0, 0xf0, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39,
	0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x6f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0x61, 0xc1, 0x63,
	0x39, 0xc6, 0x03, 0x8f, 0xe5, 0x18, 0x4f, 0x3c, 0x96, 0x63, 0xbc, 0xf0, 0x58, 0x8e, 0xe1, 0xc6,
	0x63, 0x39, 0x06, 0x2e, 0xe1, 0xe4, 0xfc, 0x5c, 0x3d, 0x34, 0xab, 0x9c, 0xf8, 0xe0, 0x16, 0x05,
	0x80, 0x84, 0x02, 0x18, 0xa3, 0x58, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x7f, 0x30, 0x32, 0x2e, 0x62,
	0x62, 0x76, 0x0f, 0x70, 0x5a, 0xc5, 0x24, 0xe7, 0x0e, 0xd1, 0x13, 0x00, 0xd5, 0xa3, 0x17, 0x9e,
	0x9a, 0x93, 0xe3, 0x9d, 0x97, 0x5f, 0x9e, 0x17, 0x02, 0x52, 0x99, 0xc4, 0x06, 0x36, 0xcc, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x83, 0xee, 0x65, 0xe9, 0x00, 0x00, 0x00,
}

func (m *FieldMask) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FieldMask) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Paths) > 0 {
		for _, s := range m.Paths {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeVarintFieldMask(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *FieldMask) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Paths) > 0 {
		for _, s := range m.Paths {
			l = len(s)
			n += 1 + l + sovFieldMask(uint64(l))
		}
	}
	return n
}

func sovFieldMask(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFieldMask(x uint64) (n int) {
	return sovFieldMask(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *FieldMask) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&FieldMask{`,
		`Paths:` + fmt.Sprintf("%v", this.Paths) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringFieldMask(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *FieldMask) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFieldMask
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
			return fmt.Errorf("proto: FieldMask: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FieldMask: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paths", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFieldMask
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
				return ErrInvalidLengthFieldMask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFieldMask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Paths = append(m.Paths, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFieldMask(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFieldMask
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFieldMask
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
func skipFieldMask(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFieldMask
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
					return 0, ErrIntOverflowFieldMask
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
					return 0, ErrIntOverflowFieldMask
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
				return 0, ErrInvalidLengthFieldMask
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthFieldMask
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFieldMask
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
				next, err := skipFieldMask(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthFieldMask
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
	ErrInvalidLengthFieldMask = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFieldMask   = fmt.Errorf("proto: integer overflow")
)
