package gemini2flash

import (
	"context"
	"fmt"

	"cloud.google.com/go/auth/credentials"

	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/library/tools/format"
	"google.golang.org/genai"

	"github.com/kade-chen/mcenter/apps/vertex"
	"github.com/kade-chen/mcenter/apps/vertex/provider"
)

var _ provider.Issuer = (*Gemini2Flash)(nil)

func init() {
	provider.Registe(&Gemini2Flash{})
}

type Gemini2Flash struct {
	gemini_clients map[vertex.GEMINI2_LOCATION]*genai.Client
}

func (g *Gemini2Flash) GrantModel() vertex.GRANT_MODEL {
	return vertex.GRANT_MODEL_GEMINI_2_0_Flash
}

func (g *Gemini2Flash) Init() error {
	g.gemini_clients = make(map[vertex.GEMINI2_LOCATION]*genai.Client)
	// data, err := os.ReadFile(ioc.Config().Get(provider.AppName).(*provider.Vertex).ServiceAccount_GEMINI_2_0_Flash)
	// if err != nil {
	// 	return err
	// }
	// creds, err := google.CredentialsFromJSON(context.Background(), data, "https://www.googleapis.com/auth/cloud-platform")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _ = creds
	creds1, _ := credentials.DetectDefault(&credentials.DetectOptions{
		Scopes:          []string{"https://www.googleapis.com/auth/cloud-platform"},
		CredentialsFile: ioc.Config().Get(provider.AppName).(*provider.Vertex).ServiceAccount_GEMINI_2_0_Flash,
	})

	for i := range make([]int, len(vertex.GET_GEMINI2_LOCATION_STRING)) {
		fmt.Println(vertex.Get_GEMINI2_LOCATION_STRING(vertex.GEMINI2_LOCATION(i)))
		// _ = creds1
		// t.log.Info().Msg("speech to text v1 client initializing...")
		// t.log.Info().Msg("speech to text v1 client successfully initlialized")
		client, err := genai.NewClient(context.Background(), &genai.ClientConfig{
			//目前不支持，另外这个apikey，不能backend，projectid,location,credentials都填，否则会报错
			//https://console.cloud.google.com/apis/credentials?project=kade-poc (Vertex AI API/Generative Language API)
			//https://console.developers.google.com/apis/api/generativelanguage.googleapis.com/overview?project=499036589398
			// APIKey:   "AIzaSyAWtMUuGHBtUCSH829TcHiCO21m_8w08lU",
			Backend:     genai.BackendVertexAI,
			Project:     ioc.Config().Get(provider.AppName).(*provider.Vertex).PROJECT_GEMINI_2_0_Flash,
			Location:    vertex.Get_GEMINI2_LOCATION_STRING(vertex.GEMINI2_LOCATION(i)),
			Credentials: creds1,

			HTTPOptions: genai.HTTPOptions{APIVersion: "v1"},
		})

		if err != nil {
			fmt.Println("Error initializing client:--------")
			return err
		}
		// g.clientsv2 = client
		g.gemini_clients[vertex.GEMINI2_LOCATION(i)] = client
	}
	// if g.clientsv2 == nil { // 检查 g.clientsv2 是否为 nil
	// 	return exception.NewNotFound(("Error: clientsv2 is nil"))
	// }
	return nil
}

func (g *Gemini2Flash) ModelIssue(ctx context.Context, se *vertex.GenaiSetting, part ...genai.Part) *genai.GenerateContentResponse {

	clinet := g.gemini_clients[vertex.GEMINI2_LOCATION(7)]
	re, err := clinet.Models.GenerateContent(ctx, "gemini-2.0-flash-001", []*genai.Content{
		{
			Parts: []*genai.Part{
				{
					Text: "k8s是什么",
				},
			},
			Role: "user",
		},
	}, nil)
	fmt.Println(format.ToJSON(re))
	fmt.Println(format.ToJSON(err))
	return nil
}
