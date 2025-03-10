// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.28.3
// source: mcenter/apps/token/pb/rpc.proto

package token

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

type DESCRIBY_BY int32

const (
	// 通过access token查看
	DESCRIBY_BY_ACCESS_TOKEN DESCRIBY_BY = 0
	// 通过刷新token查询
	DESCRIBY_BY_REFRESH_TOKEN DESCRIBY_BY = 1
)

// Enum value maps for DESCRIBY_BY.
var (
	DESCRIBY_BY_name = map[int32]string{
		0: "ACCESS_TOKEN",
		1: "REFRESH_TOKEN",
	}
	DESCRIBY_BY_value = map[string]int32{
		"ACCESS_TOKEN":  0,
		"REFRESH_TOKEN": 1,
	}
)

func (x DESCRIBY_BY) Enum() *DESCRIBY_BY {
	p := new(DESCRIBY_BY)
	*p = x
	return p
}

func (x DESCRIBY_BY) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DESCRIBY_BY) Descriptor() protoreflect.EnumDescriptor {
	return file_mcenter_apps_token_pb_rpc_proto_enumTypes[0].Descriptor()
}

func (DESCRIBY_BY) Type() protoreflect.EnumType {
	return &file_mcenter_apps_token_pb_rpc_proto_enumTypes[0]
}

func (x DESCRIBY_BY) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DESCRIBY_BY.Descriptor instead.
func (DESCRIBY_BY) EnumDescriptor() ([]byte, []int) {
	return file_mcenter_apps_token_pb_rpc_proto_rawDescGZIP(), []int{0}
}

type ValicateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 令牌
	//
	//	@gotags: json:"access_token"
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token"`
	// 令牌
	//
	//	@gotags: json:"access_token_name"
	ACCESS_TOKEN_NAME string `protobuf:"bytes,2,opt,name=ACCESS_TOKEN_NAME,json=ACCESSTOKENNAME,proto3" json:"access_token_name"`
}

func (x *ValicateTokenRequest) Reset() {
	*x = ValicateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_token_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValicateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValicateTokenRequest) ProtoMessage() {}

func (x *ValicateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_token_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValicateTokenRequest.ProtoReflect.Descriptor instead.
func (*ValicateTokenRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_token_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *ValicateTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *ValicateTokenRequest) GetACCESS_TOKEN_NAME() string {
	if x != nil {
		return x.ACCESS_TOKEN_NAME
	}
	return ""
}

// VerifyCodeRequest 验证码校验请求
type VerifyCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户名
	// @gotags: json:"username" validate:"required"
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username" validate:"required"`
	// 验证码
	// @gotags: json:"code" validate:"required"
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code" validate:"required"`
}

func (x *VerifyCodeRequest) Reset() {
	*x = VerifyCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_token_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyCodeRequest) ProtoMessage() {}

