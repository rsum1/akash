// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/market/v1beta1/lease.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// State is an enum which refers to state of lease
type Lease_State int32

const (
	// Prefix should start with 0 in enum. So declaring dummy state
	LeaseStateInvalid Lease_State = 0
	// LeaseActive denotes state for lease active
	LeaseActive Lease_State = 1
	// LeaseInsufficientFunds denotes state for lease insufficient_funds
	LeaseInsufficientFunds Lease_State = 2
	// LeaseClosed denotes state for lease closed
	LeaseClosed Lease_State = 3
)

var Lease_State_name = map[int32]string{
	0: "invalid",
	1: "active",
	2: "insufficient_funds",
	3: "closed",
}

var Lease_State_value = map[string]int32{
	"invalid":            0,
	"active":             1,
	"insufficient_funds": 2,
	"closed":             3,
}

func (x Lease_State) String() string {
	return proto.EnumName(Lease_State_name, int32(x))
}

func (Lease_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b81e11575e79ba08, []int{1, 0}
}

// LeaseID stores bid details of lease
type LeaseID struct {
	Owner    github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=owner,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"owner" yaml:"owner"`
	DSeq     uint64                                        `protobuf:"varint,2,opt,name=dseq,proto3" json:"dseq" yaml:"dseq"`
	GSeq     uint32                                        `protobuf:"varint,3,opt,name=gseq,proto3" json:"gseq" yaml:"gseq"`
	OSeq     uint32                                        `protobuf:"varint,4,opt,name=oseq,proto3" json:"oseq" yaml:"oseq"`
	Provider github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,5,opt,name=provider,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"provider" yaml:"provider"`
}

func (m *LeaseID) Reset()      { *m = LeaseID{} }
func (*LeaseID) ProtoMessage() {}
func (*LeaseID) Descriptor() ([]byte, []int) {
	return fileDescriptor_b81e11575e79ba08, []int{0}
}
func (m *LeaseID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LeaseID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaseID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LeaseID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseID.Merge(m, src)
}
func (m *LeaseID) XXX_Size() int {
	return m.Size()
}
func (m *LeaseID) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseID.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseID proto.InternalMessageInfo

func (m *LeaseID) GetOwner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *LeaseID) GetDSeq() uint64 {
	if m != nil {
		return m.DSeq
	}
	return 0
}

func (m *LeaseID) GetGSeq() uint32 {
	if m != nil {
		return m.GSeq
	}
	return 0
}

func (m *LeaseID) GetOSeq() uint32 {
	if m != nil {
		return m.OSeq
	}
	return 0
}

func (m *LeaseID) GetProvider() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Provider
	}
	return nil
}

// Lease stores LeaseID, state of lease and price
type Lease struct {
	LeaseID LeaseID     `protobuf:"bytes,1,opt,name=lease_id,json=leaseId,proto3" json:"id" yaml:"id"`
	State   Lease_State `protobuf:"varint,2,opt,name=state,proto3,enum=akash.market.v1beta1.Lease_State" json:"state" yaml:"state"`
	Price   types.Coin  `protobuf:"bytes,3,opt,name=price,proto3" json:"price" yaml:"price"`
}

func (m *Lease) Reset()      { *m = Lease{} }
func (*Lease) ProtoMessage() {}
func (*Lease) Descriptor() ([]byte, []int) {
	return fileDescriptor_b81e11575e79ba08, []int{1}
}
func (m *Lease) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Lease) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Lease.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Lease) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lease.Merge(m, src)
}
func (m *Lease) XXX_Size() int {
	return m.Size()
}
func (m *Lease) XXX_DiscardUnknown() {
	xxx_messageInfo_Lease.DiscardUnknown(m)
}

var xxx_messageInfo_Lease proto.InternalMessageInfo

func (m *Lease) GetLeaseID() LeaseID {
	if m != nil {
		return m.LeaseID
	}
	return LeaseID{}
}

