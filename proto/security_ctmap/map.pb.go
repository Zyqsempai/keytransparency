// Code generated by protoc-gen-go.
// source: proto/security_ctmap/map.proto
// DO NOT EDIT!

/*
Package security_ctmap is a generated protocol buffer package.

It is generated from these files:
	proto/security_ctmap/map.proto

It has these top-level messages:
	EpochHead
	Step
	Entry
	PublicKey
	SignedEpochHead
	SignedEntryUpdate
*/
package security_ctmap

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import security_protobuf "github.com/google/e2e-key-server/proto/security_protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// EpochHead is the head node of the Merkle Tree as well as additional metadata
// for the tree.
type EpochHead struct {
	// Realm is the domain...
	Realm string `protobuf:"bytes,1,opt,name=realm" json:"realm,omitempty"`
	// Epoch number
	Epoch int64 `protobuf:"varint,2,opt,name=epoch" json:"epoch,omitempty"`
	// Root is the value of the root node of the merkle tree.
	Root []byte `protobuf:"bytes,3,opt,name=root,proto3" json:"root,omitempty"`
	// IssueTime is the time when this epoch was released. All epochs for the
	// same keyserver MUST have non-decreasing IssueTimes.
	IssueTime *security_protobuf.Timestamp `protobuf:"bytes,4,opt,name=issue_time" json:"issue_time,omitempty"`
	// Hash of previous SEH. SHA512_256.
	PreviousHash []byte `protobuf:"bytes,5,opt,name=previous_hash,proto3" json:"previous_hash,omitempty"`
}

func (m *EpochHead) Reset()                    { *m = EpochHead{} }
func (m *EpochHead) String() string            { return proto.CompactTextString(m) }
func (*EpochHead) ProtoMessage()               {}
func (*EpochHead) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EpochHead) GetIssueTime() *security_protobuf.Timestamp {
	if m != nil {
		return m.IssueTime
	}
	return nil
}

// Step is a combined, ordered list of SignedEntryUpdates and SignedEpochHeads
// which are made available to verifiers.
type Step struct {
	// Types that are valid to be assigned to Type:
	//	*Step_EntryChanged
	//	*Step_SignedEpochHead
	Type isStep_Type `protobuf_oneof:"type"`
	// epoch of this udpate.
	Epoch int64 `protobuf:"varint,3,opt,name=epoch" json:"epoch,omitempty"`
	// commitment_timestamp is the ordered commitment_timestamp of this step.
	CommitmentTimestamp int64 `protobuf:"varint,4,opt,name=commitment_timestamp" json:"commitment_timestamp,omitempty"`
}

func (m *Step) Reset()                    { *m = Step{} }
func (m *Step) String() string            { return proto.CompactTextString(m) }
func (*Step) ProtoMessage()               {}
func (*Step) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isStep_Type interface {
	isStep_Type()
}

type Step_EntryChanged struct {
	EntryChanged []byte `protobuf:"bytes,1,opt,name=entry_changed,proto3,oneof"`
}
type Step_SignedEpochHead struct {
	SignedEpochHead *SignedEpochHead `protobuf:"bytes,2,opt,name=signed_epoch_head,oneof"`
}

func (*Step_EntryChanged) isStep_Type()    {}
func (*Step_SignedEpochHead) isStep_Type() {}

func (m *Step) GetType() isStep_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Step) GetEntryChanged() []byte {
	if x, ok := m.GetType().(*Step_EntryChanged); ok {
		return x.EntryChanged
	}
	return nil
}

func (m *Step) GetSignedEpochHead() *SignedEpochHead {
	if x, ok := m.GetType().(*Step_SignedEpochHead); ok {
		return x.SignedEpochHead
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Step) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Step_OneofMarshaler, _Step_OneofUnmarshaler, _Step_OneofSizer, []interface{}{
		(*Step_EntryChanged)(nil),
		(*Step_SignedEpochHead)(nil),
	}
}

func _Step_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Step)
	// type
	switch x := m.Type.(type) {
	case *Step_EntryChanged:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.EntryChanged)
	case *Step_SignedEpochHead:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SignedEpochHead); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Step.Type has unexpected type %T", x)
	}
	return nil
}

func _Step_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Step)
	switch tag {
	case 1: // type.entry_changed
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Type = &Step_EntryChanged{x}
		return true, err
	case 2: // type.signed_epoch_head
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SignedEpochHead)
		err := b.DecodeMessage(msg)
		m.Type = &Step_SignedEpochHead{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Step_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Step)
	// type
	switch x := m.Type.(type) {
	case *Step_EntryChanged:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.EntryChanged)))
		n += len(x.EntryChanged)
	case *Step_SignedEpochHead:
		s := proto.Size(x.SignedEpochHead)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Entry is the leaf node object for the Merkle Tree. Its unique index in the
