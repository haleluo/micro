// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: broker/service/proto/broker.proto

package go_micro_broker

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

// Api Endpoints for Broker service

func NewBrokerEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Broker service

type BrokerService interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...client.CallOption) (*Empty, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...client.CallOption) (Broker_SubscribeService, error)
}

type brokerService struct {
	c    client.Client
	name string
}

func NewBrokerService(name string, c client.Client) BrokerService {
	return &brokerService{
		c:    c,
		name: name,
	}
}

func (c *brokerService) Publish(ctx context.Context, in *PublishRequest, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "Broker.Publish", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerService) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...client.CallOption) (Broker_SubscribeService, error) {
	req := c.c.NewRequest(c.name, "Broker.Subscribe", &SubscribeRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &brokerServiceSubscribe{stream}, nil
}

type Broker_SubscribeService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*Message, error)
}

type brokerServiceSubscribe struct {
	stream client.Stream
}

func (x *brokerServiceSubscribe) Close() error {
	return x.stream.Close()
}

func (x *brokerServiceSubscribe) Context() context.Context {
	return x.stream.Context()
}

func (x *brokerServiceSubscribe) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *brokerServiceSubscribe) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *brokerServiceSubscribe) Recv() (*Message, error) {
	m := new(Message)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Broker service

type BrokerHandler interface {
	Publish(context.Context, *PublishRequest, *Empty) error
	Subscribe(context.Context, *SubscribeRequest, Broker_SubscribeStream) error
}

func RegisterBrokerHandler(s server.Server, hdlr BrokerHandler, opts ...server.HandlerOption) error {
	type broker interface {
		Publish(ctx context.Context, in *PublishRequest, out *Empty) error
		Subscribe(ctx context.Context, stream server.Stream) error
	}
	type Broker struct {
		broker
	}
	h := &brokerHandler{hdlr}
	return s.Handle(s.NewHandler(&Broker{h}, opts...))
}

type brokerHandler struct {
	BrokerHandler
}

func (h *brokerHandler) Publish(ctx context.Context, in *PublishRequest, out *Empty) error {
	return h.BrokerHandler.Publish(ctx, in, out)
}

func (h *brokerHandler) Subscribe(ctx context.Context, stream server.Stream) error {
	m := new(SubscribeRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.BrokerHandler.Subscribe(ctx, m, &brokerSubscribeStream{stream})
}

type Broker_SubscribeStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Message) error
}

type brokerSubscribeStream struct {
	stream server.Stream
}

func (x *brokerSubscribeStream) Close() error {
	return x.stream.Close()
}

func (x *brokerSubscribeStream) Context() context.Context {
	return x.stream.Context()
}

func (x *brokerSubscribeStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *brokerSubscribeStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *brokerSubscribeStream) Send(m *Message) error {
	return x.stream.Send(m)
}
