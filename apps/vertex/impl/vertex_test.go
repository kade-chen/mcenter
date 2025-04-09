package impl_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/kade-chen/library/ioc"
	// _ "github.com/kade-chen/mcenter/apps" //registry all impl/api
	"github.com/kade-chen/mcenter/apps/vertex"
	_ "github.com/kade-chen/mcenter/apps/vertex/impl"
	_ "github.com/kade-chen/mcenter/apps/vertex/provider/registry"
)

var (
	ctx  = context.Background()
	impl vertex.ServiceGemini2
)

func Test_Issue_Toke(t *testing.T) {
	fmt.Println("-----", len(vertex.GET_GEMINI2_LOCATION_STRING))
	impl.NoStreamingGenerateContent(ctx, nil)
}

func init() {
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	req.ConfigFile.Path = "/Users/kade.chen/go-kade-project/github/mcenter/etc/config.toml"
	ioc.DevelopmentSetup(req)
	impl = ioc.Controller().Get(vertex.AppName2).(vertex.ServiceGemini2)
}
