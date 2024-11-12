// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: mcenter/apps/roles/pb/roles_userbindings.proto

package roles

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

type CreateUserBindingRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: bson:"_id" json:"username_id"
	UsernameId string `protobuf:"bytes,1,opt,name=username_id,json=usernameId,proto3" json:"username_id" bson:"_id"`
	// 分页参数
	// @gotags: bson:"role_id" json:"role_id"
	RoleId string `protobuf:"bytes,2,opt,name=role_id,json=roleId,proto3" json:"role_id" bson:"role_id"`
}

func (x *CreateUserBindingRoleRequest) Reset() {
	*x = CreateUserBindingRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_roles_pb_roles_userbindings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserBindingRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserBindingRoleRequest) ProtoMessage() {}

func (x *CreateUserBindingRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_roles_pb_roles_userbindings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserBindingRoleRequest.ProtoReflect.Descriptor instead.
func (*CreateUserBindingRoleRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_roles_pb_roles_userbindings_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserBindingRoleRequest) GetUsernameId() string {
	if x != nil {
		return x.UsernameId
	}
	return ""
}

func (x *CreateUserBindingRoleRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

type UserBindingRoles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: bson:"_id" json:"username_id"
	UsernameId string `protobuf:"bytes,1,opt,name=username_id,json=usernameId,proto3" json:"username_id" bson:"_id"`
	// 分页参数
	// @gotags: bson:"role_id" json:"role_id"
	RoleId []string `protobuf:"bytes,2,rep,name=role_id,json=roleId,proto3" json:"role_id" bson:"role_id"`
}

func (x *UserBindingRoles) Reset() {
	*x = UserBindingRoles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_roles_pb_roles_userbindings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBindingRoles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBindingRoles) ProtoMessage() {}

func (x *UserBindingRoles) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_roles_pb_roles_userbindings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBindingRoles.ProtoReflect.Descriptor instead.
func (*UserBindingRoles) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_roles_pb_roles_userbindings_proto_rawDescGZIP(), []int{1}
}

func (x *UserBindingRoles) GetUsernameId() string {
	if x != nil {
		return x.UsernameId
	}
	return ""
}

func (x *UserBindingRoles) GetRoleId() []string {
	if x != nil {
		return x.RoleId
	}
	return nil
}

var File_mcenter_apps_roles_pb_roles_userbindings_proto protoreflect.FileDescriptor

var file_mcenter_apps_roles_pb_roles_userbindings_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x67, 0x6f, 0x6b, 0x61, 0x64, 0x65, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x58, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64,
	0x22, 0x4c, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x6f, 0x6c, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x32, 0x74,
	0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x6d, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x32, 0x2e, 0x67, 0x6f, 0x6b, 0x61, 0x64,
	0x65, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x67,
	0x6f, 0x6b, 0x61, 0x64, 0x65, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x6f, 0x6c, 0x65, 0x73, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6b, 0x61, 0x64, 0x65, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mcenter_apps_roles_pb_roles_userbindings_proto_rawDescOnce sync.Once
	file_mcenter_apps_roles_pb_roles_userbindings_proto_rawDescData = file_mcenter_apps_roles_pb_roles_userbindings_proto_rawDesc
)

func file_mcenter_apps_roles_pb_roles_userbindings_proto_rawDescGZIP() []byte {
	file_mcenter_apps_roles_pb_roles_userbindings_proto_rawDescOnce.Do(func() {
		file_mcenter_apps_roles_pb_roles_userbindings_proto_rawDescData = protoimpl.X.CompressGZIP(file_mcenter_apps_roles_pb_roles_userbindings_proto_rawDescData)
	})
	return file_mcenter_apps_roles_pb_roles_userbindings_proto_rawDescData
}

var file_mcenter_apps_roles_pb_roles_userbindings_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mcenter_apps_roles_pb_roles_userbindings_proto_goTypes = []interface{}{
	(*CreateUserBindingRoleRequest)(nil), // 0: gokade.mcenter.roles.CreateUserBindingRoleRequest
	(*UserBindingRoles)(nil),             // 1: gokade.mcenter.roles.UserBindingRoles
}
var file_mcenter_apps_roles_pb_roles_userbindings_proto_depIdxs = []int32{
	0, // 0: gokade.mcenter.roles.RPC.UserBindingRole:input_type -> gokade.mcenter.roles.CreateUserBindingRoleRequest
	1, // 1: gokade.mcenter.roles.RPC.UserBindingRole:output_type -> gokade.mcenter.roles.UserBindingRoles
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mcenter_apps_roles_pb_roles_userbindings_proto_init() }
func file_mcenter_apps_roles_pb_roles_userbindings_proto_init() {
	if File_mcenter_apps_roles_pb_roles_userbindings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mcenter_apps_roles_pb_roles_userbindings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserBindingRoleRequest); i {
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
		file_mcenter_apps_roles_pb_roles_userbindings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBindingRoles); i {
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
			RawDescriptor: file_mcenter_apps_roles_pb_roles_userbindings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mcenter_apps_roles_pb_roles_userbindings_proto_goTypes,
		DependencyIndexes: file_mcenter_apps_roles_pb_roles_userbindings_proto_depIdxs,
		MessageInfos:      file_mcenter_apps_roles_pb_roles_userbindings_proto_msgTypes,
	}.Build()
	File_mcenter_apps_roles_pb_roles_userbindings_proto = out.File
	file_mcenter_apps_roles_pb_roles_userbindings_proto_rawDesc = nil
	file_mcenter_apps_roles_pb_roles_userbindings_proto_goTypes = nil
	file_mcenter_apps_roles_pb_roles_userbindings_proto_depIdxs = nil
}
