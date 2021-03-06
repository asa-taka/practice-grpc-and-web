// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

package todo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type Todo_Record_Status int32

const (
	Todo_Record_STATUS_TODO       Todo_Record_Status = 0
	Todo_Record_STATUS_INPROGRESS Todo_Record_Status = 1
	Todo_Record_STATUS_DONE       Todo_Record_Status = 2
)

var Todo_Record_Status_name = map[int32]string{
	0: "STATUS_TODO",
	1: "STATUS_INPROGRESS",
	2: "STATUS_DONE",
}
var Todo_Record_Status_value = map[string]int32{
	"STATUS_TODO":       0,
	"STATUS_INPROGRESS": 1,
	"STATUS_DONE":       2,
}

func (x Todo_Record_Status) String() string {
	return proto.EnumName(Todo_Record_Status_name, int32(x))
}
func (Todo_Record_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_todo_25e824b26792a881, []int{0, 0, 0}
}

type Todo struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_25e824b26792a881, []int{0}
}
func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (dst *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(dst, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

type Todo_Record struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Title                string               `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Detail               string               `protobuf:"bytes,4,opt,name=detail,proto3" json:"detail,omitempty"`
	Deadline             *timestamp.Timestamp `protobuf:"bytes,5,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Status               Todo_Record_Status   `protobuf:"varint,6,opt,name=status,proto3,enum=todo.Todo_Record_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Todo_Record) Reset()         { *m = Todo_Record{} }
func (m *Todo_Record) String() string { return proto.CompactTextString(m) }
func (*Todo_Record) ProtoMessage()    {}
func (*Todo_Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_25e824b26792a881, []int{0, 0}
}
func (m *Todo_Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo_Record.Unmarshal(m, b)
}
func (m *Todo_Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo_Record.Marshal(b, m, deterministic)
}
func (dst *Todo_Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo_Record.Merge(dst, src)
}
func (m *Todo_Record) XXX_Size() int {
	return xxx_messageInfo_Todo_Record.Size(m)
}
func (m *Todo_Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Todo_Record proto.InternalMessageInfo

func (m *Todo_Record) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Todo_Record) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Todo_Record) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Todo_Record) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func (m *Todo_Record) GetDeadline() *timestamp.Timestamp {
	if m != nil {
		return m.Deadline
	}
	return nil
}

func (m *Todo_Record) GetStatus() Todo_Record_Status {
	if m != nil {
		return m.Status
	}
	return Todo_Record_STATUS_TODO
}

type Todo_RecordInput struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Detail               string               `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
	Deadline             *timestamp.Timestamp `protobuf:"bytes,4,opt,name=deadline,proto3" json:"deadline,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Todo_RecordInput) Reset()         { *m = Todo_RecordInput{} }
func (m *Todo_RecordInput) String() string { return proto.CompactTextString(m) }
func (*Todo_RecordInput) ProtoMessage()    {}
func (*Todo_RecordInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_25e824b26792a881, []int{0, 1}
}
func (m *Todo_RecordInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo_RecordInput.Unmarshal(m, b)
}
func (m *Todo_RecordInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo_RecordInput.Marshal(b, m, deterministic)
}
func (dst *Todo_RecordInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo_RecordInput.Merge(dst, src)
}
func (m *Todo_RecordInput) XXX_Size() int {
	return xxx_messageInfo_Todo_RecordInput.Size(m)
}
func (m *Todo_RecordInput) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo_RecordInput.DiscardUnknown(m)
}

var xxx_messageInfo_Todo_RecordInput proto.InternalMessageInfo

func (m *Todo_RecordInput) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Todo_RecordInput) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Todo_RecordInput) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func (m *Todo_RecordInput) GetDeadline() *timestamp.Timestamp {
	if m != nil {
		return m.Deadline
	}
	return nil
}

