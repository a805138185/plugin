// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dpos_msg.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	dpos_msg.proto

It has these top-level messages:
	VoteItem
	DPosVote
	DPosVoteReply
	DPosNotify
*/
package types

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

type VoteItem struct {
	VotedNodeIndex   int32  `protobuf:"varint,1,opt,name=votedNodeIndex" json:"votedNodeIndex,omitempty"`
	VotedNodeAddress []byte `protobuf:"bytes,2,opt,name=votedNodeAddress,proto3" json:"votedNodeAddress,omitempty"`
	CycleStart       int64  `protobuf:"varint,3,opt,name=cycleStart" json:"cycleStart,omitempty"`
	CycleStop        int64  `protobuf:"varint,4,opt,name=cycleStop" json:"cycleStop,omitempty"`
	PeriodStart      int64  `protobuf:"varint,5,opt,name=periodStart" json:"periodStart,omitempty"`
	PeriodStop       int64  `protobuf:"varint,6,opt,name=periodStop" json:"periodStop,omitempty"`
	Height           int64  `protobuf:"varint,7,opt,name=height" json:"height,omitempty"`
	VoteID           []byte `protobuf:"bytes,8,opt,name=voteID,proto3" json:"voteID,omitempty"`
}

func (m *VoteItem) Reset()                    { *m = VoteItem{} }
func (m *VoteItem) String() string            { return proto.CompactTextString(m) }
func (*VoteItem) ProtoMessage()               {}
func (*VoteItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *VoteItem) GetVotedNodeIndex() int32 {
	if m != nil {
		return m.VotedNodeIndex
	}
	return 0
}

func (m *VoteItem) GetVotedNodeAddress() []byte {
	if m != nil {
		return m.VotedNodeAddress
	}
	return nil
}

func (m *VoteItem) GetCycleStart() int64 {
	if m != nil {
		return m.CycleStart
	}
	return 0
}

func (m *VoteItem) GetCycleStop() int64 {
	if m != nil {
		return m.CycleStop
	}
	return 0
}

func (m *VoteItem) GetPeriodStart() int64 {
	if m != nil {
		return m.PeriodStart
	}
	return 0
}

func (m *VoteItem) GetPeriodStop() int64 {
	if m != nil {
		return m.PeriodStop
	}
	return 0
}

func (m *VoteItem) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *VoteItem) GetVoteID() []byte {
	if m != nil {
		return m.VoteID
	}
	return nil
}

