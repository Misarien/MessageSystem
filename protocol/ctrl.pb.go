// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ctrl.proto

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

type CtrlSendRequest_CtrlSendType int32

const (
	CtrlSendRequest_CREATE_GROUP     CtrlSendRequest_CtrlSendType = 0
	CtrlSendRequest_GROUP_ADDMEMBERS CtrlSendRequest_CtrlSendType = 1
	CtrlSendRequest_GROUP_EXIT       CtrlSendRequest_CtrlSendType = 3
	CtrlSendRequest_GROUP_DELETE     CtrlSendRequest_CtrlSendType = 4
	CtrlSendRequest_MSG_BACK         CtrlSendRequest_CtrlSendType = 5
)

var CtrlSendRequest_CtrlSendType_name = map[int32]string{
	0: "CREATE_GROUP",
	1: "GROUP_ADDMEMBERS",
	3: "GROUP_EXIT",
	4: "GROUP_DELETE",
	5: "MSG_BACK",
}

var CtrlSendRequest_CtrlSendType_value = map[string]int32{
	"CREATE_GROUP":     0,
	"GROUP_ADDMEMBERS": 1,
	"GROUP_EXIT":       3,
	"GROUP_DELETE":     4,
	"MSG_BACK":         5,
}

func (x CtrlSendRequest_CtrlSendType) String() string {
	return proto.EnumName(CtrlSendRequest_CtrlSendType_name, int32(x))
}

func (CtrlSendRequest_CtrlSendType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{0, 0}
}

type CtrPushRequest_CtrlPushType int32

const (
	CtrPushRequest_CREATE_GROUP     CtrPushRequest_CtrlPushType = 0
	CtrPushRequest_GROUP_ADDMEMBERS CtrPushRequest_CtrlPushType = 1
	CtrPushRequest_GROUP_EXIT       CtrPushRequest_CtrlPushType = 3
	CtrPushRequest_GROUP_DELETE     CtrPushRequest_CtrlPushType = 4
)

var CtrPushRequest_CtrlPushType_name = map[int32]string{
	0: "CREATE_GROUP",
	1: "GROUP_ADDMEMBERS",
	3: "GROUP_EXIT",
	4: "GROUP_DELETE",
}

var CtrPushRequest_CtrlPushType_value = map[string]int32{
	"CREATE_GROUP":     0,
	"GROUP_ADDMEMBERS": 1,
	"GROUP_EXIT":       3,
	"GROUP_DELETE":     4,
}

func (x CtrPushRequest_CtrlPushType) String() string {
	return proto.EnumName(CtrPushRequest_CtrlPushType_name, int32(x))
}

func (CtrPushRequest_CtrlPushType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{2, 0}
}

type CtrlSendRequest struct {
	From                 int64                        `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	Type                 CtrlSendRequest_CtrlSendType `protobuf:"varint,2,opt,name=type,proto3,enum=Protocol.CtrlSendRequest_CtrlSendType" json:"type,omitempty"`
	To                   int64                        `protobuf:"varint,3,opt,name=to,proto3" json:"to,omitempty"`
	Msgid                int64                        `protobuf:"varint,4,opt,name=msgid,proto3" json:"msgid,omitempty"`
	Content              string                       `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *CtrlSendRequest) Reset()         { *m = CtrlSendRequest{} }
func (m *CtrlSendRequest) String() string { return proto.CompactTextString(m) }
func (*CtrlSendRequest) ProtoMessage()    {}
func (*CtrlSendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{0}
}

func (m *CtrlSendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CtrlSendRequest.Unmarshal(m, b)
}
func (m *CtrlSendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CtrlSendRequest.Marshal(b, m, deterministic)
}
func (m *CtrlSendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CtrlSendRequest.Merge(m, src)
}
func (m *CtrlSendRequest) XXX_Size() int {
	return xxx_messageInfo_CtrlSendRequest.Size(m)
}
func (m *CtrlSendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CtrlSendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CtrlSendRequest proto.InternalMessageInfo

func (m *CtrlSendRequest) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *CtrlSendRequest) GetType() CtrlSendRequest_CtrlSendType {
	if m != nil {
		return m.Type
	}
	return CtrlSendRequest_CREATE_GROUP
}

func (m *CtrlSendRequest) GetTo() int64 {
	if m != nil {
		return m.To
	}
	return 0
}

func (m *CtrlSendRequest) GetMsgid() int64 {
	if m != nil {
		return m.Msgid
	}
	return 0
}

