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
	ioc.Api().Registry(&gemini2Handler{})
}

type gemini2Handler struct {
	ioc.ObjectImpl
	gemini2 vertex.ServiceGemini2
	log     *zerolog.Logger
	// user_binding_roles *mongo.Collection
	// role               *mongo.Collection
	policy policy.Service
}

func (g *gemini2Handler) Name() string {
	return vertex.AppName2
}

func (g *gemini2Handler) Version() string {
	return "v1"
}

func (g *gemini2Handler) Init() error {
	g.log = logs.Sub(g.Name())
	g.log.Debug().Msgf("The gemini application log was successfully initialized")
	// db := ioc_mongo.DB()
	// u.role = db.Collection("roles")
	g.gemini2 = ioc.Controller().Get(vertex.AppName2).(vertex.ServiceGemini2)
	g.policy = ioc.Controller().Get(policy.AppName).(policy.Service)
	// u.user_binding_roles = db.Collection("user_binding_roles")
	g.Registry()
	return nil
}

func (g *gemini2Handler) Registry() {
	tags := []string{"Gemini2.0"}
	ws := gorestful.InitRouter(g)
	ws.Route(ws.POST("/streaming").To(g.stream).
		Doc("Gemini2.0 Streaming Output").
		Metadata(restfulspec.KeyOpenAPITags, tags). //标签
		Reads(vertex.Gemini2Config{}).
		Writes(vertex.Gemini2Config{}).
		Returns(http.StatusOK, "ok", []genai.Content{}).
		Notes("Gemini2.0 Streaming Output").
		Filter(middlewares.NewTokenAuther().Auth_Login))

	ws.Route(ws.POST("/Nostreaming").To(g.noStream).
		Doc("Gemini NoStreaming Output").
		Metadata(restfulspec.KeyOpenAPITags, tags). //标签
		Reads(vertex.GenaiSetting{}).
		Writes([]genai.Part{}).
		Returns(http.StatusOK, "ok", []genai.Part{}).
		Notes("Gemini NoStreaming Output").
		Filter(middlewares.NewTokenAuther().Auth_Login))
}
