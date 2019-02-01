// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol.proto

package Protocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type Message_Type int32

const (
	Message_ACK          Message_Type = 0
	Message_REQUEST      Message_Type = 1
	Message_NOTIFICATION Message_Type = 2
)

var Message_Type_name = map[int32]string{
	0: "ACK",
	1: "REQUEST",
	2: "NOTIFICATION",
}

var Message_Type_value = map[string]int32{
	"ACK":          0,
	"REQUEST":      1,
	"NOTIFICATION": 2,
}

func (x Message_Type) String() string {
	return proto.EnumName(Message_Type_name, int32(x))
}

func (Message_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{0, 0}
}

type Message_CtrlType int32

const (
	Message_NONE             Message_CtrlType = 0
	Message_CREATE_GROUP     Message_CtrlType = 1
	Message_GROUP_ADDMEMBERS Message_CtrlType = 2
	Message_MSG_BACK         Message_CtrlType = 3
	Message_CREATE_SESSION   Message_CtrlType = 4
)

var Message_CtrlType_name = map[int32]string{
	0: "NONE",
	1: "CREATE_GROUP",
	2: "GROUP_ADDMEMBERS",
	3: "MSG_BACK",
	4: "CREATE_SESSION",
}

var Message_CtrlType_value = map[string]int32{
	"NONE":             0,
	"CREATE_GROUP":     1,
	"GROUP_ADDMEMBERS": 2,
	"MSG_BACK":         3,
	"CREATE_SESSION":   4,
}

func (x Message_CtrlType) String() string {
	return proto.EnumName(Message_CtrlType_name, int32(x))
}

func (Message_CtrlType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{0, 1}
}

type Message_ContentType int32

const (
	Message_TEXT Message_ContentType = 0
	Message_IMG  Message_ContentType = 1
	Message_FILE Message_ContentType = 2
)

var Message_ContentType_name = map[int32]string{
	0: "TEXT",
	1: "IMG",
	2: "FILE",
}

var Message_ContentType_value = map[string]int32{
	"TEXT": 0,
	"IMG":  1,
	"FILE": 2,
}

func (x Message_ContentType) String() string {
	return proto.EnumName(Message_ContentType_name, int32(x))
}

func (Message_ContentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{0, 2}
}

type Message struct {
	Type                 Message_Type        `protobuf:"varint,1,opt,name=type,proto3,enum=Protocol.Message_Type" json:"type,omitempty"`
	Cmd                  Message_CtrlType    `protobuf:"varint,2,opt,name=cmd,proto3,enum=Protocol.Message_CtrlType" json:"cmd,omitempty"`
	From                 int64               `protobuf:"varint,3,opt,name=from,proto3" json:"from,omitempty"`
	To                   int64               `protobuf:"varint,4,opt,name=to,proto3" json:"to,omitempty"`
	Isgroup              bool                `protobuf:"varint,5,opt,name=isgroup,proto3" json:"isgroup,omitempty"`
	Content              string              `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	ContentType          Message_ContentType `protobuf:"varint,7,opt,name=content_type,json=contentType,proto3,enum=Protocol.Message_ContentType" json:"content_type,omitempty"`
	Msgid                int64               `protobuf:"varint,8,opt,name=msgid,proto3" json:"msgid,omitempty"`
	Time                 int64               `protobuf:"varint,9,opt,name=time,proto3" json:"time,omitempty"`
	Userlist             []int64             `protobuf:"varint,10,rep,packed,name=userlist,proto3" json:"userlist,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetType() Message_Type {
	if m != nil {
		return m.Type
	}
	return Message_ACK
}

func (m *Message) GetCmd() Message_CtrlType {
	if m != nil {
		return m.Cmd
	}
	return Message_NONE
}

func (m *Message) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *Message) GetTo() int64 {
	if m != nil {
		return m.To
	}
	return 0
}

func (m *Message) GetIsgroup() bool {
	if m != nil {
		return m.Isgroup
	}
	return false
}

func (m *Message) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Message) GetContentType() Message_ContentType {
	if m != nil {
		return m.ContentType
	}
	return Message_TEXT
}

func (m *Message) GetMsgid() int64 {
	if m != nil {
		return m.Msgid
	}
	return 0
}

func (m *Message) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Message) GetUserlist() []int64 {
	if m != nil {
		return m.Userlist
	}
	return nil
}

func init() {
	proto.RegisterEnum("Protocol.Message_Type", Message_Type_name, Message_Type_value)
	proto.RegisterEnum("Protocol.Message_CtrlType", Message_CtrlType_name, Message_CtrlType_value)
	proto.RegisterEnum("Protocol.Message_ContentType", Message_ContentType_name, Message_ContentType_value)
	proto.RegisterType((*Message)(nil), "Protocol.Message")
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptor_2bc2336598a3f7e0) }

