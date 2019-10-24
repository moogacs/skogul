// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gnmi_ext.proto

package gen

/*
Package gnmi_ext defines a set of extensions messages which can be optionally
included with the request and response messages of gNMI RPCs. A set of
well-known extensions are defined within this file, along with a registry for
extensions defined outside of this package.
*/

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

// RegisteredExtension is an enumeration acting as a registry for extensions
// defined by external sources.
type ExtensionID int32

const (
	ExtensionID_EID_UNSET ExtensionID = 0
	// Juniper Telemetry header
	ExtensionID_EID_JUNIPER_TELEMETRY_HEADER ExtensionID = 1
	// An experimental extension that may be used during prototyping of a new
	// extension.
	ExtensionID_EID_EXPERIMENTAL ExtensionID = 999
)

var ExtensionID_name = map[int32]string{
	0:   "EID_UNSET",
	1:   "EID_JUNIPER_TELEMETRY_HEADER",
	999: "EID_EXPERIMENTAL",
}
var ExtensionID_value = map[string]int32{
	"EID_UNSET":                    0,
	"EID_JUNIPER_TELEMETRY_HEADER": 1,
	"EID_EXPERIMENTAL":             999,
}

func (x ExtensionID) String() string {
	return proto.EnumName(ExtensionID_name, int32(x))
}
func (ExtensionID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_gnmi_ext_0110643c2535c262, []int{0}
}

// The Extension message contains a single gNMI extension.
type Extension struct {
	// Types that are valid to be assigned to Ext:
	//	*Extension_RegisteredExt
	//	*Extension_MasterArbitration
	Ext                  isExtension_Ext `protobuf_oneof:"ext"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Extension) Reset()         { *m = Extension{} }
func (m *Extension) String() string { return proto.CompactTextString(m) }
func (*Extension) ProtoMessage()    {}
func (*Extension) Descriptor() ([]byte, []int) {
	return fileDescriptor_gnmi_ext_0110643c2535c262, []int{0}
}
func (m *Extension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Extension.Unmarshal(m, b)
}
func (m *Extension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Extension.Marshal(b, m, deterministic)
}
func (dst *Extension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Extension.Merge(dst, src)
}
func (m *Extension) XXX_Size() int {
	return xxx_messageInfo_Extension.Size(m)
}
func (m *Extension) XXX_DiscardUnknown() {
	xxx_messageInfo_Extension.DiscardUnknown(m)
}

var xxx_messageInfo_Extension proto.InternalMessageInfo

type isExtension_Ext interface {
	isExtension_Ext()
}

type Extension_RegisteredExt struct {
	RegisteredExt *RegisteredExtension `protobuf:"bytes,1,opt,name=registered_ext,json=registeredExt,proto3,oneof"`
}

type Extension_MasterArbitration struct {
	MasterArbitration *MasterArbitration `protobuf:"bytes,2,opt,name=master_arbitration,json=masterArbitration,proto3,oneof"`
}

func (*Extension_RegisteredExt) isExtension_Ext() {}

func (*Extension_MasterArbitration) isExtension_Ext() {}

func (m *Extension) GetExt() isExtension_Ext {
	if m != nil {
		return m.Ext
	}
	return nil
}

func (m *Extension) GetRegisteredExt() *RegisteredExtension {
	if x, ok := m.GetExt().(*Extension_RegisteredExt); ok {
		return x.RegisteredExt
	}
	return nil
}

func (m *Extension) GetMasterArbitration() *MasterArbitration {
	if x, ok := m.GetExt().(*Extension_MasterArbitration); ok {
		return x.MasterArbitration
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Extension) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Extension_OneofMarshaler, _Extension_OneofUnmarshaler, _Extension_OneofSizer, []interface{}{
		(*Extension_RegisteredExt)(nil),
		(*Extension_MasterArbitration)(nil),
	}
}

func _Extension_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Extension)
	// ext
	switch x := m.Ext.(type) {
	case *Extension_RegisteredExt:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RegisteredExt); err != nil {
			return err
		}
	case *Extension_MasterArbitration:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MasterArbitration); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Extension.Ext has unexpected type %T", x)
	}
	return nil
}

func _Extension_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Extension)
	switch tag {
	case 1: // ext.registered_ext
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RegisteredExtension)
		err := b.DecodeMessage(msg)
		m.Ext = &Extension_RegisteredExt{msg}
		return true, err
	case 2: // ext.master_arbitration
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MasterArbitration)
		err := b.DecodeMessage(msg)
		m.Ext = &Extension_MasterArbitration{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Extension_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Extension)
	// ext
	switch x := m.Ext.(type) {
	case *Extension_RegisteredExt:
		s := proto.Size(x.RegisteredExt)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Extension_MasterArbitration:
		s := proto.Size(x.MasterArbitration)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The RegisteredExtension message defines an extension which is defined outside
// of this file.
type RegisteredExtension struct {
	Id                   ExtensionID `protobuf:"varint,1,opt,name=id,proto3,enum=gnmi_ext.ExtensionID" json:"id,omitempty"`
	Msg                  []byte      `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RegisteredExtension) Reset()         { *m = RegisteredExtension{} }
