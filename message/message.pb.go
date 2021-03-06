// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/message/message.proto

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	github.com/google/cloudprober/message/message.proto

It has these top-level messages:
	Constants
	DataNode
	Msg
*/
package message

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

type DataNode_Type int32

const (
	DataNode_UNKNOWN DataNode_Type = 0
	DataNode_CLIENT  DataNode_Type = 1
	DataNode_SERVER  DataNode_Type = 2
)

var DataNode_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "CLIENT",
	2: "SERVER",
}
var DataNode_Type_value = map[string]int32{
	"UNKNOWN": 0,
	"CLIENT":  1,
	"SERVER":  2,
}

func (x DataNode_Type) Enum() *DataNode_Type {
	p := new(DataNode_Type)
	*p = x
	return p
}
func (x DataNode_Type) String() string {
	return proto.EnumName(DataNode_Type_name, int32(x))
}
func (x *DataNode_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DataNode_Type_value, data, "DataNode_Type")
	if err != nil {
		return err
	}
	*x = DataNode_Type(value)
	return nil
}
func (DataNode_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

// Constants defines constants with default values.
type Constants struct {
	Magic            *uint64 `protobuf:"varint,1,opt,name=magic,def=257787339638762" json:"magic,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Constants) Reset()                    { *m = Constants{} }
func (m *Constants) String() string            { return proto.CompactTextString(m) }
func (*Constants) ProtoMessage()               {}
func (*Constants) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_Constants_Magic uint64 = 257787339638762

func (m *Constants) GetMagic() uint64 {
	if m != nil && m.Magic != nil {
		return *m.Magic
	}
	return Default_Constants_Magic
}

// Datanode is something that see's a message AND can modify it.
type DataNode struct {
	Type *DataNode_Type `protobuf:"varint,1,opt,name=type,enum=message.DataNode_Type,def=1" json:"type,omitempty"`
	Name *string        `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// 8 bytes of timestamp in pcap-friendly network byte order.
	TimestampUsec    []byte `protobuf:"bytes,3,opt,name=timestamp_usec,json=timestampUsec" json:"timestamp_usec,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DataNode) Reset()                    { *m = DataNode{} }
func (m *DataNode) String() string            { return proto.CompactTextString(m) }
func (*DataNode) ProtoMessage()               {}
func (*DataNode) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

const Default_DataNode_Type DataNode_Type = DataNode_CLIENT

func (m *DataNode) GetType() DataNode_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_DataNode_Type
}

func (m *DataNode) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *DataNode) GetTimestampUsec() []byte {
	if m != nil {
		return m.TimestampUsec
	}
	return nil
}

// Msg is a message sent over the network.
// magic, seq, src and dst are required fields.
type Msg struct {
	Magic *uint64 `protobuf:"fixed64,1,opt,name=magic" json:"magic,omitempty"`
	// 8 bytes of sequence in pcap-friendly network byte order.
	Seq []byte `protobuf:"bytes,2,opt,name=seq" json:"seq,omitempty"`
	// Datanodes seen by this message.
	Src              *DataNode   `protobuf:"bytes,3,opt,name=src" json:"src,omitempty"`
	Dst              *DataNode   `protobuf:"bytes,4,opt,name=dst" json:"dst,omitempty"`
	Nodes            []*DataNode `protobuf:"bytes,5,rep,name=nodes" json:"nodes,omitempty"`
	Pad              []byte      `protobuf:"bytes,99,opt,name=pad" json:"pad,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Msg) Reset()                    { *m = Msg{} }
func (m *Msg) String() string            { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()               {}
func (*Msg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Msg) GetMagic() uint64 {
	if m != nil && m.Magic != nil {
		return *m.Magic
	}
	return 0
}

func (m *Msg) GetSeq() []byte {
	if m != nil {
		return m.Seq
	}
	return nil
}

func (m *Msg) GetSrc() *DataNode {
	if m != nil {
		return m.Src
	}
	return nil
}

func (m *Msg) GetDst() *DataNode {
	if m != nil {
		return m.Dst
	}
	return nil
}

func (m *Msg) GetNodes() []*DataNode {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *Msg) GetPad() []byte {
	if m != nil {
		return m.Pad
	}
	return nil
}

func init() {
	proto.RegisterType((*Constants)(nil), "message.Constants")
	proto.RegisterType((*DataNode)(nil), "message.DataNode")
	proto.RegisterType((*Msg)(nil), "message.Msg")
	proto.RegisterEnum("message.DataNode_Type", DataNode_Type_name, DataNode_Type_value)
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/message/message.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0xd1, 0x4e, 0xc2, 0x30,
	0x18, 0x85, 0x2d, 0x1b, 0x20, 0x3f, 0x88, 0xb3, 0x31, 0x66, 0x97, 0xcb, 0x0c, 0x91, 0xc4, 0x64,
	0x24, 0x5b, 0x64, 0xc8, 0x2d, 0x72, 0x61, 0xd4, 0x99, 0x54, 0xd0, 0x4b, 0x53, 0xd6, 0x66, 0x92,
	0xb0, 0xb5, 0xae, 0xe5, 0x82, 0x17, 0xf2, 0x01, 0x7c, 0x42, 0xb3, 0x8d, 0x19, 0x13, 0xf5, 0xaa,
	0xa7, 0x3d, 0x5f, 0xff, 0xff, 0x1c, 0x08, 0x92, 0xb5, 0x7e, 0xdb, 0xae, 0xbc, 0x58, 0xa4, 0xa3,
	0x44, 0x88, 0x64, 0xc3, 0x47, 0xf1, 0x46, 0x6c, 0x99, 0xcc, 0xc5, 0x8a, 0xe7, 0xa3, 0x94, 0x2b,
	0x45, 0x13, 0x5e, 0x9f, 0x9e, 0xcc, 0x85, 0x16, 0xb8, 0xbd, 0xbf, 0xba, 0x3e, 0x74, 0x66, 0x22,
	0x53, 0x9a, 0x66, 0x5a, 0xe1, 0x01, 0x34, 0x53, 0x9a, 0xac, 0x63, 0x1b, 0x39, 0x68, 0x68, 0x4e,
	0x8f, 0xfd, 0xab, 0x30, 0x9c, 0x84, 0x41, 0x70, 0x3d, 0x0e, 0x26, 0xe1, 0xd8, 0x27, 0x95, 0xeb,
	0x7e, 0x20, 0x38, 0xbc, 0xa1, 0x9a, 0x46, 0x82, 0x71, 0xec, 0x83, 0xa9, 0x77, 0x92, 0x97, 0x5f,
	0xfa, 0xfe, 0x99, 0x57, 0xef, 0xa9, 0x01, 0x6f, 0xb1, 0x93, 0x7c, 0xda, 0x9a, 0xdd, 0xdf, 0xce,
	0xa3, 0x05, 0x29, 0x59, 0x8c, 0xc1, 0xcc, 0x68, 0xca, 0xed, 0x86, 0x83, 0x86, 0x1d, 0x52, 0x6a,
	0x3c, 0x80, 0xbe, 0x5e, 0xa7, 0x5c, 0x69, 0x9a, 0xca, 0xd7, 0xad, 0xe2, 0xb1, 0x6d, 0x38, 0x68,
	0xd8, 0x23, 0x47, 0xdf, 0xaf, 0x4b, 0xc5, 0x63, 0xf7, 0x12, 0xcc, 0x62, 0x20, 0xee, 0x42, 0x7b,
	0x19, 0xdd, 0x45, 0x8f, 0x2f, 0x91, 0x75, 0x80, 0x01, 0xf6, 0xf3, 0x2d, 0x54, 0xe8, 0xa7, 0x39,
	0x79, 0x9e, 0x13, 0xab, 0xe1, 0x7e, 0x22, 0x30, 0x1e, 0x54, 0x82, 0x4f, 0x7f, 0xf6, 0x6a, 0xed,
	0x6b, 0x60, 0x0b, 0x0c, 0xc5, 0xdf, 0xcb, 0x10, 0x3d, 0x52, 0x48, 0x7c, 0x0e, 0x86, 0xca, 0xab,
	0xc5, 0x5d, 0xff, 0xe4, 0x57, 0x15, 0x52, 0xb8, 0x05, 0xc4, 0x94, 0xb6, 0xcd, 0x7f, 0x21, 0xa6,
	0x34, 0xbe, 0x80, 0x66, 0x26, 0x18, 0x57, 0x76, 0xd3, 0x31, 0xfe, 0xc6, 0x2a, 0xbf, 0x08, 0x21,
	0x29, 0xb3, 0xe3, 0x2a, 0x84, 0xa4, 0xec, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x2a, 0x4b, 0xfe,
	0xd0, 0x01, 0x00, 0x00,
}
