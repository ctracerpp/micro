// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/pay.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/pay.proto

It has these top-level messages:
	GetPayInfoReq
	GetPayInfoRes
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Pay service

type PayService interface {
	GetPayInfo(ctx context.Context, in *GetPayInfoReq, opts ...client.CallOption) (*GetPayInfoRes, error)
}

type payService struct {
	c    client.Client
	name string
}

func NewPayService(name string, c client.Client) PayService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "proto"
	}
	return &payService{
		c:    c,
		name: name,
	}
}

func (c *payService) GetPayInfo(ctx context.Context, in *GetPayInfoReq, opts ...client.CallOption) (*GetPayInfoRes, error) {
	req := c.c.NewRequest(c.name, "Pay.GetPayInfo", in)
	out := new(GetPayInfoRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Pay service

type PayHandler interface {
	GetPayInfo(context.Context, *GetPayInfoReq, *GetPayInfoRes) error
}

func RegisterPayHandler(s server.Server, hdlr PayHandler, opts ...server.HandlerOption) error {
	type pay interface {
		GetPayInfo(ctx context.Context, in *GetPayInfoReq, out *GetPayInfoRes) error
	}
	type Pay struct {
		pay
	}
	h := &payHandler{hdlr}
	return s.Handle(s.NewHandler(&Pay{h}, opts...))
}

type payHandler struct {
	PayHandler
}

func (h *payHandler) GetPayInfo(ctx context.Context, in *GetPayInfoReq, out *GetPayInfoRes) error {
	return h.PayHandler.GetPayInfo(ctx, in, out)
}