func (m *Lease) GetState() Lease_State {
	if m != nil {
		return m.State
	}
	return LeaseStateInvalid
}

func (m *Lease) GetPrice() types.Coin {
	if m != nil {
		return m.Price
	}
	return types.Coin{}
}

// LeaseFilters defines flags for lease list filter
type LeaseFilters struct {
	Owner    github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=owner,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"owner" yaml:"owner"`
	DSeq     uint64                                        `protobuf:"varint,2,opt,name=dseq,proto3" json:"dseq" yaml:"dseq"`
	GSeq     uint32                                        `protobuf:"varint,3,opt,name=gseq,proto3" json:"gseq" yaml:"gseq"`
	OSeq     uint32                                        `protobuf:"varint,4,opt,name=oseq,proto3" json:"oseq" yaml:"oseq"`
	Provider github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,5,opt,name=provider,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"provider" yaml:"provider"`
	State    string                                        `protobuf:"bytes,6,opt,name=state,proto3" json:"state" yaml:"state"`
}

func (m *LeaseFilters) Reset()         { *m = LeaseFilters{} }
func (m *LeaseFilters) String() string { return proto.CompactTextString(m) }
func (*LeaseFilters) ProtoMessage()    {}
func (*LeaseFilters) Descriptor() ([]byte, []int) {
	return fileDescriptor_b81e11575e79ba08, []int{2}
}
func (m *LeaseFilters) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LeaseFilters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaseFilters.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LeaseFilters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseFilters.Merge(m, src)
}
func (m *LeaseFilters) XXX_Size() int {
	return m.Size()
}
func (m *LeaseFilters) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseFilters.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseFilters proto.InternalMessageInfo

func (m *LeaseFilters) GetOwner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *LeaseFilters) GetDSeq() uint64 {
	if m != nil {
		return m.DSeq
	}
	return 0
}

func (m *LeaseFilters) GetGSeq() uint32 {
	if m != nil {
		return m.GSeq
	}
	return 0
}

func (m *LeaseFilters) GetOSeq() uint32 {
	if m != nil {
		return m.OSeq
	}
	return 0
}

func (m *LeaseFilters) GetProvider() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Provider
	}
	return nil
}

func (m *LeaseFilters) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func init() {
	proto.RegisterEnum("akash.market.v1beta1.Lease_State", Lease_State_name, Lease_State_value)
	proto.RegisterType((*LeaseID)(nil), "akash.market.v1beta1.LeaseID")
	proto.RegisterType((*Lease)(nil), "akash.market.v1beta1.Lease")
	proto.RegisterType((*LeaseFilters)(nil), "akash.market.v1beta1.LeaseFilters")
}

func init() { proto.RegisterFile("akash/market/v1beta1/lease.proto", fileDescriptor_b81e11575e79ba08) }

