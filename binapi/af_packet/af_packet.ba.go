// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              22.02-rc0~101-ga2b3e19df
// source: /usr/share/vpp/api/core/af_packet.api.json

// Package af_packet contains generated bindings for API file af_packet.api.
//
// Contents:
//  10 messages
//
package af_packet

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	ethernet_types "github.com/edwarnicke/govpp/binapi/ethernet_types"
	interface_types "github.com/edwarnicke/govpp/binapi/interface_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "af_packet"
	APIVersion = "2.0.0"
	VersionCrc = 0x589bd50e
)

// AfPacketCreate defines message 'af_packet_create'.
type AfPacketCreate struct {
	HwAddr          ethernet_types.MacAddress `binapi:"mac_address,name=hw_addr" json:"hw_addr,omitempty"`
	UseRandomHwAddr bool                      `binapi:"bool,name=use_random_hw_addr" json:"use_random_hw_addr,omitempty"`
	HostIfName      string                    `binapi:"string[64],name=host_if_name" json:"host_if_name,omitempty"`
}

func (m *AfPacketCreate) Reset()               { *m = AfPacketCreate{} }
func (*AfPacketCreate) GetMessageName() string { return "af_packet_create" }
func (*AfPacketCreate) GetCrcString() string   { return "a190415f" }
func (*AfPacketCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AfPacketCreate) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 6 // m.HwAddr
	size += 1     // m.UseRandomHwAddr
	size += 64    // m.HostIfName
	return size
}
func (m *AfPacketCreate) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.HwAddr[:], 6)
	buf.EncodeBool(m.UseRandomHwAddr)
	buf.EncodeString(m.HostIfName, 64)
	return buf.Bytes(), nil
}
func (m *AfPacketCreate) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	copy(m.HwAddr[:], buf.DecodeBytes(6))
	m.UseRandomHwAddr = buf.DecodeBool()
	m.HostIfName = buf.DecodeString(64)
	return nil
}

// AfPacketCreateReply defines message 'af_packet_create_reply'.
type AfPacketCreateReply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *AfPacketCreateReply) Reset()               { *m = AfPacketCreateReply{} }
func (*AfPacketCreateReply) GetMessageName() string { return "af_packet_create_reply" }
func (*AfPacketCreateReply) GetCrcString() string   { return "5383d31f" }
func (*AfPacketCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AfPacketCreateReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *AfPacketCreateReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *AfPacketCreateReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// AfPacketCreateV2 defines message 'af_packet_create_v2'.
type AfPacketCreateV2 struct {
	HwAddr           ethernet_types.MacAddress `binapi:"mac_address,name=hw_addr" json:"hw_addr,omitempty"`
	UseRandomHwAddr  bool                      `binapi:"bool,name=use_random_hw_addr" json:"use_random_hw_addr,omitempty"`
	HostIfName       string                    `binapi:"string[64],name=host_if_name" json:"host_if_name,omitempty"`
	RxFrameSize      uint32                    `binapi:"u32,name=rx_frame_size" json:"rx_frame_size,omitempty"`
	TxFrameSize      uint32                    `binapi:"u32,name=tx_frame_size" json:"tx_frame_size,omitempty"`
	RxFramesPerBlock uint32                    `binapi:"u32,name=rx_frames_per_block" json:"rx_frames_per_block,omitempty"`
	TxFramesPerBlock uint32                    `binapi:"u32,name=tx_frames_per_block" json:"tx_frames_per_block,omitempty"`
	Flags            uint32                    `binapi:"u32,name=flags" json:"flags,omitempty"`
	NumRxQueues      uint16                    `binapi:"u16,name=num_rx_queues,default=1" json:"num_rx_queues,omitempty"`
}

func (m *AfPacketCreateV2) Reset()               { *m = AfPacketCreateV2{} }
func (*AfPacketCreateV2) GetMessageName() string { return "af_packet_create_v2" }
func (*AfPacketCreateV2) GetCrcString() string   { return "4aff0436" }
func (*AfPacketCreateV2) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AfPacketCreateV2) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 6 // m.HwAddr
	size += 1     // m.UseRandomHwAddr
	size += 64    // m.HostIfName
	size += 4     // m.RxFrameSize
	size += 4     // m.TxFrameSize
	size += 4     // m.RxFramesPerBlock
	size += 4     // m.TxFramesPerBlock
	size += 4     // m.Flags
	size += 2     // m.NumRxQueues
	return size
}
func (m *AfPacketCreateV2) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.HwAddr[:], 6)
	buf.EncodeBool(m.UseRandomHwAddr)
	buf.EncodeString(m.HostIfName, 64)
	buf.EncodeUint32(m.RxFrameSize)
	buf.EncodeUint32(m.TxFrameSize)
	buf.EncodeUint32(m.RxFramesPerBlock)
	buf.EncodeUint32(m.TxFramesPerBlock)
	buf.EncodeUint32(m.Flags)
	buf.EncodeUint16(m.NumRxQueues)
	return buf.Bytes(), nil
}
func (m *AfPacketCreateV2) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	copy(m.HwAddr[:], buf.DecodeBytes(6))
	m.UseRandomHwAddr = buf.DecodeBool()
	m.HostIfName = buf.DecodeString(64)
	m.RxFrameSize = buf.DecodeUint32()
	m.TxFrameSize = buf.DecodeUint32()
	m.RxFramesPerBlock = buf.DecodeUint32()
	m.TxFramesPerBlock = buf.DecodeUint32()
	m.Flags = buf.DecodeUint32()
	m.NumRxQueues = buf.DecodeUint16()
	return nil
}

