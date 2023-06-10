// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.9.2
// source: authen_and_post.proto

package authen_and_post

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

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName     string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserPassword string `protobuf:"bytes,3,opt,name=user_password,json=userPassword,proto3" json:"user_password,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authen_and_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_authen_and_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_authen_and_post_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserInfo) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserInfo) GetUserPassword() string {
	if x != nil {
		return x.UserPassword
	}
	return ""
}

type UserDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName  string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Dob       int64  `protobuf:"varint,5,opt,name=dob,proto3" json:"dob,omitempty"`
	Email     string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UserDetailInfo) Reset() {
	*x = UserDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authen_and_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetailInfo) ProtoMessage() {}

func (x *UserDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_authen_and_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetailInfo.ProtoReflect.Descriptor instead.
func (*UserDetailInfo) Descriptor() ([]byte, []int) {
	return file_authen_and_post_proto_rawDescGZIP(), []int{1}
}

func (x *UserDetailInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserDetailInfo) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserDetailInfo) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserDetailInfo) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserDetailInfo) GetDob() int64 {
	if x != nil {
		return x.Dob
	}
	return 0
}

func (x *UserDetailInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type UserResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status uint32          `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Info   *UserDetailInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *UserResult) Reset() {
	*x = UserResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authen_and_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResult) ProtoMessage() {}

func (x *UserResult) ProtoReflect() protoreflect.Message {
	mi := &file_authen_and_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResult.ProtoReflect.Descriptor instead.
func (*UserResult) Descriptor() ([]byte, []int) {
	return file_authen_and_post_proto_rawDescGZIP(), []int{2}
}

func (x *UserResult) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UserResult) GetInfo() *UserDetailInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type UserFollower struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Followers []*UserInfo `protobuf:"bytes,1,rep,name=followers,proto3" json:"followers,omitempty"`
}

func (x *UserFollower) Reset() {
	*x = UserFollower{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authen_and_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFollower) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFollower) ProtoMessage() {}

