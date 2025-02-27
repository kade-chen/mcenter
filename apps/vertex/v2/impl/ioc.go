package impl

import (
	"context"
	"log"
	"os"

	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/library/ioc"
	logs "github.com/kade-chen/library/ioc/config/log"
	"github.com/kade-chen/mcenter/apps/vertex"
	"github.com/rs/zerolog"
	"golang.org/x/oauth2/google"
	gemini2 "google.golang.org/genai"
)

var _ vertex.ServiceGemini2 = (*gemini)(nil)

func init() {
	ioc.Controller().Registry(&gemini{})
}

type gemini struct {
	ioc.ObjectImpl
	ProjectID string `toml:"project_id" json:"project_id" yaml:"project_id"  env:"project_id"`
	// service account directory path name
	ServiceAccount string `toml:"service_account" json:"service_account" yaml:"service_account" env:"service_account"`
	log            *zerolog.Logger
	// clients        map[vertex.LOCATION]*genai.Client
	clientsv2 *gemini2.Client
}

func (g *gemini) Name() string {
	return vertex.AppName2
}

func (g *gemini) Init() error {
	g.log = logs.Sub(vertex.AppName2)

	data, err := os.ReadFile(g.ServiceAccount)
	if err != nil {
		return err
	}
	creds, err := google.CredentialsFromJSON(context.Background(), data, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		log.Fatal(err)
	}
	// t.log.Info().Msg("speech to text v1 client initializing...")
	// t.log.Info().Msg("speech to text v1 client successfully initlialized")
	client, err := gemini2.NewClient(context.Background(), &gemini2.ClientConfig{
		//目前不支持，另外这个apikey，不能backend，projectid,location,credentials都填，否则会报错
		//https://console.cloud.google.com/apis/credentials?project=kade-poc (Vertex AI API/Generative Language API)
		//https://console.developers.google.com/apis/api/generativelanguage.googleapis.com/overview?project=499036589398
		// APIKey:   "AIzaSyAWtMUuGHBtUCSH829TcHiCO21m_8w08lU",
		Backend:     gemini2.BackendVertexAI,
		Project:     g.ProjectID,
		Location:    "us-central1",
		Credentials: creds,
		// HTTPOptions: gemini2.HTTPOptions{APIVersion: "v1"},
	})
	if err != nil {
		return err
	}
	g.clientsv2 = client

	if g.clientsv2 == nil { // 检查 g.clientsv2 是否为 nil
		return exception.NewNotFound(("Error: clientsv2 is nil"))
	}
	// for i := 24; i < 26; i++ {
	// 	// for i := 0; i < len(VertexLocationToString); i++ {
	// 	client, err := genai.NewClient(context.Background(), g.ProjectID, vertex.GetLocationString(vertex.LOCATION(i)), option.WithCredentialsFile(g.ServiceAccount))
	// 	if err != nil {
	// 		g.log.Error().Err(err).Msgf("Failed to initialize client for region: %s", err)
	// 		return err
	// 	}
	// 	// g.clients[(vertex.LOCATION(i))] = client
	// 	// g.log.Info().Msgf("genai client for region: %s successfully initlialized", vertex.GetLocationString(vertex.LOCATION(i)))
	// }
	return nil
}
