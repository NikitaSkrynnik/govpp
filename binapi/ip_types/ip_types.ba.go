// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.11.0
//  VPP:              24.10.0-4~g16238f529
// source: core/ip_types.api.json

// Package ip_types contains generated bindings for API file ip_types.api.
//
// Contents:
// -  5 aliases
// -  5 enums
// -  8 structs
// -  1 union
package ip_types

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "ip_types"
	APIVersion = "3.0.0"
	VersionCrc = 0xfee023ed
)

// AddressFamily defines enum 'address_family'.
type AddressFamily uint8

const (
	ADDRESS_IP4 AddressFamily = 0
	ADDRESS_IP6 AddressFamily = 1
)

var (
	AddressFamily_name = map[uint8]string{
		0: "ADDRESS_IP4",
		1: "ADDRESS_IP6",
	}
	AddressFamily_value = map[string]uint8{
		"ADDRESS_IP4": 0,
		"ADDRESS_IP6": 1,
	}
)

func (x AddressFamily) String() string {
	s, ok := AddressFamily_name[uint8(x)]
	if ok {
		return s
	}
	return "AddressFamily(" + strconv.Itoa(int(x)) + ")"
}

// IPDscp defines enum 'ip_dscp'.
type IPDscp uint8

const (
	IP_API_DSCP_CS0  IPDscp = 0
	IP_API_DSCP_CS1  IPDscp = 8
	IP_API_DSCP_AF11 IPDscp = 10
	IP_API_DSCP_AF12 IPDscp = 12
	IP_API_DSCP_AF13 IPDscp = 14
	IP_API_DSCP_CS2  IPDscp = 16
	IP_API_DSCP_AF21 IPDscp = 18
	IP_API_DSCP_AF22 IPDscp = 20
	IP_API_DSCP_AF23 IPDscp = 22
	IP_API_DSCP_CS3  IPDscp = 24
	IP_API_DSCP_AF31 IPDscp = 26
	IP_API_DSCP_AF32 IPDscp = 28
	IP_API_DSCP_AF33 IPDscp = 30
	IP_API_DSCP_CS4  IPDscp = 32
	IP_API_DSCP_AF41 IPDscp = 34
	IP_API_DSCP_AF42 IPDscp = 36
	IP_API_DSCP_AF43 IPDscp = 38
	IP_API_DSCP_CS5  IPDscp = 40
	IP_API_DSCP_EF   IPDscp = 46
	IP_API_DSCP_CS6  IPDscp = 48
	IP_API_DSCP_CS7  IPDscp = 50
)

var (
	IPDscp_name = map[uint8]string{
		0:  "IP_API_DSCP_CS0",
		8:  "IP_API_DSCP_CS1",
		10: "IP_API_DSCP_AF11",
		12: "IP_API_DSCP_AF12",
		14: "IP_API_DSCP_AF13",
		16: "IP_API_DSCP_CS2",
		18: "IP_API_DSCP_AF21",
		20: "IP_API_DSCP_AF22",
		22: "IP_API_DSCP_AF23",
		24: "IP_API_DSCP_CS3",
		26: "IP_API_DSCP_AF31",
		28: "IP_API_DSCP_AF32",
		30: "IP_API_DSCP_AF33",
		32: "IP_API_DSCP_CS4",
		34: "IP_API_DSCP_AF41",
		36: "IP_API_DSCP_AF42",
		38: "IP_API_DSCP_AF43",
		40: "IP_API_DSCP_CS5",
		46: "IP_API_DSCP_EF",
		48: "IP_API_DSCP_CS6",
		50: "IP_API_DSCP_CS7",
	}
	IPDscp_value = map[string]uint8{
		"IP_API_DSCP_CS0":  0,
		"IP_API_DSCP_CS1":  8,
		"IP_API_DSCP_AF11": 10,
		"IP_API_DSCP_AF12": 12,
		"IP_API_DSCP_AF13": 14,
		"IP_API_DSCP_CS2":  16,
		"IP_API_DSCP_AF21": 18,
		"IP_API_DSCP_AF22": 20,
		"IP_API_DSCP_AF23": 22,
		"IP_API_DSCP_CS3":  24,
		"IP_API_DSCP_AF31": 26,
		"IP_API_DSCP_AF32": 28,
		"IP_API_DSCP_AF33": 30,
		"IP_API_DSCP_CS4":  32,
		"IP_API_DSCP_AF41": 34,
		"IP_API_DSCP_AF42": 36,
		"IP_API_DSCP_AF43": 38,
		"IP_API_DSCP_CS5":  40,
		"IP_API_DSCP_EF":   46,
		"IP_API_DSCP_CS6":  48,
		"IP_API_DSCP_CS7":  50,
	}
)

