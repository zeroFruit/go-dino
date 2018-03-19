// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protodata.proto

/*
Package dinoproto2 is a generated protocol buffer package.

It is generated from these files:
	protodata.proto

It has these top-level messages:
	Animal
*/
package dinoproto2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Animal struct {
	Id               *int32  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	AnimalType       *string `protobuf:"bytes,2,req,name=animal_type" json:"animal_type,omitempty"`
	Nickname         *string `protobuf:"bytes,3,opt,name=nickname" json:"nickname,omitempty"`
	Zone             *int32  `protobuf:"varint,4,req,name=zone" json:"zone,omitempty"`
	Age              *int32  `protobuf:"varint,5,opt,name=age" json:"age,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Animal) Reset()                    { *m = Animal{} }
func (m *Animal) String() string            { return proto.CompactTextString(m) }
func (*Animal) ProtoMessage()               {}
func (*Animal) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Animal) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Animal) GetAnimalType() string {
	if m != nil && m.AnimalType != nil {
		return *m.AnimalType
	}
	return ""
}

func (m *Animal) GetNickname() string {
	if m != nil && m.Nickname != nil {
		return *m.Nickname
	}
	return ""
}

func (m *Animal) GetZone() int32 {
	if m != nil && m.Zone != nil {
		return *m.Zone
	}
	return 0
}

func (m *Animal) GetAge() int32 {
	if m != nil && m.Age != nil {
		return *m.Age
	}
	return 0
}

func init() {
	proto.RegisterType((*Animal)(nil), "dinoproto2.animal")
}

func init() { proto.RegisterFile("protodata.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x49, 0x2c, 0x49, 0xd4, 0x03, 0xb3, 0x84, 0xb8, 0x52, 0x32, 0xf3, 0xf2, 0xc1, 0x4c,
	0x23, 0xa5, 0x30, 0x2e, 0xb6, 0xc4, 0xbc, 0xcc, 0xdc, 0xc4, 0x1c, 0x21, 0x2e, 0x2e, 0xa6, 0xcc,
	0x14, 0x09, 0x46, 0x05, 0x26, 0x0d, 0x56, 0x21, 0x61, 0x2e, 0x6e, 0x88, 0x68, 0x7c, 0x49, 0x65,
	0x41, 0xaa, 0x04, 0x93, 0x02, 0x93, 0x06, 0xa7, 0x90, 0x00, 0x17, 0x47, 0x5e, 0x66, 0x72, 0x76,
	0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0xa7, 0x10, 0x0f, 0x17, 0x4b, 0x55, 0x7e,
	0x5e, 0xaa, 0x04, 0x0b, 0x58, 0x13, 0x37, 0x17, 0x73, 0x62, 0x7a, 0xaa, 0x04, 0xab, 0x02, 0xa3,
	0x06, 0x2b, 0x20, 0x00, 0x00, 0xff, 0xff, 0x86, 0xbe, 0x61, 0xf8, 0x75, 0x00, 0x00, 0x00,
}