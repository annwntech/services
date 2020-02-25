// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/events/events.proto

package go_micro_srv_events

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Events service

type EventsService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Events_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Events_PingPongService, error)
}

type eventsService struct {
	c    client.Client
	name string
}

func NewEventsService(name string, c client.Client) EventsService {
	return &eventsService{
		c:    c,
		name: name,
	}
}

func (c *eventsService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Events.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Events_StreamService, error) {
	req := c.c.NewRequest(c.name, "Events.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &eventsServiceStream{stream}, nil
}

type Events_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type eventsServiceStream struct {
	stream client.Stream
}

func (x *eventsServiceStream) Close() error {
	return x.stream.Close()
}

func (x *eventsServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *eventsServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *eventsServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *eventsServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventsService) PingPong(ctx context.Context, opts ...client.CallOption) (Events_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Events.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &eventsServicePingPong{stream}, nil
}

type Events_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type eventsServicePingPong struct {
	stream client.Stream
}

func (x *eventsServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *eventsServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *eventsServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *eventsServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *eventsServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *eventsServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Events service

type EventsHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Events_StreamStream) error
	PingPong(context.Context, Events_PingPongStream) error
}

func RegisterEventsHandler(s server.Server, hdlr EventsHandler, opts ...server.HandlerOption) error {
	type events interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Events struct {
		events
	}
	h := &eventsHandler{hdlr}
	return s.Handle(s.NewHandler(&Events{h}, opts...))
}

type eventsHandler struct {
	EventsHandler
}

func (h *eventsHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.EventsHandler.Call(ctx, in, out)
}

func (h *eventsHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.EventsHandler.Stream(ctx, m, &eventsStreamStream{stream})
}

type Events_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type eventsStreamStream struct {
	stream server.Stream
}

func (x *eventsStreamStream) Close() error {
	return x.stream.Close()
}

func (x *eventsStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *eventsStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *eventsStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *eventsStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *eventsHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.EventsHandler.PingPong(ctx, &eventsPingPongStream{stream})
}

type Events_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type eventsPingPongStream struct {
	stream server.Stream
}

func (x *eventsPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *eventsPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *eventsPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *eventsPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *eventsPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *eventsPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
