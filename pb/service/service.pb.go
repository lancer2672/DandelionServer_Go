// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: service/service.proto

package service

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	model "github.com/lancer2672/DandelionServer_Go/pb/model"
	request "github.com/lancer2672/DandelionServer_Go/pb/request"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_service_proto protoreflect.FileDescriptor

var file_service_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1b, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2d, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2f, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x6f, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9c, 0x08, 0x0a, 0x09, 0x44, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x69,
	0x6f, 0x6e, 0x12, 0x66, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x62, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1a, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e,
	0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x5b,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x12,
	0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x65, 0x6e, 0x72,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x12, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x22, 0x07, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x3a, 0x01,
	0x2a, 0x12, 0x55, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12,
	0x07, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6f, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x42, 0x79, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x1b,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x42, 0x79, 0x47,
	0x65, 0x6e, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x42, 0x79, 0x47, 0x65, 0x6e, 0x72,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1a, 0x12, 0x18, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x72, 0x65,
	0x2f, 0x7b, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x69, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x42, 0x79, 0x53, 0x65, 0x72, 0x69, 0x65, 0x12,
	0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x42, 0x79,
	0x53, 0x65, 0x72, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x42, 0x79, 0x53, 0x65, 0x72,
	0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x14, 0x12, 0x12, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x69,
	0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x62, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63,
	0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x63, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x59, 0x0a, 0x0c, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2f, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x6c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x74, 0x65,
	0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x56, 0x6f, 0x74, 0x65, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x73,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x76,
	0x6f, 0x74, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x7d, 0x42, 0xa0, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x32, 0x36, 0x37, 0x32, 0x2f, 0x44, 0x61,
	0x6e, 0x64, 0x65, 0x6c, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x47, 0x6f,
	0x2f, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x92, 0x41, 0x68, 0x12, 0x66,
	0x0a, 0x15, 0x44, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x69, 0x6f, 0x6e, 0x20, 0x47, 0x52, 0x50, 0x43,
	0x20, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x48, 0x0a, 0x14, 0x44, 0x61, 0x6e, 0x64, 0x65,
	0x6c, 0x69, 0x6f, 0x6e, 0x5f, 0x47, 0x6f, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x30, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x32, 0x36, 0x37, 0x32, 0x2f, 0x44,
	0x61, 0x6e, 0x64, 0x65, 0x6c, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x47,
	0x6f, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_service_proto_goTypes = []interface{}{
	(*request.CreateMovieHistoryRequest)(nil), // 0: pb.CreateMovieHistoryRequest
	(*request.GetMovieHistoryRequest)(nil),    // 1: pb.GetMovieHistoryRequest
	(*request.GetListGenresRequest)(nil),      // 2: pb.GetListGenresRequest
	(*request.CreateMovieRequest)(nil),        // 3: pb.CreateMovieRequest
	(*request.GetListMoviesRequest)(nil),      // 4: pb.GetListMoviesRequest
	(*request.GetMovieRequest)(nil),           // 5: pb.GetMovieRequest
	(*request.GetMoviesByGenreRequest)(nil),   // 6: pb.GetMoviesByGenreRequest
	(*request.GetMoviesBySerieRequest)(nil),   // 7: pb.GetMoviesBySerieRequest
	(*request.GetRecentMoviesRequest)(nil),    // 8: pb.GetRecentMoviesRequest
	(*request.SearchMoviesRequest)(nil),       // 9: pb.SearchMoviesRequest
	(*request.GetVotesByUserRequest)(nil),     // 10: pb.GetVotesByUserRequest
	(*empty.Empty)(nil),                       // 11: google.protobuf.Empty
	(*request.GetMovieHistoryResponse)(nil),   // 12: pb.GetMovieHistoryResponse
	(*request.GetListGenresResponse)(nil),     // 13: pb.GetListGenresResponse
	(*model.Movie)(nil),                       // 14: pb.Movie
	(*request.GetListMoviesResponse)(nil),     // 15: pb.GetListMoviesResponse
	(*request.GetMoviesByGenreResponse)(nil),  // 16: pb.GetMoviesByGenreResponse
	(*request.GetMoviesBySerieResponse)(nil),  // 17: pb.GetMoviesBySerieResponse
	(*request.GetRecentMoviesResponse)(nil),   // 18: pb.GetRecentMoviesResponse
	(*request.SearchMoviesResponse)(nil),      // 19: pb.SearchMoviesResponse
	(*request.GetVotesByUserResponse)(nil),    // 20: pb.GetVotesByUserResponse
}
var file_service_service_proto_depIdxs = []int32{
	0,  // 0: pb.Dandelion.CreateMovieHistory:input_type -> pb.CreateMovieHistoryRequest
	1,  // 1: pb.Dandelion.GetMovieHistory:input_type -> pb.GetMovieHistoryRequest
	2,  // 2: pb.Dandelion.GetListGenres:input_type -> pb.GetListGenresRequest
	3,  // 3: pb.Dandelion.CreateMovie:input_type -> pb.CreateMovieRequest
	4,  // 4: pb.Dandelion.GetListMovies:input_type -> pb.GetListMoviesRequest
	5,  // 5: pb.Dandelion.GetMovie:input_type -> pb.GetMovieRequest
	6,  // 6: pb.Dandelion.GetMoviesByGenre:input_type -> pb.GetMoviesByGenreRequest
	7,  // 7: pb.Dandelion.GetMoviesBySerie:input_type -> pb.GetMoviesBySerieRequest
	8,  // 8: pb.Dandelion.GetRecentMovies:input_type -> pb.GetRecentMoviesRequest
	9,  // 9: pb.Dandelion.SearchMovies:input_type -> pb.SearchMoviesRequest
	10, // 10: pb.Dandelion.GetVotesByUser:input_type -> pb.GetVotesByUserRequest
	11, // 11: pb.Dandelion.CreateMovieHistory:output_type -> google.protobuf.Empty
	12, // 12: pb.Dandelion.GetMovieHistory:output_type -> pb.GetMovieHistoryResponse
	13, // 13: pb.Dandelion.GetListGenres:output_type -> pb.GetListGenresResponse
	14, // 14: pb.Dandelion.CreateMovie:output_type -> pb.Movie
	15, // 15: pb.Dandelion.GetListMovies:output_type -> pb.GetListMoviesResponse
	14, // 16: pb.Dandelion.GetMovie:output_type -> pb.Movie
	16, // 17: pb.Dandelion.GetMoviesByGenre:output_type -> pb.GetMoviesByGenreResponse
	17, // 18: pb.Dandelion.GetMoviesBySerie:output_type -> pb.GetMoviesBySerieResponse
	18, // 19: pb.Dandelion.GetRecentMovies:output_type -> pb.GetRecentMoviesResponse
	19, // 20: pb.Dandelion.SearchMovies:output_type -> pb.SearchMoviesResponse
	20, // 21: pb.Dandelion.GetVotesByUser:output_type -> pb.GetVotesByUserResponse
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_service_proto_init() }
func file_service_service_proto_init() {
	if File_service_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_service_proto_goTypes,
		DependencyIndexes: file_service_service_proto_depIdxs,
	}.Build()
	File_service_service_proto = out.File
	file_service_service_proto_rawDesc = nil
	file_service_service_proto_goTypes = nil
	file_service_service_proto_depIdxs = nil
}
