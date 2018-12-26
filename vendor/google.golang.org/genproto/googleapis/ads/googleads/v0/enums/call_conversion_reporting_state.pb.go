// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/call_conversion_reporting_state.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

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

// Possible data types for a call conversion action state.
type CallConversionReportingStateEnum_CallConversionReportingState int32

const (
	// Not specified.
	CallConversionReportingStateEnum_UNSPECIFIED CallConversionReportingStateEnum_CallConversionReportingState = 0
	// Used for return value only. Represents value unknown in this version.
	CallConversionReportingStateEnum_UNKNOWN CallConversionReportingStateEnum_CallConversionReportingState = 1
	// Call conversion action is disabled.
	CallConversionReportingStateEnum_DISABLED CallConversionReportingStateEnum_CallConversionReportingState = 2
	// Call conversion action will use call conversion type set at the
	// account level.
	CallConversionReportingStateEnum_USE_ACCOUNT_LEVEL_CALL_CONVERSION_ACTION CallConversionReportingStateEnum_CallConversionReportingState = 3
	// Call conversion action will use call conversion type set at the resource
	// (call only ads/call extensions) level.
	CallConversionReportingStateEnum_USE_RESOURCE_LEVEL_CALL_CONVERSION_ACTION CallConversionReportingStateEnum_CallConversionReportingState = 4
)

var CallConversionReportingStateEnum_CallConversionReportingState_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "DISABLED",
	3: "USE_ACCOUNT_LEVEL_CALL_CONVERSION_ACTION",
	4: "USE_RESOURCE_LEVEL_CALL_CONVERSION_ACTION",
}
var CallConversionReportingStateEnum_CallConversionReportingState_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"DISABLED":    2,
	"USE_ACCOUNT_LEVEL_CALL_CONVERSION_ACTION":  3,
	"USE_RESOURCE_LEVEL_CALL_CONVERSION_ACTION": 4,
}

func (x CallConversionReportingStateEnum_CallConversionReportingState) String() string {
	return proto.EnumName(CallConversionReportingStateEnum_CallConversionReportingState_name, int32(x))
}
func (CallConversionReportingStateEnum_CallConversionReportingState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_call_conversion_reporting_state_7cdcaa6987dc2f7b, []int{0, 0}
}

// Container for enum describing possible data types for call conversion
// reporting state.
type CallConversionReportingStateEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallConversionReportingStateEnum) Reset()         { *m = CallConversionReportingStateEnum{} }
func (m *CallConversionReportingStateEnum) String() string { return proto.CompactTextString(m) }
func (*CallConversionReportingStateEnum) ProtoMessage()    {}
func (*CallConversionReportingStateEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_call_conversion_reporting_state_7cdcaa6987dc2f7b, []int{0}
}
func (m *CallConversionReportingStateEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallConversionReportingStateEnum.Unmarshal(m, b)
}
func (m *CallConversionReportingStateEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallConversionReportingStateEnum.Marshal(b, m, deterministic)
}
func (dst *CallConversionReportingStateEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallConversionReportingStateEnum.Merge(dst, src)
}
func (m *CallConversionReportingStateEnum) XXX_Size() int {
	return xxx_messageInfo_CallConversionReportingStateEnum.Size(m)
}
func (m *CallConversionReportingStateEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CallConversionReportingStateEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CallConversionReportingStateEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CallConversionReportingStateEnum)(nil), "google.ads.googleads.v0.enums.CallConversionReportingStateEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.CallConversionReportingStateEnum_CallConversionReportingState", CallConversionReportingStateEnum_CallConversionReportingState_name, CallConversionReportingStateEnum_CallConversionReportingState_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/call_conversion_reporting_state.proto", fileDescriptor_call_conversion_reporting_state_7cdcaa6987dc2f7b)
}

