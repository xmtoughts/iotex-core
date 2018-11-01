// Code generated by protoc-gen-go. DO NOT EDIT.
// source: action.proto

package iproto

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

type TransferPb struct {
	// used by state-based model
	Amount               []byte   `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Sender               string   `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient            string   `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Payload              []byte   `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	SenderPubKey         []byte   `protobuf:"bytes,5,opt,name=senderPubKey,proto3" json:"senderPubKey,omitempty"`
	IsCoinbase           bool     `protobuf:"varint,6,opt,name=isCoinbase,proto3" json:"isCoinbase,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferPb) Reset()         { *m = TransferPb{} }
func (m *TransferPb) String() string { return proto.CompactTextString(m) }
func (*TransferPb) ProtoMessage()    {}
func (*TransferPb) Descriptor() ([]byte, []int) {
	return fileDescriptor_action_6883fb52cdc7bfb0, []int{0}
}
func (m *TransferPb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferPb.Unmarshal(m, b)
}
func (m *TransferPb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferPb.Marshal(b, m, deterministic)
}
func (dst *TransferPb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferPb.Merge(dst, src)
}
func (m *TransferPb) XXX_Size() int {
	return xxx_messageInfo_TransferPb.Size(m)
}
func (m *TransferPb) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferPb.DiscardUnknown(m)
}

var xxx_messageInfo_TransferPb proto.InternalMessageInfo

func (m *TransferPb) GetAmount() []byte {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *TransferPb) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *TransferPb) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *TransferPb) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *TransferPb) GetSenderPubKey() []byte {
	if m != nil {
		return m.SenderPubKey
	}
	return nil
}

func (m *TransferPb) GetIsCoinbase() bool {
	if m != nil {
		return m.IsCoinbase
	}
	return false
}

type VotePb struct {
	Timestamp            uint64   `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	SelfPubkey           []byte   `protobuf:"bytes,2,opt,name=selfPubkey,proto3" json:"selfPubkey,omitempty"`
	VoterAddress         string   `protobuf:"bytes,3,opt,name=voterAddress,proto3" json:"voterAddress,omitempty"`
	VoteeAddress         string   `protobuf:"bytes,4,opt,name=voteeAddress,proto3" json:"voteeAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VotePb) Reset()         { *m = VotePb{} }
func (m *VotePb) String() string { return proto.CompactTextString(m) }
func (*VotePb) ProtoMessage()    {}
func (*VotePb) Descriptor() ([]byte, []int) {
	return fileDescriptor_action_6883fb52cdc7bfb0, []int{1}
}
func (m *VotePb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VotePb.Unmarshal(m, b)
}
func (m *VotePb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VotePb.Marshal(b, m, deterministic)
}
func (dst *VotePb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VotePb.Merge(dst, src)
}
func (m *VotePb) XXX_Size() int {
	return xxx_messageInfo_VotePb.Size(m)
}
func (m *VotePb) XXX_DiscardUnknown() {
	xxx_messageInfo_VotePb.DiscardUnknown(m)
}

var xxx_messageInfo_VotePb proto.InternalMessageInfo

func (m *VotePb) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *VotePb) GetSelfPubkey() []byte {
	if m != nil {
		return m.SelfPubkey
	}
	return nil
}

func (m *VotePb) GetVoterAddress() string {
	if m != nil {
		return m.VoterAddress
	}
	return ""
}

func (m *VotePb) GetVoteeAddress() string {
	if m != nil {
		return m.VoteeAddress
	}
	return ""
}

type ExecutionPb struct {
	Amount               []byte   `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Executor             string   `protobuf:"bytes,2,opt,name=executor,proto3" json:"executor,omitempty"`
	Contract             string   `protobuf:"bytes,3,opt,name=contract,proto3" json:"contract,omitempty"`
	ExecutorPubKey       []byte   `protobuf:"bytes,4,opt,name=executorPubKey,proto3" json:"executorPubKey,omitempty"`
	Data                 []byte   `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecutionPb) Reset()         { *m = ExecutionPb{} }
func (m *ExecutionPb) String() string { return proto.CompactTextString(m) }
func (*ExecutionPb) ProtoMessage()    {}
func (*ExecutionPb) Descriptor() ([]byte, []int) {
	return fileDescriptor_action_6883fb52cdc7bfb0, []int{2}
}
func (m *ExecutionPb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionPb.Unmarshal(m, b)
}
func (m *ExecutionPb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionPb.Marshal(b, m, deterministic)
}
func (dst *ExecutionPb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionPb.Merge(dst, src)
}
func (m *ExecutionPb) XXX_Size() int {
	return xxx_messageInfo_ExecutionPb.Size(m)
}
func (m *ExecutionPb) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionPb.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionPb proto.InternalMessageInfo

func (m *ExecutionPb) GetAmount() []byte {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *ExecutionPb) GetExecutor() string {
	if m != nil {
		return m.Executor
	}
	return ""
}

func (m *ExecutionPb) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *ExecutionPb) GetExecutorPubKey() []byte {
	if m != nil {
		return m.ExecutorPubKey
	}
	return nil
}

func (m *ExecutionPb) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type SecretProposalPb struct {
	Sender               string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient            string   `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Secret               []uint32 `protobuf:"varint,3,rep,packed,name=secret,proto3" json:"secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecretProposalPb) Reset()         { *m = SecretProposalPb{} }
func (m *SecretProposalPb) String() string { return proto.CompactTextString(m) }
func (*SecretProposalPb) ProtoMessage()    {}
func (*SecretProposalPb) Descriptor() ([]byte, []int) {
	return fileDescriptor_action_6883fb52cdc7bfb0, []int{3}
}
func (m *SecretProposalPb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecretProposalPb.Unmarshal(m, b)
}
func (m *SecretProposalPb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecretProposalPb.Marshal(b, m, deterministic)
}
func (dst *SecretProposalPb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecretProposalPb.Merge(dst, src)
}
func (m *SecretProposalPb) XXX_Size() int {
	return xxx_messageInfo_SecretProposalPb.Size(m)
}
func (m *SecretProposalPb) XXX_DiscardUnknown() {
	xxx_messageInfo_SecretProposalPb.DiscardUnknown(m)
}

var xxx_messageInfo_SecretProposalPb proto.InternalMessageInfo

func (m *SecretProposalPb) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *SecretProposalPb) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *SecretProposalPb) GetSecret() []uint32 {
	if m != nil {
		return m.Secret
	}
	return nil
}

type SecretWitnessPb struct {
	Sender               string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Witness              [][]byte `protobuf:"bytes,2,rep,name=witness,proto3" json:"witness,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecretWitnessPb) Reset()         { *m = SecretWitnessPb{} }
func (m *SecretWitnessPb) String() string { return proto.CompactTextString(m) }
func (*SecretWitnessPb) ProtoMessage()    {}
func (*SecretWitnessPb) Descriptor() ([]byte, []int) {
	return fileDescriptor_action_6883fb52cdc7bfb0, []int{4}
}
func (m *SecretWitnessPb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecretWitnessPb.Unmarshal(m, b)
}
func (m *SecretWitnessPb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecretWitnessPb.Marshal(b, m, deterministic)
}
func (dst *SecretWitnessPb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecretWitnessPb.Merge(dst, src)
}
func (m *SecretWitnessPb) XXX_Size() int {
	return xxx_messageInfo_SecretWitnessPb.Size(m)
}
func (m *SecretWitnessPb) XXX_DiscardUnknown() {
	xxx_messageInfo_SecretWitnessPb.DiscardUnknown(m)
}

var xxx_messageInfo_SecretWitnessPb proto.InternalMessageInfo

func (m *SecretWitnessPb) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *SecretWitnessPb) GetWitness() [][]byte {
	if m != nil {
		return m.Witness
	}
	return nil
}

