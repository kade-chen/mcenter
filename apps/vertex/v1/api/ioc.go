package api

import (
	"net/http"

	"cloud.google.com/go/vertexai/genai"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/library/ioc/config/gorestful"
	logs "github.com/kade-chen/library/ioc/config/log"
	"github.com/kade-chen/mcenter/apps/policy"
	"github.com/kade-chen/mcenter/apps/vertex"
	"github.com/kade-chen/mcenter/middlewares"
	"github.com/rs/zerolog"
)

func init() {
	ioc.Api().Registry(&geminiHandler{})
}

type geminiHandler struct {
	ioc.ObjectImpl
	gemini vertex.Service
	log    *zerolog.Logger
	// user_binding_roles *mongo.Collection
	// role               *mongo.Collection
	policy policy.Service
}

func (g *geminiHandler) Name() string {
	return vertex.AppName3
}

func (g *geminiHandler) Version() string {
	return "v1"
}

func (g *geminiHandler) Init() error {
	g.log = logs.Sub(g.Name())
	g.log.Debug().Msgf("The gemini application log was successfully initialized")
	// db := ioc_mongo.DB()
	// u.role = db.Collection("roles")
	g.gemini = ioc.Controller().Get(vertex.AppName3).(vertex.Service)
	g.policy = ioc.Controller().Get(policy.AppName).(policy.Service)
	// u.user_binding_roles = db.Collection("user_binding_roles")
	g.Registry()
	return nil
}

func (g *geminiHandler) Registry() {
	tags := []string{"Gemini"}
	ws := gorestful.InitRouter(g)
	ws.Route(ws.GET("/streaming").To(g.streamHandler).
		Doc("Gemini Streaming Output").
		Metadata(restfulspec.KeyOpenAPITags, tags). //标签
		Reads(vertex.GenaiSetting{}).
		Writes([]genai.Part{}).
		Returns(http.StatusOK, "ok", []genai.Part{}).
		Notes("Gemini Streaming Output").
		Filter(middlewares.NewTokenAuther().Auth_Login))

	ws.Route(ws.GET("/Nostreaming").To(g.noStreamHandler).
		Doc("Gemini NoStreaming Output").
		Metadata(restfulspec.KeyOpenAPITags, tags). //标签
		Reads(vertex.GenaiSetting{}).
		Writes([]genai.Part{}).
		Returns(http.StatusOK, "ok", []genai.Part{}).
		Notes("Gemini NoStreaming Output").
		Filter(middlewares.NewTokenAuther().Auth_Login))
}
