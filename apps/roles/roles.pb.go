// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.28.3
// source: mcenter/apps/roles/pb/roles.proto

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

type Roles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// roles _id
	// @gotags: bson:"_id" json:"Name"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name" bson:"_id"`
	// roles title
	// @gotags: json:"title"
	Title string `protobuf:"bytes,2,opt,name=Title,proto3" json:"title"`
	// roles stage
	// @gotags: json:"stage"
	Stage string `protobuf:"bytes,3,opt,name=Stage,proto3" json:"stage"`
	// roles description
	// @gotags: json:"description"
	Description string `protobuf:"bytes,4,opt,name=Description,proto3" json:"description"`
	// roles included permissions
	// @gotags: json:"included_permissions"
	IncludedPermissions []string `protobuf:"bytes,5,rep,name=IncludedPermissions,proto3" json:"included_permissions"`
	// roles deleted
	// @gotags: json:"deleted"
	Deleted bool `protobuf:"varint,6,opt,name=Deleted,proto3" json:"deleted"`
}

func (x *Roles) Reset() {
	*x = Roles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_roles_pb_roles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Roles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Roles) ProtoMessage() {}

func (x *Roles) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_roles_pb_roles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Roles.ProtoReflect.Descriptor instead.
func (*Roles) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_roles_pb_roles_proto_rawDescGZIP(), []int{0}
}

func (x *Roles) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Roles) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Roles) GetStage() string {
	if x != nil {
		return x.Stage
	}
	return ""
}

func (x *Roles) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Roles) GetIncludedPermissions() []string {
	if x != nil {
		return x.IncludedPermissions
	}
	return nil
}

func (x *Roles) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type CreateRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// roles _id
	// @gotags: bson:"_id" json:"Name"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name" bson:"_id"`
	// roles title
	// @gotags: json:"title"
	Title string `protobuf:"bytes,2,opt,name=Title,proto3" json:"title"`
	// roles stage
	// @gotags: json:"stage"
	Stage string `protobuf:"bytes,3,opt,name=Stage,proto3" json:"stage"`
	// roles description
	// @gotags: json:"description"
	Description string `protobuf:"bytes,4,opt,name=Description,proto3" json:"description"`
	// roles included permissions
	// @gotags: json:"included_permissions"
	IncludedPermissions []string `protobuf:"bytes,5,rep,name=IncludedPermissions,proto3" json:"included_permissions"`
	// roles deleted
	// @gotags: json:"deleted"
	Deleted bool `protobuf:"varint,6,opt,name=Deleted,proto3" json:"deleted"`
}

func (x *CreateRoleRequest) Reset() {
	*x = CreateRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_roles_pb_roles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleRequest) ProtoMessage() {}

func (x *CreateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_roles_pb_roles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleRequest.ProtoReflect.Descriptor instead.
func (*CreateRoleRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_roles_pb_roles_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRoleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRoleRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateRoleRequest) GetStage() string {
	if x != nil {
		return x.Stage
	}
	return ""
}

func (x *CreateRoleRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateRoleRequest) GetIncludedPermissions() []string {
	if x != nil {
		return x.IncludedPermissions
	}
	return nil
}

func (x *CreateRoleRequest) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

var File_mcenter_apps_roles_pb_roles_proto protoreflect.FileDescriptor

var file_mcenter_apps_roles_pb_roles_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x67, 0x6f, 0x6b, 0x61, 0x64, 0x65, 0x2e, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0xb5, 0x01, 0x0a, 0x05, 0x52, 0x6f,
	0x6c, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x13, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x13, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x22, 0xc1, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x13, 0x49, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x64, 0x65, 0x2d, 0x63, 0x68, 0x65, 0x6e, 0x2f, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mcenter_apps_roles_pb_roles_proto_rawDescOnce sync.Once
	file_mcenter_apps_roles_pb_roles_proto_rawDescData = file_mcenter_apps_roles_pb_roles_proto_rawDesc
)

func file_mcenter_apps_roles_pb_roles_proto_rawDescGZIP() []byte {
	file_mcenter_apps_roles_pb_roles_proto_rawDescOnce.Do(func() {
		file_mcenter_apps_roles_pb_roles_proto_rawDescData = protoimpl.X.CompressGZIP(file_mcenter_apps_roles_pb_roles_proto_rawDescData)
	})
	return file_mcenter_apps_roles_pb_roles_proto_rawDescData
}

var file_mcenter_apps_roles_pb_roles_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mcenter_apps_roles_pb_roles_proto_goTypes = []interface{}{
	(*Roles)(nil),             // 0: gokade.mcenter.roles.Roles
	(*CreateRoleRequest)(nil), // 1: gokade.mcenter.roles.CreateRoleRequest
}
var file_mcenter_apps_roles_pb_roles_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mcenter_apps_roles_pb_roles_proto_init() }
func file_mcenter_apps_roles_pb_roles_proto_init() {
	if File_mcenter_apps_roles_pb_roles_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mcenter_apps_roles_pb_roles_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Roles); i {
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
		file_mcenter_apps_roles_pb_roles_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleRequest); i {
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
			RawDescriptor: file_mcenter_apps_roles_pb_roles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mcenter_apps_roles_pb_roles_proto_goTypes,
		DependencyIndexes: file_mcenter_apps_roles_pb_roles_proto_depIdxs,
		MessageInfos:      file_mcenter_apps_roles_pb_roles_proto_msgTypes,
	}.Build()
	File_mcenter_apps_roles_pb_roles_proto = out.File
	file_mcenter_apps_roles_pb_roles_proto_rawDesc = nil
	file_mcenter_apps_roles_pb_roles_proto_goTypes = nil
	file_mcenter_apps_roles_pb_roles_proto_depIdxs = nil
}