func (x IPDscp) String() string {
	s, ok := IPDscp_name[uint8(x)]
	if ok {
		return s
	}
	return "IPDscp(" + strconv.Itoa(int(x)) + ")"
}

// IPEcn defines enum 'ip_ecn'.
type IPEcn uint8

const (
	IP_API_ECN_NONE IPEcn = 0
	IP_API_ECN_ECT0 IPEcn = 1
	IP_API_ECN_ECT1 IPEcn = 2
	IP_API_ECN_CE   IPEcn = 3
)

var (
	IPEcn_name = map[uint8]string{
		0: "IP_API_ECN_NONE",
		1: "IP_API_ECN_ECT0",
		2: "IP_API_ECN_ECT1",
		3: "IP_API_ECN_CE",
	}
	IPEcn_value = map[string]uint8{
		"IP_API_ECN_NONE": 0,
		"IP_API_ECN_ECT0": 1,
		"IP_API_ECN_ECT1": 2,
		"IP_API_ECN_CE":   3,
	}
)

func (x IPEcn) String() string {
	s, ok := IPEcn_name[uint8(x)]
	if ok {
		return s
	}
	return "IPEcn(" + strconv.Itoa(int(x)) + ")"
}

// IPFeatureLocation defines enum 'ip_feature_location'.
type IPFeatureLocation uint8

const (
	IP_API_FEATURE_INPUT  IPFeatureLocation = 0
	IP_API_FEATURE_OUTPUT IPFeatureLocation = 1
	IP_API_FEATURE_LOCAL  IPFeatureLocation = 2
	IP_API_FEATURE_PUNT   IPFeatureLocation = 3
	IP_API_FEATURE_DROP   IPFeatureLocation = 4
)

var (
	IPFeatureLocation_name = map[uint8]string{
		0: "IP_API_FEATURE_INPUT",
		1: "IP_API_FEATURE_OUTPUT",
		2: "IP_API_FEATURE_LOCAL",
		3: "IP_API_FEATURE_PUNT",
		4: "IP_API_FEATURE_DROP",
	}
	IPFeatureLocation_value = map[string]uint8{
		"IP_API_FEATURE_INPUT":  0,
		"IP_API_FEATURE_OUTPUT": 1,
		"IP_API_FEATURE_LOCAL":  2,
		"IP_API_FEATURE_PUNT":   3,
		"IP_API_FEATURE_DROP":   4,
	}
)

func (x IPFeatureLocation) String() string {
	s, ok := IPFeatureLocation_name[uint8(x)]
	if ok {
		return s
	}
	return "IPFeatureLocation(" + strconv.Itoa(int(x)) + ")"
}

// IPProto defines enum 'ip_proto'.
type IPProto uint8

const (
	IP_API_PROTO_HOPOPT   IPProto = 0
	IP_API_PROTO_ICMP     IPProto = 1
	IP_API_PROTO_IGMP     IPProto = 2
	IP_API_PROTO_TCP      IPProto = 6
	IP_API_PROTO_UDP      IPProto = 17
	IP_API_PROTO_GRE      IPProto = 47
	IP_API_PROTO_ESP      IPProto = 50
	IP_API_PROTO_AH       IPProto = 51
	IP_API_PROTO_ICMP6    IPProto = 58
	IP_API_PROTO_EIGRP    IPProto = 88
	IP_API_PROTO_OSPF     IPProto = 89
	IP_API_PROTO_SCTP     IPProto = 132
	IP_API_PROTO_RESERVED IPProto = 255
)

