// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

/*
Package todo is a generated protocol buffer package.

It is generated from these files:
	todo.proto

It has these top-level messages:
	Task
	TaskTitle
	TaskList
	Void
*/
package todo

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

type Task struct {
	Title string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Done  bool   `protobuf:"varint,2,opt,name=done" json:"done,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Task) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Task) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

type TaskTitle struct {
	Title string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
}

func (m *TaskTitle) Reset()                    { *m = TaskTitle{} }
func (m *TaskTitle) String() string            { return proto.CompactTextString(m) }
func (*TaskTitle) ProtoMessage()               {}
func (*TaskTitle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TaskTitle) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type TaskList struct {
	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks" json:"tasks,omitempty"`
}

func (m *TaskList) Reset()                    { *m = TaskList{} }
func (m *TaskList) String() string            { return proto.CompactTextString(m) }
func (*TaskList) ProtoMessage()               {}
func (*TaskList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TaskList) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type Void struct {
}

func (m *Void) Reset()                    { *m = Void{} }
func (m *Void) String() string            { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()               {}
func (*Void) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Task)(nil), "todo.Task")
	proto.RegisterType((*TaskTitle)(nil), "todo.TaskTitle")
	proto.RegisterType((*TaskList)(nil), "todo.TaskList")
	proto.RegisterType((*Void)(nil), "todo.Void")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Tasks service

type TasksClient interface {
	List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*TaskList, error)
	Add(ctx context.Context, in *TaskTitle, opts ...grpc.CallOption) (*Task, error)
	Done(ctx context.Context, in *TaskTitle, opts ...grpc.CallOption) (*Task, error)
}

type tasksClient struct {
	cc *grpc.ClientConn
}

func NewTasksClient(cc *grpc.ClientConn) TasksClient {
	return &tasksClient{cc}
}

func (c *tasksClient) List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*TaskList, error) {
	out := new(TaskList)
	err := grpc.Invoke(ctx, "/todo.Tasks/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) Add(ctx context.Context, in *TaskTitle, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := grpc.Invoke(ctx, "/todo.Tasks/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) Done(ctx context.Context, in *TaskTitle, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := grpc.Invoke(ctx, "/todo.Tasks/Done", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Tasks service

type TasksServer interface {
	List(context.Context, *Void) (*TaskList, error)
	Add(context.Context, *TaskTitle) (*Task, error)
	Done(context.Context, *TaskTitle) (*Task, error)
}

func RegisterTasksServer(s *grpc.Server, srv TasksServer) {
	s.RegisterService(&_Tasks_serviceDesc, srv)
}

func _Tasks_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.Tasks/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).List(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskTitle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.Tasks/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).Add(ctx, req.(*TaskTitle))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_Done_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskTitle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).Done(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.Tasks/Done",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).Done(ctx, req.(*TaskTitle))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tasks_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.Tasks",
	HandlerType: (*TasksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Tasks_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Tasks_Add_Handler,
		},
		{
			MethodName: "Done",
			Handler:    _Tasks_Done_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0xab, 0x83, 0x30,
	0x14, 0x46, 0xcd, 0x33, 0x8a, 0x7e, 0x0f, 0xde, 0x83, 0x4b, 0x07, 0x71, 0x4a, 0x83, 0x05, 0x87,
	0x22, 0xc5, 0xfe, 0x82, 0x42, 0xc7, 0x4e, 0x22, 0xdd, 0x2d, 0x71, 0x10, 0x8b, 0x29, 0x4d, 0xa6,
	0xfe, 0xfa, 0x92, 0x58, 0xb0, 0x43, 0xa1, 0xdb, 0x4d, 0xce, 0x39, 0xe4, 0x12, 0xc0, 0x6a, 0xa5,
	0xab, 0xdb, 0x5d, 0x5b, 0x4d, 0xdc, 0xcd, 0x72, 0x07, 0xde, 0x76, 0x66, 0xa4, 0x15, 0x22, 0x3b,
	0xd8, 0x6b, 0x9f, 0x31, 0xc1, 0xca, 0xb4, 0x99, 0x0f, 0x44, 0xe0, 0x4a, 0x4f, 0x7d, 0xf6, 0x23,
	0x58, 0x99, 0x34, 0x7e, 0x96, 0x6b, 0xa4, 0xae, 0x68, 0xbd, 0xf0, 0x31, 0x93, 0x5b, 0x24, 0x4e,
	0x39, 0x0d, 0xc6, 0x92, 0x40, 0x64, 0x3b, 0x33, 0x9a, 0x8c, 0x89, 0xb0, 0xfc, 0xad, 0x51, 0xf9,
	0x15, 0x1c, 0x6e, 0x66, 0x20, 0x63, 0xf0, 0xb3, 0x1e, 0x54, 0xfd, 0x40, 0xe4, 0xae, 0x0d, 0x15,
	0xe0, 0x3e, 0x7d, 0xb9, 0x0e, 0xe6, 0x7f, 0x4b, 0xe7, 0x98, 0x0c, 0xa8, 0x40, 0x78, 0x50, 0x8a,
	0xfe, 0x17, 0xe0, 0x57, 0xca, 0xdf, 0x5e, 0x90, 0x01, 0x6d, 0xc0, 0x8f, 0x7a, 0xea, 0xbf, 0x68,
	0x97, 0xd8, 0xff, 0xc9, 0xfe, 0x19, 0x00, 0x00, 0xff, 0xff, 0x54, 0xf7, 0xe8, 0x06, 0x21, 0x01,
	0x00, 0x00,
}
