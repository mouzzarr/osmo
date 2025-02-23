// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/gamm/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
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

// GenesisState defines the gamm module's genesis state.
type GenesisState struct {
	Pools []*types.Any `protobuf:"bytes,1,rep,name=pools,proto3" json:"pools,omitempty"`
	// will be renamed to next_pool_id in an upcoming version
	NextPoolNumber uint64 `protobuf:"varint,2,opt,name=next_pool_number,json=nextPoolNumber,proto3" json:"next_pool_number,omitempty"`
	Params         Params `protobuf:"bytes,3,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a324eb7f1dd793e, []int{0}
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

func (m *GenesisState) GetPools() []*types.Any {
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

type Params struct {
	PoolCreationFee      github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=pool_creation_fee,json=poolCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"pool_creation_fee" yaml:"pool_creation_fee"`
	EnableGlobalPoolFees bool                                     `protobuf:"varint,2,opt,name=enable_global_pool_fees,json=enableGlobalPoolFees,proto3" json:"enable_global_pool_fees,omitempty"`
	GlobalFees           GlobalFees                               `protobuf:"bytes,3,opt,name=global_fees,json=globalFees,proto3" json:"global_fees" yaml:"global_fees"`
	TakerFee             github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,4,opt,name=taker_fee,json=takerFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"taker_fee" yaml:"taker_fee"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a324eb7f1dd793e, []int{1}
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

func (m *Params) GetEnableGlobalPoolFees() bool {
	if m != nil {
		return m.EnableGlobalPoolFees
	}
	return false
}

func (m *Params) GetGlobalFees() GlobalFees {
	if m != nil {
		return m.GlobalFees
	}
	return GlobalFees{}
}

type GlobalFees struct {
	SwapFee github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=swap_fee,json=swapFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"swap_fee" yaml:"swap_fee"`
	ExitFee github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=exit_fee,json=exitFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"exit_fee" yaml:"exit_fee"`
}

func (m *GlobalFees) Reset()         { *m = GlobalFees{} }
func (m *GlobalFees) String() string { return proto.CompactTextString(m) }
func (*GlobalFees) ProtoMessage()    {}
func (*GlobalFees) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a324eb7f1dd793e, []int{2}
}
func (m *GlobalFees) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GlobalFees) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GlobalFees.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GlobalFees) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalFees.Merge(m, src)
}
func (m *GlobalFees) XXX_Size() int {
	return m.Size()
}
func (m *GlobalFees) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalFees.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalFees proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisState)(nil), "osmosis.gamm.v1beta1.GenesisState")
	proto.RegisterType((*Params)(nil), "osmosis.gamm.v1beta1.Params")
	proto.RegisterType((*GlobalFees)(nil), "osmosis.gamm.v1beta1.GlobalFees")
}

func init() {
	proto.RegisterFile("osmosis/gamm/v1beta1/genesis.proto", fileDescriptor_5a324eb7f1dd793e)
}