// DPosVote Dpos共识的节点投票，为达成共识用。
type DPosVote struct {
	VoteItem         *VoteItem `protobuf:"bytes,1,opt,name=voteItem" json:"voteItem,omitempty"`
	VoteTimestamp    int64     `protobuf:"varint,2,opt,name=voteTimestamp" json:"voteTimestamp,omitempty"`
	VoterNodeIndex   int32     `protobuf:"varint,3,opt,name=voterNodeIndex" json:"voterNodeIndex,omitempty"`
	VoterNodeAddress []byte    `protobuf:"bytes,4,opt,name=voterNodeAddress,proto3" json:"voterNodeAddress,omitempty"`
	Signature        []byte    `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *DPosVote) Reset()                    { *m = DPosVote{} }
func (m *DPosVote) String() string            { return proto.CompactTextString(m) }
func (*DPosVote) ProtoMessage()               {}
func (*DPosVote) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DPosVote) GetVoteItem() *VoteItem {
	if m != nil {
		return m.VoteItem
	}
	return nil
}

func (m *DPosVote) GetVoteTimestamp() int64 {
	if m != nil {
		return m.VoteTimestamp
	}
	return 0
}

func (m *DPosVote) GetVoterNodeIndex() int32 {
	if m != nil {
		return m.VoterNodeIndex
	}
	return 0
}

func (m *DPosVote) GetVoterNodeAddress() []byte {
	if m != nil {
		return m.VoterNodeAddress
	}
	return nil
}

func (m *DPosVote) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type DPosVoteReply struct {
	Vote *DPosVote `protobuf:"bytes,1,opt,name=vote" json:"vote,omitempty"`
}

func (m *DPosVoteReply) Reset()                    { *m = DPosVoteReply{} }
func (m *DPosVoteReply) String() string            { return proto.CompactTextString(m) }
func (*DPosVoteReply) ProtoMessage()               {}
func (*DPosVoteReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DPosVoteReply) GetVote() *DPosVote {
	if m != nil {
		return m.Vote
	}
	return nil
}

// DPosNotify Dpos委托节点出块周期结束时，通知其他节点进行高度确认及新节点投票。
type DPosNotify struct {
	Vote              *VoteItem `protobuf:"bytes,1,opt,name=vote" json:"vote,omitempty"`
	HeightStop        int64     `protobuf:"varint,2,opt,name=heightStop" json:"heightStop,omitempty"`
	HashStop          []byte    `protobuf:"bytes,3,opt,name=hashStop,proto3" json:"hashStop,omitempty"`
	NotifyTimestamp   int64     `protobuf:"varint,4,opt,name=notifyTimestamp" json:"notifyTimestamp,omitempty"`
	NotifyNodeIndex   int32     `protobuf:"varint,5,opt,name=notifyNodeIndex" json:"notifyNodeIndex,omitempty"`
	NotifyNodeAddress []byte    `protobuf:"bytes,6,opt,name=notifyNodeAddress,proto3" json:"notifyNodeAddress,omitempty"`
	Signature         []byte    `protobuf:"bytes,7,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *DPosNotify) Reset()                    { *m = DPosNotify{} }
func (m *DPosNotify) String() string            { return proto.CompactTextString(m) }
func (*DPosNotify) ProtoMessage()               {}
func (*DPosNotify) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DPosNotify) GetVote() *VoteItem {
	if m != nil {
		return m.Vote
	}
	return nil
}

func (m *DPosNotify) GetHeightStop() int64 {
	if m != nil {
		return m.HeightStop
	}
	return 0
}

func (m *DPosNotify) GetHashStop() []byte {
	if m != nil {
		return m.HashStop
	}
	return nil
}

func (m *DPosNotify) GetNotifyTimestamp() int64 {
	if m != nil {
		return m.NotifyTimestamp
	}
	return 0
}

func (m *DPosNotify) GetNotifyNodeIndex() int32 {
	if m != nil {
		return m.NotifyNodeIndex
	}
	return 0
}

func (m *DPosNotify) GetNotifyNodeAddress() []byte {
	if m != nil {
		return m.NotifyNodeAddress
	}
	return nil
}

func (m *DPosNotify) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*VoteItem)(nil), "types.VoteItem")
	proto.RegisterType((*DPosVote)(nil), "types.DPosVote")
	proto.RegisterType((*DPosVoteReply)(nil), "types.DPosVoteReply")
	proto.RegisterType((*DPosNotify)(nil), "types.DPosNotify")
}

