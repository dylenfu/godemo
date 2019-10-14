// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/go-micro/api/proto/api.proto

package go_api

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

type Pair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}
func (*Pair) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_17a7876430d97ebd, []int{0}
}
func (m *Pair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pair.Unmarshal(m, b)
}
func (m *Pair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pair.Marshal(b, m, deterministic)
}
func (dst *Pair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pair.Merge(dst, src)
}
func (m *Pair) XXX_Size() int {
	return xxx_messageInfo_Pair.Size(m)
}
func (m *Pair) XXX_DiscardUnknown() {
	xxx_messageInfo_Pair.DiscardUnknown(m)
}

var xxx_messageInfo_Pair proto.InternalMessageInfo

func (m *Pair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Pair) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

// A HTTP request as RPC
// Forward by the api handler
type Request struct {
	Method               string           `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Path                 string           `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Header               map[string]*Pair `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Get                  map[string]*Pair `protobuf:"bytes,4,rep,name=get,proto3" json:"get,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Post                 map[string]*Pair `protobuf:"bytes,5,rep,name=post,proto3" json:"post,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 string           `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
	Url                  string           `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_17a7876430d97ebd, []int{1}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Request) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetGet() map[string]*Pair {
	if m != nil {
		return m.Get
	}
	return nil
}

