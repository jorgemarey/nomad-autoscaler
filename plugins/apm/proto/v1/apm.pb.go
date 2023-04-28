// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: plugins/apm/proto/v1/apm.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	v1 "github.com/hashicorp/nomad-autoscaler/plugins/shared/proto/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     string        `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	TimeRange *v1.TimeRange `protobuf:"bytes,2,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_apm_proto_v1_apm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_apm_proto_v1_apm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_plugins_apm_proto_v1_apm_proto_rawDescGZIP(), []int{0}
}

func (x *QueryRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *QueryRequest) GetTimeRange() *v1.TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

type QueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimestampedMetric []*v1.TimestampedMetric `protobuf:"bytes,1,rep,name=timestamped_metric,json=timestampedMetric,proto3" json:"timestamped_metric,omitempty"`
}

func (x *QueryResponse) Reset() {
	*x = QueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_apm_proto_v1_apm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResponse) ProtoMessage() {}

func (x *QueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_apm_proto_v1_apm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResponse.ProtoReflect.Descriptor instead.
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return file_plugins_apm_proto_v1_apm_proto_rawDescGZIP(), []int{1}
}

func (x *QueryResponse) GetTimestampedMetric() []*v1.TimestampedMetric {
	if x != nil {
		return x.TimestampedMetric
	}
	return nil
}

type QueryMultipleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     string        `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	TimeRange *v1.TimeRange `protobuf:"bytes,2,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
}

func (x *QueryMultipleRequest) Reset() {
	*x = QueryMultipleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_apm_proto_v1_apm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMultipleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMultipleRequest) ProtoMessage() {}

func (x *QueryMultipleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_apm_proto_v1_apm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMultipleRequest.ProtoReflect.Descriptor instead.
func (*QueryMultipleRequest) Descriptor() ([]byte, []int) {
	return file_plugins_apm_proto_v1_apm_proto_rawDescGZIP(), []int{2}
}

func (x *QueryMultipleRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *QueryMultipleRequest) GetTimeRange() *v1.TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

type QueryMultipleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimestampedMetric []*QueryResponse `protobuf:"bytes,1,rep,name=timestamped_metric,json=timestampedMetric,proto3" json:"timestamped_metric,omitempty"`
}

func (x *QueryMultipleResponse) Reset() {
	*x = QueryMultipleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_apm_proto_v1_apm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMultipleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMultipleResponse) ProtoMessage() {}

func (x *QueryMultipleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_apm_proto_v1_apm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMultipleResponse.ProtoReflect.Descriptor instead.
func (*QueryMultipleResponse) Descriptor() ([]byte, []int) {
	return file_plugins_apm_proto_v1_apm_proto_rawDescGZIP(), []int{3}
}

func (x *QueryMultipleResponse) GetTimestampedMetric() []*QueryResponse {
	if x != nil {
		return x.TimestampedMetric
	}
	return nil
}

var File_plugins_apm_proto_v1_apm_proto protoreflect.FileDescriptor

var file_plugins_apm_proto_v1_apm_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x61,
	0x64, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x24, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x0c, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x5c, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e,
	0x6e, 0x6f, 0x6d, 0x61, 0x64, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x85, 0x01,
	0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x74, 0x0a, 0x12, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x65, 0x64, 0x5f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x68, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x61, 0x64, 0x5f, 0x61, 0x75,
	0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x52, 0x11, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x65, 0x64, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x22, 0x8a, 0x01, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x5c, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x61, 0x64, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6d, 0x0a, 0x12,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x61, 0x64, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x61, 0x70,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x11, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x32, 0xc0, 0x02, 0x0a, 0x10,
	0x41, 0x50, 0x4d, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x88, 0x01, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3d, 0x2e, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x61, 0x64, 0x5f, 0x61, 0x75, 0x74,
	0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e,
	0x61, 0x70, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x61, 0x64, 0x5f, 0x61, 0x75, 0x74, 0x6f,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x61,
	0x70, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa0, 0x01, 0x0a, 0x0d,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x45, 0x2e,
	0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x61, 0x64, 0x5f,
	0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x46, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x2e, 0x6e, 0x6f, 0x6d, 0x61, 0x64, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07,
	0x5a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugins_apm_proto_v1_apm_proto_rawDescOnce sync.Once
	file_plugins_apm_proto_v1_apm_proto_rawDescData = file_plugins_apm_proto_v1_apm_proto_rawDesc
)

func file_plugins_apm_proto_v1_apm_proto_rawDescGZIP() []byte {
	file_plugins_apm_proto_v1_apm_proto_rawDescOnce.Do(func() {
		file_plugins_apm_proto_v1_apm_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugins_apm_proto_v1_apm_proto_rawDescData)
	})
	return file_plugins_apm_proto_v1_apm_proto_rawDescData
}

