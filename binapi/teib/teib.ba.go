// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.8.0
//  VPP:              23.10-rc0~166-gb034cda73
// source: core/teib.api.json

// Package teib contains generated bindings for API file teib.api.
//
// Contents:
// -  1 struct
// -  4 messages
package teib

import (
	interface_types "github.com/networkservicemesh/govpp/binapi/interface_types"
	ip_types "github.com/networkservicemesh/govpp/binapi/ip_types"
	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "teib"
	APIVersion = "1.0.0"
	VersionCrc = 0x14ded985
)

// TeibEntry defines type 'teib_entry'.
type TeibEntry struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Peer      ip_types.Address               `binapi:"address,name=peer" json:"peer,omitempty"`
	Nh        ip_types.Address               `binapi:"address,name=nh" json:"nh,omitempty"`
	NhTableID uint32                         `binapi:"u32,name=nh_table_id" json:"nh_table_id,omitempty"`
}

// TeibDetails defines message 'teib_details'.
type TeibDetails struct {
	Entry TeibEntry `binapi:"teib_entry,name=entry" json:"entry,omitempty"`
}

func (m *TeibDetails) Reset()               { *m = TeibDetails{} }
func (*TeibDetails) GetMessageName() string { return "teib_details" }
func (*TeibDetails) GetCrcString() string   { return "981ee1a1" }
func (*TeibDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *TeibDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.Entry.SwIfIndex
	size += 1      // m.Entry.Peer.Af
	size += 1 * 16 // m.Entry.Peer.Un
	size += 1      // m.Entry.Nh.Af
	size += 1 * 16 // m.Entry.Nh.Un
	size += 4      // m.Entry.NhTableID
	return size
}
func (m *TeibDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.Entry.SwIfIndex))
	buf.EncodeUint8(uint8(m.Entry.Peer.Af))
	buf.EncodeBytes(m.Entry.Peer.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(uint8(m.Entry.Nh.Af))
	buf.EncodeBytes(m.Entry.Nh.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(m.Entry.NhTableID)
	return buf.Bytes(), nil
}
func (m *TeibDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Entry.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Entry.Peer.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Entry.Peer.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Entry.Nh.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Entry.Nh.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Entry.NhTableID = buf.DecodeUint32()
	return nil
}

// TeibDump defines message 'teib_dump'.
type TeibDump struct{}

func (m *TeibDump) Reset()               { *m = TeibDump{} }
func (*TeibDump) GetMessageName() string { return "teib_dump" }
func (*TeibDump) GetCrcString() string   { return "51077d14" }
func (*TeibDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TeibDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *TeibDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *TeibDump) Unmarshal(b []byte) error {
	return nil
}

// TEIB Entry
//   - sw_if_index
//
// TeibEntryAddDel defines message 'teib_entry_add_del'.
type TeibEntryAddDel struct {
	IsAdd uint8     `binapi:"u8,name=is_add" json:"is_add,omitempty"`
	Entry TeibEntry `binapi:"teib_entry,name=entry" json:"entry,omitempty"`
}

func (m *TeibEntryAddDel) Reset()               { *m = TeibEntryAddDel{} }
func (*TeibEntryAddDel) GetMessageName() string { return "teib_entry_add_del" }
func (*TeibEntryAddDel) GetCrcString() string   { return "8016cfd2" }
func (*TeibEntryAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TeibEntryAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.IsAdd
	size += 4      // m.Entry.SwIfIndex
	size += 1      // m.Entry.Peer.Af
	size += 1 * 16 // m.Entry.Peer.Un
	size += 1      // m.Entry.Nh.Af
	size += 1 * 16 // m.Entry.Nh.Un
	size += 4      // m.Entry.NhTableID
	return size
}
func (m *TeibEntryAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.IsAdd)
	buf.EncodeUint32(uint32(m.Entry.SwIfIndex))
	buf.EncodeUint8(uint8(m.Entry.Peer.Af))
	buf.EncodeBytes(m.Entry.Peer.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(uint8(m.Entry.Nh.Af))
	buf.EncodeBytes(m.Entry.Nh.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(m.Entry.NhTableID)
	return buf.Bytes(), nil
}
func (m *TeibEntryAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeUint8()
	m.Entry.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Entry.Peer.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Entry.Peer.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Entry.Nh.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Entry.Nh.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Entry.NhTableID = buf.DecodeUint32()
	return nil
}

// TeibEntryAddDelReply defines message 'teib_entry_add_del_reply'.
type TeibEntryAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *TeibEntryAddDelReply) Reset()               { *m = TeibEntryAddDelReply{} }
func (*TeibEntryAddDelReply) GetMessageName() string { return "teib_entry_add_del_reply" }
func (*TeibEntryAddDelReply) GetCrcString() string   { return "e8d4e804" }
func (*TeibEntryAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *TeibEntryAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *TeibEntryAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *TeibEntryAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_teib_binapi_init() }
func file_teib_binapi_init() {
	api.RegisterMessage((*TeibDetails)(nil), "teib_details_981ee1a1")
	api.RegisterMessage((*TeibDump)(nil), "teib_dump_51077d14")
	api.RegisterMessage((*TeibEntryAddDel)(nil), "teib_entry_add_del_8016cfd2")
	api.RegisterMessage((*TeibEntryAddDelReply)(nil), "teib_entry_add_del_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*TeibDetails)(nil),
		(*TeibDump)(nil),
		(*TeibEntryAddDel)(nil),
		(*TeibEntryAddDelReply)(nil),
	}
}