func (m *Request) GetPost() map[string]*Pair {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *Request) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Request) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// A HTTP response as RPC
// Expected response for the api handler
type Response struct {
	StatusCode           int32            `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Header               map[string]*Pair `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 string           `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_17a7876430d97ebd, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Response) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Response) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

// A HTTP event as RPC
// Forwarded by the event handler
type Event struct {
	// e.g login
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// uuid
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// unix timestamp of event
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// event headers
	Header map[string]*Pair `protobuf:"bytes,4,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// the event data
	Data                 string   `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_17a7876430d97ebd, []int{3}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Event) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Event) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Pair)(nil), "go.api.Pair")
	proto.RegisterType((*Request)(nil), "go.api.Request")
	proto.RegisterMapType((map[string]*Pair)(nil), "go.api.Request.GetEntry")
	proto.RegisterMapType((map[string]*Pair)(nil), "go.api.Request.HeaderEntry")
	proto.RegisterMapType((map[string]*Pair)(nil), "go.api.Request.PostEntry")
	proto.RegisterType((*Response)(nil), "go.api.Response")
	proto.RegisterMapType((map[string]*Pair)(nil), "go.api.Response.HeaderEntry")
	proto.RegisterType((*Event)(nil), "go.api.Event")
	proto.RegisterMapType((map[string]*Pair)(nil), "go.api.Event.HeaderEntry")
}

func init() {
	proto.RegisterFile("github.com/micro/go-micro/api/proto/api.proto", fileDescriptor_api_17a7876430d97ebd)
}

var fileDescriptor_api_17a7876430d97ebd = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x55, 0xe2, 0x24, 0x6d, 0x66, 0x11, 0x42, 0x3e, 0x20, 0x53, 0x2a, 0xb4, 0xca, 0x85, 0x15,
	0x52, 0x13, 0x68, 0x39, 0x20, 0xae, 0xb0, 0x2a, 0xc7, 0xca, 0x7f, 0xe0, 0x6d, 0xac, 0xc4, 0x62,
	0x13, 0x9b, 0xd8, 0xa9, 0xb4, 0x1f, 0xc7, 0x81, 0xcf, 0xe0, 0x6f, 0x90, 0x27, 0xde, 0xdd, 0xb2,
	0x5a, 0x2e, 0x74, 0x6f, 0x2f, 0xf6, 0x9b, 0x37, 0x6f, 0xde, 0x38, 0xf0, 0xb6, 0x51, 0xae, 0x1d,
	0x57, 0xe5, 0xbd, 0xee, 0xaa, 0x4e, 0xdd, 0x0f, 0xba, 0x6a, 0xf4, 0x95, 0x30, 0xaa, 0x32, 0x83,
	0x76, 0xba, 0x12, 0x46, 0x95, 0x88, 0x68, 0xd6, 0xe8, 0x52, 0x18, 0x55, 0xbc, 0x87, 0xe4, 0x4e,
	0xa8, 0x81, 0xbe, 0x00, 0xf2, 0x5d, 0x6e, 0x58, 0x34, 0x8f, 0x16, 0x39, 0xf7, 0x90, 0xbe, 0x84,
	0xec, 0x41, 0xac, 0x47, 0x69, 0x59, 0x3c, 0x27, 0x8b, 0x9c, 0x87, 0xaf, 0xe2, 0x17, 0x81, 0x33,
	0x2e, 0x7f, 0x8c, 0xd2, 0x3a, 0xcf, 0xe9, 0xa4, 0x6b, 0x75, 0x1d, 0x0a, 0xc3, 0x17, 0xa5, 0x90,
	0x18, 0xe1, 0x5a, 0x16, 0xe3, 0x29, 0x62, 0x7a, 0x03, 0x59, 0x2b, 0x45, 0x2d, 0x07, 0x46, 0xe6,
	0x64, 0x31, 0xbb, 0x7e, 0x5d, 0x4e, 0x16, 0xca, 0x20, 0x56, 0x7e, 0xc3, 0xdb, 0x65, 0xef, 0x86,
	0x0d, 0x0f, 0x54, 0xfa, 0x0e, 0x48, 0x23, 0x1d, 0x4b, 0xb0, 0x82, 0x1d, 0x56, 0xdc, 0x4a, 0x37,
	0xd1, 0x3d, 0x89, 0x5e, 0x41, 0x62, 0xb4, 0x75, 0x2c, 0x45, 0xf2, 0xab, 0x43, 0xf2, 0x9d, 0xb6,
	0x81, 0x8d, 0x34, 0xef, 0x71, 0xa5, 0xeb, 0x0d, 0xcb, 0x26, 0x8f, 0x1e, 0xfb, 0x14, 0xc6, 0x61,
	0xcd, 0xce, 0xa6, 0x14, 0xc6, 0x61, 0x7d, 0x71, 0x0b, 0xb3, 0x47, 0xbe, 0x8e, 0xc4, 0x54, 0x40,
	0x8a, 0xc1, 0xe0, 0xac, 0xb3, 0xeb, 0x67, 0xdb, 0xb6, 0x3e, 0x55, 0x3e, 0x5d, 0x7d, 0x8e, 0x3f,
	0x45, 0x17, 0x5f, 0xe1, 0x7c, 0x6b, 0xf7, 0x09, 0x2a, 0x4b, 0xc8, 0x77, 0x73, 0xfc, 0xbf, 0x4c,
	0xf1, 0x33, 0x82, 0x73, 0x2e, 0xad, 0xd1, 0xbd, 0x95, 0xf4, 0x0d, 0x80, 0x75, 0xc2, 0x8d, 0xf6,
	0x8b, 0xae, 0x25, 0xaa, 0xa5, 0xfc, 0xd1, 0x09, 0xfd, 0xb8, 0x5b, 0x5c, 0x8c, 0xc9, 0x5e, 0xee,
	0x93, 0x9d, 0x14, 0x8e, 0x6e, 0x6e, 0x1b, 0x2f, 0xd9, 0xc7, 0x7b, 0xb2, 0x30, 0x8b, 0xdf, 0x11,
	0xa4, 0xcb, 0x07, 0xd9, 0xe3, 0x16, 0x7b, 0xd1, 0xc9, 0x20, 0x82, 0x98, 0x3e, 0x87, 0x58, 0xd5,
	0xe1, 0xed, 0xc5, 0xaa, 0xa6, 0x97, 0x90, 0x3b, 0xd5, 0x49, 0xeb, 0x44, 0x67, 0xd0, 0x0f, 0xe1,
	0xfb, 0x03, 0xfa, 0x61, 0x37, 0x5e, 0xf2, 0xf7, 0xc3, 0xc1, 0x06, 0xff, 0x9a, 0xad, 0x16, 0x4e,
	0xb0, 0x74, 0x6a, 0xea, 0xf1, 0xc9, 0x66, 0x5b, 0x65, 0xf8, 0x83, 0xde, 0xfc, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x7a, 0xb4, 0xd4, 0x8f, 0xcb, 0x03, 0x00, 0x00,
}
