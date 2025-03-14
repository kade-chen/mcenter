// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.28.3
// source: mcenter/apps/user/pb/rpc.proto

package user

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

type DESCRIBE_BY int32

const (
	// 通过UserId查询用户
	DESCRIBE_BY_USER_ID DESCRIBE_BY = 0
	// 通过Username查询用户
	DESCRIBE_BY_USER_NAME DESCRIBE_BY = 1
	// 通过飞书UserId查询用户
	DESCRIBE_BY_FEISHU_USER_ID DESCRIBE_BY = 2
)

// Enum value maps for DESCRIBE_BY.
var (
	DESCRIBE_BY_name = map[int32]string{
		0: "USER_ID",
		1: "USER_NAME",
		2: "FEISHU_USER_ID",
	}
	DESCRIBE_BY_value = map[string]int32{
		"USER_ID":        0,
		"USER_NAME":      1,
		"FEISHU_USER_ID": 2,
	}
)

func (x DESCRIBE_BY) Enum() *DESCRIBE_BY {
	p := new(DESCRIBE_BY)
	*p = x
	return p
}

func (x DESCRIBE_BY) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DESCRIBE_BY) Descriptor() protoreflect.EnumDescriptor {
	return file_mcenter_apps_user_pb_rpc_proto_enumTypes[0].Descriptor()
}

func (DESCRIBE_BY) Type() protoreflect.EnumType {
	return &file_mcenter_apps_user_pb_rpc_proto_enumTypes[0]
}

func (x DESCRIBE_BY) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DESCRIBE_BY.Descriptor instead.
func (DESCRIBE_BY) EnumDescriptor() ([]byte, []int) {
	return file_mcenter_apps_user_pb_rpc_proto_rawDescGZIP(), []int{0}
}

// QueryUserRequest 获取子账号列表
type QueryUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: json:"page"
	Page *PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 用户所属Domain
	// @gotags: json:"domain" validate:"required"
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain" validate:"required"`
	// 账号提供方
	// @gotags: json:"provider"
	Provider *PROVIDER `protobuf:"varint,3,opt,name=provider,proto3,enum=gokade.mcenter.user.PROVIDER,oneof" json:"provider"`
	// 用户类型
	// @gotags: json:"type"
	Type *TYPE `protobuf:"varint,4,opt,name=type,proto3,enum=gokade.mcenter.user.TYPE,oneof" json:"type"`
	// 通过Id
	// @gotags: json:"user_ids"
	UserIds []string `protobuf:"bytes,5,rep,name=user_ids,json=userIds,proto3" json:"user_ids"`
	// 额外需要查询出来的用户
	// @gotags: json:"extra_user_ids"
	ExtraUserIds []string `protobuf:"bytes,7,rep,name=extra_user_ids,json=extraUserIds,proto3" json:"extra_user_ids"`
	// 根据标签过滤用户
	// @gotags: json:"labels"
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 是否获取数据
	// @gotags: json:"skip_items"
	SkipItems bool `protobuf:"varint,8,opt,name=skip_items,json=skipItems,proto3" json:"skip_items"`
	// 关键字查询
	// @gotags: json:"keywords"
	Keywords string `protobuf:"bytes,9,opt,name=keywords,proto3" json:"keywords"`
}

func (x *QueryUserRequest) Reset() {
	*x = QueryUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_user_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUserRequest) ProtoMessage() {}

func (x *QueryUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_user_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUserRequest.ProtoReflect.Descriptor instead.
func (*QueryUserRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_user_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *QueryUserRequest) GetPage() *PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryUserRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *QueryUserRequest) GetProvider() PROVIDER {
	if x != nil && x.Provider != nil {
		return *x.Provider
	}
	return PROVIDER_LOCAL
}

func (x *QueryUserRequest) GetType() TYPE {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return TYPE_SUB
}

func (x *QueryUserRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *QueryUserRequest) GetExtraUserIds() []string {
	if x != nil {
		return x.ExtraUserIds
	}
	return nil
}

func (x *QueryUserRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *QueryUserRequest) GetSkipItems() bool {
	if x != nil {
		return x.SkipItems
	}
	return false
}

func (x *QueryUserRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

// DescribeUserRequest 查询用户详情
type DescribeUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 查询方式
	// @gotags: json:"describe_by"
	DescribeBy DESCRIBE_BY `protobuf:"varint,1,opt,name=describe_by,json=describeBy,proto3,enum=gokade.mcenter.user.DESCRIBE_BY" json:"describe_by"`
	// 用户账号id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 用户账号
	// @gotags: json:"username"
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username"`
}

func (x *DescribeUserRequest) Reset() {
	*x = DescribeUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_user_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeUserRequest) ProtoMessage() {}

