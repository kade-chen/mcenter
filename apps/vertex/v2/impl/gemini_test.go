package impl_test

import (
	"context"
	"testing"

	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/mcenter/apps/vertex"
	_ "github.com/kade-chen/mcenter/apps/vertex/v2/impl"
	// "google.golang.org/appengine/log"
)

var (
	ctx  = context.Background()
	impl vertex.ServiceGemini2
)

func TestNoStreamingGenerateContent(t *testing.T) {
	impl.NoStreamingGenerateContent(ctx)
	// a := format.ToJSON(resp)
	// fmt.Println(a, err)
}

func init() {
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	req.ConfigFile.Path = "/Users/kade.chen/go-kade-project/github/mcenter/etc/config.toml"
	ioc.DevelopmentSetup(req)
	impl = ioc.Controller().Get(vertex.AppName2).(vertex.ServiceGemini2)
}