var fileDescriptor_b81e11575e79ba08 = []byte{
	// 636 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x54, 0x3f, 0x6f, 0xd3, 0x4e,
	0x18, 0xb6, 0x9b, 0xa4, 0x7f, 0x2e, 0xfd, 0xfd, 0x1a, 0xac, 0x82, 0x5a, 0xa3, 0xfa, 0x8c, 0x07,
	0x94, 0xa5, 0xb6, 0x9a, 0x6e, 0x9d, 0x68, 0x5a, 0x15, 0x15, 0x21, 0x40, 0xe9, 0x86, 0x90, 0x8a,
	0xe3, 0xbb, 0xba, 0xa7, 0x38, 0xbe, 0xd4, 0xe7, 0x06, 0xba, 0x32, 0xa1, 0x4e, 0x8c, 0x2c, 0x15,
	0x95, 0xf8, 0x0a, 0xf0, 0x1d, 0x3a, 0x76, 0x64, 0x3a, 0xa1, 0x74, 0x41, 0x19, 0x33, 0x32, 0xa1,
	0x7b, 0x2f, 0x49, 0x53, 0x09, 0x81, 0xc4, 0xc0, 0xc4, 0xe4, 0xdc, 0xf3, 0x3e, 0xcf, 0x73, 0x6f,
	0xee, 0x79, 0xef, 0x90, 0x1b, 0xb6, 0x42, 0x71, 0x18, 0xb4, 0xc3, 0xac, 0x45, 0xf3, 0xa0, 0xbb,
	0xd6, 0xa4, 0x79, 0xb8, 0x16, 0x24, 0x34, 0x14, 0xd4, 0xef, 0x64, 0x3c, 0xe7, 0xd6, 0x22, 0x30,
	0x7c, 0xcd, 0xf0, 0x87, 0x0c, 0x7b, 0x31, 0xe6, 0x31, 0x07, 0x42, 0xa0, 0x7e, 0x69, 0xae, 0xed,
	0x44, 0x5c, 0xb4, 0xb9, 0x08, 0x9a, 0xa1, 0xa0, 0x63, 0xb3, 0x88, 0xb3, 0x54, 0xd7, 0xbd, 0x37,
	0x05, 0x34, 0xf3, 0x58, 0x79, 0xef, 0x6e, 0x5b, 0x2f, 0x51, 0x89, 0xbf, 0x4a, 0x69, 0xb6, 0x64,
	0xba, 0x66, 0x75, 0xbe, 0xfe, 0xa8, 0x2f, 0xb1, 0x06, 0x06, 0x12, 0xcf, 0x9f, 0x84, 0xed, 0x64,
	0xc3, 0x83, 0xa5, 0xf7, 0x5d, 0xe2, 0xd5, 0x98, 0xe5, 0x87, 0xc7, 0x4d, 0x3f, 0xe2, 0xed, 0x60,
	0xb8, 0x85, 0xfe, 0xac, 0x0a, 0xd2, 0x0a, 0xf2, 0x93, 0x0e, 0x15, 0xfe, 0x66, 0x14, 0x6d, 0x12,
	0x92, 0x51, 0x21, 0x1a, 0xda, 0xc7, 0x5a, 0x47, 0x45, 0x22, 0xe8, 0xd1, 0xd2, 0x94, 0x6b, 0x56,
	0x8b, 0x75, 0xdc, 0x93, 0xb8, 0xb8, 0xbd, 0x47, 0x8f, 0xfa, 0x12, 0x03, 0x3e, 0x90, 0xb8, 0xac,
	0xf7, 0x51, 0x2b, 0xaf, 0x01, 0xa0, 0x12, 0xc5, 0x4a, 0x54, 0x70, 0xcd, 0xea, 0x7f, 0x5a, 0xf4,
	0x70, 0x28, 0x8a, 0x6f, 0x88, 0x62, 0x2d, 0x8a, 0x87, 0x22, 0xae, 0x44, 0xc5, 0x6b, 0xd1, 0xd3,
	0xa1, 0x88, 0xdf, 0x10, 0x71, 0x2d, 0x52, 0x1f, 0x2b, 0x41, 0xb3, 0x9d, 0x8c, 0x77, 0x19, 0xa1,
	0xd9, 0x52, 0x09, 0xce, 0xe0, 0x59, 0x5f, 0xe2, 0x31, 0x36, 0x90, 0x78, 0x41, 0x8b, 0x46, 0xc8,
	0x1f, 0x9c, 0xc4, 0xd8, 0x6d, 0x63, 0xf6, 0xfd, 0x39, 0x36, 0xbe, 0x9d, 0x63, 0xd3, 0xfb, 0x5c,
	0x40, 0x25, 0x08, 0xc1, 0x7a, 0x81, 0x66, 0x21, 0xe9, 0x7d, 0x46, 0x20, 0x85, 0x72, 0x6d, 0xc5,
	0xff, 0x59, 0xda, 0xfe, 0x30, 0xb3, 0xba, 0x77, 0x21, 0xb1, 0xd1, 0x93, 0x78, 0x14, 0x62, 0x5f,
	0xe2, 0x29, 0x46, 0x06, 0x12, 0xcf, 0xe9, 0x4e, 0x19, 0xf1, 0x1a, 0x33, 0x60, 0xb9, 0x4b, 0xac,
	0x06, 0x2a, 0x89, 0x3c, 0xcc, 0x29, 0x9c, 0xff, 0xff, 0xb5, 0x7b, 0xbf, 0xb0, 0xf6, 0xf7, 0x14,
	0xb1, 0xbe, 0xac, 0x66, 0x00, 0x34, 0xd7, 0x33, 0x00, 0x4b, 0xaf, 0xa1, 0x61, 0xeb, 0x09, 0x2a,
	0x75, 0x32, 0x16, 0x51, 0x88, 0xa7, 0x5c, 0x5b, 0xf6, 0xf5, 0x9f, 0xf7, 0xd5, 0xc0, 0x8d, 0x2d,
	0xb7, 0x38, 0x4b, 0xeb, 0x2b, 0xaa, 0x55, 0xe5, 0x07, 0xfc, 0x6b, 0x3f, 0x58, 0x7a, 0x0d, 0x0d,
	0x7b, 0x1f, 0x4c, 0x54, 0x82, 0xbd, 0x2d, 0x0f, 0xcd, 0xb0, 0xb4, 0x1b, 0x26, 0x8c, 0x54, 0x0c,
	0xfb, 0xf6, 0xe9, 0x99, 0x7b, 0x0b, 0x3a, 0x83, 0xe2, 0xae, 0x2e, 0x58, 0x77, 0xd1, 0x74, 0x18,
	0xe5, 0xac, 0x4b, 0x2b, 0xa6, 0xbd, 0x70, 0x7a, 0xe6, 0x96, 0x81, 0xb2, 0x09, 0x90, 0x55, 0x43,
	0x16, 0x4b, 0xc5, 0xf1, 0xc1, 0x01, 0x8b, 0x18, 0x4d, 0xf3, 0xfd, 0x83, 0xe3, 0x94, 0x88, 0xca,
	0x94, 0x6d, 0x9f, 0x9e, 0xb9, 0x77, 0xf4, 0x79, 0x4d, 0x94, 0x77, 0x54, 0x55, 0x19, 0x46, 0x09,
	0x17, 0x94, 0x54, 0x0a, 0x13, 0x86, 0x5b, 0x00, 0xd9, 0xc5, 0xb7, 0x1f, 0x1d, 0x63, 0x9c, 0x9b,
	0xe1, 0x7d, 0x2a, 0xa0, 0x79, 0xa8, 0xef, 0xb0, 0x24, 0xa7, 0x99, 0xf8, 0x77, 0x83, 0xfe, 0xea,
	0x0d, 0xb2, 0x82, 0xd1, 0x3c, 0x4f, 0xbb, 0x66, 0x75, 0xee, 0xf7, 0xc3, 0xba, 0x51, 0x54, 0xb1,
	0xd5, 0x1f, 0x5c, 0xf4, 0x1c, 0xf3, 0xb2, 0xe7, 0x98, 0x5f, 0x7b, 0x8e, 0xf9, 0xee, 0xca, 0x31,
	0x2e, 0xaf, 0x1c, 0xe3, 0xcb, 0x95, 0x63, 0x3c, 0xbf, 0x3f, 0xd1, 0x09, 0xef, 0x66, 0x51, 0xd2,
	0x0a, 0xf4, 0x6b, 0xfc, 0x7a, 0xf4, 0x1e, 0x43, 0x37, 0xcd, 0x69, 0x78, 0x3c, 0xd7, 0x7f, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xc8, 0x78, 0x00, 0xc3, 0xac, 0x05, 0x00, 0x00,
}

