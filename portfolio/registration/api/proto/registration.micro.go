// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/registration.proto

package registration

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Client API for Registration service

type RegistrationService interface {
	Signup(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Count(ctx context.Context, in *CountRequest, opts ...client.CallOption) (*CountResponse, error)
}

type registrationService struct {
	c    client.Client
	name string
}

func NewRegistrationService(name string, c client.Client) RegistrationService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "registration"
	}
	return &registrationService{
		c:    c,
		name: name,
	}
}

func (c *registrationService) Signup(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Registration.Signup", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationService) Count(ctx context.Context, in *CountRequest, opts ...client.CallOption) (*CountResponse, error) {
	req := c.c.NewRequest(c.name, "Registration.Count", in)
	out := new(CountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Registration service

type RegistrationHandler interface {
	Signup(context.Context, *User, *Response) error
	Count(context.Context, *CountRequest, *CountResponse) error
}

func RegisterRegistrationHandler(s server.Server, hdlr RegistrationHandler, opts ...server.HandlerOption) error {
	type registration interface {
		Signup(ctx context.Context, in *User, out *Response) error
		Count(ctx context.Context, in *CountRequest, out *CountResponse) error
	}
	type Registration struct {
		registration
	}
	h := &registrationHandler{hdlr}
	return s.Handle(s.NewHandler(&Registration{h}, opts...))
}

type registrationHandler struct {
	RegistrationHandler
}

func (h *registrationHandler) Signup(ctx context.Context, in *User, out *Response) error {
	return h.RegistrationHandler.Signup(ctx, in, out)
}

func (h *registrationHandler) Count(ctx context.Context, in *CountRequest, out *CountResponse) error {
	return h.RegistrationHandler.Count(ctx, in, out)
}