package api

import (
	"fmt"
	"net/http"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/library/http/response"
	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/library/ioc/config/gorestful"
	"github.com/kade-chen/mcenter/apps/vertex"
	"github.com/kade-chen/mcenter/middlewares"
	"google.golang.org/genai"
)

func init() {
	ioc.Api().Registry(&vertexHandler{})
}

type vertexHandler struct {
	ioc.ObjectImpl
	gemini vertex.ServiceGemini3
	// log    *zerolog.Logger
	// user_binding_roles *mongo.Collection
	// role               *mongo.Collection
	// policy policy.Service
}

func (g *vertexHandler) Name() string {
	return vertex.AppName
}

func (g *vertexHandler) Version() string {
	return "v1"
}

func (g *vertexHandler) Init() error {
	// g.log = logs.Sub(g.Name())
	// g.log.Debug().Msgf("The gemini application log was successfully initialized")
	// db := ioc_mongo.DB()
	// // u.role = db.Collection("roles")
	g.gemini = ioc.Controller().Get(vertex.AppName).(vertex.ServiceGemini3)
	// g.policy = ioc.Controller().Get(policy.AppName).(policy.Service)
	// u.user_binding_roles = db.Collection("user_binding_roles")
	g.Registry()
	return nil
}

func (g *vertexHandler) Registry() {
	tags := []string{"Gemini1111111"}
	ws := gorestful.InitRouter(g)
	ws.Route(ws.GET("/nostream").To(g.nostreamHandler).
		Doc("test Streaming Output").
		Metadata(restfulspec.KeyOpenAPITags, tags). //标签
		Reads(vertex.GenaiSetting{}).
		Writes([]genai.Part{}).
		Returns(http.StatusOK, "ok", []genai.Part{}).
		Notes("test Streaming Output").
		Filter(middlewares.NewTokenAuther().Auth_Login))

	ws.Route(ws.GET("/stream").To(g.streamHandler).
		Doc("test1 Streaming Output").
		Metadata(restfulspec.KeyOpenAPITags, tags). //标签
		Reads(vertex.GenaiSetting{}).
		Writes([]genai.Part{}).
		Returns(http.StatusOK, "ok", []genai.Part{}).
		Notes("test1 Streaming Output").
		Filter(middlewares.NewTokenAuther().Auth_Login))
}

// streaming generate content for api
func (g *vertexHandler) nostreamHandler(r *restful.Request, w *restful.Response) {
	ppp := vertex.NewGemini2Config()
	ppp.ModelName = "gemini-2.0-flash-001"
	ppp.Contents = []*genai.Content{
		{
			Parts: []*genai.Part{
				{
					Text: "k8s是什么",
				},
			},
			Role: "user",
		},
	}
	a, err := g.gemini.NoStreamingGenerateContent(r.Request.Context(), ppp)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, a)
}

func (g *vertexHandler) streamHandler(r *restful.Request, w *restful.Response) {
	ppp := vertex.NewGemini2Config()
	ppp.ModelName = "gemini-2.0-flash-001"
	ppp.Contents = []*genai.Content{
		{
			Parts: []*genai.Part{
				{
					Text: "k8s是什么",
				},
			},
			Role: "user",
		},
	}
	ass, err := g.gemini.StreamingGenerateContent(r.Request.Context(), ppp)

	if err != nil {
		response.Failed(w, err)
		return
	}

	for a, err := range ass {
		if err != nil {
			response.Failed(w, exception.NewInternalServerError("ERROR: %v", err.Error()))
			return
		}
		fmt.Fprint(w, a.Candidates[0].Content.Parts[0].Text)
		// fmt.Fprintf(os.Stdout, "%v", p) == fmt.Print(p)
		// 立即发送
		if f, ok := w.ResponseWriter.(http.Flusher); ok {
			f.Flush()
		}
	}
	// response.Success(w, result)
}