func (x *VerifyCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_token_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyCodeRequest.ProtoReflect.Descriptor instead.
func (*VerifyCodeRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_token_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyCodeRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *VerifyCodeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type RevolkTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 令牌
	// @gotags: json:"access_token"
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token"`
	// 刷新令牌
	// @gotags: json:"refresh_token"
	RefreshToken string `protobuf:"bytes,5,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token"`
	// 令牌名字
	// @gotags: json:"access_token_name"
	ACCESS_TOKEN_NAME string `protobuf:"bytes,2,opt,name=ACCESS_TOKEN_NAME,json=ACCESSTOKENNAME,proto3" json:"access_token_name"`
}

func (x *RevolkTokenRequest) Reset() {
	*x = RevolkTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_token_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevolkTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevolkTokenRequest) ProtoMessage() {}

func (x *RevolkTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_token_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevolkTokenRequest.ProtoReflect.Descriptor instead.
func (*RevolkTokenRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_token_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *RevolkTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *RevolkTokenRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *RevolkTokenRequest) GetACCESS_TOKEN_NAME() string {
	if x != nil {
		return x.ACCESS_TOKEN_NAME
	}
	return ""
}

var File_mcenter_apps_token_pb_rpc_proto protoreflect.FileDescriptor

var file_mcenter_apps_token_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x67, 0x6f, 0x6b, 0x61, 0x64, 0x65, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x21, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x70,
	0x62, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x14,
	0x56, 0x61, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x41, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x4e,
	0x41, 0x4d, 0x45, 0x22, 0x43, 0x0a, 0x11, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x88, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x76,
	0x6f, 0x6c, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x41, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x4e,
	0x41, 0x4d, 0x45, 0x2a, 0x32, 0x0a, 0x0b, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x59, 0x5f,
	0x42, 0x59, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x4f, 0x4b,
	0x45, 0x4e, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45, 0x46, 0x52, 0x45, 0x53, 0x48, 0x5f,
	0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x01, 0x32, 0xb2, 0x01, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12,
	0x58, 0x0a, 0x0d, 0x56, 0x61, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x2a, 0x2e, 0x67, 0x6f, 0x6b, 0x61, 0x64, 0x65, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67,
	0x6f, 0x6b, 0x61, 0x64, 0x65, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x51, 0x0a, 0x0a, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x27, 0x2e, 0x67, 0x6f, 0x6b, 0x61, 0x64, 0x65,
	0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x67, 0x6f, 0x6b, 0x61, 0x64, 0x65, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x29, 0x5a, 0x27,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x64, 0x65, 0x2d,
	0x63, 0x68, 0x65, 0x6e, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mcenter_apps_token_pb_rpc_proto_rawDescOnce sync.Once
	file_mcenter_apps_token_pb_rpc_proto_rawDescData = file_mcenter_apps_token_pb_rpc_proto_rawDesc
)

func file_mcenter_apps_token_pb_rpc_proto_rawDescGZIP() []byte {
	file_mcenter_apps_token_pb_rpc_proto_rawDescOnce.Do(func() {
		file_mcenter_apps_token_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_mcenter_apps_token_pb_rpc_proto_rawDescData)
	})
	return file_mcenter_apps_token_pb_rpc_proto_rawDescData
}

var file_mcenter_apps_token_pb_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mcenter_apps_token_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mcenter_apps_token_pb_rpc_proto_goTypes = []interface{}{
	(DESCRIBY_BY)(0),             // 0: gokade.mcenter.token.DESCRIBY_BY
	(*ValicateTokenRequest)(nil), // 1: gokade.mcenter.token.ValicateTokenRequest
	(*VerifyCodeRequest)(nil),    // 2: gokade.mcenter.token.VerifyCodeRequest
	(*RevolkTokenRequest)(nil),   // 3: gokade.mcenter.token.RevolkTokenRequest
	(*Token)(nil),                // 4: gokade.mcenter.token.Token
	(*Code)(nil),                 // 5: gokade.mcenter.token.Code
}
var file_mcenter_apps_token_pb_rpc_proto_depIdxs = []int32{
	1, // 0: gokade.mcenter.token.RPC.ValicateToken:input_type -> gokade.mcenter.token.ValicateTokenRequest
	2, // 1: gokade.mcenter.token.RPC.VerifyCode:input_type -> gokade.mcenter.token.VerifyCodeRequest
	4, // 2: gokade.mcenter.token.RPC.ValicateToken:output_type -> gokade.mcenter.token.Token
	5, // 3: gokade.mcenter.token.RPC.VerifyCode:output_type -> gokade.mcenter.token.Code
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mcenter_apps_token_pb_rpc_proto_init() }
func file_mcenter_apps_token_pb_rpc_proto_init() {
	if File_mcenter_apps_token_pb_rpc_proto != nil {
		return
	}
	file_mcenter_apps_token_pb_token_proto_init()
	file_mcenter_apps_token_pb_code_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mcenter_apps_token_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValicateTokenRequest); i {
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
		file_mcenter_apps_token_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyCodeRequest); i {
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
		file_mcenter_apps_token_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevolkTokenRequest); i {
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
			RawDescriptor: file_mcenter_apps_token_pb_rpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mcenter_apps_token_pb_rpc_proto_goTypes,
		DependencyIndexes: file_mcenter_apps_token_pb_rpc_proto_depIdxs,
		EnumInfos:         file_mcenter_apps_token_pb_rpc_proto_enumTypes,
		MessageInfos:      file_mcenter_apps_token_pb_rpc_proto_msgTypes,
	}.Build()
	File_mcenter_apps_token_pb_rpc_proto = out.File
	file_mcenter_apps_token_pb_rpc_proto_rawDesc = nil
	file_mcenter_apps_token_pb_rpc_proto_goTypes = nil
	file_mcenter_apps_token_pb_rpc_proto_depIdxs = nil
}
