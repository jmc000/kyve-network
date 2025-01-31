// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kyve/delegation/v1beta1/events.proto

package types

import (
	fmt "fmt"
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

// EventDelegate is an event emitted when someone delegates to a protocol node.
// emitted_by: MsgDelegate
type EventDelegate struct {
	// address is the account address of the delegator.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// staker is the account address of the protocol node.
	Staker string `protobuf:"bytes,2,opt,name=staker,proto3" json:"staker,omitempty"`
	// amount ...
	Amount uint64 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *EventDelegate) Reset()         { *m = EventDelegate{} }
func (m *EventDelegate) String() string { return proto.CompactTextString(m) }
func (*EventDelegate) ProtoMessage()    {}
func (*EventDelegate) Descriptor() ([]byte, []int) {
	return fileDescriptor_d01988a9108a2e89, []int{0}
}
func (m *EventDelegate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventDelegate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventDelegate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventDelegate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventDelegate.Merge(m, src)
}
func (m *EventDelegate) XXX_Size() int {
	return m.Size()
}
func (m *EventDelegate) XXX_DiscardUnknown() {
	xxx_messageInfo_EventDelegate.DiscardUnknown(m)
}

var xxx_messageInfo_EventDelegate proto.InternalMessageInfo

func (m *EventDelegate) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *EventDelegate) GetStaker() string {
	if m != nil {
		return m.Staker
	}
	return ""
}

func (m *EventDelegate) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// EventUndelegate is an event emitted when someone undelegates from a protocol node.
// emitted_by: EndBlock
type EventUndelegate struct {
	// address is the account address of the delegator.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// staker is the account address of the protocol node.
	Staker string `protobuf:"bytes,2,opt,name=staker,proto3" json:"staker,omitempty"`
	// amount ...
	Amount uint64 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *EventUndelegate) Reset()         { *m = EventUndelegate{} }
func (m *EventUndelegate) String() string { return proto.CompactTextString(m) }
func (*EventUndelegate) ProtoMessage()    {}
func (*EventUndelegate) Descriptor() ([]byte, []int) {
	return fileDescriptor_d01988a9108a2e89, []int{1}
}
func (m *EventUndelegate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUndelegate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUndelegate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUndelegate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUndelegate.Merge(m, src)
}
func (m *EventUndelegate) XXX_Size() int {
	return m.Size()
}
func (m *EventUndelegate) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUndelegate.DiscardUnknown(m)
}

var xxx_messageInfo_EventUndelegate proto.InternalMessageInfo

func (m *EventUndelegate) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *EventUndelegate) GetStaker() string {
	if m != nil {
		return m.Staker
	}
	return ""
}