// AfPacketCreateV2Reply defines message 'af_packet_create_v2_reply'.
type AfPacketCreateV2Reply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *AfPacketCreateV2Reply) Reset()               { *m = AfPacketCreateV2Reply{} }
func (*AfPacketCreateV2Reply) GetMessageName() string { return "af_packet_create_v2_reply" }
func (*AfPacketCreateV2Reply) GetCrcString() string   { return "5383d31f" }
func (*AfPacketCreateV2Reply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AfPacketCreateV2Reply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *AfPacketCreateV2Reply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *AfPacketCreateV2Reply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// AfPacketDelete defines message 'af_packet_delete'.
type AfPacketDelete struct {
	HostIfName string `binapi:"string[64],name=host_if_name" json:"host_if_name,omitempty"`
}

func (m *AfPacketDelete) Reset()               { *m = AfPacketDelete{} }
func (*AfPacketDelete) GetMessageName() string { return "af_packet_delete" }
func (*AfPacketDelete) GetCrcString() string   { return "863fa648" }
func (*AfPacketDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AfPacketDelete) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 64 // m.HostIfName
	return size
}
func (m *AfPacketDelete) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.HostIfName, 64)
	return buf.Bytes(), nil
}
func (m *AfPacketDelete) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.HostIfName = buf.DecodeString(64)
	return nil
}

// AfPacketDeleteReply defines message 'af_packet_delete_reply'.
type AfPacketDeleteReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *AfPacketDeleteReply) Reset()               { *m = AfPacketDeleteReply{} }
func (*AfPacketDeleteReply) GetMessageName() string { return "af_packet_delete_reply" }
func (*AfPacketDeleteReply) GetCrcString() string   { return "e8d4e804" }
func (*AfPacketDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AfPacketDeleteReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *AfPacketDeleteReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *AfPacketDeleteReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// AfPacketDetails defines message 'af_packet_details'.
type AfPacketDetails struct {
	SwIfIndex  interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	HostIfName string                         `binapi:"string[64],name=host_if_name" json:"host_if_name,omitempty"`
}

func (m *AfPacketDetails) Reset()               { *m = AfPacketDetails{} }
func (*AfPacketDetails) GetMessageName() string { return "af_packet_details" }
func (*AfPacketDetails) GetCrcString() string   { return "58c7c042" }
func (*AfPacketDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AfPacketDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4  // m.SwIfIndex
	size += 64 // m.HostIfName
	return size
}
func (m *AfPacketDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeString(m.HostIfName, 64)
	return buf.Bytes(), nil
}
func (m *AfPacketDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.HostIfName = buf.DecodeString(64)
	return nil
}

// AfPacketDump defines message 'af_packet_dump'.
type AfPacketDump struct{}

func (m *AfPacketDump) Reset()               { *m = AfPacketDump{} }
func (*AfPacketDump) GetMessageName() string { return "af_packet_dump" }
func (*AfPacketDump) GetCrcString() string   { return "51077d14" }
func (*AfPacketDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AfPacketDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *AfPacketDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *AfPacketDump) Unmarshal(b []byte) error {
	return nil
}

// AfPacketSetL4CksumOffload defines message 'af_packet_set_l4_cksum_offload'.
type AfPacketSetL4CksumOffload struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Set       bool                           `binapi:"bool,name=set" json:"set,omitempty"`
}

