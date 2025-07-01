package api

import (
	"net/http"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/library/ioc/config/gorestful"
	logs "github.com/kade-chen/library/ioc/config/log"
	"github.com/kade-chen/mcenter/apps/policy"
	"github.com/kade-chen/mcenter/apps/vertex"
	"github.com/kade-chen/mcenter/middlewares"
	"github.com/rs/zerolog"
	"google.golang.org/genai"
)

func init() {
	ioc.Api().Registry(&vertexHandler{})
}

type vertexHandler struct {
	ioc.ObjectImpl
	gemini vertex.Service
	log    *zerolog.Logger
	// user_binding_roles *mongo.Collection
	// role               *mongo.Collection
	policy policy.Service
}

func (g *vertexHandler) Name() string {
	return vertex.AppName
}

func (g *vertexHandler) Version() string {
	return "v1"
}

func (g *vertexHandler) Init() error {
	g.log = logs.Sub(g.Name())
	g.log.Debug().Msgf("The gemini application log was successfully initialized")
	// db := ioc_mongo.DB()
	// // u.role = db.Collection("roles")
	g.gemini = ioc.Controller().Get(vertex.AppName).(vertex.Service)
	g.policy = ioc.Controller().Get(policy.AppName).(policy.Service)
	// u.user_binding_roles = db.Collection("user_binding_roles")
	g.Registry()
	return nil
}

func (g *vertexHandler) Registry() {
	tags := []string{"Vertex Plotform"}
	ws := gorestful.InitRouter(g)
	ws.Route(ws.GET("/no-streaming").To(g.nostreamHandler).
		Doc("Vertex NoStreaming Output").
		Metadata(restfulspec.KeyOpenAPITags, tags). //标签
		Reads(vertex.Gemini_Config{}).
		Writes([]genai.Part{}).
		Returns(http.StatusOK, "ok", []genai.GenerateContentResponse{}).
		Notes("Vertex NoStreaming Output").
		Filter(middlewares.NewTokenAuther().Auth_Login))

	ws.Route(ws.GET("/streaming").To(g.streamHandler).
		Doc("Vertex Streaming Output").
		Metadata(restfulspec.KeyOpenAPITags, tags). //标签
		Reads(vertex.Gemini_Config{}).
		Writes([]genai.Part{}).
		Returns(http.StatusOK, "ok", []genai.GenerateContentResponse{}).
		Notes("Vertex Streaming Output").
		Filter(middlewares.NewTokenAuther().Auth_Login))
}
