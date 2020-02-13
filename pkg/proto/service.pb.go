// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package proto

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

type Command struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Args                 []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Command) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

type JobReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobReq) Reset()         { *m = JobReq{} }
func (m *JobReq) String() string { return proto.CompactTextString(m) }
func (*JobReq) ProtoMessage()    {}
func (*JobReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *JobReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobReq.Unmarshal(m, b)
}
func (m *JobReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobReq.Marshal(b, m, deterministic)
}
func (m *JobReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobReq.Merge(m, src)
}
func (m *JobReq) XXX_Size() int {
	return xxx_messageInfo_JobReq.Size(m)
}
func (m *JobReq) XXX_DiscardUnknown() {
	xxx_messageInfo_JobReq.DiscardUnknown(m)
}

var xxx_messageInfo_JobReq proto.InternalMessageInfo

func (m *JobReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Job struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (m *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(m, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Job) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type Log struct {
	Line                 string   `protobuf:"bytes,1,opt,name=line,proto3" json:"line,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Log.Unmarshal(m, b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Log.Marshal(b, m, deterministic)
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return xxx_messageInfo_Log.Size(m)
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

func init() {
	proto.RegisterType((*Command)(nil), "proto.Command")
	proto.RegisterType((*JobReq)(nil), "proto.JobReq")
	proto.RegisterType((*Job)(nil), "proto.Job")
	proto.RegisterType((*Log)(nil), "proto.Log")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8e, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0x27, 0xcd, 0x18, 0xf1, 0xc2, 0xcc, 0xe2, 0x2e, 0xa4, 0xba, 0x1a, 0x02, 0x8a, 0x1b,
	0x07, 0x7f, 0x1e, 0x61, 0x04, 0xa1, 0xcc, 0xaa, 0xf8, 0x02, 0x89, 0xbd, 0x94, 0x40, 0xda, 0x6b,
	0xd3, 0xe8, 0x03, 0xf9, 0xa4, 0x92, 0x18, 0x14, 0x0a, 0xe2, 0x2a, 0xe7, 0x7c, 0x7c, 0xe4, 0x5c,
	0xd8, 0xcc, 0x14, 0x3e, 0xdc, 0x2b, 0xed, 0xdf, 0x02, 0x47, 0xc6, 0x93, 0xfc, 0xe8, 0x7b, 0x38,
	0x3d, 0xf0, 0x30, 0x98, 0xb1, 0x43, 0x84, 0xf5, 0x68, 0x06, 0xaa, 0xc5, 0x4e, 0xdc, 0x9c, 0xb5,
	0x39, 0x27, 0x66, 0x42, 0x3f, 0xd7, 0xd5, 0x4e, 0x26, 0x96, 0xb2, 0xae, 0x41, 0x35, 0x6c, 0x5b,
	0x9a, 0x70, 0x0b, 0x95, 0xeb, 0x8a, 0x5f, 0xb9, 0x4e, 0xdf, 0x82, 0x6c, 0xd8, 0x2e, 0x31, 0x9e,
	0x83, 0x9a, 0xa3, 0x89, 0xef, 0xe9, 0x9b, 0xc4, 0x4a, 0xd3, 0x17, 0x20, 0x8f, 0xdc, 0xa7, 0x0d,
	0xef, 0xc6, 0x9f, 0xdd, 0x94, 0x1f, 0x3e, 0x45, 0x1e, 0xf1, 0x14, 0xf1, 0x1a, 0xd4, 0x21, 0x90,
	0x89, 0x84, 0xdb, 0xef, 0xd3, 0xf7, 0xe5, 0xe0, 0x4b, 0x28, 0xbd, 0x61, 0xab, 0x57, 0xa8, 0x41,
	0x3e, 0x53, 0xc4, 0xcd, 0x2f, 0x6c, 0x69, 0x5a, 0x38, 0x57, 0xa0, 0x9e, 0xc8, 0x53, 0xa4, 0xff,
	0xb4, 0xf5, 0x8b, 0x71, 0xfe, 0x2f, 0xe9, 0xc8, 0xbd, 0x5e, 0xdd, 0x09, 0xab, 0x72, 0x7d, 0xfc,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0x58, 0xad, 0x15, 0x2a, 0x5a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// JobletClient is the client API for Joblet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobletClient interface {
	Create(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Job, error)
	Get(ctx context.Context, in *JobReq, opts ...grpc.CallOption) (*Job, error)
	Delete(ctx context.Context, in *JobReq, opts ...grpc.CallOption) (*Job, error)
	Tail(ctx context.Context, in *JobReq, opts ...grpc.CallOption) (Joblet_TailClient, error)
}

type jobletClient struct {
	cc grpc.ClientConnInterface
}

func NewJobletClient(cc grpc.ClientConnInterface) JobletClient {
	return &jobletClient{cc}
}

func (c *jobletClient) Create(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := c.cc.Invoke(ctx, "/proto.Joblet/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobletClient) Get(ctx context.Context, in *JobReq, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := c.cc.Invoke(ctx, "/proto.Joblet/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobletClient) Delete(ctx context.Context, in *JobReq, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := c.cc.Invoke(ctx, "/proto.Joblet/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobletClient) Tail(ctx context.Context, in *JobReq, opts ...grpc.CallOption) (Joblet_TailClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Joblet_serviceDesc.Streams[0], "/proto.Joblet/Tail", opts...)
	if err != nil {
		return nil, err
	}
	x := &jobletTailClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Joblet_TailClient interface {
	Recv() (*Log, error)
	grpc.ClientStream
}

type jobletTailClient struct {
	grpc.ClientStream
}

func (x *jobletTailClient) Recv() (*Log, error) {
	m := new(Log)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// JobletServer is the server API for Joblet service.
type JobletServer interface {
	Create(context.Context, *Command) (*Job, error)
	Get(context.Context, *JobReq) (*Job, error)
	Delete(context.Context, *JobReq) (*Job, error)
	Tail(*JobReq, Joblet_TailServer) error
}

// UnimplementedJobletServer can be embedded to have forward compatible implementations.
type UnimplementedJobletServer struct {
}

func (*UnimplementedJobletServer) Create(ctx context.Context, req *Command) (*Job, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedJobletServer) Get(ctx context.Context, req *JobReq) (*Job, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedJobletServer) Delete(ctx context.Context, req *JobReq) (*Job, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedJobletServer) Tail(req *JobReq, srv Joblet_TailServer) error {
	return status.Errorf(codes.Unimplemented, "method Tail not implemented")
}

func RegisterJobletServer(s *grpc.Server, srv JobletServer) {
	s.RegisterService(&_Joblet_serviceDesc, srv)
}

func _Joblet_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobletServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Joblet/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobletServer).Create(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Joblet_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobletServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Joblet/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobletServer).Get(ctx, req.(*JobReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Joblet_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobletServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Joblet/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobletServer).Delete(ctx, req.(*JobReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Joblet_Tail_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JobReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(JobletServer).Tail(m, &jobletTailServer{stream})
}

type Joblet_TailServer interface {
	Send(*Log) error
	grpc.ServerStream
}

type jobletTailServer struct {
	grpc.ServerStream
}

func (x *jobletTailServer) Send(m *Log) error {
	return x.ServerStream.SendMsg(m)
}

var _Joblet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Joblet",
	HandlerType: (*JobletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Joblet_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Joblet_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Joblet_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tail",
			Handler:       _Joblet_Tail_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}