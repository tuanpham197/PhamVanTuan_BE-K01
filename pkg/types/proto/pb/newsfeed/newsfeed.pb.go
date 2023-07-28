// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.3
// source: newsfeed.proto

package newsfeed

import (
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

type NewsfeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *NewsfeedRequest) Reset() {
	*x = NewsfeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_newsfeed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsfeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsfeedRequest) ProtoMessage() {}

func (x *NewsfeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_newsfeed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsfeedRequest.ProtoReflect.Descriptor instead.
func (*NewsfeedRequest) Descriptor() ([]byte, []int) {
	return file_newsfeed_proto_rawDescGZIP(), []int{0}
}

func (x *NewsfeedRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type NewsfeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostIds []int64 `protobuf:"varint,1,rep,packed,name=post_ids,json=postIds,proto3" json:"post_ids,omitempty"`
}

func (x *NewsfeedResponse) Reset() {
	*x = NewsfeedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_newsfeed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsfeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsfeedResponse) ProtoMessage() {}

func (x *NewsfeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_newsfeed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsfeedResponse.ProtoReflect.Descriptor instead.
func (*NewsfeedResponse) Descriptor() ([]byte, []int) {
	return file_newsfeed_proto_rawDescGZIP(), []int{1}
}

func (x *NewsfeedResponse) GetPostIds() []int64 {
	if x != nil {
		return x.PostIds
	}
	return nil
}

var File_newsfeed_proto protoreflect.FileDescriptor

var file_newsfeed_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x73, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6e, 0x65, 0x77, 0x73, 0x66, 0x65, 0x65, 0x64, 0x22, 0x2a, 0x0a, 0x0f, 0x4e, 0x65,
	0x77, 0x73, 0x66, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x10, 0x4e, 0x65, 0x77, 0x73, 0x66, 0x65,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x70, 0x6f,
	0x73, 0x74, 0x49, 0x64, 0x73, 0x32, 0x4f, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x73, 0x66, 0x65, 0x65,
	0x64, 0x12, 0x43, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x73, 0x66, 0x65, 0x65, 0x64, 0x12, 0x19, 0x2e,
	0x6e, 0x65, 0x77, 0x73, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x66, 0x65, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x66,
	0x65, 0x65, 0x64, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x66, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x6e, 0x65, 0x77, 0x73,
	0x66, 0x65, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_newsfeed_proto_rawDescOnce sync.Once
	file_newsfeed_proto_rawDescData = file_newsfeed_proto_rawDesc
)

func file_newsfeed_proto_rawDescGZIP() []byte {
	file_newsfeed_proto_rawDescOnce.Do(func() {
		file_newsfeed_proto_rawDescData = protoimpl.X.CompressGZIP(file_newsfeed_proto_rawDescData)
	})
	return file_newsfeed_proto_rawDescData
}

var file_newsfeed_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_newsfeed_proto_goTypes = []interface{}{
	(*NewsfeedRequest)(nil),  // 0: newsfeed.NewsfeedRequest
	(*NewsfeedResponse)(nil), // 1: newsfeed.NewsfeedResponse
}
var file_newsfeed_proto_depIdxs = []int32{
	0, // 0: newsfeed.Newsfeed.Newsfeed:input_type -> newsfeed.NewsfeedRequest
	1, // 1: newsfeed.Newsfeed.Newsfeed:output_type -> newsfeed.NewsfeedResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_newsfeed_proto_init() }
func file_newsfeed_proto_init() {
	if File_newsfeed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_newsfeed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsfeedRequest); i {
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
		file_newsfeed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsfeedResponse); i {
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
			RawDescriptor: file_newsfeed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_newsfeed_proto_goTypes,
		DependencyIndexes: file_newsfeed_proto_depIdxs,
		MessageInfos:      file_newsfeed_proto_msgTypes,
	}.Build()
	File_newsfeed_proto = out.File
	file_newsfeed_proto_rawDesc = nil
	file_newsfeed_proto_goTypes = nil
	file_newsfeed_proto_depIdxs = nil
}