var (
	IPProto_name = map[uint8]string{
		0:   "IP_API_PROTO_HOPOPT",
		1:   "IP_API_PROTO_ICMP",
		2:   "IP_API_PROTO_IGMP",
		6:   "IP_API_PROTO_TCP",
		17:  "IP_API_PROTO_UDP",
		47:  "IP_API_PROTO_GRE",
		50:  "IP_API_PROTO_ESP",
		51:  "IP_API_PROTO_AH",
		58:  "IP_API_PROTO_ICMP6",
		88:  "IP_API_PROTO_EIGRP",
		89:  "IP_API_PROTO_OSPF",
		132: "IP_API_PROTO_SCTP",
		255: "IP_API_PROTO_RESERVED",
	}
	IPProto_value = map[string]uint8{
		"IP_API_PROTO_HOPOPT":   0,
		"IP_API_PROTO_ICMP":     1,
		"IP_API_PROTO_IGMP":     2,
		"IP_API_PROTO_TCP":      6,
		"IP_API_PROTO_UDP":      17,
		"IP_API_PROTO_GRE":      47,
		"IP_API_PROTO_ESP":      50,
		"IP_API_PROTO_AH":       51,
		"IP_API_PROTO_ICMP6":    58,
		"IP_API_PROTO_EIGRP":    88,
		"IP_API_PROTO_OSPF":     89,
		"IP_API_PROTO_SCTP":     132,
		"IP_API_PROTO_RESERVED": 255,
	}
)

func (x IPProto) String() string {
	s, ok := IPProto_name[uint8(x)]
	if ok {
		return s
	}
	return "IPProto(" + strconv.Itoa(int(x)) + ")"
}

// AddressWithPrefix defines alias 'address_with_prefix'.
type AddressWithPrefix Prefix

func NewAddressWithPrefix(network net.IPNet) AddressWithPrefix {
	prefix := NewPrefix(network)
	return AddressWithPrefix(prefix)
}

func ParseAddressWithPrefix(s string) (AddressWithPrefix, error) {
	prefix, err := ParsePrefix(s)
	if err != nil {
		return AddressWithPrefix{}, err
	}
	return AddressWithPrefix(prefix), nil
}

func (x AddressWithPrefix) ToIPNet() *net.IPNet {
	return Prefix(x).ToIPNet()
}

func (x AddressWithPrefix) String() string {
	return Prefix(x).String()
}

func (x *AddressWithPrefix) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *AddressWithPrefix) UnmarshalText(text []byte) error {
	prefix, err := ParseAddressWithPrefix(string(text))
	if err != nil {
		return err
	}
	*x = prefix
	return nil
}

// IP4Address defines alias 'ip4_address'.
type IP4Address [4]uint8

func NewIP4Address(ip net.IP) IP4Address {
	var ipaddr IP4Address
	copy(ipaddr[:], ip.To4())
	return ipaddr
}

func ParseIP4Address(s string) (IP4Address, error) {
	ip := net.ParseIP(s).To4()
	if ip == nil {
		return IP4Address{}, fmt.Errorf("invalid IP4 address: %s", s)
	}
	var ipaddr IP4Address
	copy(ipaddr[:], ip.To4())
	return ipaddr, nil
}

func (x IP4Address) ToIP() net.IP {
	return net.IP(x[:]).To4()
}

func (x IP4Address) String() string {
	return x.ToIP().String()
}

func (x *IP4Address) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *IP4Address) UnmarshalText(text []byte) error {
	ipaddr, err := ParseIP4Address(string(text))
	if err != nil {
		return err
	}
	*x = ipaddr
	return nil
}

