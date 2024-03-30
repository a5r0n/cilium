// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/cilium/cilium/pkg/k8s/slim/k8s/apis/meta/v1beta1/generated.proto

package v1beta1

import (
	fmt "fmt"

	v1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/apis/meta/v1"
	proto "github.com/gogo/protobuf/proto"

	math "math"
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

func (m *PartialObjectMetadataList) Reset()         { *m = PartialObjectMetadataList{} }
func (m *PartialObjectMetadataList) String() string { return proto.CompactTextString(m) }
func (*PartialObjectMetadataList) ProtoMessage()    {}
func (*PartialObjectMetadataList) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a84ae209524fd15, []int{0}
}
func (m *PartialObjectMetadataList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartialObjectMetadataList.Unmarshal(m, b)
}
func (m *PartialObjectMetadataList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartialObjectMetadataList.Marshal(b, m, deterministic)
}
func (m *PartialObjectMetadataList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartialObjectMetadataList.Merge(m, src)
}
func (m *PartialObjectMetadataList) XXX_Size() int {
	return xxx_messageInfo_PartialObjectMetadataList.Size(m)
}
func (m *PartialObjectMetadataList) XXX_DiscardUnknown() {
	xxx_messageInfo_PartialObjectMetadataList.DiscardUnknown(m)
}

var xxx_messageInfo_PartialObjectMetadataList proto.InternalMessageInfo

func (m *PartialObjectMetadataList) GetMetadata() *v1.ListMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *PartialObjectMetadataList) GetItems() []*v1.PartialObjectMetadata {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*PartialObjectMetadataList)(nil), "github.com.cilium.cilium.pkg.k8s.slim.k8s.apis.meta.v1beta1.PartialObjectMetadataList")
}

func init() {
	proto.RegisterFile("github.com/cilium/cilium/pkg/k8s/slim/k8s/apis/meta/v1beta1/generated.proto", fileDescriptor_1a84ae209524fd15)
}

var fileDescriptor_1a84ae209524fd15 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0xd1, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x06, 0x60, 0xa2, 0x08, 0x92, 0xde, 0x72, 0xaa, 0x3d, 0x15, 0x4f, 0x3d, 0xcd, 0xd2, 0x1e,
	0xa4, 0x50, 0x04, 0xf1, 0xe0, 0x45, 0x8b, 0xd2, 0xa3, 0x78, 0x99, 0x6c, 0x87, 0x64, 0x4c, 0x26,
	0x59, 0x76, 0x27, 0x82, 0x2f, 0xeb, 0xb3, 0xc8, 0xa6, 0xb5, 0x42, 0xf0, 0x14, 0x4f, 0xf3, 0x2f,
	0xec, 0x7e, 0xfc, 0xcb, 0xa4, 0x8f, 0x05, 0x6b, 0xd9, 0xe5, 0x60, 0x5b, 0x31, 0x96, 0x6b, 0xee,
	0x4e, 0xc3, 0x55, 0x85, 0xa9, 0xd6, 0xc1, 0x84, 0x9a, 0xa5, 0x0f, 0xe8, 0x38, 0x18, 0x21, 0x45,
	0xf3, 0xb1, 0xcc, 0x49, 0x71, 0x69, 0x0a, 0x6a, 0xc8, 0xa3, 0xd2, 0x1e, 0x9c, 0x6f, 0xb5, 0xcd,
	0x36, 0xbf, 0x18, 0x1c, 0x94, 0x9f, 0xe1, 0xaa, 0x02, 0xaa, 0x75, 0x80, 0x88, 0xf5, 0x21, 0x62,
	0x10, 0x31, 0x38, 0x62, 0xb3, 0x87, 0x71, 0x4d, 0x86, 0x25, 0x66, 0x37, 0x91, 0xe7, 0x36, 0x5e,
	0x12, 0xb4, 0x25, 0x37, 0xe4, 0x3f, 0x7b, 0xc2, 0x77, 0x8d, 0xb2, 0x90, 0x09, 0xb6, 0x24, 0xc1,
	0xe1, 0xbb, 0xeb, 0xaf, 0x24, 0xbd, 0x7a, 0x41, 0xaf, 0x8c, 0xf5, 0x73, 0xfe, 0x4e, 0x56, 0xb7,
	0xa4, 0xb8, 0x47, 0xc5, 0x27, 0x0e, 0x9a, 0xbd, 0xa5, 0x97, 0x72, 0x3c, 0x4f, 0xcf, 0xe6, 0xc9,
	0x62, 0xb2, 0xba, 0x83, 0x71, 0xbf, 0x85, 0xe8, 0x45, 0x7b, 0x77, 0x12, 0x33, 0x9b, 0x5e, 0xb0,
	0x92, 0x84, 0x69, 0x32, 0x3f, 0x5f, 0x4c, 0x56, 0xdb, 0xb1, 0xf4, 0x9f, 0xfd, 0x77, 0x07, 0xfb,
	0xfe, 0xf6, 0x75, 0xf3, 0x8f, 0x65, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x97, 0x72, 0x2a,
	0x2a, 0x02, 0x00, 0x00,
}