func (m *EventUndelegate) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// EventRedelegate is an event emitted when someone redelegates from one protocol node to another.
// emitted_by: MsgRedelegate
type EventRedelegate struct {
	// address is the account address of the delegator.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// from_staker ...
	FromStaker string `protobuf:"bytes,2,opt,name=from_staker,json=fromStaker,proto3" json:"from_staker,omitempty"`
	// to_staker is the account address of the new staker in the the pool
	ToStaker string `protobuf:"bytes,3,opt,name=to_staker,json=toStaker,proto3" json:"to_staker,omitempty"`
	// amount ...
	Amount uint64 `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *EventRedelegate) Reset()         { *m = EventRedelegate{} }
func (m *EventRedelegate) String() string { return proto.CompactTextString(m) }
func (*EventRedelegate) ProtoMessage()    {}
func (*EventRedelegate) Descriptor() ([]byte, []int) {
	return fileDescriptor_d01988a9108a2e89, []int{2}
}
func (m *EventRedelegate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRedelegate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRedelegate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRedelegate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRedelegate.Merge(m, src)
}
func (m *EventRedelegate) XXX_Size() int {
	return m.Size()
}
func (m *EventRedelegate) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRedelegate.DiscardUnknown(m)
}

var xxx_messageInfo_EventRedelegate proto.InternalMessageInfo

func (m *EventRedelegate) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *EventRedelegate) GetFromStaker() string {
	if m != nil {
		return m.FromStaker
	}
	return ""
}

func (m *EventRedelegate) GetToStaker() string {
	if m != nil {
		return m.ToStaker
	}
	return ""
}

func (m *EventRedelegate) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// EventWithdrawRewards ...
// emitted_by: MsgRedelegate, MsgDelegate, MsgWithdrawRewards, EndBlock
type EventWithdrawRewards struct {
	// address is the account address of the delegator.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// staker is the account address of the protocol node the users withdraws from.
	Staker string `protobuf:"bytes,2,opt,name=staker,proto3" json:"staker,omitempty"`
	// amount ...
	Amount uint64 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *EventWithdrawRewards) Reset()         { *m = EventWithdrawRewards{} }
func (m *EventWithdrawRewards) String() string { return proto.CompactTextString(m) }
func (*EventWithdrawRewards) ProtoMessage()    {}
func (*EventWithdrawRewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_d01988a9108a2e89, []int{3}
}
func (m *EventWithdrawRewards) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventWithdrawRewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventWithdrawRewards.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventWithdrawRewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventWithdrawRewards.Merge(m, src)
}
func (m *EventWithdrawRewards) XXX_Size() int {
	return m.Size()
}
func (m *EventWithdrawRewards) XXX_DiscardUnknown() {
	xxx_messageInfo_EventWithdrawRewards.DiscardUnknown(m)
}

var xxx_messageInfo_EventWithdrawRewards proto.InternalMessageInfo

func (m *EventWithdrawRewards) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *EventWithdrawRewards) GetStaker() string {
	if m != nil {
		return m.Staker
	}
	return ""
}

func (m *EventWithdrawRewards) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// EventSlash is an event emitted when a protocol node is slashed.
// emitted_by: MsgSubmitBundleProposal, EndBlock
type EventSlash struct {
	// pool_id is the unique ID of the pool.
	PoolId uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	// staker is the account address of the protocol node.
	Staker string `protobuf:"bytes,2,opt,name=staker,proto3" json:"staker,omitempty"`
	// amount ...
	Amount uint64 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	// slash_type
	SlashType SlashType `protobuf:"varint,4,opt,name=slash_type,json=slashType,proto3,enum=kyve.delegation.v1beta1.SlashType" json:"slash_type,omitempty"`
}

func (m *EventSlash) Reset()         { *m = EventSlash{} }
func (m *EventSlash) String() string { return proto.CompactTextString(m) }
func (*EventSlash) ProtoMessage()    {}
func (*EventSlash) Descriptor() ([]byte, []int) {
	return fileDescriptor_d01988a9108a2e89, []int{4}
}
func (m *EventSlash) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSlash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventSlash.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventSlash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSlash.Merge(m, src)
}
func (m *EventSlash) XXX_Size() int {
	return m.Size()
}
func (m *EventSlash) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSlash.DiscardUnknown(m)
}

var xxx_messageInfo_EventSlash proto.InternalMessageInfo

func (m *EventSlash) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *EventSlash) GetStaker() string {
	if m != nil {
		return m.Staker
	}
	return ""
}

func (m *EventSlash) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *EventSlash) GetSlashType() SlashType {
	if m != nil {
		return m.SlashType
	}
	return SLASH_TYPE_UNSPECIFIED
}

func init() {
	proto.RegisterType((*EventDelegate)(nil), "kyve.delegation.v1beta1.EventDelegate")
	proto.RegisterType((*EventUndelegate)(nil), "kyve.delegation.v1beta1.EventUndelegate")
	proto.RegisterType((*EventRedelegate)(nil), "kyve.delegation.v1beta1.EventRedelegate")
	proto.RegisterType((*EventWithdrawRewards)(nil), "kyve.delegation.v1beta1.EventWithdrawRewards")
	proto.RegisterType((*EventSlash)(nil), "kyve.delegation.v1beta1.EventSlash")
}

func init() {
	proto.RegisterFile("kyve/delegation/v1beta1/events.proto", fileDescriptor_d01988a9108a2e89)
}

var fileDescriptor_d01988a9108a2e89 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x4e, 0xf2, 0x40,
	0x14, 0xc5, 0x99, 0x0f, 0x02, 0x1f, 0xd7, 0xa8, 0x49, 0x63, 0xa4, 0xd1, 0xa4, 0x12, 0xe2, 0x82,
	0x55, 0x1b, 0xf4, 0x09, 0x34, 0xb2, 0x20, 0x26, 0x2e, 0x8a, 0x7f, 0x82, 0x2e, 0x70, 0x60, 0xae,
	0xb4, 0x01, 0x3a, 0xcd, 0xcc, 0x00, 0xb2, 0xf4, 0x0d, 0x5c, 0xfa, 0x48, 0x2e, 0x59, 0xba, 0x34,
	0xf0, 0x22, 0x66, 0x86, 0x12, 0xc1, 0x84, 0x44, 0x13, 0x76, 0x3d, 0x73, 0x4e, 0x7f, 0xe7, 0xde,
	0xe4, 0xc2, 0x71, 0x77, 0x3c, 0x44, 0x8f, 0x61, 0x0f, 0x3b, 0x54, 0x85, 0x3c, 0xf2, 0x86, 0x95,
	0x16, 0x2a, 0x5a, 0xf1, 0x70, 0x88, 0x91, 0x92, 0x6e, 0x2c, 0xb8, 0xe2, 0x56, 0x41, 0xa7, 0xdc,
	0xef, 0x94, 0x9b, 0xa4, 0x0e, 0xca, 0xeb, 0x7e, 0x5f, 0xca, 0x1a, 0x44, 0xa9, 0x01, 0xdb, 0x55,
	0x8d, 0xbc, 0x98, 0x1b, 0x68, 0xd9, 0x90, 0xa3, 0x8c, 0x09, 0x94, 0xd2, 0x26, 0x45, 0x52, 0xce,
	0xfb, 0x0b, 0x69, 0xed, 0x43, 0x56, 0x2a, 0xda, 0x45, 0x61, 0xff, 0x33, 0x46, 0xa2, 0xf4, 0x3b,
	0xed, 0xf3, 0x41, 0xa4, 0xec, 0x74, 0x91, 0x94, 0x33, 0x7e, 0xa2, 0x4a, 0x0f, 0xb0, 0x6b, 0xd0,
	0x37, 0x11, 0xdb, 0x3c, 0xfc, 0x85, 0x24, 0x74, 0x1f, 0x7f, 0x41, 0x3f, 0x82, 0xad, 0x27, 0xc1,
	0xfb, 0xcd, 0x95, 0x0a, 0xd0, 0x4f, 0xf5, 0x79, 0xcd, 0x21, 0xe4, 0x15, 0x5f, 0xd8, 0x69, 0x63,
	0xff, 0x57, 0xbc, 0xfe, 0x73, 0x86, 0xcc, 0xca, 0x0c, 0x8f, 0xb0, 0x67, 0x46, 0xb8, 0x0b, 0x55,
	0xc0, 0x04, 0x1d, 0xf9, 0x38, 0xa2, 0x82, 0xc9, 0x0d, 0x6e, 0xf9, 0x46, 0x00, 0x4c, 0x45, 0xbd,
	0x47, 0x65, 0x60, 0x15, 0x20, 0x17, 0x73, 0xde, 0x6b, 0x86, 0xcc, 0x80, 0x33, 0x7e, 0x56, 0xcb,
	0x1a, 0xfb, 0x2b, 0xd7, 0x3a, 0x03, 0x90, 0x9a, 0xd8, 0x54, 0xe3, 0x18, 0xcd, 0x56, 0x3b, 0x27,
	0x25, 0x77, 0xcd, 0x35, 0xb9, 0xa6, 0xfc, 0x7a, 0x1c, 0xa3, 0x9f, 0x97, 0x8b, 0xcf, 0xf3, 0xda,
	0xfb, 0xd4, 0x21, 0x93, 0xa9, 0x43, 0x3e, 0xa7, 0x0e, 0x79, 0x9d, 0x39, 0xa9, 0xc9, 0xcc, 0x49,
	0x7d, 0xcc, 0x9c, 0xd4, 0xbd, 0xd7, 0x09, 0x55, 0x30, 0x68, 0xb9, 0x6d, 0xde, 0xf7, 0x2e, 0x1b,
	0xb7, 0xd5, 0x2b, 0x54, 0x23, 0x2e, 0xba, 0x5e, 0x3b, 0xa0, 0x61, 0xe4, 0x3d, 0x2f, 0x9f, 0xa5,
	0xae, 0x97, 0xad, 0xac, 0x39, 0xc5, 0xd3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x35, 0xe5, 0x5e,
	0x91, 0xf5, 0x02, 0x00, 0x00,
}

func (m *EventDelegate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventDelegate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventDelegate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Staker) > 0 {
		i -= len(m.Staker)
		copy(dAtA[i:], m.Staker)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Staker)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventUndelegate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUndelegate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUndelegate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Staker) > 0 {
		i -= len(m.Staker)
		copy(dAtA[i:], m.Staker)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Staker)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventRedelegate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRedelegate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRedelegate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ToStaker) > 0 {
		i -= len(m.ToStaker)
		copy(dAtA[i:], m.ToStaker)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ToStaker)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.FromStaker) > 0 {
		i -= len(m.FromStaker)
		copy(dAtA[i:], m.FromStaker)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.FromStaker)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventWithdrawRewards) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventWithdrawRewards) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventWithdrawRewards) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Staker) > 0 {
		i -= len(m.Staker)
		copy(dAtA[i:], m.Staker)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Staker)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventSlash) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSlash) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSlash) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SlashType != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.SlashType))
		i--
		dAtA[i] = 0x20
	}
	if m.Amount != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Staker) > 0 {
		i -= len(m.Staker)
		copy(dAtA[i:], m.Staker)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Staker)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventDelegate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Staker)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovEvents(uint64(m.Amount))
	}
	return n
}

func (m *EventUndelegate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Staker)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovEvents(uint64(m.Amount))
	}
	return n
}

func (m *EventRedelegate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.FromStaker)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.ToStaker)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovEvents(uint64(m.Amount))
	}
	return n
}

func (m *EventWithdrawRewards) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Staker)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovEvents(uint64(m.Amount))
	}
	return n
}

func (m *EventSlash) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovEvents(uint64(m.PoolId))
	}
	l = len(m.Staker)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovEvents(uint64(m.Amount))
	}
	if m.SlashType != 0 {
		n += 1 + sovEvents(uint64(m.SlashType))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventDelegate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventDelegate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventDelegate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Staker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Staker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventUndelegate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventUndelegate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUndelegate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Staker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Staker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventRedelegate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventRedelegate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRedelegate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromStaker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromStaker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToStaker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToStaker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventWithdrawRewards) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventWithdrawRewards: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventWithdrawRewards: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Staker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Staker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventSlash) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventSlash: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSlash: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Staker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Staker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashType", wireType)
			}
			m.SlashType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SlashType |= SlashType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