// IP4AddressWithPrefix defines alias 'ip4_address_with_prefix'.
type IP4AddressWithPrefix IP4Prefix

// IP6Address defines alias 'ip6_address'.
type IP6Address [16]uint8

func NewIP6Address(ip net.IP) IP6Address {
	var ipaddr IP6Address
	copy(ipaddr[:], ip.To16())
	return ipaddr
}

func ParseIP6Address(s string) (IP6Address, error) {
	ip := net.ParseIP(s).To16()
	if ip == nil {
		return IP6Address{}, fmt.Errorf("invalid IP6 address: %s", s)
	}
	var ipaddr IP6Address
	copy(ipaddr[:], ip.To16())
	return ipaddr, nil
}

func (x IP6Address) ToIP() net.IP {
	return net.IP(x[:]).To16()
}

func (x IP6Address) String() string {
	return x.ToIP().String()
}

func (x *IP6Address) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *IP6Address) UnmarshalText(text []byte) error {
	ipaddr, err := ParseIP6Address(string(text))
	if err != nil {
		return err
	}
	*x = ipaddr
	return nil
}

// IP6AddressWithPrefix defines alias 'ip6_address_with_prefix'.
type IP6AddressWithPrefix IP6Prefix

// Address defines type 'address'.
type Address struct {
	Af AddressFamily `binapi:"address_family,name=af" json:"af,omitempty"`
	Un AddressUnion  `binapi:"address_union,name=un" json:"un,omitempty"`
}

func NewAddress(ip net.IP) Address {
	var addr Address
	if ip.To4() == nil {
		addr.Af = ADDRESS_IP6
		var ip6 IP6Address
		copy(ip6[:], ip.To16())
		addr.Un.SetIP6(ip6)
	} else {
		addr.Af = ADDRESS_IP4
		var ip4 IP4Address
		copy(ip4[:], ip.To4())
		addr.Un.SetIP4(ip4)
	}
	return addr
}

func ParseAddress(s string) (Address, error) {
	ip := net.ParseIP(s)
	if ip == nil {
		return Address{}, fmt.Errorf("invalid IP address: %s", s)
	}
	return NewAddress(ip), nil
}

func (x Address) ToIP() net.IP {
	if x.Af == ADDRESS_IP6 {
		ip6 := x.Un.GetIP6()
		return net.IP(ip6[:]).To16()
	} else {
		ip4 := x.Un.GetIP4()
		return net.IP(ip4[:]).To4()
	}
}

func (x Address) String() string {
	return x.ToIP().String()
}

func (x *Address) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *Address) UnmarshalText(text []byte) error {
	addr, err := ParseAddress(string(text))
	if err != nil {
		return err
	}
	*x = addr
	return nil
}

// IP4AddressAndMask defines type 'ip4_address_and_mask'.
type IP4AddressAndMask struct {
	Addr IP4Address `binapi:"ip4_address,name=addr" json:"addr,omitempty"`
	Mask IP4Address `binapi:"ip4_address,name=mask" json:"mask,omitempty"`
}

// IP4Prefix defines type 'ip4_prefix'.
type IP4Prefix struct {
	Address IP4Address `binapi:"ip4_address,name=address" json:"address,omitempty"`
	Len     uint8      `binapi:"u8,name=len" json:"len,omitempty"`
}

func NewIP4Prefix(network net.IPNet) IP4Prefix {
	var prefix IP4Prefix
	maskSize, _ := network.Mask.Size()
	prefix.Len = byte(maskSize)
	prefix.Address = NewIP4Address(network.IP)
	return prefix
}

