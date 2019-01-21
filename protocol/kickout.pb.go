// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kickout.proto

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

type KickoutRequest_Reasion int32

const (
	KickoutRequest_OTHER_LOGIN KickoutRequest_Reasion = 0
)

var KickoutRequest_Reasion_name = map[int32]string{
	0: "OTHER_LOGIN",
}

var KickoutRequest_Reasion_value = map[string]int32{
	"OTHER_LOGIN": 0,
}

func (x KickoutRequest_Reasion) String() string {
	return proto.EnumName(KickoutRequest_Reasion_name, int32(x))
}

func (KickoutRequest_Reasion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e21ae08ae73a01df, []int{0, 0}
}

type KickoutRequest struct {
	Reasion              KickoutRequest_Reasion `protobuf:"varint,1,opt,name=reasion,proto3,enum=Protocol.KickoutRequest_Reasion" json:"reasion,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *KickoutRequest) Reset()         { *m = KickoutRequest{} }
func (m *KickoutRequest) String() string { return proto.CompactTextString(m) }
func (*KickoutRequest) ProtoMessage()    {}
func (*KickoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e21ae08ae73a01df, []int{0}
}

func (m *KickoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KickoutRequest.Unmarshal(m, b)
}
func (m *KickoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KickoutRequest.Marshal(b, m, deterministic)
}
func (m *KickoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KickoutRequest.Merge(m, src)
}
func (m *KickoutRequest) XXX_Size() int {
	return xxx_messageInfo_KickoutRequest.Size(m)
}
func (m *KickoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KickoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KickoutRequest proto.InternalMessageInfo

func (m *KickoutRequest) GetReasion() KickoutRequest_Reasion {
	if m != nil {
		return m.Reasion
	}
	return KickoutRequest_OTHER_LOGIN
}

func init() {
	proto.RegisterEnum("Protocol.KickoutRequest_Reasion", KickoutRequest_Reasion_name, KickoutRequest_Reasion_value)
	proto.RegisterType((*KickoutRequest)(nil), "Protocol.KickoutRequest")
}

func init() { proto.RegisterFile("kickout.proto", fileDescriptor_e21ae08ae73a01df) }

var fileDescriptor_e21ae08ae73a01df = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0xce, 0x4c, 0xce,
	0xce, 0x2f, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x08, 0x00, 0x51, 0xc9, 0xf9,
	0x39, 0x4a, 0x19, 0x5c, 0x7c, 0xde, 0x10, 0xa9, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21,
	0x2b, 0x2e, 0xf6, 0xa2, 0xd4, 0xc4, 0xe2, 0xcc, 0xfc, 0x3c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x3e,
	0x23, 0x05, 0x3d, 0x98, 0x6a, 0x3d, 0x54, 0xa5, 0x7a, 0x41, 0x10, 0x75, 0x41, 0x30, 0x0d, 0x4a,
	0x52, 0x5c, 0xec, 0x50, 0x31, 0x21, 0x7e, 0x2e, 0x6e, 0xff, 0x10, 0x0f, 0xd7, 0xa0, 0x78, 0x1f,
	0x7f, 0x77, 0x4f, 0x3f, 0x01, 0x86, 0x24, 0x36, 0xb0, 0xd5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xfe, 0xb7, 0x8f, 0xda, 0x8b, 0x00, 0x00, 0x00,
}
