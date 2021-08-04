// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: runtime/service/proto/runtime.proto

package go_micro_runtime

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/haleluo/micro/v2/api"
	client "github.com/haleluo/micro/v2/client"
	server "github.com/haleluo/micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Runtime service

func NewRuntimeEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Runtime service

type RuntimeService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Logs(ctx context.Context, in *LogsRequest, opts ...client.CallOption) (Runtime_LogsService, error)
}

type runtimeService struct {
	c    client.Client
	name string
}

func NewRuntimeService(name string, c client.Client) RuntimeService {
	return &runtimeService{
		c:    c,
		name: name,
	}
}

func (c *runtimeService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "Runtime.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "Runtime.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Runtime.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "Runtime.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeService) Logs(ctx context.Context, in *LogsRequest, opts ...client.CallOption) (Runtime_LogsService, error) {
	req := c.c.NewRequest(c.name, "Runtime.Logs", &LogsRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &runtimeServiceLogs{stream}, nil
}

type Runtime_LogsService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*LogRecord, error)
}

type runtimeServiceLogs struct {
	stream client.Stream
}

func (x *runtimeServiceLogs) Close() error {
	return x.stream.Close()
}

func (x *runtimeServiceLogs) Context() context.Context {
	return x.stream.Context()
}

func (x *runtimeServiceLogs) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *runtimeServiceLogs) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *runtimeServiceLogs) Recv() (*LogRecord, error) {
	m := new(LogRecord)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Runtime service

type RuntimeHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Logs(context.Context, *LogsRequest, Runtime_LogsStream) error
}

func RegisterRuntimeHandler(s server.Server, hdlr RuntimeHandler, opts ...server.HandlerOption) error {
	type runtime interface {
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		Logs(ctx context.Context, stream server.Stream) error
	}
	type Runtime struct {
		runtime
	}
	h := &runtimeHandler{hdlr}
	return s.Handle(s.NewHandler(&Runtime{h}, opts...))
}

type runtimeHandler struct {
	RuntimeHandler
}

func (h *runtimeHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.RuntimeHandler.Create(ctx, in, out)
}

func (h *runtimeHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.RuntimeHandler.Read(ctx, in, out)
}

func (h *runtimeHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.RuntimeHandler.Delete(ctx, in, out)
}

func (h *runtimeHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.RuntimeHandler.Update(ctx, in, out)
}

func (h *runtimeHandler) Logs(ctx context.Context, stream server.Stream) error {
	m := new(LogsRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.RuntimeHandler.Logs(ctx, m, &runtimeLogsStream{stream})
}

type Runtime_LogsStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*LogRecord) error
}

type runtimeLogsStream struct {
	stream server.Stream
}

func (x *runtimeLogsStream) Close() error {
	return x.stream.Close()
}

func (x *runtimeLogsStream) Context() context.Context {
	return x.stream.Context()
}

func (x *runtimeLogsStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *runtimeLogsStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *runtimeLogsStream) Send(m *LogRecord) error {
	return x.stream.Send(m)
}
