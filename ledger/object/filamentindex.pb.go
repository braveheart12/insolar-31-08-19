// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ledger/object/filamentindex.proto

package object

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_insolar_insolar_insolar "github.com/insolar/insolar/insolar"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type FilamentIndex struct {
	XPolymorph       int32                                          `protobuf:"varint,16,opt,name=__polymorph,json=Polymorph,proto3" json:"__polymorph,omitempty"`
	ObjID            github_com_insolar_insolar_insolar.ID          `protobuf:"bytes,20,opt,name=ObjID,proto3,customtype=github.com/insolar/insolar/insolar.ID" json:"ObjID"`
	Lifeline         Lifeline                                       `protobuf:"bytes,21,opt,name=Lifeline,proto3" json:"Lifeline"`
	LifelineLastUsed github_com_insolar_insolar_insolar.PulseNumber `protobuf:"varint,22,opt,name=LifelineLastUsed,proto3,customtype=github.com/insolar/insolar/insolar.PulseNumber" json:"LifelineLastUsed"`
	PendingRecords   []github_com_insolar_insolar_insolar.ID        `protobuf:"bytes,23,rep,name=PendingRecords,proto3,customtype=github.com/insolar/insolar/insolar.ID" json:"PendingRecords"`
}

func (m *FilamentIndex) Reset()      { *m = FilamentIndex{} }
func (*FilamentIndex) ProtoMessage() {}
func (*FilamentIndex) Descriptor() ([]byte, []int) {
	return fileDescriptor_714fa1835dbaf271, []int{0}
}
func (m *FilamentIndex) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FilamentIndex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FilamentIndex.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FilamentIndex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilamentIndex.Merge(m, src)
}
func (m *FilamentIndex) XXX_Size() int {
	return m.Size()
}
func (m *FilamentIndex) XXX_DiscardUnknown() {
	xxx_messageInfo_FilamentIndex.DiscardUnknown(m)
}

var xxx_messageInfo_FilamentIndex proto.InternalMessageInfo

func (m *FilamentIndex) GetXPolymorph() int32 {
	if m != nil {
		return m.XPolymorph
	}
	return 0
}

func (m *FilamentIndex) GetLifeline() Lifeline {
	if m != nil {
		return m.Lifeline
	}
	return Lifeline{}
}

func init() {
	proto.RegisterType((*FilamentIndex)(nil), "object.FilamentIndex")
}

func init() { proto.RegisterFile("ledger/object/filamentindex.proto", fileDescriptor_714fa1835dbaf271) }

var fileDescriptor_714fa1835dbaf271 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcc, 0x49, 0x4d, 0x49,
	0x4f, 0x2d, 0xd2, 0xcf, 0x4f, 0xca, 0x4a, 0x4d, 0x2e, 0xd1, 0x4f, 0xcb, 0xcc, 0x49, 0xcc, 0x4d,
	0xcd, 0x2b, 0xc9, 0xcc, 0x4b, 0x49, 0xad, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83,
	0xc8, 0x49, 0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7,
	0xa7, 0xe7, 0xeb, 0x83, 0xa5, 0x93, 0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0x68, 0x93,
	0x92, 0x41, 0x35, 0x39, 0x27, 0x33, 0x2d, 0x35, 0x27, 0x33, 0x2f, 0x15, 0x22, 0xab, 0xf4, 0x92,
	0x89, 0x8b, 0xd7, 0x0d, 0x6a, 0x99, 0x27, 0xc8, 0x32, 0x21, 0x39, 0x2e, 0xee, 0xf8, 0xf8, 0x82,
	0xfc, 0x9c, 0xca, 0xdc, 0xfc, 0xa2, 0x82, 0x0c, 0x09, 0x01, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0xce,
	0x00, 0x98, 0x80, 0x90, 0x33, 0x17, 0xab, 0x7f, 0x52, 0x96, 0xa7, 0x8b, 0x84, 0x88, 0x02, 0xa3,
	0x06, 0x8f, 0x93, 0xee, 0x89, 0x7b, 0xf2, 0x0c, 0xb7, 0xee, 0xc9, 0xab, 0x22, 0xb9, 0x2a, 0x33,
	0xaf, 0x38, 0x3f, 0x27, 0xb1, 0x08, 0x9d, 0xd6, 0xf3, 0x74, 0x09, 0x82, 0xe8, 0x15, 0x32, 0xe2,
	0xe2, 0xf0, 0x81, 0x3a, 0x44, 0x42, 0x54, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x40, 0x0f, 0xe2, 0x40,
	0x3d, 0x98, 0xb8, 0x13, 0x0b, 0xc8, 0xe4, 0x20, 0xb8, 0x3a, 0xa1, 0x24, 0x2e, 0x01, 0x18, 0xdb,
	0x27, 0xb1, 0xb8, 0x24, 0xb4, 0x38, 0x35, 0x45, 0x42, 0x4c, 0x81, 0x51, 0x83, 0xd7, 0xc9, 0x0c,
	0xea, 0x06, 0x3d, 0x22, 0xdc, 0x10, 0x50, 0x9a, 0x53, 0x9c, 0xea, 0x57, 0x9a, 0x9b, 0x94, 0x5a,
	0x14, 0x84, 0x61, 0x9e, 0x50, 0x28, 0x17, 0x5f, 0x40, 0x6a, 0x5e, 0x4a, 0x66, 0x5e, 0x7a, 0x50,
	0x6a, 0x72, 0x7e, 0x51, 0x4a, 0xb1, 0x84, 0xb8, 0x02, 0x33, 0xe9, 0xbe, 0x44, 0x33, 0xc4, 0x8a,
	0xe5, 0xc5, 0x02, 0x79, 0x06, 0x27, 0x93, 0x0b, 0x0f, 0xe5, 0x18, 0x6e, 0x3c, 0x94, 0x63, 0xf8,
	0xf0, 0x50, 0x8e, 0xb1, 0xe1, 0x91, 0x1c, 0xe3, 0x8a, 0x47, 0x72, 0x8c, 0x27, 0x1e, 0xc9, 0x31,
	0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8b, 0x47, 0x72, 0x0c, 0x1f, 0x1e, 0xc9,
	0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x49, 0x6c,
	0xe0, 0x88, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x6a, 0x43, 0x90, 0x22, 0x02, 0x00,
	0x00,
}