func (m *AfPacketSetL4CksumOffload) Reset()               { *m = AfPacketSetL4CksumOffload{} }
func (*AfPacketSetL4CksumOffload) GetMessageName() string { return "af_packet_set_l4_cksum_offload" }
func (*AfPacketSetL4CksumOffload) GetCrcString() string   { return "319cd5c8" }
func (*AfPacketSetL4CksumOffload) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AfPacketSetL4CksumOffload) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.Set
	return size
}
func (m *AfPacketSetL4CksumOffload) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.Set)
	return buf.Bytes(), nil
}
func (m *AfPacketSetL4CksumOffload) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Set = buf.DecodeBool()
	return nil
}

// AfPacketSetL4CksumOffloadReply defines message 'af_packet_set_l4_cksum_offload_reply'.
type AfPacketSetL4CksumOffloadReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *AfPacketSetL4CksumOffloadReply) Reset() { *m = AfPacketSetL4CksumOffloadReply{} }
func (*AfPacketSetL4CksumOffloadReply) GetMessageName() string {
	return "af_packet_set_l4_cksum_offload_reply"
}
func (*AfPacketSetL4CksumOffloadReply) GetCrcString() string { return "e8d4e804" }
func (*AfPacketSetL4CksumOffloadReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AfPacketSetL4CksumOffloadReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *AfPacketSetL4CksumOffloadReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *AfPacketSetL4CksumOffloadReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_af_packet_binapi_init() }
func file_af_packet_binapi_init() {
	api.RegisterMessage((*AfPacketCreate)(nil), "af_packet_create_a190415f")
	api.RegisterMessage((*AfPacketCreateReply)(nil), "af_packet_create_reply_5383d31f")
	api.RegisterMessage((*AfPacketCreateV2)(nil), "af_packet_create_v2_4aff0436")
	api.RegisterMessage((*AfPacketCreateV2Reply)(nil), "af_packet_create_v2_reply_5383d31f")
	api.RegisterMessage((*AfPacketDelete)(nil), "af_packet_delete_863fa648")
	api.RegisterMessage((*AfPacketDeleteReply)(nil), "af_packet_delete_reply_e8d4e804")
	api.RegisterMessage((*AfPacketDetails)(nil), "af_packet_details_58c7c042")
	api.RegisterMessage((*AfPacketDump)(nil), "af_packet_dump_51077d14")
	api.RegisterMessage((*AfPacketSetL4CksumOffload)(nil), "af_packet_set_l4_cksum_offload_319cd5c8")
	api.RegisterMessage((*AfPacketSetL4CksumOffloadReply)(nil), "af_packet_set_l4_cksum_offload_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*AfPacketCreate)(nil),
		(*AfPacketCreateReply)(nil),
		(*AfPacketCreateV2)(nil),
		(*AfPacketCreateV2Reply)(nil),
		(*AfPacketDelete)(nil),
		(*AfPacketDeleteReply)(nil),
		(*AfPacketDetails)(nil),
		(*AfPacketDump)(nil),
		(*AfPacketSetL4CksumOffload)(nil),
		(*AfPacketSetL4CksumOffloadReply)(nil),
	}
}
