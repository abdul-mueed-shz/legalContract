// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: msp/identities.proto

package msp

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This struct represents an Identity
// (with its MSP identifier) to be used
// to serialize it and deserialize it
type SerializedIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The identifier of the associated membership service provider
	Mspid string `protobuf:"bytes,1,opt,name=mspid,proto3" json:"mspid,omitempty"`
	// the Identity, serialized according to the rules of its MPS
	IdBytes []byte `protobuf:"bytes,2,opt,name=id_bytes,json=idBytes,proto3" json:"id_bytes,omitempty"`
}

func (x *SerializedIdentity) Reset() {
	*x = SerializedIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msp_identities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SerializedIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SerializedIdentity) ProtoMessage() {}

func (x *SerializedIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_msp_identities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SerializedIdentity.ProtoReflect.Descriptor instead.
func (*SerializedIdentity) Descriptor() ([]byte, []int) {
	return file_msp_identities_proto_rawDescGZIP(), []int{0}
}

func (x *SerializedIdentity) GetMspid() string {
	if x != nil {
		return x.Mspid
	}
	return ""
}

func (x *SerializedIdentity) GetIdBytes() []byte {
	if x != nil {
		return x.IdBytes
	}
	return nil
}

// This struct represents an Idemix Identity
// to be used to serialize it and deserialize it.
// The IdemixMSP will first serialize an idemix identity to bytes using
// this proto, and then uses these bytes as id_bytes in SerializedIdentity
type SerializedIdemixIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// nym_x is the X-component of the pseudonym elliptic curve point.
	// It is a []byte representation of an amcl.BIG
	// The pseudonym can be seen as a public key of the identity, it is used to verify signatures.
	NymX []byte `protobuf:"bytes,1,opt,name=nym_x,json=nymX,proto3" json:"nym_x,omitempty"`
	// nym_y is the Y-component of the pseudonym elliptic curve point.
	// It is a []byte representation of an amcl.BIG
	// The pseudonym can be seen as a public key of the identity, it is used to verify signatures.
	NymY []byte `protobuf:"bytes,2,opt,name=nym_y,json=nymY,proto3" json:"nym_y,omitempty"`
	// ou contains the organizational unit of the idemix identity
	Ou []byte `protobuf:"bytes,3,opt,name=ou,proto3" json:"ou,omitempty"`
	// role contains the role of this identity (e.g., ADMIN or MEMBER)
	Role []byte `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	// proof contains the cryptographic evidence that this identity is valid
	Proof []byte `protobuf:"bytes,5,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (x *SerializedIdemixIdentity) Reset() {
	*x = SerializedIdemixIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msp_identities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SerializedIdemixIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SerializedIdemixIdentity) ProtoMessage() {}

func (x *SerializedIdemixIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_msp_identities_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SerializedIdemixIdentity.ProtoReflect.Descriptor instead.
func (*SerializedIdemixIdentity) Descriptor() ([]byte, []int) {
	return file_msp_identities_proto_rawDescGZIP(), []int{1}
}

func (x *SerializedIdemixIdentity) GetNymX() []byte {
	if x != nil {
		return x.NymX
	}
	return nil
}

func (x *SerializedIdemixIdentity) GetNymY() []byte {
	if x != nil {
		return x.NymY
	}
	return nil
}

func (x *SerializedIdemixIdentity) GetOu() []byte {
	if x != nil {
		return x.Ou
	}
	return nil
}

func (x *SerializedIdemixIdentity) GetRole() []byte {
	if x != nil {
		return x.Role
	}
	return nil
}

func (x *SerializedIdemixIdentity) GetProof() []byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

var File_msp_identities_proto protoreflect.FileDescriptor

var file_msp_identities_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x73, 0x70, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x73, 0x70, 0x22, 0x45, 0x0a, 0x12, 0x53,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6d, 0x73, 0x70, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x64, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x69, 0x64, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x22, 0x7e, 0x0a, 0x18, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64,
	0x49, 0x64, 0x65, 0x6d, 0x69, 0x78, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x13,
	0x0a, 0x05, 0x6e, 0x79, 0x6d, 0x5f, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6e,
	0x79, 0x6d, 0x58, 0x12, 0x13, 0x0a, 0x05, 0x6e, 0x79, 0x6d, 0x5f, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x6e, 0x79, 0x6d, 0x59, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x75, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x6f, 0x75, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x72, 0x6f,
	0x6f, 0x66, 0x42, 0x93, 0x01, 0x0a, 0x21, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x42, 0x0f, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2d, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6d, 0x73, 0x70, 0xa2, 0x02,
	0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x03, 0x4d, 0x73, 0x70, 0xca, 0x02, 0x03, 0x4d, 0x73, 0x70,
	0xe2, 0x02, 0x0f, 0x4d, 0x73, 0x70, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x03, 0x4d, 0x73, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msp_identities_proto_rawDescOnce sync.Once
	file_msp_identities_proto_rawDescData = file_msp_identities_proto_rawDesc
)

func file_msp_identities_proto_rawDescGZIP() []byte {
	file_msp_identities_proto_rawDescOnce.Do(func() {
		file_msp_identities_proto_rawDescData = protoimpl.X.CompressGZIP(file_msp_identities_proto_rawDescData)
	})
	return file_msp_identities_proto_rawDescData
}

var file_msp_identities_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_msp_identities_proto_goTypes = []interface{}{
	(*SerializedIdentity)(nil),       // 0: msp.SerializedIdentity
	(*SerializedIdemixIdentity)(nil), // 1: msp.SerializedIdemixIdentity
}
var file_msp_identities_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_msp_identities_proto_init() }
func file_msp_identities_proto_init() {
	if File_msp_identities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msp_identities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SerializedIdentity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msp_identities_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SerializedIdemixIdentity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msp_identities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msp_identities_proto_goTypes,
		DependencyIndexes: file_msp_identities_proto_depIdxs,
		MessageInfos:      file_msp_identities_proto_msgTypes,
	}.Build()
	File_msp_identities_proto = out.File
	file_msp_identities_proto_rawDesc = nil
	file_msp_identities_proto_goTypes = nil
	file_msp_identities_proto_depIdxs = nil
}