func (m *RegisteredExtension) String() string { return proto.CompactTextString(m) }
func (*RegisteredExtension) ProtoMessage()    {}
func (*RegisteredExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_gnmi_ext_0110643c2535c262, []int{1}
}
func (m *RegisteredExtension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisteredExtension.Unmarshal(m, b)
}
func (m *RegisteredExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisteredExtension.Marshal(b, m, deterministic)
}
func (dst *RegisteredExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisteredExtension.Merge(dst, src)
}
func (m *RegisteredExtension) XXX_Size() int {
	return xxx_messageInfo_RegisteredExtension.Size(m)
}
func (m *RegisteredExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisteredExtension.DiscardUnknown(m)
}

var xxx_messageInfo_RegisteredExtension proto.InternalMessageInfo

func (m *RegisteredExtension) GetId() ExtensionID {
	if m != nil {
		return m.Id
	}
	return ExtensionID_EID_UNSET
}

func (m *RegisteredExtension) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

// MasterArbitration is used to select the master among multiple gNMI clients
// with the same Roles. The client with the largest election_id is honored as
// the master.
// The document about gNMI master arbitration can be found at
// https://github.com/openconfig/reference/blob/master/rpc/gnmi/gnmi-master-arbitration.md
type MasterArbitration struct {
	Role                 *Role    `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	ElectionId           *Uint128 `protobuf:"bytes,2,opt,name=election_id,json=electionId,proto3" json:"election_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MasterArbitration) Reset()         { *m = MasterArbitration{} }
func (m *MasterArbitration) String() string { return proto.CompactTextString(m) }
func (*MasterArbitration) ProtoMessage()    {}
func (*MasterArbitration) Descriptor() ([]byte, []int) {
	return fileDescriptor_gnmi_ext_0110643c2535c262, []int{2}
}
func (m *MasterArbitration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MasterArbitration.Unmarshal(m, b)
}
func (m *MasterArbitration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MasterArbitration.Marshal(b, m, deterministic)
}
func (dst *MasterArbitration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MasterArbitration.Merge(dst, src)
}
func (m *MasterArbitration) XXX_Size() int {
	return xxx_messageInfo_MasterArbitration.Size(m)
}
func (m *MasterArbitration) XXX_DiscardUnknown() {
	xxx_messageInfo_MasterArbitration.DiscardUnknown(m)
}

var xxx_messageInfo_MasterArbitration proto.InternalMessageInfo

func (m *MasterArbitration) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *MasterArbitration) GetElectionId() *Uint128 {
	if m != nil {
		return m.ElectionId
	}
	return nil
}

