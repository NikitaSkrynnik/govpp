// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.8.0
//  VPP:              23.10-rc0~166-gb034cda73
// source: plugins/http_static.api.json

// Package http_static contains generated bindings for API file http_static.api.
//
// Contents:
// -  2 messages
package http_static

import (
	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "http_static"
	APIVersion = "2.1.0"
	VersionCrc = 0xd29e72e9
)

// Configure and enable the static http server
//   - fifo_size - size (in bytes) of the session FIFOs
//   - cache_size_limit - size (in bytes) of the in-memory file data cache
//   - prealloc_fifos - number of preallocated fifos (usually 0)
//   - private_segment_size - fifo segment size (usually 0)
//   - www_root - html root path
//   - uri - bind URI, defaults to "tcp://0.0.0.0/80"
//
// HTTPStaticEnable defines message 'http_static_enable'.
type HTTPStaticEnable struct {
	FifoSize           uint32 `binapi:"u32,name=fifo_size" json:"fifo_size,omitempty"`
	CacheSizeLimit     uint32 `binapi:"u32,name=cache_size_limit" json:"cache_size_limit,omitempty"`
	PreallocFifos      uint32 `binapi:"u32,name=prealloc_fifos" json:"prealloc_fifos,omitempty"`
	PrivateSegmentSize uint32 `binapi:"u32,name=private_segment_size" json:"private_segment_size,omitempty"`
	WwwRoot            string `binapi:"string[256],name=www_root" json:"www_root,omitempty"`
	URI                string `binapi:"string[256],name=uri" json:"uri,omitempty"`
}

func (m *HTTPStaticEnable) Reset()               { *m = HTTPStaticEnable{} }
func (*HTTPStaticEnable) GetMessageName() string { return "http_static_enable" }
func (*HTTPStaticEnable) GetCrcString() string   { return "075f8292" }
func (*HTTPStaticEnable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *HTTPStaticEnable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4   // m.FifoSize
	size += 4   // m.CacheSizeLimit
	size += 4   // m.PreallocFifos
	size += 4   // m.PrivateSegmentSize
	size += 256 // m.WwwRoot
	size += 256 // m.URI
	return size
}
func (m *HTTPStaticEnable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.FifoSize)
	buf.EncodeUint32(m.CacheSizeLimit)
	buf.EncodeUint32(m.PreallocFifos)
	buf.EncodeUint32(m.PrivateSegmentSize)
	buf.EncodeString(m.WwwRoot, 256)
	buf.EncodeString(m.URI, 256)
	return buf.Bytes(), nil
}
func (m *HTTPStaticEnable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.FifoSize = buf.DecodeUint32()
	m.CacheSizeLimit = buf.DecodeUint32()
	m.PreallocFifos = buf.DecodeUint32()
	m.PrivateSegmentSize = buf.DecodeUint32()
	m.WwwRoot = buf.DecodeString(256)
	m.URI = buf.DecodeString(256)
	return nil
}

// HTTPStaticEnableReply defines message 'http_static_enable_reply'.
type HTTPStaticEnableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *HTTPStaticEnableReply) Reset()               { *m = HTTPStaticEnableReply{} }
func (*HTTPStaticEnableReply) GetMessageName() string { return "http_static_enable_reply" }
func (*HTTPStaticEnableReply) GetCrcString() string   { return "e8d4e804" }
func (*HTTPStaticEnableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *HTTPStaticEnableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *HTTPStaticEnableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *HTTPStaticEnableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_http_static_binapi_init() }
func file_http_static_binapi_init() {
	api.RegisterMessage((*HTTPStaticEnable)(nil), "http_static_enable_075f8292")
	api.RegisterMessage((*HTTPStaticEnableReply)(nil), "http_static_enable_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*HTTPStaticEnable)(nil),
		(*HTTPStaticEnableReply)(nil),
	}
}