func ParseIP4Prefix(s string) (prefix IP4Prefix, err error) {
	hasPrefix := strings.Contains(s, "/")
	if hasPrefix {
		ip, network, err := net.ParseCIDR(s)
		if err != nil {
			return IP4Prefix{}, fmt.Errorf("invalid IP4 %s: %s", s, err)
		}
		maskSize, _ := network.Mask.Size()
		prefix.Len = byte(maskSize)
		prefix.Address, err = ParseIP4Address(ip.String())
		if err != nil {
			return IP4Prefix{}, fmt.Errorf("invalid IP4 %s: %s", s, err)
		}
	} else {
		ip := net.ParseIP(s)
		defaultMaskSize, _ := net.CIDRMask(32, 32).Size()
		if ip.To4() == nil {
			defaultMaskSize, _ = net.CIDRMask(128, 128).Size()
		}
		prefix.Len = byte(defaultMaskSize)
		prefix.Address, err = ParseIP4Address(ip.String())
		if err != nil {
			return IP4Prefix{}, fmt.Errorf("invalid IP4 %s: %s", s, err)
		}
	}
	return prefix, nil
}

func (x IP4Prefix) ToIPNet() *net.IPNet {
	mask := net.CIDRMask(int(x.Len), 32)
	ipnet := &net.IPNet{IP: x.Address.ToIP(), Mask: mask}
	return ipnet
}

func (x IP4Prefix) String() string {
	ip := x.Address.String()
	return ip + "/" + strconv.Itoa(int(x.Len))
}

func (x *IP4Prefix) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *IP4Prefix) UnmarshalText(text []byte) error {
	prefix, err := ParseIP4Prefix(string(text))
	if err != nil {
		return err
	}
	*x = prefix
	return nil
}

// IP6AddressAndMask defines type 'ip6_address_and_mask'.
type IP6AddressAndMask struct {
	Addr IP6Address `binapi:"ip6_address,name=addr" json:"addr,omitempty"`
	Mask IP6Address `binapi:"ip6_address,name=mask" json:"mask,omitempty"`
}

// IP6Prefix defines type 'ip6_prefix'.
type IP6Prefix struct {
	Address IP6Address `binapi:"ip6_address,name=address" json:"address,omitempty"`
	Len     uint8      `binapi:"u8,name=len" json:"len,omitempty"`
}

func NewIP6Prefix(network net.IPNet) IP6Prefix {
	var prefix IP6Prefix
	maskSize, _ := network.Mask.Size()
	prefix.Len = byte(maskSize)
	prefix.Address = NewIP6Address(network.IP)
	return prefix
}

func ParseIP6Prefix(s string) (prefix IP6Prefix, err error) {
	hasPrefix := strings.Contains(s, "/")
	if hasPrefix {
		ip, network, err := net.ParseCIDR(s)
		if err != nil {
			return IP6Prefix{}, fmt.Errorf("invalid IP6 %s: %s", s, err)
		}
		maskSize, _ := network.Mask.Size()
		prefix.Len = byte(maskSize)
		prefix.Address, err = ParseIP6Address(ip.String())
		if err != nil {
			return IP6Prefix{}, fmt.Errorf("invalid IP6 %s: %s", s, err)
		}
	} else {
		ip := net.ParseIP(s)
		defaultMaskSize, _ := net.CIDRMask(32, 32).Size()
		if ip.To4() == nil {
			defaultMaskSize, _ = net.CIDRMask(128, 128).Size()
		}
		prefix.Len = byte(defaultMaskSize)
		prefix.Address, err = ParseIP6Address(ip.String())
		if err != nil {
			return IP6Prefix{}, fmt.Errorf("invalid IP6 %s: %s", s, err)
		}
	}
	return prefix, nil
}

func (x IP6Prefix) ToIPNet() *net.IPNet {
	mask := net.CIDRMask(int(x.Len), 128)
	ipnet := &net.IPNet{IP: x.Address.ToIP(), Mask: mask}
	return ipnet
}

func (x IP6Prefix) String() string {
	ip := x.Address.String()
	return ip + "/" + strconv.Itoa(int(x.Len))
}

func (x *IP6Prefix) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *IP6Prefix) UnmarshalText(text []byte) error {
	prefix, err := ParseIP6Prefix(string(text))
	if err != nil {
		return err
	}
	*x = prefix
	return nil
}

