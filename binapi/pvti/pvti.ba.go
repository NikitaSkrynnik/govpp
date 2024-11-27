// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.11.0
//  VPP:              24.10.0-4~g16238f529
// source: plugins/pvti.api.json

// Package pvti contains generated bindings for API file pvti.api.
//
// Contents:
// -  1 struct
// -  6 messages
package pvti

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
	APIFile    = "pvti"
	APIVersion = "0.0.1"
	VersionCrc = 0xb319860
)

// PvtiTunnel defines type 'pvti_tunnel'.
type PvtiTunnel struct {
	SwIfIndex              interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	LocalIP                ip_types.Address               `binapi:"address,name=local_ip" json:"local_ip,omitempty"`
	LocalPort              uint16                         `binapi:"u16,name=local_port" json:"local_port,omitempty"`
	RemoteIP               ip_types.Address               `binapi:"address,name=remote_ip" json:"remote_ip,omitempty"`
	PeerAddressFromPayload bool                           `binapi:"bool,name=peer_address_from_payload" json:"peer_address_from_payload,omitempty"`
	RemotePort             uint16                         `binapi:"u16,name=remote_port" json:"remote_port,omitempty"`
	UnderlayMtu            uint16                         `binapi:"u16,name=underlay_mtu" json:"underlay_mtu,omitempty"`
	UnderlayFibIndex       uint32                         `binapi:"u32,name=underlay_fib_index" json:"underlay_fib_index,omitempty"`
}

// @brief API to enable / disable pvti on an interface
//   - enable_disable - 1 to enable, 0 to disable the feature
//   - sw_if_index - interface handle
//
// PvtiInterfaceCreate defines message 'pvti_interface_create'.
// InProgress: the message form may change in the future versions
type PvtiInterfaceCreate struct {
	Interface PvtiTunnel `binapi:"pvti_tunnel,name=interface" json:"interface,omitempty"`
}

func (m *PvtiInterfaceCreate) Reset()               { *m = PvtiInterfaceCreate{} }
func (*PvtiInterfaceCreate) GetMessageName() string { return "pvti_interface_create" }
func (*PvtiInterfaceCreate) GetCrcString() string   { return "a1e95595" }
func (*PvtiInterfaceCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PvtiInterfaceCreate) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.Interface.SwIfIndex
	size += 1      // m.Interface.LocalIP.Af
	size += 1 * 16 // m.Interface.LocalIP.Un
	size += 2      // m.Interface.LocalPort
	size += 1      // m.Interface.RemoteIP.Af
	size += 1 * 16 // m.Interface.RemoteIP.Un
	size += 1      // m.Interface.PeerAddressFromPayload
	size += 2      // m.Interface.RemotePort
	size += 2      // m.Interface.UnderlayMtu
	size += 4      // m.Interface.UnderlayFibIndex
	return size
}
func (m *PvtiInterfaceCreate) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.Interface.SwIfIndex))
	buf.EncodeUint8(uint8(m.Interface.LocalIP.Af))
	buf.EncodeBytes(m.Interface.LocalIP.Un.XXX_UnionData[:], 16)
	buf.EncodeUint16(m.Interface.LocalPort)
	buf.EncodeUint8(uint8(m.Interface.RemoteIP.Af))
	buf.EncodeBytes(m.Interface.RemoteIP.Un.XXX_UnionData[:], 16)
	buf.EncodeBool(m.Interface.PeerAddressFromPayload)
	buf.EncodeUint16(m.Interface.RemotePort)
	buf.EncodeUint16(m.Interface.UnderlayMtu)
	buf.EncodeUint32(m.Interface.UnderlayFibIndex)
	return buf.Bytes(), nil
}
func (m *PvtiInterfaceCreate) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Interface.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Interface.LocalIP.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Interface.LocalIP.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Interface.LocalPort = buf.DecodeUint16()
	m.Interface.RemoteIP.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Interface.RemoteIP.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Interface.PeerAddressFromPayload = buf.DecodeBool()
	m.Interface.RemotePort = buf.DecodeUint16()
	m.Interface.UnderlayMtu = buf.DecodeUint16()
	m.Interface.UnderlayFibIndex = buf.DecodeUint32()
	return nil
}

