package impl

import (
	"context"

	speech "cloud.google.com/go/speech/apiv1"
	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/library/ioc"
	logs "github.com/kade-chen/library/ioc/config/log"
	"github.com/kade-chen/mcenter/apps/stt"
	"github.com/rs/zerolog"
	"google.golang.org/api/option"
)

var _ stt.Service = (*speechToText)(nil)

func init() {
	ioc.Controller().Registry(&speechToText{})
}

type speechToText struct {
	ioc.ObjectImpl
	log    *zerolog.Logger
	client *speech.Client
	// service account directory path name
	ServiceAccount string `toml:"service_account" json:"service_account" yaml:"service_account" env:"service_account"`
}

func (t *speechToText) Name() string {
	return stt.AppNameV1
}

func (t *speechToText) Init() error {

	t.log = logs.Sub(stt.AppNameV1)

	t.log.Info().Msg("speech to text v1 client initializing...")
	client, err := speech.NewClient(context.Background(), option.WithCredentialsFile(t.ServiceAccount))
	if err != nil {
		t.log.Error().Msgf("speech to text client init failed , error: %v", err)
		return exception.NewUnauthorized("speech to text client init failed, error: %v", err)
	}

	t.client = client

	t.log.Info().Msg("speech to text v1 client successfully initlialized")

	// 初始化 PortAudio
	// err = portaudio.Initialize()
	// if err != nil {
	// 	log.Fatalf("Failed to initialize portaudio: %v", err)
	// }
	return nil
}
