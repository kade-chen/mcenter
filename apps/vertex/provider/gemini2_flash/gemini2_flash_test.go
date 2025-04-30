package gemini2flash_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/library/tools/format"
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
	ppp.ModelName = "gemini-2.0-flash-001"
	ppp.Contents = []*genai.Content{
		{
			Parts: []*genai.Part{
				{
					Text: "k8s是什么",
				},
			},
			Role: "user",
		},
	}
	a, err := impl.NoStreamingGenerateContent(ctx, ppp)
	if err != nil {
		panic(err)
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