func (x *DescribeUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_user_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeUserRequest.ProtoReflect.Descriptor instead.
func (*DescribeUserRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_user_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeUserRequest) GetDescribeBy() DESCRIBE_BY {
	if x != nil {
		return x.DescribeBy
	}
	return DESCRIBE_BY_USER_ID
}

func (x *DescribeUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DescribeUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户账号id
	// @gotags: json:"user_ids" validate:"required,lte=60"
	UserIds []string `protobuf:"bytes,2,rep,name=user_ids,json=userIds,proto3" json:"user_ids" validate:"required,lte=60"`
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_user_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_user_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_user_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteUserRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

var File_mcenter_apps_user_pb_rpc_proto protoreflect.FileDescriptor

var file_mcenter_apps_user_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x67, 0x6f, 0x6b, 0x61, 0x64, 0x65, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x70, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x03, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6b,
	0x61, 0x64, 0x65, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x3e, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x67,
	0x6f, 0x6b, 0x61, 0x64, 0x65, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x48, 0x00, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6b, 0x61, 0x64,
	0x65, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x54,
	0x59, 0x50, 0x45, 0x48, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x78, 0x74, 0x72, 0x61, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12,
	0x49, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x67, 0x6f, 0x6b, 0x61, 0x64, 0x65, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6b,
	0x69, 0x70, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x73, 0x6b, 0x69, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6b, 0x61, 0x64, 0x65, 0x2e, 0x6d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49,
	0x42, 0x45, 0x5f, 0x42, 0x59, 0x52, 0x0a, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x2a, 0x3d, 0x0a,
	0x0b, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x42, 0x59, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x45, 0x49, 0x53,
	0x48, 0x55, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x10, 0x02, 0x32, 0xab, 0x01, 0x0a,
	0x03, 0x52, 0x50, 0x43, 0x12, 0x4f, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x25, 0x2e, 0x67, 0x6f, 0x6b, 0x61, 0x64, 0x65, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6b, 0x61, 0x64, 0x65,
	0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x74, 0x12, 0x53, 0x0a, 0x0c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x28, 0x2e, 0x67, 0x6f, 0x6b, 0x61, 0x64, 0x65, 0x2e, 0x6d,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x67, 0x6f, 0x6b, 0x61, 0x64, 0x65, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x64, 0x65, 0x2d, 0x63, 0x68,
	0x65, 0x6e, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mcenter_apps_user_pb_rpc_proto_rawDescOnce sync.Once
	file_mcenter_apps_user_pb_rpc_proto_rawDescData = file_mcenter_apps_user_pb_rpc_proto_rawDesc
)

func file_mcenter_apps_user_pb_rpc_proto_rawDescGZIP() []byte {
	file_mcenter_apps_user_pb_rpc_proto_rawDescOnce.Do(func() {
		file_mcenter_apps_user_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_mcenter_apps_user_pb_rpc_proto_rawDescData)
	})
	return file_mcenter_apps_user_pb_rpc_proto_rawDescData
}

var file_mcenter_apps_user_pb_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mcenter_apps_user_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mcenter_apps_user_pb_rpc_proto_goTypes = []interface{}{
	(DESCRIBE_BY)(0),            // 0: gokade.mcenter.user.DESCRIBE_BY
	(*QueryUserRequest)(nil),    // 1: gokade.mcenter.user.QueryUserRequest
	(*DescribeUserRequest)(nil), // 2: gokade.mcenter.user.DescribeUserRequest
	(*DeleteUserRequest)(nil),   // 3: gokade.mcenter.user.DeleteUserRequest
	nil,                         // 4: gokade.mcenter.user.QueryUserRequest.LabelsEntry
	(*PageRequest)(nil),         // 5: gokade.mcenter.user.PageRequest
	(PROVIDER)(0),               // 6: gokade.mcenter.user.PROVIDER
	(TYPE)(0),                   // 7: gokade.mcenter.user.TYPE
	(*UserSet)(nil),             // 8: gokade.mcenter.user.UserSet
	(*User)(nil),                // 9: gokade.mcenter.user.User
}
var file_mcenter_apps_user_pb_rpc_proto_depIdxs = []int32{
	5, // 0: gokade.mcenter.user.QueryUserRequest.page:type_name -> gokade.mcenter.user.PageRequest
	6, // 1: gokade.mcenter.user.QueryUserRequest.provider:type_name -> gokade.mcenter.user.PROVIDER
	7, // 2: gokade.mcenter.user.QueryUserRequest.type:type_name -> gokade.mcenter.user.TYPE
	4, // 3: gokade.mcenter.user.QueryUserRequest.labels:type_name -> gokade.mcenter.user.QueryUserRequest.LabelsEntry
	0, // 4: gokade.mcenter.user.DescribeUserRequest.describe_by:type_name -> gokade.mcenter.user.DESCRIBE_BY
	1, // 5: gokade.mcenter.user.RPC.ListUser:input_type -> gokade.mcenter.user.QueryUserRequest
	2, // 6: gokade.mcenter.user.RPC.DescribeUser:input_type -> gokade.mcenter.user.DescribeUserRequest
	8, // 7: gokade.mcenter.user.RPC.ListUser:output_type -> gokade.mcenter.user.UserSet
	9, // 8: gokade.mcenter.user.RPC.DescribeUser:output_type -> gokade.mcenter.user.User
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_mcenter_apps_user_pb_rpc_proto_init() }
func file_mcenter_apps_user_pb_rpc_proto_init() {
	if File_mcenter_apps_user_pb_rpc_proto != nil {
		return
	}
	file_mcenter_apps_user_pb_user_proto_init()
	file_mcenter_apps_user_pb_page_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mcenter_apps_user_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUserRequest); i {
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
		file_mcenter_apps_user_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeUserRequest); i {
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
		file_mcenter_apps_user_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserRequest); i {
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
	file_mcenter_apps_user_pb_rpc_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mcenter_apps_user_pb_rpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mcenter_apps_user_pb_rpc_proto_goTypes,
		DependencyIndexes: file_mcenter_apps_user_pb_rpc_proto_depIdxs,
		EnumInfos:         file_mcenter_apps_user_pb_rpc_proto_enumTypes,
		MessageInfos:      file_mcenter_apps_user_pb_rpc_proto_msgTypes,
	}.Build()
	File_mcenter_apps_user_pb_rpc_proto = out.File
	file_mcenter_apps_user_pb_rpc_proto_rawDesc = nil
	file_mcenter_apps_user_pb_rpc_proto_goTypes = nil
	file_mcenter_apps_user_pb_rpc_proto_depIdxs = nil
}
