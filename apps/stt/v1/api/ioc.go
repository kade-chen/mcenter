package api

import (
	"net/http"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/gorilla/websocket"
	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/library/ioc/config/gorestful"
	logs "github.com/kade-chen/library/ioc/config/log"
	"github.com/kade-chen/mcenter/apps/policy"
	"github.com/kade-chen/mcenter/apps/stt"
	"github.com/kade-chen/mcenter/middlewares"
	"github.com/rs/zerolog"
)

func init() {
	ioc.Api().Registry(&speechToTextHandler{})
}

type speechToTextHandler struct {
	ioc.ObjectImpl
	stt stt.Service
	log *zerolog.Logger
	// user_binding_roles *mongo.Collection
	// role               *mongo.Collection
	policy policy.Service
}

func (u *speechToTextHandler) Init() error {
	u.log = logs.Sub(stt.AppNameV1)

	// db := ioc_mongo.DB()
	// u.role = db.Collection("roles")
	u.stt = ioc.Controller().Get(stt.AppNameV1).(stt.Service)
	u.policy = ioc.Controller().Get(policy.AppName).(policy.Service)
	// u.user_binding_roles = db.Collection("user_binding_roles")
	u.Registry()
	return nil
}

func (u *speechToTextHandler) Name() string {
	return stt.AppNameV1
}

func (u *speechToTextHandler) Version() string {
	return "v1"
}

func (u *speechToTextHandler) Registry() {
	tags := []string{"Speech To Text V1 Client"}
	ws := gorestful.InitRouter(u)
	ws.Route(ws.GET("/ws").To(u.streamingRecognize).
		Doc("Websocket Streaming Recognize").
		Metadata(restfulspec.KeyOpenAPITags, tags). //标签
		Notes("Websocket Streaming Recognize").
		Filter(middlewares.NewTokenAuther().Auth_Login))

	// Writes(user.User{}).
	// Reads(user.CreateUserRequest{}).
	// Returns(200, "Ok", user.User{}).
	// )

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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 在这里可以根据请求的 Origin 进行判断，返回 true 表示允许连接
		// 返回 false 表示拒绝连接
		return true // 允许所有请求
	},
}
