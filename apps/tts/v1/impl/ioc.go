package impl

import (
	"context"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/mcenter/apps/tts"
	"github.com/rs/zerolog"
	"google.golang.org/api/option"

	logs "github.com/kade-chen/library/ioc/config/log"
)

var _ tts.Service = (*textToSpeechV1)(nil)

func init() {
	ioc.Controller().Registry(&textToSpeechV1{})
}

type textToSpeechV1 struct {
	ioc.ObjectImpl
	log    *zerolog.Logger
	client *texttospeech.Client
	// service account directory path name
	ServiceAccount string `toml:"service_account" json:"service_account" yaml:"service_account" env:"service_account"`
}

func (t *textToSpeechV1) Name() string {
	return tts.AppNameV1
}

func (t *textToSpeechV1) Init() error {

	t.log = logs.Sub(tts.AppNameV1)

	t.log.Info().Msg("text to speech v1 client initializing...")
	client, err := texttospeech.NewClient(context.Background(), option.WithCredentialsFile(t.ServiceAccount))
	if err != nil {
		t.log.Error().Msgf("text to speech client init failed , error: %v", err)
		return exception.NewUnauthorized("text to speech client init failed, error: %v", err)
	}
	t.log.Info().Msg("text to speech v1 client successfully initlialized")

	t.client = client
	return nil
}
