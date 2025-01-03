// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.28.3
// source: mcenter/apps/domain/pb/wechat_work.proto

package domain

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

type WechatWorkConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 开启企业微信认证
	// @gotags: bson:"enabled" json:"enabled"
	Enabled bool `protobuf:"varint,6,opt,name=enabled,proto3" json:"enabled" bson:"enabled"`
	// 企业微信企业Id, Oauth2.0应用获取应用Token时需要
	// @gotags: bson:"corp_id" json:"corp_id" env:"WECHAT_WORK_CORP_ID"
	CorpId string `protobuf:"bytes,1,opt,name=corp_id,json=corpId,proto3" json:"corp_id" bson:"corp_id" env:"WECHAT_WORK_CORP_ID"`
	// 企业微信应用Id, Oauth2.0时 也叫client_id
	// @gotags: bson:"agent_id" json:"agent_id" env:"WECHAT_WORK_CLIENT_ID"
	AgentId string `protobuf:"bytes,2,opt,name=agent_id,json=agentId,proto3" json:"agent_id" bson:"agent_id" env:"WECHAT_WORK_CLIENT_ID"`
	// 企业微信应用凭证, Oauth2.0时 也叫client_secret
	// @gotags: bson:"app_secret" json:"app_secret" env:"WECHAT_WORK_CLIENT_SECRET"
	AppSecret string `protobuf:"bytes,3,opt,name=app_secret,json=appSecret,proto3" json:"app_secret" bson:"app_secret" env:"WECHAT_WORK_CLIENT_SECRET"`
	// Oauth2.0时, 应用服务地址页面
	// @gotags: bson:"redirect_uri" json:"redirect_uri" env:"WECHAT_WORK_REDIRECT_URI"
	RedirectUri string `protobuf:"bytes,4,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri" bson:"redirect_uri" env:"WECHAT_WORK_REDIRECT_URI"`
	// 企业微信应用凭证(缓存使用)
	// @gotags: bson:"access_token" json:"access_token"
	AccessToken *WechatWorkAccessToken `protobuf:"bytes,5,opt,name=access_token,json=accessToken,proto3" json:"access_token" bson:"access_token"`
}

func (x *WechatWorkConfig) Reset() {
	*x = WechatWorkConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_domain_pb_wechat_work_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WechatWorkConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WechatWorkConfig) ProtoMessage() {}

func (x *WechatWorkConfig) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_domain_pb_wechat_work_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WechatWorkConfig.ProtoReflect.Descriptor instead.
func (*WechatWorkConfig) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_domain_pb_wechat_work_proto_rawDescGZIP(), []int{0}
}

func (x *WechatWorkConfig) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *WechatWorkConfig) GetCorpId() string {
	if x != nil {
		return x.CorpId
	}
	return ""
}

func (x *WechatWorkConfig) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *WechatWorkConfig) GetAppSecret() string {
	if x != nil {
		return x.AppSecret
	}
	return ""
}

func (x *WechatWorkConfig) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

func (x *WechatWorkConfig) GetAccessToken() *WechatWorkAccessToken {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

type WechatWorkAccessToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 应用服务访问凭证
	// @gotags: bson:"access_token" json:"access_token" env:"WECHAT_WORK_ACCESS_TOKEN"
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token" bson:"access_token" env:"WECHAT_WORK_ACCESS_TOKEN"`
	// 凭证过期时间
	// @gotags: bson:"expires_in" json:"expires_in"
	ExpiresIn int64 `protobuf:"varint,2,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in" bson:"expires_in"`
}

func (x *WechatWorkAccessToken) Reset() {
	*x = WechatWorkAccessToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_domain_pb_wechat_work_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WechatWorkAccessToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WechatWorkAccessToken) ProtoMessage() {}