var file_plugins_apm_proto_v1_apm_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_plugins_apm_proto_v1_apm_proto_goTypes = []interface{}{
	(*QueryRequest)(nil),          // 0: hashicorp.nomad_autoscaler.plugins.apm.proto.v1.QueryRequest
	(*QueryResponse)(nil),         // 1: hashicorp.nomad_autoscaler.plugins.apm.proto.v1.QueryResponse
	(*QueryMultipleRequest)(nil),  // 2: hashicorp.nomad_autoscaler.plugins.apm.proto.v1.QueryMultipleRequest
	(*QueryMultipleResponse)(nil), // 3: hashicorp.nomad_autoscaler.plugins.apm.proto.v1.QueryMultipleResponse
	(*v1.TimeRange)(nil),          // 4: hashicorp.nomad_autoscaler.plugins.shared.proto.v1.TimeRange
	(*v1.TimestampedMetric)(nil),  // 5: hashicorp.nomad_autoscaler.plugins.shared.proto.v1.TimestampedMetric
}
var file_plugins_apm_proto_v1_apm_proto_depIdxs = []int32{
	4, // 0: hashicorp.nomad_autoscaler.plugins.apm.proto.v1.QueryRequest.time_range:type_name -> hashicorp.nomad_autoscaler.plugins.shared.proto.v1.TimeRange
	5, // 1: hashicorp.nomad_autoscaler.plugins.apm.proto.v1.QueryResponse.timestamped_metric:type_name -> hashicorp.nomad_autoscaler.plugins.shared.proto.v1.TimestampedMetric
	4, // 2: hashicorp.nomad_autoscaler.plugins.apm.proto.v1.QueryMultipleRequest.time_range:type_name -> hashicorp.nomad_autoscaler.plugins.shared.proto.v1.TimeRange
	1, // 3: hashicorp.nomad_autoscaler.plugins.apm.proto.v1.QueryMultipleResponse.timestamped_metric:type_name -> hashicorp.nomad_autoscaler.plugins.apm.proto.v1.QueryResponse
	0, // 4: hashicorp.nomad_autoscaler.plugins.apm.proto.v1.APMPluginService.Query:input_type -> hashicorp.nomad_autoscaler.plugins.apm.proto.v1.QueryRequest
	2, // 5: hashicorp.nomad_autoscaler.plugins.apm.proto.v1.APMPluginService.QueryMultiple:input_type -> hashicorp.nomad_autoscaler.plugins.apm.proto.v1.QueryMultipleRequest
	1, // 6: hashicorp.nomad_autoscaler.plugins.apm.proto.v1.APMPluginService.Query:output_type -> hashicorp.nomad_autoscaler.plugins.apm.proto.v1.QueryResponse
	3, // 7: hashicorp.nomad_autoscaler.plugins.apm.proto.v1.APMPluginService.QueryMultiple:output_type -> hashicorp.nomad_autoscaler.plugins.apm.proto.v1.QueryMultipleResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_plugins_apm_proto_v1_apm_proto_init() }
func file_plugins_apm_proto_v1_apm_proto_init() {
	if File_plugins_apm_proto_v1_apm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugins_apm_proto_v1_apm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_plugins_apm_proto_v1_apm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_plugins_apm_proto_v1_apm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMultipleRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_plugins_apm_proto_v1_apm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMultipleResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_plugins_apm_proto_v1_apm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plugins_apm_proto_v1_apm_proto_goTypes,
		DependencyIndexes: file_plugins_apm_proto_v1_apm_proto_depIdxs,
		MessageInfos:      file_plugins_apm_proto_v1_apm_proto_msgTypes,
	}.Build()
	File_plugins_apm_proto_v1_apm_proto = out.File
	file_plugins_apm_proto_v1_apm_proto_rawDesc = nil
	file_plugins_apm_proto_v1_apm_proto_goTypes = nil
	file_plugins_apm_proto_v1_apm_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// APMPluginServiceClient is the client API for APMPluginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APMPluginServiceClient interface {
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	QueryMultiple(ctx context.Context, in *QueryMultipleRequest, opts ...grpc.CallOption) (*QueryMultipleResponse, error)
}

type aPMPluginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAPMPluginServiceClient(cc grpc.ClientConnInterface) APMPluginServiceClient {
	return &aPMPluginServiceClient{cc}
}

func (c *aPMPluginServiceClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.nomad_autoscaler.plugins.apm.proto.v1.APMPluginService/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPMPluginServiceClient) QueryMultiple(ctx context.Context, in *QueryMultipleRequest, opts ...grpc.CallOption) (*QueryMultipleResponse, error) {
	out := new(QueryMultipleResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.nomad_autoscaler.plugins.apm.proto.v1.APMPluginService/QueryMultiple", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APMPluginServiceServer is the server API for APMPluginService service.
type APMPluginServiceServer interface {
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	QueryMultiple(context.Context, *QueryMultipleRequest) (*QueryMultipleResponse, error)
}

// UnimplementedAPMPluginServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAPMPluginServiceServer struct {
}

func (*UnimplementedAPMPluginServiceServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (*UnimplementedAPMPluginServiceServer) QueryMultiple(context.Context, *QueryMultipleRequest) (*QueryMultipleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryMultiple not implemented")
}

func RegisterAPMPluginServiceServer(s *grpc.Server, srv APMPluginServiceServer) {
	s.RegisterService(&_APMPluginService_serviceDesc, srv)
}

func _APMPluginService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APMPluginServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.nomad_autoscaler.plugins.apm.proto.v1.APMPluginService/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APMPluginServiceServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APMPluginService_QueryMultiple_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMultipleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APMPluginServiceServer).QueryMultiple(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.nomad_autoscaler.plugins.apm.proto.v1.APMPluginService/QueryMultiple",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APMPluginServiceServer).QueryMultiple(ctx, req.(*QueryMultipleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _APMPluginService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hashicorp.nomad_autoscaler.plugins.apm.proto.v1.APMPluginService",
	HandlerType: (*APMPluginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _APMPluginService_Query_Handler,
		},
		{
			MethodName: "QueryMultiple",
			Handler:    _APMPluginService_QueryMultiple_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugins/apm/proto/v1/apm.proto",
}