type Todo_GetRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo_GetRequest) Reset()         { *m = Todo_GetRequest{} }
func (m *Todo_GetRequest) String() string { return proto.CompactTextString(m) }
func (*Todo_GetRequest) ProtoMessage()    {}
func (*Todo_GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_25e824b26792a881, []int{0, 2}
}
func (m *Todo_GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo_GetRequest.Unmarshal(m, b)
}
func (m *Todo_GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *Todo_GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo_GetRequest.Merge(dst, src)
}
func (m *Todo_GetRequest) XXX_Size() int {
	return xxx_messageInfo_Todo_GetRequest.Size(m)
}
func (m *Todo_GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_Todo_GetRequest proto.InternalMessageInfo

func (m *Todo_GetRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Todo_GetResponse struct {
	Record               *Todo_Record `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Todo_GetResponse) Reset()         { *m = Todo_GetResponse{} }
func (m *Todo_GetResponse) String() string { return proto.CompactTextString(m) }
func (*Todo_GetResponse) ProtoMessage()    {}
func (*Todo_GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_25e824b26792a881, []int{0, 3}
}
func (m *Todo_GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo_GetResponse.Unmarshal(m, b)
}
func (m *Todo_GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo_GetResponse.Marshal(b, m, deterministic)
}
func (dst *Todo_GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo_GetResponse.Merge(dst, src)
}
func (m *Todo_GetResponse) XXX_Size() int {
	return xxx_messageInfo_Todo_GetResponse.Size(m)
}
func (m *Todo_GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Todo_GetResponse proto.InternalMessageInfo

func (m *Todo_GetResponse) GetRecord() *Todo_Record {
	if m != nil {
		return m.Record
	}
	return nil
}

type Todo_QueryRequest struct {
	Start                *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End                  *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Todo_QueryRequest) Reset()         { *m = Todo_QueryRequest{} }
func (m *Todo_QueryRequest) String() string { return proto.CompactTextString(m) }
func (*Todo_QueryRequest) ProtoMessage()    {}
func (*Todo_QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_25e824b26792a881, []int{0, 4}
}
func (m *Todo_QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo_QueryRequest.Unmarshal(m, b)
}
func (m *Todo_QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo_QueryRequest.Marshal(b, m, deterministic)
}
func (dst *Todo_QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo_QueryRequest.Merge(dst, src)
}
func (m *Todo_QueryRequest) XXX_Size() int {
	return xxx_messageInfo_Todo_QueryRequest.Size(m)
}
func (m *Todo_QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_Todo_QueryRequest proto.InternalMessageInfo

func (m *Todo_QueryRequest) GetStart() *timestamp.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *Todo_QueryRequest) GetEnd() *timestamp.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

type Todo_QueryResponse struct {
	Records              []*Todo_Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Todo_QueryResponse) Reset()         { *m = Todo_QueryResponse{} }
func (m *Todo_QueryResponse) String() string { return proto.CompactTextString(m) }
func (*Todo_QueryResponse) ProtoMessage()    {}
func (*Todo_QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_25e824b26792a881, []int{0, 5}
}
func (m *Todo_QueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo_QueryResponse.Unmarshal(m, b)
}
func (m *Todo_QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo_QueryResponse.Marshal(b, m, deterministic)
}
func (dst *Todo_QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo_QueryResponse.Merge(dst, src)
}
func (m *Todo_QueryResponse) XXX_Size() int {
	return xxx_messageInfo_Todo_QueryResponse.Size(m)
}
func (m *Todo_QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Todo_QueryResponse proto.InternalMessageInfo

func (m *Todo_QueryResponse) GetRecords() []*Todo_Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type Todo_CreateRequest struct {
	Record               *Todo_RecordInput `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Todo_CreateRequest) Reset()         { *m = Todo_CreateRequest{} }
func (m *Todo_CreateRequest) String() string { return proto.CompactTextString(m) }
func (*Todo_CreateRequest) ProtoMessage()    {}
func (*Todo_CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_25e824b26792a881, []int{0, 6}
}
func (m *Todo_CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo_CreateRequest.Unmarshal(m, b)
}
func (m *Todo_CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo_CreateRequest.Marshal(b, m, deterministic)
}
func (dst *Todo_CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo_CreateRequest.Merge(dst, src)
}
func (m *Todo_CreateRequest) XXX_Size() int {
	return xxx_messageInfo_Todo_CreateRequest.Size(m)
}
func (m *Todo_CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_Todo_CreateRequest proto.InternalMessageInfo

func (m *Todo_CreateRequest) GetRecord() *Todo_RecordInput {
	if m != nil {
		return m.Record
	}
	return nil
}

type Todo_UpdateRequest struct {
	Record               *Todo_RecordInput `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	Id                   int32             `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Todo_UpdateRequest) Reset()         { *m = Todo_UpdateRequest{} }
func (m *Todo_UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*Todo_UpdateRequest) ProtoMessage()    {}
func (*Todo_UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_25e824b26792a881, []int{0, 7}
}
func (m *Todo_UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo_UpdateRequest.Unmarshal(m, b)
}
func (m *Todo_UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo_UpdateRequest.Marshal(b, m, deterministic)
}
func (dst *Todo_UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo_UpdateRequest.Merge(dst, src)
}
func (m *Todo_UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_Todo_UpdateRequest.Size(m)
}
func (m *Todo_UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_Todo_UpdateRequest proto.InternalMessageInfo

func (m *Todo_UpdateRequest) GetRecord() *Todo_RecordInput {
	if m != nil {
		return m.Record
	}
	return nil
}

func (m *Todo_UpdateRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Todo_DeleteRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo_DeleteRequest) Reset()         { *m = Todo_DeleteRequest{} }
func (m *Todo_DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*Todo_DeleteRequest) ProtoMessage()    {}
func (*Todo_DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_25e824b26792a881, []int{0, 8}
}
func (m *Todo_DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo_DeleteRequest.Unmarshal(m, b)
}
func (m *Todo_DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo_DeleteRequest.Marshal(b, m, deterministic)
}
func (dst *Todo_DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo_DeleteRequest.Merge(dst, src)
}
func (m *Todo_DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_Todo_DeleteRequest.Size(m)
}
func (m *Todo_DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_Todo_DeleteRequest proto.InternalMessageInfo

func (m *Todo_DeleteRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Todo_MutateResponse struct {
	Record               *Todo_Record `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Todo_MutateResponse) Reset()         { *m = Todo_MutateResponse{} }
func (m *Todo_MutateResponse) String() string { return proto.CompactTextString(m) }
func (*Todo_MutateResponse) ProtoMessage()    {}
func (*Todo_MutateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_25e824b26792a881, []int{0, 9}
}
func (m *Todo_MutateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo_MutateResponse.Unmarshal(m, b)
}
func (m *Todo_MutateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo_MutateResponse.Marshal(b, m, deterministic)
}
func (dst *Todo_MutateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo_MutateResponse.Merge(dst, src)
}
func (m *Todo_MutateResponse) XXX_Size() int {
	return xxx_messageInfo_Todo_MutateResponse.Size(m)
}
func (m *Todo_MutateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo_MutateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Todo_MutateResponse proto.InternalMessageInfo

func (m *Todo_MutateResponse) GetRecord() *Todo_Record {
	if m != nil {
		return m.Record
	}
	return nil
}

func init() {
	proto.RegisterType((*Todo)(nil), "todo.Todo")
	proto.RegisterType((*Todo_Record)(nil), "todo.Todo.Record")
	proto.RegisterType((*Todo_RecordInput)(nil), "todo.Todo.RecordInput")
	proto.RegisterType((*Todo_GetRequest)(nil), "todo.Todo.GetRequest")
	proto.RegisterType((*Todo_GetResponse)(nil), "todo.Todo.GetResponse")
	proto.RegisterType((*Todo_QueryRequest)(nil), "todo.Todo.QueryRequest")
	proto.RegisterType((*Todo_QueryResponse)(nil), "todo.Todo.QueryResponse")
	proto.RegisterType((*Todo_CreateRequest)(nil), "todo.Todo.CreateRequest")
	proto.RegisterType((*Todo_UpdateRequest)(nil), "todo.Todo.UpdateRequest")
	proto.RegisterType((*Todo_DeleteRequest)(nil), "todo.Todo.DeleteRequest")
	proto.RegisterType((*Todo_MutateResponse)(nil), "todo.Todo.MutateResponse")
	proto.RegisterEnum("todo.Todo_Record_Status", Todo_Record_Status_name, Todo_Record_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoServiceClient interface {
	QueryTodo(ctx context.Context, in *Todo_QueryResponse, opts ...grpc.CallOption) (*Todo_QueryRequest, error)
	GetTodo(ctx context.Context, in *Todo_GetRequest, opts ...grpc.CallOption) (*Todo_GetResponse, error)
	CreateTodo(ctx context.Context, in *Todo_CreateRequest, opts ...grpc.CallOption) (*Todo_MutateResponse, error)
	UpdateTodo(ctx context.Context, in *Todo_UpdateRequest, opts ...grpc.CallOption) (*Todo_MutateResponse, error)
	DeleteTodo(ctx context.Context, in *Todo_DeleteRequest, opts ...grpc.CallOption) (*Todo_MutateResponse, error)
}

type todoServiceClient struct {
	cc *grpc.ClientConn
}

func NewTodoServiceClient(cc *grpc.ClientConn) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) QueryTodo(ctx context.Context, in *Todo_QueryResponse, opts ...grpc.CallOption) (*Todo_QueryRequest, error) {
	out := new(Todo_QueryRequest)
	err := c.cc.Invoke(ctx, "/todo.TodoService/QueryTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetTodo(ctx context.Context, in *Todo_GetRequest, opts ...grpc.CallOption) (*Todo_GetResponse, error) {
	out := new(Todo_GetResponse)
	err := c.cc.Invoke(ctx, "/todo.TodoService/GetTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) CreateTodo(ctx context.Context, in *Todo_CreateRequest, opts ...grpc.CallOption) (*Todo_MutateResponse, error) {
	out := new(Todo_MutateResponse)
	err := c.cc.Invoke(ctx, "/todo.TodoService/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) UpdateTodo(ctx context.Context, in *Todo_UpdateRequest, opts ...grpc.CallOption) (*Todo_MutateResponse, error) {
	out := new(Todo_MutateResponse)
	err := c.cc.Invoke(ctx, "/todo.TodoService/UpdateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) DeleteTodo(ctx context.Context, in *Todo_DeleteRequest, opts ...grpc.CallOption) (*Todo_MutateResponse, error) {
	out := new(Todo_MutateResponse)
	err := c.cc.Invoke(ctx, "/todo.TodoService/DeleteTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	QueryTodo(context.Context, *Todo_QueryResponse) (*Todo_QueryRequest, error)
	GetTodo(context.Context, *Todo_GetRequest) (*Todo_GetResponse, error)
	CreateTodo(context.Context, *Todo_CreateRequest) (*Todo_MutateResponse, error)
	UpdateTodo(context.Context, *Todo_UpdateRequest) (*Todo_MutateResponse, error)
	DeleteTodo(context.Context, *Todo_DeleteRequest) (*Todo_MutateResponse, error)
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_QueryTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo_QueryResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).QueryTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/QueryTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).QueryTodo(ctx, req.(*Todo_QueryResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo_GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/GetTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetTodo(ctx, req.(*Todo_GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo_CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CreateTodo(ctx, req.(*Todo_CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_UpdateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo_UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).UpdateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/UpdateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).UpdateTodo(ctx, req.(*Todo_UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_DeleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo_DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).DeleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/DeleteTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).DeleteTodo(ctx, req.(*Todo_DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryTodo",
			Handler:    _TodoService_QueryTodo_Handler,
		},
		{
			MethodName: "GetTodo",
			Handler:    _TodoService_GetTodo_Handler,
		},
		{
			MethodName: "CreateTodo",
			Handler:    _TodoService_CreateTodo_Handler,
		},
		{
			MethodName: "UpdateTodo",
			Handler:    _TodoService_UpdateTodo_Handler,
		},
		{
			MethodName: "DeleteTodo",
			Handler:    _TodoService_DeleteTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor_todo_25e824b26792a881) }

var fileDescriptor_todo_25e824b26792a881 = []byte{
	// 707 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x41, 0x6f, 0xd3, 0x48,
	0x18, 0xdd, 0x71, 0x12, 0xb7, 0xfd, 0xb2, 0xe9, 0xa6, 0xa3, 0x4d, 0xeb, 0x9d, 0x05, 0x6a, 0x85,
	0x4b, 0x28, 0x8d, 0x9d, 0x06, 0x09, 0x41, 0x41, 0x42, 0x86, 0x96, 0x52, 0x21, 0x1a, 0x70, 0x52,
	0xc1, 0x89, 0xe2, 0xda, 0x43, 0x64, 0x29, 0xb5, 0x8d, 0x3d, 0x69, 0x85, 0x10, 0x17, 0xb8, 0x72,
	0x2a, 0xff, 0x2c, 0xfc, 0x05, 0x24, 0x7e, 0x04, 0x17, 0xe4, 0x99, 0x49, 0x9a, 0xc4, 0xa1, 0x14,
	0x4e, 0x89, 0xbf, 0x79, 0xf3, 0xde, 0xfb, 0xde, 0xb3, 0x01, 0x58, 0xe8, 0x85, 0x46, 0x14, 0x87,
	0x2c, 0xc4, 0xf9, 0xf4, 0x3f, 0x59, 0xed, 0x86, 0x61, 0xb7, 0x47, 0x4d, 0x3e, 0x3b, 0xec, 0xbf,
	0x36, 0x99, 0x7f, 0x44, 0x13, 0xe6, 0x1c, 0x45, 0x02, 0x46, 0x2e, 0x49, 0x80, 0x13, 0xf9, 0xa6,
	0x13, 0x04, 0x21, 0x73, 0x98, 0x1f, 0x06, 0x89, 0x3c, 0x5d, 0xe7, 0x3f, 0x6e, 0xbd, 0x4b, 0x83,
	0x7a, 0x72, 0xe2, 0x74, 0xbb, 0x34, 0x36, 0xc3, 0x88, 0x23, 0xb2, 0xe8, 0xea, 0xa7, 0x79, 0xc8,
	0x77, 0x52, 0xd5, 0xef, 0x0a, 0xa8, 0x36, 0x75, 0xc3, 0xd8, 0xc3, 0x8b, 0xa0, 0xf8, 0x9e, 0x86,
	0x74, 0x54, 0x2b, 0xd8, 0x8a, 0xef, 0xe1, 0xdb, 0x00, 0x6e, 0x4c, 0x1d, 0x46, 0xbd, 0x03, 0x87,
	0x69, 0x8a, 0x8e, 0x6a, 0xc5, 0x26, 0x31, 0x84, 0x09, 0x63, 0xe8, 0xd2, 0xe8, 0x0c, 0x5d, 0xda,
	0x0b, 0x12, 0x6d, 0x31, 0xfc, 0x2f, 0x14, 0x98, 0xcf, 0x7a, 0x54, 0xcb, 0xe9, 0xa8, 0xb6, 0x60,
	0x8b, 0x07, 0xbc, 0x0c, 0xaa, 0x47, 0x99, 0xe3, 0xf7, 0xb4, 0x3c, 0x1f, 0xcb, 0x27, 0x7c, 0x13,
	0xe6, 0x3d, 0xea, 0x78, 0x3d, 0x3f, 0xa0, 0x5a, 0xe1, 0x97, 0x32, 0x23, 0x2c, 0x6e, 0x80, 0x9a,
	0x30, 0x87, 0xf5, 0x13, 0x4d, 0xd5, 0x51, 0x6d, 0xb1, 0xa9, 0x19, 0x3c, 0xd4, 0x74, 0x2f, 0x43,
	0xec, 0x64, 0xb4, 0xf9, 0xb9, 0x2d, 0x71, 0x55, 0x0b, 0x54, 0x31, 0xc1, 0xff, 0x40, 0xb1, 0xdd,
	0xb1, 0x3a, 0xfb, 0xed, 0x83, 0x4e, 0x6b, 0xab, 0x55, 0xfe, 0x0b, 0x57, 0x60, 0x49, 0x0e, 0x76,
	0xf7, 0x9e, 0xda, 0xad, 0x1d, 0x7b, 0xbb, 0xdd, 0x2e, 0xa3, 0x31, 0xdc, 0x56, 0x6b, 0x6f, 0xbb,
	0xac, 0x6c, 0x3e, 0x3e, 0xb5, 0x1e, 0xc1, 0xc3, 0xb5, 0x92, 0xa5, 0xa7, 0x3a, 0xba, 0xd0, 0x19,
	0x20, 0xc5, 0xf7, 0x06, 0x68, 0x2c, 0xaf, 0x01, 0x12, 0x4b, 0x0f, 0x90, 0xdc, 0x72, 0x80, 0x46,
	0xc6, 0x07, 0x48, 0xfa, 0x21, 0x1f, 0x11, 0x14, 0x05, 0xc3, 0x6e, 0x10, 0xf5, 0x59, 0xa6, 0x82,
	0x51, 0x8e, 0xca, 0xec, 0x1c, 0x73, 0x3f, 0xcd, 0x31, 0x7f, 0xf1, 0x1c, 0x49, 0x0d, 0x60, 0x87,
	0x32, 0x9b, 0xbe, 0xe9, 0xd3, 0x24, 0xe3, 0x61, 0x13, 0x4e, 0xad, 0x39, 0x28, 0xf0, 0x0d, 0xc9,
	0x2d, 0x28, 0x72, 0x64, 0x12, 0x85, 0x41, 0x42, 0xf1, 0x35, 0x50, 0x63, 0xee, 0x9e, 0xc3, 0x8b,
	0xcd, 0xa5, 0x4c, 0x01, 0xb6, 0x04, 0x90, 0x00, 0xfe, 0x7e, 0xd6, 0xa7, 0xf1, 0xdb, 0xa1, 0x4a,
	0x03, 0x0a, 0x09, 0x73, 0x62, 0x26, 0x6f, 0x9e, 0x67, 0x54, 0x00, 0xf1, 0x3a, 0xe4, 0x68, 0xe0,
	0x5d, 0xe0, 0x3d, 0x4c, 0x61, 0xe4, 0x2e, 0x94, 0xa4, 0x9e, 0xf4, 0x7a, 0x1d, 0xe6, 0x84, 0x95,
	0x44, 0x43, 0x7a, 0x6e, 0xb6, 0xd9, 0x21, 0x82, 0xdc, 0x83, 0xd2, 0x03, 0x5e, 0xe5, 0xd0, 0xae,
	0x31, 0xb5, 0xe9, 0x72, 0xe6, 0x32, 0x2f, 0x70, 0xb4, 0x6e, 0x0b, 0x4a, 0xfb, 0x91, 0xf7, 0xe7,
	0x04, 0xb2, 0x05, 0x65, 0xd8, 0x02, 0x59, 0x85, 0xd2, 0x16, 0xed, 0xd1, 0x33, 0xc2, 0xa9, 0x9a,
	0xc8, 0x1d, 0x58, 0x7c, 0xd2, 0x67, 0x5c, 0xf1, 0xb7, 0xdb, 0x69, 0x7e, 0xcb, 0x41, 0x31, 0x9d,
	0xb7, 0x69, 0x7c, 0xec, 0xbb, 0x14, 0xbf, 0x80, 0x05, 0x9e, 0x5e, 0x3a, 0xc3, 0xe3, 0x9f, 0xd5,
	0x44, 0xa6, 0x64, 0x25, 0x7b, 0xc2, 0xcd, 0x55, 0xb5, 0x0f, 0x5f, 0xbe, 0x7e, 0x56, 0x30, 0x2e,
	0x9b, 0xc7, 0x1b, 0x66, 0x8a, 0x31, 0x65, 0xb2, 0xf8, 0x39, 0xcc, 0xed, 0x50, 0xc6, 0x79, 0x2b,
	0x63, 0xb7, 0xcf, 0xde, 0x3f, 0xb2, 0x3c, 0x3d, 0x16, 0x62, 0xd5, 0xcb, 0x9c, 0x73, 0x05, 0x57,
	0xa6, 0x39, 0xcd, 0x77, 0xbe, 0xf7, 0x1e, 0xbf, 0x04, 0x10, 0x95, 0x65, 0x3c, 0x4f, 0x34, 0x49,
	0xfe, 0x1b, 0x3b, 0x99, 0x0c, 0xac, 0xfa, 0x3f, 0x57, 0xa8, 0x54, 0x33, 0xae, 0x37, 0xd1, 0x1a,
	0x76, 0x01, 0x44, 0xa3, 0x19, 0xfe, 0x89, 0xa2, 0xcf, 0xe3, 0xd7, 0x39, 0x3f, 0x21, 0xb3, 0x37,
	0x48, 0x45, 0x5e, 0x01, 0x88, 0x96, 0x33, 0x22, 0x13, 0xe5, 0x9f, 0x27, 0x22, 0x63, 0x5a, 0x9b,
	0x2d, 0x72, 0xff, 0xea, 0xa9, 0xa5, 0xe3, 0x2b, 0x50, 0x89, 0x62, 0xc7, 0x65, 0xbe, 0x4b, 0xeb,
	0xdd, 0x38, 0x72, 0xeb, 0x4e, 0xe0, 0xd5, 0x4f, 0xe8, 0x61, 0xb3, 0xd0, 0x30, 0x1a, 0xc6, 0xc6,
	0xa1, 0xca, 0xbf, 0xaa, 0x1b, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x50, 0xe3, 0x72, 0xa5,
	0x06, 0x00, 0x00,
}