// tree is identified by a hash of an verifiable unpredictable function on the
// user_id.
type Entry struct {
	// Index is the location in the merkle tree for this entry.
	// If signing keys are not unique per user, we need to tie updates to a
	// particular profile.
	Index []byte `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	// UpdateCount prevents replaying old signed EntryUpdates.
	// not nessesarilly incremented by only one each update.
	UpdateCount uint64 `protobuf:"varint,2,opt,name=update_count" json:"update_count,omitempty"`
	// EntryKey allows verifiers to validate updates to Entry.
	EntryKey []*PublicKey `protobuf:"bytes,3,rep,name=entry_key" json:"entry_key,omitempty"`
	// profile_commitment is a cryptographic commitment to the Profile of the form
	// HMAC(profile_commitment_key, serialized_profile)
	ProfileCommitment []byte `protobuf:"bytes,4,opt,name=profile_commitment,proto3" json:"profile_commitment,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Entry) GetEntryKey() []*PublicKey {
	if m != nil {
		return m.EntryKey
	}
	return nil
}

// PublicKey defines a key this domain uses to sign EpochHeads with.
type PublicKey struct {
	// KeyFormats from Keyczar.
	//
	// Types that are valid to be assigned to KeyType:
	//	*PublicKey_Ed25519
	//	*PublicKey_RsaVerifyingSha256_2048
	//	*PublicKey_EcdsaVerifyingP256
	KeyType isPublicKey_KeyType `protobuf_oneof:"key_type"`
}

func (m *PublicKey) Reset()                    { *m = PublicKey{} }
func (m *PublicKey) String() string            { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()               {}
func (*PublicKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isPublicKey_KeyType interface {
	isPublicKey_KeyType()
}

type PublicKey_Ed25519 struct {
	Ed25519 []byte `protobuf:"bytes,1,opt,name=ed25519,proto3,oneof"`
}
type PublicKey_RsaVerifyingSha256_2048 struct {
	RsaVerifyingSha256_2048 []byte `protobuf:"bytes,2,opt,name=rsa_verifying_sha256_2048,proto3,oneof"`
}
type PublicKey_EcdsaVerifyingP256 struct {
	EcdsaVerifyingP256 []byte `protobuf:"bytes,3,opt,name=ecdsa_verifying_p256,proto3,oneof"`
}

func (*PublicKey_Ed25519) isPublicKey_KeyType()                 {}
func (*PublicKey_RsaVerifyingSha256_2048) isPublicKey_KeyType() {}
func (*PublicKey_EcdsaVerifyingP256) isPublicKey_KeyType()      {}

func (m *PublicKey) GetKeyType() isPublicKey_KeyType {
	if m != nil {
		return m.KeyType
	}
	return nil
}

func (m *PublicKey) GetEd25519() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_Ed25519); ok {
		return x.Ed25519
	}
	return nil
}

func (m *PublicKey) GetRsaVerifyingSha256_2048() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_RsaVerifyingSha256_2048); ok {
		return x.RsaVerifyingSha256_2048
	}
	return nil
}

func (m *PublicKey) GetEcdsaVerifyingP256() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_EcdsaVerifyingP256); ok {
		return x.EcdsaVerifyingP256
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PublicKey) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PublicKey_OneofMarshaler, _PublicKey_OneofUnmarshaler, _PublicKey_OneofSizer, []interface{}{
		(*PublicKey_Ed25519)(nil),
		(*PublicKey_RsaVerifyingSha256_2048)(nil),
		(*PublicKey_EcdsaVerifyingP256)(nil),
	}
}

func _PublicKey_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PublicKey)
	// key_type
	switch x := m.KeyType.(type) {
	case *PublicKey_Ed25519:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Ed25519)
	case *PublicKey_RsaVerifyingSha256_2048:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.RsaVerifyingSha256_2048)
	case *PublicKey_EcdsaVerifyingP256:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.EcdsaVerifyingP256)
	case nil:
	default:
		return fmt.Errorf("PublicKey.KeyType has unexpected type %T", x)
	}
	return nil
}

func _PublicKey_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PublicKey)
	switch tag {
	case 1: // key_type.ed25519
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_Ed25519{x}
		return true, err
	case 2: // key_type.rsa_verifying_sha256_2048
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_RsaVerifyingSha256_2048{x}
		return true, err
	case 3: // key_type.ecdsa_verifying_p256
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_EcdsaVerifyingP256{x}
		return true, err
	default:
		return false, nil
	}
}

func _PublicKey_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PublicKey)
	// key_type
	switch x := m.KeyType.(type) {
	case *PublicKey_Ed25519:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Ed25519)))
		n += len(x.Ed25519)
	case *PublicKey_RsaVerifyingSha256_2048:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.RsaVerifyingSha256_2048)))
		n += len(x.RsaVerifyingSha256_2048)
	case *PublicKey_EcdsaVerifyingP256:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.EcdsaVerifyingP256)))
		n += len(x.EcdsaVerifyingP256)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// SignedEpochHead represents a signed state of the Merkel tree.