func (x *UserFollower) ProtoReflect() protoreflect.Message {
	mi := &file_authen_and_post_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFollower.ProtoReflect.Descriptor instead.
func (*UserFollower) Descriptor() ([]byte, []int) {
	return file_authen_and_post_proto_rawDescGZIP(), []int{3}
}

func (x *UserFollower) GetFollowers() []*UserInfo {
	if x != nil {
		return x.Followers
	}
	return nil
}

type GetPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId int64 `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
}

func (x *GetPostRequest) Reset() {
	*x = GetPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authen_and_post_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostRequest) ProtoMessage() {}

func (x *GetPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authen_and_post_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostRequest.ProtoReflect.Descriptor instead.
func (*GetPostRequest) Descriptor() ([]byte, []int) {
	return file_authen_and_post_proto_rawDescGZIP(), []int{4}
}

func (x *GetPostRequest) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId           int64  `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	UserId           int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ContentText      string `protobuf:"bytes,3,opt,name=content_text,json=contentText,proto3" json:"content_text,omitempty"`
	ContentImagePath string `protobuf:"bytes,4,opt,name=content_image_path,json=contentImagePath,proto3" json:"content_image_path,omitempty"`
	Visible          bool   `protobuf:"varint,5,opt,name=visible,proto3" json:"visible,omitempty"`
	CreatedAt        int64  `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authen_and_post_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_authen_and_post_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_authen_and_post_proto_rawDescGZIP(), []int{5}
}

func (x *Post) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *Post) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Post) GetContentText() string {
	if x != nil {
		return x.ContentText
	}
	return ""
}

func (x *Post) GetContentImagePath() string {
	if x != nil {
		return x.ContentImagePath
	}
	return ""
}

func (x *Post) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

func (x *Post) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

var File_authen_and_post_proto protoreflect.FileDescriptor

var file_authen_and_post_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x5f,
	0x61, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x65, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0xaa, 0x01, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x64, 0x6f, 0x62, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x59, 0x0a, 0x0a,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x33, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x70, 0x6f,
	0x73, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x47, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73,
	0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0xc2, 0x01, 0x0a, 0x04,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x78, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62,
	0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x32, 0x9e, 0x03, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x41, 0x6e, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x53, 0x0a, 0x17, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x5f, 0x61, 0x6e, 0x64,
	0x5f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1b,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x74,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x4c, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1b, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x08, 0x45,
	0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x5f, 0x61,
	0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x22,
	0x00, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x5f, 0x61, 0x6e,
	0x64, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authen_and_post_proto_rawDescOnce sync.Once
	file_authen_and_post_proto_rawDescData = file_authen_and_post_proto_rawDesc
)

func file_authen_and_post_proto_rawDescGZIP() []byte {
	file_authen_and_post_proto_rawDescOnce.Do(func() {
		file_authen_and_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_authen_and_post_proto_rawDescData)
	})
	return file_authen_and_post_proto_rawDescData
}

var file_authen_and_post_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_authen_and_post_proto_goTypes = []interface{}{
	(*UserInfo)(nil),       // 0: authen_and_post.UserInfo
	(*UserDetailInfo)(nil), // 1: authen_and_post.UserDetailInfo
	(*UserResult)(nil),     // 2: authen_and_post.UserResult
	(*UserFollower)(nil),   // 3: authen_and_post.UserFollower
	(*GetPostRequest)(nil), // 4: authen_and_post.GetPostRequest
	(*Post)(nil),           // 5: authen_and_post.Post
}
var file_authen_and_post_proto_depIdxs = []int32{
	1, // 0: authen_and_post.UserResult.info:type_name -> authen_and_post.UserDetailInfo
	0, // 1: authen_and_post.UserFollower.followers:type_name -> authen_and_post.UserInfo
	0, // 2: authen_and_post.AuthenticateAndPost.CheckUserAuthentication:input_type -> authen_and_post.UserInfo
	1, // 3: authen_and_post.AuthenticateAndPost.CreateUser:input_type -> authen_and_post.UserDetailInfo
	1, // 4: authen_and_post.AuthenticateAndPost.EditUser:input_type -> authen_and_post.UserDetailInfo
	0, // 5: authen_and_post.AuthenticateAndPost.GetUserFollower:input_type -> authen_and_post.UserInfo
	4, // 6: authen_and_post.AuthenticateAndPost.GetPostDetail:input_type -> authen_and_post.GetPostRequest
	2, // 7: authen_and_post.AuthenticateAndPost.CheckUserAuthentication:output_type -> authen_and_post.UserResult
	2, // 8: authen_and_post.AuthenticateAndPost.CreateUser:output_type -> authen_and_post.UserResult
	2, // 9: authen_and_post.AuthenticateAndPost.EditUser:output_type -> authen_and_post.UserResult
	3, // 10: authen_and_post.AuthenticateAndPost.GetUserFollower:output_type -> authen_and_post.UserFollower
	5, // 11: authen_and_post.AuthenticateAndPost.GetPostDetail:output_type -> authen_and_post.Post
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_authen_and_post_proto_init() }
func file_authen_and_post_proto_init() {
	if File_authen_and_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authen_and_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_authen_and_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDetailInfo); i {
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
		file_authen_and_post_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResult); i {
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
		file_authen_and_post_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFollower); i {
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
		file_authen_and_post_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostRequest); i {
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
		file_authen_and_post_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
			RawDescriptor: file_authen_and_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authen_and_post_proto_goTypes,
		DependencyIndexes: file_authen_and_post_proto_depIdxs,
		MessageInfos:      file_authen_and_post_proto_msgTypes,
	}.Build()
	File_authen_and_post_proto = out.File
	file_authen_and_post_proto_rawDesc = nil
	file_authen_and_post_proto_goTypes = nil
	file_authen_and_post_proto_depIdxs = nil
}
