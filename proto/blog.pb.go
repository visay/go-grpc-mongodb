// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/blog.proto

/*
Package blogpb is a generated protocol buffer package.

It is generated from these files:
	proto/blog.proto

It has these top-level messages:
	Blog
	CreateBlogReq
	CreateBlogRes
	ReadBlogReq
	ReadBlogRes
	UpdateBlogReq
	UpdateBlogRes
	DeleteBlogReq
	DeleteBlogRes
*/
package blogpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Blog struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	AuthorId string `protobuf:"bytes,2,opt,name=author_id,json=authorId" json:"author_id,omitempty"`
	Title    string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Content  string `protobuf:"bytes,4,opt,name=content" json:"content,omitempty"`
}

func (m *Blog) Reset()                    { *m = Blog{} }
func (m *Blog) String() string            { return proto.CompactTextString(m) }
func (*Blog) ProtoMessage()               {}
func (*Blog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Blog) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Blog) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *Blog) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Blog) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type CreateBlogReq struct {
	Blog *Blog `protobuf:"bytes,1,opt,name=blog" json:"blog,omitempty"`
}

func (m *CreateBlogReq) Reset()                    { *m = CreateBlogReq{} }
func (m *CreateBlogReq) String() string            { return proto.CompactTextString(m) }
func (*CreateBlogReq) ProtoMessage()               {}
func (*CreateBlogReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateBlogReq) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type CreateBlogRes struct {
	Blog *Blog `protobuf:"bytes,1,opt,name=blog" json:"blog,omitempty"`
}

