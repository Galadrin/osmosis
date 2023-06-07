// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/gamm/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types1 "github.com/cosmos/cosmos-sdk/codec/types"
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

// Params holds parameters for the incentives module
type Params struct {
	PoolCreationFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=pool_creation_fee,json=poolCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"pool_creation_fee" yaml:"pool_creation_fee"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a324eb7f1dd793e, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetPoolCreationFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.PoolCreationFee
	}
	return nil
}

// GenesisState defines the gamm module's genesis state.
type GenesisState struct {
	Pools []*types1.Any `protobuf:"bytes,1,rep,name=pools,proto3" json:"pools,omitempty"`
	// will be renamed to next_pool_id in an upcoming version
	NextPoolNumber   uint64            `protobuf:"varint,2,opt,name=next_pool_number,json=nextPoolNumber,proto3" json:"next_pool_number,omitempty"`
	Params           Params            `protobuf:"bytes,3,opt,name=params,proto3" json:"params"`
	MigrationRecords *MigrationRecords `protobuf:"bytes,4,opt,name=migration_records,json=migrationRecords,proto3" json:"migration_records,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a324eb7f1dd793e, []int{1}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetPools() []*types1.Any {
	if m != nil {
		return m.Pools
	}
	return nil
}

func (m *GenesisState) GetNextPoolNumber() uint64 {
	if m != nil {
		return m.NextPoolNumber
	}
	return 0
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetMigrationRecords() *MigrationRecords {
	if m != nil {
		return m.MigrationRecords
	}
	return nil
}

// MigrationRecords contains all the links between balancer and concentrated
// pools
type MigrationRecords struct {
	BalancerToConcentratedPoolLinks []BalancerToConcentratedPoolLink `protobuf:"bytes,1,rep,name=balancer_to_concentrated_pool_links,json=balancerToConcentratedPoolLinks,proto3" json:"balancer_to_concentrated_pool_links"`
}

func (m *MigrationRecords) Reset()         { *m = MigrationRecords{} }
func (m *MigrationRecords) String() string { return proto.CompactTextString(m) }
func (*MigrationRecords) ProtoMessage()    {}
func (*MigrationRecords) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a324eb7f1dd793e, []int{2}
}
func (m *MigrationRecords) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MigrationRecords) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MigrationRecords.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MigrationRecords) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MigrationRecords.Merge(m, src)
}
func (m *MigrationRecords) XXX_Size() int {
	return m.Size()
}
func (m *MigrationRecords) XXX_DiscardUnknown() {
	xxx_messageInfo_MigrationRecords.DiscardUnknown(m)
}

var xxx_messageInfo_MigrationRecords proto.InternalMessageInfo

func (m *MigrationRecords) GetBalancerToConcentratedPoolLinks() []BalancerToConcentratedPoolLink {
	if m != nil {
		return m.BalancerToConcentratedPoolLinks
	}
	return nil
}

// BalancerToConcentratedPoolLink defines a single link between a single
// balancer pool and a single concentrated liquidity pool. This link is used to
// allow a balancer pool to migrate to a single canonical full range
// concentrated liquidity pool position
// A balancer pool can be linked to a maximum of one cl pool, and a cl pool can
// be linked to a maximum of one balancer pool.
type BalancerToConcentratedPoolLink struct {
	BalancerPoolId uint64 `protobuf:"varint,1,opt,name=balancer_pool_id,json=balancerPoolId,proto3" json:"balancer_pool_id,omitempty"`
	ClPoolId       uint64 `protobuf:"varint,2,opt,name=cl_pool_id,json=clPoolId,proto3" json:"cl_pool_id,omitempty"`
}

func (m *BalancerToConcentratedPoolLink) Reset()         { *m = BalancerToConcentratedPoolLink{} }
func (m *BalancerToConcentratedPoolLink) String() string { return proto.CompactTextString(m) }
func (*BalancerToConcentratedPoolLink) ProtoMessage()    {}
func (*BalancerToConcentratedPoolLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a324eb7f1dd793e, []int{3}
}
func (m *BalancerToConcentratedPoolLink) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BalancerToConcentratedPoolLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BalancerToConcentratedPoolLink.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BalancerToConcentratedPoolLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalancerToConcentratedPoolLink.Merge(m, src)
}
func (m *BalancerToConcentratedPoolLink) XXX_Size() int {
	return m.Size()
}
func (m *BalancerToConcentratedPoolLink) XXX_DiscardUnknown() {
	xxx_messageInfo_BalancerToConcentratedPoolLink.DiscardUnknown(m)
}

var xxx_messageInfo_BalancerToConcentratedPoolLink proto.InternalMessageInfo

func (m *BalancerToConcentratedPoolLink) GetBalancerPoolId() uint64 {
	if m != nil {
		return m.BalancerPoolId
	}
	return 0
}

func (m *BalancerToConcentratedPoolLink) GetClPoolId() uint64 {
	if m != nil {
		return m.ClPoolId
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "osmosis.gamm.v1beta1.Params")
	proto.RegisterType((*GenesisState)(nil), "osmosis.gamm.v1beta1.GenesisState")
	proto.RegisterType((*MigrationRecords)(nil), "osmosis.gamm.v1beta1.MigrationRecords")
	proto.RegisterType((*BalancerToConcentratedPoolLink)(nil), "osmosis.gamm.v1beta1.BalancerToConcentratedPoolLink")
}

func init() {
	proto.RegisterFile("osmosis/gamm/v1beta1/genesis.proto", fileDescriptor_5a324eb7f1dd793e)
}

var fileDescriptor_5a324eb7f1dd793e = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x3f, 0x6f, 0xd3, 0x4e,
	0x18, 0xc7, 0x73, 0x6d, 0x1a, 0xfd, 0x7e, 0x57, 0x04, 0xa9, 0x95, 0xc1, 0xad, 0x2a, 0x27, 0x32,
	0x12, 0xca, 0x92, 0x33, 0x29, 0x7f, 0x86, 0x6c, 0x24, 0x12, 0x08, 0x54, 0x50, 0xe5, 0x32, 0xb1,
	0x58, 0x67, 0xe7, 0x6a, 0xac, 0xd8, 0xf7, 0x44, 0xbe, 0x4b, 0xd5, 0x6c, 0x8c, 0x8c, 0x48, 0xcc,
	0x20, 0x66, 0x66, 0x5e, 0x44, 0xc5, 0xd4, 0x91, 0xa9, 0xa0, 0x64, 0x61, 0xe6, 0x15, 0xa0, 0xfb,
	0xe3, 0xa8, 0x2a, 0x51, 0x27, 0xfb, 0x79, 0x9e, 0xcf, 0xf3, 0xf5, 0xf3, 0x7c, 0xef, 0x8c, 0x7d,
	0x10, 0x05, 0x88, 0x4c, 0x04, 0x29, 0x2d, 0x8a, 0xe0, 0xb4, 0x1f, 0x33, 0x49, 0xfb, 0x41, 0xca,
	0x38, 0x13, 0x99, 0x20, 0xd3, 0x12, 0x24, 0x38, 0x2d, 0xcb, 0x10, 0xc5, 0x10, 0xcb, 0xec, 0xb5,
	0x52, 0x48, 0x41, 0x03, 0x81, 0x7a, 0x33, 0xec, 0xde, 0x6e, 0x0a, 0x90, 0xe6, 0x2c, 0xd0, 0x51,
	0x3c, 0x3b, 0x09, 0x28, 0x9f, 0x57, 0xa5, 0x44, 0xeb, 0x44, 0xa6, 0xc7, 0x04, 0xb6, 0xe4, 0x99,
	0x28, 0x88, 0xa9, 0x60, 0xab, 0x21, 0x12, 0xc8, 0xb8, 0xa9, 0xfb, 0x9f, 0x11, 0x6e, 0x1c, 0xd1,
	0x92, 0x16, 0xc2, 0xf9, 0x88, 0xf0, 0xce, 0x14, 0x20, 0x8f, 0x92, 0x92, 0x51, 0x99, 0x01, 0x8f,
	0x4e, 0x18, 0x73, 0x51, 0x67, 0xb3, 0xbb, 0x7d, 0xb0, 0x4b, 0xac, 0xaa, 0xd2, 0xa9, 0x06, 0x25,
	0x23, 0xc8, 0xf8, 0xf0, 0xf0, 0xfc, 0xb2, 0x5d, 0xfb, 0x73, 0xd9, 0x76, 0xe7, 0xb4, 0xc8, 0x07,
	0xfe, 0x3f, 0x0a, 0xfe, 0xd7, 0x9f, 0xed, 0x6e, 0x9a, 0xc9, 0xb7, 0xb3, 0x98, 0x24, 0x50, 0xd8,
	0xf1, 0xec, 0xa3, 0x27, 0xc6, 0x93, 0x40, 0xce, 0xa7, 0x4c, 0x68, 0x31, 0x11, 0xde, 0x51, 0xfd,
	0x23, 0xdb, 0xfe, 0x94, 0x31, 0xff, 0xdd, 0x06, 0xbe, 0xf5, 0xcc, 0x98, 0x76, 0x2c, 0xa9, 0x64,
	0xce, 0x23, 0xbc, 0xa5, 0x18, 0x61, 0x27, 0x6b, 0x11, 0xe3, 0x0b, 0xa9, 0x7c, 0x21, 0x4f, 0xf8,
	0x7c, 0xf8, 0xff, 0xf7, 0x6f, 0xbd, 0xad, 0x23, 0x80, 0xfc, 0x79, 0x68, 0x68, 0xa7, 0x8b, 0x9b,
	0x9c, 0x9d, 0xc9, 0x48, 0xcf, 0xc7, 0x67, 0x45, 0xcc, 0x4a, 0x77, 0xa3, 0x83, 0xba, 0xf5, 0xf0,
	0xb6, 0xca, 0x2b, 0xf6, 0x95, 0xce, 0x3a, 0x03, 0xdc, 0x98, 0x6a, 0x47, 0xdc, 0xcd, 0x0e, 0xea,
	0x6e, 0x1f, 0xec, 0x93, 0x75, 0xa7, 0x44, 0x8c, 0x6b, 0xc3, 0xba, 0x5a, 0x3f, 0xb4, 0x1d, 0xce,
	0x31, 0xde, 0x29, 0xb2, 0xb4, 0x34, 0xcb, 0x97, 0x2c, 0x81, 0x72, 0x2c, 0xdc, 0xba, 0x96, 0xb9,
	0xb7, 0x5e, 0xe6, 0x65, 0x85, 0x87, 0x86, 0x0e, 0x9b, 0xc5, 0xb5, 0x8c, 0xff, 0x09, 0xe1, 0xe6,
	0x75, 0xcc, 0x79, 0x8f, 0xf0, 0xdd, 0x98, 0xe6, 0x94, 0x27, 0xac, 0x8c, 0x24, 0x44, 0x09, 0xf0,
	0x84, 0x71, 0x59, 0x52, 0xc9, 0xc6, 0x66, 0xc9, 0x3c, 0xe3, 0x93, 0xca, 0xa5, 0x87, 0xeb, 0x3f,
	0x3e, 0xb4, 0x02, 0xaf, 0x61, 0x74, 0xa5, 0x5d, 0x79, 0x71, 0x98, 0xf1, 0x89, 0xdd, 0xad, 0x1d,
	0xdf, 0x48, 0x09, 0x9f, 0x63, 0xef, 0x66, 0x21, 0x65, 0xfe, 0x6a, 0x56, 0x3d, 0x5b, 0x36, 0x76,
	0x91, 0x31, 0xbf, 0xca, 0xeb, 0xc3, 0x1a, 0x3b, 0xfb, 0x18, 0x27, 0xf9, 0x8a, 0x31, 0x07, 0xf4,
	0x5f, 0x92, 0x9b, 0xea, 0xa0, 0xfe, 0xfb, 0x4b, 0x1b, 0x0d, 0x5f, 0x9c, 0x2f, 0x3c, 0x74, 0xb1,
	0xf0, 0xd0, 0xaf, 0x85, 0x87, 0x3e, 0x2c, 0xbd, 0xda, 0xc5, 0xd2, 0xab, 0xfd, 0x58, 0x7a, 0xb5,
	0x37, 0xf7, 0xaf, 0xdc, 0x33, 0xbb, 0x70, 0x2f, 0xa7, 0xb1, 0xa8, 0x82, 0xe0, 0xb4, 0xff, 0x38,
	0x38, 0x33, 0x7f, 0xa4, 0xbe, 0x75, 0x71, 0x43, 0x5f, 0x9b, 0x07, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xa6, 0x5b, 0xa5, 0xbb, 0xae, 0x03, 0x00, 0x00,
}

func (this *BalancerToConcentratedPoolLink) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BalancerToConcentratedPoolLink)
	if !ok {
		that2, ok := that.(BalancerToConcentratedPoolLink)
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
	if this.BalancerPoolId != that1.BalancerPoolId {
		return false
	}
	if this.ClPoolId != that1.ClPoolId {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PoolCreationFee) > 0 {
		for iNdEx := len(m.PoolCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MigrationRecords != nil {
		{
			size, err := m.MigrationRecords.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.NextPoolNumber != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextPoolNumber))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Pools) > 0 {
		for iNdEx := len(m.Pools) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pools[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MigrationRecords) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MigrationRecords) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MigrationRecords) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BalancerToConcentratedPoolLinks) > 0 {
		for iNdEx := len(m.BalancerToConcentratedPoolLinks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BalancerToConcentratedPoolLinks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *BalancerToConcentratedPoolLink) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BalancerToConcentratedPoolLink) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BalancerToConcentratedPoolLink) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ClPoolId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ClPoolId))
		i--
		dAtA[i] = 0x10
	}
	if m.BalancerPoolId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.BalancerPoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PoolCreationFee) > 0 {
		for _, e := range m.PoolCreationFee {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Pools) > 0 {
		for _, e := range m.Pools {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.NextPoolNumber != 0 {
		n += 1 + sovGenesis(uint64(m.NextPoolNumber))
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.MigrationRecords != nil {
		l = m.MigrationRecords.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *MigrationRecords) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.BalancerToConcentratedPoolLinks) > 0 {
		for _, e := range m.BalancerToConcentratedPoolLinks {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *BalancerToConcentratedPoolLink) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BalancerPoolId != 0 {
		n += 1 + sovGenesis(uint64(m.BalancerPoolId))
	}
	if m.ClPoolId != 0 {
		n += 1 + sovGenesis(uint64(m.ClPoolId))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolCreationFee = append(m.PoolCreationFee, types.Coin{})
			if err := m.PoolCreationFee[len(m.PoolCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pools", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pools = append(m.Pools, &types1.Any{})
			if err := m.Pools[len(m.Pools)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextPoolNumber", wireType)
			}
			m.NextPoolNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextPoolNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MigrationRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MigrationRecords == nil {
				m.MigrationRecords = &MigrationRecords{}
			}
			if err := m.MigrationRecords.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *MigrationRecords) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: MigrationRecords: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MigrationRecords: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BalancerToConcentratedPoolLinks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BalancerToConcentratedPoolLinks = append(m.BalancerToConcentratedPoolLinks, BalancerToConcentratedPoolLink{})
			if err := m.BalancerToConcentratedPoolLinks[len(m.BalancerToConcentratedPoolLinks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *BalancerToConcentratedPoolLink) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: BalancerToConcentratedPoolLink: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BalancerToConcentratedPoolLink: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BalancerPoolId", wireType)
			}
			m.BalancerPoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BalancerPoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClPoolId", wireType)
			}
			m.ClPoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClPoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
