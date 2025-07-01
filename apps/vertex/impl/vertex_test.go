package impl_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/library/tools/format"
	"google.golang.org/genai"

	// _ "github.com/kade-chen/mcenter/apps" //registry all impl/api
	"github.com/kade-chen/mcenter/apps/vertex"
	// _ "github.com/kade-chen/mcenter/apps/vertex/impl"

	_ "github.com/kade-chen/mcenter/apps/vertex/provider/registry"
)

var (
	ctx  = context.Background()
	impl vertex.Service
)

func Test_Issue_Toke(t *testing.T) {
	// fmt.Println("-----", len(vertex.GET_GEMINI2_LOCATION_STRING))
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
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	req.ConfigFile.Path = "/Users/kade.chen/go-kade-project/github/mcenter/etc/config.toml"
	ioc.DevelopmentSetup(req)
	impl = ioc.Controller().Get(vertex.AppName).(vertex.Service)
}