type SignedEpochHead struct {
	// Serialized EpochHead.
	EpochHead []byte `protobuf:"bytes,1,opt,name=epoch_head,proto3" json:"epoch_head,omitempty"`
	// Signature of head, using the signature type of the key.
	// keyed by the first 64 bits bytes of the hash of the key.
	Signatures map[uint64][]byte `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty" protobuf_key:"fixed64,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *SignedEpochHead) Reset()                    { *m = SignedEpochHead{} }
func (m *SignedEpochHead) String() string            { return proto.CompactTextString(m) }
func (*SignedEpochHead) ProtoMessage()               {}
func (*SignedEpochHead) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SignedEpochHead) GetSignatures() map[uint64][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}

// SignedEntryUpdate is what users provide to update their profiles.
// A serialized SignedEntryUpdate is used as the leaf value in the MerkleTree.
type SignedEntryUpdate struct {
	// NewEntry is the serialized protobuf Entry.
	NewEntry []byte `protobuf:"bytes,1,opt,name=new_entry,proto3" json:"new_entry,omitempty"`
	// Signature of entry, by the entry_key inside entry AND the old key from the
	// previous epoch. The first proves ownership of new epoch key, and the
	// second proves the the correct owner is making this change.
	Signatures map[uint64][]byte `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty" protobuf_key:"fixed64,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *SignedEntryUpdate) Reset()                    { *m = SignedEntryUpdate{} }