// PvtiInterfaceCreateReply defines message 'pvti_interface_create_reply'.
// InProgress: the message form may change in the future versions
type PvtiInterfaceCreateReply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *PvtiInterfaceCreateReply) Reset()               { *m = PvtiInterfaceCreateReply{} }
func (*PvtiInterfaceCreateReply) GetMessageName() string { return "pvti_interface_create_reply" }
func (*PvtiInterfaceCreateReply) GetCrcString() string   { return "5383d31f" }
func (*PvtiInterfaceCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PvtiInterfaceCreateReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *PvtiInterfaceCreateReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *PvtiInterfaceCreateReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// PvtiInterfaceDelete defines message 'pvti_interface_delete'.
// InProgress: the message form may change in the future versions
type PvtiInterfaceDelete struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *PvtiInterfaceDelete) Reset()               { *m = PvtiInterfaceDelete{} }
func (*PvtiInterfaceDelete) GetMessageName() string { return "pvti_interface_delete" }
func (*PvtiInterfaceDelete) GetCrcString() string   { return "f9e6675e" }
func (*PvtiInterfaceDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PvtiInterfaceDelete) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *PvtiInterfaceDelete) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *PvtiInterfaceDelete) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// PvtiInterfaceDeleteReply defines message 'pvti_interface_delete_reply'.
// InProgress: the message form may change in the future versions
type PvtiInterfaceDeleteReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *PvtiInterfaceDeleteReply) Reset()               { *m = PvtiInterfaceDeleteReply{} }
func (*PvtiInterfaceDeleteReply) GetMessageName() string { return "pvti_interface_delete_reply" }
func (*PvtiInterfaceDeleteReply) GetCrcString() string   { return "e8d4e804" }
func (*PvtiInterfaceDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PvtiInterfaceDeleteReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *PvtiInterfaceDeleteReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *PvtiInterfaceDeleteReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// PvtiInterfaceDetails defines message 'pvti_interface_details'.
// InProgress: the message form may change in the future versions
type PvtiInterfaceDetails struct {
	Interface PvtiTunnel `binapi:"pvti_tunnel,name=interface" json:"interface,omitempty"`
}

func (m *PvtiInterfaceDetails) Reset()               { *m = PvtiInterfaceDetails{} }
func (*PvtiInterfaceDetails) GetMessageName() string { return "pvti_interface_details" }
func (*PvtiInterfaceDetails) GetCrcString() string   { return "a26072b7" }
func (*PvtiInterfaceDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PvtiInterfaceDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.Interface.SwIfIndex
	size += 1      // m.Interface.LocalIP.Af
	size += 1 * 16 // m.Interface.LocalIP.Un
	size += 2      // m.Interface.LocalPort
	size += 1      // m.Interface.RemoteIP.Af
	size += 1 * 16 // m.Interface.RemoteIP.Un
	size += 1      // m.Interface.PeerAddressFromPayload
	size += 2      // m.Interface.RemotePort
	size += 2      // m.Interface.UnderlayMtu
	size += 4      // m.Interface.UnderlayFibIndex
	return size
}
func (m *PvtiInterfaceDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.Interface.SwIfIndex))
	buf.EncodeUint8(uint8(m.Interface.LocalIP.Af))
	buf.EncodeBytes(m.Interface.LocalIP.Un.XXX_UnionData[:], 16)
	buf.EncodeUint16(m.Interface.LocalPort)
	buf.EncodeUint8(uint8(m.Interface.RemoteIP.Af))
	buf.EncodeBytes(m.Interface.RemoteIP.Un.XXX_UnionData[:], 16)
	buf.EncodeBool(m.Interface.PeerAddressFromPayload)
	buf.EncodeUint16(m.Interface.RemotePort)
	buf.EncodeUint16(m.Interface.UnderlayMtu)
	buf.EncodeUint32(m.Interface.UnderlayFibIndex)
	return buf.Bytes(), nil
}
func (m *PvtiInterfaceDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Interface.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Interface.LocalIP.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Interface.LocalIP.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Interface.LocalPort = buf.DecodeUint16()
	m.Interface.RemoteIP.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Interface.RemoteIP.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Interface.PeerAddressFromPayload = buf.DecodeBool()
	m.Interface.RemotePort = buf.DecodeUint16()
	m.Interface.UnderlayMtu = buf.DecodeUint16()
	m.Interface.UnderlayFibIndex = buf.DecodeUint32()
	return nil
}

// PvtiInterfaceDump defines message 'pvti_interface_dump'.
// InProgress: the message form may change in the future versions
type PvtiInterfaceDump struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *PvtiInterfaceDump) Reset()               { *m = PvtiInterfaceDump{} }
func (*PvtiInterfaceDump) GetMessageName() string { return "pvti_interface_dump" }
func (*PvtiInterfaceDump) GetCrcString() string   { return "f9e6675e" }
func (*PvtiInterfaceDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PvtiInterfaceDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *PvtiInterfaceDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *PvtiInterfaceDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

func init() { file_pvti_binapi_init() }
func file_pvti_binapi_init() {
	api.RegisterMessage((*PvtiInterfaceCreate)(nil), "pvti_interface_create_a1e95595")
	api.RegisterMessage((*PvtiInterfaceCreateReply)(nil), "pvti_interface_create_reply_5383d31f")
	api.RegisterMessage((*PvtiInterfaceDelete)(nil), "pvti_interface_delete_f9e6675e")
	api.RegisterMessage((*PvtiInterfaceDeleteReply)(nil), "pvti_interface_delete_reply_e8d4e804")
	api.RegisterMessage((*PvtiInterfaceDetails)(nil), "pvti_interface_details_a26072b7")
	api.RegisterMessage((*PvtiInterfaceDump)(nil), "pvti_interface_dump_f9e6675e")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*PvtiInterfaceCreate)(nil),
		(*PvtiInterfaceCreateReply)(nil),
		(*PvtiInterfaceDelete)(nil),
		(*PvtiInterfaceDeleteReply)(nil),
		(*PvtiInterfaceDetails)(nil),
		(*PvtiInterfaceDump)(nil),
	}
}