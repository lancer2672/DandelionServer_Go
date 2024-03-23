// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: request/vote.proto

package request

import (
	_ "github.com/lancer2672/DandelionServer_Go/pb/model"
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

type GetVotesByUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Vote `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetVotesByUserResponse) Reset() {
	*x = GetVotesByUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_vote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVotesByUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVotesByUserResponse) ProtoMessage() {}

func (x *GetVotesByUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_request_vote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVotesByUserResponse.ProtoReflect.Descriptor instead.
func (*GetVotesByUserResponse) Descriptor() ([]byte, []int) {
	return file_request_vote_proto_rawDescGZIP(), []int{0}
}

func (x *GetVotesByUserResponse) GetData() []*Vote {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetVotesByUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetVotesByUserRequest) Reset() {
	*x = GetVotesByUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_vote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVotesByUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVotesByUserRequest) ProtoMessage() {}

func (x *GetVotesByUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_vote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVotesByUserRequest.ProtoReflect.Descriptor instead.
func (*GetVotesByUserRequest) Descriptor() ([]byte, []int) {
	return file_request_vote_proto_rawDescGZIP(), []int{1}
}

func (x *GetVotesByUserRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Vote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId  int32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MovieId int32 `protobuf:"varint,3,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
	Stars   int32 `protobuf:"varint,4,opt,name=stars,proto3" json:"stars,omitempty"`
}

func (x *Vote) Reset() {
	*x = Vote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_vote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vote) ProtoMessage() {}

func (x *Vote) ProtoReflect() protoreflect.Message {
	mi := &file_request_vote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vote.ProtoReflect.Descriptor instead.
func (*Vote) Descriptor() ([]byte, []int) {
	return file_request_vote_proto_rawDescGZIP(), []int{2}
}

func (x *Vote) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Vote) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Vote) GetMovieId() int32 {
	if x != nil {
		return x.MovieId
	}
	return 0
}

func (x *Vote) GetStars() int32 {
	if x != nil {
		return x.Stars
	}
	return 0
}

var File_request_vote_proto protoreflect.FileDescriptor

var file_request_vote_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x67, 0x65, 0x6e, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x30, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x60, 0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x73, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x32, 0x36, 0x37, 0x32,
	0x2f, 0x44, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x47, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_request_vote_proto_rawDescOnce sync.Once
	file_request_vote_proto_rawDescData = file_request_vote_proto_rawDesc
)

func file_request_vote_proto_rawDescGZIP() []byte {
	file_request_vote_proto_rawDescOnce.Do(func() {
		file_request_vote_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_vote_proto_rawDescData)
	})
	return file_request_vote_proto_rawDescData
}

var file_request_vote_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_request_vote_proto_goTypes = []interface{}{
	(*GetVotesByUserResponse)(nil), // 0: pb.GetVotesByUserResponse
	(*GetVotesByUserRequest)(nil),  // 1: pb.GetVotesByUserRequest
	(*Vote)(nil),                   // 2: pb.Vote
}
var file_request_vote_proto_depIdxs = []int32{
	2, // 0: pb.GetVotesByUserResponse.data:type_name -> pb.Vote
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_request_vote_proto_init() }
func file_request_vote_proto_init() {
	if File_request_vote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_request_vote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVotesByUserResponse); i {
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
		file_request_vote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVotesByUserRequest); i {
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
		file_request_vote_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vote); i {
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
			RawDescriptor: file_request_vote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_request_vote_proto_goTypes,
		DependencyIndexes: file_request_vote_proto_depIdxs,
		MessageInfos:      file_request_vote_proto_msgTypes,
	}.Build()
	File_request_vote_proto = out.File
	file_request_vote_proto_rawDesc = nil
	file_request_vote_proto_goTypes = nil
	file_request_vote_proto_depIdxs = nil
}
