package gemini2flash_test

import (
	"context"
	"testing"

	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/mcenter/apps/vertex"
	"github.com/kade-chen/mcenter/apps/vertex/provider"
	_ "github.com/kade-chen/mcenter/apps/vertex/provider/all"
)

var (
	impl provider.ModelIssuer
	ctx  = context.Background()
)

func TestMain(t *testing.T) {
	impl.ModelIssue(ctx, nil)
}

func init() {
	// os.Setenv("DEBUG", "true") //this debug conflicts with another vs debug
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	req.ConfigFile.Path = "/Users/kade.chen/go-kade-project/github/mcenter/etc/config.toml"
	ioc.DevelopmentSetup(req)
	impl = provider.GetModelIssuer(vertex.GRANT_MODEL_GEMINI_2_0_Flash)
}
