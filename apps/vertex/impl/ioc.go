package impl

import (
	"context"

	"cloud.google.com/go/vertexai/genai"
	"github.com/kade-chen/library/ioc"
	logs "github.com/kade-chen/library/ioc/config/log"
	"github.com/kade-chen/mcenter/apps/vertex"
	"github.com/rs/zerolog"
	"google.golang.org/api/option"
)

var _ vertex.Service = (*gemini)(nil)

func init() {
	ioc.Controller().Registry(&gemini{
		clients: make(map[vertex.LOCATION]*genai.Client, len(vertex.VertexLocationToString)*2),
	})
}

type gemini struct {
	ioc.ObjectImpl
	ProjectID string `toml:"project_id" json:"project_id" yaml:"project_id"  env:"project_id"`
	// service account directory path name
	ServiceAccount string `toml:"service_account" json:"service_account" yaml:"service_account" env:"service_account"`
	log            *zerolog.Logger
	clients        map[vertex.LOCATION]*genai.Client
}

func (g *gemini) Name() string {
	return vertex.AppName
}

func (g *gemini) Init() error {

	g.log = logs.Sub(vertex.AppName)

	// t.log.Info().Msg("speech to text v1 client initializing...")

	// t.log.Info().Msg("speech to text v1 client successfully initlialized")

	for i := 0; i < 2; i++ {
		// for i := 0; i < len(VertexLocationToString); i++ {
		client, err := genai.NewClient(context.Background(), g.ProjectID, vertex.GetLocationString(vertex.LOCATION(i)), option.WithCredentialsFile(g.ServiceAccount))
		if err != nil {
			g.log.Error().Err(err).Msgf("Failed to initialize client for region: %s", err)
			return err
		}
		g.clients[(vertex.LOCATION(i))] = client
		g.log.Info().Msgf("genai client for region: %s successfully initlialized", vertex.GetLocationString(vertex.LOCATION(i)))
	}

	return nil
}
