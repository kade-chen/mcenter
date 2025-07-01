package gemini25flashpreview

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

var _ provider.Issuer = (*Gemini_25_flash_preview)(nil)

func init() {
	provider.Registe(&Gemini_25_flash_preview{})
}

type Gemini_25_flash_preview struct {
	gemini_clients map[vertex.Gemini_25_FLASH_PREVIEW_LOCATION]*genai.Client
}

func (g *Gemini_25_flash_preview) GrantModel() vertex.GRANT_MODEL {
	return vertex.GRANT_MODEL_Gemini_2_5_Flash_Preview
}

func (g *Gemini_25_flash_preview) getCredentials(custom_ProjectId_bool bool, custom_ProjectId, custom_serviceAccount string) (project_id, credentials_file string) {
	if custom_ProjectId_bool {
		return custom_ProjectId, custom_serviceAccount
	}
	return ioc.Config().Get(provider.AppName).(*provider.Vertex).Default_Project_ID, ioc.Config().Get(provider.AppName).(*provider.Vertex).Default_Service_Account_Name
}

func (g *Gemini_25_flash_preview) Init() error {
	g.gemini_clients = make(map[vertex.Gemini_25_FLASH_PREVIEW_LOCATION]*genai.Client)
	//https://pkg.go.dev/cloud.google.com/go/auth/credentials
	var creds1 *auth.Credentials
	project_id, credentials_file := g.getCredentials(ioc.Config().Get(provider.AppName).(*provider.Vertex).Gemini_25_Flash_Preview_Project_ID_Bool, ioc.Config().Get(provider.AppName).(*provider.Vertex).Gemini_25_Flash_Preview_Project_ID, ioc.Config().Get(provider.AppName).(*provider.Vertex).Gemini_25_Flash_Preview_Service_Account_Name)
	creds1, _ = credentials.DetectDefault(&credentials.DetectOptions{
		Scopes: []string{"https://www.googleapis.com/auth/cloud-platform"},
		// CredentialsJSON: data,
		CredentialsFile: credentials_file,
	})
	for i := range make([]int, len(vertex.Gemini_25_FLASH_PREVIEW_String_To_LOCATION)) {
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
			Backend: genai.BackendVertexAI,
			Project: project_id,
			// Project:     ioc.Config().Get(provider.AppName).(*provider.Vertex).ServiceAccount_GEMINI_2_5_Flash_Preview,
			Location:    vertex.Gemini_25_FLASH_PREVIEW_LOCATION_To_String[vertex.Gemini_25_FLASH_PREVIEW_LOCATION(i)],
			Credentials: creds1,

			HTTPOptions: genai.HTTPOptions{APIVersion: "v1"},
		})

		if err != nil {
			fmt.Println("Error initializing client:--------", err)
			return err
		}
		// g.clientsv2 = client
		g.gemini_clients[vertex.Gemini_25_FLASH_PREVIEW_LOCATION(i)] = client
	}
	// if g.clientsv2 == nil { // 检查 g.clientsv2 是否为 nil
	// 	return exception.NewNotFound(("Error: clientsv2 is nil"))
	// }

	return nil
}
