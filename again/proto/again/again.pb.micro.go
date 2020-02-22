// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/again/again.proto

package go_micro_srv_again

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

// Client API for Again service

type AgainService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Again_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Again_PingPongService, error)
}

type againService struct {
	c    client.Client
	name string
}

func NewAgainService(name string, c client.Client) AgainService {
	return &againService{
		c:    c,
		name: name,
	}
}

func (c *againService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Again.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *againService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Again_StreamService, error) {
	req := c.c.NewRequest(c.name, "Again.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &againServiceStream{stream}, nil
}

type Again_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type againServiceStream struct {
	stream client.Stream
}

func (x *againServiceStream) Close() error {
	return x.stream.Close()
}

func (x *againServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *againServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *againServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *againServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *againService) PingPong(ctx context.Context, opts ...client.CallOption) (Again_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Again.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &againServicePingPong{stream}, nil
}

type Again_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type againServicePingPong struct {
	stream client.Stream
}

func (x *againServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *againServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *againServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *againServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *againServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *againServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Again service

type AgainHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Again_StreamStream) error
	PingPong(context.Context, Again_PingPongStream) error
}

func RegisterAgainHandler(s server.Server, hdlr AgainHandler, opts ...server.HandlerOption) error {
	type again interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Again struct {
		again
	}
	h := &againHandler{hdlr}
	return s.Handle(s.NewHandler(&Again{h}, opts...))
}

type againHandler struct {
	AgainHandler
}

func (h *againHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.AgainHandler.Call(ctx, in, out)
}

func (h *againHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.AgainHandler.Stream(ctx, m, &againStreamStream{stream})
}

type Again_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type againStreamStream struct {
	stream server.Stream
}

func (x *againStreamStream) Close() error {
	return x.stream.Close()
}

func (x *againStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *againStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *againStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *againStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *againHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.AgainHandler.PingPong(ctx, &againPingPongStream{stream})
}

type Again_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type againPingPongStream struct {
	stream server.Stream
}

func (x *againPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *againPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *againPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *againPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *againPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *againPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
