package impl

import (
	"context"

	speech "cloud.google.com/go/speech/apiv2"
	"cloud.google.com/go/speech/apiv2/speechpb"
	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/library/ioc/config/grpc"
	logs "github.com/kade-chen/library/ioc/config/log"
	"github.com/kade-chen/mcenter/apps/stt"
	"github.com/rs/zerolog"
	"google.golang.org/api/option"
)

var _ stt.ServiceV2 = (*speechToTextV2)(nil)

func init() {
	ioc.Controller().Registry(&speechToTextV2{
		clinets: make(map[stt.ENDPOINT]*speech.Client, len(stt.ENDPOINTS_ENDPOINT)),
	})
}

type speechToTextV2 struct {
	ioc.ObjectImpl
	log       *zerolog.Logger
	clinets   map[stt.ENDPOINT]*speech.Client
	ProjectID string `toml:"project_id" json:"project_id" yaml:"project_id"  env:"project_id"`
	// service account directory path name
	ServiceAccount string `toml:"service_account" json:"service_account" yaml:"service_account" env:"service_account"`
	// cc             speechpb.SpeechClient
	cc speechpb.UnimplementedSpeechServer
}

func (s *speechToTextV2) Name() string {
	return stt.AppNameV2
}

func (s *speechToTextV2) Init() error {
	speechpb.RegisterSpeechServer(grpc.Get().Server(), &s.cc)
	// a := speechpb.NewSpeechClient()
	s.log = logs.Sub(stt.AppNameV2)

	s.log.Info().Msg("speech to text v2 client initializing...")

	for i := 0; i < len(stt.ENDPOINTS_ENDPOINT); i++ {
		client, err := speech.NewClient(context.Background(), option.WithCredentialsFile(s.ServiceAccount), option.WithEndpoint(stt.ENDPOINTS_STRING[(stt.ENDPOINT(i))]))
		if err != nil {
			s.log.Error().Msgf("speech to text v2 client init failed , error: %v", err)
			return exception.NewUnauthorized("speech to text client v2 init failed, error: %v", err)
		}
		s.clinets[stt.ENDPOINT(i)] = client
		s.log.Info().Msgf("speech to text v2 client for region: %s successfully initlialized", stt.ENDPOINTS_STRING[(stt.ENDPOINT(i))])
	}
	s.log.Info().Msg("speech to text v2 client successfully initlialized")

	return nil
}