func (m *CreateBlogRes) Reset()                    { *m = CreateBlogRes{} }
func (m *CreateBlogRes) String() string            { return proto.CompactTextString(m) }
func (*CreateBlogRes) ProtoMessage()               {}
func (*CreateBlogRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateBlogRes) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type ReadBlogReq struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ReadBlogReq) Reset()                    { *m = ReadBlogReq{} }
func (m *ReadBlogReq) String() string            { return proto.CompactTextString(m) }
func (*ReadBlogReq) ProtoMessage()               {}
func (*ReadBlogReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadBlogReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadBlogRes struct {
	Blog *Blog `protobuf:"bytes,1,opt,name=blog" json:"blog,omitempty"`
}

func (m *ReadBlogRes) Reset()                    { *m = ReadBlogRes{} }
func (m *ReadBlogRes) String() string            { return proto.CompactTextString(m) }
func (*ReadBlogRes) ProtoMessage()               {}
func (*ReadBlogRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReadBlogRes) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type UpdateBlogReq struct {
	Blog *Blog `protobuf:"bytes,1,opt,name=blog" json:"blog,omitempty"`
}

func (m *UpdateBlogReq) Reset()                    { *m = UpdateBlogReq{} }
func (m *UpdateBlogReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateBlogReq) ProtoMessage()               {}
func (*UpdateBlogReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateBlogReq) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type UpdateBlogRes struct {
	Blog *Blog `protobuf:"bytes,1,opt,name=blog" json:"blog,omitempty"`
}

func (m *UpdateBlogRes) Reset()                    { *m = UpdateBlogRes{} }
func (m *UpdateBlogRes) String() string            { return proto.CompactTextString(m) }
func (*UpdateBlogRes) ProtoMessage()               {}
func (*UpdateBlogRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdateBlogRes) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type DeleteBlogReq struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteBlogReq) Reset()                    { *m = DeleteBlogReq{} }
func (m *DeleteBlogReq) String() string            { return proto.CompactTextString(m) }
func (*DeleteBlogReq) ProtoMessage()               {}
func (*DeleteBlogReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DeleteBlogReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteBlogRes struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *DeleteBlogRes) Reset()                    { *m = DeleteBlogRes{} }
func (m *DeleteBlogRes) String() string            { return proto.CompactTextString(m) }
func (*DeleteBlogRes) ProtoMessage()               {}
func (*DeleteBlogRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DeleteBlogRes) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*Blog)(nil), "blog.Blog")
	proto.RegisterType((*CreateBlogReq)(nil), "blog.CreateBlogReq")
	proto.RegisterType((*CreateBlogRes)(nil), "blog.CreateBlogRes")
	proto.RegisterType((*ReadBlogReq)(nil), "blog.ReadBlogReq")
	proto.RegisterType((*ReadBlogRes)(nil), "blog.ReadBlogRes")
	proto.RegisterType((*UpdateBlogReq)(nil), "blog.UpdateBlogReq")
	proto.RegisterType((*UpdateBlogRes)(nil), "blog.UpdateBlogRes")
	proto.RegisterType((*DeleteBlogReq)(nil), "blog.DeleteBlogReq")
	proto.RegisterType((*DeleteBlogRes)(nil), "blog.DeleteBlogRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BlogService service

type BlogServiceClient interface {
	CreateBlog(ctx context.Context, in *CreateBlogReq, opts ...grpc.CallOption) (*CreateBlogRes, error)
	ReadBlog(ctx context.Context, in *ReadBlogReq, opts ...grpc.CallOption) (*ReadBlogRes, error)
	UpdateBlog(ctx context.Context, in *UpdateBlogReq, opts ...grpc.CallOption) (*UpdateBlogRes, error)
	DeleteBlog(ctx context.Context, in *DeleteBlogReq, opts ...grpc.CallOption) (*DeleteBlogRes, error)
}

type blogServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlogServiceClient(cc *grpc.ClientConn) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) CreateBlog(ctx context.Context, in *CreateBlogReq, opts ...grpc.CallOption) (*CreateBlogRes, error) {
	out := new(CreateBlogRes)
	err := grpc.Invoke(ctx, "/blog.BlogService/CreateBlog", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ReadBlog(ctx context.Context, in *ReadBlogReq, opts ...grpc.CallOption) (*ReadBlogRes, error) {
	out := new(ReadBlogRes)
	err := grpc.Invoke(ctx, "/blog.BlogService/ReadBlog", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) UpdateBlog(ctx context.Context, in *UpdateBlogReq, opts ...grpc.CallOption) (*UpdateBlogRes, error) {
	out := new(UpdateBlogRes)
	err := grpc.Invoke(ctx, "/blog.BlogService/UpdateBlog", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) DeleteBlog(ctx context.Context, in *DeleteBlogReq, opts ...grpc.CallOption) (*DeleteBlogRes, error) {
	out := new(DeleteBlogRes)
	err := grpc.Invoke(ctx, "/blog.BlogService/DeleteBlog", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BlogService service

type BlogServiceServer interface {
	CreateBlog(context.Context, *CreateBlogReq) (*CreateBlogRes, error)
	ReadBlog(context.Context, *ReadBlogReq) (*ReadBlogRes, error)
	UpdateBlog(context.Context, *UpdateBlogReq) (*UpdateBlogRes, error)
	DeleteBlog(context.Context, *DeleteBlogReq) (*DeleteBlogRes, error)
}

func RegisterBlogServiceServer(s *grpc.Server, srv BlogServiceServer) {
	s.RegisterService(&_BlogService_serviceDesc, srv)
}

func _BlogService_CreateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).CreateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/CreateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).CreateBlog(ctx, req.(*CreateBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ReadBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).ReadBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/ReadBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).ReadBlog(ctx, req.(*ReadBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_UpdateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).UpdateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/UpdateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).UpdateBlog(ctx, req.(*UpdateBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_DeleteBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).DeleteBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/DeleteBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).DeleteBlog(ctx, req.(*DeleteBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blog.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlog",
			Handler:    _BlogService_CreateBlog_Handler,
		},
		{
			MethodName: "ReadBlog",
			Handler:    _BlogService_ReadBlog_Handler,
		},
		{
			MethodName: "UpdateBlog",
			Handler:    _BlogService_UpdateBlog_Handler,
		},
		{
			MethodName: "DeleteBlog",
			Handler:    _BlogService_DeleteBlog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/blog.proto",
}

func init() { proto.RegisterFile("proto/blog.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4e, 0x83, 0x40,
	0x10, 0xc6, 0x03, 0x62, 0xa5, 0x43, 0x6a, 0x74, 0xf4, 0x40, 0x30, 0xfe, 0x09, 0x27, 0x3d, 0xd8,
	0x9a, 0x9a, 0xf8, 0x00, 0xd5, 0x8b, 0x57, 0x8c, 0x17, 0x2f, 0x06, 0xd8, 0x49, 0x25, 0x21, 0x5d,
	0x64, 0xb7, 0xbe, 0xb2, 0xaf, 0x61, 0x76, 0xc8, 0x86, 0x3f, 0x92, 0x54, 0x6f, 0x3b, 0xdf, 0xcc,
	0xc7, 0x6f, 0xe6, 0x0b, 0x70, 0x54, 0xd5, 0x52, 0xcb, 0x45, 0x56, 0xca, 0xf5, 0x9c, 0x9f, 0xe8,
	0x99, 0x77, 0x9c, 0x83, 0xb7, 0x2a, 0xe5, 0x1a, 0x0f, 0xc1, 0x2d, 0x44, 0xe8, 0x5c, 0x39, 0xd7,
	0xd3, 0xc4, 0x2d, 0x04, 0x9e, 0xc1, 0x34, 0xdd, 0xea, 0x0f, 0x59, 0xbf, 0x17, 0x22, 0x74, 0x59,
	0xf6, 0x1b, 0xe1, 0x59, 0xe0, 0x29, 0xec, 0xeb, 0x42, 0x97, 0x14, 0xee, 0x71, 0xa3, 0x29, 0x30,
	0x84, 0x83, 0x5c, 0x6e, 0x34, 0x6d, 0x74, 0xe8, 0xb1, 0x6e, 0xcb, 0x78, 0x01, 0xb3, 0xc7, 0x9a,
	0x52, 0x4d, 0x06, 0x95, 0xd0, 0x27, 0x5e, 0x00, 0xd3, 0x99, 0x17, 0x2c, 0x61, 0xce, 0x6b, 0x71,
	0xb3, 0xd9, 0x6a, 0x60, 0x50, 0x3b, 0x0d, 0xe7, 0x10, 0x24, 0x94, 0x0a, 0xfb, 0xfd, 0xc1, 0x35,
	0xf1, 0x6d, 0xb7, 0xad, 0xfe, 0x82, 0x7f, 0xad, 0xc4, 0xff, 0xf6, 0xed, 0x1a, 0x76, 0x13, 0x2e,
	0x61, 0xf6, 0x44, 0x25, 0xb5, 0x84, 0xe1, 0xc6, 0x37, 0xfd, 0x01, 0x65, 0xd2, 0x55, 0xdb, 0x3c,
	0x27, 0xa5, 0x78, 0xca, 0x4f, 0x6c, 0xb9, 0xfc, 0x76, 0x20, 0x30, 0x53, 0x2f, 0x54, 0x7f, 0x15,
	0x39, 0xe1, 0x03, 0x40, 0x1b, 0x1e, 0x9e, 0x34, 0xec, 0x5e, 0xfe, 0xd1, 0x88, 0xa8, 0xf0, 0x0e,
	0x7c, 0x1b, 0x12, 0x1e, 0x37, 0x03, 0x9d, 0x4c, 0xa3, 0x5f, 0x92, 0x32, 0xa4, 0xf6, 0x6c, 0x4b,
	0xea, 0x25, 0x17, 0x8d, 0x88, 0xec, 0x6b, 0x8f, 0xb3, 0xbe, 0x5e, 0x1e, 0xd1, 0x88, 0xa8, 0x56,
	0xfe, 0xdb, 0xc4, 0xa8, 0x55, 0x96, 0x4d, 0xf8, 0x1f, 0xbe, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff,
	0xb4, 0x3c, 0x48, 0xea, 0xd7, 0x02, 0x00, 0x00,
}