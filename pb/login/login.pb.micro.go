// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/login.proto

package login

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Login service

func NewLoginEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Login service

type LoginService interface {
	// 定義 API Interface，Greet 為此 API 的 Name，
	// 代表 給 Greet API Request 當參數，並返回 Response
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
}

type loginService struct {
	c    client.Client
	name string
}

func NewLoginService(name string, c client.Client) LoginService {
	return &loginService{
		c:    c,
		name: name,
	}
}

func (c *loginService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "Login.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Login service

type LoginHandler interface {
	// 定義 API Interface，Greet 為此 API 的 Name，
	// 代表 給 Greet API Request 當參數，並返回 Response
	Login(context.Context, *LoginRequest, *LoginResponse) error
}

func RegisterLoginHandler(s server.Server, hdlr LoginHandler, opts ...server.HandlerOption) error {
	type login interface {
		Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error
	}
	type Login struct {
		login
	}
	h := &loginHandler{hdlr}
	return s.Handle(s.NewHandler(&Login{h}, opts...))
}

type loginHandler struct {
	LoginHandler
}

func (h *loginHandler) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.LoginHandler.Login(ctx, in, out)
}
