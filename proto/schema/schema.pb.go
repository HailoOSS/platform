// Code generated by protoc-gen-go.
// source: github.com/HailoOSS/platform/proto/schema/schema.proto
// DO NOT EDIT!

/*
Package com_HailoOSS_kernel_platform_schema is a generated protocol buffer package.

It is generated from these files:
	github.com/HailoOSS/platform/proto/schema/schema.proto

It has these top-level messages:
	Request
	Response
*/
package com_HailoOSS_kernel_platform_schema

import proto "github.com/HailoOSS/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Request struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

type Response struct {
	Schema           *string `protobuf:"bytes,1,req,name=schema" json:"schema,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetSchema() string {
	if m != nil && m.Schema != nil {
		return *m.Schema
	}
	return ""
}

func init() {
}