type StartSubChainPb struct {
	// TODO: chainID chould be assigned by system and returned via a receipt
	ChainID              uint32   `protobuf:"varint,1,opt,name=chainID,proto3" json:"chainID,omitempty"`
	SecurityDeposit      []byte   `protobuf:"bytes,2,opt,name=securityDeposit,proto3" json:"securityDeposit,omitempty"`
	OperationDeposit     []byte   `protobuf:"bytes,3,opt,name=operationDeposit,proto3" json:"operationDeposit,omitempty"`
	StartHeight          uint64   `protobuf:"varint,4,opt,name=startHeight,proto3" json:"startHeight,omitempty"`
	ParentHeightOffset   uint64   `protobuf:"varint,5,opt,name=parentHeightOffset,proto3" json:"parentHeightOffset,omitempty"`
	OwnerAddress         string   `protobuf:"bytes,6,opt,name=ownerAddress,proto3" json:"ownerAddress,omitempty"`
	OwnerPublicKey       []byte   `protobuf:"bytes,7,opt,name=ownerPublicKey,proto3" json:"ownerPublicKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartSubChainPb) Reset()         { *m = StartSubChainPb{} }
func (m *StartSubChainPb) String() string { return proto.CompactTextString(m) }
func (*StartSubChainPb) ProtoMessage()    {}
func (*StartSubChainPb) Descriptor() ([]byte, []int) {
	return fileDescriptor_action_6883fb52cdc7bfb0, []int{5}
}
func (m *StartSubChainPb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartSubChainPb.Unmarshal(m, b)
}
func (m *StartSubChainPb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartSubChainPb.Marshal(b, m, deterministic)
}
func (dst *StartSubChainPb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartSubChainPb.Merge(dst, src)
}
func (m *StartSubChainPb) XXX_Size() int {
	return xxx_messageInfo_StartSubChainPb.Size(m)
}
func (m *StartSubChainPb) XXX_DiscardUnknown() {
	xxx_messageInfo_StartSubChainPb.DiscardUnknown(m)
}

var xxx_messageInfo_StartSubChainPb proto.InternalMessageInfo

func (m *StartSubChainPb) GetChainID() uint32 {
	if m != nil {
		return m.ChainID
	}
	return 0
}

func (m *StartSubChainPb) GetSecurityDeposit() []byte {
	if m != nil {
		return m.SecurityDeposit
	}
	return nil
}

func (m *StartSubChainPb) GetOperationDeposit() []byte {
	if m != nil {
		return m.OperationDeposit
	}
	return nil
}

func (m *StartSubChainPb) GetStartHeight() uint64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *StartSubChainPb) GetParentHeightOffset() uint64 {
	if m != nil {
		return m.ParentHeightOffset
	}
	return 0
}

func (m *StartSubChainPb) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
	}
	return ""
}

func (m *StartSubChainPb) GetOwnerPublicKey() []byte {
	if m != nil {
		return m.OwnerPublicKey
	}
	return nil
}

type StopSubChainPb struct {
	ChainID              uint32   `protobuf:"varint,1,opt,name=chainID,proto3" json:"chainID,omitempty"`
	StopHeight           uint64   `protobuf:"varint,2,opt,name=stopHeight,proto3" json:"stopHeight,omitempty"`
	Owner                string   `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	OwnerPublicKey       []byte   `protobuf:"bytes,4,opt,name=ownerPublicKey,proto3" json:"ownerPublicKey,omitempty"`
	SubChainAddress      string   `protobuf:"bytes,5,opt,name=subChainAddress,proto3" json:"subChainAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopSubChainPb) Reset()         { *m = StopSubChainPb{} }
func (m *StopSubChainPb) String() string { return proto.CompactTextString(m) }
func (*StopSubChainPb) ProtoMessage()    {}
func (*StopSubChainPb) Descriptor() ([]byte, []int) {
	return fileDescriptor_action_6883fb52cdc7bfb0, []int{6}
}
func (m *StopSubChainPb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopSubChainPb.Unmarshal(m, b)
}
func (m *StopSubChainPb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopSubChainPb.Marshal(b, m, deterministic)
}
func (dst *StopSubChainPb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopSubChainPb.Merge(dst, src)
}
func (m *StopSubChainPb) XXX_Size() int {
	return xxx_messageInfo_StopSubChainPb.Size(m)
}
func (m *StopSubChainPb) XXX_DiscardUnknown() {
	xxx_messageInfo_StopSubChainPb.DiscardUnknown(m)
}

var xxx_messageInfo_StopSubChainPb proto.InternalMessageInfo

func (m *StopSubChainPb) GetChainID() uint32 {
	if m != nil {
		return m.ChainID
	}
	return 0
}

func (m *StopSubChainPb) GetStopHeight() uint64 {
	if m != nil {
		return m.StopHeight
	}
	return 0
}

func (m *StopSubChainPb) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *StopSubChainPb) GetOwnerPublicKey() []byte {
	if m != nil {
		return m.OwnerPublicKey
	}
	return nil
}

func (m *StopSubChainPb) GetSubChainAddress() string {
	if m != nil {
		return m.SubChainAddress
	}
	return ""
}

type PutBlockPb struct {
	SubChainAddress      string            `protobuf:"bytes,1,opt,name=subChainAddress,proto3" json:"subChainAddress,omitempty"`
	Height               uint64            `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Roots                map[string][]byte `protobuf:"bytes,3,rep,name=roots,proto3" json:"roots,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ProducerAddress      string            `protobuf:"bytes,4,opt,name=producerAddress,proto3" json:"producerAddress,omitempty"`
	ProducerPublicKey    []byte            `protobuf:"bytes,5,opt,name=producerPublicKey,proto3" json:"producerPublicKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PutBlockPb) Reset()         { *m = PutBlockPb{} }
func (m *PutBlockPb) String() string { return proto.CompactTextString(m) }
func (*PutBlockPb) ProtoMessage()    {}
func (*PutBlockPb) Descriptor() ([]byte, []int) {
	return fileDescriptor_action_6883fb52cdc7bfb0, []int{7}
}
func (m *PutBlockPb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutBlockPb.Unmarshal(m, b)
}
func (m *PutBlockPb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutBlockPb.Marshal(b, m, deterministic)
}
func (dst *PutBlockPb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutBlockPb.Merge(dst, src)
}
func (m *PutBlockPb) XXX_Size() int {
	return xxx_messageInfo_PutBlockPb.Size(m)
}
func (m *PutBlockPb) XXX_DiscardUnknown() {
	xxx_messageInfo_PutBlockPb.DiscardUnknown(m)
}

var xxx_messageInfo_PutBlockPb proto.InternalMessageInfo

func (m *PutBlockPb) GetSubChainAddress() string {
	if m != nil {
		return m.SubChainAddress
	}
	return ""
}

func (m *PutBlockPb) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *PutBlockPb) GetRoots() map[string][]byte {
	if m != nil {
		return m.Roots
	}
	return nil
}