func (m *CtrlSendRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type CtrlSendResponse struct {
	Msgid                int64    `protobuf:"varint,1,opt,name=msgid,proto3" json:"msgid,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CtrlSendResponse) Reset()         { *m = CtrlSendResponse{} }
func (m *CtrlSendResponse) String() string { return proto.CompactTextString(m) }
func (*CtrlSendResponse) ProtoMessage()    {}
func (*CtrlSendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{1}
}

func (m *CtrlSendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CtrlSendResponse.Unmarshal(m, b)
}
func (m *CtrlSendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CtrlSendResponse.Marshal(b, m, deterministic)
}
func (m *CtrlSendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CtrlSendResponse.Merge(m, src)
}
func (m *CtrlSendResponse) XXX_Size() int {
	return xxx_messageInfo_CtrlSendResponse.Size(m)
}
func (m *CtrlSendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CtrlSendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CtrlSendResponse proto.InternalMessageInfo

func (m *CtrlSendResponse) GetMsgid() int64 {
	if m != nil {
		return m.Msgid
	}
	return 0
}

func (m *CtrlSendResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *CtrlSendResponse) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

////////////
type CtrPushRequest struct {
	From                 int64                       `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	Type                 CtrPushRequest_CtrlPushType `protobuf:"varint,2,opt,name=type,proto3,enum=Protocol.CtrPushRequest_CtrlPushType" json:"type,omitempty"`
	To                   int64                       `protobuf:"varint,3,opt,name=to,proto3" json:"to,omitempty"`
	Msgid                int64                       `protobuf:"varint,4,opt,name=msgid,proto3" json:"msgid,omitempty"`
	Content              string                      `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *CtrPushRequest) Reset()         { *m = CtrPushRequest{} }
func (m *CtrPushRequest) String() string { return proto.CompactTextString(m) }
func (*CtrPushRequest) ProtoMessage()    {}
func (*CtrPushRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{2}
}

func (m *CtrPushRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CtrPushRequest.Unmarshal(m, b)
}
func (m *CtrPushRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CtrPushRequest.Marshal(b, m, deterministic)
}
func (m *CtrPushRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CtrPushRequest.Merge(m, src)
}
func (m *CtrPushRequest) XXX_Size() int {
	return xxx_messageInfo_CtrPushRequest.Size(m)
}
func (m *CtrPushRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CtrPushRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CtrPushRequest proto.InternalMessageInfo

func (m *CtrPushRequest) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *CtrPushRequest) GetType() CtrPushRequest_CtrlPushType {
	if m != nil {
		return m.Type
	}
	return CtrPushRequest_CREATE_GROUP
}

func (m *CtrPushRequest) GetTo() int64 {
	if m != nil {
		return m.To
	}
	return 0
}

func (m *CtrPushRequest) GetMsgid() int64 {
	if m != nil {
		return m.Msgid
	}
	return 0
}

func (m *CtrPushRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type CtrlPushResponse struct {
	Msgid                int64    `protobuf:"varint,1,opt,name=msgid,proto3" json:"msgid,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CtrlPushResponse) Reset()         { *m = CtrlPushResponse{} }
func (m *CtrlPushResponse) String() string { return proto.CompactTextString(m) }
func (*CtrlPushResponse) ProtoMessage()    {}
func (*CtrlPushResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{3}
}

func (m *CtrlPushResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CtrlPushResponse.Unmarshal(m, b)
}
func (m *CtrlPushResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CtrlPushResponse.Marshal(b, m, deterministic)
}
func (m *CtrlPushResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CtrlPushResponse.Merge(m, src)
}
func (m *CtrlPushResponse) XXX_Size() int {
	return xxx_messageInfo_CtrlPushResponse.Size(m)
}
func (m *CtrlPushResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CtrlPushResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CtrlPushResponse proto.InternalMessageInfo

func (m *CtrlPushResponse) GetMsgid() int64 {
	if m != nil {
		return m.Msgid
	}
	return 0
}

func (m *CtrlPushResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterEnum("Protocol.CtrlSendRequest_CtrlSendType", CtrlSendRequest_CtrlSendType_name, CtrlSendRequest_CtrlSendType_value)
	proto.RegisterEnum("Protocol.CtrPushRequest_CtrlPushType", CtrPushRequest_CtrlPushType_name, CtrPushRequest_CtrlPushType_value)
	proto.RegisterType((*CtrlSendRequest)(nil), "Protocol.CtrlSendRequest")
	proto.RegisterType((*CtrlSendResponse)(nil), "Protocol.CtrlSendResponse")
	proto.RegisterType((*CtrPushRequest)(nil), "Protocol.CtrPushRequest")
	proto.RegisterType((*CtrlPushResponse)(nil), "Protocol.CtrlPushResponse")
}

func init() { proto.RegisterFile("ctrl.proto", fileDescriptor_a0572e205f89e843) }

var fileDescriptor_a0572e205f89e843 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xb1, 0x4f, 0x02, 0x31,
	0x14, 0xc6, 0xed, 0xdd, 0xa1, 0xf0, 0x42, 0xce, 0xa6, 0x61, 0xe8, 0xe0, 0x40, 0x2e, 0xd1, 0x30,
	0x31, 0xe8, 0xa4, 0x1b, 0x1c, 0x95, 0x18, 0x25, 0x92, 0x72, 0x26, 0x6c, 0xa8, 0x50, 0x94, 0x84,
	0xbb, 0x9e, 0xd7, 0xc7, 0xc0, 0x9f, 0xe0, 0x9f, 0xec, 0x66, 0x68, 0x83, 0x1c, 0x0e, 0x0e, 0x86,
	0xa9, 0xef, 0x7b, 0xc9, 0xf7, 0xe5, 0x7d, 0xbf, 0x14, 0x60, 0x8a, 0xc5, 0xb2, 0x9d, 0x17, 0x1a,
	0x35, 0xab, 0x0e, 0x37, 0xcf, 0x54, 0x2f, 0xa3, 0x4f, 0x0f, 0x4e, 0x63, 0x2c, 0x96, 0x23, 0x95,
	0xcd, 0xa4, 0xfa, 0x58, 0x29, 0x83, 0x8c, 0x41, 0x30, 0x2f, 0x74, 0xca, 0x49, 0x93, 0xb4, 0x7c,
	0x69, 0x67, 0x76, 0x03, 0x01, 0xae, 0x73, 0xc5, 0xbd, 0x26, 0x69, 0x85, 0x97, 0x17, 0xed, 0x6d,
	0x40, 0xfb, 0x97, 0xf9, 0x47, 0x27, 0xeb, 0x5c, 0x49, 0xeb, 0x61, 0x21, 0x78, 0xa8, 0xb9, 0x6f,
	0xd3, 0x3c, 0xd4, 0xac, 0x01, 0x95, 0xd4, 0xbc, 0x2d, 0x66, 0x3c, 0xb0, 0x2b, 0x27, 0x18, 0x87,
	0x93, 0xa9, 0xce, 0x50, 0x65, 0xc8, 0x2b, 0x4d, 0xd2, 0xaa, 0xc9, 0xad, 0x8c, 0xe6, 0x50, 0x2f,
	0xa7, 0x32, 0x0a, 0xf5, 0x58, 0x8a, 0x4e, 0x22, 0x26, 0x7d, 0xf9, 0xf8, 0x34, 0xa4, 0x47, 0xac,
	0x01, 0xd4, 0x8e, 0x93, 0x4e, 0xaf, 0x37, 0x10, 0x83, 0xae, 0x90, 0x23, 0x4a, 0x58, 0x08, 0xe0,
	0xb6, 0x62, 0x7c, 0x97, 0x50, 0x7f, 0xe3, 0x73, 0xba, 0x27, 0x1e, 0x44, 0x22, 0x68, 0xc0, 0xea,
	0x50, 0x1d, 0x8c, 0xfa, 0x93, 0x6e, 0x27, 0xbe, 0xa7, 0x95, 0xe8, 0x19, 0xe8, 0xae, 0x8d, 0xc9,
	0x75, 0x66, 0xd4, 0xee, 0x56, 0x52, 0xbe, 0xf5, 0x0c, 0x6a, 0xb8, 0x48, 0x95, 0xc1, 0x97, 0x34,
	0xb7, 0x48, 0x7c, 0xb9, 0x5b, 0x94, 0x9b, 0xf8, 0xfb, 0x4d, 0xbe, 0x08, 0x84, 0x31, 0x16, 0xc3,
	0x95, 0x79, 0xff, 0x0b, 0xf6, 0xf5, 0x1e, 0xec, 0xf3, 0x3d, 0xd8, 0x25, 0xaf, 0x65, 0xbd, 0xd1,
	0x07, 0x64, 0x3d, 0x76, 0xac, 0xb7, 0xa9, 0x87, 0x63, 0x1d, 0xdd, 0x3a, 0xba, 0xee, 0xfe, 0xff,
	0xd3, 0x7d, 0x3d, 0xb6, 0x5f, 0xf8, 0xea, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x37, 0xe8, 0x2f,
	0xd0, 0x02, 0x00, 0x00,
}