var fileDescriptor_5a324eb7f1dd793e = []byte{
	// 550 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x8e, 0x9b, 0x34, 0x4d, 0x36, 0x88, 0x16, 0x2b, 0x12, 0x69, 0x84, 0x9c, 0xc8, 0x07, 0xe4,
	0x4b, 0x76, 0x49, 0x51, 0x2f, 0xbd, 0xd5, 0x45, 0xa9, 0x40, 0x08, 0x55, 0xe6, 0x86, 0x8a, 0xac,
	0x75, 0x98, 0x18, 0xab, 0xb6, 0x37, 0xca, 0x6e, 0x4a, 0xf2, 0x16, 0x48, 0x88, 0x97, 0x80, 0x2b,
	0x0f, 0x11, 0x71, 0xea, 0x11, 0x71, 0x08, 0x28, 0x79, 0x83, 0x3e, 0x01, 0xda, 0x1f, 0x87, 0x4a,
	0xf4, 0x00, 0xe2, 0xe4, 0x19, 0xcf, 0xf7, 0x33, 0xf3, 0x59, 0x46, 0x2e, 0xe3, 0x19, 0xe3, 0x09,
	0x27, 0x31, 0xcd, 0x32, 0x72, 0xd9, 0x8f, 0x40, 0xd0, 0x3e, 0x89, 0x21, 0x07, 0x9e, 0x70, 0x3c,
	0x9e, 0x30, 0xc1, 0xec, 0xa6, 0xc1, 0x60, 0x89, 0xc1, 0x06, 0xd3, 0x6e, 0xc6, 0x2c, 0x66, 0x0a,
	0x40, 0x64, 0xa5, 0xb1, 0xed, 0xfd, 0x98, 0xb1, 0x38, 0x05, 0xa2, 0xba, 0x68, 0x3a, 0x22, 0x34,
	0x9f, 0x17, 0xa3, 0xa1, 0xd2, 0x09, 0x35, 0x47, 0x37, 0x66, 0xe4, 0xe8, 0x8e, 0x44, 0x94, 0xc3,
	0x66, 0x89, 0x21, 0x4b, 0x72, 0x3d, 0x77, 0x3f, 0x5b, 0xe8, 0xce, 0xa9, 0xde, 0xe9, 0xa5, 0xa0,
	0x02, 0xec, 0x43, 0xb4, 0x3d, 0x66, 0x2c, 0xe5, 0x2d, 0xab, 0x5b, 0xf6, 0x1a, 0x07, 0x4d, 0xac,
	0x6d, 0x71, 0x61, 0x8b, 0x8f, 0xf3, 0xb9, 0x5f, 0xff, 0xfa, 0xa5, 0xb7, 0x7d, 0xc6, 0x58, 0xfa,
	0x34, 0xd0, 0x68, 0xdb, 0x43, 0x7b, 0x39, 0xcc, 0x44, 0x28, 0xbb, 0x30, 0x9f, 0x66, 0x11, 0x4c,
	0x5a, 0x5b, 0x5d, 0xcb, 0xab, 0x04, 0x77, 0xe5, 0x7b, 0x89, 0x7d, 0xa1, 0xde, 0xda, 0x47, 0xa8,
	0x3a, 0xa6, 0x13, 0x9a, 0xf1, 0x56, 0xb9, 0x6b, 0x79, 0x8d, 0x83, 0x07, 0xf8, 0xb6, 0x10, 0xf0,
	0x99, 0xc2, 0xf8, 0x95, 0xc5, 0xb2, 0x53, 0x0a, 0x0c, 0xc3, 0xfd, 0x58, 0x46, 0x55, 0x3d, 0xb0,
	0x3f, 0x58, 0xe8, 0x9e, 0x32, 0x1b, 0x4e, 0x80, 0x8a, 0x84, 0xe5, 0xe1, 0x08, 0xc0, 0x2c, 0xbd,
	0x8f, 0x4d, 0x06, 0xf2, 0xea, 0x8d, 0xe2, 0x09, 0x4b, 0x72, 0xff, 0xb9, 0xd4, 0xbb, 0x5e, 0x76,
	0x5a, 0x73, 0x9a, 0xa5, 0x47, 0xee, 0x1f, 0x0a, 0xee, 0xa7, 0x1f, 0x1d, 0x2f, 0x4e, 0xc4, 0xdb,
	0x69, 0x84, 0x87, 0x2c, 0x33, 0x61, 0x9a, 0x47, 0x8f, 0xbf, 0xb9, 0x20, 0x62, 0x3e, 0x06, 0xae,
	0xc4, 0x78, 0xb0, 0x2b, 0xf9, 0x27, 0x86, 0x3e, 0x00, 0x99, 0xde, 0x7d, 0xc8, 0x69, 0x94, 0x42,
	0x18, 0xa7, 0x2c, 0xa2, 0xa9, 0xce, 0x63, 0x04, 0xc0, 0x55, 0x1a, 0xb5, 0xa0, 0xa9, 0xc7, 0xa7,
	0x6a, 0x2a, 0x53, 0x19, 0x00, 0x70, 0xfb, 0x35, 0x6a, 0x18, 0xbc, 0x82, 0xea, 0x60, 0xba, 0xb7,
	0x07, 0xa3, 0xa9, 0x92, 0xe6, 0xb7, 0xcd, 0x31, 0xb6, 0x3e, 0xe6, 0x86, 0x84, 0x1b, 0xa0, 0x78,
	0x83, 0xb3, 0x43, 0x54, 0x17, 0xf4, 0x02, 0x26, 0x2a, 0xa2, 0x4a, 0xd7, 0xf2, 0xea, 0xbe, 0x2f,
	0xa9, 0xdf, 0x97, 0x9d, 0x87, 0x7f, 0x71, 0xeb, 0x13, 0x18, 0x5e, 0x2f, 0x3b, 0x7b, 0xda, 0x64,
	0x23, 0xe4, 0x06, 0x35, 0x55, 0x0f, 0x00, 0xdc, 0x85, 0x85, 0xd0, 0xef, 0xbd, 0xec, 0x73, 0x54,
	0xe3, 0xef, 0xe8, 0xd8, 0x7c, 0x11, 0x69, 0x77, 0xfc, 0xcf, 0x76, 0xbb, 0xda, 0xae, 0xd0, 0x71,
	0x83, 0x1d, 0x59, 0xca, 0x8c, 0xcf, 0x51, 0x0d, 0x66, 0x89, 0x50, 0xea, 0x5b, 0xff, 0xa7, 0x5e,
	0xe8, 0xb8, 0xc1, 0x8e, 0x2c, 0x07, 0x00, 0xfe, 0xb3, 0xc5, 0xca, 0xb1, 0xae, 0x56, 0x8e, 0xf5,
	0x73, 0xe5, 0x58, 0xef, 0xd7, 0x4e, 0xe9, 0x6a, 0xed, 0x94, 0xbe, 0xad, 0x9d, 0xd2, 0xab, 0x47,
	0x37, 0xd4, 0xcd, 0x97, 0xe9, 0xa5, 0x34, 0xe2, 0x45, 0x43, 0x2e, 0xfb, 0x87, 0x64, 0xa6, 0x7f,
	0x77, 0xe5, 0x15, 0x55, 0xd5, 0x4f, 0xf3, 0xf8, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x46, 0x35,
	0xad, 0x1b, 0x0b, 0x04, 0x00, 0x00,
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
	{
		size := m.TakerFee.Size()
		i -= size
		if _, err := m.TakerFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.GlobalFees.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.EnableGlobalPoolFees {
		i--
		if m.EnableGlobalPoolFees {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
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

func (m *GlobalFees) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GlobalFees) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GlobalFees) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ExitFee.Size()
		i -= size
		if _, err := m.ExitFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.SwapFee.Size()
		i -= size
		if _, err := m.SwapFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	return n
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
	if m.EnableGlobalPoolFees {
		n += 2
	}
	l = m.GlobalFees.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.TakerFee.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *GlobalFees) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.SwapFee.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.ExitFee.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
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
			m.Pools = append(m.Pools, &types.Any{})
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
			m.PoolCreationFee = append(m.PoolCreationFee, types1.Coin{})
			if err := m.PoolCreationFee[len(m.PoolCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableGlobalPoolFees", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableGlobalPoolFees = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalFees", wireType)
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
			if err := m.GlobalFees.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakerFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TakerFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *GlobalFees) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GlobalFees: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GlobalFees: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExitFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExitFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
