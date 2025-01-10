package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/library/ioc/config/gorestful"
	logs "github.com/kade-chen/library/ioc/config/log"
	"github.com/kade-chen/mcenter/apps/policy"
	"github.com/kade-chen/mcenter/apps/stt"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

func init() {
	ioc.Api().Registry(&speechToTextV2Handler{})
}

type speechToTextV2Handler struct {
	ioc.ObjectImpl
	stt stt.ServiceV2
	log *zerolog.Logger
	// user_binding_roles *mongo.Collection
	// role               *mongo.Collection
	policy policy.Service
	grpc   *grpc.ClientConn
}

func (u *speechToTextV2Handler) Init() error {
	u.log = logs.Sub(stt.AppNameV2)

	// db := ioc_mongo.DB()
	// u.role = db.Collection("roles")
	u.stt = ioc.Controller().Get(stt.AppNameV2).(stt.ServiceV2)
	u.policy = ioc.Controller().Get(policy.AppName).(policy.Service)
	// u.user_binding_roles = db.Collection("user_binding_roles")

	grpcClientConn, err := u.getGrpcClientConn()
	if err != nil {
		return err
	}

	u.grpc = grpcClientConn

	u.Registry()
	return nil
}

func (u *speechToTextV2Handler) Name() string {
	return stt.AppNameV2
}

func (u *speechToTextV2Handler) Version() string {
	return "v1"
}

func (u *speechToTextV2Handler) getGrpcClientConn() (*grpc.ClientConn, error) {
	// 1. 加载服务账号 JSON 文件
	// credentialsFilePath := u.stt.GetServiceAccount()
	perRPCCreds, err := oauth.NewServiceAccountFromFile(u.stt.GetServiceAccount(), "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		u.log.Debug().Msgf("Failed to create credentials: %v\n", err)
		return nil, exception.NewIocApiRegisterFailed("Failed to create credentials: %v", err)
	}
	// 2. 创建 gRPC 客户端连接
	conn1, err := grpc.NewClient(
		"speech.googleapis.com:443",
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
		grpc.WithPerRPCCredentials(perRPCCreds),
	)
	if err != nil {
		u.log.Debug().Msgf("Failed to create gRPC connection: %v\n", err)
		return nil, exception.NewIocApiRegisterFailed("Failed to create gRPC connection: %v", err)
	}
	return conn1, nil
}

func (u *speechToTextV2Handler) Registry() {
	tags := []string{"Speech To Text V2 Client"}
	ws := gorestful.InitRouter(u)
	ws.Route(ws.GET("/ws").To(u.streamingRecognize).
		Doc("Websocket Streaming Recognize").
		Metadata(restfulspec.KeyOpenAPITags, tags). //标签
		Notes("Websocket Streaming Recognize"))
	// Filter(middlewares.NewTokenAuther().Auth_Login))

	//LocalStreamingRecognize
	// ws.Route(ws.GET("/").To(u.localStreamingRecognize).
	// 	Doc("create user").
	// 	Metadata(restfulspec.KeyOpenAPITags, tags). //标签
	// 	Notes("create user"))
	// Writes(user.User{}).
	// Reads(user.CreateUserRequest{}).
	// Returns(200, "Ok", user.User{}).
	// )
}