// Representation of unsigned 128-bit integer.
type Uint128 struct {
	High                 uint64   `protobuf:"varint,1,opt,name=high,proto3" json:"high,omitempty"`
	Low                  uint64   `protobuf:"varint,2,opt,name=low,proto3" json:"low,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Uint128) Reset()         { *m = Uint128{} }
func (m *Uint128) String() string { return proto.CompactTextString(m) }
func (*Uint128) ProtoMessage()    {}
func (*Uint128) Descriptor() ([]byte, []int) {
	return fileDescriptor_gnmi_ext_0110643c2535c262, []int{3}
}
func (m *Uint128) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Uint128.Unmarshal(m, b)
}
func (m *Uint128) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Uint128.Marshal(b, m, deterministic)
}
func (dst *Uint128) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Uint128.Merge(dst, src)
}
func (m *Uint128) XXX_Size() int {
	return xxx_messageInfo_Uint128.Size(m)
}
func (m *Uint128) XXX_DiscardUnknown() {
	xxx_messageInfo_Uint128.DiscardUnknown(m)
}

var xxx_messageInfo_Uint128 proto.InternalMessageInfo

func (m *Uint128) GetHigh() uint64 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *Uint128) GetLow() uint64 {
	if m != nil {
		return m.Low
	}
	return 0
}

// There can be one master for each role. The role is identified by its id.
type Role struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_gnmi_ext_0110643c2535c262, []int{4}
}
func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (dst *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(dst, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Extension)(nil), "gnmi_ext.Extension")
	proto.RegisterType((*RegisteredExtension)(nil), "gnmi_ext.RegisteredExtension")
	proto.RegisterType((*MasterArbitration)(nil), "gnmi_ext.MasterArbitration")
	proto.RegisterType((*Uint128)(nil), "gnmi_ext.Uint128")
	proto.RegisterType((*Role)(nil), "gnmi_ext.Role")
	proto.RegisterEnum("gnmi_ext.ExtensionID", ExtensionID_name, ExtensionID_value)
}

func init() { proto.RegisterFile("gnmi_ext.proto", fileDescriptor_gnmi_ext_0110643c2535c262) }

var fileDescriptor_gnmi_ext_0110643c2535c262 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x9b, 0x36, 0x5a, 0x7b, 0x6b, 0x43, 0x7a, 0xa5, 0x52, 0x50, 0xa1, 0x04, 0x04, 0x71,
	0x51, 0x31, 0x6e, 0xdc, 0x56, 0x3a, 0xd2, 0x48, 0x1b, 0xca, 0x98, 0x80, 0xae, 0x42, 0x6b, 0x86,
	0x74, 0x30, 0x3f, 0x32, 0x1d, 0xb0, 0x8f, 0xe4, 0x9b, 0xf9, 0x1a, 0x32, 0x43, 0x9b, 0x14, 0x75,
	0x77, 0x7f, 0xbe, 0x39, 0xf7, 0x1c, 0x18, 0xb0, 0x92, 0x3c, 0xe3, 0x11, 0xdb, 0xc8, 0xe1, 0x87,
	0x28, 0x64, 0x81, 0x47, 0xbb, 0xde, 0xf9, 0x32, 0xa0, 0x45, 0x36, 0x92, 0xe5, 0x6b, 0x5e, 0xe4,
	0xf8, 0x08, 0x96, 0x60, 0x09, 0x5f, 0x4b, 0x26, 0x58, 0xac, 0xf6, 0x7d, 0x63, 0x60, 0x5c, 0xb5,
	0xdd, 0x8b, 0x61, 0x29, 0x40, 0xcb, 0x7d, 0xf9, 0x6c, 0x52, 0xa3, 0x1d, 0xb1, 0x3f, 0xc6, 0x29,
	0x60, 0xb6, 0x50, 0x6d, 0xb4, 0x10, 0x4b, 0x2e, 0xc5, 0x42, 0xf2, 0x22, 0xef, 0xd7, 0xb5, 0xd6,
	0x59, 0xa5, 0x35, 0xd3, 0xcc, 0xa8, 0x42, 0x26, 0x35, 0xda, 0xcd, 0x7e, 0x0f, 0x1f, 0x0e, 0xa0,
	0xa1, 0xac, 0xfa, 0x70, 0xf2, 0xcf, 0x71, 0xbc, 0x84, 0x3a, 0x8f, 0xb5, 0x4f, 0xcb, 0xed, 0x55,
	0xda, 0x25, 0xe0, 0x8d, 0x69, 0x9d, 0xc7, 0x68, 0x43, 0x23, 0x5b, 0x27, 0xda, 0xc3, 0x31, 0x55,
	0xa5, 0xf3, 0x0e, 0xdd, 0x3f, 0x06, 0xd0, 0x01, 0x53, 0x14, 0x29, 0xdb, 0xe6, 0xb6, 0xf6, 0x72,
	0x17, 0x29, 0xa3, 0x7a, 0x87, 0x2e, 0xb4, 0x59, 0xca, 0xde, 0x14, 0x1f, 0xf1, 0x78, 0x1b, 0xab,
	0x5b, 0xa1, 0x21, 0xcf, 0xe5, 0xad, 0x7b, 0x4f, 0x61, 0x47, 0x79, 0xb1, 0x73, 0x03, 0xcd, 0xed,
	0x18, 0x11, 0xcc, 0x15, 0x4f, 0x56, 0xfa, 0x84, 0x49, 0x75, 0xad, 0xdc, 0xa5, 0xc5, 0xa7, 0x96,
	0x32, 0xa9, 0x2a, 0x9d, 0x53, 0x30, 0xd5, 0x49, 0xb4, 0xca, 0x78, 0x2d, 0x95, 0xe3, 0x3a, 0x84,
	0xf6, 0x5e, 0x34, 0xec, 0x40, 0x8b, 0x78, 0xe3, 0x28, 0xf4, 0x9f, 0x49, 0x60, 0xd7, 0x70, 0x00,
	0xe7, 0xaa, 0x7d, 0x0a, 0x7d, 0x6f, 0x4e, 0x68, 0x14, 0x90, 0x29, 0x99, 0x91, 0x80, 0xbe, 0x46,
	0x13, 0x32, 0x1a, 0x13, 0x6a, 0x1b, 0xd8, 0x03, 0x5b, 0x11, 0xe4, 0x65, 0x4e, 0xa8, 0x37, 0x23,
	0x7e, 0x30, 0x9a, 0xda, 0xdf, 0xcd, 0xe5, 0xa1, 0xfe, 0x18, 0x77, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x39, 0xa1, 0xea, 0xa7, 0x2a, 0x02, 0x00, 0x00,
}