func (m *SignedEntryUpdate) String() string            { return proto.CompactTextString(m) }
func (*SignedEntryUpdate) ProtoMessage()               {}
func (*SignedEntryUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SignedEntryUpdate) GetSignatures() map[uint64][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func init() {
	proto.RegisterType((*EpochHead)(nil), "security_ctmap.EpochHead")
	proto.RegisterType((*Step)(nil), "security_ctmap.Step")
	proto.RegisterType((*Entry)(nil), "security_ctmap.Entry")
	proto.RegisterType((*PublicKey)(nil), "security_ctmap.PublicKey")
	proto.RegisterType((*SignedEpochHead)(nil), "security_ctmap.SignedEpochHead")
	proto.RegisterType((*SignedEntryUpdate)(nil), "security_ctmap.SignedEntryUpdate")
}

var fileDescriptor0 = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x97, 0xa5, 0x2b, 0xf4, 0xb5, 0x05, 0xd5, 0x2a, 0xa2, 0xab, 0xa6, 0x81, 0xca, 0x85,
	0x03, 0x4b, 0xd6, 0x40, 0xd1, 0x40, 0x9c, 0x86, 0x2a, 0x55, 0xe2, 0x82, 0x34, 0x38, 0x5b, 0x6e,
	0xf2, 0x9a, 0x58, 0x34, 0x71, 0x14, 0xdb, 0x85, 0xde, 0xf8, 0x1b, 0xb8, 0x23, 0xf1, 0xa7, 0x62,
	0x3b, 0x6d, 0xb7, 0x56, 0x20, 0x2e, 0x1c, 0x72, 0x78, 0xf6, 0xf3, 0xfb, 0x7e, 0xdf, 0xf7, 0x02,
	0xe7, 0x65, 0x25, 0x94, 0x08, 0x25, 0xc6, 0xba, 0xe2, 0x6a, 0x4d, 0x63, 0x95, 0xb3, 0x32, 0x34,
	0x5f, 0xe0, 0x2e, 0xc8, 0x83, 0xfd, 0x9b, 0xe1, 0x2c, 0xe5, 0x2a, 0xd3, 0xf3, 0x20, 0x16, 0x79,
	0x98, 0x0a, 0x91, 0x2e, 0x31, 0xc4, 0x08, 0x2f, 0xbe, 0xe0, 0xfa, 0x42, 0x62, 0xb5, 0xc2, 0x2a,
	0x3c, 0x18, 0xe8, 0xca, 0xb9, 0x5e, 0x84, 0x8a, 0xe7, 0x28, 0x15, 0xcb, 0x37, 0x93, 0x47, 0xdf,
	0x3d, 0x68, 0x4d, 0x4b, 0x11, 0x67, 0x33, 0x64, 0x09, 0xe9, 0xc2, 0x49, 0x85, 0x6c, 0x99, 0x0f,
	0xbc, 0xa7, 0xde, 0xf3, 0x96, 0x2d, 0xd1, 0xde, 0x0d, 0x8e, 0x4d, 0xe9, 0x93, 0x0e, 0x34, 0x2a,
	0x21, 0xd4, 0xc0, 0x37, 0x55, 0x87, 0x5c, 0x02, 0x70, 0x29, 0x35, 0x52, 0x3b, 0x72, 0xd0, 0x30,
	0x67, 0xed, 0xe8, 0x2c, 0xd8, 0x2a, 0x06, 0x5b, 0xc5, 0xe0, 0xd3, 0x56, 0x91, 0x3c, 0x82, 0x6e,
	0x59, 0xe1, 0x8a, 0x0b, 0x2d, 0x69, 0xc6, 0x64, 0x36, 0x38, 0xb1, 0x83, 0x46, 0x3f, 0x3c, 0x68,
	0xdc, 0x28, 0x2c, 0xc9, 0x63, 0xe8, 0x62, 0xa1, 0x2a, 0x63, 0x32, 0x63, 0x45, 0x8a, 0x89, 0xa3,
	0xe8, 0xcc, 0x8e, 0xc8, 0x3b, 0xe8, 0x49, 0x9e, 0x16, 0x98, 0x50, 0x87, 0x43, 0x33, 0xc3, 0xea,
	0x98, 0xda, 0xd1, 0x93, 0x60, 0x3f, 0x9a, 0xe0, 0xc6, 0x35, 0xee, 0x2c, 0x99, 0xd7, 0x3b, 0x17,
	0xbe, 0x73, 0x71, 0x06, 0x7d, 0x13, 0x5b, 0xce, 0x55, 0x6e, 0xc4, 0xe8, 0x2e, 0x0f, 0xe7, 0xc0,
	0xbf, 0x6e, 0x42, 0x43, 0xad, 0x4b, 0x1c, 0xad, 0xe0, 0x64, 0x6a, 0x59, 0xec, 0x6b, 0x5e, 0x24,
	0xf8, 0xad, 0x86, 0x21, 0x7d, 0xe8, 0xe8, 0x32, 0x61, 0x0a, 0x69, 0x2c, 0x74, 0xa1, 0x1c, 0x45,
	0x83, 0xbc, 0x80, 0x56, 0x4d, 0x6e, 0x36, 0x60, 0x64, 0x7c, 0x03, 0x76, 0x7a, 0x08, 0xf6, 0x51,
	0xcf, 0x97, 0x3c, 0xfe, 0x80, 0x6b, 0x32, 0x04, 0x62, 0xd2, 0x59, 0xf0, 0xa5, 0x1d, 0xb2, 0x25,
	0x71, 0xfa, 0x9d, 0x91, 0x80, 0xd6, 0x6d, 0x63, 0x0f, 0xee, 0x61, 0x12, 0x4d, 0x26, 0xe3, 0x37,
	0xbb, 0x28, 0x9e, 0xc1, 0x69, 0x25, 0x19, 0x35, 0x2b, 0xe6, 0x8b, 0x35, 0x2f, 0x52, 0x2a, 0x33,
	0x16, 0x4d, 0x5e, 0xd3, 0xe8, 0xf2, 0xd5, 0x95, 0x83, 0xb1, 0x4d, 0xe7, 0xd0, 0xc7, 0x38, 0xd9,
	0x6b, 0x2b, 0x4d, 0x53, 0xbd, 0xb8, 0xd9, 0xd1, 0x35, 0xc0, 0x7d, 0x03, 0x4a, 0x9d, 0xd1, 0x9f,
	0x1e, 0x3c, 0x3c, 0xc8, 0x8c, 0x10, 0x80, 0x3b, 0x41, 0xd7, 0xc6, 0xdf, 0x03, 0xd8, 0x1d, 0x30,
	0xa5, 0x2b, 0x94, 0x46, 0xc9, 0x7a, 0x0c, 0xff, 0x11, 0xbe, 0xab, 0xeb, 0x17, 0x2e, 0xcc, 0xe1,
	0xb8, 0xd6, 0xba, 0x73, 0x44, 0xda, 0xe0, 0xdb, 0xd0, 0xac, 0x48, 0xd3, 0x86, 0xbd, 0x62, 0x4b,
	0x8d, 0xb5, 0x93, 0xb7, 0xc7, 0x57, 0xde, 0xe8, 0x97, 0x07, 0xbd, 0xcd, 0x58, 0xdb, 0xff, 0xd9,
	0x85, 0x6f, 0x92, 0x69, 0x15, 0xf8, 0x95, 0xba, 0xd0, 0x37, 0x80, 0xd3, 0x3f, 0x00, 0x8e, 0xff,
	0x02, 0x78, 0x3b, 0xe9, 0x3f, 0x20, 0xce, 0x9b, 0xee, 0x5f, 0x7f, 0xf9, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0xf1, 0xe6, 0xcf, 0x6a, 0xc6, 0x03, 0x00, 0x00,
}
