package gemini2flash

import (
	"context"
	"fmt"

	"cloud.google.com/go/auth"
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

func (g *Gemini2Flash) getCredentials(custom_ProjectId_bool bool, custom_ProjectId, custom_serviceAccount string) (project_id, credentials_file string) {
	if custom_ProjectId_bool {
		return custom_ProjectId, custom_serviceAccount
	}
	return ioc.Config().Get(provider.AppName).(*provider.Vertex).Default_Project_ID, ioc.Config().Get(provider.AppName).(*provider.Vertex).Default_Service_Account_Name
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
	var creds1 *auth.Credentials
	project_id, credentials_file := g.getCredentials(ioc.Config().Get(provider.AppName).(*provider.Vertex).Gemini_25_Flash_Preview_Project_ID_Bool, ioc.Config().Get(provider.AppName).(*provider.Vertex).Gemini_25_Flash_Preview_Project_ID, ioc.Config().Get(provider.AppName).(*provider.Vertex).Gemini_25_Flash_Preview_Service_Account_Name)
	creds1, _ = credentials.DetectDefault(&credentials.DetectOptions{
		Scopes: []string{"https://www.googleapis.com/auth/cloud-platform"},
		// CredentialsJSON: data,
		CredentialsFile: credentials_file,
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
			Project:     project_id,
			Location:    vertex.Get_GEMINI2_LOCATION_STRING(vertex.GEMINI_2_LOCATION(i)),
			Credentials: creds1,

			HTTPOptions: genai.HTTPOptions{APIVersion: "v1"},
		})

		if err != nil {
			fmt.Println("Error initializing client1:--------")
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