func (this *FilamentIndex) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&object.FilamentIndex{")
	s = append(s, "XPolymorph: "+fmt.Sprintf("%#v", this.XPolymorph)+",\n")
	s = append(s, "ObjID: "+fmt.Sprintf("%#v", this.ObjID)+",\n")
	s = append(s, "Lifeline: "+strings.Replace(this.Lifeline.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "LifelineLastUsed: "+fmt.Sprintf("%#v", this.LifelineLastUsed)+",\n")
	s = append(s, "PendingRecords: "+fmt.Sprintf("%#v", this.PendingRecords)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringFilamentindex(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *FilamentIndex) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FilamentIndex) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XPolymorph != 0 {
		dAtA[i] = 0x80
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintFilamentindex(dAtA, i, uint64(m.XPolymorph))
	}
	dAtA[i] = 0xa2
	i++
	dAtA[i] = 0x1
	i++
	i = encodeVarintFilamentindex(dAtA, i, uint64(m.ObjID.Size()))
	n1, err := m.ObjID.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0xaa
	i++
	dAtA[i] = 0x1
	i++
	i = encodeVarintFilamentindex(dAtA, i, uint64(m.Lifeline.Size()))
	n2, err := m.Lifeline.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.LifelineLastUsed != 0 {
		dAtA[i] = 0xb0
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintFilamentindex(dAtA, i, uint64(m.LifelineLastUsed))
	}
	if len(m.PendingRecords) > 0 {
		for _, msg := range m.PendingRecords {
			dAtA[i] = 0xba
			i++
			dAtA[i] = 0x1
			i++
			i = encodeVarintFilamentindex(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintFilamentindex(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *FilamentIndex) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XPolymorph != 0 {
		n += 2 + sovFilamentindex(uint64(m.XPolymorph))
	}
	l = m.ObjID.Size()
	n += 2 + l + sovFilamentindex(uint64(l))
	l = m.Lifeline.Size()
	n += 2 + l + sovFilamentindex(uint64(l))
	if m.LifelineLastUsed != 0 {
		n += 2 + sovFilamentindex(uint64(m.LifelineLastUsed))
	}
	if len(m.PendingRecords) > 0 {
		for _, e := range m.PendingRecords {
			l = e.Size()
			n += 2 + l + sovFilamentindex(uint64(l))
		}
	}
	return n
}

func sovFilamentindex(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFilamentindex(x uint64) (n int) {
	return sovFilamentindex(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *FilamentIndex) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&FilamentIndex{`,
		`XPolymorph:` + fmt.Sprintf("%v", this.XPolymorph) + `,`,
		`ObjID:` + fmt.Sprintf("%v", this.ObjID) + `,`,
		`Lifeline:` + strings.Replace(strings.Replace(this.Lifeline.String(), "Lifeline", "Lifeline", 1), `&`, ``, 1) + `,`,
		`LifelineLastUsed:` + fmt.Sprintf("%v", this.LifelineLastUsed) + `,`,
		`PendingRecords:` + fmt.Sprintf("%v", this.PendingRecords) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringFilamentindex(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *FilamentIndex) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFilamentindex
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
			return fmt.Errorf("proto: FilamentIndex: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FilamentIndex: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field XPolymorph", wireType)
			}
			m.XPolymorph = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilamentindex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.XPolymorph |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilamentindex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFilamentindex
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFilamentindex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lifeline", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilamentindex
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
				return ErrInvalidLengthFilamentindex
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFilamentindex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Lifeline.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 22:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LifelineLastUsed", wireType)
			}
			m.LifelineLastUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilamentindex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LifelineLastUsed |= github_com_insolar_insolar_insolar.PulseNumber(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 23:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingRecords", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilamentindex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFilamentindex
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFilamentindex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_insolar_insolar_insolar.ID
			m.PendingRecords = append(m.PendingRecords, v)
			if err := m.PendingRecords[len(m.PendingRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFilamentindex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFilamentindex
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFilamentindex
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
func skipFilamentindex(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFilamentindex
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
					return 0, ErrIntOverflowFilamentindex
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
					return 0, ErrIntOverflowFilamentindex
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
				return 0, ErrInvalidLengthFilamentindex
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthFilamentindex
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFilamentindex
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
				next, err := skipFilamentindex(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthFilamentindex
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
	ErrInvalidLengthFilamentindex = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFilamentindex   = fmt.Errorf("proto: integer overflow")
)
