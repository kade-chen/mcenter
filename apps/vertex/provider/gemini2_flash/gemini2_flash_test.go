package gemini2flash_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/library/tools/format"
	"github.com/kade-chen/library/tools/generics"
	"github.com/kade-chen/mcenter/apps/vertex"
	"github.com/kade-chen/mcenter/apps/vertex/provider" //ioc
	"google.golang.org/genai"
	// _ "github.com/kade-chen/mcenter/apps/vertex/provider/registry"
	// _ "github.com/kade-chen/mcenter/apps/vertex/provider/gemini2_flash"
)

var (
	impl provider.ModelIssuer
	ctx  = context.Background()
)

func TestMain(t *testing.T) {
	fmt.Println(ioc.Config().List())
	fmt.Println(provider.ProviderRegistry())
	ppp := vertex.NewGemini2Config()
	ppp.ModelName = "gemini-2.5-pro-preview-05-06"
	// ppp.ModelName = "gemini-2.5-flash-preview-04-17"
	// ppp.GenerateContentConfig.ThinkingConfig.IncludeThoughts = true
	// ppp.GenerateContentConfig.Seed = generics.Generics[int32](1)
	ppp.GenerateContentConfig.TopK = generics.Generics[float32](20.0)
	ppp.GenerateContentConfig.TopP = generics.Generics[float32](0.5)
	ppp.GenerateContentConfig.Temperature = generics.Generics[float32](0.9)
	// ppp.GenerateContentConfig.MaxOutputTokens = 65535

	// ppp.GenerateContentConfig.ThinkingConfig.ThinkingBudget = generics.Generics[int32](3000)
	// ppp.ModelName = "gemini-2.0-flash-001"
	//"gemini-2.5-pro-preview-05-06"
	// imgData, _ := os.ReadFile("/Users/kade.chen/Downloads/s_0003_src.mp3")
	ppp.Contents = []*genai.Content{
		// {
		// 	Parts: []*genai.Part{
		// 		// {
		// 		// 	InlineData: &genai.Blob{
		// 		// 		MIMEType: "audio/mpeg",
		// 		// 		Data:     imgData,
		// 		// 	},
		// 		// },
		// 		// {
		// 		// 	FileData: &genai.FileData{
		// 		// 		FileURI:  "gs://kadeccc/s_0003_src.mp3",
		// 		// 		MIMEType: "audio/mpeg",
		// 		// 	},
		// 		// },
		// 	},
		// 	Role: "user",
		// },
		{
			Parts: []*genai.Part{
				{
					Text: "简述下vertex ai ,genai studio 的关系和区别",
				},
			},
			Role: "user",
		},
	}
	a, err := impl.NoStreamingGenerateContent(ctx, ppp)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(format.ToJSON(a))
}

func init() {
	// os.Setenv("DEBUG", "true") //this debug conflicts with another vs debug
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	req.ConfigFile.Path = "/Users/kade.chen/go-kade-project/github/mcenter/etc/config.toml"
	ioc.DevelopmentSetup(req)
	if err := provider.Init(); err != nil {
		panic(err)
	}
	// err := registry.ProviderInit() //
	// if err != nil {
	// 	fmt.Println("1234", err)
	// 	// panic(err)
	// }
	impl = provider.GetModelIssuer(vertex.GRANT_MODEL_GEMINI_2_0_Flash_001)
}
