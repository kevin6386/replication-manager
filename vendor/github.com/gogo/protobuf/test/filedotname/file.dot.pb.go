// Code generated by protoc-gen-gogo.
// source: file.dot.proto
// DO NOT EDIT!

/*
Package filedotname is a generated protocol buffer package.

It is generated from these files:
	file.dot.proto

It has these top-level messages:
	M
*/
package filedotname

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_gogo_protobuf_protoc_gen_gogo_descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import compress_gzip "compress/gzip"
import bytes "bytes"
import io_ioutil "io/ioutil"

import strings "strings"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type M struct {
	A                *string `protobuf:"bytes,1,opt,name=a" json:"a,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *M) Reset()                    { *m = M{} }
func (*M) ProtoMessage()               {}
func (*M) Descriptor() ([]byte, []int) { return fileDescriptorFileDot, []int{0} }

func init() {
	proto.RegisterType((*M)(nil), "filedotname.M")
}
func (this *M) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return FileDotDescription()
}
func FileDotDescription() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	d := &github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet{}
	var gzipped = []byte{
		// 3368 bytes of a gzipped FileDescriptorSet
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x5a, 0x5b, 0x6c, 0x1c, 0x57,
		0x19, 0xf6, 0xec, 0xc5, 0xde, 0xfd, 0x77, 0xbd, 0x1e, 0x8f, 0x5d, 0x67, 0xe3, 0xb6, 0x1b, 0xc7,
		0xbd, 0xb9, 0x29, 0x75, 0x4a, 0x9a, 0xa4, 0xc9, 0x86, 0x36, 0x5a, 0xdb, 0x1b, 0xd7, 0x91, 0x6f,
		0xcc, 0xda, 0x6d, 0x5a, 0x1e, 0x46, 0xc7, 0xb3, 0x67, 0xd7, 0x93, 0xcc, 0xce, 0x2c, 0x33, 0xb3,
		0x49, 0x9c, 0xa7, 0xa0, 0x72, 0x51, 0x85, 0xb8, 0x23, 0xd1, 0x3b, 0xb4, 0x12, 0xb4, 0x14, 0x0a,
		0x2d, 0x37, 0x21, 0x9e, 0x2a, 0xa1, 0x42, 0x9f, 0x50, 0xe1, 0x89, 0x07, 0x1e, 0x1a, 0x53, 0x89,
		0x02, 0x05, 0x8a, 0x14, 0x89, 0x4a, 0x7d, 0x41, 0xe7, 0x36, 0x3b, 0x7b, 0x71, 0x66, 0x5d, 0xa9,
		0x94, 0x27, 0xef, 0xf9, 0xcf, 0xff, 0x7d, 0xf3, 0x9f, 0xff, 0xfc, 0xe7, 0xff, 0xff, 0x39, 0x63,
		0xf8, 0xf5, 0xc7, 0x61, 0xa2, 0x6a, 0xdb, 0x55, 0x13, 0x1f, 0xac, 0x3b, 0xb6, 0x67, 0x6f, 0x34,
		0x2a, 0x07, 0xcb, 0xd8, 0xd5, 0x1d, 0xa3, 0xee, 0xd9, 0xce, 0x34, 0x95, 0x29, 0x43, 0x4c, 0x63,
		0x5a, 0x68, 0x4c, 0x2e, 0xc1, 0xf0, 0x29, 0xc3, 0xc4, 0x73, 0xbe, 0x62, 0x09, 0x7b, 0xca, 0x31,
		0x88, 0x55, 0x0c, 0x13, 0x67, 0xa5, 0x89, 0xe8, 0x54, 0xea, 0xd0, 0xcd, 0xd3, 0x6d, 0xa0, 0xe9,
		0x56, 0xc4, 0x2a, 0x11, 0xab, 0x14, 0x31, 0xf9, 0x56, 0x0c, 0x46, 0xba, 0xcc, 0x2a, 0x0a, 0xc4,
		0x2c, 0x54, 0x23, 0x8c, 0xd2, 0x54, 0x52, 0xa5, 0xbf, 0x95, 0x2c, 0x0c, 0xd4, 0x91, 0x7e, 0x0e,
		0x55, 0x71, 0x36, 0x42, 0xc5, 0x62, 0xa8, 0xe4, 0x00, 0xca, 0xb8, 0x8e, 0xad, 0x32, 0xb6, 0xf4,
		0xad, 0x6c, 0x74, 0x22, 0x3a, 0x95, 0x54, 0x03, 0x12, 0xe5, 0x0e, 0x18, 0xae, 0x37, 0x36, 0x4c,
		0x43, 0xd7, 0x02, 0x6a, 0x30, 0x11, 0x9d, 0x8a, 0xab, 0x32, 0x9b, 0x98, 0x6b, 0x2a, 0xdf, 0x06,
		0x43, 0x17, 0x30, 0x3a, 0x17, 0x54, 0x4d, 0x51, 0xd5, 0x0c, 0x11, 0x07, 0x14, 0x67, 0x21, 0x5d,
		0xc3, 0xae, 0x8b, 0xaa, 0x58, 0xf3, 0xb6, 0xea, 0x38, 0x1b, 0xa3, 0xab, 0x9f, 0xe8, 0x58, 0x7d,
		0xfb, 0xca, 0x53, 0x1c, 0xb5, 0xb6, 0x55, 0xc7, 0x4a, 0x01, 0x92, 0xd8, 0x6a, 0xd4, 0x18, 0x43,
		0x7c, 0x07, 0xff, 0x15, 0xad, 0x46, 0xad, 0x9d, 0x25, 0x41, 0x60, 0x9c, 0x62, 0xc0, 0xc5, 0xce,
		0x79, 0x43, 0xc7, 0xd9, 0x7e, 0x4a, 0x70, 0x5b, 0x07, 0x41, 0x89, 0xcd, 0xb7, 0x73, 0x08, 0x9c,
		0x32, 0x0b, 0x49, 0x7c, 0xd1, 0xc3, 0x96, 0x6b, 0xd8, 0x56, 0x76, 0x80, 0x92, 0xdc, 0xd2, 0x65,
		0x17, 0xb1, 0x59, 0x6e, 0xa7, 0x68, 0xe2, 0x94, 0xa3, 0x30, 0x60, 0xd7, 0x3d, 0xc3, 0xb6, 0xdc,
		0x6c, 0x62, 0x42, 0x9a, 0x4a, 0x1d, 0xba, 0xa1, 0x6b, 0x20, 0xac, 0x30, 0x1d, 0x55, 0x28, 0x2b,
		0x0b, 0x20, 0xbb, 0x76, 0xc3, 0xd1, 0xb1, 0xa6, 0xdb, 0x65, 0xac, 0x19, 0x56, 0xc5, 0xce, 0x26,
		0x29, 0xc1, 0xbe, 0xce, 0x85, 0x50, 0xc5, 0x59, 0xbb, 0x8c, 0x17, 0xac, 0x8a, 0xad, 0x66, 0xdc,
		0x96, 0xb1, 0x32, 0x06, 0xfd, 0xee, 0x96, 0xe5, 0xa1, 0x8b, 0xd9, 0x34, 0x8d, 0x10, 0x3e, 0x9a,
		0xfc, 0x4f, 0x1c, 0x86, 0x7a, 0x09, 0xb1, 0x13, 0x10, 0xaf, 0x90, 0x55, 0x66, 0x23, 0xbb, 0xf1,
		0x01, 0xc3, 0xb4, 0x3a, 0xb1, 0xff, 0x03, 0x3a, 0xb1, 0x00, 0x29, 0x0b, 0xbb, 0x1e, 0x2e, 0xb3,
		0x88, 0x88, 0xf6, 0x18, 0x53, 0xc0, 0x40, 0x9d, 0x21, 0x15, 0xfb, 0x40, 0x21, 0x75, 0x06, 0x86,
		0x7c, 0x93, 0x34, 0x07, 0x59, 0x55, 0x11, 0x9b, 0x07, 0xc3, 0x2c, 0x99, 0x2e, 0x0a, 0x9c, 0x4a,
		0x60, 0x6a, 0x06, 0xb7, 0x8c, 0x95, 0x39, 0x00, 0xdb, 0xc2, 0x76, 0x45, 0x2b, 0x63, 0xdd, 0xcc,
		0x26, 0x76, 0xf0, 0xd2, 0x0a, 0x51, 0xe9, 0xf0, 0x92, 0xcd, 0xa4, 0xba, 0xa9, 0x1c, 0x6f, 0x86,
		0xda, 0xc0, 0x0e, 0x91, 0xb2, 0xc4, 0x0e, 0x59, 0x47, 0xb4, 0xad, 0x43, 0xc6, 0xc1, 0x24, 0xee,
		0x71, 0x99, 0xaf, 0x2c, 0x49, 0x8d, 0x98, 0x0e, 0x5d, 0x99, 0xca, 0x61, 0x6c, 0x61, 0x83, 0x4e,
		0x70, 0xa8, 0xdc, 0x04, 0xbe, 0x40, 0xa3, 0x61, 0x05, 0x34, 0x0b, 0xa5, 0x85, 0x70, 0x19, 0xd5,
		0xf0, 0xf8, 0x31, 0xc8, 0xb4, 0xba, 0x47, 0x19, 0x85, 0xb8, 0xeb, 0x21, 0xc7, 0xa3, 0x51, 0x18,
		0x57, 0xd9, 0x40, 0x91, 0x21, 0x8a, 0xad, 0x32, 0xcd, 0x72, 0x71, 0x95, 0xfc, 0x1c, 0xbf, 0x07,
		0x06, 0x5b, 0x1e, 0xdf, 0x2b, 0x70, 0xf2, 0xb1, 0x7e, 0x18, 0xed, 0x16, 0x73, 0x5d, 0xc3, 0x7f,
		0x0c, 0xfa, 0xad, 0x46, 0x6d, 0x03, 0x3b, 0xd9, 0x28, 0x65, 0xe0, 0x23, 0xa5, 0x00, 0x71, 0x13,
		0x6d, 0x60, 0x33, 0x1b, 0x9b, 0x90, 0xa6, 0x32, 0x87, 0xee, 0xe8, 0x29, 0xaa, 0xa7, 0x17, 0x09,
		0x44, 0x65, 0x48, 0xe5, 0x3e, 0x88, 0xf1, 0x14, 0x47, 0x18, 0x0e, 0xf4, 0xc6, 0x40, 0x62, 0x51,
		0xa5, 0x38, 0xe5, 0x7a, 0x48, 0x92, 0xbf, 0xcc, 0xb7, 0xfd, 0xd4, 0xe6, 0x04, 0x11, 0x10, 0xbf,
		0x2a, 0xe3, 0x90, 0xa0, 0x61, 0x56, 0xc6, 0xa2, 0x34, 0xf8, 0x63, 0xb2, 0x31, 0x65, 0x5c, 0x41,
		0x0d, 0xd3, 0xd3, 0xce, 0x23, 0xb3, 0x81, 0x69, 0xc0, 0x24, 0xd5, 0x34, 0x17, 0x3e, 0x40, 0x64,
		0xca, 0x3e, 0x48, 0xb1, 0xa8, 0x34, 0xac, 0x32, 0xbe, 0x48, 0xb3, 0x4f, 0x5c, 0x65, 0x81, 0xba,
		0x40, 0x24, 0xe4, 0xf1, 0x67, 0x5d, 0xdb, 0x12, 0x5b, 0x4b, 0x1f, 0x41, 0x04, 0xf4, 0xf1, 0xf7,
		0xb4, 0x27, 0xbe, 0x1b, 0xbb, 0x2f, 0xaf, 0x3d, 0x16, 0x27, 0x7f, 0x11, 0x81, 0x18, 0x3d, 0x6f,
		0x43, 0x90, 0x5a, 0x7b, 0x68, 0xb5, 0xa8, 0xcd, 0xad, 0xac, 0xcf, 0x2c, 0x16, 0x65, 0x49, 0xc9,
		0x00, 0x50, 0xc1, 0xa9, 0xc5, 0x95, 0xc2, 0x9a, 0x1c, 0xf1, 0xc7, 0x0b, 0xcb, 0x6b, 0x47, 0x0f,
		0xcb, 0x51, 0x1f, 0xb0, 0xce, 0x04, 0xb1, 0xa0, 0xc2, 0xdd, 0x87, 0xe4, 0xb8, 0x22, 0x43, 0x9a,
		0x11, 0x2c, 0x9c, 0x29, 0xce, 0x1d, 0x3d, 0x2c, 0xf7, 0xb7, 0x4a, 0xee, 0x3e, 0x24, 0x0f, 0x28,
		0x83, 0x90, 0xa4, 0x92, 0x99, 0x95, 0x95, 0x45, 0x39, 0xe1, 0x73, 0x96, 0xd6, 0xd4, 0x85, 0xe5,
		0x79, 0x39, 0xe9, 0x73, 0xce, 0xab, 0x2b, 0xeb, 0xab, 0x32, 0xf8, 0x0c, 0x4b, 0xc5, 0x52, 0xa9,
		0x30, 0x5f, 0x94, 0x53, 0xbe, 0xc6, 0xcc, 0x43, 0x6b, 0xc5, 0x92, 0x9c, 0x6e, 0x31, 0xeb, 0xee,
		0x43, 0xf2, 0xa0, 0xff, 0x88, 0xe2, 0xf2, 0xfa, 0x92, 0x9c, 0x51, 0x86, 0x61, 0x90, 0x3d, 0x42,
		0x18, 0x31, 0xd4, 0x26, 0x3a, 0x7a, 0x58, 0x96, 0x9b, 0x86, 0x30, 0x96, 0xe1, 0x16, 0xc1, 0xd1,
		0xc3, 0xb2, 0x32, 0x39, 0x0b, 0x71, 0x1a, 0x5d, 0x8a, 0x02, 0x99, 0xc5, 0xc2, 0x4c, 0x71, 0x51,
		0x5b, 0x59, 0x5d, 0x5b, 0x58, 0x59, 0x2e, 0x2c, 0xca, 0x52, 0x53, 0xa6, 0x16, 0x3f, 0xb9, 0xbe,
		0xa0, 0x16, 0xe7, 0xe4, 0x48, 0x50, 0xb6, 0x5a, 0x2c, 0xac, 0x15, 0xe7, 0xe4, 0xe8, 0xe4, 0x01,
		0x18, 0xed, 0x96, 0x67, 0xba, 0x9d, 0x8c, 0xc9, 0xe7, 0x24, 0x18, 0xe9, 0x92, 0x32, 0xbb, 0x9e,
		0xa2, 0x93, 0x10, 0x67, 0x91, 0xc6, 0x8a, 0xc8, 0xed, 0x5d, 0x73, 0x2f, 0x8d, 0xbb, 0x8e, 0x42,
		0x42, 0x71, 0xc1, 0x42, 0x1a, 0xdd, 0xa1, 0x90, 0x12, 0x8a, 0x8e, 0x70, 0x7a, 0x44, 0x82, 0xec,
		0x4e, 0xdc, 0x21, 0xe7, 0x3d, 0xd2, 0x72, 0xde, 0x4f, 0xb4, 0x1b, 0xb0, 0x7f, 0xe7, 0x35, 0x74,
		0x58, 0xf1, 0xbc, 0x04, 0x63, 0xdd, 0xfb, 0x8d, 0xae, 0x36, 0xdc, 0x07, 0xfd, 0x35, 0xec, 0x6d,
		0xda, 0xa2, 0xe6, 0xde, 0xda, 0x25, 0x93, 0x93, 0xe9, 0x76, 0x5f, 0x71, 0x54, 0xb0, 0x14, 0x44,
		0x77, 0x6a, 0x1a, 0x98, 0x35, 0x1d, 0x96, 0x3e, 0x1a, 0x81, 0xeb, 0xba, 0x92, 0x77, 0x35, 0xf4,
		0x46, 0x00, 0xc3, 0xaa, 0x37, 0x3c, 0x56, 0x57, 0x59, 0x9a, 0x49, 0x52, 0x09, 0x3d, 0xc2, 0x24,
		0x85, 0x34, 0x3c, 0x7f, 0x3e, 0x4a, 0xe7, 0x81, 0x89, 0xa8, 0xc2, 0xb1, 0xa6, 0xa1, 0x31, 0x6a,
		0x68, 0x6e, 0x87, 0x95, 0x76, 0x94, 0xac, 0xbb, 0x40, 0xd6, 0x4d, 0x03, 0x5b, 0x9e, 0xe6, 0x7a,
		0x0e, 0x46, 0x35, 0xc3, 0xaa, 0xd2, 0x3c, 0x9a, 0xc8, 0xc7, 0x2b, 0xc8, 0x74, 0xb1, 0x3a, 0xc4,
		0xa6, 0x4b, 0x62, 0x96, 0x20, 0x68, 0xb1, 0x70, 0x02, 0x88, 0xfe, 0x16, 0x04, 0x9b, 0xf6, 0x11,
		0x93, 0x7f, 0x18, 0x80, 0x54, 0xa0, 0x3b, 0x53, 0xf6, 0x43, 0xfa, 0x2c, 0x3a, 0x8f, 0x34, 0xd1,
		0x71, 0x33, 0x4f, 0xa4, 0x88, 0x6c, 0x95, 0x77, 0xdd, 0x77, 0xc1, 0x28, 0x55, 0xb1, 0x1b, 0x1e,
		0x76, 0x34, 0xdd, 0x44, 0xae, 0x4b, 0x9d, 0x96, 0xa0, 0xaa, 0x0a, 0x99, 0x5b, 0x21, 0x53, 0xb3,
		0x62, 0x46, 0x39, 0x02, 0x23, 0x14, 0x51, 0x6b, 0x98, 0x9e, 0x51, 0x37, 0xb1, 0x46, 0xde, 0x01,
		0x5c, 0x9a, 0x4f, 0x7d, 0xcb, 0x86, 0x89, 0xc6, 0x12, 0x57, 0x20, 0x16, 0xb9, 0xca, 0x3c, 0xdc,
		0x48, 0x61, 0x55, 0x6c, 0x61, 0x07, 0x79, 0x58, 0xc3, 0x9f, 0x6e, 0x20, 0xd3, 0xd5, 0x90, 0x55,
		0xd6, 0x36, 0x91, 0xbb, 0x99, 0x1d, 0x0d, 0x12, 0xec, 0x25, 0xba, 0xf3, 0x5c, 0xb5, 0x48, 0x35,
		0x0b, 0x56, 0xf9, 0x7e, 0xe4, 0x6e, 0x2a, 0x79, 0x18, 0xa3, 0x44, 0xae, 0xe7, 0x18, 0x56, 0x55,
		0xd3, 0x37, 0xb1, 0x7e, 0x4e, 0x6b, 0x78, 0x95, 0x63, 0xd9, 0xeb, 0x83, 0x0c, 0xd4, 0xc8, 0x12,
		0xd5, 0x99, 0x25, 0x2a, 0xeb, 0x5e, 0xe5, 0x98, 0x52, 0x82, 0x34, 0xd9, 0x8f, 0x9a, 0x71, 0x09,
		0x6b, 0x15, 0xdb, 0xa1, 0x35, 0x22, 0xd3, 0xe5, 0x70, 0x07, 0x9c, 0x38, 0xbd, 0xc2, 0x01, 0x4b,
		0x76, 0x19, 0xe7, 0xe3, 0xa5, 0xd5, 0x62, 0x71, 0x4e, 0x4d, 0x09, 0x96, 0x53, 0xb6, 0x43, 0x62,
		0xaa, 0x6a, 0xfb, 0x3e, 0x4e, 0xb1, 0x98, 0xaa, 0xda, 0xc2, 0xc3, 0x47, 0x60, 0x44, 0xd7, 0xd9,
		0xb2, 0x0d, 0x5d, 0xe3, 0xcd, 0xba, 0x9b, 0x95, 0x5b, 0xfc, 0xa5, 0xeb, 0xf3, 0x4c, 0x81, 0x87,
		0xb9, 0xab, 0x1c, 0x87, 0xeb, 0x9a, 0xfe, 0x0a, 0x02, 0x87, 0x3b, 0x56, 0xd9, 0x0e, 0x3d, 0x02,
		0x23, 0xf5, 0xad, 0x4e, 0xa0, 0xd2, 0xf2, 0xc4, 0xfa, 0x56, 0x3b, 0xec, 0x16, 0xfa, 0x02, 0xe6,
		0x60, 0x1d, 0x79, 0xb8, 0x9c, 0xdd, 0x13, 0xd4, 0x0e, 0x4c, 0x28, 0x07, 0x41, 0xd6, 0x75, 0x0d,
		0x5b, 0x68, 0xc3, 0xc4, 0x1a, 0x72, 0xb0, 0x85, 0xdc, 0xec, 0xbe, 0xa0, 0x72, 0x46, 0xd7, 0x8b,
		0x74, 0xb6, 0x40, 0x27, 0x95, 0x03, 0x30, 0x6c, 0x6f, 0x9c, 0xd5, 0x59, 0x70, 0x69, 0x75, 0x07,
		0x57, 0x8c, 0x8b, 0xd9, 0x9b, 0xa9, 0x9b, 0x86, 0xc8, 0x04, 0x0d, 0xad, 0x55, 0x2a, 0x56, 0x6e,
		0x07, 0x59, 0x77, 0x37, 0x91, 0x53, 0xa7, 0x45, 0xda, 0xad, 0x23, 0x1d, 0x67, 0x6f, 0x61, 0xaa,
		0x4c, 0xbe, 0x2c, 0xc4, 0x4a, 0x11, 0xf6, 0x91, 0xc5, 0x5b, 0xc8, 0xb2, 0xb5, 0x86, 0x8b, 0xb5,
		0xa6, 0x89, 0xfe, 0x5e, 0xdc, 0x4a, 0xcc, 0x52, 0x6f, 0x10, 0x6a, 0xeb, 0x2e, 0x9e, 0xf3, 0x95,
		0xc4, 0xf6, 0x9c, 0x81, 0xd1, 0x86, 0x65, 0x58, 0x1e, 0x76, 0xea, 0x0e, 0x26, 0x60, 0x76, 0x60,
		0xb3, 0x7f, 0x19, 0xd8, 0xa1, 0xe9, 0x5e, 0x0f, 0x6a, 0xb3, 0x20, 0x51, 0x47, 0x1a, 0x9d, 0xc2,
		0xc9, 0x3c, 0xa4, 0x83, 0xb1, 0xa3, 0x24, 0x81, 0x45, 0x8f, 0x2c, 0x91, 0x8a, 0x3a, 0xbb, 0x32,
		0x47, 0x6a, 0xe1, 0xc3, 0x45, 0x39, 0x42, 0x6a, 0xf2, 0xe2, 0xc2, 0x5a, 0x51, 0x53, 0xd7, 0x97,
		0xd7, 0x16, 0x96, 0x8a, 0x72, 0xf4, 0x40, 0x32, 0xf1, 0xf6, 0x80, 0x7c, 0xf9, 0xf2, 0xe5, 0xcb,
		0x91, 0xc9, 0xd7, 0x22, 0x90, 0x69, 0xed, 0x83, 0x95, 0x4f, 0xc0, 0x1e, 0xf1, 0xd2, 0xea, 0x62,
		0x4f, 0xbb, 0x60, 0x38, 0x34, 0x9c, 0x6b, 0x88, 0x75, 0x92, 0xfe, 0x4e, 0x8c, 0x72, 0xad, 0x12,
		0xf6, 0x1e, 0x34, 0x1c, 0x12, 0xac, 0x35, 0xe4, 0x29, 0x8b, 0xb0, 0xcf, 0xb2, 0x35, 0xd7, 0x43,
		0x56, 0x19, 0x39, 0x65, 0xad, 0x79, 0x5d, 0xa0, 0x21, 0x5d, 0xc7, 0xae, 0x6b, 0xb3, 0x4a, 0xe2,
		0xb3, 0xdc, 0x60, 0xd9, 0x25, 0xae, 0xdc, 0x4c, 0xb1, 0x05, 0xae, 0xda, 0x16, 0x35, 0xd1, 0x9d,
		0xa2, 0xe6, 0x7a, 0x48, 0xd6, 0x50, 0x5d, 0xc3, 0x96, 0xe7, 0x6c, 0xd1, 0xee, 0x2d, 0xa1, 0x26,
		0x6a, 0xa8, 0x5e, 0x24, 0xe3, 0x0f, 0x6f, 0x0f, 0x82, 0x7e, 0xfc, 0x53, 0x14, 0xd2, 0xc1, 0x0e,
		0x8e, 0x34, 0xc4, 0x3a, 0x4d, 0xf3, 0x12, 0xcd, 0x02, 0x37, 0x5d, 0xb3, 0xdf, 0x9b, 0x9e, 0x25,
		0xf9, 0x3f, 0xdf, 0xcf, 0xfa, 0x2a, 0x95, 0x21, 0x49, 0xed, 0x25, 0xb1, 0x86, 0x59, 0xb7, 0x9e,
		0x50, 0xf9, 0x48, 0x99, 0x87, 0xfe, 0xb3, 0x2e, 0xe5, 0xee, 0xa7, 0xdc, 0x37, 0x5f, 0x9b, 0xfb,
		0x74, 0x89, 0x92, 0x27, 0x4f, 0x97, 0xb4, 0xe5, 0x15, 0x75, 0xa9, 0xb0, 0xa8, 0x72, 0xb8, 0xb2,
		0x17, 0x62, 0x26, 0xba, 0xb4, 0xd5, 0x5a, 0x29, 0xa8, 0xa8, 0x57, 0xc7, 0xef, 0x85, 0xd8, 0x05,
		0x8c, 0xce, 0xb5, 0xe6, 0x67, 0x2a, 0xfa, 0x10, 0x43, 0xff, 0x20, 0xc4, 0xa9, 0xbf, 0x14, 0x00,
		0xee, 0x31, 0xb9, 0x4f, 0x49, 0x40, 0x6c, 0x76, 0x45, 0x25, 0xe1, 0x2f, 0x43, 0x9a, 0x49, 0xb5,
		0xd5, 0x85, 0xe2, 0x6c, 0x51, 0x8e, 0x4c, 0x1e, 0x81, 0x7e, 0xe6, 0x04, 0x72, 0x34, 0x7c, 0x37,
		0xc8, 0x7d, 0x7c, 0xc8, 0x39, 0x24, 0x31, 0xbb, 0xbe, 0x34, 0x53, 0x54, 0xe5, 0x48, 0x70, 0x7b,
		0x7f, 0x25, 0x41, 0x2a, 0xd0, 0x50, 0x91, 0x52, 0x8e, 0x4c, 0xd3, 0xbe, 0xa0, 0x21, 0xd3, 0x40,
		0x2e, 0xdf, 0x1f, 0xa0, 0xa2, 0x02, 0x91, 0xf4, 0xea, 0xbf, 0xff, 0x49, 0x6c, 0x3e, 0x23, 0x81,
		0xdc, 0xde, 0x8c, 0xb5, 0x19, 0x28, 0x7d, 0xa4, 0x06, 0x3e, 0x25, 0x41, 0xa6, 0xb5, 0x03, 0x6b,
		0x33, 0x6f, 0xff, 0x47, 0x6a, 0xde, 0x93, 0x12, 0x0c, 0xb6, 0xf4, 0x5d, 0xff, 0x57, 0xd6, 0x3d,
		0x11, 0x85, 0x91, 0x2e, 0x38, 0xa5, 0xc0, 0x1b, 0x54, 0xd6, 0x33, 0xdf, 0xd9, 0xcb, 0xb3, 0xa6,
		0x49, 0xfd, 0x5b, 0x45, 0x8e, 0xc7, 0xfb, 0xd9, 0xdb, 0x41, 0x36, 0xca, 0xd8, 0xf2, 0x8c, 0x8a,
		0x81, 0x1d, 0xfe, 0x6e, 0xcc, 0xba, 0xd6, 0xa1, 0xa6, 0x9c, 0xbd, 0x1e, 0x7f, 0x0c, 0x94, 0xba,
		0xed, 0x1a, 0x9e, 0x71, 0x1e, 0x6b, 0x86, 0x25, 0x5e, 0xa4, 0x49, 0x17, 0x1b, 0x53, 0x65, 0x31,
		0xb3, 0x60, 0x79, 0xbe, 0xb6, 0x85, 0xab, 0xa8, 0x4d, 0x9b, 0xa4, 0xa1, 0xa8, 0x2a, 0x8b, 0x19,
		0x5f, 0x7b, 0x3f, 0xa4, 0xcb, 0x76, 0x83, 0x34, 0x04, 0x4c, 0x8f, 0x64, 0x3d, 0x49, 0x4d, 0x31,
		0x99, 0xaf, 0xc2, 0x3b, 0xb6, 0xe6, 0x1b, 0x7c, 0x5a, 0x4d, 0x31, 0x19, 0x53, 0xb9, 0x0d, 0x86,
		0x50, 0xb5, 0xea, 0x10, 0x72, 0x41, 0xc4, 0xda, 0xd0, 0x8c, 0x2f, 0xa6, 0x8a, 0xe3, 0xa7, 0x21,
		0x21, 0xfc, 0x40, 0x0a, 0x0b, 0xf1, 0x84, 0x56, 0x67, 0xf7, 0x28, 0x11, 0xf2, 0x52, 0x6f, 0x89,
		0xc9, 0xfd, 0x90, 0x36, 0x5c, 0xad, 0x79, 0xa1, 0x17, 0x99, 0x88, 0x4c, 0x25, 0xd4, 0x94, 0xe1,
		0xfa, 0x37, 0x38, 0x93, 0xcf, 0x47, 0x20, 0xd3, 0x7a, 0x21, 0xa9, 0xcc, 0x41, 0xc2, 0xb4, 0x75,
		0x44, 0x03, 0x81, 0xdd, 0x86, 0x4f, 0x85, 0xdc, 0x61, 0x4e, 0x2f, 0x72, 0x7d, 0xd5, 0x47, 0x8e,
		0xff, 0x4e, 0x82, 0x84, 0x10, 0x2b, 0x63, 0x10, 0xab, 0x23, 0x6f, 0x93, 0xd2, 0xc5, 0x67, 0x22,
		0xb2, 0xa4, 0xd2, 0x31, 0x91, 0xbb, 0x75, 0x64, 0xd1, 0x10, 0xe0, 0x72, 0x32, 0x26, 0xfb, 0x6a,
		0x62, 0x54, 0xa6, 0x0d, 0xae, 0x5d, 0xab, 0x61, 0xcb, 0x73, 0xc5, 0xbe, 0x72, 0xf9, 0x2c, 0x17,
		0x2b, 0x77, 0xc0, 0xb0, 0xe7, 0x20, 0xc3, 0x6c, 0xd1, 0x8d, 0x51, 0x5d, 0x59, 0x4c, 0xf8, 0xca,
		0x79, 0xd8, 0x2b, 0x78, 0xcb, 0xd8, 0x43, 0xfa, 0x26, 0x2e, 0x37, 0x41, 0xfd, 0xf4, 0xb6, 0x6b,
		0x0f, 0x57, 0x98, 0xe3, 0xf3, 0x02, 0x3b, 0x73, 0x06, 0x46, 0x74, 0xbb, 0xd6, 0xee, 0x89, 0x19,
		0xb9, 0xed, 0xbd, 0xcb, 0xbd, 0x5f, 0x7a, 0x18, 0x9a, 0x4d, 0xc5, 0x73, 0x91, 0xe8, 0xfc, 0xea,
		0xcc, 0x8b, 0x91, 0xf1, 0x79, 0x86, 0x5b, 0x15, 0x1e, 0x54, 0x71, 0xc5, 0xc4, 0x3a, 0xf1, 0x0e,
		0x3c, 0x7b, 0x13, 0xdc, 0x59, 0x35, 0xbc, 0xcd, 0xc6, 0xc6, 0xb4, 0x6e, 0xd7, 0x0e, 0x56, 0xed,
		0xaa, 0xdd, 0xfc, 0x9c, 0x41, 0x46, 0x74, 0x40, 0x7f, 0xf1, 0x4f, 0x1a, 0x49, 0x5f, 0x3a, 0x1e,
		0xfa, 0xfd, 0x23, 0xbf, 0x0c, 0x23, 0x5c, 0x59, 0xa3, 0x77, 0xaa, 0xac, 0x05, 0x55, 0xae, 0xf9,
		0x42, 0x9e, 0x7d, 0xe5, 0x2d, 0x5a, 0x12, 0xd4, 0x61, 0x0e, 0x25, 0x73, 0xac, 0x49, 0xcd, 0xab,
		0x70, 0x5d, 0x0b, 0x1f, 0x8b, 0x61, 0xec, 0x84, 0x30, 0xbe, 0xc6, 0x19, 0x47, 0x02, 0x8c, 0x25,
		0x0e, 0xcd, 0xcf, 0xc2, 0xe0, 0x6e, 0xb8, 0x7e, 0xc3, 0xb9, 0xd2, 0x38, 0x48, 0x32, 0x0f, 0x43,
		0x94, 0x44, 0x6f, 0xb8, 0x9e, 0x5d, 0xa3, 0x09, 0xe2, 0xda, 0x34, 0xbf, 0x7d, 0x8b, 0x05, 0x55,
		0x86, 0xc0, 0x66, 0x7d, 0x54, 0xfe, 0x01, 0x18, 0x25, 0x12, 0x7a, 0x06, 0x83, 0x6c, 0xe1, 0x57,
		0x08, 0xd9, 0xdf, 0x3f, 0xc2, 0x62, 0x6f, 0xc4, 0x27, 0x08, 0xf0, 0x06, 0x76, 0xa2, 0x8a, 0x3d,
		0x0f, 0x3b, 0xae, 0x86, 0x4c, 0x53, 0xb9, 0xe6, 0x37, 0x86, 0xec, 0xe3, 0xef, 0xb4, 0xee, 0xc4,
		0x3c, 0x43, 0x16, 0x4c, 0x33, 0xbf, 0x0e, 0x7b, 0xba, 0xec, 0x6c, 0x0f, 0x9c, 0x4f, 0x70, 0xce,
		0xd1, 0x8e, 0xdd, 0x25, 0xb4, 0xab, 0x20, 0xe4, 0xfe, 0x7e, 0xf4, 0xc0, 0xf9, 0x24, 0xe7, 0x54,
		0x38, 0x56, 0x6c, 0x0b, 0x61, 0x3c, 0x0d, 0xc3, 0xe7, 0xb1, 0xb3, 0x61, 0xbb, 0xfc, 0xbd, 0xb7,
		0x07, 0xba, 0xa7, 0x38, 0xdd, 0x10, 0x07, 0xd2, 0xb7, 0x60, 0xc2, 0x75, 0x1c, 0x12, 0x15, 0xa4,
		0xe3, 0x1e, 0x28, 0x9e, 0xe6, 0x14, 0x03, 0x44, 0x9f, 0x40, 0x0b, 0x90, 0xae, 0xda, 0x3c, 0x0d,
		0x87, 0xc3, 0x9f, 0xe1, 0xf0, 0x94, 0xc0, 0x70, 0x8a, 0xba, 0x5d, 0x6f, 0x98, 0x24, 0x47, 0x87,
		0x53, 0x7c, 0x5b, 0x50, 0x08, 0x0c, 0xa7, 0xd8, 0x85, 0x5b, 0xbf, 0x23, 0x28, 0xdc, 0x80, 0x3f,
		0x4f, 0x42, 0xca, 0xb6, 0xcc, 0x2d, 0xdb, 0xea, 0xc5, 0x88, 0x67, 0x39, 0x03, 0x70, 0x08, 0x21,
		0x38, 0x01, 0xc9, 0x5e, 0x37, 0xe2, 0xbb, 0x1c, 0x9e, 0xc0, 0x62, 0x07, 0xe6, 0x61, 0x48, 0x24,
		0x19, 0xc3, 0xb6, 0x7a, 0xa0, 0xf8, 0x1e, 0xa7, 0xc8, 0x04, 0x60, 0x7c, 0x19, 0x1e, 0x76, 0xbd,
		0x2a, 0xee, 0x85, 0xe4, 0x79, 0xb1, 0x0c, 0x0e, 0xe1, 0xae, 0xdc, 0xc0, 0x96, 0xbe, 0xd9, 0x1b,
		0xc3, 0x0b, 0xc2, 0x95, 0x02, 0x43, 0x28, 0x66, 0x61, 0xb0, 0x86, 0x1c, 0x77, 0x13, 0x99, 0x3d,
		0x6d, 0xc7, 0xf7, 0x39, 0x47, 0xda, 0x07, 0x71, 0x8f, 0x34, 0xac, 0xdd, 0xd0, 0xbc, 0x28, 0x3c,
		0x12, 0x80, 0xf1, 0xa3, 0xe7, 0x7a, 0xf4, 0x6a, 0x61, 0x37, 0x6c, 0x3f, 0x10, 0x47, 0x8f, 0x61,
		0x97, 0x82, 0x8c, 0x27, 0x20, 0xe9, 0x1a, 0x97, 0x7a, 0xa2, 0xf9, 0xa1, 0xd8, 0x69, 0x0a, 0x20,
		0xe0, 0x87, 0x60, 0x6f, 0xd7, 0x54, 0xdf, 0x03, 0xd9, 0x4b, 0x9c, 0x6c, 0xac, 0x4b, 0xba, 0xe7,
		0x29, 0x61, 0xb7, 0x94, 0x3f, 0x12, 0x29, 0x01, 0xb7, 0x71, 0xad, 0x92, 0x36, 0xd6, 0x45, 0x95,
		0xdd, 0x79, 0xed, 0xc7, 0xc2, 0x6b, 0x0c, 0xdb, 0xe2, 0xb5, 0x35, 0x18, 0xe3, 0x8c, 0xbb, 0xdb,
		0xd7, 0x97, 0x45, 0x62, 0x65, 0xe8, 0xf5, 0xd6, 0xdd, 0xfd, 0x14, 0x8c, 0xfb, 0xee, 0x14, 0x1d,
		0x98, 0xab, 0xd5, 0x50, 0xbd, 0x07, 0xe6, 0x57, 0x38, 0xb3, 0xc8, 0xf8, 0x7e, 0x0b, 0xe7, 0x2e,
		0xa1, 0x3a, 0x21, 0x3f, 0x03, 0x59, 0x41, 0xde, 0xb0, 0x1c, 0xac, 0xdb, 0x55, 0xcb, 0xb8, 0x84,
		0xcb, 0x3d, 0x50, 0xff, 0xa4, 0x6d, 0xab, 0xd6, 0x03, 0x70, 0xc2, 0xbc, 0x00, 0xb2, 0xdf, 0x6f,
		0x68, 0x46, 0xad, 0x6e, 0x3b, 0x5e, 0x08, 0xe3, 0x4f, 0xc5, 0x4e, 0xf9, 0xb8, 0x05, 0x0a, 0xcb,
		0x17, 0x21, 0x43, 0x87, 0xbd, 0x86, 0xe4, 0xcf, 0x38, 0xd1, 0x60, 0x13, 0xc5, 0x13, 0x87, 0x6e,
		0xd7, 0xea, 0xc8, 0xe9, 0x25, 0xff, 0xfd, 0x5c, 0x24, 0x0e, 0x0e, 0x61, 0xd1, 0x37, 0xd4, 0x56,
		0x89, 0x95, 0xb0, 0xcf, 0xaf, 0xd9, 0xcf, 0x5c, 0xe5, 0x67, 0xb6, 0xb5, 0x10, 0xe7, 0x17, 0x89,
		0x7b, 0x5a, 0xcb, 0x65, 0x38, 0xd9, 0x23, 0x57, 0x7d, 0x0f, 0xb5, 0x54, 0xcb, 0xfc, 0x29, 0x18,
		0x6c, 0x29, 0x95, 0xe1, 0x54, 0x9f, 0xe5, 0x54, 0xe9, 0x60, 0xa5, 0xcc, 0x1f, 0x81, 0x18, 0x29,
		0x7b, 0xe1, 0xf0, 0xcf, 0x71, 0x38, 0x55, 0xcf, 0xdf, 0x0b, 0x09, 0x51, 0xee, 0xc2, 0xa1, 0x9f,
		0xe7, 0x50, 0x1f, 0x42, 0xe0, 0xa2, 0xd4, 0x85, 0xc3, 0xbf, 0x20, 0xe0, 0x02, 0x42, 0xe0, 0xbd,
		0xbb, 0xf0, 0xd5, 0x2f, 0xc6, 0x78, 0xba, 0x12, 0xbe, 0x3b, 0x01, 0x03, 0xbc, 0xc6, 0x85, 0xa3,
		0x1f, 0xe5, 0x0f, 0x17, 0x88, 0xfc, 0x3d, 0x10, 0xef, 0xd1, 0xe1, 0x5f, 0xe2, 0x50, 0xa6, 0x9f,
		0x9f, 0x85, 0x54, 0xa0, 0xae, 0x85, 0xc3, 0xbf, 0xcc, 0xe1, 0x41, 0x14, 0x31, 0x9d, 0xd7, 0xb5,
		0x70, 0x82, 0xaf, 0x08, 0xd3, 0x39, 0x82, 0xb8, 0x4d, 0x94, 0xb4, 0x70, 0xf4, 0x57, 0x85, 0xd7,
		0x05, 0x24, 0x7f, 0x12, 0x92, 0x7e, 0x9a, 0x0a, 0xc7, 0x7f, 0x8d, 0xe3, 0x9b, 0x18, 0xe2, 0x81,
		0x40, 0x9a, 0x0c, 0xa7, 0xf8, 0xba, 0xf0, 0x40, 0x00, 0x45, 0x8e, 0x51, 0x7b, 0xe9, 0x0b, 0x67,
		0xfa, 0x86, 0x38, 0x46, 0x6d, 0x95, 0x8f, 0xec, 0x26, 0xcd, 0x16, 0xe1, 0x14, 0xdf, 0x14, 0xbb,
		0x49, 0xf5, 0x89, 0x19, 0xed, 0xb5, 0x24, 0x9c, 0xe3, 0x5b, 0xc2, 0x8c, 0xb6, 0x52, 0x92, 0x5f,
		0x05, 0xa5, 0xb3, 0x8e, 0x84, 0xf3, 0x3d, 0xc6, 0xf9, 0x86, 0x3b, 0xca, 0x48, 0xfe, 0x41, 0x18,
		0xeb, 0x5e, 0x43, 0xc2, 0x59, 0x1f, 0xbf, 0xda, 0xd6, 0xf5, 0x07, 0x4b, 0x48, 0x7e, 0xad, 0xd9,
		0xf5, 0x07, 0xeb, 0x47, 0x38, 0xed, 0x13, 0x57, 0x5b, 0x5f, 0xec, 0x82, 0xe5, 0x23, 0x5f, 0x00,
		0x68, 0xa6, 0xee, 0x70, 0xae, 0xa7, 0x38, 0x57, 0x00, 0x44, 0x8e, 0x06, 0xcf, 0xdc, 0xe1, 0xf8,
		0xa7, 0xc5, 0xd1, 0xe0, 0x88, 0xfc, 0x09, 0x48, 0x58, 0x0d, 0xd3, 0x24, 0xc1, 0xa1, 0x5c, 0xfb,
		0x5f, 0x1a, 0xb2, 0x7f, 0x7d, 0x9f, 0x1f, 0x0c, 0x01, 0xc8, 0x1f, 0x81, 0x38, 0xae, 0x6d, 0xe0,
		0x72, 0x18, 0xf2, 0x6f, 0xef, 0x8b, 0x84, 0x40, 0xb4, 0xf3, 0x27, 0x01, 0xd8, 0x4b, 0x23, 0xbd,
		0xc3, 0x0e, 0xc1, 0xfe, 0xfd, 0x7d, 0xfe, 0x99, 0xb5, 0x09, 0x69, 0x12, 0xb0, 0x8f, 0xb6, 0xd7,
		0x26, 0x78, 0xa7, 0x95, 0x80, 0xbe, 0x68, 0x1e, 0x87, 0x81, 0xb3, 0xae, 0x6d, 0x79, 0xa8, 0x1a,
		0x86, 0xfe, 0x07, 0x47, 0x0b, 0x7d, 0xe2, 0xb0, 0x9a, 0xed, 0x60, 0x0f, 0x55, 0xdd, 0x30, 0xec,
		0x3f, 0x39, 0xd6, 0x07, 0x10, 0xb0, 0x8e, 0x5c, 0xaf, 0x97, 0x75, 0xff, 0x4b, 0x80, 0x05, 0x80,
		0x18, 0x4d, 0x7e, 0x9f, 0xc3, 0x5b, 0x61, 0xd8, 0x77, 0x85, 0xd1, 0x5c, 0x3f, 0x7f, 0x2f, 0x24,
		0xc9, 0x4f, 0xf6, 0xaf, 0x07, 0x21, 0xe0, 0x7f, 0x73, 0x70, 0x13, 0x31, 0xb3, 0xbf, 0xfb, 0xd5,
		0x0e, 0xcc, 0xdb, 0xf3, 0x36, 0xbb, 0xd4, 0x81, 0x97, 0x24, 0xc8, 0x54, 0x0c, 0x13, 0x4f, 0x97,
		0x6d, 0x8f, 0x5f, 0xc2, 0xa4, 0xc8, 0xb8, 0x6c, 0x7b, 0xc4, 0xe3, 0xe3, 0xbb, 0xbb, 0xc0, 0x99,
		0x1c, 0x06, 0x69, 0x49, 0x49, 0x83, 0x84, 0xf8, 0x47, 0x69, 0x09, 0xcd, 0x2c, 0xbe, 0x7e, 0x25,
		0xd7, 0xf7, 0xc6, 0x95, 0x5c, 0xdf, 0x1f, 0xaf, 0xe4, 0xfa, 0xde, 0xbc, 0x92, 0x93, 0xde, 0xbe,
		0x92, 0x93, 0xde, 0xbd, 0x92, 0x93, 0xde, 0xbb, 0x92, 0x93, 0x2e, 0x6f, 0xe7, 0xa4, 0x17, 0xb6,
		0x73, 0xd2, 0xcb, 0xdb, 0x39, 0xe9, 0x97, 0xdb, 0x39, 0xe9, 0xd5, 0xed, 0x9c, 0xf4, 0xfa, 0x76,
		0xae, 0xef, 0x8d, 0xed, 0x5c, 0xdf, 0x9b, 0xdb, 0x39, 0xe9, 0xed, 0xed, 0x5c, 0xdf, 0xbb, 0xdb,
		0x39, 0xe9, 0xbd, 0xed, 0x5c, 0xdf, 0xe5, 0x3f, 0xe7, 0xfa, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff,
		0xe7, 0x43, 0x74, 0x16, 0x1a, 0x2b, 0x00, 0x00,
	}
	r := bytes.NewReader(gzipped)
	gzipr, err := compress_gzip.NewReader(r)
	if err != nil {
		panic(err)
	}
	ungzipped, err := io_ioutil.ReadAll(gzipr)
	if err != nil {
		panic(err)
	}
	if err := github_com_gogo_protobuf_proto.Unmarshal(ungzipped, d); err != nil {
		panic(err)
	}
	return d
}
func (this *M) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*M)
	if !ok {
		that2, ok := that.(M)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *M")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *M but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *M but is not nil && this == nil")
	}
	if this.A != nil && that1.A != nil {
		if *this.A != *that1.A {
			return fmt.Errorf("A this(%v) Not Equal that(%v)", *this.A, *that1.A)
		}
	} else if this.A != nil {
		return fmt.Errorf("this.A == nil && that.A != nil")
	} else if that1.A != nil {
		return fmt.Errorf("A this(%v) Not Equal that(%v)", this.A, that1.A)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *M) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*M)
	if !ok {
		that2, ok := that.(M)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.A != nil && that1.A != nil {
		if *this.A != *that1.A {
			return false
		}
	} else if this.A != nil {
		return false
	} else if that1.A != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

type MFace interface {
	Proto() github_com_gogo_protobuf_proto.Message
	GetA() *string
}

func (this *M) Proto() github_com_gogo_protobuf_proto.Message {
	return this
}

func (this *M) TestProto() github_com_gogo_protobuf_proto.Message {
	return NewMFromFace(this)
}

func (this *M) GetA() *string {
	return this.A
}

func NewMFromFace(that MFace) *M {
	this := &M{}
	this.A = that.GetA()
	return this
}

func (this *M) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&filedotname.M{")
	if this.A != nil {
		s = append(s, "A: "+valueToGoStringFileDot(this.A, "string")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringFileDot(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringFileDot(m github_com_gogo_protobuf_proto.Message) string {
	e := github_com_gogo_protobuf_proto.GetUnsafeExtensionsMap(m)
	if e == nil {
		return "nil"
	}
	s := "proto.NewUnsafeXXX_InternalExtensions(map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "})"
	return s
}
func NewPopulatedM(r randyFileDot, easy bool) *M {
	this := &M{}
	if r.Intn(10) != 0 {
		v1 := randStringFileDot(r)
		this.A = &v1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedFileDot(r, 2)
	}
	return this
}

type randyFileDot interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneFileDot(r randyFileDot) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringFileDot(r randyFileDot) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneFileDot(r)
	}
	return string(tmps)
}
func randUnrecognizedFileDot(r randyFileDot, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldFileDot(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldFileDot(data []byte, r randyFileDot, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateFileDot(data, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		data = encodeVarintPopulateFileDot(data, uint64(v3))
	case 1:
		data = encodeVarintPopulateFileDot(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateFileDot(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateFileDot(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateFileDot(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateFileDot(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *M) Size() (n int) {
	var l int
	_ = l
	if m.A != nil {
		l = len(*m.A)
		n += 1 + l + sovFileDot(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFileDot(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFileDot(x uint64) (n int) {
	return sovFileDot(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *M) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&M{`,
		`A:` + valueToStringFileDot(this.A) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringFileDot(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}

func init() { proto.RegisterFile("file.dot.proto", fileDescriptorFileDot) }

var fileDescriptorFileDot = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x24, 0xcb, 0xaf, 0x6e, 0xc2, 0x50,
	0x1c, 0xc5, 0xf1, 0xdf, 0x91, 0xeb, 0x96, 0x25, 0xab, 0x5a, 0x26, 0x4e, 0x96, 0xa9, 0x99, 0xb5,
	0xef, 0x30, 0x0d, 0x86, 0x37, 0x68, 0xe9, 0x1f, 0x9a, 0x50, 0x2e, 0x21, 0xb7, 0xbe, 0x8f, 0x83,
	0x44, 0x22, 0x91, 0x95, 0x95, 0xc8, 0xde, 0x1f, 0xa6, 0xb2, 0xb2, 0x92, 0x70, 0x71, 0xe7, 0x93,
	0x9c, 0x6f, 0xf0, 0x5e, 0x54, 0xdb, 0x3c, 0xca, 0x8c, 0x8d, 0xf6, 0x07, 0x63, 0x4d, 0xf8, 0xfa,
	0x70, 0x66, 0xec, 0x2e, 0xa9, 0xf3, 0xaf, 0xbf, 0xb2, 0xb2, 0x9b, 0x26, 0x8d, 0xd6, 0xa6, 0x8e,
	0x4b, 0x53, 0x9a, 0xd8, 0x7f, 0xd2, 0xa6, 0xf0, 0xf2, 0xf0, 0xeb, 0xd9, 0xfe, 0x7c, 0x04, 0x58,
	0x86, 0x6f, 0x01, 0x92, 0x4f, 0x7c, 0xe3, 0xf7, 0x65, 0x85, 0xe4, 0x7f, 0xd1, 0x39, 0x4a, 0xef,
	0x28, 0x57, 0x47, 0x19, 0x1c, 0x31, 0x3a, 0x62, 0x72, 0xc4, 0xec, 0x88, 0x56, 0x89, 0xa3, 0x12,
	0x27, 0x25, 0xce, 0x4a, 0x5c, 0x94, 0xe8, 0x94, 0xd2, 0x2b, 0x65, 0x50, 0x62, 0x54, 0xca, 0xa4,
	0xc4, 0xac, 0x94, 0xf6, 0x46, 0xb9, 0x07, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x59, 0x32, 0x8a, 0xad,
	0x00, 0x00, 0x00,
}