// Mprefix defines type 'mprefix'.
type Mprefix struct {
	Af               AddressFamily `binapi:"address_family,name=af" json:"af,omitempty"`
	GrpAddressLength uint16        `binapi:"u16,name=grp_address_length" json:"grp_address_length,omitempty"`
	GrpAddress       AddressUnion  `binapi:"address_union,name=grp_address" json:"grp_address,omitempty"`
	SrcAddress       AddressUnion  `binapi:"address_union,name=src_address" json:"src_address,omitempty"`
}

// Prefix defines type 'prefix'.
type Prefix struct {
	Address Address `binapi:"address,name=address" json:"address,omitempty"`
	Len     uint8   `binapi:"u8,name=len" json:"len,omitempty"`
}

func NewPrefix(network net.IPNet) Prefix {
	var prefix Prefix
	maskSize, _ := network.Mask.Size()
	prefix.Len = byte(maskSize)
	prefix.Address = NewAddress(network.IP)
	return prefix
}

func ParsePrefix(ip string) (prefix Prefix, err error) {
	hasPrefix := strings.Contains(ip, "/")
	if hasPrefix {
		netIP, network, err := net.ParseCIDR(ip)
		if err != nil {
			return Prefix{}, fmt.Errorf("invalid IP %s: %s", ip, err)
		}
		maskSize, _ := network.Mask.Size()
		prefix.Len = byte(maskSize)
		prefix.Address, err = ParseAddress(netIP.String())
		if err != nil {
			return Prefix{}, fmt.Errorf("invalid IP %s: %s", ip, err)
		}
	} else {
		netIP := net.ParseIP(ip)
		defaultMaskSize, _ := net.CIDRMask(32, 32).Size()
		if netIP.To4() == nil {
			defaultMaskSize, _ = net.CIDRMask(128, 128).Size()
		}
		prefix.Len = byte(defaultMaskSize)
		prefix.Address, err = ParseAddress(netIP.String())
		if err != nil {
			return Prefix{}, fmt.Errorf("invalid IP %s: %s", ip, err)
		}
	}
	return prefix, nil
}

func (x Prefix) ToIPNet() *net.IPNet {
	var mask net.IPMask
	if x.Address.Af == ADDRESS_IP4 {
		mask = net.CIDRMask(int(x.Len), 32)
	} else {
		mask = net.CIDRMask(int(x.Len), 128)
	}
	ipnet := &net.IPNet{IP: x.Address.ToIP(), Mask: mask}
	return ipnet
}

func (x Prefix) String() string {
	ip := x.Address.String()
	return ip + "/" + strconv.Itoa(int(x.Len))
}

func (x *Prefix) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *Prefix) UnmarshalText(text []byte) error {
	prefix, err := ParsePrefix(string(text))
	if err != nil {
		return err
	}
	*x = prefix
	return nil
}

// PrefixMatcher defines type 'prefix_matcher'.
type PrefixMatcher struct {
	Le uint8 `binapi:"u8,name=le" json:"le,omitempty"`
	Ge uint8 `binapi:"u8,name=ge" json:"ge,omitempty"`
}

// AddressUnion defines union 'address_union'.
type AddressUnion struct {
	// AddressUnion can be one of:
	// - IP4 *IP4Address
	// - IP6 *IP6Address
	XXX_UnionData [16]byte
}

func AddressUnionIP4(a IP4Address) (u AddressUnion) {
	u.SetIP4(a)
	return
}
func (u *AddressUnion) SetIP4(a IP4Address) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeBytes(a[:], 4)
}
func (u *AddressUnion) GetIP4() (a IP4Address) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	copy(a[:], buf.DecodeBytes(4))
	return
}

func AddressUnionIP6(a IP6Address) (u AddressUnion) {
	u.SetIP6(a)
	return
}
func (u *AddressUnion) SetIP6(a IP6Address) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeBytes(a[:], 16)
}
func (u *AddressUnion) GetIP6() (a IP6Address) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	copy(a[:], buf.DecodeBytes(16))
	return
}