var fileDescriptor_2bc2336598a3f7e0 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xd1, 0x6a, 0x9c, 0x40,
	0x14, 0x86, 0x77, 0x1c, 0x13, 0xcd, 0xd9, 0x65, 0x19, 0x0e, 0xa1, 0x0c, 0x81, 0x82, 0xec, 0x95,
	0x84, 0xe2, 0x45, 0xfb, 0x02, 0x35, 0x66, 0xb2, 0x48, 0xab, 0xa6, 0xa3, 0x81, 0x5e, 0x55, 0x5a,
	0x63, 0x17, 0x61, 0xcd, 0x88, 0x4e, 0x2e, 0xf2, 0xf4, 0x2d, 0x33, 0xab, 0x6d, 0xa1, 0xb9, 0xfb,
	0xff, 0xf1, 0xf3, 0x9c, 0x4f, 0x19, 0xd8, 0x0e, 0xa3, 0xd2, 0xaa, 0x51, 0xc7, 0xc8, 0x06, 0xf4,
	0xef, 0xe7, 0xbe, 0xfb, 0x45, 0xc1, 0xcb, 0xda, 0x69, 0xfa, 0x7e, 0x68, 0xf1, 0x1a, 0x5c, 0xfd,
	0x32, 0xb4, 0x9c, 0x04, 0x24, 0xdc, 0xbe, 0x7f, 0x13, 0x2d, 0x50, 0x34, 0x03, 0x51, 0xf5, 0x32,
	0xb4, 0xd2, 0x32, 0xf8, 0x0e, 0x68, 0xd3, 0x3f, 0x72, 0xc7, 0xa2, 0x57, 0xff, 0xa3, 0x89, 0x1e,
	0x8f, 0x16, 0x37, 0x18, 0x22, 0xb8, 0x3f, 0x47, 0xd5, 0x73, 0x1a, 0x90, 0x90, 0x4a, 0x9b, 0x71,
	0x0b, 0x8e, 0x56, 0xdc, 0xb5, 0x27, 0x8e, 0x56, 0xc8, 0xc1, 0xeb, 0xa6, 0xc3, 0xa8, 0x9e, 0x07,
	0x7e, 0x16, 0x90, 0xd0, 0x97, 0x4b, 0x35, 0x4f, 0x1a, 0xf5, 0xa4, 0xdb, 0x27, 0xcd, 0xcf, 0x03,
	0x12, 0x5e, 0xc8, 0xa5, 0xe2, 0x47, 0xd8, 0xcc, 0xb1, 0xb6, 0xe6, 0x9e, 0xd5, 0x79, 0xfb, 0x8a,
	0xce, 0x89, 0xb2, 0x46, 0xeb, 0xe6, 0x6f, 0xc1, 0x4b, 0x38, 0xeb, 0xa7, 0x43, 0xf7, 0xc8, 0x7d,
	0x2b, 0x72, 0x2a, 0xc6, 0x57, 0x77, 0x7d, 0xcb, 0x2f, 0x4e, 0xbe, 0x26, 0xe3, 0x15, 0xf8, 0xcf,
	0x53, 0x3b, 0x1e, 0xbb, 0x49, 0x73, 0x08, 0x68, 0x48, 0xe5, 0x9f, 0xbe, 0x8b, 0xc0, 0xb5, 0xd3,
	0x3c, 0xa0, 0x71, 0xf2, 0x89, 0xad, 0x70, 0x0d, 0x9e, 0x14, 0x5f, 0x1e, 0x44, 0x59, 0x31, 0x82,
	0x0c, 0x36, 0x79, 0x51, 0xa5, 0x77, 0x69, 0x12, 0x57, 0x69, 0x91, 0x33, 0x67, 0xf7, 0x0d, 0xfc,
	0xe5, 0x07, 0xa1, 0x0f, 0x6e, 0x5e, 0xe4, 0x82, 0xad, 0x0c, 0x97, 0x48, 0x11, 0x57, 0xa2, 0xde,
	0xcb, 0xe2, 0xe1, 0x9e, 0x11, 0xbc, 0x04, 0x66, 0x63, 0x1d, 0xdf, 0xde, 0x66, 0x22, 0xbb, 0x11,
	0xb2, 0x64, 0x0e, 0x6e, 0xc0, 0xcf, 0xca, 0x7d, 0x7d, 0x63, 0x56, 0x51, 0x44, 0xd8, 0xce, 0x6f,
	0x95, 0xa2, 0x2c, 0xcd, 0x7c, 0x77, 0x77, 0x0d, 0xeb, 0x7f, 0xbe, 0xd8, 0xac, 0xa8, 0xc4, 0xd7,
	0x8a, 0xad, 0x8c, 0x60, 0x9a, 0xed, 0x19, 0x31, 0x47, 0x77, 0xe9, 0x67, 0xc1, 0x9c, 0x1f, 0xe7,
	0xf6, 0x4a, 0x7c, 0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x97, 0xa9, 0x38, 0x0e, 0x24, 0x02, 0x00,
	0x00,
}