func (m *PutBlockPb) GetProducerAddress() string {
	if m != nil {
		return m.ProducerAddress
	}
	return ""
}

func (m *PutBlockPb) GetProducerPublicKey() []byte {
	if m != nil {
		return m.ProducerPublicKey
	}
	return nil
}

type ActionPb struct {
	Version   uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Nonce     uint64 `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	GasLimit  uint64 `protobuf:"varint,3,opt,name=gasLimit,proto3" json:"gasLimit,omitempty"`
	GasPrice  []byte `protobuf:"bytes,4,opt,name=gasPrice,proto3" json:"gasPrice,omitempty"`
	Signature []byte `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	// Types that are valid to be assigned to Action:
	//	*ActionPb_Transfer
	//	*ActionPb_Vote
	//	*ActionPb_Execution
	//	*ActionPb_SecretProposal
	//	*ActionPb_SecretWitness
	//	*ActionPb_StartSubChain
	//	*ActionPb_StopSubChain
	//	*ActionPb_PutBlock
	Action               isActionPb_Action `protobuf_oneof:"action"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ActionPb) Reset()         { *m = ActionPb{} }
func (m *ActionPb) String() string { return proto.CompactTextString(m) }
func (*ActionPb) ProtoMessage()    {}
func (*ActionPb) Descriptor() ([]byte, []int) {
	return fileDescriptor_action_6883fb52cdc7bfb0, []int{8}
}
func (m *ActionPb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionPb.Unmarshal(m, b)
}
func (m *ActionPb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionPb.Marshal(b, m, deterministic)
}
func (dst *ActionPb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionPb.Merge(dst, src)
}
func (m *ActionPb) XXX_Size() int {
	return xxx_messageInfo_ActionPb.Size(m)
}
func (m *ActionPb) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionPb.DiscardUnknown(m)
}

var xxx_messageInfo_ActionPb proto.InternalMessageInfo

type isActionPb_Action interface {
	isActionPb_Action()
}

type ActionPb_Transfer struct {
	Transfer *TransferPb `protobuf:"bytes,10,opt,name=transfer,proto3,oneof"`
}
type ActionPb_Vote struct {
	Vote *VotePb `protobuf:"bytes,11,opt,name=vote,proto3,oneof"`
}
type ActionPb_Execution struct {
	Execution *ExecutionPb `protobuf:"bytes,12,opt,name=execution,proto3,oneof"`
}
type ActionPb_SecretProposal struct {
	SecretProposal *SecretProposalPb `protobuf:"bytes,13,opt,name=secretProposal,proto3,oneof"`
}
type ActionPb_SecretWitness struct {
	SecretWitness *SecretWitnessPb `protobuf:"bytes,14,opt,name=secretWitness,proto3,oneof"`
}
type ActionPb_StartSubChain struct {
	StartSubChain *StartSubChainPb `protobuf:"bytes,15,opt,name=startSubChain,proto3,oneof"`
}
type ActionPb_StopSubChain struct {
	StopSubChain *StopSubChainPb `protobuf:"bytes,16,opt,name=stopSubChain,proto3,oneof"`
}
type ActionPb_PutBlock struct {
	PutBlock *PutBlockPb `protobuf:"bytes,17,opt,name=putBlock,proto3,oneof"`
}

func (*ActionPb_Transfer) isActionPb_Action()       {}
func (*ActionPb_Vote) isActionPb_Action()           {}
func (*ActionPb_Execution) isActionPb_Action()      {}
func (*ActionPb_SecretProposal) isActionPb_Action() {}
func (*ActionPb_SecretWitness) isActionPb_Action()  {}
func (*ActionPb_StartSubChain) isActionPb_Action()  {}
func (*ActionPb_StopSubChain) isActionPb_Action()   {}
func (*ActionPb_PutBlock) isActionPb_Action()       {}

func (m *ActionPb) GetAction() isActionPb_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *ActionPb) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ActionPb) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *ActionPb) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *ActionPb) GetGasPrice() []byte {
	if m != nil {
		return m.GasPrice
	}
	return nil
}

func (m *ActionPb) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *ActionPb) GetTransfer() *TransferPb {
	if x, ok := m.GetAction().(*ActionPb_Transfer); ok {
		return x.Transfer
	}
	return nil
}

func (m *ActionPb) GetVote() *VotePb {
	if x, ok := m.GetAction().(*ActionPb_Vote); ok {
		return x.Vote
	}
	return nil
}

func (m *ActionPb) GetExecution() *ExecutionPb {
	if x, ok := m.GetAction().(*ActionPb_Execution); ok {
		return x.Execution
	}
	return nil
}

func (m *ActionPb) GetSecretProposal() *SecretProposalPb {
	if x, ok := m.GetAction().(*ActionPb_SecretProposal); ok {
		return x.SecretProposal
	}
	return nil
}

func (m *ActionPb) GetSecretWitness() *SecretWitnessPb {
	if x, ok := m.GetAction().(*ActionPb_SecretWitness); ok {
		return x.SecretWitness
	}
	return nil
}

func (m *ActionPb) GetStartSubChain() *StartSubChainPb {
	if x, ok := m.GetAction().(*ActionPb_StartSubChain); ok {
		return x.StartSubChain
	}
	return nil
}

func (m *ActionPb) GetStopSubChain() *StopSubChainPb {
	if x, ok := m.GetAction().(*ActionPb_StopSubChain); ok {
		return x.StopSubChain
	}
	return nil
}

func (m *ActionPb) GetPutBlock() *PutBlockPb {
	if x, ok := m.GetAction().(*ActionPb_PutBlock); ok {
		return x.PutBlock
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ActionPb) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ActionPb_OneofMarshaler, _ActionPb_OneofUnmarshaler, _ActionPb_OneofSizer, []interface{}{
		(*ActionPb_Transfer)(nil),
		(*ActionPb_Vote)(nil),
		(*ActionPb_Execution)(nil),
		(*ActionPb_SecretProposal)(nil),
		(*ActionPb_SecretWitness)(nil),
		(*ActionPb_StartSubChain)(nil),
		(*ActionPb_StopSubChain)(nil),
		(*ActionPb_PutBlock)(nil),
	}
}

func _ActionPb_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ActionPb)
	// action
	switch x := m.Action.(type) {
	case *ActionPb_Transfer:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Transfer); err != nil {
			return err
		}
	case *ActionPb_Vote:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Vote); err != nil {
			return err
		}
	case *ActionPb_Execution:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Execution); err != nil {
			return err
		}
	case *ActionPb_SecretProposal:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SecretProposal); err != nil {
			return err
		}
	case *ActionPb_SecretWitness:
		b.EncodeVarint(14<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SecretWitness); err != nil {
			return err
		}
	case *ActionPb_StartSubChain:
		b.EncodeVarint(15<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StartSubChain); err != nil {
			return err
		}
	case *ActionPb_StopSubChain:
		b.EncodeVarint(16<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StopSubChain); err != nil {
			return err
		}
	case *ActionPb_PutBlock:
		b.EncodeVarint(17<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PutBlock); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ActionPb.Action has unexpected type %T", x)
	}
	return nil
}

func _ActionPb_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ActionPb)
	switch tag {
	case 10: // action.transfer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TransferPb)
		err := b.DecodeMessage(msg)
		m.Action = &ActionPb_Transfer{msg}
		return true, err
	case 11: // action.vote
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(VotePb)
		err := b.DecodeMessage(msg)
		m.Action = &ActionPb_Vote{msg}
		return true, err
	case 12: // action.execution
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ExecutionPb)
		err := b.DecodeMessage(msg)
		m.Action = &ActionPb_Execution{msg}
		return true, err
	case 13: // action.secretProposal
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SecretProposalPb)
		err := b.DecodeMessage(msg)
		m.Action = &ActionPb_SecretProposal{msg}
		return true, err
	case 14: // action.secretWitness
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SecretWitnessPb)
		err := b.DecodeMessage(msg)
		m.Action = &ActionPb_SecretWitness{msg}
		return true, err
	case 15: // action.startSubChain
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StartSubChainPb)
		err := b.DecodeMessage(msg)
		m.Action = &ActionPb_StartSubChain{msg}
		return true, err
	case 16: // action.stopSubChain
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StopSubChainPb)
		err := b.DecodeMessage(msg)
		m.Action = &ActionPb_StopSubChain{msg}
		return true, err
	case 17: // action.putBlock
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PutBlockPb)
		err := b.DecodeMessage(msg)
		m.Action = &ActionPb_PutBlock{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ActionPb_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ActionPb)
	// action
	switch x := m.Action.(type) {
	case *ActionPb_Transfer:
		s := proto.Size(x.Transfer)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ActionPb_Vote:
		s := proto.Size(x.Vote)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ActionPb_Execution:
		s := proto.Size(x.Execution)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ActionPb_SecretProposal:
		s := proto.Size(x.SecretProposal)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ActionPb_SecretWitness:
		s := proto.Size(x.SecretWitness)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ActionPb_StartSubChain:
		s := proto.Size(x.StartSubChain)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ActionPb_StopSubChain:
		s := proto.Size(x.StopSubChain)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ActionPb_PutBlock:
		s := proto.Size(x.PutBlock)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*TransferPb)(nil), "iproto.TransferPb")
	proto.RegisterType((*VotePb)(nil), "iproto.VotePb")
	proto.RegisterType((*ExecutionPb)(nil), "iproto.ExecutionPb")
	proto.RegisterType((*SecretProposalPb)(nil), "iproto.SecretProposalPb")
	proto.RegisterType((*SecretWitnessPb)(nil), "iproto.SecretWitnessPb")
	proto.RegisterType((*StartSubChainPb)(nil), "iproto.StartSubChainPb")
	proto.RegisterType((*StopSubChainPb)(nil), "iproto.StopSubChainPb")
	proto.RegisterType((*PutBlockPb)(nil), "iproto.PutBlockPb")
	proto.RegisterMapType((map[string][]byte)(nil), "iproto.PutBlockPb.RootsEntry")
	proto.RegisterType((*ActionPb)(nil), "iproto.ActionPb")
}

func init() { proto.RegisterFile("action.proto", fileDescriptor_action_6883fb52cdc7bfb0) }

var fileDescriptor_action_6883fb52cdc7bfb0 = []byte{
	// 852 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x8a, 0x23, 0x45,
	0x14, 0x4e, 0x27, 0x3d, 0x49, 0xe6, 0x24, 0x93, 0x99, 0x2d, 0x65, 0x2d, 0x16, 0x95, 0xd0, 0x88,
	0x04, 0x91, 0x20, 0x3b, 0x37, 0x8b, 0x08, 0xb2, 0x33, 0xbb, 0x10, 0x51, 0xb0, 0xa9, 0x88, 0xde,
	0x5a, 0xdd, 0xa9, 0xcc, 0x14, 0x9b, 0x74, 0x35, 0x55, 0xd5, 0xb3, 0xe6, 0x25, 0xbc, 0xf3, 0xc2,
	0x2b, 0xdf, 0xc1, 0x2b, 0x5f, 0xc6, 0x77, 0x91, 0xfa, 0x4b, 0xff, 0x6c, 0x1c, 0xf6, 0x2a, 0xf9,
	0xbe, 0xfa, 0xea, 0xd4, 0xa9, 0xef, 0x9c, 0x3a, 0x0d, 0x53, 0x9a, 0x6b, 0x2e, 0x8a, 0x65, 0x29,
	0x85, 0x16, 0x68, 0xc8, 0xed, 0x6f, 0xf2, 0x4f, 0x04, 0xf0, 0x93, 0xa4, 0x85, 0xda, 0x32, 0x99,
	0x66, 0xe8, 0x29, 0x0c, 0xe9, 0x5e, 0x54, 0x85, 0xc6, 0xd1, 0x3c, 0x5a, 0x4c, 0x89, 0x47, 0x86,
	0x57, 0xac, 0xd8, 0x30, 0x89, 0xfb, 0xf3, 0x68, 0x71, 0x4e, 0x3c, 0x42, 0x1f, 0xc3, 0xb9, 0x64,
	0x39, 0x2f, 0x39, 0x2b, 0x34, 0x1e, 0xd8, 0xa5, 0x9a, 0x40, 0x18, 0x46, 0x25, 0x3d, 0xec, 0x04,
	0xdd, 0xe0, 0xd8, 0x86, 0x0b, 0x10, 0x25, 0x30, 0x75, 0x11, 0xd2, 0x2a, 0xfb, 0x9e, 0x1d, 0xf0,
	0x99, 0x5d, 0x6e, 0x71, 0xe8, 0x53, 0x00, 0xae, 0x6e, 0x05, 0x2f, 0x32, 0xaa, 0x18, 0x1e, 0xce,
	0xa3, 0xc5, 0x98, 0x34, 0x98, 0xe4, 0xf7, 0x08, 0x86, 0x3f, 0x0b, 0xcd, 0xd2, 0xcc, 0xa4, 0xa1,
	0xf9, 0x9e, 0x29, 0x4d, 0xf7, 0xa5, 0xcd, 0x3c, 0x26, 0x35, 0x61, 0x02, 0x29, 0xb6, 0xdb, 0xa6,
	0x55, 0xf6, 0x86, 0x1d, 0xec, 0x05, 0xa6, 0xa4, 0xc1, 0x98, 0x64, 0x1e, 0x84, 0x66, 0xf2, 0xe5,
	0x66, 0x23, 0x99, 0x52, 0xfe, 0x1e, 0x2d, 0x2e, 0x68, 0x58, 0xd0, 0xc4, 0xb5, 0x26, 0x70, 0xc9,
	0x9f, 0x11, 0x4c, 0x5e, 0xff, 0xc6, 0xf2, 0xca, 0xf8, 0xfc, 0x88, 0x99, 0xcf, 0x60, 0xcc, 0xac,
	0x4c, 0x04, 0x3b, 0x8f, 0xd8, 0xac, 0xe5, 0xa2, 0xd0, 0x92, 0xe6, 0xc1, 0xcf, 0x23, 0x46, 0x9f,
	0xc3, 0x2c, 0xe8, 0xbc, 0x6d, 0xce, 0xd5, 0x0e, 0x8b, 0x10, 0xc4, 0x1b, 0xaa, 0xa9, 0x37, 0xd5,
	0xfe, 0x4f, 0x7e, 0x85, 0xab, 0x35, 0xcb, 0x25, 0xd3, 0xa9, 0x14, 0xa5, 0x50, 0x74, 0xe7, 0xf2,
	0xf3, 0x45, 0x8d, 0xfe, 0xbf, 0xa8, 0xfd, 0x6e, 0x51, 0xed, 0x2e, 0x13, 0x09, 0x0f, 0xe6, 0x83,
	0xc5, 0x05, 0xf1, 0x28, 0xb9, 0x85, 0x4b, 0x77, 0xc2, 0x2f, 0x5c, 0x17, 0x4c, 0xa9, 0x47, 0x0e,
	0xc0, 0x30, 0x7a, 0xeb, 0x44, 0xb8, 0x3f, 0x1f, 0x98, 0xbe, 0xf0, 0x30, 0xf9, 0xab, 0x0f, 0x97,
	0x6b, 0x4d, 0xa5, 0x5e, 0x57, 0xd9, 0xed, 0x3d, 0xe5, 0xc6, 0x46, 0x0c, 0xa3, 0xdc, 0xfc, 0xfd,
	0xee, 0x95, 0x0d, 0x73, 0x41, 0x02, 0x44, 0x0b, 0xb8, 0x54, 0x2c, 0xaf, 0x24, 0xd7, 0x87, 0x57,
	0xac, 0x14, 0x8a, 0x6b, 0x5f, 0xdd, 0x2e, 0x8d, 0xbe, 0x80, 0x2b, 0x51, 0x32, 0x49, 0x4d, 0x65,
	0x82, 0x74, 0x60, 0xa5, 0xef, 0xf0, 0x68, 0x0e, 0x13, 0x65, 0x52, 0x58, 0x31, 0x7e, 0x77, 0xaf,
	0xad, 0xc7, 0x31, 0x69, 0x52, 0x68, 0x09, 0xa8, 0xa4, 0x92, 0x15, 0x1e, 0xff, 0xb8, 0xdd, 0x2a,
	0xa6, 0xad, 0xdd, 0x31, 0x39, 0xb1, 0x62, 0x9a, 0x47, 0xbc, 0x2d, 0xea, 0x06, 0x1b, 0xba, 0xe6,
	0x69, 0x72, 0xa6, 0xb8, 0x16, 0xa7, 0x55, 0xb6, 0xe3, 0xb9, 0x29, 0xee, 0xc8, 0x15, 0xb7, 0xcd,
	0x26, 0x7f, 0x47, 0x30, 0x5b, 0x6b, 0x51, 0xbe, 0x97, 0x41, 0xa6, 0xf3, 0xb5, 0x28, 0xfd, 0x4d,
	0xfa, 0x36, 0xc1, 0x06, 0x83, 0x3e, 0x84, 0x33, 0x1b, 0xde, 0xb7, 0x9a, 0x03, 0x27, 0x52, 0x89,
	0x4f, 0xa5, 0x62, 0xed, 0xf7, 0x59, 0x84, 0x9b, 0x9d, 0xd9, 0x38, 0x5d, 0x3a, 0xf9, 0xa3, 0x0f,
	0x90, 0x56, 0xfa, 0x66, 0x27, 0xf2, 0x37, 0x69, 0x76, 0x6a, 0x63, 0x74, 0x72, 0xa3, 0xe9, 0xa0,
	0xfb, 0x66, 0xf2, 0x1e, 0xa1, 0x6b, 0x38, 0x93, 0x42, 0x68, 0x65, 0x7b, 0x70, 0xf2, 0xfc, 0x93,
	0xa5, 0x1b, 0x67, 0xcb, 0xfa, 0x90, 0x25, 0x31, 0xeb, 0xaf, 0x0b, 0x2d, 0x0f, 0xc4, 0x69, 0xcd,
	0xb1, 0xa5, 0x14, 0x9b, 0x2a, 0xaf, 0x2b, 0xe1, 0x9e, 0x71, 0x97, 0x46, 0x5f, 0xc2, 0x93, 0x40,
	0xd5, 0x26, 0xb8, 0xe7, 0xf4, 0xee, 0xc2, 0xb3, 0x17, 0x00, 0xf5, 0x61, 0xe8, 0x0a, 0x06, 0x66,
	0xcc, 0xb8, 0x0b, 0x99, 0xbf, 0xc6, 0xe5, 0x07, 0xba, 0xab, 0x98, 0x6f, 0x4e, 0x07, 0xbe, 0xee,
	0xbf, 0x88, 0x92, 0x7f, 0x63, 0x18, 0xbf, 0xcc, 0xfd, 0xb8, 0xc0, 0x30, 0x7a, 0x60, 0x52, 0x71,
	0x51, 0x84, 0x32, 0x7a, 0x68, 0x02, 0x14, 0xa2, 0xc8, 0x99, 0x37, 0xc1, 0x01, 0x33, 0x2a, 0xee,
	0xa8, 0xfa, 0x81, 0xef, 0x7d, 0x2f, 0xc7, 0xe4, 0x88, 0xfd, 0x5a, 0x2a, 0x79, 0xce, 0x7c, 0xf1,
	0x8e, 0xd8, 0x3c, 0x6f, 0xc5, 0xef, 0x0a, 0xaa, 0x2b, 0xc9, 0xfc, 0xa5, 0x6a, 0x02, 0x7d, 0x05,
	0x63, 0xed, 0xbf, 0x07, 0x18, 0xe6, 0xd1, 0x62, 0xf2, 0x1c, 0x05, 0x73, 0xeb, 0xef, 0xc4, 0xaa,
	0x47, 0x8e, 0x2a, 0xf4, 0x19, 0xc4, 0x66, 0x0c, 0xe2, 0x89, 0x55, 0xcf, 0x82, 0xda, 0x8d, 0xe6,
	0x55, 0x8f, 0xd8, 0x55, 0x74, 0x0d, 0xe7, 0x2c, 0xcc, 0x46, 0x3c, 0xb5, 0xd2, 0x0f, 0x82, 0xb4,
	0x31, 0x34, 0x57, 0x3d, 0x52, 0xeb, 0xd0, 0x0d, 0xcc, 0x54, 0x6b, 0x6a, 0xe1, 0x0b, 0xbb, 0x13,
	0x87, 0x9d, 0xdd, 0x99, 0xb6, 0xea, 0x91, 0xce, 0x0e, 0xf4, 0x2d, 0x5c, 0xa8, 0xe6, 0x5c, 0xc2,
	0x33, 0x1b, 0xe2, 0xa3, 0x76, 0x88, 0xe3, 0xd0, 0x5a, 0xf5, 0x48, 0x5b, 0x6f, 0x03, 0x34, 0x47,
	0x12, 0xbe, 0xec, 0x04, 0x68, 0xcf, 0x2b, 0x1b, 0xa0, 0x49, 0xa1, 0x6f, 0x60, 0xaa, 0x1a, 0x2f,
	0x16, 0x5f, 0xd9, 0xfd, 0x4f, 0xeb, 0xfd, 0xcd, 0xd7, 0xbc, 0xea, 0x91, 0x96, 0xda, 0x14, 0xa4,
	0xf4, 0x5d, 0x8d, 0x9f, 0xb4, 0x0b, 0x52, 0x77, 0xbb, 0x29, 0x48, 0x50, 0xdd, 0x8c, 0x61, 0xe8,
	0xbe, 0xf5, 0xd9, 0xd0, 0xea, 0xae, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x27, 0xdf, 0x9c, 0x8e,
	0xfc, 0x07, 0x00, 0x00,
}