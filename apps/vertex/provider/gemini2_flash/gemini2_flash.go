package gemini2flash

import (
	"context"
	"fmt"
	"iter"

	"cloud.google.com/go/auth/credentials"

	"github.com/kade-chen/library/ioc"
	"google.golang.org/genai"

	"github.com/kade-chen/mcenter/apps/vertex"
	"github.com/kade-chen/mcenter/apps/vertex/provider"
)

var _ provider.Issuer = (*Gemini2Flash)(nil)

func init() {
	provider.Registe(&Gemini2Flash{})
}

type Gemini2Flash struct {
	gemini_clients map[vertex.GEMINI_2_LOCATION]*genai.Client
}

func (g *Gemini2Flash) GrantModel() vertex.GRANT_MODEL {
	return vertex.GRANT_MODEL_GEMINI_2_0_Flash_001
}

func (g *Gemini2Flash) Init() error {
	g.gemini_clients = make(map[vertex.GEMINI_2_LOCATION]*genai.Client)
	// data, err := os.ReadFile(ioc.Config().Get(provider.AppName).(*provider.Vertex).ServiceAccount_GEMINI_2_0_Flash)
	// if err != nil {
	// 	return err
	// }
	// creds, err := google.CredentialsFromJSON(context.Background(), data, "https://www.googleapis.com/auth/cloud-platform")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _ = creds
	//https://pkg.go.dev/cloud.google.com/go/auth/credentials
	creds1, _ := credentials.DetectDefault(&credentials.DetectOptions{
		Scopes: []string{"https://www.googleapis.com/auth/cloud-platform"},
		// CredentialsJSON: data,
		CredentialsFile: ioc.Config().Get(provider.AppName).(*provider.Vertex).ServiceAccount_GEMINI_2_0_Flash,
	})

	for i := range make([]int, len(vertex.GET_GEMINI2_LOCATION_STRING)) {
		//echo init regional
		// fmt.Println(vertex.Get_GEMINI2_LOCATION_STRING(vertex.GEMINI_2_LOCATION(i)))
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
			Location:    vertex.Get_GEMINI2_LOCATION_STRING(vertex.GEMINI_2_LOCATION(i)),
			Credentials: creds1,

			HTTPOptions: genai.HTTPOptions{APIVersion: "v1"},
		})

		if err != nil {
			fmt.Println("Error initializing client:--------")
			return err
		}
		// g.clientsv2 = client
		g.gemini_clients[vertex.GEMINI_2_LOCATION(i)] = client
	}
	// if g.clientsv2 == nil { // 检查 g.clientsv2 是否为 nil
	// 	return exception.NewNotFound(("Error: clientsv2 is nil"))
	// }
	return nil
}

func (g *Gemini2Flash) NoStreamingGenerateContent(ctx context.Context, gemini2Config *vertex.Gemini2Config) (*genai.GenerateContentResponse, error) {
	// fmt.Println("----", gemini2Config.Gemini_2_Regional)
	gemini2Config.Gemini_2_Regional = vertex.GEMINI2_LOCATION_Global
	clinet := g.gemini_clients[gemini2Config.Gemini_2_Regional]
	// // fmt.Println(format.ToJSON(re))
	// fmt.Println(format.ToJSON(err))
	return clinet.Models.GenerateContent(ctx, gemini2Config.ModelName, gemini2Config.Contents, gemini2Config.GenerateContentConfig)
}

func (g *Gemini2Flash) StreamingGenerateContent(ctx context.Context, gemini2Config *vertex.Gemini2Config) iter.Seq2[*genai.GenerateContentResponse, error] {
	// fmt.Println("----", gemini2Config.Gemini_2_Regional)
	gemini2Config.Gemini_2_Regional = vertex.GEMINI2_LOCATION_Global
	clinet := g.gemini_clients[gemini2Config.Gemini_2_Regional]

	return clinet.Models.GenerateContentStream(ctx, gemini2Config.ModelName, gemini2Config.Contents, gemini2Config.GenerateContentConfig)
}