var fileDescriptor_call_conversion_reporting_state_7cdcaa6987dc2f7b = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0xff, 0xdd, 0xfe, 0xa8, 0x64, 0x82, 0xa5, 0x67, 0x07, 0x6e, 0x27, 0x05, 0x4d, 0x0b,
	0x1e, 0x3d, 0xa5, 0x59, 0x1c, 0xc5, 0x92, 0x8e, 0x76, 0xad, 0x20, 0x85, 0x50, 0xd7, 0x12, 0x06,
	0x59, 0x32, 0x9a, 0x6e, 0x1f, 0xc8, 0x8b, 0xe0, 0x47, 0x11, 0x4f, 0x7e, 0x22, 0x49, 0xeb, 0x76,
	0xb3, 0x97, 0xf0, 0xc0, 0xf3, 0xe4, 0xf7, 0xf2, 0x3e, 0x2f, 0xc0, 0x5c, 0x29, 0x2e, 0x2a, 0xb7,
	0x28, 0xb5, 0xdb, 0x49, 0xa3, 0xf6, 0x9e, 0x5b, 0xc9, 0xdd, 0x46, 0xbb, 0xab, 0x42, 0x08, 0xb6,
	0x52, 0x72, 0x5f, 0xd5, 0x7a, 0xad, 0x24, 0xab, 0xab, 0xad, 0xaa, 0x9b, 0xb5, 0xe4, 0x4c, 0x37,
	0x45, 0x53, 0xc1, 0x6d, 0xad, 0x1a, 0xe5, 0x8c, 0xbb, 0x9f, 0xb0, 0x28, 0x35, 0x3c, 0x42, 0xe0,
	0xde, 0x83, 0x2d, 0x64, 0xfa, 0x65, 0x81, 0x2b, 0x5c, 0x08, 0x81, 0x8f, 0x9c, 0xf8, 0x80, 0x49,
	0x0c, 0x85, 0xc8, 0xdd, 0x66, 0xfa, 0x6e, 0x81, 0xcb, 0xbe, 0x90, 0x73, 0x01, 0x46, 0x29, 0x4d,
	0x16, 0x04, 0x07, 0x8f, 0x01, 0x99, 0xd9, 0xff, 0x9c, 0x11, 0x38, 0x4d, 0xe9, 0x13, 0x8d, 0x9e,
	0xa9, 0x6d, 0x39, 0xe7, 0xe0, 0x6c, 0x16, 0x24, 0xc8, 0x0f, 0xc9, 0xcc, 0x1e, 0x38, 0xb7, 0xe0,
	0x3a, 0x4d, 0x08, 0x43, 0x18, 0x47, 0x29, 0x5d, 0xb2, 0x90, 0x64, 0x24, 0x64, 0x18, 0x85, 0x21,
	0xc3, 0x11, 0xcd, 0x48, 0x9c, 0x04, 0x11, 0x65, 0x08, 0x2f, 0x83, 0x88, 0xda, 0x43, 0xe7, 0x0e,
	0xdc, 0x98, 0x74, 0x4c, 0x92, 0x28, 0x8d, 0x31, 0xe9, 0x8f, 0xff, 0xf7, 0xbf, 0x2d, 0x30, 0x59,
	0xa9, 0x0d, 0xec, 0x5d, 0xda, 0x9f, 0xf4, 0x2d, 0xb3, 0x30, 0xb5, 0x2d, 0xac, 0x17, 0xff, 0x97,
	0xc1, 0x95, 0x28, 0x24, 0x87, 0xaa, 0xe6, 0x2e, 0xaf, 0x64, 0x5b, 0xea, 0xe1, 0x1a, 0xdb, 0xb5,
	0xfe, 0xe3, 0x38, 0x0f, 0xed, 0xfb, 0x36, 0x18, 0xce, 0x11, 0xfa, 0x18, 0x8c, 0xe7, 0x1d, 0x0a,
	0x95, 0x1a, 0x76, 0xd2, 0xa8, 0xcc, 0x83, 0xa6, 0x5d, 0xfd, 0x79, 0xf0, 0x73, 0x54, 0xea, 0xfc,
	0xe8, 0xe7, 0x99, 0x97, 0xb7, 0xfe, 0xeb, 0x49, 0x3b, 0xf4, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff,
	0xc1, 0xfc, 0x20, 0x51, 0x10, 0x02, 0x00, 0x00,
}