func (x *WechatWorkAccessToken) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_domain_pb_wechat_work_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WechatWorkAccessToken.ProtoReflect.Descriptor instead.
func (*WechatWorkAccessToken) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_domain_pb_wechat_work_proto_rawDescGZIP(), []int{1}
}

func (x *WechatWorkAccessToken) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *WechatWorkAccessToken) GetExpiresIn() int64 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

var File_mcenter_apps_domain_pb_wechat_work_proto protoreflect.FileDescriptor

var file_mcenter_apps_domain_pb_wechat_work_proto_rawDesc = []byte{
	0x0a, 0x28, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x5f,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x6b, 0x61,
	0x64, 0x65, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x22, 0xf3, 0x01, 0x0a, 0x10, 0x57, 0x65, 0x63, 0x68, 0x61, 0x74, 0x57, 0x6f, 0x72, 0x6b,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x63, 0x6f, 0x72, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x6f, 0x72, 0x70, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x5f, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f,
	0x75, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x55, 0x72, 0x69, 0x12, 0x4f, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67,
	0x6f, 0x6b, 0x61, 0x64, 0x65, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x57, 0x65, 0x63, 0x68, 0x61, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x59, 0x0a, 0x15, 0x57, 0x65, 0x63, 0x68, 0x61,
	0x74, 0x57, 0x6f, 0x72, 0x6b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x69,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x49, 0x6e, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x61, 0x64, 0x65, 0x2d, 0x63, 0x68, 0x65, 0x6e, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mcenter_apps_domain_pb_wechat_work_proto_rawDescOnce sync.Once
	file_mcenter_apps_domain_pb_wechat_work_proto_rawDescData = file_mcenter_apps_domain_pb_wechat_work_proto_rawDesc
)

func file_mcenter_apps_domain_pb_wechat_work_proto_rawDescGZIP() []byte {
	file_mcenter_apps_domain_pb_wechat_work_proto_rawDescOnce.Do(func() {
		file_mcenter_apps_domain_pb_wechat_work_proto_rawDescData = protoimpl.X.CompressGZIP(file_mcenter_apps_domain_pb_wechat_work_proto_rawDescData)
	})
	return file_mcenter_apps_domain_pb_wechat_work_proto_rawDescData
}

var file_mcenter_apps_domain_pb_wechat_work_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mcenter_apps_domain_pb_wechat_work_proto_goTypes = []interface{}{
	(*WechatWorkConfig)(nil),      // 0: gokade.mcenter.domain.WechatWorkConfig
	(*WechatWorkAccessToken)(nil), // 1: gokade.mcenter.domain.WechatWorkAccessToken
}
var file_mcenter_apps_domain_pb_wechat_work_proto_depIdxs = []int32{
	1, // 0: gokade.mcenter.domain.WechatWorkConfig.access_token:type_name -> gokade.mcenter.domain.WechatWorkAccessToken
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mcenter_apps_domain_pb_wechat_work_proto_init() }
func file_mcenter_apps_domain_pb_wechat_work_proto_init() {
	if File_mcenter_apps_domain_pb_wechat_work_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mcenter_apps_domain_pb_wechat_work_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WechatWorkConfig); i {
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
		file_mcenter_apps_domain_pb_wechat_work_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WechatWorkAccessToken); i {
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
			RawDescriptor: file_mcenter_apps_domain_pb_wechat_work_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mcenter_apps_domain_pb_wechat_work_proto_goTypes,
		DependencyIndexes: file_mcenter_apps_domain_pb_wechat_work_proto_depIdxs,
		MessageInfos:      file_mcenter_apps_domain_pb_wechat_work_proto_msgTypes,
	}.Build()
	File_mcenter_apps_domain_pb_wechat_work_proto = out.File
	file_mcenter_apps_domain_pb_wechat_work_proto_rawDesc = nil
	file_mcenter_apps_domain_pb_wechat_work_proto_goTypes = nil
	file_mcenter_apps_domain_pb_wechat_work_proto_depIdxs = nil
}