func (this *LeaseID) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LeaseID)
	if !ok {
		that2, ok := that.(LeaseID)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Owner, that1.Owner) {
		return false
	}
	if this.DSeq != that1.DSeq {
		return false
	}
	if this.GSeq != that1.GSeq {
		return false
	}
	if this.OSeq != that1.OSeq {
		return false
	}
	if !bytes.Equal(this.Provider, that1.Provider) {
		return false
	}
	return true
}
func (m *LeaseID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintLease(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x2a
	}
	if m.OSeq != 0 {
		i = encodeVarintLease(dAtA, i, uint64(m.OSeq))
		i--
		dAtA[i] = 0x20
	}
	if m.GSeq != 0 {
		i = encodeVarintLease(dAtA, i, uint64(m.GSeq))
		i--
		dAtA[i] = 0x18
	}
	if m.DSeq != 0 {
		i = encodeVarintLease(dAtA, i, uint64(m.DSeq))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintLease(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Lease) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Lease) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Lease) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLease(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.State != 0 {
		i = encodeVarintLease(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.LeaseID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLease(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *LeaseFilters) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseFilters) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseFilters) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintLease(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintLease(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x2a
	}
	if m.OSeq != 0 {
		i = encodeVarintLease(dAtA, i, uint64(m.OSeq))
		i--
		dAtA[i] = 0x20
	}
	if m.GSeq != 0 {
		i = encodeVarintLease(dAtA, i, uint64(m.GSeq))
		i--
		dAtA[i] = 0x18
	}
	if m.DSeq != 0 {
		i = encodeVarintLease(dAtA, i, uint64(m.DSeq))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintLease(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLease(dAtA []byte, offset int, v uint64) int {
	offset -= sovLease(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LeaseID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovLease(uint64(l))
	}
	if m.DSeq != 0 {
		n += 1 + sovLease(uint64(m.DSeq))
	}
	if m.GSeq != 0 {
		n += 1 + sovLease(uint64(m.GSeq))
	}
	if m.OSeq != 0 {
		n += 1 + sovLease(uint64(m.OSeq))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovLease(uint64(l))
	}
	return n
}

func (m *Lease) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.LeaseID.Size()
	n += 1 + l + sovLease(uint64(l))
	if m.State != 0 {
		n += 1 + sovLease(uint64(m.State))
	}
	l = m.Price.Size()
	n += 1 + l + sovLease(uint64(l))
	return n
}

func (m *LeaseFilters) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovLease(uint64(l))
	}
	if m.DSeq != 0 {
		n += 1 + sovLease(uint64(m.DSeq))
	}
	if m.GSeq != 0 {
		n += 1 + sovLease(uint64(m.GSeq))
	}
	if m.OSeq != 0 {
		n += 1 + sovLease(uint64(m.OSeq))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovLease(uint64(l))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovLease(uint64(l))
	}
	return n
}

func sovLease(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLease(x uint64) (n int) {
	return sovLease(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LeaseID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLease
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LeaseID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLease
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLease
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DSeq", wireType)
			}
			m.DSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DSeq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GSeq", wireType)
			}
			m.GSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GSeq |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OSeq", wireType)
			}
			m.OSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OSeq |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLease
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLease
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = append(m.Provider[:0], dAtA[iNdEx:postIndex]...)
			if m.Provider == nil {
				m.Provider = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLease(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLease
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLease
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Lease) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLease
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Lease: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Lease: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaseID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLease
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLease
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LeaseID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= Lease_State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLease
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLease
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLease(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLease
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLease
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LeaseFilters) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLease
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LeaseFilters: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseFilters: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLease
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLease
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DSeq", wireType)
			}
			m.DSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DSeq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GSeq", wireType)
			}
			m.GSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GSeq |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OSeq", wireType)
			}
			m.OSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OSeq |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLease
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLease
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = append(m.Provider[:0], dAtA[iNdEx:postIndex]...)
			if m.Provider == nil {
				m.Provider = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLease
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLease
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLease(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLease
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLease
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLease(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLease
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLease
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLease
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthLease
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLease
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLease
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLease        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLease          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLease = fmt.Errorf("proto: unexpected end of group")
)