// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grep.proto

package generated

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BookInfo struct {
	Book                 string   `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
	Word                 string   `protobuf:"bytes,2,opt,name=word,proto3" json:"word,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookInfo) Reset()         { *m = BookInfo{} }
func (m *BookInfo) String() string { return proto.CompactTextString(m) }
func (*BookInfo) ProtoMessage()    {}
func (*BookInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_502de7f9a522a36d, []int{0}
}

func (m *BookInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookInfo.Unmarshal(m, b)
}
func (m *BookInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookInfo.Marshal(b, m, deterministic)
}
func (m *BookInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookInfo.Merge(m, src)
}
func (m *BookInfo) XXX_Size() int {
	return xxx_messageInfo_BookInfo.Size(m)
}
func (m *BookInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BookInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BookInfo proto.InternalMessageInfo

func (m *BookInfo) GetBook() string {
	if m != nil {
		return m.Book
	}
	return ""
}

func (m *BookInfo) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

type Occurrences struct {
	Book                 string   `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
	Word                 string   `protobuf:"bytes,2,opt,name=word,proto3" json:"word,omitempty"`
	Total                int64    `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Occurrences) Reset()         { *m = Occurrences{} }
func (m *Occurrences) String() string { return proto.CompactTextString(m) }
func (*Occurrences) ProtoMessage()    {}
func (*Occurrences) Descriptor() ([]byte, []int) {
	return fileDescriptor_502de7f9a522a36d, []int{1}
}

func (m *Occurrences) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Occurrences.Unmarshal(m, b)
}
func (m *Occurrences) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Occurrences.Marshal(b, m, deterministic)
}
func (m *Occurrences) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Occurrences.Merge(m, src)
}
func (m *Occurrences) XXX_Size() int {
	return xxx_messageInfo_Occurrences.Size(m)
}
func (m *Occurrences) XXX_DiscardUnknown() {
	xxx_messageInfo_Occurrences.DiscardUnknown(m)
}

var xxx_messageInfo_Occurrences proto.InternalMessageInfo

func (m *Occurrences) GetBook() string {
	if m != nil {
		return m.Book
	}
	return ""
}

func (m *Occurrences) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

func (m *Occurrences) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*BookInfo)(nil), "grep.BookInfo")
	proto.RegisterType((*Occurrences)(nil), "grep.Occurrences")
}

func init() {
	proto.RegisterFile("grep.proto", fileDescriptor_502de7f9a522a36d)
}

var fileDescriptor_502de7f9a522a36d = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2f, 0x4a, 0x2d,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x8c, 0xb8, 0x38, 0x9c, 0xf2,
	0xf3, 0xb3, 0x3d, 0xf3, 0xd2, 0xf2, 0x85, 0x84, 0xb8, 0x58, 0x92, 0xf2, 0xf3, 0xb3, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x90, 0x58, 0x79, 0x7e, 0x51, 0x8a, 0x04, 0x13, 0x44,
	0x0c, 0xc4, 0x56, 0xf2, 0xe6, 0xe2, 0xf6, 0x4f, 0x4e, 0x2e, 0x2d, 0x2a, 0x4a, 0xcd, 0x4b, 0x4e,
	0x2d, 0x26, 0x56, 0x9b, 0x90, 0x08, 0x17, 0x6b, 0x49, 0x7e, 0x49, 0x62, 0x8e, 0x04, 0xb3, 0x02,
	0xa3, 0x06, 0x73, 0x10, 0x84, 0x63, 0x64, 0xc6, 0xc5, 0xee, 0x5e, 0x94, 0x5a, 0x50, 0x90, 0x5a,
	0x24, 0xa4, 0xcd, 0xc5, 0x02, 0x62, 0x0a, 0xf1, 0xe9, 0x81, 0x9d, 0x09, 0x73, 0x97, 0x94, 0x20,
	0x84, 0x8f, 0x64, 0xa7, 0x12, 0x83, 0x93, 0x48, 0x94, 0x50, 0x66, 0x5e, 0x49, 0x6a, 0x51, 0x5e,
	0x62, 0x8e, 0x7e, 0x7a, 0x6a, 0x5e, 0x6a, 0x51, 0x62, 0x49, 0x6a, 0x4a, 0x12, 0x1b, 0xd8, 0x6f,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x4d, 0x61, 0x93, 0xe9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GrepperClient is the client API for Grepper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GrepperClient interface {
	// Grep returns the number of occurrences of a given word in a book.
	Grep(ctx context.Context, in *BookInfo, opts ...grpc.CallOption) (*Occurrences, error)
}

type grepperClient struct {
	cc grpc.ClientConnInterface
}

func NewGrepperClient(cc grpc.ClientConnInterface) GrepperClient {
	return &grepperClient{cc}
}

func (c *grepperClient) Grep(ctx context.Context, in *BookInfo, opts ...grpc.CallOption) (*Occurrences, error) {
	out := new(Occurrences)
	err := c.cc.Invoke(ctx, "/grep.Grepper/Grep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrepperServer is the server API for Grepper service.
type GrepperServer interface {
	// Grep returns the number of occurrences of a given word in a book.
	Grep(context.Context, *BookInfo) (*Occurrences, error)
}

// UnimplementedGrepperServer can be embedded to have forward compatible implementations.
type UnimplementedGrepperServer struct {
}

func (*UnimplementedGrepperServer) Grep(ctx context.Context, req *BookInfo) (*Occurrences, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Grep not implemented")
}

func RegisterGrepperServer(s *grpc.Server, srv GrepperServer) {
	s.RegisterService(&_Grepper_serviceDesc, srv)
}

func _Grepper_Grep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrepperServer).Grep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grep.Grepper/Grep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrepperServer).Grep(ctx, req.(*BookInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Grepper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grep.Grepper",
	HandlerType: (*GrepperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Grep",
			Handler:    _Grepper_Grep_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grep.proto",
}