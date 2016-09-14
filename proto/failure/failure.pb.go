// Code generated by protoc-gen-go.
// source: github.com/HailoOSS/platform/proto/failure/failure.proto
// DO NOT EDIT!

/*
Package com_HailoOSS_kernel_platform_failure is a generated protocol buffer package.

It is generated from these files:
	github.com/HailoOSS/platform/proto/failure/failure.proto

It has these top-level messages:
	Failure
*/
package com_HailoOSS_kernel_platform_failure

import proto "github.com/HailoOSS/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Failure struct {
	ServiceName      *string `protobuf:"bytes,1,req,name=serviceName" json:"serviceName,omitempty"`
	ServiceVersion   *uint64 `protobuf:"varint,2,req,name=serviceVersion" json:"serviceVersion,omitempty"`
	InstanceId       *string `protobuf:"bytes,3,req,name=instanceId" json:"instanceId,omitempty"`
	AzName           *string `protobuf:"bytes,4,req,name=azName" json:"azName,omitempty"`
	Hostname         *string `protobuf:"bytes,5,req,name=hostname" json:"hostname,omitempty"`
	Timestamp        *int64  `protobuf:"varint,6,req,name=timestamp" json:"timestamp,omitempty"`
	Uptime           *int64  `protobuf:"varint,7,opt,name=uptime" json:"uptime,omitempty"`
	Type             *string `protobuf:"bytes,8,req,name=type" json:"type,omitempty"`
	Reason           *string `protobuf:"bytes,9,req,name=reason" json:"reason,omitempty"`
	Stack            *string `protobuf:"bytes,10,opt,name=stack" json:"stack,omitempty"`
	Info             *string `protobuf:"bytes,11,opt,name=info" json:"info,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Failure) Reset()         { *m = Failure{} }
func (m *Failure) String() string { return proto.CompactTextString(m) }
func (*Failure) ProtoMessage()    {}

func (m *Failure) GetServiceName() string {
	if m != nil && m.ServiceName != nil {
		return *m.ServiceName
	}
	return ""
}

func (m *Failure) GetServiceVersion() uint64 {
	if m != nil && m.ServiceVersion != nil {
		return *m.ServiceVersion
	}
	return 0
}

func (m *Failure) GetInstanceId() string {
	if m != nil && m.InstanceId != nil {
		return *m.InstanceId
	}
	return ""
}

func (m *Failure) GetAzName() string {
	if m != nil && m.AzName != nil {
		return *m.AzName
	}
	return ""
}

func (m *Failure) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *Failure) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *Failure) GetUptime() int64 {
	if m != nil && m.Uptime != nil {
		return *m.Uptime
	}
	return 0
}

func (m *Failure) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Failure) GetReason() string {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return ""
}

func (m *Failure) GetStack() string {
	if m != nil && m.Stack != nil {
		return *m.Stack
	}
	return ""
}

func (m *Failure) GetInfo() string {
	if m != nil && m.Info != nil {
		return *m.Info
	}
	return ""
}

func init() {
}