func init() { proto.RegisterFile("dpos_msg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0xc7, 0x49, 0xd3, 0xa4, 0xb9, 0xa7, 0x5f, 0xf7, 0xce, 0xe2, 0x12, 0x2e, 0x17, 0x09, 0x51,
	0x24, 0xa8, 0x74, 0xa1, 0xbe, 0x80, 0xd0, 0x4d, 0x37, 0x45, 0xa2, 0xb8, 0x95, 0xd8, 0x8c, 0x4d,
	0xa0, 0xe9, 0x0c, 0x33, 0x63, 0xb1, 0x0f, 0xa1, 0xaf, 0xe6, 0x2b, 0xc9, 0x9c, 0x7c, 0x4c, 0x9a,
	0xe2, 0xae, 0xe7, 0xf7, 0x9f, 0x8f, 0x73, 0x7e, 0xd3, 0xc0, 0x24, 0xe5, 0x4c, 0x3e, 0x17, 0x72,
	0x3d, 0xe3, 0x82, 0x29, 0x46, 0x1c, 0xb5, 0xe7, 0x54, 0x86, 0x9f, 0x3d, 0xf0, 0x9e, 0x98, 0xa2,
	0x0b, 0x45, 0x0b, 0x72, 0x0e, 0x93, 0x1d, 0x53, 0x34, 0x5d, 0xb2, 0x94, 0x2e, 0xb6, 0x29, 0x7d,
	0xf7, 0xad, 0xc0, 0x8a, 0x9c, 0xb8, 0x43, 0xc9, 0x05, 0xfc, 0x6e, 0xc8, 0x5d, 0x9a, 0x0a, 0x2a,
	0xa5, 0xdf, 0x0b, 0xac, 0x68, 0x14, 0x1f, 0x71, 0x72, 0x02, 0xb0, 0xda, 0xaf, 0x36, 0xf4, 0x41,
	0x25, 0x42, 0xf9, 0x76, 0x60, 0x45, 0x76, 0xdc, 0x22, 0xe4, 0x3f, 0xfc, 0xaa, 0x2a, 0xc6, 0xfd,
	0x3e, 0xc6, 0x06, 0x90, 0x00, 0x86, 0x9c, 0x8a, 0x9c, 0xa5, 0xe5, 0x76, 0x07, 0xf3, 0x36, 0xd2,
	0xe7, 0xd7, 0x25, 0xe3, 0xbe, 0x5b, 0x9e, 0x6f, 0x08, 0xf9, 0x0b, 0x6e, 0x46, 0xf3, 0x75, 0xa6,
	0xfc, 0x01, 0x66, 0x55, 0xa5, 0xb9, 0xee, 0x75, 0x31, 0xf7, 0x3d, 0xec, 0xbc, 0xaa, 0xc2, 0x2f,
	0x0b, 0xbc, 0xf9, 0x3d, 0x93, 0x5a, 0x0a, 0xb9, 0x04, 0x6f, 0x57, 0xc9, 0x41, 0x15, 0xc3, 0xeb,
	0xe9, 0x0c, 0xbd, 0xcd, 0x6a, 0x67, 0x71, 0xb3, 0x80, 0x9c, 0xc1, 0x58, 0xff, 0x7e, 0xcc, 0x0b,
	0x2a, 0x55, 0x52, 0x70, 0x54, 0x62, 0xc7, 0x87, 0xb0, 0x76, 0x2c, 0x8c, 0x63, 0xdb, 0x38, 0x16,
	0x47, 0x8e, 0x45, 0xdb, 0x71, 0xdf, 0x38, 0x6e, 0x73, 0xed, 0x50, 0xe6, 0xeb, 0x6d, 0xa2, 0xde,
	0x04, 0x45, 0x47, 0xa3, 0xd8, 0x80, 0xf0, 0x16, 0xc6, 0xf5, 0x40, 0x31, 0xe5, 0x9b, 0x3d, 0x39,
	0x85, 0xbe, 0x3e, 0xa2, 0x33, 0x51, 0xb3, 0x06, 0xc3, 0xf0, 0xa3, 0x07, 0xa0, 0xd1, 0x92, 0xa9,
	0xfc, 0xf5, 0xa7, 0x3d, 0x8d, 0x05, 0x0c, 0xf5, 0x5b, 0x94, 0x76, 0xf1, 0x2d, 0xca, 0xf1, 0x5b,
	0x84, 0xfc, 0x03, 0x2f, 0x4b, 0x64, 0x86, 0xa9, 0x8d, 0x6d, 0x36, 0x35, 0x89, 0x60, 0xba, 0xc5,
	0xab, 0x8c, 0xbf, 0xf2, 0xdf, 0xd0, 0xc5, 0x66, 0xa5, 0x51, 0xe8, 0xa0, 0xc2, 0x2e, 0x26, 0x57,
	0xf0, 0xc7, 0xa0, 0x5a, 0xa2, 0x8b, 0x17, 0x1f, 0x07, 0x87, 0x16, 0x07, 0x1d, 0x8b, 0x2f, 0x2e,
	0x7e, 0x36, 0x37, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xf4, 0x71, 0xcb, 0x48, 0x03, 0x00,
	0x00,
}
