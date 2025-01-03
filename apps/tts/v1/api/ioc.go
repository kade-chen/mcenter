package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/library/ioc/config/gorestful"
	logs "github.com/kade-chen/library/ioc/config/log"
	"github.com/kade-chen/mcenter/apps/policy"
	"github.com/kade-chen/mcenter/apps/tts"
	"github.com/kade-chen/mcenter/middlewares"
	"github.com/rs/zerolog"
)

func init() {
	ioc.Api().Registry(&textToSpeechV1Handler{})
}

type textToSpeechV1Handler struct {
	ioc.ObjectImpl
	log    *zerolog.Logger
	ttsV1  tts.Service
	policy policy.Service
}

func (t *textToSpeechV1Handler) Init() error {
	// db := ioc_mongo.DB()
	t.log = logs.Sub(tts.AppNameV1)
	// u.role = db.Collection("roles")
	// u.user = ioc.Controller().Get(user.AppName).(user.Service)
	t.policy = ioc.Controller().Get(policy.AppName).(policy.Service)
	// u.user_binding_roles = db.Collection("user_binding_roles")
	t.ttsV1 = ioc.Controller().Get(tts.AppNameV1).(tts.Service)
	t.Registry()
	return nil
}

func (t *textToSpeechV1Handler) Name() string {
	return tts.AppNameV1
}

func (t *textToSpeechV1Handler) Version() string {
	return "v1"
}

func (t *textToSpeechV1Handler) Registry() {
	tags := []string{"Text To Speech V1 Client"}
	ws := gorestful.InitRouter(t)
	ws.Route(ws.POST("/").To(t.textToSpeech).
		Doc("Text To Speech").
		Metadata(restfulspec.KeyOpenAPITags, tags). //标签
		Notes("Text To Speech").
		Writes(tts.SynthesizeSpeechRequestBody{}).
		// Reads(tts.NewSynthesizeSpeechRequest().AudioConfig).
		// Returns(200, "Ok", texttospeechpb.SynthesizeSpeechRequest{}).
		Filter(middlewares.NewTokenAuther().Auth_Login))
}
